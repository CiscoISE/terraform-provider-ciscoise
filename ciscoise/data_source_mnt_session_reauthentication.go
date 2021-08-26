package ciscoise

import (
	"context"

	"ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceMntSessionReauthentication() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceMntSessionReauthenticationRead,
		Schema: map[string]*schema.Schema{
			"psn_nam_e": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"end_poi_ntm_ac": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"rea_uth_typ_e": &schema.Schema{
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

func dataSourceMntSessionReauthenticationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vPSNNAME := d.Get("psn_nam_e")
	vENDPOINTMac := d.Get("end_poi_ntm_ac")
	vREAuthTYPE := d.Get("rea_uth_typ_e")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: SessionReauthenticationByMac")
		vvPSNNAME := vPSNNAME.(string)
		vvENDPOINTMac := vENDPOINTMac.(string)
		vvREAuthTYPE := vREAuthTYPE.(string)

		response1, err := client.Misc.SessionReauthenticationByMac(vvPSNNAME, vvENDPOINTMac, vvREAuthTYPE)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing SessionReauthenticationByMac", err,
				"Failure at SessionReauthenticationByMac, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		if err := d.Set("item", response1.String()); err != nil {
			diags = append(diags, diagError(
				"Failure when setting SessionReauthenticationByMac response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}