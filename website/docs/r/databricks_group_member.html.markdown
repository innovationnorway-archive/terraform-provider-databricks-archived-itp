---
layout: "databricks"
page_title: "Databricks: databricks_group_member"
sidebar_current: "docs-databricks-resource-group-member"
description: |-
  Add a user or group to a group.
---

# databricks_group_member

Add a user or group to a group. 

## Example Usage

```hcl
resource "databricks_group_member" "example" {
  parent_name = "example"
  user_name   = "user@example.com"
}
```

## Argument Reference

The following arguments are supported:

* `parent_name` - (Required) Name of the parent group to which the new member will be added.

* `user_name` - (Optional) - A user name.

* `group_name` - (Optional) A group name.

-> **NOTE:** Either a `user_name` or `group_name` must be specified - but not both.
