package databricks

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/innovationnorway/go-databricks/groups"
)

func TestAccDatabricksGroupMember_basic(t *testing.T) {
	resourceName := "databricks_group_member.test"
	groupName := acctest.RandString(6)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDatabricksGroupMemberDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDatabricksGroupMemberBasic(groupName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "parent_name", groupName),
					resource.TestCheckResourceAttrSet(resourceName, "user_name"),
				),
			},
		},
	})
}

func testAccCheckDatabricksGroupMemberDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "databricks_group_member" {
			continue
		}

		parentName := rs.Primary.Attributes["parent_name"]
		groupName := rs.Primary.Attributes["group_name"]
		username := rs.Primary.Attributes["user_name"]

		principalName := groups.PrincipalName{
			GroupName: &groupName,
			UserName:  &username,
		}

		client := testAccProvider.Meta().(*Meta).Groups
		ctx := testAccProvider.Meta().(*Meta).StopContext
		resp, err := client.ListMembers(ctx, parentName)
		if err != nil {
			if resp.IsHTTPStatus(404) || !isPrincipalMemberOf(principalName, resp.Members) {
				return nil
			}
			return err
		}

		return fmt.Errorf("Databricks group member still exists:\n%#v", resp)
	}

	return nil
}

func testAccDatabricksGroupMemberBasic(groupName string) string {
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
