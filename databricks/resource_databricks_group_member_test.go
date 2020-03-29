package databricks

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccDatabricksGroupMember(t *testing.T) {
	resourceName := "databricks_group_member.test"
	groupName := acctest.RandString(6)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDatabricksGroupMember(groupName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "parent_name", groupName),
					resource.TestCheckResourceAttrSet(resourceName, "user_name"),
				),
			},
		},
	})
}

func testAccDatabricksGroupMember(groupName string) string {
	return fmt.Sprintf(`
data "databricks_group_members" "test" {
  name = "admins"
}

resource "databricks_group" "test" {
  name = "%s"
}

resource "databricks_group_member" "test" {
  parent_name = databricks_group.test.name
  user_name   = data.databricks_group_members.test.members[0].user_name
}
`, groupName)
}
