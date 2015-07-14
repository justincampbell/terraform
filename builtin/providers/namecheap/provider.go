package namecheap

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"user": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("NAMECHEAP_API_USER", nil),
				Description: "The username for a Namecheap account.",
			},

			"token": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("NAMECHEAP_API_TOKEN", nil),
				Description: "An API token associated with the API user.",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"namecheap_record": resourceNamcheapRecord(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func resourceNamcheapRecord() *schema.Resource {
	return nil
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		User:  d.Get("user").(string),
		Token: d.Get("token").(string),
	}

	return config.Client()
}
