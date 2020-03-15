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

data "databricks_cluster" "example" {
  cluster_id = "0308-153622-deity853"
}
```

## Argument Reference

The following arguments are supported in the `provider` block:

* `host` - (Required) A Databricks host (should begin with `https://`). This is the URL of the Databricks instance. It can also be sourced from the `DATABRICKS_HOST` environment variable.

* `token` - (Required) A [personal access token](https://docs.databricks.com/dev-tools/api/latest/authentication.html#authentication). This is used to access Databricks REST APIs. It can also be sourced from the `DATABRICKS_TOKEN` environment variable.
