package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceSgMappingGroupDeploy() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on IPToSGTMappingGroup.

- This data source action allows the client to deploy an IP to SGT mapping group by ID.
Only one Deploy process can run at any given time
`,

		ReadContext: dataSourceSgMappingGroupDeployRead,
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

func dataSourceSgMappingGroupDeployRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: DeployIPToSgtMappingGroupByID")
		vvID := vID.(string)

		response1, err := client.IPToSgtMappingGroup.DeployIPToSgtMappingGroupByID(vvID)

		if err != nil || response1 == nil {
			if response1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", response1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing DeployIPToSgtMappingGroupByID", err, response1.String(),
					"Failure at DeployIPToSgtMappingGroupByID, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing DeployIPToSgtMappingGroupByID", err,
				"Failure at DeployIPToSgtMappingGroupByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %s", response1.String())

		if err := d.Set("item", response1.String()); err != nil {
			diags = append(diags, diagError(
				"Failure when setting DeployIPToSgtMappingGroupByID response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}
