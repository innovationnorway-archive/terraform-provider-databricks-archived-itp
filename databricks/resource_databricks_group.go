package databricks

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/innovationnorway/go-databricks/groups"
)

func resourceDatabricksGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceDatabricksGroupCreate,
		Read:   resourceDatabricksGroupRead,
		Delete: resourceDatabricksGroupDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
		},
	}
}

func resourceDatabricksGroupCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Meta).Groups
	ctx := meta.(*Meta).StopContext

	name := d.Get("name").(string)

	attributes := groups.Attributes{
		GroupName: &name,
	}

	resp, err := client.Create(ctx, attributes)
	if err != nil {
		return fmt.Errorf("unable to create group: %s", err)
	}

	d.Set("name", resp.GroupName)
	d.SetId(*resp.GroupName)

	return resourceDatabricksGroupRead(d, meta)
}

func resourceDatabricksGroupRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Meta).Groups
	ctx := meta.(*Meta).StopContext

	name := d.Get("name").(string)

	resp, err := client.ListMembers(ctx, name)
	if err != nil {
		if resp.IsHTTPStatus(404) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("unable to get group members: %s", err)
	}

	d.Set("name", name)

	return nil
}

func resourceDatabricksGroupDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Meta).Groups
	ctx := meta.(*Meta).StopContext

	name := d.Get("name").(string)

	attributes := groups.DeleteAttributes{
		GroupName: &name,
	}

	_, err := client.Delete(ctx, attributes)
	if err != nil {
		return fmt.Errorf("unable to delete group: %s", err)
	}

	d.SetId("")

	return nil
}
