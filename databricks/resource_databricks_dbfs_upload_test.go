package databricks

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccDatabricksDbfsUpload_basic(t *testing.T) {
	resourceName := "databricks_dbfs_upload.test"
	path := "/mnt/foo/bar.txt"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDatabricksDbfsUploadDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDatabricksDbfsUploadBasic(path),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "path", path),
				),
			},
		},
	})
}

func testAccCheckDatabricksDbfsUploadDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "databricks_dbfs_upload" {
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

		return fmt.Errorf("DBFS file still exists:\n%#v", resp)
	}

	return nil
}

func testAccDatabricksDbfsUploadBasic(path string) string {
	return fmt.Sprintf(`
resource "databricks_dbfs_upload" "test" {
  path     = "%s"
  contents = base64encode("print(\"Hello, world!\")")
}
`, path)
}
