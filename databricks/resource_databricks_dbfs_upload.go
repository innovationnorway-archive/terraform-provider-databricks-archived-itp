package databricks

import (
	"fmt"

	"github.com/Azure/go-autorest/autorest/to"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/innovationnorway/go-databricks/dbfs"
)

func resourceDatabricksDbfsUpload() *schema.Resource {
	return &schema.Resource{
		Create: resourceDatabricksDbfsUploadCreate,
		Read:   resourceDatabricksDbfsUploadRead,
		Update: resourceDatabricksDbfsUploadUpdate,
		Delete: resourceDatabricksDbfsUploadDelete,

		Schema: map[string]*schema.Schema{
			"path": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},

			"contents": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsBase64,
			},
		},
	}
}

func resourceDatabricksDbfsUploadCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Meta).Dbfs
	ctx := meta.(*Meta).StopContext

	path := d.Get("path").(string)

	attributes := dbfs.PutAttributes{
		Path: &path,
	}

	if v, ok := d.GetOk("contents"); ok {
		attributes.Contents = to.StringPtr(v.(string))
	}

	_, err := client.Put(ctx, attributes)
	if err != nil {
		return fmt.Errorf("unable to upload file: %s", err)
	}

	d.SetId(path)

	return resourceDatabricksDbfsUploadRead(d, meta)
}

func resourceDatabricksDbfsUploadRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Meta).Dbfs
	ctx := meta.(*Meta).StopContext

	path := d.Get("path").(string)

	resp, err := client.GetStatus(ctx, path)
	if err != nil {
		if resp.IsHTTPStatus(404) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("unable to get file status: %s", err)
	}

	d.Set("path", resp.Path)

	return nil
}

func resourceDatabricksDbfsUploadUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Meta).Dbfs
	ctx := meta.(*Meta).StopContext

	path := d.Get("path").(string)

	attributes := dbfs.PutAttributes{
		Path:      &path,
		Overwrite: to.BoolPtr(true),
	}

	if v, ok := d.GetOk("contents"); ok {
		attributes.Contents = to.StringPtr(v.(string))
	}

	_, err := client.Put(ctx, attributes)
	if err != nil {
		return fmt.Errorf("unable to upload file: %s", err)
	}

	return resourceDatabricksDbfsUploadRead(d, meta)
}

func resourceDatabricksDbfsUploadDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Meta).Dbfs
	ctx := meta.(*Meta).StopContext

	path := d.Get("path").(string)

	attributes := dbfs.DeleteAttributes{
		Path: &path,
	}

	_, err := client.Delete(ctx, attributes)
	if err != nil {
		return fmt.Errorf("unable to delete file: %s", err)
	}

	d.SetId("")

	return nil
}
