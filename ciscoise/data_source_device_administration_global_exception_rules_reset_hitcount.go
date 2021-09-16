package ciscoise

import (
	"context"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceDeviceAdministrationGlobalExceptionRulesResetHitcount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Device Administration - Authorization Global Exception Rules.

- Device Admin Reset HitCount for Global Exceptions`,

		ReadContext: dataSourceDeviceAdministrationGlobalExceptionRulesResetHitcountRead,
		Schema: map[string]*schema.Schema{
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

func dataSourceDeviceAdministrationGlobalExceptionRulesResetHitcountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: ResetHitCountsDeviceAdminGlobalExceptions")

		response1, _, err := client.DeviceAdministrationAuthorizationGlobalExceptionRules.ResetHitCountsDeviceAdminGlobalExceptions()

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing ResetHitCountsDeviceAdminGlobalExceptions", err,
				"Failure at ResetHitCountsDeviceAdminGlobalExceptions, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		vItem1 := flattenDeviceAdministrationAuthorizationGlobalExceptionRulesResetHitCountsDeviceAdminGlobalExceptionsItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting ResetHitCountsDeviceAdminGlobalExceptions response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDeviceAdministrationAuthorizationGlobalExceptionRulesResetHitCountsDeviceAdminGlobalExceptionsItem(item *isegosdk.ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesResetHitCountsDeviceAdminGlobalExceptions) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["message"] = item.Message
	return []map[string]interface{}{
		respItem,
	}
}
