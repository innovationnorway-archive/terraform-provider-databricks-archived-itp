---
layout: "databricks"
page_title: "Databricks: databricks_workspace_import"
sidebar_current: "docs-databricks-resource-workspace-import"
description: |-
  Import a notebook.
---

# databricks_workspace_import

Import a notebook.

## Example Usage

```hcl
resource "databricks_workspace_import" "example" {
  path     = "/Shared/example"
  language = "PYTHON"
  content  = filebase64("example.py")
}
```

## Argument Reference

The following arguments are supported:

* `path` - (Required) The absolute path of the notebook.

* `content` - (Required) The base64-encoded content. This has a limit of 10 MB.

* `format` - (Optional) This specifies the format of the file to be imported. By default, this is `SOURCE`. However it may be one of: `SOURCE`, `HTML`, `JUPYTER`, `DBC`.

* `language` - (Optional) The language. If format is set to `SOURCE`, this field is required; otherwise, it will be ignored. Possible values are: `SCALA`, `PYTHON`, `SQL`, `R`.

## Attributes Reference

The following attributes are exported:

* `object_id` - A unique identifier for the notebook.
