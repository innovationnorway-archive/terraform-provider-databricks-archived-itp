package databricks

import (
	"context"
	"fmt"
	"net/url"

	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/terraform/httpclient"
	"github.com/innovationnorway/go-azure/auth"
	"github.com/innovationnorway/go-databricks/clusters"
	"github.com/innovationnorway/go-databricks/databricks"
	"github.com/innovationnorway/go-databricks/groups"
	"github.com/innovationnorway/go-databricks/workspace"
	"github.com/innovationnorway/terraform-provider-databricks/version"
)

const TerraformProviderUserAgent = "terraform-provider-databricks"

type Config struct {
	Token            string
	Host             string
	Azure            *AzureConfig
	OrganizationID   string
	terraformVersion string
}

type AzureConfig struct {
	WorkspaceID      string
	ServicePrincipal *AzureServicePrincipalConfig
}

type AzureServicePrincipalConfig struct {
	ClientID     string
	ClientSecret string
	TenantID     string
	Environment  string
}

type Meta struct {
	Clusters    clusters.BaseClient
	Groups      groups.BaseClient
	Workspace   workspace.BaseClient
	StopContext context.Context
}

func (c *Config) Client() (*Meta, error) {
	u, err := url.Parse(c.Host)
	if err != nil {
		return nil, fmt.Errorf("unable to parse URL: %s", err)
	}

	if u.Scheme == "" {
		u.Scheme = "https"
	}

	if u.Path == "" {
		u.Path = databricks.DefaultBaseURI
	}

	authorizer, err := c.getAuthorizer()
	if err != nil {
		return nil, fmt.Errorf("unable to get auth: %s", err)
	}

	return c.createClients(u.String(), authorizer)
}

func (c *Config) createClients(baseURI string, authorizer autorest.Authorizer) (*Meta, error) {
	meta := Meta{}

	meta.Clusters = clusters.NewWithBaseURI(baseURI)
	configureClient(&meta.Clusters.Client, authorizer, c.terraformVersion)

	meta.Groups = groups.NewWithBaseURI(baseURI)
	configureClient(&meta.Groups.Client, authorizer, c.terraformVersion)

	meta.Workspace = workspace.NewWithBaseURI(baseURI)
	configureClient(&meta.Workspace.Client, authorizer, c.terraformVersion)

	return &meta, nil
}

func configureClient(client *autorest.Client, authorizer autorest.Authorizer, tfVersion string) {
	client.Authorizer = authorizer
	client.UserAgent = getUserAgent(tfVersion)
	client.ResponseInspector = databricks.WithError()
}

func (c *Config) getAuthorizer() (autorest.Authorizer, error) {
	if c.Azure != nil {
		if c.Azure.ServicePrincipal != nil {
			config := auth.Config{
				ClientID:     c.Azure.ServicePrincipal.ClientID,
				ClientSecret: c.Azure.ServicePrincipal.ClientSecret,
				TenantID:     c.Azure.ServicePrincipal.TenantID,
				Environment:  c.Azure.ServicePrincipal.Environment,
			}

			managementToken, err := auth.GetToken(config)
			if err != nil {
				return nil, err
			}

			config.Resource = databricks.AzureDatabricksApplicationID
			token, err := auth.GetToken(config)
			if err != nil {
				return nil, err
			}

			return databricks.NewTokenAuthorizerWithServicePrincipal(
				token.OAuthToken(),
				managementToken.OAuthToken(),
				c.Azure.WorkspaceID), nil
		}

		config := auth.Config{Resource: databricks.AzureDatabricksApplicationID}
		token, err := auth.GetToken(config)
		if err != nil {
			return nil, err
		}

		return databricks.NewTokenAuthorizerWithWorkspaceID(token.OAuthToken(), c.Azure.WorkspaceID), nil
	}

	if c.OrganizationID != "" {
		return databricks.NewTokenAuthorizerWithOrgID(c.Token, c.OrganizationID), nil
	}

	return databricks.NewTokenAuthorizer(c.Token), nil
}

func getUserAgent(terraformVersion string) string {
	terraformUserAgent := httpclient.TerraformUserAgent(terraformVersion)
	providerUserAgent := fmt.Sprintf("%s/%s", TerraformProviderUserAgent, version.ProviderVersion)
	return fmt.Sprintf("%s %s", terraformUserAgent, providerUserAgent)
}
