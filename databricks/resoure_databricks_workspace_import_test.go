package databricks

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccDatabricksWorkspaceImport_basic(t *testing.T) {
	resourceName := "databricks_workspace_import.test"
	path := fmt.Sprintf("/Shared/%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDatabricksWorkspaceImportDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDatabricksWorkspaceImportBasic(path),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "path", path),
				),
			},
		},
	})
}

func testAccCheckDatabricksWorkspaceImportDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "databricks_workspace_import" {
			continue
		}

		client := testAccProvider.Meta().(*Meta).Workspace
		ctx := testAccProvider.Meta().(*Meta).StopContext
		resp, err := client.GetStatus(ctx, rs.Primary.ID)
		if err != nil {
			if resp.IsHTTPStatus(404) {
				return nil
			}
			return err
		}

		return fmt.Errorf("Databricks notebook still exists:\n%#v", resp)
	}

	return nil
}

func testAccDatabricksWorkspaceImportBasic(path string) string {
	return fmt.Sprintf(`
resource "databricks_workspace_import" "test" {
  path     = "%s"
  content  = base64encode("print(\"Hello, world!\")")
  language = "PYTHON"
}
`, path)
}
