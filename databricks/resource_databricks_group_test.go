package databricks

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccDatabricksGroup_basic(t *testing.T) {
	resourceName := "databricks_group.test"
	groupName := acctest.RandString(6)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDatabricksGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDatabricksGroupConfig(groupName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", groupName),
				),
			},
		},
	})
}

func testAccCheckDatabricksGroupDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "databricks_group" {
			continue
		}

		client := testAccProvider.Meta().(*Meta).Groups
		ctx := testAccProvider.Meta().(*Meta).StopContext
		resp, err := client.ListMembers(ctx, rs.Primary.ID)
		if err != nil {
			if resp.IsHTTPStatus(404) {
				return nil
			}

			return err
		}

		return fmt.Errorf("Databricks group still exists:\n%#v", resp)
	}

	return nil
}

func testAccDatabricksGroupConfig(groupName string) string {
	return fmt.Sprintf(`
resource "databricks_group" "test" {
  name = "%s"
}
`, groupName)
}
