package ciscoise

import (
	"context"

	"log"

	isegosdk "ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceSgMappingDeploy() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on IPToSGTMapping.

- This data source action allows the client to deploy an IP to SGT mapping by ID.
Only one Deploy process can run at any given time
`,

		ReadContext: dataSourceSgMappingDeployRead,
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

func dataSourceSgMappingDeployRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: DeployIPToSgtMappingByID")
		vvID := vID.(string)

		response1, err := client.IPToSgtMapping.DeployIPToSgtMappingByID(vvID)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing DeployIPToSgtMappingByID", err,
				"Failure at DeployIPToSgtMappingByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %s", response1.String())

		if err := d.Set("item", response1.String()); err != nil {
			diags = append(diags, diagError(
				"Failure when setting DeployIPToSgtMappingByID response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}
