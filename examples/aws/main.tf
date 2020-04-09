provider "databricks" {
  host  = var.databricks_host
  token = var.databricks_token
}

resource "databricks_cluster" "example" {
  cluster_name  = "example"
  spark_version = "6.3.x-scala2.11"
  node_type_id  = "m4.large"

  autoscale {
    min_workers = 2
    max_workers = 8
  }

  aws_attributes {
    first_on_demand        = 1
    availability           = "SPOT_WITH_FALLBACK"
    spot_bid_price_percent = 100
    ebs_volume_type        = "GENERAL_PURPOSE_SSD"
    ebs_volume_count       = 3
    ebs_volume_size        = 100
  }

  autotermination_minutes = 120
}
