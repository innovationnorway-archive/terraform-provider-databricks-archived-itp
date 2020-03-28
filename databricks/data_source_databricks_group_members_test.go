package databricks

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccDataSourceDatabricksGroupMembers(t *testing.T) {
	resourceName := "data.databricks_group_members.test"
	groupName := "admins"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceDatabricksGroupMembers(groupName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", groupName),
				),
			},
		},
	})
}

func testAccDataSourceDatabricksGroupMembers(groupName string) string {
	return fmt.Sprintf(`
data "databricks_group_members" "test" {
  name = "%s"
}
`, groupName)
}
