package databricks

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/innovationnorway/go-databricks/groups"
)

func dataSourceDatabricksGroupMembers() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDatabricksGroupMembersRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},

			"members": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"user_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"group_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceDatabricksGroupMembersRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Meta).Groups
	ctx := meta.(*Meta).StopContext

	name := d.Get("name").(string)

	resp, err := client.ListMembers(ctx, name)
	if err != nil {
		return fmt.Errorf("unable to get group members: %s", err)
	}

	d.Set("members", flattenGroupMembers(resp.Members))

	d.SetId(name)

	return nil
}

func flattenGroupMembers(input *[]groups.PrincipalName) []interface{} {
	result := make([]interface{}, 0)

	if input == nil {
		return []interface{}{}
	}

	for _, item := range *input {
		values := make(map[string]interface{})

		if item.UserName != nil {
			values["user_name"] = *item.UserName
		}

		if item.GroupName != nil {
			values["group_name"] = *item.GroupName
		}

		result = append(result, values)
	}

	return result
}
