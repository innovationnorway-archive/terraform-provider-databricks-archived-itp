---
layout: "databricks"
page_title: "Databricks: databricks_cluster"
sidebar_current: "docs-databricks-datasource-cluster"
description: |-
  Get details about a Databricks cluster.
---

# databricks_cluster

Use this data source to retrieve information for a cluster. Clusters can be described while they are running or up to 30 days after they are terminated.

## Example Usage

```hcl
data "databricks_cluster" "example" {
  cluster_id = "0308-153622-deity853"
}
```

## Argument Reference

The following arguments are supported:

* `cluster_id` - (Required) The ID of the cluster about which to retrieve information.

## Attributes Reference

* `num_workers` - The number of worker nodes that this cluster should have.

* `autoscale` - A `autoscale` block as defined below. Parameters needed in order to automatically scale clusters up and down based on load.

* `creator_user_name` - The creator user name.

* `driver` - A `driver` block as defined below. Node on which the Spark driver resides.

* `executors` - A `driver` block as defined below. Nodes on which the Spark executors reside.

* `spark_context_id` - A canonical SparkContext identifier. This value changes when the Spark driver restarts. 

* `jdbc_port` - Port on which Spark JDBC server is listening in the driver node.

* `cluster_name` - The name of the cluster.

* `spark_version` - The runtime version of the cluster.

* `spark_conf` - User-specified Spark configuration key-value pairs.

* `aws_attributes` - A `aws_attributes` block as defined below. Attributes related to clusters running on Amazon Web Services.

* `node_type_id` - This field encodes, through a single value, the resources available to each of the Spark nodes in this cluster.

* `driver_node_type_id` - The node type of the Spark driver.

* `ssh_public_keys` - A list of SSH public key contents that will be added to each Spark node in this cluster. 

* `custom_tags` - Additional tags for cluster resources.

* `cluster_log_conf` - A `cluster_log_conf` block as defined below. The configuration for delivering Spark logs to a long-term storage destination.

* `init_scripts` - A `init_scripts` block as defined below. The configuration for storing init scripts. 

* `docker_image` - A `docker_image` block as defined below. Docker image for a custom container.

* `spark_env_vars` - User-specified environment variable key-value pairs.

* `autotermination_minutes` - Automatically terminates the cluster after it is inactive for this time in minutes.

* `enable_elastic_disk` - Whether [autoscaling local storage](https://docs.databricks.com/clusters/configure.html#autoscaling-local-storage) is enabled.

* `instance_pool_id` - The ID of the instance pool to which the cluster belongs.

* `cluster_source` - Determines whether the cluster was created by a user through the UI, by the Databricks Jobs scheduler, or through an API request.

* `state` - The state of the cluster.

* `state_message` - A message associated with the most recent state transition.

* `start_time` - Time (in epoch milliseconds) when the cluster creation request was received (when the cluster entered a `PENDING` state).

* `terminated_time` - Time (in epoch milliseconds) when the cluster was terminated, if applicable.

* `last_state_loss_time` - Time when the cluster driver last lost its state (due to a restart or driver failure).

* `last_activity_time` - Time (in epoch milliseconds) when the cluster was last active. A cluster is active if there is at least one command that has not finished on the cluster.

* `cluster_memory_mb` - Total amount of cluster memory, in megabytes.

* `cluster_cores` - Number of CPU cores available for this cluster.

* `default_tags` - Tags that are added by Databricks.

* `cluster_log_status` - A `cluster_log_status` block as defined below. Cluster log delivery status.

* `termination_reason` - A `termination_reason` block as defined below. Information about why the cluster was terminated.

---

A `autoscale` block exports the following:

* `min_workers` - The minimum number of workers to which the cluster can scale down when underutilized. It is also the initial number of workers the cluster will have after creation.

* `max_workers` - The maximum number of workers to which the cluster can scale up when overloaded.

---

A `driver` block exports the following:

* `private_ip` - Private IP address (typically a 10.x.x.x address) of the Spark node. This is different from the private IP address of the host instance.

* `public_dns` - Public DNS address of this node. This address can be used to access the Spark JDBC server on the driver node.

* `node_id` - Globally unique identifier for this node.

* `instance_id` - Globally unique identifier for the host instance from the cloud provider.

* `start_timestamp` - The timestamp (in millisecond) when the Spark node is launched.

* `node_aws_attributes` - A `node_aws_attributes` block as defined below. Attributes specific to AWS for a Spark node.

* `host_private_ip` - The private IP address of the host instance.

---

A `node_aws_attributes` block exports the following:

* `is_spot` - Whether this node is on an Amazon spot instance.

---

A `aws_attributes` block exports the following:

* `first_on_demand` - The first `first_on_demand` nodes of the cluster will be placed on on-demand instances. If this value is greater than 0, the cluster driver node will be placed on an on-demand instance. If this value is greater than or equal to the current cluster size, all nodes will be placed on on-demand instances.

* `availability` - Availability type used for all subsequent nodes past the `first_on_demand` ones.

* `zone_id` - Identifier for the availability zone/datacenter in which the cluster resides.

* `instance_profile_arn` - Nodes for this cluster will only be placed on AWS instances with this instance profile.

* `spot_bid_price_percent` - The max price for AWS spot instances, as a percentage of the corresponding instance type’s on-demand price.

* `ebs_volume_type` - The type of EBS volumes that will be launched with this cluster.

* `ebs_volume_count` - The number of volumes launched for each instance.

* `ebs_volume_size` - The size of each EBS volume (in GiB) launched for each instance.

---

A `cluster_log_conf` block exports the following:

* `dbfs` - A `dbfs` block as defined below. DBFS location of cluster log.

* `s3` - A `s3` block as defined below. S3 location of cluster log.

---

A `init_scripts` block exports the following:

* `dbfs` - A `dbfs` block as defined below. DBFS location of init script.

* `s3` - A `s3` block as defined below. S3 location of init script.

---

A `dbfs` block exports the following:

* `destination` - DBFS destination, e.g. `dbfs:/my/path`.

---

A `s3` block exports the following:

* `destination` - S3 destination, e.g. `s3://my-bucket/some-prefix`.

* `region` - S3 region, e.g. `us-west-2`.

* `endpoint` - S3 endpoint, e.g. `https://s3-us-west-2.amazonaws.com`.

* `enable_encryption` - Wheter server side encryption is enabled.

* `encryption_type` - The encryption type, it could be `sse-s3` or `sse-kms`.

* `kms_key` - KMS key used if encryption is enabled and encryption type is set to `sse-kms`.

* `canned_acl` - Canned access control list, e.g. `bucket-owner-full-control`.

---

A `docker_image` block exports the following:

* `url` - URL for the Docker image.

* `basic_auth` - A `docker_basic_auth` block as defined below. Basic authentication information for Docker repository.

---

A `docker_basic_auth` block exports the following:

* `username` - User name for the Docker repository.

* `password` - Password for the Docker repository.

---

A `cluster_log_status` block exports the following:

* `last_attempted` - The timestamp of last attempt. If the last attempt fails, `last_exception` will contain the exception in the last attempt.

* `last_exception` - The exception thrown in the last attempt, it would be empty if there is no exception in last attempted.

---

A `termination_reason` block exports the following:

* `code` - Status code indicating why a cluster was terminated.

* `parameters` - List of parameters that provide additional information about why a cluster was terminated.
