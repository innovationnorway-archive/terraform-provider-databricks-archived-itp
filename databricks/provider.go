package databricks

import (
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
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("DATABRICKS_TOKEN", nil),
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"databricks_cluster": dataSourceDatabricksCluster(),
		},

		ResourcesMap: map[string]*schema.Resource{
			//"databricks_cluster": resourceDatabricksCluster(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		Token: d.Get("token").(string),
		Host:  d.Get("host").(string),
	}

	return config.Client()
}
