---
layout: "databricks"
page_title: "Provider: Databricks"
sidebar_current: "docs-databricks-index"
description: |-
  The Databricks provider is used to interact with Databricks services.
---

# Databricks Provider

The Databricks provider is used to interact with [Databricks](https://databricks.com/) services.

Use the navigation to the left to read about the available resources.

## Example Usage

```hcl
provider "databricks" {
  host  = var.databricks_host
  token = var.databricks_token
}

resource "databricks_cluster" "example" {
  cluster_name  = "example"
  spark_version = "6.3.x-scala2.11"
  node_type_id  = "Standard_DS3_v2"

  autoscale {
    min_workers = 2
    max_workers = 8
  }

  autotermination_minutes = 120
}
```

## Argument Reference

The following arguments are supported in the `provider` block:

* `host` - (Required) A Databricks host (should begin with `https://`). This is the URL of the Databricks instance. It can also be sourced from the `DATABRICKS_HOST` environment variable.

* `token` - (Required) A [personal access token](https://docs.databricks.com/dev-tools/api/latest/authentication.html#authentication). This is used to access Databricks REST APIs. It can also be sourced from the `DATABRICKS_TOKEN` environment variable.

* `organization_id` - (Optional) A workspace organization ID. The random number after `o=` in the workspace URL is the organization ID. It can also be sourced from the `DATABRICKS_ORGANIZATION_ID` environment variable.

* `azure` - (Optional) A `azure` block supports the following arguments:

  * `workspace_id` - (Required) The resource ID for the Azure Databricks workspace. It can also be sourced from the `DATABRICKS_AZURE_WORKSPACE_ID` environment variable.

  * `service_principal` - (Optional) A `service_principal` block supports the following arguments:

    * `client_id` - (Required) - The Application (client) ID for the Service Principal. It can also be sourced from the `DATABRICKS_AZURE_CLIENT_ID` environment variable.
    
    * `client_secret` - (Required) The client secret (password) for the Service Principal. It can also be sourced from the `DATABRICKS_AZURE_CLIENT_SECRET` environment variable.
    
    * `tenant_id` - (Required) The Directory (tenant) ID used for the Service Principal. It can also be sourced from the `DATABRICKS_AZURE_TENANT_ID` environment variable.
