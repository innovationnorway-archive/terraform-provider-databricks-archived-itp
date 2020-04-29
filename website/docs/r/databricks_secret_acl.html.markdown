---
layout: "databricks"
page_title: "Databricks: databricks_secret_acl"
sidebar_current: "docs-databricks-resource-secret-acl"
description: |-
  Create or overwrite the ACL associated with the given principal.
---

# databricks_secret_acl

Create or overwrite the ACL associated with the given principal (user or group) on the specified scope point.

## Example Usage

```hcl
resource "databricks_secret_scope" "example" {
  scope = "example"
}
	
resource "databricks_secret_acl" "example" {
  scope      = databricks_secret_scope.example.scope
  principal  = "data-scientists"
  permission = "READ"
}
```

## Argument Reference

The following arguments are supported:

* `scope` - (Required) The name of the scope to apply permissions to.

* `principal` - (Required) The principal to which the permission is applied.

* `permission` - (Optional) The permission level applied to the principal. Possible values are: `READ`, `WRITE` and `MANAGE`.
