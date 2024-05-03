package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceConfiguration() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Configuration.

- Retrieve configuration information for LSD settings page
`,

		ReadContext: dataSourceConfigurationRead,
		Schema: map[string]*schema.Schema{
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"enable_epo": &schema.Schema{
							Description: `To enable/disable LSD ownership settings`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"enable_rcm": &schema.Schema{
							Description: `To enable/disable random mac(RCM) settings. Please note that this flag will be set to false if enableEPO flag is false`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceConfigurationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetConfiguration")

		response1, restyResp1, err := client.Configuration.GetConfiguration()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetConfiguration", err,
				"Failure at GetConfiguration, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenConfigurationGetConfigurationItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetConfiguration response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenConfigurationGetConfigurationItem(item *isegosdk.ResponseConfigurationGetConfiguration) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["enable_epo"] = boolPtrToString(item.EnableEPO)
	respItem["enable_rcm"] = boolPtrToString(item.EnableRCM)
	return []map[string]interface{}{
		respItem,
	}
}
