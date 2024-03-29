package ciscoise

import (
	"context"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceMntSessionByIP() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Misc.

- Sessions by Endpoint IP
`,

		ReadContext: dataSourceMntSessionByIPRead,
		Schema: map[string]*schema.Schema{
			"endpoint_ipv4": &schema.Schema{
				Description: `endpoint_ipv4 path parameter.`,
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

func dataSourceMntSessionByIPRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics
	vEndpointIPv4 := d.Get("endpoint_ipv4")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetSessionsByEndpointIP")
		vvEndpointIPv4 := vEndpointIPv4.(string)

		response1, err := client.Misc.GetSessionsByEndpointIP(vvEndpointIPv4)

		if err != nil || response1 == nil {
			if response1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", response1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing GetSessionsByEndpointIP", err, response1.String(),
					"Failure at GetSessionsByEndpointIP, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetSessionsByEndpointIP", err,
				"Failure at GetSessionsByEndpointIP, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %s", response1.String())

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
