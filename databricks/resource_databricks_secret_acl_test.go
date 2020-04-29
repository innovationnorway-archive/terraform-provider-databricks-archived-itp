package databricks

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/innovationnorway/go-databricks/secrets"
)

func TestAccDatabricksSecretACL_basic(t *testing.T) {
	resourceName := "databricks_secret_acl.test"
	scope := acctest.RandString(6)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDatabricksSecretACLDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDatabricksSecretACLConfig(scope),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "scope", scope),
					resource.TestCheckResourceAttr(resourceName, "principal", "admins"),
					resource.TestCheckResourceAttr(resourceName, "permission", "MANAGE"),
				),
			},
		},
	})
}

func testAccCheckDatabricksSecretACLDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "databricks_secret_acl" {
			continue
		}

		client := testAccProvider.Meta().(*Meta).Secrets
		ctx := testAccProvider.Meta().(*Meta).StopContext

		scope := rs.Primary.Attributes["scope"]
		principal := rs.Primary.Attributes["principal"]

		attributes := secrets.AclsAttributes{
			Scope:     &scope,
			Principal: &principal,
		}

		resp, err := client.GetAcls(ctx, attributes)
		if err != nil {
			if resp.IsHTTPStatus(404) {
				return nil
			}
			return err
		}

		return fmt.Errorf("Secret ACL still exists:\n%#v", resp)
	}

	return nil
}

func testAccDatabricksSecretACLConfig(scope string) string {
	return fmt.Sprintf(`
resource "databricks_secret_scope" "test" {
  scope = "%s"
}
	
resource "databricks_secret_acl" "test" {
  scope      = databricks_secret_scope.test.scope
  principal  = "admins"
  permission = "MANAGE"
}
`, scope)
}
