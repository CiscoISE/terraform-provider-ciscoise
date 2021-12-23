package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceNodeSecondaryToPrimary() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Node Deployment.

- Execute this API in the secondary PAN in the cluster to promote the node to primary PAN.  Ensure that the API Gateway
setting is enabled in the secondary PAN.
 Approximate execution time 300 seconds.
`,

		ReadContext: dataSourceNodeSecondaryToPrimaryRead,
		Schema: map[string]*schema.Schema{
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"success": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"message": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"version": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceNodeSecondaryToPrimaryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: PromoteNode")

		response1, restyResp1, err := client.NodeDeployment.PromoteNode()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing PromoteNode", err,
				"Failure at PromoteNode, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenNodeDeploymentPromoteNodeItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting PromoteNode response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenNodeDeploymentPromoteNodeItem(item *isegosdk.ResponseNodeDeploymentPromoteNode) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["success"] = flattenNodeDeploymentPromoteNodeItemSuccess(item.Success)
	respItem["version"] = item.Version
	return []map[string]interface{}{
		respItem,
	}
}

func flattenNodeDeploymentPromoteNodeItemSuccess(item *isegosdk.ResponseNodeDeploymentPromoteNodeSuccess) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["message"] = item.Message

	return []map[string]interface{}{
		respItem,
	}

}
