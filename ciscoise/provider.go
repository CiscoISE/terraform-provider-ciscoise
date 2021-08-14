package ciscoise

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider definition of schema(configuration), resources(CRUD) operations and dataSources(query)
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"base_url": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ISE_BASE_URL", nil),
			},
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("ISE_USERNAME", nil),
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("ISE_PASSWORD", nil),
			},
			"debug": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				DefaultFunc:  schema.EnvDefaultFunc("ISE_DEBUG", "false"),
				ValidateFunc: validateStringHasValueFunc([]string{"true", "false"}),
			},
			"ssl_verify": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Sensitive:    true,
				DefaultFunc:  schema.EnvDefaultFunc("ISE_SSL_VERIFY", "true"),
				ValidateFunc: validateStringHasValueFunc([]string{"true", "false"}),
			},
			"use_api_gateway": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				DefaultFunc:  schema.EnvDefaultFunc("ISE_USE_API_GATEWAY", "false"),
				ValidateFunc: validateStringHasValueFunc([]string{"true", "false"}),
			},
			"use_csrf_token": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				DefaultFunc:  schema.EnvDefaultFunc("ISE_USE_CSRF_TOKEN", "false"),
				ValidateFunc: validateStringHasValueFunc([]string{"true", "false"}),
			},
		},
		ResourcesMap: map[string]*schema.Resource{},
		DataSourcesMap: map[string]*schema.Resource{
			"ise_telemetry_info": dataSourceTelemetryInfo(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}
