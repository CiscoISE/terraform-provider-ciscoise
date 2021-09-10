package ciscoise

import (
	"context"

	"reflect"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceNodePromotion() *schema.Resource {
	return &schema.Resource{
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

		response1, _, err := client.NodeDeployment.PromoteNode(request1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing PromoteNode", err,
				"Failure at PromoteNode, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

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
	if v, ok := d.GetOkExists(key + ".promotion_type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".promotion_type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".promotion_type"))) {
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
