package databricks

import (
	"context"
	"fmt"
	"net/url"

	"github.com/Azure/go-autorest/autorest"
	azureAuth "github.com/innovationnorway/go-azure/auth"
	"github.com/innovationnorway/go-databricks/auth"
	"github.com/innovationnorway/go-databricks/clusters"
	"github.com/innovationnorway/go-databricks/groups"
)

type Config struct {
	Token string
	Host  string
	Azure *AzureConfig
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
		u.Path = clusters.DefaultBaseURI
	}

	authorizer, err := c.getAuthorizer()
	if err != nil {
		return nil, fmt.Errorf("unable to get auth: %s", err)
	}

	return c.configureClients(u.String(), authorizer)
}

func (c *Config) configureClients(baseURI string, authorizer autorest.Authorizer) (*Meta, error) {
	meta := Meta{}

	meta.Clusters = clusters.NewWithBaseURI(baseURI)
	meta.Clusters.Authorizer = authorizer

	meta.Groups = groups.NewWithBaseURI(baseURI)
	meta.Groups.Authorizer = authorizer

	return &meta, nil
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
