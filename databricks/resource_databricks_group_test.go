package databricks

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccDatabricksGroup(t *testing.T) {
	resourceName := "databricks_group.test"
	groupName := acctest.RandString(6)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDatabricksGroup(groupName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", groupName),
				),
			},
		},
	})
}

func testAccDatabricksGroup(groupName string) string {
	return fmt.Sprintf(`
resource "databricks_group" "test" {
  name = "%s"
}
`, groupName)
}
