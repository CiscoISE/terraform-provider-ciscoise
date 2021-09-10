package ciscoise

import (
	"context"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Config is the configuration structure used to instantiate a
// new Cisco Identity Services Engine client.
type Config struct {
	BaseURL       string
	Username      string
	Password      string
	Debug         string
	SSLVerify     string
	UseAPIGateway string
	UseCSRFToken  string
}

// NewClient returns a new Cisco Identity Services Engine client.
func (c *Config) NewClient() (*isegosdk.Client, error) {
	return isegosdk.NewClientWithOptions(c.BaseURL,
		c.Username, c.Password,
		c.Debug, c.SSLVerify,
		c.UseAPIGateway, c.UseCSRFToken,
	)
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics

	config := Config{
		BaseURL:       d.Get("base_url").(string),
		Username:      d.Get("username").(string),
		Password:      d.Get("password").(string),
		Debug:         d.Get("debug").(string),
		SSLVerify:     d.Get("ssl_verify").(string),
		UseAPIGateway: d.Get("use_api_gateway").(string),
		UseCSRFToken:  d.Get("use_csrf_token").(string),
	}

	client, err := config.NewClient()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create Cisco Identity Services Engine client",
			Detail:   err.Error(),
		})
		return nil, diags
	}
	return client, diags
}
