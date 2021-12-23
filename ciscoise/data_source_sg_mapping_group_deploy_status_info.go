package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceSgMappingGroupDeployStatusInfo() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on IPToSGTMappingGroup.

- This data source action allows the client to get the IP to SGT mapping group deployment status.
Deploy Status will show last Deploy command output. The information will be saved until the next Deploy command
`,

		ReadContext: dataSourceSgMappingGroupDeployStatusInfoRead,
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

func dataSourceSgMappingGroupDeployStatusInfoRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDeployStatusIPToSgtMappingGroup")

		response1, restyResp1, err := client.IPToSgtMappingGroup.GetDeployStatusIPToSgtMappingGroup()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDeployStatusIPToSgtMappingGroup", err,
				"Failure at GetDeployStatusIPToSgtMappingGroup, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenIPToSgtMappingGroupGetDeployStatusIPToSgtMappingGroupItem(response1.OperationResult)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeployStatusIPToSgtMappingGroup response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenIPToSgtMappingGroupGetDeployStatusIPToSgtMappingGroupItem(item *isegosdk.ResponseIPToSgtMappingGroupGetDeployStatusIPToSgtMappingGroupOperationResult) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["result_value"] = flattenIPToSgtMappingGroupGetDeployStatusIPToSgtMappingGroupItemResultValue(item.ResultValue)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenIPToSgtMappingGroupGetDeployStatusIPToSgtMappingGroupItemResultValue(items *[]isegosdk.ResponseIPToSgtMappingGroupGetDeployStatusIPToSgtMappingGroupOperationResultResultValue) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["value"] = item.Value
		respItem["name"] = item.Name
		respItems = append(respItems, respItem)
	}
	return respItems
}
