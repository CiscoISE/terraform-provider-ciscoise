package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDataconnectSettings() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Dataconnect Services.

- This data source retrieves the status of the Dataconnect feature.
`,

		ReadContext: dataSourceDataconnectSettingsRead,
		Schema: map[string]*schema.Schema{
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"is_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_password_changed": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"password_expires_in_days": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"password_expires_on": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceDataconnectSettingsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDataconnectService")

		response1, restyResp1, err := client.DataconnectServices.GetDataconnectService()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetDataconnectService", err,
				"Failure at GetDataconnectService, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenDataconnectServicesGetDataconnectServiceItemResponse(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDataconnectService response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDataconnectServicesGetDataconnectServiceItemResponse(item *isegosdk.ResponseDataconnectServicesGetDataconnectServiceResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["is_enabled"] = boolPtrToString(item.IsEnabled)
	respItem["is_password_changed"] = boolPtrToString(item.IsPasswordChanged)
	respItem["password_expires_in_days"] = item.PasswordExpiresInDays
	respItem["password_expires_on"] = item.PasswordExpiresOn

	return []map[string]interface{}{
		respItem,
	}

}
