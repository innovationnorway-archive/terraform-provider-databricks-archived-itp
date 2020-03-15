package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/innovationnorway/terraform-provider-databricks/databricks"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: databricks.Provider})
}
