package databricks

import (
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("DATABRICKS_HOST", nil),
			},

			"token": {
				Type:         schema.TypeString,
				Optional:     true,
				DefaultFunc:  schema.EnvDefaultFunc("DATABRICKS_TOKEN", nil),
				ExactlyOneOf: []string{"token", "azure"},
			},

			"azure": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"workspace_id": {
							Type:        schema.TypeString,
							Required:    true,
							DefaultFunc: schema.EnvDefaultFunc("DATABRICKS_AZURE_WORKSPACE_ID", nil),
						},

						"service_principal": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"client_id": {
										Type:        schema.TypeString,
										Required:    true,
										DefaultFunc: schema.EnvDefaultFunc("DATABRICKS_AZURE_CLIENT_ID", nil),
									},

									"client_secret": {
										Type:        schema.TypeString,
										Required:    true,
										Sensitive:   true,
										DefaultFunc: schema.EnvDefaultFunc("DATABRICKS_AZURE_CLIENT_SECRET", nil),
									},

									"tenant_id": {
										Type:        schema.TypeString,
										Required:    true,
										DefaultFunc: schema.EnvDefaultFunc("DATABRICKS_AZURE_TENANT_ID", nil),
									},

									"environment": {
										Type:        schema.TypeString,
										Required:    true,
										DefaultFunc: schema.EnvDefaultFunc("DATABRICKS_AZURE_ENVIRONMENT", azure.PublicCloud.Name),
									},
								},
							},
						},
					},
				},
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"databricks_cluster": dataSourceDatabricksCluster(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"databricks_cluster": resourceDatabricksCluster(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		Token: d.Get("token").(string),
		Host:  d.Get("host").(string),
	}

	if v, ok := d.GetOk("azure"); ok {
		config.Azure = expandAzureConfig(v.([]interface{}))
	}

	return config.Client()
}

func expandAzureConfig(input []interface{}) *AzureConfig {
	if len(input) == 0 {
		return nil
	}

	values := input[0].(map[string]interface{})

	config := AzureConfig{}

	if v, ok := values["workspace_id"]; ok {
		config.WorkspaceID = v.(string)
	}

	if v, ok := values["service_principal"]; ok {
		servicePrincipal := expandAzureServicePrincipalConfig(v.([]interface{}))
		config.ServicePrincipal = servicePrincipal
	}

	return &config
}

func expandAzureServicePrincipalConfig(input []interface{}) *AzureServicePrincipalConfig {
	if len(input) == 0 {
		return nil
	}

	values := input[0].(map[string]interface{})

	config := AzureServicePrincipalConfig{}

	if v, ok := values["service_principal"]; ok {
		config.ClientID = v.(string)
	}

	if v, ok := values["client_secret"]; ok {
		config.ClientSecret = v.(string)
	}

	if v, ok := values["tenant_id"]; ok {
		config.TenantID = v.(string)
	}

	if v, ok := values["environment"]; ok {
		config.Environment = v.(string)
	}

	return &config
}
