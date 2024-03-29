package ciscoise

import (
	"context"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceMntSessionByNasIP() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Misc.

- Sessions by NAS IP
`,

		ReadContext: dataSourceMntSessionByNasIPRead,
		Schema: map[string]*schema.Schema{
			"nas_ipv4": &schema.Schema{
				Description: `nas_ipv4 path parameter.`,
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

func dataSourceMntSessionByNasIPRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics
	vNasIPv4 := d.Get("nas_ipv4")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetSessionsByNasIP")
		vvNasIPv4 := vNasIPv4.(string)

		response1, err := client.Misc.GetSessionsByNasIP(vvNasIPv4)

		if err != nil || response1 == nil {
			if response1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", response1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing GetSessionsByNasIP", err, response1.String(),
					"Failure at GetSessionsByNasIP, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetSessionsByNasIP", err,
				"Failure at GetSessionsByNasIP, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %s", response1.String())

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
