package ciscoise

import (
	"context"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceMntSessionDisconnect() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Misc.

- Session Disconnect
`,

		ReadContext: dataSourceMntSessionDisconnectRead,
		Schema: map[string]*schema.Schema{
			"dis_con_nec_tty_pe": &schema.Schema{
				Description: `DISCONNECT_TYPE path parameter.`,
				Type:        schema.TypeString,
				Required:    true,
			},
			"end_poi_nti_p": &schema.Schema{
				Description: `ENDPOINT_IP path parameter.`,
				Type:        schema.TypeString,
				Required:    true,
			},
			"mac": &schema.Schema{
				Description: `MAC path parameter.`,
				Type:        schema.TypeString,
				Required:    true,
			},
			"nas_ipv4": &schema.Schema{
				Description: `NAS_IPV4 path parameter.`,
				Type:        schema.TypeString,
				Required:    true,
			},
			"psn_nam_e": &schema.Schema{
				Description: `PSN_NAME path parameter.`,
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

func dataSourceMntSessionDisconnectRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics
	vENDPOINTIP := d.Get("end_poi_nti_p")
	vPSNNAME := d.Get("psn_nam_e")
	vMAC := d.Get("mac")
	vDISCONNECTTYPE := d.Get("dis_con_nec_tty_pe")
	vNASIPV4 := d.Get("nas_ipv4")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: SessionDisconnect")
		vvENDPOINTIP := vENDPOINTIP.(string)
		vvPSNNAME := vPSNNAME.(string)
		vvMAC := vMAC.(string)
		vvDISCONNECTTYPE := vDISCONNECTTYPE.(string)
		vvNASIPV4 := vNASIPV4.(string)

		response1, err := client.Misc.SessionDisconnect(vvENDPOINTIP, vvPSNNAME, vvMAC, vvDISCONNECTTYPE, vvNASIPV4)

		if err != nil || response1 == nil {
			if response1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", response1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing SessionDisconnect", err, response1.String(),
					"Failure at SessionDisconnect, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing SessionDisconnect", err,
				"Failure at SessionDisconnect, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %s", response1.String())

		if err := d.Set("item", response1.String()); err != nil {
			diags = append(diags, diagError(
				"Failure when setting SessionDisconnect response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}
