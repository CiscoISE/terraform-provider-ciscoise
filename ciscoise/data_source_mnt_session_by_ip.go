package ciscoise

import (
	"context"

	"ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceMntSessionByIP() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceMntSessionByIPRead,
		Schema: map[string]*schema.Schema{
			"endpoint_ipv4": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceMntSessionByIPRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vEndpointIPv4 := d.Get("endpoint_ipv4")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetSessionsByEndpointIP")
		vvEndpointIPv4 := vEndpointIPv4.(string)

		response1, err := client.Misc.GetSessionsByEndpointIP(vvEndpointIPv4)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetSessionsByEndpointIP", err,
				"Failure at GetSessionsByEndpointIP, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		if err := d.Set("item", response1.String()); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSessionsByEndpointIP response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}
