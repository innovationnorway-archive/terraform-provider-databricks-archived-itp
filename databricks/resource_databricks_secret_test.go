package databricks

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/innovationnorway/go-databricks/secrets"
)

func TestAccDatabricksSecret_basic(t *testing.T) {
	resourceName := "databricks_secret.test"
	scope := acctest.RandString(6)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDatabricksSecretDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDatabricksSecretConfig(scope),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "scope", scope),
					resource.TestCheckResourceAttr(resourceName, "key", "foo"),
					resource.TestCheckResourceAttr(resourceName, "string_value", "bar"),
				),
			},
		},
	})
}

func testAccCheckDatabricksSecretDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "databricks_secret" {
			continue
		}

		client := testAccProvider.Meta().(*Meta).Secrets
		ctx := testAccProvider.Meta().(*Meta).StopContext

		scope := rs.Primary.Attributes["scope"]
		key := rs.Primary.Attributes["key"]

		attributes := secrets.ListSecretsAttributes{
			Scope: &scope,
		}

		resp, err := client.List(ctx, attributes)
		if err != nil {
			if resp.IsHTTPStatus(404) || !isExistingSecret(key, resp.SecretsProperty) {
				return nil
			}

			return err
		}

		return fmt.Errorf("Secret still exists:\n%#v", resp)
	}

	return nil
}

func testAccDatabricksSecretConfig(scope string) string {
	return fmt.Sprintf(`
resource "databricks_secret_scope" "test" {
  scope = "%s"
}

resource "databricks_secret" "test" {
  scope        = databricks_secret_scope.test.scope
  key          = "foo"
  string_value = "bar"
}
`, scope)
}
