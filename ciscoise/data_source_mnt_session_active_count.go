package ciscoise

import (
	"context"

	"ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceMntSessionActiveCount() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceMntSessionActiveCountRead,
		Schema: map[string]*schema.Schema{
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"count": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceMntSessionActiveCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetActiveCount")

		response1, _, err := client.Misc.GetActiveCount()

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetActiveCount", err,
				"Failure at GetActiveCount, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		vItem1 := flattenMiscGetActiveCountItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetActiveCount response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenMiscGetActiveCountItem(item *isegosdk.ResponseMiscGetActiveCount) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["count"] = item.Count
	return []map[string]interface{}{
		respItem,
	}
}