package databricks

import (
	"fmt"

	"github.com/Azure/go-autorest/autorest/to"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/innovationnorway/go-databricks/groups"
)

func resourceDatabricksGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceDatabricksGroupCreate,
		Read:   resourceDatabricksGroupRead,
		Delete: resourceDatabricksGroupDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
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

func expandGroupMembers(input []interface{}) *[]groups.PrincipalName {
	if len(input) == 0 {
		return nil
	}

	results := make([]groups.PrincipalName, 0)

	for _, item := range input {
		values := item.(map[string]interface{})
		result := groups.PrincipalName{}

		if v, ok := values["user_name"]; ok {
			result.UserName = to.StringPtr(v.(string))
		}

		if v, ok := values["group_name"]; ok {
			result.GroupName = to.StringPtr(v.(string))
		}

		results = append(results, result)
	}

	return &results
}
