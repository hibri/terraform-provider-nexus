package nexus

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hibri/nexus"
)

func IQConnection() *schema.Resource {
	return &schema.Resource{
		Create: resourceCreate,
		Read:   resourceRead,
		Update: resourceUpdate,
		Delete: resourceDelete,

		Schema: map[string]*schema.Schema{
			"enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"authentication_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"timeout_seconds": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default: 1,
			},


		},
	}
}

func resourceDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCreate(d *schema.ResourceData, m interface{}) error {
	//The iq connection is not created, it is empty and already exists
	// we are doing an update of it
	return resourceUpdate(d, m)
}

func resourceRead(d *schema.ResourceData, m interface{}) error {
	nexusClient := m.(nexus.NexusClient)
	iqConnectionClient := nexus.NewIQConnectionClient(nexusClient)
	iqConnection, err := iqConnectionClient.Get()

	if err != nil {
		d.Set("enabled", iqConnection.Enabled)
		d.Set("url", iqConnection.Url)
		d.Set("username", iqConnection.Username)
		d.Set("password", iqConnection.Password)
		d.Set("url", iqConnection.Url)
		d.Set("authentication_type", iqConnection.AuthenticationType)
		d.Set("timeout_seconds", iqConnection.TimeoutSeconds)
	}
	return nil
}

func resourceUpdate(d *schema.ResourceData, m interface{}) error {

	iqConnection := nexus.IQConnection{}
	url := d.Get("url").(string)

	iqConnection.Url = url
	iqConnection.Username = d.Get("username").(string)
	iqConnection.Password = d.Get("password").(string)
	iqConnection.AuthenticationType = d.Get("authentication_type").(string)
	iqConnection.Enabled = d.Get("enabled").(bool)
	iqConnection.TimeoutSeconds = d.Get("timeout_seconds").(int)

	nexusClient := m.(nexus.NexusClient)
	iqConnectionClient := nexus.NewIQConnectionClient(nexusClient)

	iqConnection, err := iqConnectionClient.Update(iqConnection)
	if err == nil {
		d.SetId(url)
		return resourceRead(d, m)
	}

	return err

}
