---
layout: "databricks"
page_title: "Databricks: databricks_secret"
sidebar_current: "docs-databricks-resource-secret"
description: |-
  Create a new secret.
---

# databricks_secret

Create a secret under the provided scope with the given name. If a secret already exists with the same name, this overwrites the existing secret's value.

## Example Usage

```hcl
resource "databricks_secret_scope" "example" {
  scope = "example"
}

resource "databricks_secret" "example" {
  scope        = databricks_secret_scope.example.scope
  key          = "foo"
  string_value = "bar"
}
```

## Argument Reference

The following arguments are supported:

* `scope` - (Required) The name of the scope to which the secret will be associated with.

* `key` - (Required) A unique name to identify the secret.

* `string_value` - (Optional) The value of the secret stored in UTF-8 (MB4) form.

* `bytes_value` - (Optional) The value of the secret stored as bytes.

-> **NOTE:** Either a `string_value` or `bytes_value` must be specified - but not both.
