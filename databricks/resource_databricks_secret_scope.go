package databricks

import (
	"fmt"

	"github.com/Azure/go-autorest/autorest/to"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/innovationnorway/go-databricks/secrets"
)

func resourceDatabricksSecretScope() *schema.Resource {
	return &schema.Resource{
		Create: resourceDatabricksSecretScopeCreate,

		Schema: map[string]*schema.Schema{
			"scope": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"initial_manage_principal": {
				Type:         schema.TypeString,
				Required:     false,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
		},
	}
}

func resourceDatabricksSecretScopeCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Meta).Secrets
	ctx := meta.(*Meta).StopContext

	scope := d.Get("scope").(string)
	initialManagePrincipal := d.Get("initial_manage_principal").(string)

	attributes := secrets.CreateScopeRequest{
		Scope: &scope,
	}

	if initialManagePrincipal != "" {
		attributes.InitialManagePrincipal = &initialManagePrincipal
	}

	_, err := client.CreateScope(ctx, attributes)
	if err != nil {
		return fmt.Errorf("unable to create secret scope: %s", err)
	}

	d.Set("scope", attributes.Scope)
	d.Set("initial_manage_principal", attributes.InitialManagePrincipal)
	d.SetId(to.String(attributes.Scope))

	return nil
}
