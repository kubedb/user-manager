package postgres

import (
	"encoding/json"
	"fmt"

	vaultapi "github.com/hashicorp/vault/api"
	api "github.com/kubedb/apimachinery/apis/authorization/v1alpha1"
	configapi "github.com/kubedb/apimachinery/apis/config/v1alpha1"
	"github.com/pkg/errors"
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	appcat_cs "kmodules.xyz/custom-resources/client/clientset/versioned/typed/appcatalog/v1alpha1"
)

type PostgresRole struct {
	config       *configapi.PostgresConfiguration
	secret       *core.Secret
	pgRole       *api.PostgresRole
	vaultClient  *vaultapi.Client
	kubeClient   kubernetes.Interface
	databasePath string
}

func NewPostgresRole(kClient kubernetes.Interface, appClient appcat_cs.AppcatalogV1alpha1Interface, v *vaultapi.Client, pgRole *api.PostgresRole, databasePath string) (*PostgresRole, error) {
	ref := pgRole.Spec.DatabaseRef
	dApp, err := appClient.AppBindings(ref.Namespace).Get(ref.Name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	secretRef := dApp.Spec.Secret
	if secretRef == nil {
		return nil, errors.New("database secret is not provided")
	}

	sr, err := kClient.CoreV1().Secrets(pgRole.Namespace).Get(secretRef.Name, metav1.GetOptions{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to get database secret")
	}

	cf := &configapi.PostgresConfiguration{}
	if dApp.Spec.Parameters != nil {
		err := json.Unmarshal(dApp.Spec.Parameters.Raw, cf)
		if err != nil {
			return nil, errors.Wrap(err, "failed to unmarshal database parameter")
		}
	}
	cf.SetDefaults()

	return &PostgresRole{
		config:       cf,
		secret:       sr,
		pgRole:       pgRole,
		vaultClient:  v,
		kubeClient:   kClient,
		databasePath: databasePath,
	}, nil
}

// https://www.vaultproject.io/api/secret/databases/index.html#configure-connection
// https://www.vaultproject.io/api/secret/databases/postgresql.html#configure-connection
//
// CreateConfig creates database configuration
func (p *PostgresRole) CreateConfig() error {
	if p.config == nil {
		return errors.New("database config is nil")
	}
	if p.secret == nil {
		return errors.New("database config is nil")
	}

	dRef := p.pgRole.Spec.DatabaseRef
	path := fmt.Sprintf("/v1/%s/config/%s", p.databasePath, dRef.Name)
	req := p.vaultClient.NewRequest("POST", path)
	payload := map[string]interface{}{
		"plugin_name":   p.config.PluginName,
		"allowed_roles": p.config.AllowedRoles,
	}

	data := p.secret.Data
	if val, ok := data["username"]; ok {
		payload["username"] = string(val)
	}
	if val, ok := data["password"]; ok {
		payload["password"] = string(val)
	}
	// TODO: get connection url from config parameters
	if val, ok := data["connection_url"]; ok {
		payload["connection_url"] = string(val)
	}

	if p.config.MaxOpenConnections > 0 {
		payload["max_open_connections"] = p.config.MaxOpenConnections
	}
	if p.config.MaxIdleConnections > 0 {
		payload["max_idle_connections"] = p.config.MaxIdleConnections
	}
	if p.config.MaxConnectionLifetime != "" {
		payload["max_connection_lifetime"] = p.config.MaxConnectionLifetime
	}

	err := req.SetJSONBody(payload)
	if err != nil {
		return errors.WithStack(err)
	}
	_, err = p.vaultClient.RawRequest(req)
	return err
}

// https://www.vaultproject.io/api/secret/databases/index.html#create-role
//
// CreateRole creates role
func (p *PostgresRole) CreateRole() error {
	name := p.pgRole.Name
	pg := p.pgRole.Spec

	path := fmt.Sprintf("/v1/%s/roles/%s", p.databasePath, name)
	req := p.vaultClient.NewRequest("POST", path)

	payload := map[string]interface{}{
		"db_name":             pg.DatabaseRef.Name,
		"creation_statements": pg.CreationStatements,
	}

	if len(pg.RevocationStatements) > 0 {
		payload["revocation_statements"] = pg.RevocationStatements
	}
	if len(pg.RollbackStatements) > 0 {
		payload["rollback_statements"] = pg.RollbackStatements
	}
	if len(pg.RenewStatements) > 0 {
		payload["renew_statements"] = pg.RenewStatements
	}
	if pg.DefaultTTL != "" {
		payload["default_ttl"] = pg.DefaultTTL
	}
	if pg.MaxTTL != "" {
		payload["max_ttl"] = pg.MaxTTL
	}

	err := req.SetJSONBody(payload)
	if err != nil {
		return errors.WithStack(err)
	}

	_, err = p.vaultClient.RawRequest(req)
	if err != nil {
		return errors.Wrapf(err, "failed to create database role %s for config %s", name, pg.DBName)
	}

	return nil
}
