package ciscoise

import (
	"context"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceSgMappingDeployStatusInfo() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSgMappingDeployStatusInfoRead,
		Schema: map[string]*schema.Schema{
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"result_value": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"value": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceSgMappingDeployStatusInfoRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetDeployStatusIPToSgtMapping")

		response1, _, err := client.IPToSgtMapping.GetDeployStatusIPToSgtMapping()

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDeployStatusIPToSgtMapping", err,
				"Failure at GetDeployStatusIPToSgtMapping, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		vItem1 := flattenIPToSgtMappingGetDeployStatusIPToSgtMappingItem(response1.OperationResult)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeployStatusIPToSgtMapping response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenIPToSgtMappingGetDeployStatusIPToSgtMappingItem(item *isegosdk.ResponseIPToSgtMappingGetDeployStatusIPToSgtMappingOperationResult) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["result_value"] = flattenIPToSgtMappingGetDeployStatusIPToSgtMappingItemResultValue(item.ResultValue)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenIPToSgtMappingGetDeployStatusIPToSgtMappingItemResultValue(items *[]isegosdk.ResponseIPToSgtMappingGetDeployStatusIPToSgtMappingOperationResultResultValue) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["value"] = item.Value
		respItem["name"] = item.Name
	}
	return respItems

}
