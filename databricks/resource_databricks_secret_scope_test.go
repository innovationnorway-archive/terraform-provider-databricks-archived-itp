package databricks

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccDatabricksSecretScope_basic(t *testing.T) {
	resourceName := "databricks_secret_scope.test"
	scope := acctest.RandString(6)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDatabricksSecretScopeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDatabricksSecretScopeConfig(scope),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "scope", scope),
					resource.TestCheckResourceAttr(resourceName, "initial_manage_principal", "users"),
				),
			},
		},
	})
}

func testAccCheckDatabricksSecretScopeDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "databricks_secret_scope" {
			continue
		}

		client := testAccProvider.Meta().(*Meta).Secrets
		ctx := testAccProvider.Meta().(*Meta).StopContext

		scope := rs.Primary.Attributes["scope"]

		resp, err := client.ListScopes(ctx)
		if err != nil {
			return err
		}

		if !isExistingScope(scope, resp.Scopes) {
			return nil
		}

		return fmt.Errorf("Secret scope still exists:\n%#v", resp)
	}

	return nil
}

func testAccDatabricksSecretScopeConfig(scope string) string {
	return fmt.Sprintf(`
resource "databricks_secret_scope" "test" {
  scope                    = "%s"
  initial_manage_principal = "users"
}
`, scope)
}
