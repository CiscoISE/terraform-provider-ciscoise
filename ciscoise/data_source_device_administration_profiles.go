package ciscoise

import (
	"context"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDeviceAdministrationProfiles() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDeviceAdministrationProfilesRead,
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

func dataSourceDeviceAdministrationProfilesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetDeviceAdminProfiles")

		response1, _, err := client.DeviceAdministrationProfiles.GetDeviceAdminProfiles()

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDeviceAdminProfiles", err,
				"Failure at GetDeviceAdminProfiles, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		vItems1 := flattenDeviceAdministrationProfilesGetDeviceAdminProfilesItems(response1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceAdminProfiles response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDeviceAdministrationProfilesGetDeviceAdminProfilesItems(items *[]isegosdk.ResponseDeviceAdministrationProfilesGetDeviceAdminProfiles) []map[string]interface{} {
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
