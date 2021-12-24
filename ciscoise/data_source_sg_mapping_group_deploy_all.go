package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceSgMappingGroupDeployAll() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on IPToSGTMappingGroup.

- This data source action allows the client to deploy all the IP to SGT mapping groups.
Only one Deploy process can run at any given time
`,

		ReadContext: dataSourceSgMappingGroupDeployAllRead,
		Schema: map[string]*schema.Schema{
			"item": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceSgMappingGroupDeployAllRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: DeployAllIPToSgtMappingGroup")

		response1, err := client.IPToSgtMappingGroup.DeployAllIPToSgtMappingGroup()

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing DeployAllIPToSgtMappingGroup", err,
				"Failure at DeployAllIPToSgtMappingGroup, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %s", response1.String())

		if err := d.Set("item", response1.String()); err != nil {
			diags = append(diags, diagError(
				"Failure when setting DeployAllIPToSgtMappingGroup response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}
