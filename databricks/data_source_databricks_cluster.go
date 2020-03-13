package databricks

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/innovationnorway/go-databricks/models"
	"github.com/innovationnorway/go-databricks/plumbing/clusters"
)

func dataSourceDatabricksCluster() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDatabricksClusterRead,

		Schema: map[string]*schema.Schema{
			"cluster_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			"num_workers": {
				Type:     schema.TypeInt,
				Computed: true,
			},

			"autoscale": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"min_workers": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"max_workers": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},

			"creator_user_name": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"driver": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"private_ip": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"public_dns": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"node_id": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"instance_id": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"start_timestamp": {
							Type:     schema.TypeInt,
							Computed: true,
						},

						"node_aws_attributes": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"is_spot": {
										Type:     schema.TypeBool,
										Computed: true,
									},
								},
							},
						},

						"host_private_ip": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			"executors": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"private_ip": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"public_dns": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"node_id": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"instance_id": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"start_timestamp": {
							Type:     schema.TypeInt,
							Computed: true,
						},

						"node_aws_attributes": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"is_spot": {
										Type:     schema.TypeBool,
										Computed: true,
									},
								},
							},
						},

						"host_private_ip": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			"cluster_name": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"spark_version": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"spark_conf": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"aws_attributes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"first_on_demand": {
							Type:     schema.TypeInt,
							Computed: true,
						},

						"availability": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"zone_id": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"instance_profile_arn": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"spot_bid_price_percent": {
							Type:     schema.TypeInt,
							Computed: true,
						},

						"ebs_volume_type": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"ebs_volume_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},

						"ebs_volume_size": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},

			"node_type_id": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"driver_node_type_id": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"ssh_public_keys": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"custom_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"cluster_log_conf": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dbfs": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"destination": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"s3": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"destination": {
										Type:     schema.TypeString,
										Computed: true,
									},

									"region": {
										Type:     schema.TypeString,
										Computed: true,
									},

									"endpoint": {
										Type:     schema.TypeString,
										Computed: true,
									},

									"enable_encryption": {
										Type:     schema.TypeBool,
										Computed: true,
									},

									"encryption_type": {
										Type:     schema.TypeString,
										Computed: true,
									},

									"kms_key": {
										Type:     schema.TypeString,
										Computed: true,
									},

									"canned_acl": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},

			"init_scripts": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dbfs": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"destination": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"s3": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"destination": {
										Type:     schema.TypeString,
										Computed: true,
									},

									"region": {
										Type:     schema.TypeString,
										Computed: true,
									},

									"endpoint": {
										Type:     schema.TypeString,
										Computed: true,
									},

									"enable_encryption": {
										Type:     schema.TypeBool,
										Computed: true,
									},

									"encryption_type": {
										Type:     schema.TypeString,
										Computed: true,
									},

									"kms_key": {
										Type:     schema.TypeString,
										Computed: true,
									},

									"canned_acl": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},

			"docker_image": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"url": {
							Type:     schema.TypeInt,
							Computed: true,
						},

						"basic_auth": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"username": {
										Type:     schema.TypeString,
										Computed: true,
									},

									"password": {
										Type:      schema.TypeString,
										Computed:  true,
										Sensitive: true,
									},
								},
							},
						},
					},
				},
			},

			"spark_env_vars": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"autotermination_minutes": {
				Type:     schema.TypeInt,
				Computed: true,
			},

			"enable_elastic_disk": {
				Type:     schema.TypeBool,
				Computed: true,
			},

			"instance_pool_id": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"cluster_source": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"state_message": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"start_time": {
				Type:     schema.TypeInt,
				Computed: true,
			},

			"terminated_time": {
				Type:     schema.TypeInt,
				Computed: true,
			},

			"last_state_loss_time": {
				Type:     schema.TypeInt,
				Computed: true,
			},

			"last_activity_time": {
				Type:     schema.TypeInt,
				Computed: true,
			},

			"cluster_memory_mb": {
				Type:     schema.TypeInt,
				Computed: true,
			},

			"cluster_cores": {
				Type:     schema.TypeFloat,
				Computed: true,
			},

			"default_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"cluster_log_status": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"last_attempted": {
							Type:     schema.TypeInt,
							Computed: true,
						},

						"last_exception": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			"termination_reason": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"code": {
							Type:     schema.TypeInt,
							Computed: true,
						},

						"parameters": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceDatabricksClusterRead(d *schema.ResourceData, meta interface{}) error {
	params := clusters.NewGetParams()
	params.ClusterID = d.Get("cluster_id").(string)

	m := meta.(*Meta)
	resp, err := m.Databricks.Clusters.Get(params, m.AuthInfo)
	if err != nil {
		return fmt.Errorf("unable to get cluster: %s", err)
	}

	d.Set("num_workers", resp.Payload.NumWorkers)
	d.Set("autoscale", flattenAutoscale(resp.Payload.Autoscale))
	d.Set("cluster_id", resp.Payload.ClusterID)
	d.Set("creator_user_name", resp.Payload.CreatorUserName)
	d.Set("driver", flattenDriver(resp.Payload.Driver))
	d.Set("executors", flattenExecutors(resp.Payload.Executors))
	d.Set("spark_context_id", resp.Payload.SparkContextID)
	d.Set("jdbc_port", resp.Payload.JdbcPort)
	d.Set("cluster_name", resp.Payload.ClusterName)
	d.Set("spark_version", resp.Payload.SparkVersion)
	d.Set("spark_conf", resp.Payload.SparkConf)
	d.Set("aws_attributes", flattenAwsAttributes(resp.Payload.AwsAttributes))
	d.Set("node_type_id", resp.Payload.NodeTypeID)
	d.Set("driver_node_type_id", resp.Payload.DriverNodeTypeID)
	d.Set("ssh_public_keys", resp.Payload.SSHPublicKeys)
	d.Set("custom_tags", resp.Payload.CustomTags)
	d.Set("cluster_log_conf", flattenClusterLogConf(resp.Payload.ClusterLogConf))
	d.Set("init_scripts", flattenInitScripts(resp.Payload.InitScripts))
	d.Set("docker_image", flattenDockerImage(resp.Payload.DockerImage))
	d.Set("spark_env_vars", resp.Payload.SparkEnvVars)
	d.Set("autotermination_minutes", resp.Payload.AutoterminationMinutes)
	d.Set("enable_elastic_disk", resp.Payload.EnableElasticDisk)
	d.Set("instance_pool_id", resp.Payload.InstancePoolID)
	d.Set("cluster_source", resp.Payload.ClusterSource)
	d.Set("state", resp.Payload.State)
	d.Set("state_message", resp.Payload.StateMessage)
	d.Set("start_time", resp.Payload.StartTime)
	d.Set("terminated_time", resp.Payload.TerminatedTime)
	d.Set("last_state_loss_time", resp.Payload.LastStateLossTime)
	d.Set("last_activity_time", resp.Payload.LastActivityTime)
	d.Set("cluster_memory_mb", resp.Payload.ClusterMemoryMb)
	d.Set("cluster_cores", resp.Payload.ClusterCores)
	d.Set("default_tags", resp.Payload.DefaultTags)
	d.Set("cluster_log_status", flattenClusterLogStatus(resp.Payload.ClusterLogStatus))
	d.Set("termination_reason", flattenTerminationReason(resp.Payload.TerminationReason))

	d.SetId(resp.Payload.ClusterID)

	return nil
}

func flattenAutoscale(input *models.AutoScale) []interface{} {
	if input == nil {
		return []interface{}{}
	}

	values := make(map[string]interface{})

	values["min_workers"] = input.MinWorkers
	values["max_workers"] = input.MaxWorkers

	return []interface{}{values}
}

func flattenExecutors(input []*models.SparkNode) []interface{} {
	result := make([]interface{}, 0)

	if input == nil {
		return []interface{}{}
	}

	for _, executor := range input {
		values := make(map[string]interface{})

		values["private_ip"] = executor.PrivateIP
		values["public_dns"] = executor.PublicDNS
		values["node_id"] = executor.NodeID
		values["instance_id"] = executor.InstanceID
		values["start_timestamp"] = executor.StartTimestamp
		values["node_aws_attributes"] = flattenNodeAwsAttributes(executor.NodeAwsAttributes)
		values["host_private_ip"] = executor.HostPrivateIP

		result = append(result, values)
	}

	return result
}

func flattenDriver(input *models.SparkNode) []interface{} {
	if input == nil {
		return []interface{}{}
	}

	values := make(map[string]interface{})

	values["private_ip"] = input.PrivateIP
	values["public_dns"] = input.PublicDNS
	values["node_id"] = input.NodeID
	values["instance_id"] = input.InstanceID
	values["start_timestamp"] = input.StartTimestamp
	values["node_aws_attributes"] = flattenNodeAwsAttributes(input.NodeAwsAttributes)
	values["host_private_ip"] = input.HostPrivateIP

	return []interface{}{values}
}

func flattenNodeAwsAttributes(input *models.SparkNodeAwsAttributes) []interface{} {
	if input == nil {
		return []interface{}{}
	}

	values := make(map[string]interface{})

	values["is_spot"] = input.IsSpot

	return []interface{}{values}
}

func flattenAwsAttributes(input *models.AwsAttributes) []interface{} {
	if input == nil {
		return []interface{}{}
	}

	values := make(map[string]interface{})

	values["first_on_demand"] = input.FirstOnDemand
	values["availability"] = input.Availability
	values["zone_id"] = input.ZoneID
	values["instance_profile_arn"] = input.InstanceProfileArn
	values["ebs_volume_type"] = input.EbsVolumeType
	values["ebs_volume_count"] = input.EbsVolumeCount
	values["ebs_volume_size"] = input.EbsVolumeSize

	return []interface{}{values}
}

func flattenClusterLogConf(input *models.ClusterLogConf) []interface{} {
	if input == nil {
		return []interface{}{}
	}

	values := make(map[string]interface{})

	values["dbfs"] = flattenStorageInfoDbfs(input.Dbfs)
	values["s3"] = flattenStorageInfoS3(input.S3)

	return []interface{}{values}
}

func flattenInitScripts(input []*models.InitScriptInfo) []interface{} {
	result := make([]interface{}, 0)

	if input == nil {
		return []interface{}{}
	}

	for _, item := range input {
		values := make(map[string]interface{})

		values["dbfs"] = flattenStorageInfoDbfs(item.Dbfs)
		values["s3"] = flattenStorageInfoS3(item.S3)

		result = append(result, values)
	}

	return result
}

func flattenStorageInfoDbfs(input *models.DbfsStorageInfo) []interface{} {
	if input == nil {
		return []interface{}{}
	}

	values := make(map[string]interface{})

	values["destination"] = input.Destination

	return []interface{}{values}
}

func flattenStorageInfoS3(input *models.S3StorageInfo) []interface{} {
	if input == nil {
		return []interface{}{}
	}

	values := make(map[string]interface{})

	values["destination"] = input.Destination
	values["region"] = input.Region
	values["endpoint"] = input.Endpoint
	values["enable_encryption"] = input.EnableEncryption
	values["encryption_type"] = input.EncryptionType
	values["kms_key"] = input.KmsKey
	values["canned_acl"] = input.CannedACL

	return []interface{}{values}
}

func flattenDockerImage(input *models.DockerImage) []interface{} {
	if input == nil {
		return []interface{}{}
	}

	values := make(map[string]interface{})

	values["url"] = input.URL
	values["basic_auth"] = flattenDockerBasicAuth(input.BasicAuth)

	return []interface{}{values}
}

func flattenDockerBasicAuth(input *models.DockerBasicAuth) []interface{} {
	if input == nil {
		return []interface{}{}
	}

	values := make(map[string]interface{})

	values["username"] = input.Username
	values["password"] = input.Password

	return []interface{}{values}
}

func flattenClusterLogStatus(input *models.LogSyncStatus) []interface{} {
	if input == nil {
		return []interface{}{}
	}

	values := make(map[string]interface{})

	values["last_attempted"] = input.LastAttempted
	values["last_exception"] = input.LastException

	return []interface{}{values}
}

func flattenTerminationReason(input *models.TerminationReason) []interface{} {
	if input == nil {
		return []interface{}{}
	}

	values := make(map[string]interface{})

	values["code"] = input.Code
	values["parameters"] = input.Parameters

	return []interface{}{values}
}
