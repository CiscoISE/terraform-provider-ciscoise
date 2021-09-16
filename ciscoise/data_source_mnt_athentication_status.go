package ciscoise

import (
	"context"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceMntAthenticationStatus() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Misc.

AuthenticationStatus by MAC`,

		ReadContext: dataSourceMntAthenticationStatusRead,
		Schema: map[string]*schema.Schema{
			"mac": &schema.Schema{
				Description: `MAC path parameter.`,
				Type:        schema.TypeString,
				Required:    true,
			},
			"rec_ord_s": &schema.Schema{
				Description: `RECORDS path parameter.`,
				Type:        schema.TypeString,
				Required:    true,
			},
			"sec_ond_s": &schema.Schema{
				Description: `SECONDS path parameter.`,
				Type:        schema.TypeString,
				Required:    true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceMntAthenticationStatusRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vMAC := d.Get("mac")
	vSECONDS := d.Get("sec_ond_s")
	vRECORDS := d.Get("rec_ord_s")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetAuthenticationStatusByMac")
		vvMAC := vMAC.(string)
		vvSECONDS := vSECONDS.(string)
		vvRECORDS := vRECORDS.(string)

		response1, err := client.Misc.GetAuthenticationStatusByMac(vvMAC, vvSECONDS, vvRECORDS)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetAuthenticationStatusByMac", err,
				"Failure at GetAuthenticationStatusByMac, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		if err := d.Set("item", response1.String()); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAuthenticationStatusByMac response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}
