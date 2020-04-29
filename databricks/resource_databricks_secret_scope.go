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
		Read:   resourceDatabricksSecretScopeRead,
		Delete: resourceDatabricksSecretScopeDelete,

		Schema: map[string]*schema.Schema{
			"scope": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"initial_manage_principal": {
				Type:         schema.TypeString,
				Optional:     true,
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

	attributes := secrets.CreateScopeAttributes{
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

	return resourceDatabricksSecretScopeRead(d, meta)
}

func resourceDatabricksSecretScopeRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Meta).Secrets
	ctx := meta.(*Meta).StopContext

	scope := d.Get("scope").(string)

	resp, err := client.ListScopes(ctx)
	if err != nil {
		return fmt.Errorf("unable to get scope: %s", err)
	}
	if !isExistingScope(scope, resp.Scopes) {
		d.SetId("")
		return nil
	}

	d.Set("scope", scope)

	return nil
}

func resourceDatabricksSecretScopeDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Meta).Secrets
	ctx := meta.(*Meta).StopContext

	scope := d.Get("scope").(string)

	attributes := secrets.DeleteScopeAttributes{
		Scope: &scope,
	}

	_, err := client.DeleteScope(ctx, attributes)
	if err != nil {
		return fmt.Errorf("unable to delete scope: %s", err)
	}

	d.SetId("")

	return nil
}

func isExistingScope(scope string, scopes *[]secrets.ScopeAttributes) bool {
	if scopes == nil {
		return false
	}

	for _, item := range *scopes {
		if scope == *item.Name {
			return true
		}
	}

	return false
}
