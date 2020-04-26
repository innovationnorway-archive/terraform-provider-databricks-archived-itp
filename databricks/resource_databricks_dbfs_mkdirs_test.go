package databricks

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccDatabricksDbfsMkdirs_basic(t *testing.T) {
	resourceName := "databricks_dbfs_mkdirs.test"
	path := "/mnt/foo"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDatabricksDbfsMkdirsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDatabricksDbfsMkdirsConfig(path),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "path", path),
				),
			},
		},
	})
}

func testAccCheckDatabricksDbfsMkdirsDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "databricks_dbfs_mkdirs" {
			continue
		}

		client := testAccProvider.Meta().(*Meta).Dbfs
		ctx := testAccProvider.Meta().(*Meta).StopContext
		resp, err := client.GetStatus(ctx, rs.Primary.ID)
		if err != nil {
			if resp.IsHTTPStatus(404) {
				return nil
			}

			return err
		}

		return fmt.Errorf("DBFS directory still exists:\n%#v", resp)
	}

	return nil
}

func testAccDatabricksDbfsMkdirsConfig(path string) string {
	return fmt.Sprintf(`
resource "databricks_dbfs_mkdirs" "test" {
  path = "%s"
}
`, path)
}
