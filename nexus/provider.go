package nexus

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hibri/nexus"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"base_url": {
				Type:         schema.TypeString,
				Optional:     true,
				DefaultFunc:  schema.EnvDefaultFunc("NEXUS_BASE_URL", ""),
				Description:  descriptions["base_url"],

			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("NEXUS_USERNAME", nil),
				Description: descriptions["username"],
			},

			"password": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("NEXUS_PASSWORD", nil),
				Description: descriptions["password"],
			},

			"insecure": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: descriptions["insecure"],
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"nexus_iq_connection": IQConnection(),
		},
		ConfigureFunc: providerConfigure,
	}
}

var descriptions map[string]string

func init() {
	descriptions = map[string]string{
		"base_url": "The Nexus  URL",
		"username": "Username for a Nexus account",
		"password": "Password for the Nexus account",
		"insecure": "Disable SSL verification of API calls",
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	nexusClient := nexus.NexusClient{
		BaseUrl:  d.Get("base_url").(string),
		UserName: d.Get("username").(string),
		Password: d.Get("password").(string),
		Insecure: d.Get("insecure").(bool),
	}

	return nexusClient, nil
}
