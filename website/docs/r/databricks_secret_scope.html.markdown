---
layout: "databricks"
page_title: "Databricks: databricks_secret_scope"
sidebar_current: "docs-databricks-resource-secret-scope"
description: |-
  Create a secret scope.
---

# databricks_secret_scope

Create a [Databricks-backed](https://docs.databricks.com/dev-tools/api/latest/secrets.html#secretscopebackendtype) secret scope in which secrets are stored in Databricks-managed storage and encrypted with a cloud-based specific encryption key.

## Example Usage

```hcl
resource "databricks_secret_scope" "example" {
  scope = "example"
}
```

## Argument Reference

The following arguments are supported:

* `scope` - (Required) Scope name requested by the user. Scope names are unique.

* `initial_manage_principal` - (Optional) The principal that is initially granted `MANAGE` permission to the created scope.

-> **NOTE:** If `initial_manage_principal` is specified, the initial ACL applied to the scope is applied to the supplied principal (user or group) with `MANAGE` permissions. The only supported principal for this option is the group `users`, which contains all users in the workspace. If `initial_manage_principal` is not specified, the initial ACL with `MANAGE` permission applied to the scope is assigned to the API request issuer's user identity.
