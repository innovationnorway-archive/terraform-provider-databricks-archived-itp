package databricks

import (
	"fmt"

	"github.com/Azure/go-autorest/autorest/to"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/innovationnorway/go-databricks/dbfs"
)

func resourceDatabricksDbfsMkdirs() *schema.Resource {
	return &schema.Resource{
		Create: resourceDatabricksDbfsMkdirsCreate,
		Read:   resourceDatabricksDbfsMkdirsRead,
		Delete: resourceDatabricksDbfsMkdirsDelete,

		Schema: map[string]*schema.Schema{
			"path": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
		},
	}
}

func resourceDatabricksDbfsMkdirsCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Meta).Dbfs
	ctx := meta.(*Meta).StopContext

	path := d.Get("path").(string)

	attributes := dbfs.MkdirsAttributes{
		Path: to.StringPtr(path),
	}

	_, err := client.Mkdirs(ctx, attributes)
	if err != nil {
		return fmt.Errorf("unable to create directory: %s", err)
	}

	d.SetId(path)

	return resourceDatabricksDbfsMkdirsRead(d, meta)
}

func resourceDatabricksDbfsMkdirsRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Meta).Dbfs
	ctx := meta.(*Meta).StopContext

	resp, err := client.GetStatus(ctx, d.Id())
	if err != nil {
		if resp.IsHTTPStatus(404) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("unable to get directory: %s", err)
	}

	d.Set("path", resp.Path)

	return nil
}

func resourceDatabricksDbfsMkdirsDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Meta).Dbfs
	ctx := meta.(*Meta).StopContext

	attributes := dbfs.DeleteAttributes{
		Path:      to.StringPtr(d.Get("path").(string)),
		Recursive: to.BoolPtr(true),
	}

	_, err := client.Delete(ctx, attributes)
	if err != nil {
		return fmt.Errorf("unable to delete directory: %s", err)
	}

	d.SetId("")

	return nil
}
