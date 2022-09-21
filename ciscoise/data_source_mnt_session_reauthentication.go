package ciscoise

import (
	"context"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceMntSessionReauthentication() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Misc.

- Session Reauthentication by MAC
`,

		ReadContext: dataSourceMntSessionReauthenticationRead,
		Schema: map[string]*schema.Schema{
			"end_poi_ntm_ac": &schema.Schema{
				Description: `ENDPOINT_MAC path parameter.`,
				Type:        schema.TypeString,
				Required:    true,
			},
			"psn_nam_e": &schema.Schema{
				Description: `PSN_NAME path parameter.`,
				Type:        schema.TypeString,
				Required:    true,
			},
			"rea_uth_typ_e": &schema.Schema{
				Description: `REAUTH_TYPE path parameter.`,
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

func dataSourceMntSessionReauthenticationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics
	vPSNNAME := d.Get("psn_nam_e")
	vENDPOINTMac := d.Get("end_poi_ntm_ac")
	vREAuthTYPE := d.Get("rea_uth_typ_e")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: SessionReauthenticationByMac")
		vvPSNNAME := vPSNNAME.(string)
		vvENDPOINTMac := vENDPOINTMac.(string)
		vvREAuthTYPE := vREAuthTYPE.(string)

		response1, err := client.Misc.SessionReauthenticationByMac(vvPSNNAME, vvENDPOINTMac, vvREAuthTYPE)

		if err != nil || response1 == nil {
			if response1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", response1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing SessionReauthenticationByMac", err, response1.String(),
					"Failure at SessionReauthenticationByMac, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing SessionReauthenticationByMac", err,
				"Failure at SessionReauthenticationByMac, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %s", response1.String())

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
