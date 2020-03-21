package databricks

import (
	"fmt"
	"net/url"

	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/Azure/go-autorest/autorest/azure/cli"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	cleanhttp "github.com/hashicorp/go-cleanhttp"
	"github.com/hashicorp/terraform/helper/logging"
	"github.com/innovationnorway/go-databricks/plumbing"
	"github.com/innovationnorway/go-databricks/porcelain"
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
	Databricks *plumbing.Databricks
	AuthInfo   runtime.ClientAuthInfoWriter
}

func (c *Config) Client() (interface{}, error) {
	u, err := url.Parse(c.Host)
	if err != nil {
		return nil, fmt.Errorf("Error parsing host: %s", err)
	}

	if u.Scheme == "" {
		u.Scheme = "https"
	}

	if u.Path == "" {
		u.Path = plumbing.DefaultBasePath
	}

	transport := httptransport.New(u.Host, u.Path, []string{u.Scheme})
	transport.Transport = logging.NewTransport("Databricks", cleanhttp.DefaultTransport())
	transport.Transport = porcelain.NewUserAgentTransport(transport.Transport, "Terraform")

	authInfo := httptransport.BearerToken(c.Token)

	if c.Azure != nil {
		authInfo, err = getAzureAuth(c.Azure)
		if err != nil {
			return nil, err
		}
	}

	return &Meta{
		Databricks: plumbing.New(transport, strfmt.Default),
		AuthInfo:   authInfo,
	}, nil
}

func getAzureAuth(c *AzureConfig) (runtime.ClientAuthInfoWriter, error) {
	var authInfo runtime.ClientAuthInfoWriter
	if c.ServicePrincipal != nil && c.ServicePrincipal.ClientSecret != "" {
		authConfig := auth.NewClientCredentialsConfig(
			c.ServicePrincipal.ClientID,
			c.ServicePrincipal.ClientSecret,
			c.ServicePrincipal.TenantID)
		managementToken, err := authConfig.ServicePrincipalToken()
		if err != nil {
			return nil, err
		}

		authConfig.Resource = porcelain.AzureDatabricksApplicationID
		token, err := authConfig.ServicePrincipalToken()
		if err != nil {
			return nil, err
		}

		authInfo = porcelain.AzureAuth(token.OAuthToken(), managementToken.OAuthToken(), c.WorkspaceID)
	} else {
		managementToken, err := cli.GetTokenFromCLI(azure.PublicCloud.ResourceManagerEndpoint)
		if err != nil {
			return nil, err
		}

		token, err := cli.GetTokenFromCLI(porcelain.AzureDatabricksApplicationID)
		if err != nil {
			return nil, err
		}

		authInfo = porcelain.AzureAuth(token.AccessToken, managementToken.AccessToken, c.WorkspaceID)
	}

	return authInfo, nil
}
