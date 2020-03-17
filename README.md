# Terraform Provider for Databricks

[![Build Status](https://travis-ci.com/innovationnorway/terraform-provider-databricks.svg?branch=master)](https://travis-ci.com/innovationnorway/terraform-provider-databricks)

## Requirements

-	[Terraform](https://www.terraform.io/downloads.html) >= 0.12.x
-	[Go](https://golang.org/doc/install) >= 1.13

## Usage

```hcl
provider "databricks" {
  host  = var.databricks_host
  token = var.databricks_token
}

resource "databricks_cluster" "example" {
  cluster_name  = "example"
  spark_version = "6.3.x-scala2.11"
  node_type_id  = "Standard_DS3_v2"

  autoscale {
    min_workers = 2
    max_workers = 8
  }

  autotermination_minutes = 120
}
```

## Contributing

To build the provider:

```sh
$ go build
```

To test the provider:

```sh
$ go test -v ./...
```

To run all acceptance tests:

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ TF_ACC=1 go test -v ./...
```

To run a subset of acceptance tests:

```sh
$ TF_ACC=1 go test -v ./... -run=TestAccDatabricksCluster
```

The following environment variables must be set prior to running acceptance tests:

- `DATABRICKS_HOST`
- `DATABRICKS_TOKEN`
