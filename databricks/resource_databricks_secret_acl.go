package databricks

import (
	"fmt"

	"github.com/Azure/go-autorest/autorest/to"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/innovationnorway/go-databricks/secrets"
)

func resourceDatabricksSecretAcl() *schema.Resource {
	return &schema.Resource{
		Create: resourceDatabricksSecretAclCreate,
		Read:   resourceDatabricksSecretAclRead,
		Delete: resourceDatabricksSecretAclDelete,

		Schema: map[string]*schema.Schema{
			"scope": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"principal": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"permission": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					string(secrets.READ),
					string(secrets.WRITE),
					string(secrets.MANAGE),
				}, false),
			},
		},
	}
}

func resourceDatabricksSecretAclCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Meta).Secrets
	ctx := meta.(*Meta).StopContext

	scope := d.Get("scope").(string)
	principal := d.Get("principal").(string)
	permission := d.Get("permission").(string)

	attributes := secrets.PutSecretAclsAttributes{
		Scope:      &scope,
		Principal:  &principal,
		Permission: secrets.Permission(permission),
	}

	_, err := client.PutAcls(ctx, attributes)
	if err != nil {
		return fmt.Errorf("unable to create acl: %s", err)
	}

	d.Set("scope", attributes.Scope)
	d.Set("principal", attributes.Principal)
	d.Set("permission", string(attributes.Permission))
	d.SetId(to.String(attributes.Scope) + "-" + to.String(attributes.Principal))

	return resourceDatabricksSecretAclRead(d, meta)
}

func resourceDatabricksSecretAclRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Meta).Secrets
	ctx := meta.(*Meta).StopContext

	scope := d.Get("scope").(string)
	principal := d.Get("principal").(string)
	permission := d.Get("permission").(string)

	attributes := secrets.AclsAttributes{
		Scope:     &scope,
		Principal: &principal,
	}

	resp, err := client.GetAcls(ctx, attributes)
	if err != nil {
		if resp.IsHTTPStatus(404) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("unable to get acl: %s", err)
	}

	d.Set("scope", scope)
	d.Set("principal", principal)
	d.Set("permission", permission)

	return nil
}

func resourceDatabricksSecretAclDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Meta).Secrets
	ctx := meta.(*Meta).StopContext

	scope := d.Get("scope").(string)
	principal := d.Get("principal").(string)

	attributes := secrets.AclsAttributes{
		Scope:     &scope,
		Principal: &principal,
	}

	_, err := client.DeleteAcls(ctx, attributes)
	if err != nil {
		return fmt.Errorf("unable to remove acl: %s", err)
	}

	d.SetId("")

	return nil
}
