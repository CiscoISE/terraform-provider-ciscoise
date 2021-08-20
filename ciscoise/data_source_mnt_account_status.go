package ciscoise

import (
	"context"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceMntAccountStatus() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceMntAccountStatusRead,
		Schema: map[string]*schema.Schema{
			"mac": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"duration": &schema.Schema{
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

func dataSourceMntAccountStatusRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vMac := d.Get("mac")
	vDuration := d.Get("duration")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetAccountStatusByMac")
		vvMac := vMac.(string)
		vvDuration := vDuration.(string)

		response1, err := client.Misc.GetAccountStatusByMac(vvMac, vvDuration)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetAccountStatusByMac", err,
				"Failure at GetAccountStatusByMac, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		if err := d.Set("item", response1.String()); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAccountStatusByMac response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}
