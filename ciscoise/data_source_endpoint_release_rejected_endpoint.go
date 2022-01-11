package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceEndpointReleaseRejectedEndpoint() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on endpoint.

- This data source action allows the client to release a rejected endpoint.
`,

		ReadContext: dataSourceEndpointReleaseRejectedEndpointRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter.`,
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

func dataSourceEndpointReleaseRejectedEndpointRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: ReleaseRejectedEndpoint")
		vvID := vID.(string)

		response1, err := client.Endpoint.ReleaseRejectedEndpoint(vvID)

		if err != nil || response1 == nil {
			if response1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", response1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing ReleaseRejectedEndpoint", err, response1.String(),
					"Failure at ReleaseRejectedEndpoint, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing ReleaseRejectedEndpoint", err,
				"Failure at ReleaseRejectedEndpoint, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %s", response1.String())

		if err := d.Set("item", response1.String()); err != nil {
			diags = append(diags, diagError(
				"Failure when setting ReleaseRejectedEndpoint response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}
