package ciscoise

import (
	"context"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourcePxgridEndpointsByTypeInfo() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on MDM.

- 🚧 getEndpointsByType
`,

		ReadContext: dataSourcePxgridEndpointsByTypeInfoRead,
		Schema: map[string]*schema.Schema{
			"item": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourcePxgridEndpointsByTypeInfoRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetEndpointsByType")

		response1, err := client.Mdm.GetEndpointsByType()

		if err != nil || response1 == nil {
			if response1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", response1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing GetEndpointsByType", err, response1.String(),
					"Failure at GetEndpointsByType, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetEndpointsByType", err,
				"Failure at GetEndpointsByType, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %s", response1.String())

		if err := d.Set("item", response1.String()); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetEndpointsByType response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}
