package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/hibri/terraform-provider-nexus/nexus"
)

func main() {
	//client.Get()
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return nexus.Provider()
		},
	})
}
