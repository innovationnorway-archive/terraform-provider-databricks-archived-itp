---
layout: "databricks"
page_title: "Databricks: databricks_group_members"
sidebar_current: "docs-databricks-datasource-group-members"
description: |-
  Return all of the members of a particular group.
---

# databricks_group_members

Return all of the members of a particular group.

## Example Usage

```hcl
data "databricks_group_members" "example" {
  name = "admins"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The group whose members we want to retrieve..

## Attributes Reference

The following attributes are exported:

* `members` - The users and groups that belong to the given group. A list of `members` blocks as defined below.

---

A `members` block exports the following:

* `user_name` - The name of the user.

* `group_name` - The name of the group.
