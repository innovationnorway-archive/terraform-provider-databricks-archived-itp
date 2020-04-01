package databricks

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccDatabricksCluster_basic(t *testing.T) {
	resourceName := "databricks_cluster.test"
	clusterName := acctest.RandString(6)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDatabricksClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDatabricksClusterBasic(clusterName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "cluster_name", clusterName),
					resource.TestCheckResourceAttr(resourceName, "spark_version", "6.3.x-scala2.11"),
					resource.TestCheckResourceAttr(resourceName, "node_type_id", "Standard_DS3_v2"),
					resource.TestCheckResourceAttr(resourceName, "autotermination_minutes", "120"),
				),
			},
		},
	})
}

func testAccCheckDatabricksClusterDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "databricks_cluster" {
			continue
		}

		client := testAccProvider.Meta().(*Meta).Clusters
		ctx := testAccProvider.Meta().(*Meta).StopContext
		resp, err := client.Get(ctx, rs.Primary.ID)
		if err != nil {
			if resp.IsHTTPStatus(400) && isDatabricksClusterNotExistsError(err) {
				return nil
			}
			return err
		}

		return fmt.Errorf("Databricks cluster still exists:\n%#v", resp)
	}

	return nil
}

func testAccDatabricksClusterBasic(clusterName string) string {
	return fmt.Sprintf(`
resource "databricks_cluster" "test" {
  cluster_name  = "%s"
  spark_version = "6.3.x-scala2.11"
  node_type_id  = "Standard_DS3_v2"

  autoscale {
    min_workers = 1
    max_workers = 2
  }

  spark_conf = {
    "spark.databricks.delta.preview.enabled" = true
  }

  autotermination_minutes = 120
}
`, clusterName)
}
