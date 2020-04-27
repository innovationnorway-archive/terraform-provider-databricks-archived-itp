---
layout: "databricks"
page_title: "Databricks: databricks_dbfs_mkdirs"
sidebar_current: "docs-databricks-resource-dbfs-mkdirs"
description: |-
  Create the given directory and necessary parent directories.
---

# databricks_dbfs_mkdirs

Create the given directory and necessary parent directories if they do not exist.

## Example Usage

```hcl
resource "databricks_dbfs_mkdirs" "example" {
  path = "/mnt/foo"
}
```

## Argument Reference

The following arguments are supported:

* `path` - (Required) The path of the new directory. The path should be the absolute DBFS path (e.g. `/mnt/foo/`).
