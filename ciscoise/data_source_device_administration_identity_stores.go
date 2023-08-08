package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDeviceAdministrationIDentityStores() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Device Administration - Identity Stores.

- Device Admin Return list of identity stores for authentication.
 (Other CRUD APIs available through ERS)
`,

		ReadContext: dataSourceDeviceAdministrationIDentityStoresRead,
		Schema: map[string]*schema.Schema{
			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceDeviceAdministrationIDentityStoresRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDeviceAdminIDentityStores")

		response1, restyResp1, err := client.DeviceAdministrationIDentityStores.GetDeviceAdminIDentityStores()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDeviceAdminIDentityStores", err,
				"Failure at GetDeviceAdminIDentityStores, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenDeviceAdministrationIDentityStoresGetDeviceAdminIDentityStoresItems(response1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceAdminIDentityStores response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDeviceAdministrationIDentityStoresGetDeviceAdminIDentityStoresItems(items *isegosdk.ResponseDeviceAdministrationIDentityStoresGetDeviceAdminIDentityStores) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItems = append(respItems, respItem)
	}
	return respItems
}
