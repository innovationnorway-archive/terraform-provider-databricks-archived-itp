package databricks

import (
	"context"
	"fmt"
	"net/url"

	azureAuth "github.com/innovationnorway/go-azure/auth"
	"github.com/innovationnorway/go-databricks/auth"
	"github.com/innovationnorway/go-databricks/clusters"
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
	StopContext context.Context
}

func (c *Config) Client() (*Meta, error) {
	u, err := url.Parse(c.Host)
	if err != nil {
		return nil, fmt.Errorf("Error parsing host: %s", err)
	}

	if u.Scheme == "" {
		u.Scheme = "https"
	}

	if u.Path == "" {
		u.Path = clusters.DefaultBaseURI
	}

	client := clusters.New()
	client.BaseURI = u.String()

	if c.Token != "" {
		client.Authorizer = auth.NewTokenAuthorizer(c.Token)
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

		client.Authorizer = auth.NewAzureDatabricksAuthorizer(
			token.OAuthToken(), managementToken.OAuthToken(), c.Azure.WorkspaceID)
	}

	return &Meta{
		Clusters: client,
	}, nil
}
