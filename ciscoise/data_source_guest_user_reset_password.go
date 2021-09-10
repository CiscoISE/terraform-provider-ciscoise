package ciscoise

import (
	"context"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceGuestUserResetPassword() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGuestUserResetPasswordRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
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

func dataSourceGuestUserResetPasswordRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: ResetGuestUserPasswordByID")
		vvID := vID.(string)

		response1, _, err := client.GuestUser.ResetGuestUserPasswordByID(vvID)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing ResetGuestUserPasswordByID", err,
				"Failure at ResetGuestUserPasswordByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		vItem1 := flattenGuestUserResetGuestUserPasswordByIDItem(response1.OperationResult)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting ResetGuestUserPasswordByID response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenGuestUserResetGuestUserPasswordByIDItem(item *isegosdk.ResponseGuestUserResetGuestUserPasswordByIDOperationResult) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["result_value"] = flattenGuestUserResetGuestUserPasswordByIDItemResultValue(item.ResultValue)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenGuestUserResetGuestUserPasswordByIDItemResultValue(items *[]isegosdk.ResponseGuestUserResetGuestUserPasswordByIDOperationResultResultValue) []map[string]interface{} {
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
