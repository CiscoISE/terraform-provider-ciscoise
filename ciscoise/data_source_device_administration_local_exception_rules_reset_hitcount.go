package ciscoise

import (
	"context"

	"log"

	isegosdk "ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceDeviceAdministrationLocalExceptionRulesResetHitcount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Device Administration - Authorization Exception Rules.

- Device Admin Reset HitCount for local exceptions
`,

		ReadContext: dataSourceDeviceAdministrationLocalExceptionRulesResetHitcountRead,
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

func dataSourceDeviceAdministrationLocalExceptionRulesResetHitcountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vPolicyID := d.Get("policy_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: ResetHitCountsDeviceAdminLocalExceptions")
		vvPolicyID := vPolicyID.(string)

		response1, restyResp1, err := client.DeviceAdministrationAuthorizationExceptionRules.ResetHitCountsDeviceAdminLocalExceptions(vvPolicyID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing ResetHitCountsDeviceAdminLocalExceptions", err,
				"Failure at ResetHitCountsDeviceAdminLocalExceptions, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenDeviceAdministrationAuthorizationExceptionRulesResetHitCountsDeviceAdminLocalExceptionsItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting ResetHitCountsDeviceAdminLocalExceptions response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDeviceAdministrationAuthorizationExceptionRulesResetHitCountsDeviceAdminLocalExceptionsItem(item *isegosdk.ResponseDeviceAdministrationAuthorizationExceptionRulesResetHitCountsDeviceAdminLocalExceptions) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["message"] = item.Message
	return []map[string]interface{}{
		respItem,
	}
}
