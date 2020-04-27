---
layout: "databricks"
page_title: "Databricks: databricks_dbfs_upload"
sidebar_current: "docs-databricks-resource-dbfs-upload"
description: |-
  Upload a file.
---

# databricks_dbfs_upload

Upload a file.

## Example Usage

```hcl
resource "databricks_dbfs_upload" "example" {
  path     = "/mnt/foo/bar.txt"
  contents = filebase64("bar.txt")
}
```

## Argument Reference

The following arguments are supported:

* `path` - (Required) The path of the new file. The path should be the absolute DBFS path (e.g. `/mnt/foo/bar.txt`).

* `contents` - (Required) The base64-encoded content. This currently has a limit of 1 MB.

-> **NOTE:** The amount of data that can be passed using `contents` parameter is currently limited to 1 MB; `MAX_BLOCK_SIZE_EXCEEDED` is thrown if exceeded. Using streaming upload to support large files will be added in a future release.
