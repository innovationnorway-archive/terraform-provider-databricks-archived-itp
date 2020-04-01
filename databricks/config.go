package databricks

import (
	"context"
	"fmt"
	"net/url"

	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/terraform/httpclient"
	azureAuth "github.com/innovationnorway/go-azure/auth"
	"github.com/innovationnorway/go-databricks/auth"
	"github.com/innovationnorway/go-databricks/clusters"
	"github.com/innovationnorway/go-databricks/databricks"
	"github.com/innovationnorway/go-databricks/groups"
	"github.com/innovationnorway/go-databricks/workspace"
	"github.com/innovationnorway/terraform-provider-databricks/version"
)

const TerraformProviderUserAgent = "terraform-provider-databricks"

type Config struct {
	Token string
	Host  string
	Azure *AzureConfig

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
	var authorizer autorest.Authorizer

	if c.Token != "" {
		authorizer = auth.NewTokenAuthorizer(c.Token)
	}

	if c.Azure != nil {
		config := azureAuth.Config{}

		if c.Azure.ServicePrincipal != nil {
			config.ClientID = c.Azure.ServicePrincipal.ClientID
			config.ClientSecret = c.Azure.ServicePrincipal.ClientSecret
			config.TenantID = c.Azure.ServicePrincipal.TenantID
			config.Environment = c.Azure.ServicePrincipal.Environment
		}

		managementToken, err := azureAuth.GetToken(config)
		if err != nil {
			return nil, err
		}

		config.Resource = auth.AzureDatabricksApplicationID
		token, err := azureAuth.GetToken(config)
		if err != nil {
			return nil, err
		}

		authorizer = auth.NewAzureDatabricksAuthorizer(
			token.OAuthToken(),
			managementToken.OAuthToken(),
			c.Azure.WorkspaceID)
	}

	return authorizer, nil
}

func getUserAgent(terraformVersion string) string {
	terraformUserAgent := httpclient.TerraformUserAgent(terraformVersion)
	providerUserAgent := fmt.Sprintf("%s/%s", TerraformProviderUserAgent, version.ProviderVersion)
	return fmt.Sprintf("%s %s", terraformUserAgent, providerUserAgent)
}
