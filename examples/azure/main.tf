provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "example" {
  name     = "example-resources"
  location = var.location
}

resource "azurerm_databricks_workspace" "example" {
  name                = "example"
  resource_group_name = azurerm_resource_group.example.name
  location            = azurerm_resource_group.example.location
  sku                 = "trial"
}

provider "databricks" {
  host = format("https://%s.azuredatabricks.net", azurerm_databricks_workspace.example.location)

  azure {
    workspace_id = azurerm_databricks_workspace.example.id
  }
}

resource "databricks_cluster" "example" {
  cluster_name  = "example"
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
