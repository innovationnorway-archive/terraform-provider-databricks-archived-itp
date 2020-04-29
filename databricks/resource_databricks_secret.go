package databricks

import (
	"fmt"

	"github.com/Azure/go-autorest/autorest/to"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/innovationnorway/go-databricks/secrets"
)

func resourceDatabricksSecret() *schema.Resource {
	return &schema.Resource{
		Create: resourceDatabricksSecretCreate,
		Read:   resourceDatabricksSecretRead,
		Delete: resourceDatabricksSecretDelete,

		Schema: map[string]*schema.Schema{
			"scope": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"key": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"string_value": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Sensitive:    true,
			},
			"bytes_value": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Sensitive:    true,
			},
		},
	}
}

func resourceDatabricksSecretCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Meta).Secrets
	ctx := meta.(*Meta).StopContext

	scope := d.Get("scope").(string)
	key := d.Get("key").(string)
	string_value := d.Get("string_value").(string)
	bytes_value := []byte(d.Get("bytes_value").(string))

	attributes := secrets.Attributes{
		Scope: &scope,
		Key:   &key,
	}

	if string_value != "" && len(bytes_value) == 0 {
		attributes.StringValue = &string_value
	} else if string_value == "" && len(bytes_value) != 0 {
		attributes.BytesValue = &bytes_value
	} else {
		return fmt.Errorf("unable to create secret, you must specify either string_value or bytes_value")
	}

	_, err := client.Put(ctx, attributes)
	if err != nil {
		return fmt.Errorf("unable to create secret: %s", err)
	}

	d.Set("scope", attributes.Scope)
	d.Set("key", attributes.Key)
	d.SetId(to.String(attributes.Scope) + "-" + to.String(attributes.Key))

	return resourceDatabricksSecretRead(d, meta)
}

func resourceDatabricksSecretRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Meta).Secrets
	ctx := meta.(*Meta).StopContext

	scope := d.Get("scope").(string)
	key := d.Get("key").(string)

	attributes := secrets.ListSecretsAttributes{
		Scope: &scope,
	}

	resp, err := client.List(ctx, attributes)
	if err != nil {
		return fmt.Errorf("unable to get key: %s", err)
	}
	if !isExistingSecret(key, resp.SecretsProperty) {
		d.SetId("")
		return nil
	}

	d.Set("scope", scope)
	d.Set("key", key)

	return nil
}

func resourceDatabricksSecretDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Meta).Secrets
	ctx := meta.(*Meta).StopContext

	scope := d.Get("scope").(string)
	key := d.Get("key").(string)

	attributes := secrets.Attributes{
		Scope: &scope,
		Key:   &key,
	}

	_, err := client.Delete(ctx, attributes)
	if err != nil {
		return fmt.Errorf("unable to delete secret: %s", err)
	}

	d.SetId("")

	return nil
}

func isExistingSecret(key string, secrets *[]secrets.MetadataAttributes) bool {
	if secrets == nil {
		return false
	}

	for _, item := range *secrets {
		if key == *item.Key {
			return true
		}
	}

	return false
}
