package ciscoise

import (
	"context"

	"reflect"

	"log"

	isegosdk "ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceNodePromotion() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on Node Deployment.

- Changes the cluster setting by promoting a node to primary when exceuted on standalone or secondary node.
It could also be used to convert a deployment node to standalone node.
`,

		ReadContext: dataSourceNodePromotionRead,
		Schema: map[string]*schema.Schema{
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"code": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"message": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"root_cause": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"promotion_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourceNodePromotionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: PromoteNode")
		request1 := expandRequestNodePromotionPromoteNode(ctx, "", d)

		response1, restyResp1, err := client.NodeDeployment.PromoteNode(request1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
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

func expandRequestNodePromotionPromoteNode(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNodeDeploymentPromoteNode {
	request := isegosdk.RequestNodeDeploymentPromoteNode{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".promotion_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".promotion_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".promotion_type")))) {
		request.PromotionType = interfaceToString(v)
	}
	return &request
}

func flattenNodeDeploymentPromoteNodeItem(item *isegosdk.ResponseNodeDeploymentPromoteNode) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["code"] = item.Code
	respItem["message"] = item.Message
	respItem["root_cause"] = item.RootCause
	return []map[string]interface{}{
		respItem,
	}
}
