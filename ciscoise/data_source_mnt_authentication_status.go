package ciscoise

import (
	"context"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceMntAuthenticationStatus() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Misc.

- AuthenticationStatus by MAC
`,

		ReadContext: dataSourceMntAuthenticationStatusRead,
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

func dataSourceMntAuthenticationStatusRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics
	vMAC := d.Get("mac")
	vSECONDS := d.Get("sec_ond_s")
	vRECORDS := d.Get("rec_ord_s")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetAuthenticationStatusByMac")
		vvMAC := vMAC.(string)
		vvSECONDS := vSECONDS.(string)
		vvRECORDS := vRECORDS.(string)

		response1, err := client.Misc.GetAuthenticationStatusByMac(vvMAC, vvSECONDS, vvRECORDS)

		if err != nil || response1 == nil {
			if response1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", response1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing GetAuthenticationStatusByMac", err, response1.String(),
					"Failure at GetAuthenticationStatusByMac, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetAuthenticationStatusByMac", err,
				"Failure at GetAuthenticationStatusByMac, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %s", response1.String())

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
