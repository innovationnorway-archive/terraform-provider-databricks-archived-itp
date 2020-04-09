variable "databricks_host" {
  type        = string
  description = "A Databricks host. This is the URL of the Databricks instance. Example: https://<account>.cloud.databricks.com"
}

variable "databricks_token" {
  type        = string
  description = "A personal access token. This is used to access Databricks REST APIs."
}
