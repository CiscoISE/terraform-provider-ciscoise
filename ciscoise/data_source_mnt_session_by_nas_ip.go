package ciscoise

import (
	"context"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceMntSessionByNasIP() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceMntSessionByNasIPRead,
		Schema: map[string]*schema.Schema{
			"nas_ipv4": &schema.Schema{
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

func dataSourceMntSessionByNasIPRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vNasIPv4 := d.Get("nas_ipv4")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetSessionsByNasIP")
		vvNasIPv4 := vNasIPv4.(string)

		response1, err := client.Misc.GetSessionsByNasIP(vvNasIPv4)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetSessionsByNasIP", err,
				"Failure at GetSessionsByNasIP, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		if err := d.Set("item", response1.String()); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSessionsByNasIP response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}
