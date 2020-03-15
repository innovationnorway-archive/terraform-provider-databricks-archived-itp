---
layout: "databricks"
page_title: "Databricks: databricks_cluster"
sidebar_current: "docs-databricks-resource-cluster"
description: |-
  Create a new Apache Spark cluster.
---

# databricks_cluster

Create a new Apache Spark cluster. This method acquires new instances from the cloud provider if necessary.

## Example Usage

```hcl
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

## Argument Reference

The following arguments are supported:

* `spark_version` - (Required) The runtime version of the cluster. A list of available runtime versions can be retrieved by using the `databricks_cluster_runtime_versions` data source.

* `node_type_id` - (Required) This field encodes, through a single value, the resources available to each of the Spark nodes in this cluster. A list of available node types can be retrieved by using the `databricks_cluster_node_types` data source.

* `cluster_name` - (Optional) The name of the cluster. This doesn't have to be unique. If not specified at creation, the cluster name will be an empty string (`""`).

* `num_workers` - (Optional) The number of worker nodes that this cluster should have. A cluster has one Spark driver and `num_workers` executors for a total of `num_workers` + 1 Spark nodes.

* `autoscale` - (Optional) A `autoscale` block as defined below. Parameters needed in order to automatically scale clusters up and down based on load.

* `spark_conf` - (Optional) A map containing a set of optional, user-specified Spark configuration key-value pairs.

* `aws_attributes` - (Optional) A `aws_attributes` block as defined below. Attributes related to clusters running on Amazon Web Services. If not specified at cluster creation, a set of default values will be used.

* `driver_node_type_id` - (Optional) The node type of the Spark driver. This field is optional; if unset, the driver node type will be set as the same value as `node_type_id` defined above.

* `ssh_public_keys` - (Optional) A list of SSH public key contents that will be added to each Spark node in this cluster. The corresponding private keys can be used to login with the user name ubuntu on port `2200`. Up to 10 keys can be specified.

* `custom_tags` - (Optional) Additional tags for cluster resources. Databricks will tag all cluster resources (e.g., AWS instances and EBS volumes) with these tags in addition to `default_tags`.

* `cluster_log_conf` - (Optional) A `cluster_log_conf` block as defined below. The configuration for delivering Spark logs to a long-term storage destination. Only one destination can be specified for one cluster. If the conf is given, the logs will be delivered to the destination every 5 mins.

* `init_scripts` - (Optional) A `init_scripts` block as defined below. The configuration for storing init scripts. Any number of destinations can be specified. The scripts are executed sequentially in the order provided. If `cluster_log_conf` is specified, init script logs are sent to `<destination>/<cluster-ID>/init_scripts`.

* `docker_image` - (Optional) A `docker_image` block as defined below. Docker image for a [custom container](https://docs.databricks.com/clusters/custom-containers.html).

* `spark_env_vars` - (Optional) An map containing a set of optional, user-specified environment variable key-value pairs.

* `autotermination_minutes` - (Optional) Automatically terminates the cluster after it is inactive for this time in minutes. If not set, this cluster will not be automatically terminated. If specified, the threshold must be between 10 and 10000 minutes. You can also set this value to 0 to explicitly disable automatic termination.

* `enable_elastic_disk` - (Optional) Set to `true` enable [autoscaling local storage](https://docs.databricks.com/clusters/configure.html#autoscaling-local-storage) when enabled, this cluster will dynamically acquire additional disk space when its Spark workers are running low on disk space.

* `instance_pool_id` - (Optional) The ID of the instance pool to which the cluster belongs.

* `idempotency_token` - (Optional) An optional token that can be used to guarantee the idempotency of cluster creation requests. If an active cluster with the provided token already exists, the request will not create a new cluster, but it will return the ID of the existing cluster instead. The existence of a cluster with the same token is not checked against terminated clusters.

---

A `autoscale` block supports the following:

* `min_workers` - (Required) The minimum number of workers to which the cluster can scale down when underutilized. It is also the initial number of workers the cluster will have after creation.

* `max_workers` - (Required) The maximum number of workers to which the cluster can scale up when overloaded. `max_workers` must be strictly greater than `min_workers`.

---

A `aws_attributes` block supports the following:

* `first_on_demand` - The first `first_on_demand` nodes of the cluster will be placed on on-demand instances. If this value is greater than 0, the cluster driver node will be placed on an on-demand instance. If this value is greater than or equal to the current cluster size, all nodes will be placed on on-demand instances. If this value is less than the current cluster size, `first_on_demand` nodes will be placed on on-demand instances and the remainder will be placed on `availability` instances.

* `availability` - Availability type used for all subsequent nodes past the `first_on_demand` ones. Note: If `first_on_demand` is zero, this availability type will be used for the entire cluster.

* `zone_id` - Identifier for the availability zone/datacenter in which the cluster resides. This string will be of a form like `us-west-2a`. The provided availability zone must be in the same region as the Databricks deployment.

* `instance_profile_arn` - Nodes for this cluster will only be placed on AWS instances with this instance profile. If omitted, nodes will be placed on instances without an IAM instance profile.

* `spot_bid_price_percent` - The max price for AWS spot instances, as a percentage of the corresponding instance type's on-demand price. For example, if this field is set to 50, and the cluster needs a new `i3.xlarge` spot instance, then the max price is half of the price of on-demand `i3.xlarge` instances. Similarly, if this field is set to 200, the max price is twice the price of on-demand `i3.xlarge` instances. If not specified, the default value is 100.

* `ebs_volume_type` - The type of EBS volumes that will be launched with this cluster.

* `ebs_volume_count` - The number of volumes launched for each instance. You can choose up to 10 volumes. This feature is only enabled for supported node types. 

* `ebs_volume_size` - The size of each EBS volume (in GiB) launched for each instance. For general purpose SSD, this value must be within the range 100 - 4096. For throughput optimized HDD, this value must be within the range 500 - 4096.

---

A `cluster_log_conf` block supports the following:

* `dbfs` - A `dbfs` block as defined below. DBFS location of cluster log.

* `s3` - A `s3` block as defined below. S3 location of cluster log.

---

A `init_scripts` block supports the following:

* `dbfs` - A `dbfs` block as defined below. DBFS location of init script.

* `s3` - A `s3` block as defined below. S3 location of init script.

---

A `dbfs` block supports the following:

* `destination` - DBFS destination, e.g. `dbfs:/my/path`.

---

A `s3` block supports the following:

* `destination` - S3 destination, e.g. `s3://my-bucket/some-prefix` You must set cluster an IAM role and the role must have write access to the destination. You cannot use AWS keys.

* `region` - S3 region, e.g. `us-west-2`. Either `region` or `endpoint` must be set. If both are set, `endpoint` is used.

* `endpoint` - S3 endpoint, e.g. `https://s3-us-west-2.amazonaws.com`. Either `region` or `endpoint` needs to be set. If both are set, `endpoint` is used.

* `enable_encryption` - Set to `true` to enable server side encryption. Default is `false`.

* `encryption_type` - The encryption type, it could be `sse-s3` or `sse-kms`. It is used only when encryption is enabled and the default type is `sse-s3`.

* `kms_key` - KMS key used if encryption is enabled and encryption type is set to `sse-kms`.

* `canned_acl` - Set canned access control list, e.g. `bucket-owner-full-control`. If `canned_cal` is set, the cluster IAM role must have `s3:PutObjectAcl` permission on the destination bucket and prefix. 

---

A `docker_image` block supports the following:

* `url` - (Required) The URL for the Docker image.

* `basic_auth` - (Optional) A `docker_basic_auth` block as defined below. Basic authentication information for Docker repository.

---

A `docker_basic_auth` block supports the following:

* `username` - (Required) User name for the Docker repository.

* `password` - (Required) Password for the Docker repository.

---

## Attributes Reference

The following attributes are exported:

* `cluster_id` - The canonical identifier for the cluster.
