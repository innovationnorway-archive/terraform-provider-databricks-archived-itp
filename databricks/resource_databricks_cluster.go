package databricks

import (
	"fmt"

	"github.com/Azure/go-autorest/autorest/to"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/innovationnorway/go-databricks/clusters"
)

func resourceDatabricksCluster() *schema.Resource {
	return &schema.Resource{
		Create: resourceDatabricksClusterCreate,
		Read:   resourceDatabricksClusterRead,
		Update: resourceDatabricksClusterUpdate,
		Delete: resourceDatabricksClusterDelete,

		Schema: map[string]*schema.Schema{
			"num_workers": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validation.IntBetween(1, 100000),
			},

			"autoscale": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"min_workers": {
							Type:         schema.TypeInt,
							Required:     true,
							ValidateFunc: validation.IntBetween(1, 100000),
						},
						"max_workers": {
							Type:         schema.TypeInt,
							Required:     true,
							ValidateFunc: validation.IntBetween(1, 100000),
						},
					},
				},
				ExactlyOneOf: []string{"num_workers", "autoscale"},
			},

			"spark_version": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},

			"node_type_id": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},

			"cluster_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"spark_conf": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"aws_attributes": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"first_on_demand": {
							Type:     schema.TypeInt,
							Optional: true,
						},

						"availability": {
							Type:     schema.TypeString,
							Optional: true,
							ValidateFunc: validation.StringInSlice([]string{
								string(clusters.ONDEMAND),
								string(clusters.SPOT),
								string(clusters.SPOTWITHFALLBACK),
							}, false),
						},

						"zone_id": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"instance_profile_arn": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"spot_bid_price_percent": {
							Type:     schema.TypeInt,
							Optional: true,
						},

						"ebs_volume_type": {
							Type:     schema.TypeString,
							Optional: true,
							ValidateFunc: validation.StringInSlice([]string{
								string(clusters.GENERALPURPOSESSD),
								string(clusters.THROUGHPUTOPTIMIZEDHDD),
							}, false),
						},

						"ebs_volume_count": {
							Type:     schema.TypeInt,
							Optional: true,
						},

						"ebs_volume_size": {
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},

			"driver_node_type_id": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},

			"ssh_public_keys": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"custom_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"cluster_log_conf": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dbfs": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"destination": {
										Type:     schema.TypeString,
										Required: true,
									},
								},
							},
						},

						"s3": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"destination": {
										Type:     schema.TypeString,
										Required: true,
									},

									"region": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"endpoint": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"enable_encryption": {
										Type:     schema.TypeBool,
										Optional: true,
									},

									"encryption_type": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"kms_key": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"canned_acl": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
							ExactlyOneOf: []string{
								"cluster_log_conf.0.dbfs",
								"cluster_log_conf.0.s3",
							},
						},
					},
				},
			},

			"init_scripts": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dbfs": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"destination": {
										Type:     schema.TypeString,
										Required: true,
									},
								},
							},
						},

						"s3": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"destination": {
										Type:     schema.TypeString,
										Required: true,
									},

									"region": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"endpoint": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"enable_encryption": {
										Type:     schema.TypeBool,
										Optional: true,
									},

									"encryption_type": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"kms_key": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"canned_acl": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
							ExactlyOneOf: []string{
								"init_scripts.0.dbfs",
								"init_scripts.0.s3",
							},
						},
					},
				},
			},

			"docker_image": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"url": {
							Type:     schema.TypeInt,
							Required: true,
						},

						"basic_auth": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"username": {
										Type:     schema.TypeString,
										Required: true,
									},

									"password": {
										Type:      schema.TypeString,
										Required:  true,
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
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"autotermination_minutes": {
				Type:         schema.TypeInt,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.IntBetween(0, 10000),
			},

			"enable_elastic_disk": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			"instance_pool_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"idempotency_token": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"cluster_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceDatabricksClusterCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Meta).Clusters
	ctx := meta.(*Meta).StopContext

	sparkVersion := d.Get("spark_version").(string)
	nodeTypeID := d.Get("node_type_id").(string)

	attributes := clusters.Attributes{
		SparkVersion: &sparkVersion,
		NodeTypeID:   &nodeTypeID,
	}

	if v, ok := d.GetOk("num_workers"); ok {
		attributes.NumWorkers = to.Int32Ptr(v.(int32))
	}

	if v, ok := d.GetOk("autoscale"); ok {
		attributes.Autoscale = expandClusterAutoscale(v.(*schema.Set).List())
	}

	if v, ok := d.GetOk("cluster_name"); ok {
		attributes.ClusterName = to.StringPtr(v.(string))
	}

	if v, ok := d.GetOk("spark_conf"); ok {
		attributes.SparkConf = expandClusterSparkConf(v.(map[string]interface{}))
	}

	if v, ok := d.GetOk("aws_attributes"); ok {
		attributes.AwsAttributes = expandClusterAwsAttributes(v.(*schema.Set).List())
	}

	if v, ok := d.GetOk("driver_node_type_id"); ok {
		attributes.DriverNodeTypeID = to.StringPtr(v.(string))
	}

	if v, ok := d.GetOk("ssh_public_keys"); ok {
		attributes.SSHPublicKeys = to.StringSlicePtr(v.([]string))
	}

	if v, ok := d.GetOk("custom_tags"); ok {
		attributes.CustomTags = expandClusterCustomTags(v.(map[string]interface{}))
	}

	if v, ok := d.GetOk("cluster_log_conf"); ok {
		attributes.ClusterLogConf = expandClusterLogConf(v.(*schema.Set).List())
	}

	if v, ok := d.GetOk("init_scripts"); ok {
		attributes.InitScripts = expandClusterInitScripts(v.(*schema.Set).List())
	}

	if v, ok := d.GetOk("docker_image"); ok {
		attributes.DockerImage = expandClusterDockerImage(v.(*schema.Set).List())
	}

	if v, ok := d.GetOk("spark_env_vars"); ok {
		attributes.SparkEnvVars = expandClusterSparkEnvVars(v.(map[string]interface{}))
	}

	if v, ok := d.GetOk("autotermination_minutes"); ok {
		attributes.AutoterminationMinutes = to.Int32Ptr(int32(v.(int)))
	}

	if v, ok := d.GetOk("enable_elastic_disk"); ok {
		attributes.EnableElasticDisk = to.BoolPtr(v.(bool))
	}

	if v, ok := d.GetOk("instance_pool_id"); ok {
		attributes.InstancePoolID = to.StringPtr(v.(string))
	}

	if v, ok := d.GetOk("idempotency_token"); ok {
		attributes.IdempotencyToken = to.StringPtr(v.(string))
	}

	resp, err := client.Create(ctx, attributes)
	if err != nil {
		return fmt.Errorf("unable to create cluster: %s", err)
	}

	d.SetId(*resp.ClusterID)

	return resourceDatabricksClusterRead(d, meta)
}

func resourceDatabricksClusterRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Meta).Clusters
	ctx := meta.(*Meta).StopContext

	resp, err := client.Get(ctx, d.Id())
	if err != nil {
		if resp.StatusCode == 400 {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("unable to get cluster: %s", err)
	}

	d.Set("num_workers", resp.NumWorkers)
	d.Set("autoscale", flattenAutoscale(resp.Autoscale))
	d.Set("cluster_name", resp.ClusterName)
	d.Set("spark_version", resp.SparkVersion)
	d.Set("spark_conf", resp.SparkConf)
	d.Set("aws_attributes", flattenAwsAttributes(resp.AwsAttributes))
	d.Set("node_type_id", resp.NodeTypeID)
	d.Set("driver_node_type_id", resp.DriverNodeTypeID)
	d.Set("ssh_public_keys", resp.SSHPublicKeys)
	d.Set("custom_tags", resp.CustomTags)
	d.Set("cluster_log_conf", flattenClusterLogConf(resp.ClusterLogConf))
	d.Set("init_scripts", flattenInitScripts(resp.InitScripts))
	d.Set("docker_image", flattenDockerImage(resp.DockerImage))
	d.Set("spark_env_vars", resp.SparkEnvVars)
	d.Set("autotermination_minutes", resp.AutoterminationMinutes)
	d.Set("enable_elastic_disk", resp.EnableElasticDisk)
	d.Set("instance_pool_id", resp.InstancePoolID)
	d.Set("cluster_id", resp.ClusterID)

	return nil
}

func resourceDatabricksClusterUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Meta).Clusters
	ctx := meta.(*Meta).StopContext

	clusterID := d.Id()
	sparkVersion := d.Get("spark_version").(string)
	nodeTypeID := d.Get("node_type_id").(string)

	attributes := clusters.EditAttributes{
		ClusterID:    &clusterID,
		SparkVersion: &sparkVersion,
		NodeTypeID:   &nodeTypeID,
	}

	if v, ok := d.GetOk("num_workers"); ok {
		attributes.NumWorkers = to.Int32Ptr(v.(int32))
	}

	if v, ok := d.GetOk("autoscale"); ok {
		attributes.Autoscale = expandClusterAutoscale(v.(*schema.Set).List())
	}

	if v, ok := d.GetOk("cluster_name"); ok {
		attributes.ClusterName = to.StringPtr(v.(string))
	}

	if v, ok := d.GetOk("spark_conf"); ok {
		attributes.SparkConf = expandClusterSparkConf(v.(map[string]interface{}))
	}

	if v, ok := d.GetOk("aws_attributes"); ok {
		attributes.AwsAttributes = expandClusterAwsAttributes(v.(*schema.Set).List())
	}

	if v, ok := d.GetOk("driver_node_type_id"); ok {
		attributes.DriverNodeTypeID = to.StringPtr(v.(string))
	}

	if v, ok := d.GetOk("ssh_public_keys"); ok {
		attributes.SSHPublicKeys = to.StringSlicePtr(v.([]string))
	}

	if v, ok := d.GetOk("custom_tags"); ok {
		attributes.CustomTags = expandClusterCustomTags(v.(map[string]interface{}))
	}

	if v, ok := d.GetOk("cluster_log_conf"); ok {
		attributes.ClusterLogConf = expandClusterLogConf(v.(*schema.Set).List())
	}

	if v, ok := d.GetOk("init_scripts"); ok {
		attributes.InitScripts = expandClusterInitScripts(v.(*schema.Set).List())
	}

	if v, ok := d.GetOk("docker_image"); ok {
		attributes.DockerImage = expandClusterDockerImage(v.(*schema.Set).List())
	}

	if v, ok := d.GetOk("spark_env_vars"); ok {
		attributes.SparkEnvVars = expandClusterSparkEnvVars(v.(map[string]interface{}))
	}

	if v, ok := d.GetOk("autotermination_minutes"); ok {
		attributes.AutoterminationMinutes = to.Int32Ptr(int32(v.(int)))
	}

	if v, ok := d.GetOk("enable_elastic_disk"); ok {
		attributes.EnableElasticDisk = to.BoolPtr(v.(bool))
	}

	if v, ok := d.GetOk("instance_pool_id"); ok {
		attributes.InstancePoolID = to.StringPtr(v.(string))
	}

	_, err := client.Edit(ctx, attributes)
	if err != nil {
		return fmt.Errorf("unable to update cluster: %s", err)
	}

	return resourceDatabricksClusterRead(d, meta)
}

func resourceDatabricksClusterDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Meta).Clusters
	ctx := meta.(*Meta).StopContext

	clusterID := d.Id()

	attributes := clusters.DeleteAttributes{
		ClusterID: &clusterID,
	}

	_, err := client.Delete(ctx, attributes)
	if err != nil {
		return fmt.Errorf("unable to delete cluster: %s", err)
	}

	d.SetId("")

	return nil
}

func expandClusterAutoscale(input []interface{}) *clusters.AutoScale {
	if len(input) == 0 {
		return nil
	}

	values := input[0].(map[string]interface{})

	result := &clusters.AutoScale{}

	if v, ok := values["min_workers"]; ok {
		result.MinWorkers = to.Int32Ptr(int32(v.(int)))
	}

	if v, ok := values["max_workers"]; ok {
		result.MaxWorkers = to.Int32Ptr(int32(v.(int)))
	}

	return result
}

func expandClusterSparkConf(input map[string]interface{}) map[string]*string {
	result := make(map[string]*string, len(input))

	for k, v := range input {
		result[k] = to.StringPtr(v.(string))
	}

	return result
}

func expandClusterAwsAttributes(input []interface{}) *clusters.AwsAttributes {
	if len(input) == 0 {
		return nil
	}

	values := input[0].(map[string]interface{})

	result := clusters.AwsAttributes{}

	if v, ok := values["first_on_demand"]; ok {
		firstOnDemand := v.(int32)
		result.FirstOnDemand = &firstOnDemand
	}

	if v, ok := values["availability"]; ok {
		availability := clusters.Availability(v.(string))
		result.Availability = availability
	}

	if v, ok := values["zone_id"]; ok {
		zoneID := v.(string)
		result.ZoneID = &zoneID
	}

	if v, ok := values["instance_profile_arn"]; ok {
		instanceProfileArn := v.(string)
		result.InstanceProfileArn = &instanceProfileArn
	}

	if v, ok := values["spot_bid_price_percent"]; ok {
		spotBidPricePercent := v.(int32)
		result.SpotBidPricePercent = &spotBidPricePercent
	}

	if v, ok := values["availability"]; ok {
		volumeType := clusters.EbsVolumeType(v.(string))
		result.EbsVolumeType = volumeType
	}

	if v, ok := values["ebs_volume_count"]; ok {
		ebsVolumeCount := v.(int32)
		result.EbsVolumeCount = &ebsVolumeCount
	}

	if v, ok := values["ebs_volume_size"]; ok {
		ebsVolumeSize := v.(int32)
		result.EbsVolumeSize = &ebsVolumeSize
	}

	return &result
}

func expandClusterCustomTags(input map[string]interface{}) map[string]*string {
	result := make(map[string]*string, len(input))

	for k, v := range input {
		result[k] = v.(*string)
	}

	return result
}

func expandClusterLogConf(input []interface{}) *clusters.LogConf {
	if len(input) == 0 {
		return nil
	}

	values := input[0].(map[string]interface{})

	result := clusters.LogConf{}

	if v, ok := values["dbfs"]; ok {
		storageInfo := expandClusterStorageInfoDbfs(v.([]interface{}))
		result.Dbfs = storageInfo
	}

	if v, ok := values["s3"]; ok {
		storageInfo := expandClusterStorageInfoS3(v.([]interface{}))
		result.S3 = storageInfo
	}

	return &result
}

func expandClusterInitScripts(input []interface{}) *[]clusters.InitScriptInfo {
	if len(input) == 0 {
		return nil
	}

	results := make([]clusters.InitScriptInfo, 0)

	for _, item := range input {
		values := item.(map[string]interface{})
		result := clusters.InitScriptInfo{}

		if v, ok := values["dbfs"]; ok {
			storageInfo := expandClusterStorageInfoDbfs(v.([]interface{}))
			result.Dbfs = storageInfo
		}

		if v, ok := values["s3"]; ok {
			storageInfo := expandClusterStorageInfoS3(v.([]interface{}))
			result.S3 = storageInfo
		}

		results = append(results, result)
	}

	return &results
}

func expandClusterStorageInfoDbfs(input []interface{}) *clusters.DbfsStorageInfo {
	if len(input) == 0 {
		return nil
	}

	values := input[0].(map[string]interface{})

	result := clusters.DbfsStorageInfo{}

	if v, ok := values["destination"]; ok {
		destination := v.(string)
		result.Destination = &destination
	}

	return &result
}

func expandClusterStorageInfoS3(input []interface{}) *clusters.S3StorageInfo {
	if len(input) == 0 {
		return nil
	}

	values := input[0].(map[string]interface{})

	result := clusters.S3StorageInfo{}

	if v, ok := values["destination"]; ok {
		destination := v.(string)
		result.Destination = &destination
	}

	if v, ok := values["region"]; ok {
		region := v.(string)
		result.Region = &region
	}

	if v, ok := values["endpoint"]; ok {
		endpoint := v.(string)
		result.Endpoint = &endpoint
	}

	if v, ok := values["enable_encryption"]; ok {
		enableEncryption := v.(bool)
		result.EnableEncryption = &enableEncryption
	}

	if v, ok := values["encryption_type"]; ok {
		encryptionType := v.(string)
		result.EncryptionType = &encryptionType
	}

	if v, ok := values["kms_key"]; ok {
		kmsKey := v.(string)
		result.KmsKey = &kmsKey
	}

	if v, ok := values["canned_acl"]; ok {
		cannedACL := v.(string)
		result.CannedACL = &cannedACL
	}

	return &result
}

func expandClusterDockerImage(input []interface{}) *clusters.DockerImage {
	if len(input) == 0 {
		return nil
	}

	values := input[0].(map[string]interface{})

	result := clusters.DockerImage{}

	if v, ok := values["url"]; ok {
		url := v.(string)
		result.URL = &url
	}

	if v, ok := values["basic_auth"]; ok {
		basicAuth := expandClusterDockerBasicAuth(v.([]interface{}))
		result.BasicAuth = basicAuth
	}

	return &result
}

func expandClusterDockerBasicAuth(input []interface{}) *clusters.DockerBasicAuth {
	if len(input) == 0 {
		return nil
	}

	values := input[0].(map[string]interface{})

	result := clusters.DockerBasicAuth{}

	if v, ok := values["username"]; ok {
		username := v.(string)
		result.Username = &username
	}

	if v, ok := values["password"]; ok {
		password := v.(string)
		result.Password = &password
	}

	return &result
}

func expandClusterSparkEnvVars(input map[string]interface{}) map[string]*string {
	result := make(map[string]*string, len(input))

	for k, v := range input {
		result[k] = v.(*string)
	}

	return result
}
