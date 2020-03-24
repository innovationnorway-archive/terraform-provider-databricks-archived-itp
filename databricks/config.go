package databricks

import (
	"context"
	"fmt"
	"net/url"

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

	return &Meta{
		Clusters: client,
	}, nil
}
