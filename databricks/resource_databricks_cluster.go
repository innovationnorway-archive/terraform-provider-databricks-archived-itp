package databricks

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/innovationnorway/go-databricks/models"
	"github.com/innovationnorway/go-databricks/plumbing/clusters"
)

func resourceDatabricksCluster() *schema.Resource {
	return &schema.Resource{
		Create: resourceDatabricksClusterCreate,
		Read:   resourceDatabricksClusterRead,
		Update: resourceDatabricksClusterUpdate,
		Delete: resourceDatabricksClusterDelete,

		Schema: map[string]*schema.Schema{
			"num_workers": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			"autoscale": {
				Type:     schema.TypeSet,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"min_workers": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"max_workers": {
							Type:     schema.TypeInt,
							Required: true,
						},
					},
				},
				ConflictsWith: []string{"num_workers"},
			},

			"cluster_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"spark_version": {
				Type:     schema.TypeString,
				Required: true,
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
				Type:     schema.TypeSet,
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
								string(models.AwsAvailabilitySPOT),
								string(models.AwsAvailabilityONDEMAND),
								string(models.AwsAvailabilitySPOTWITHFALLBACK),
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
								string(models.EbsVolumeTypeGENERALPURPOSESSD),
								string(models.EbsVolumeTypeTHROUGHPUTOPTIMIZEDHDD),
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

			"node_type_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			"driver_node_type_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
				Type:     schema.TypeSet,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dbfs": {
							Type:     schema.TypeSet,
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
							Type:     schema.TypeSet,
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
							ConflictsWith: []string{"cluster_log_conf.dbfs"},
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
							Type:     schema.TypeSet,
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
							Type:     schema.TypeSet,
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
							ConflictsWith: []string{"init_scripts.dbfs"},
						},
					},
				},
			},

			"docker_image": {
				Type:     schema.TypeSet,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"url": {
							Type:     schema.TypeInt,
							Required: true,
						},

						"basic_auth": {
							Type:     schema.TypeSet,
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
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
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
	params := clusters.NewCreateParams()

	params.Body = &models.ClusterAttributes{
		SparkVersion: d.Get("spark_version").(string),
		NodeTypeID:   d.Get("node_type_id").(string),
	}

	if v, ok := d.GetOk("num_workers"); ok {
		params.Body.NumWorkers = int32(v.(int))
	}

	if v, ok := d.GetOk("autoscale"); ok {
		params.Body.Autoscale = expandAutoscale(v.(*schema.Set).List())
	}

	if v, ok := d.GetOk("cluster_name"); ok {
		params.Body.ClusterName = v.(string)
	}

	if v, ok := d.GetOk("spark_conf"); ok {
		params.Body.SparkConf = expandSparkConf(v.(map[string]interface{}))
	}

	if v, ok := d.GetOk("aws_attributes"); ok {
		params.Body.AwsAttributes = expandAwsAttributes(v.(*schema.Set).List())
	}

	if v, ok := d.GetOk("driver_node_type_id"); ok {
		params.Body.DriverNodeTypeID = v.(string)
	}

	if v, ok := d.GetOk("ssh_public_keys"); ok {
		params.Body.SSHPublicKeys = v.([]string)
	}

	if v, ok := d.GetOk("custom_tags"); ok {
		params.Body.CustomTags = expandCustomTags(v.(map[string]interface{}))
	}

	if v, ok := d.GetOk("cluster_log_conf"); ok {
		params.Body.ClusterLogConf = expandClusterLogConf(v.(*schema.Set).List())
	}

	if v, ok := d.GetOk("init_scripts"); ok {
		params.Body.InitScripts = expandInitScripts(v.(*schema.Set).List())
	}

	if v, ok := d.GetOk("docker_image"); ok {
		params.Body.DockerImage = expandDockerImage(v.(*schema.Set).List())
	}

	if v, ok := d.GetOk("spark_env_vars"); ok {
		params.Body.SparkEnvVars = expandSparkEnvVars(v.(map[string]interface{}))
	}

	if v, ok := d.GetOk("autotermination_minutes"); ok {
		params.Body.AutoterminationMinutes = int32(v.(int))
	}

	if v, ok := d.GetOk("enable_elastic_disk"); ok {
		params.Body.EnableElasticDisk = v.(bool)
	}

	if v, ok := d.GetOk("instance_pool_id"); ok {
		params.Body.InstancePoolID = v.(string)
	}

	if v, ok := d.GetOk("idempotency_token"); ok {
		params.Body.IdempotencyToken = v.(string)
	}

	m := meta.(*Meta)
	resp, err := m.Databricks.Clusters.Create(params, m.AuthInfo)
	if err != nil {
		return fmt.Errorf("unable to create cluster: %s", err)
	}

	d.SetId(resp.Payload.ClusterID)

	return resourceDatabricksClusterRead(d, meta)
}

func resourceDatabricksClusterRead(d *schema.ResourceData, meta interface{}) error {
	params := clusters.NewGetParams()
	params.ClusterID = d.Id()

	m := meta.(*Meta)
	resp, err := m.Databricks.Clusters.Get(params, m.AuthInfo)
	if err != nil {
		if v, ok := err.(*clusters.GetDefault); ok {
			if v.Payload.ErrorCode == models.ErrorCodeINVALIDPARAMETERVALUE {
				d.SetId("")
				return nil
			}
		}

		return fmt.Errorf("unable to get cluster: %s", err)
	}

	d.Set("num_workers", resp.Payload.NumWorkers)
	d.Set("autoscale", flattenAutoscale(resp.Payload.Autoscale))
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
	d.Set("cluster_id", resp.Payload.ClusterID)

	return nil
}

func resourceDatabricksClusterUpdate(d *schema.ResourceData, meta interface{}) error {
	params := clusters.NewEditParams()

	params.Body = clusters.EditBody{
		ClusterID:    d.Id(),
		SparkVersion: d.Get("spark_version").(string),
		NodeTypeID:   d.Get("node_type_id").(string),
	}

	if v, ok := d.GetOk("num_workers"); ok {
		params.Body.NumWorkers = int32(v.(int))
	}

	if v, ok := d.GetOk("autoscale"); ok {
		params.Body.Autoscale = expandAutoscale(v.(*schema.Set).List())
	}

	if v, ok := d.GetOk("cluster_name"); ok {
		params.Body.ClusterName = v.(string)
	}

	if v, ok := d.GetOk("spark_conf"); ok {
		params.Body.SparkConf = expandSparkConf(v.(map[string]interface{}))
	}

	if v, ok := d.GetOk("aws_attributes"); ok {
		params.Body.AwsAttributes = expandAwsAttributes(v.(*schema.Set).List())
	}

	if v, ok := d.GetOk("driver_node_type_id"); ok {
		params.Body.DriverNodeTypeID = v.(string)
	}

	if v, ok := d.GetOk("ssh_public_keys"); ok {
		params.Body.SSHPublicKeys = v.([]string)
	}

	if v, ok := d.GetOk("custom_tags"); ok {
		params.Body.CustomTags = expandCustomTags(v.(map[string]interface{}))
	}

	if v, ok := d.GetOk("cluster_log_conf"); ok {
		params.Body.ClusterLogConf = expandClusterLogConf(v.(*schema.Set).List())
	}

	if v, ok := d.GetOk("init_scripts"); ok {
		params.Body.InitScripts = expandInitScripts(v.(*schema.Set).List())
	}

	if v, ok := d.GetOk("docker_image"); ok {
		params.Body.DockerImage = expandDockerImage(v.(*schema.Set).List())
	}

	if v, ok := d.GetOk("spark_env_vars"); ok {
		params.Body.SparkEnvVars = expandSparkEnvVars(v.(map[string]interface{}))
	}

	if v, ok := d.GetOk("autotermination_minutes"); ok {
		params.Body.AutoterminationMinutes = int32(v.(int))
	}

	if v, ok := d.GetOk("enable_elastic_disk"); ok {
		params.Body.EnableElasticDisk = v.(bool)
	}

	if v, ok := d.GetOk("instance_pool_id"); ok {
		params.Body.InstancePoolID = v.(string)
	}

	m := meta.(*Meta)
	_, err := m.Databricks.Clusters.Edit(params, m.AuthInfo)
	if err != nil {
		return fmt.Errorf("unable to update cluster: %s", err)
	}

	return resourceDatabricksClusterRead(d, meta)
}

func resourceDatabricksClusterDelete(d *schema.ResourceData, meta interface{}) error {
	params := clusters.NewDeleteParams()
	params.Body = clusters.DeleteBody{
		ClusterID: d.Id(),
	}

	m := meta.(*Meta)
	_, err := m.Databricks.Clusters.Delete(params, m.AuthInfo)
	if err != nil {
		return fmt.Errorf("unable to delete cluster: %s", err)
	}

	d.SetId("")

	return nil
}

func expandAutoscale(input []interface{}) *models.AutoScale {
	if len(input) == 0 {
		return nil
	}

	values := input[0].(map[string]interface{})

	result := models.AutoScale{}

	result.MinWorkers = int32(values["min_workers"].(int))
	result.MaxWorkers = int32(values["max_workers"].(int))

	return &result
}

func expandSparkConf(input map[string]interface{}) models.SparkConfPair {
	result := make(map[string]string, len(input))

	for k, v := range input {
		result[k] = v.(string)
	}

	return result
}

func expandAwsAttributes(input []interface{}) *models.AwsAttributes {
	if len(input) == 0 {
		return nil
	}

	values := input[0].(map[string]interface{})

	result := models.AwsAttributes{}

	if v, ok := values["first_on_demand"]; ok {
		result.FirstOnDemand = int32(v.(int))
	}

	if v, ok := values["availability"]; ok {
		availability := models.AwsAvailability(v.(string))
		result.Availability = availability
	}

	if v, ok := values["zone_id"]; ok {
		result.ZoneID = v.(string)
	}

	if v, ok := values["instance_profile_arn"]; ok {
		result.InstanceProfileArn = v.(string)
	}

	if v, ok := values["spot_bid_price_percent"]; ok {
		result.SpotBidPricePercent = int32(v.(int))
	}

	if v, ok := values["availability"]; ok {
		volumeType := models.EbsVolumeType(v.(string))
		result.EbsVolumeType = volumeType
	}

	if v, ok := values["ebs_volume_count"]; ok {
		result.EbsVolumeCount = int32(v.(int))
	}

	if v, ok := values["ebs_volume_size"]; ok {
		result.EbsVolumeSize = int32(v.(int))
	}

	return &result
}

func expandCustomTags(input map[string]interface{}) models.ClusterTag {
	result := make(map[string]string, len(input))

	for k, v := range input {
		result[k] = v.(string)
	}

	return result
}

func expandClusterLogConf(input []interface{}) *models.ClusterLogConf {
	if len(input) == 0 {
		return nil
	}

	values := input[0].(map[string]interface{})

	result := models.ClusterLogConf{}

	if v, ok := values["dbfs"]; ok {
		storageInfo := expandStorageInfoDbfs(v.([]interface{}))
		result.Dbfs = storageInfo
	}

	if v, ok := values["s3"]; ok {
		storageInfo := expandStorageInfoS3(v.([]interface{}))
		result.S3 = storageInfo
	}

	return &result
}

func expandInitScripts(input []interface{}) []*models.InitScriptInfo {
	if len(input) == 0 {
		return nil
	}

	results := make([]*models.InitScriptInfo, 0)

	for _, item := range input {
		values := item.(map[string]interface{})
		result := models.InitScriptInfo{}

		if v, ok := values["dbfs"]; ok {
			storageInfo := expandStorageInfoDbfs(v.([]interface{}))
			result.Dbfs = storageInfo
		}

		if v, ok := values["s3"]; ok {
			storageInfo := expandStorageInfoS3(v.([]interface{}))
			result.S3 = storageInfo
		}

		results = append(results, &result)
	}

	return results
}

func expandStorageInfoDbfs(input []interface{}) *models.DbfsStorageInfo {
	if len(input) == 0 {
		return nil
	}

	values := input[0].(map[string]interface{})

	result := models.DbfsStorageInfo{}

	result.Destination = values["destination"].(string)

	return &result
}

func expandStorageInfoS3(input []interface{}) *models.S3StorageInfo {
	if len(input) == 0 {
		return nil
	}

	values := input[0].(map[string]interface{})

	result := models.S3StorageInfo{}

	result.Destination = values["destination"].(string)

	if v, ok := values["region"]; ok {
		result.Region = v.(string)
	}

	if v, ok := values["endpoint"]; ok {
		result.Endpoint = v.(string)
	}

	if v, ok := values["enable_encryption"]; ok {
		result.EnableEncryption = v.(bool)
	}

	if v, ok := values["encryption_type"]; ok {
		result.EncryptionType = v.(string)
	}

	if v, ok := values["kms_key"]; ok {
		result.KmsKey = v.(string)
	}

	if v, ok := values["canned_acl"]; ok {
		result.CannedACL = v.(string)
	}

	return &result
}

func expandDockerImage(input []interface{}) *models.DockerImage {
	if len(input) == 0 {
		return nil
	}

	values := input[0].(map[string]interface{})

	result := models.DockerImage{}

	result.URL = values["url"].(string)

	if v, ok := values["basic_auth"]; ok {
		basicAuth := expandDockerBasicAuth(v.([]interface{}))
		result.BasicAuth = basicAuth
	}

	return &result
}

func expandDockerBasicAuth(input []interface{}) *models.DockerBasicAuth {
	if len(input) == 0 {
		return nil
	}

	values := input[0].(map[string]interface{})

	result := models.DockerBasicAuth{}

	result.Username = values["username"].(string)
	result.Password = values["password"].(string)

	return &result
}

func expandSparkEnvVars(input map[string]interface{}) models.SparkEnvPair {
	result := make(map[string]string, len(input))

	for k, v := range input {
		result[k] = v.(string)
	}

	return result
}
