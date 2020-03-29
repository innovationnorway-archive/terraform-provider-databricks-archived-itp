package databricks

import (
	"fmt"

	"github.com/Azure/go-autorest/autorest/to"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/innovationnorway/go-databricks/workspace"
)

func resourceDatabricksWorkspaceImport() *schema.Resource {
	return &schema.Resource{
		Create: resourceDatabricksWorkspaceImportCreate,
		Read:   resourceDatabricksWorkspaceImportRead,
		Update: resourceDatabricksWorkspaceImportUpdate,
		Delete: resourceDatabricksWorkspaceImportDelete,

		Schema: map[string]*schema.Schema{
			"path": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},

			"content": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsBase64,
			},

			"format": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					string(workspace.SOURCE),
					string(workspace.HTML),
					string(workspace.JUPYTER),
					string(workspace.DBC),
				}, false),
			},

			"language": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ValidateFunc: validation.StringInSlice([]string{
					string(workspace.SCALA),
					string(workspace.PYTHON),
					string(workspace.SQL),
					string(workspace.R),
				}, false),
			},

			"object_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceDatabricksWorkspaceImportCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Meta).Workspace
	ctx := meta.(*Meta).StopContext

	path := d.Get("path").(string)

	attributes := workspace.ImportAttributes{
		Path: &path,
	}

	if v, ok := d.GetOk("format"); ok {
		attributes.Format = workspace.Format(v.(string))
	}

	if v, ok := d.GetOk("language"); ok {
		attributes.Language = workspace.Language(v.(string))
	}

	if v, ok := d.GetOk("content"); ok {
		attributes.Content = to.StringPtr(v.(string))
	}

	_, err := client.Import(ctx, attributes)
	if err != nil {
		return fmt.Errorf("unable to import object: %s", err)
	}

	d.SetId(path)

	return resourceDatabricksWorkspaceImportRead(d, meta)
}

func resourceDatabricksWorkspaceImportRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Meta).Workspace
	ctx := meta.(*Meta).StopContext

	path := d.Get("path").(string)

	resp, err := client.GetStatus(ctx, path)
	if err != nil {
		if resp.IsHTTPStatus(404) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("unable to get object status: %s", err)
	}

	d.Set("path", resp.Path)
	d.Set("language", resp.Language)
	d.Set("object_id", resp.ObjectID)

	return nil
}

func resourceDatabricksWorkspaceImportUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Meta).Workspace
	ctx := meta.(*Meta).StopContext

	path := d.Get("path").(string)

	attributes := workspace.ImportAttributes{
		Path:      &path,
		Overwrite: to.BoolPtr(true),
	}

	if v, ok := d.GetOk("format"); ok {
		attributes.Format = workspace.Format(v.(string))
	}

	if v, ok := d.GetOk("language"); ok {
		attributes.Language = workspace.Language(v.(string))
	}

	if v, ok := d.GetOk("content"); ok {
		attributes.Content = to.StringPtr(v.(string))
	}

	_, err := client.Import(ctx, attributes)
	if err != nil {
		return fmt.Errorf("unable to import object: %s", err)
	}

	return resourceDatabricksWorkspaceImportRead(d, meta)
}

func resourceDatabricksWorkspaceImportDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Meta).Workspace
	ctx := meta.(*Meta).StopContext

	path := d.Get("path").(string)

	attributes := workspace.DeleteAttributes{
		Path: &path,
	}

	_, err := client.Delete(ctx, attributes)
	if err != nil {
		return fmt.Errorf("unable to delete object: %s", err)
	}

	d.SetId("")

	return nil
}
