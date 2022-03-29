package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"nsd_api_url": {
				Type:        schema.TypeString,
				DefaultFunc: schema.EnvDefaultFunc("NSD_API_URL", nil),
			},
			"nsd_api_username": {
				Type:        schema.TypeString,
				DefaultFunc: schema.EnvDefaultFunc("NSD_API_USERNAME", nil),
			},
			"nsd_api_password": {
				Type:        schema.TypeString,
				DefaultFunc: schema.EnvDefaultFunc("NSD_API_PASSWORD", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"nsd_firewall_opening": resourceNsdFirewallOpening(),
		},
		DataSourcesMap:       map[string]*schema.Resource{},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, data *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics

	server := data.Get("nsd_api_url").(string)
	username := data.Get("nsd_api_username").(string)
	password := data.Get("nsd_api_password").(string)
	if server != "" && username != "" && password != "" {
		return &NsdApiConfig{server: server, username: username, password: password}, diags
	} else {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "You need to specify 'nsd_api_url', 'nsd_api_username' and 'nsd_api_password'.",
			Detail:   "Specify the URL, the username and password to invoke the NSD API.",
		})
		return nil, diags
	}
}

type NsdApiConfig struct {
	server   string
	username string
	password string
}
