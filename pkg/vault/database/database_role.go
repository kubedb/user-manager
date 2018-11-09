package database

import (
	"fmt"

	"encoding/json"

	"github.com/appscode/go/types"
	vaultapi "github.com/hashicorp/vault/api"
	api "github.com/kubedb/apimachinery/apis/authorization/v1alpha1"
	"github.com/kubevault/db-manager/pkg/vault/database/mongodb"
	"github.com/kubevault/db-manager/pkg/vault/database/mysql"
	"github.com/kubevault/db-manager/pkg/vault/database/postgres"
	vaultcs "github.com/kubevault/operator/pkg/vault"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	appcat "kmodules.xyz/custom-resources/apis/appcatalog/v1alpha1"
	appcat_cs "kmodules.xyz/custom-resources/client/clientset/versioned/typed/appcatalog/v1alpha1"
)

const (
	DefaultDatabasePath = "database"
)

type DatabaseRole struct {
	RoleInterface
	vaultClient *vaultapi.Client
	path        string
}

func NewDatabaseRoleForPostgres(kClient kubernetes.Interface, appClient appcat_cs.AppcatalogV1alpha1Interface, role *api.PostgresRole) (DatabaseRoleInterface, error) {
	vClient, err := vaultcs.NewClient(kClient, appClient, &appcat.AppReference{
		Name:      types.String(role.Spec.AuthManagerRef.Name),
		Namespace: types.String(role.Spec.AuthManagerRef.Namespace),
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	path, err := getDatabasePath(appClient, role.Spec.AuthManagerRef)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get database path")
	}
	d := &DatabaseRole{
		RoleInterface: postgres.NewPostgresRole(kClient, vClient, role, path),
		path:          path,
		vaultClient:   vClient,
	}
	return d, nil
}

func NewDatabaseRoleForMysql(kClient kubernetes.Interface, appClient appcat_cs.AppcatalogV1alpha1Interface, role *api.MySQLRole) (DatabaseRoleInterface, error) {
	vClient, err := vaultcs.NewClient(kClient, appClient, &appcat.AppReference{
		Name:      types.String(role.Spec.AuthManagerRef.Name),
		Namespace: types.String(role.Spec.AuthManagerRef.Namespace),
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	path, err := getDatabasePath(appClient, role.Spec.AuthManagerRef)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get database path")
	}

	d := &DatabaseRole{
		RoleInterface: mysql.NewMySQLRole(kClient, vClient, role, path),
		path:          path,
		vaultClient:   vClient,
	}
	return d, nil
}

func NewDatabaseRoleForMongodb(kClient kubernetes.Interface, appClient appcat_cs.AppcatalogV1alpha1Interface, role *api.MongoDBRole) (DatabaseRoleInterface, error) {
	vClient, err := vaultcs.NewClient(kClient, appClient, &appcat.AppReference{
		Name:      types.String(role.Spec.AuthManagerRef.Name),
		Namespace: types.String(role.Spec.AuthManagerRef.Namespace),
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	path, err := getDatabasePath(appClient, role.Spec.AuthManagerRef)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get database path")
	}

	d := &DatabaseRole{
		RoleInterface: mongodb.NewMongoDBRole(kClient, vClient, role, path),
		path:          path,
		vaultClient:   vClient,
	}
	return d, nil
}

// EnableDatabase enables database secret engine
// It first checks whether database is enabled or not
func (d *DatabaseRole) EnableDatabase() error {
	enabled, err := d.IsDatabaseEnabled()
	if err != nil {
		return err
	}

	if enabled {
		return nil
	}

	err = d.vaultClient.Sys().Mount(d.path, &vaultapi.MountInput{
		Type: "database",
	})
	if err != nil {
		return err
	}
	return nil
}

// IsDatabaseEnabled checks whether database is enabled or not
func (d *DatabaseRole) IsDatabaseEnabled() (bool, error) {
	mnt, err := d.vaultClient.Sys().ListMounts()
	if err != nil {
		return false, errors.Wrap(err, "failed to list mounted secrets engines")
	}

	mntPath := d.path + "/"
	for k := range mnt {
		if k == mntPath {
			return true, nil
		}
	}
	return false, nil
}

// https://www.vaultproject.io/api/secret/databases/index.html#delete-role
//
// DeleteRole deletes role
// It's safe to call multiple time. It doesn't give
// error even if respective role doesn't exist
func (d *DatabaseRole) DeleteRole(name string) error {
	path := fmt.Sprintf("/v1/%s/roles/%s", d.path, name)
	req := d.vaultClient.NewRequest("DELETE", path)

	_, err := d.vaultClient.RawRequest(req)
	if err != nil {
		return errors.Wrapf(err, "failed to delete database role %s", name)
	}
	return nil
}

// If database path does not exist, then use default database path
func getDatabasePath(c appcat_cs.AppcatalogV1alpha1Interface, ref api.AuthManagerRef) (string, error) {
	vApp, err := c.AppBindings(types.String(ref.Namespace)).Get(types.String(ref.Name), metav1.GetOptions{})
	if err != nil {
		return "", err
	}

	var cf struct {
		DatabasePath string `json:"database_Path,omitempty"`
	}

	if vApp.Spec.Parameters != nil {
		err := json.Unmarshal(vApp.Spec.Parameters.Raw, &cf)
		if err != nil {
			return "", err
		}
	}

	if cf.DatabasePath != "" {
		return cf.DatabasePath, nil
	}
	return DefaultDatabasePath, nil
}
