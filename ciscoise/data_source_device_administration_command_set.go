package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/kuba-mazurkiewicz/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDeviceAdministrationCommandSet() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Device Administration - Command Set.

- Device Admin Return list of command sets.
 (Other CRUD APIs available through ERS)
`,

		ReadContext: dataSourceDeviceAdministrationCommandSetRead,
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
							Description: `Command used in Device Admin authorization policies`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceDeviceAdministrationCommandSetRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDeviceAdminCommandSets")

		response1, restyResp1, err := client.DeviceAdministrationCommandSet.GetDeviceAdminCommandSets()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDeviceAdminCommandSets", err,
				"Failure at GetDeviceAdminCommandSets, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenDeviceAdministrationCommandSetGetDeviceAdminCommandSetsItems(response1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceAdminCommandSets response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDeviceAdministrationCommandSetGetDeviceAdminCommandSetsItems(items *isegosdk.ResponseDeviceAdministrationCommandSetGetDeviceAdminCommandSets) []map[string]interface{} {
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
