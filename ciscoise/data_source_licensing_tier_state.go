package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceLicensingTierState() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Licensing.

- Get tier state information
`,

		ReadContext: dataSourceLicensingTierStateRead,
		Schema: map[string]*schema.Schema{
			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"compliance": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"consumption_counter": &schema.Schema{
							Description: `Compliance counter for tier`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"days_out_of_compliance": &schema.Schema{
							Description: `Number of days tier is out of compliance`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"last_authorization": &schema.Schema{
							Description: `Last date of authorization`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceLicensingTierStateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetTierStateInfo")

		response1, restyResp1, err := client.Licensing.GetTierStateInfo()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetTierStateInfo", err,
				"Failure at GetTierStateInfo, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenLicensingGetTierStateInfoItems(response1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTierStateInfo response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenLicensingGetTierStateInfoItems(items *isegosdk.ResponseLicensingGetTierStateInfo) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["compliance"] = item.Compliance
		respItem["consumption_counter"] = item.ConsumptionCounter
		respItem["days_out_of_compliance"] = item.DaysOutOfCompliance
		respItem["last_authorization"] = item.LastAuthorization
		respItem["name"] = item.Name
		respItem["status"] = item.Status
		respItems = append(respItems, respItem)
	}
	return respItems
}
