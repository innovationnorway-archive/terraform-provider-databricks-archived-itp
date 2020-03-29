package databricks

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccDatabricksWorkspaceImport(t *testing.T) {
	resourceName := "databricks_workspace_import.test"
	path := fmt.Sprintf("/Shared/%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDatabricksWorkspaceImport(path),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "path", path),
				),
			},
		},
	})
}

func testAccDatabricksWorkspaceImport(path string) string {
	return fmt.Sprintf(`
resource "databricks_workspace_import" "test" {
  path     = "%s"
  content  = base64encode("print(\"Hello, world!\")")
  language = "PYTHON"
}
`, path)
}
