package databricks

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccDatabricksCluster(t *testing.T) {
	resourceName := "databricks_cluster.test"
	clusterName := acctest.RandString(6)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDatabricksCluster(clusterName),
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

func testAccDatabricksCluster(clusterName string) string {
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
