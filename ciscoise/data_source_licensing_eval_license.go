package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceLicensingEvalLicense() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Licensing.

- Get registration information
`,

		ReadContext: dataSourceLicensingEvalLicenseRead,
		Schema: map[string]*schema.Schema{
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"days_remaining": &schema.Schema{
							Description: `Number of eval license days left`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceLicensingEvalLicenseRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetEvalLicenseInfo")

		response1, restyResp1, err := client.Licensing.GetEvalLicenseInfo()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetEvalLicenseInfo", err,
				"Failure at GetEvalLicenseInfo, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenLicensingGetEvalLicenseInfoItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetEvalLicenseInfo response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenLicensingGetEvalLicenseInfoItem(item *isegosdk.ResponseLicensingGetEvalLicenseInfo) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["days_remaining"] = item.DaysRemaining
	return []map[string]interface{}{
		respItem,
	}
}
