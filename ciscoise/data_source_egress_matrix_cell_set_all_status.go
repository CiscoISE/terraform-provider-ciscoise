package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceEgressMatrixCellSetAllStatus() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on EgressMatrixCell.

- his API allows the client to set status of all the egress matrix cells.
`,

		ReadContext: dataSourceEgressMatrixCellSetAllStatusRead,
		Schema: map[string]*schema.Schema{
			"status": &schema.Schema{
				Description: `status path parameter.`,
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

func dataSourceEgressMatrixCellSetAllStatusRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vStatus := d.Get("status")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: SetAllCellsStatus")
		vvStatus := vStatus.(string)

		response1, err := client.EgressMatrixCell.SetAllCellsStatus(vvStatus)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing SetAllCellsStatus", err,
				"Failure at SetAllCellsStatus, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %s", response1.String())

		if err := d.Set("item", response1.String()); err != nil {
			diags = append(diags, diagError(
				"Failure when setting SetAllCellsStatus response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}
