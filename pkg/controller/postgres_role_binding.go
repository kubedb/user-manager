package controller

import (
	"fmt"
	"time"

	kutilcorev1 "github.com/appscode/kutil/core/v1"
	"github.com/appscode/kutil/tools/queue"
	"github.com/golang/glog"
	api "github.com/kubedb/user-manager/apis/authorization/v1alpha1"
	patchutil "github.com/kubedb/user-manager/client/clientset/versioned/typed/authorization/v1alpha1/util"
	"github.com/kubedb/user-manager/pkg/vault"
	"github.com/kubedb/user-manager/pkg/vault/database/postgres"
	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
)

const (
	PostgresRoleBindingFinalizer = "database.postgres.rolebinding"
)

const (
	PhaseSuccess           api.PostgresRoleBindingPhase = "Success"
	PhaseInit              api.PostgresRoleBindingPhase = "Init"
	PhaseGetCredential     api.PostgresRoleBindingPhase = "GetCredential"
	PhaseCreateSecret      api.PostgresRoleBindingPhase = "CreateSecret"
	PhaseCreateRole        api.PostgresRoleBindingPhase = "CreateRole"
	PhaseCreateRoleBinding api.PostgresRoleBindingPhase = "CreateRoleBinding"
)

func (c *UserManagerController) initPostgresRoleBindingWatcher() {
	c.postgresRoleBindingInformer = c.dbInformerFactory.Authorization().V1alpha1().PostgresRoleBindings().Informer()
	c.postgresRoleBindingQueue = queue.New(api.ResourceKindPostgresRoleBinding, c.MaxNumRequeues, c.NumThreads, c.runPostgresRoleBindingInjector)

	// TODO: add custom event handler?
	c.postgresRoleBindingInformer.AddEventHandler(queue.DefaultEventHandler(c.postgresRoleBindingQueue.GetQueue()))
	c.postgresRoleBindingLister = c.dbInformerFactory.Authorization().V1alpha1().PostgresRoleBindings().Lister()
}

func (c *UserManagerController) runPostgresRoleBindingInjector(key string) error {
	obj, exist, err := c.postgresRoleBindingInformer.GetIndexer().GetByKey(key)
	if err != nil {
		glog.Errorf("Fetching object with key %s from store failed with %v", key, err)
		return err
	}

	if !exist {
		glog.Warningf("PostgresRoleBinding %s does not exist anymore\n", key)

	} else {
		pgRoleBinding := obj.(*api.PostgresRoleBinding)

		glog.Infof("Sync/Add/Update for PostgresRoleBinding %s/%s\n", pgRoleBinding.Namespace, pgRoleBinding.Name)

		if pgRoleBinding.DeletionTimestamp != nil {
			if kutilcorev1.HasFinalizer(pgRoleBinding.ObjectMeta, PostgresRoleBindingFinalizer) {
				pg, err := postgres.NewPostgresRoleBinding(c.kubeClient, c.dbClient, pgRoleBinding)
				if err != nil {
					glog.Errorf("for postgresRoleBinding(%s/%s): %v", pgRoleBinding.Namespace, pgRoleBinding.Name, err)
				} else {
					err = pg.RevokeLease(pgRoleBinding.Status.Lease.ID)
					if err != nil {
						glog.Errorf("for postgresRoleBinding(%s/%s): %v", pgRoleBinding.Namespace, pgRoleBinding.Name, err)
					}
				}

				// remove finalizer
				_, _, err = patchutil.PatchPostgresRoleBinding(c.dbClient.AuthorizationV1alpha1(), pgRoleBinding, func(binding *api.PostgresRoleBinding) *api.PostgresRoleBinding {
					binding.ObjectMeta = kutilcorev1.RemoveFinalizer(binding.ObjectMeta, PostgresRoleBindingFinalizer)
					return binding
				})
				if err != nil {
					return errors.Wrapf(err, "failed to remove postgresRoleBinding finalizer for (%s/%s)", pgRoleBinding.Namespace, pgRoleBinding.Name)
				}
			}

		} else if !kutilcorev1.HasFinalizer(pgRoleBinding.ObjectMeta, PostgresRoleBindingFinalizer) {
			// Add finalizer
			_, _, err = patchutil.PatchPostgresRoleBinding(c.dbClient.AuthorizationV1alpha1(), pgRoleBinding, func(binding *api.PostgresRoleBinding) *api.PostgresRoleBinding {
				binding.ObjectMeta = kutilcorev1.AddFinalizer(binding.ObjectMeta, PostgresRoleBindingFinalizer)
				return binding
			})
			if err != nil {
				return errors.Wrapf(err, "failed to set postgresRoleBinding finalizer for (%s/%s)", pgRoleBinding.Namespace, pgRoleBinding.Name)
			}

		} else {
			err := c.reconcilePostgresRoleBinding(pgRoleBinding)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Will do:
//	For vault:
//	  - get postgres credential
//	  - create secret containing credential
//	  - create rbac role and role binding
//    - sync role binding
func (c *UserManagerController) reconcilePostgresRoleBinding(pgRoleBinding *api.PostgresRoleBinding) error {
	if pgRoleBinding.Status.ObservedGeneration == 0 { // initial stage
		var (
			cred *vault.DatabaseCredentials
		)

		status := pgRoleBinding.Status
		name := pgRoleBinding.Name
		namespace := pgRoleBinding.Namespace
		roleName := getPostgresRbacRoleName(name)
		roleBindingName := getPostgresRbacRoleBindingName(name)

		pgClient, err := postgres.NewPostgresRoleBinding(c.kubeClient, c.dbClient, pgRoleBinding)
		if err != nil {
			status.Conditions = []api.PostgresRoleBindingCondition{
				{
					Type:    "Available",
					Status:  corev1.ConditionFalse,
					Reason:  "Unknown",
					Message: err.Error(),
				},
			}

			err2 := c.updatedPostgresRoleBindingStatus(&status, pgRoleBinding)
			if err2 != nil {
				return errors.Wrapf(err2, "for postgresRoleBinding(%s/%s): failed to update status", namespace, name)
			}
			return errors.Wrapf(err, "for postgresRoleBinding(%s/%s)", namespace, name)
		}

		if status.Phase == "" || status.Phase == PhaseGetCredential || status.Phase == PhaseCreateSecret {
			status.Phase = PhaseGetCredential

			cred, err = pgClient.GetCredentials()
			if err != nil {
				status.Conditions = []api.PostgresRoleBindingCondition{
					{
						Type:    "Available",
						Status:  corev1.ConditionFalse,
						Reason:  "FailedToGetCredential",
						Message: err.Error(),
					},
				}

				err2 := c.updatedPostgresRoleBindingStatus(&status, pgRoleBinding)
				if err2 != nil {
					return errors.Wrapf(err2, "for postgresRoleBinding(%s/%s): failed to update status", namespace, name)
				}

				return errors.Wrapf(err, "for postgresRoleBinding(%s/%s)", namespace, name)
			}

			glog.Infof("for postgresRoleBinding(%s/%s): getting postgres credential is successful\n", namespace, name)

			// add lease info
			d := time.Duration(cred.LeaseDuration)
			status.Lease = api.LeaseData{
				ID:            cred.LeaseID,
				Duration:      cred.LeaseDuration,
				RenewDeadline: time.Now().Add(time.Second * d).Unix(),
			}

			// next phase
			status.Phase = PhaseCreateSecret
		}

		if status.Phase == PhaseCreateSecret {
			err = pgClient.CreateSecret(cred)
			if err != nil {
				err2 := pgClient.RevokeLease(cred.LeaseID)
				if err2 != nil {
					return errors.Wrapf(err2, "for postgresRoleBinding(%s/%s): failed to revoke lease", namespace, name)
				}

				status.Conditions = []api.PostgresRoleBindingCondition{
					{
						Type:    "Available",
						Status:  corev1.ConditionFalse,
						Reason:  "FailedToCreateSecret",
						Message: err.Error(),
					},
				}

				err2 = c.updatedPostgresRoleBindingStatus(&status, pgRoleBinding)
				if err2 != nil {
					return errors.Wrapf(err2, "for postgresRoleBinding(%s/%s): failed to update status", namespace, name)
				}

				return errors.Wrapf(err, "for postgresRoleBinding(%s/%s)", namespace, name)
			}
			glog.Infof("for postgresRoleBinding(%s/%s): creating secret(%s/%s) is successful\n", namespace, name, namespace, pgRoleBinding.Spec.Store.Secret)

			// next phase
			status.Phase = PhaseCreateRole
		}

		if status.Phase == PhaseCreateRole {
			err = pgClient.CreateRole(roleName, pgRoleBinding.Spec.Store.Secret)
			if err != nil {
				status.Conditions = []api.PostgresRoleBindingCondition{
					{
						Type:    "Available",
						Status:  corev1.ConditionFalse,
						Reason:  "FailedToCreateRole",
						Message: err.Error(),
					},
				}

				err2 := c.updatedPostgresRoleBindingStatus(&status, pgRoleBinding)
				if err2 != nil {
					return errors.Wrapf(err2, "for postgresRoleBinding(%s/%s): failed to update status", namespace, name)
				}

				return errors.Wrapf(err, "for postgresRoleBinding(%s/%s)", namespace, name)
			}
			glog.Infof("for postgresRoleBinding(%s/%s): creating rbac role(%s/%s) is successful\n", namespace, name, namespace, roleName)

			//next phase
			status.Phase = PhaseCreateRoleBinding
		}

		if status.Phase == PhaseCreateRoleBinding {
			err = pgClient.CreateRoleBinding(roleBindingName, roleName)
			if err != nil {
				status.Conditions = []api.PostgresRoleBindingCondition{
					{
						Type:    "Available",
						Status:  corev1.ConditionFalse,
						Reason:  "FailedToCreateRoleBinding",
						Message: err.Error(),
					},
				}

				err2 := c.updatedPostgresRoleBindingStatus(&status, pgRoleBinding)
				if err2 != nil {
					return errors.Wrapf(err2, "for postgresRoleBinding(%s/%s): failed to update status", namespace, name)
				}

				return errors.Wrapf(err, "for postgresRoleBinding(%s/%s)", namespace, name)
			}
			glog.Infof("for postgresRoleBinding(%s/%s): creating rbac role binding(%s/%s) is successful\n", namespace, name, namespace, roleBindingName)
		}

		status.Phase = PhaseSuccess
		status.Conditions = []api.PostgresRoleBindingCondition{}
		status.ObservedGeneration = pgRoleBinding.GetGeneration()

		err = c.updatedPostgresRoleBindingStatus(&status, pgRoleBinding)
		if err != nil {
			return errors.Wrapf(err, "for postgresRoleBinding(%s/%s): failed to update status", namespace, name)
		}

	} else {
		// sync role binding
		// - update role binding
		if pgRoleBinding.ObjectMeta.Generation > pgRoleBinding.Status.ObservedGeneration {
			name := pgRoleBinding.Name
			namespace := pgRoleBinding.Namespace

			pgClient, err := postgres.NewPostgresRoleBinding(c.kubeClient, c.dbClient, pgRoleBinding)
			if err != nil {
				return errors.Wrapf(err, "for postgresRoleBinding(%s/%s)", namespace, name)
			}

			err = pgClient.UpdateRoleBinding(getPostgresRbacRoleBindingName(name), namespace)
			if err != nil {
				return errors.Wrapf(err, "for postgresRoleBinding(%s/%s)", namespace, name)
			}

			status := pgRoleBinding.Status
			status.ObservedGeneration = pgRoleBinding.ObjectMeta.Generation

			err = c.updatedPostgresRoleBindingStatus(&status, pgRoleBinding)
			if err != nil {
				return errors.Wrapf(err, "for postgresRoleBinding(%s/%s)", namespace, name)
			}
		}
	}

	return nil
}

func (c *UserManagerController) updatedPostgresRoleBindingStatus(status *api.PostgresRoleBindingStatus, pgRoleBinding *api.PostgresRoleBinding) error {
	_, err := patchutil.UpdatePostgresRoleBindingStatus(c.dbClient.AuthorizationV1alpha1(), pgRoleBinding, func(s *api.PostgresRoleBindingStatus) *api.PostgresRoleBindingStatus {
		s = status
		return s
	})
	if err != nil {
		return err
	}

	return nil
}

func getPostgresRbacRoleName(name string) string {
	return fmt.Sprintf("postgresrolebinding-%s-credential-reader", name)
}

func getPostgresRbacRoleBindingName(name string) string {
	return fmt.Sprintf("postgresrolebinding-%s-credential-reader", name)
}
