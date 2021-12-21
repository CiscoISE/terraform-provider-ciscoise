package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceNetworkAccessLocalExceptionRulesResetHitcounts() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Network Access - Authorization Exception Rules.

- Network Access Reset HitCount for local exceptions
`,

		ReadContext: dataSourceNetworkAccessLocalExceptionRulesResetHitcountsRead,
		Schema: map[string]*schema.Schema{
			"policy_id": &schema.Schema{
				Description: `policyId path parameter. Policy id`,
				Type:        schema.TypeString,
				Required:    true,
			},
			"item": &schema.Schema{
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
		},
	}
}

func dataSourceNetworkAccessLocalExceptionRulesResetHitcountsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vPolicyID := d.Get("policy_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: ResetHitCountsNetworkAccessLocalExceptions")
		vvPolicyID := vPolicyID.(string)

		response1, restyResp1, err := client.NetworkAccessAuthorizationExceptionRules.ResetHitCountsNetworkAccessLocalExceptions(vvPolicyID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing ResetHitCountsNetworkAccessLocalExceptions", err,
				"Failure at ResetHitCountsNetworkAccessLocalExceptions, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenNetworkAccessAuthorizationExceptionRulesResetHitCountsNetworkAccessLocalExceptionsItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting ResetHitCountsNetworkAccessLocalExceptions response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenNetworkAccessAuthorizationExceptionRulesResetHitCountsNetworkAccessLocalExceptionsItem(item *isegosdk.ResponseNetworkAccessAuthorizationExceptionRulesResetHitCountsNetworkAccessLocalExceptions) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["message"] = item.Message
	return []map[string]interface{}{
		respItem,
	}
}
