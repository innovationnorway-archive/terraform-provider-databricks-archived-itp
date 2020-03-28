---
layout: "databricks"
page_title: "Databricks: databricks_group"
sidebar_current: "docs-databricks-resource-group"
description: |-
  Create a new group.
---

# databricks_group

Create a new group with the given name.

## Example Usage

```hcl
resource "databricks_group" "example" {
  name = "example"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the group; must be unique among groups owned by this organization.
