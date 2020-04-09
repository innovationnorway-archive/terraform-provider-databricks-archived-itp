package databricks

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccDatabricksCluster_Azure(t *testing.T) {
	resourceName := "databricks_cluster.test"
	clusterName := acctest.RandString(6)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDatabricksClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDatabricksClusterAzure(clusterName),
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

func TestAccDatabricksCluster_AWS(t *testing.T) {
	resourceName := "databricks_cluster.test"
	clusterName := acctest.RandString(6)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDatabricksClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDatabricksClusterAWS(clusterName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "cluster_name", clusterName),
					resource.TestCheckResourceAttr(resourceName, "spark_version", "6.3.x-scala2.11"),
					resource.TestCheckResourceAttr(resourceName, "node_type_id", "m4.large"),
					resource.TestCheckResourceAttr(resourceName, "autotermination_minutes", "120"),
					resource.TestCheckResourceAttr(resourceName, "aws_attributes.0.ebs_volume_type", "GENERAL_PURPOSE_SSD"),
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

func testAccDatabricksClusterAzure(clusterName string) string {
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

func testAccDatabricksClusterAWS(clusterName string) string {
	return fmt.Sprintf(`
resource "databricks_cluster" "test" {
  cluster_name  = "%s"
  spark_version = "6.3.x-scala2.11"
  node_type_id  = "m4.large"

  autoscale {
    min_workers = 2
    max_workers = 8
  }

  aws_attributes {
    first_on_demand  = 0
    ebs_volume_type  = "GENERAL_PURPOSE_SSD"
    ebs_volume_count = 1
    ebs_volume_size  = 100
  }

  autotermination_minutes = 120
}
`, clusterName)
}
