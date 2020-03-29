package databricks

import (
	"fmt"

	"github.com/Azure/go-autorest/autorest/to"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/innovationnorway/go-databricks/groups"
)

func resourceDatabricksGroupMember() *schema.Resource {
	return &schema.Resource{
		Create: resourceDatabricksGroupMemberCreate,
		Read:   resourceDatabricksGroupMemberRead,
		Delete: resourceDatabricksGroupMemberDelete,

		Schema: map[string]*schema.Schema{
			"parent_name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},

			"user_name": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},

			"group_name": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				ExactlyOneOf: []string{"user_name", "group_name"},
			},
		},
	}
}

func resourceDatabricksGroupMemberCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Meta).Groups
	ctx := meta.(*Meta).StopContext

	parentName := d.Get("parent_name").(string)
	username := d.Get("user_name").(string)
	groupName := d.Get("group_name").(string)

	attributes := groups.MemberAttributes{
		ParentName: &parentName,
	}

	if username != "" {
		attributes.UserName = &username
	}

	if groupName != "" {
		attributes.GroupName = &groupName
	}

	_, err := client.AddMember(ctx, attributes)
	if err != nil {
		return fmt.Errorf("unable to add member: %s", err)
	}

	d.SetId(getDatabricksGroupMemberID(parentName, username, groupName))

	return resourceDatabricksGroupMemberRead(d, meta)
}

func resourceDatabricksGroupMemberRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Meta).Groups
	ctx := meta.(*Meta).StopContext

	username := d.Get("user_name").(string)
	groupName := d.Get("group_name").(string)

	resp, err := client.ListParents(ctx, groupName, username)
	if err != nil {
		if resp.IsHTTPStatus(404) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("unable to get member: %s", err)
	}

	return nil
}

func resourceDatabricksGroupMemberDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Meta).Groups
	ctx := meta.(*Meta).StopContext

	parentName := d.Get("parent_name").(string)

	attributes := groups.MemberAttributes{
		ParentName: &parentName,
	}

	if v, ok := d.GetOk("user_name"); ok {
		attributes.UserName = to.StringPtr(v.(string))
	}

	if v, ok := d.GetOk("group_name"); ok {
		attributes.GroupName = to.StringPtr(v.(string))
	}

	_, err := client.RemoveMember(ctx, attributes)
	if err != nil {
		return fmt.Errorf("unable to remove member: %s", err)
	}

	d.SetId("")

	return nil
}

func getDatabricksGroupMemberID(parentName, userName, groupName string) string {
	if userName != "" {
		return fmt.Sprintf("user:%s:%s", parentName, userName)
	}

	return fmt.Sprintf("group:%s:%s", parentName, groupName)
}
