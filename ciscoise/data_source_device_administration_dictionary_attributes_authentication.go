package ciscoise

import (
	"context"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDeviceAdministrationDictionaryAttributesAuthentication() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDeviceAdministrationDictionaryAttributesAuthenticationRead,
		Schema: map[string]*schema.Schema{
			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"allowed_values": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"is_default": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"key": &schema.Schema{
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
						"data_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"dictionary_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"direction_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"internal_name": &schema.Schema{
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

func dataSourceDeviceAdministrationDictionaryAttributesAuthenticationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetDeviceAdminDictionariesAuthentication")

		response1, _, err := client.DeviceAdministrationDictionaryAttributesList.GetDeviceAdminDictionariesAuthentication()

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDeviceAdminDictionariesAuthentication", err,
				"Failure at GetDeviceAdminDictionariesAuthentication, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		vItems1 := flattenDeviceAdministrationDictionaryAttributesListGetDeviceAdminDictionariesAuthenticationItems(&response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceAdminDictionariesAuthentication response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDeviceAdministrationDictionaryAttributesListGetDeviceAdminDictionariesAuthenticationItems(items *[]isegosdk.ResponseDeviceAdministrationDictionaryAttributesListGetDeviceAdminDictionariesAuthenticationResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["allowed_values"] = flattenDeviceAdministrationDictionaryAttributesListGetDeviceAdminDictionariesAuthenticationItemsAllowedValues(item.AllowedValues)
		respItem["data_type"] = item.DataType
		respItem["description"] = item.Description
		respItem["dictionary_name"] = item.DictionaryName
		respItem["direction_type"] = item.DirectionType
		respItem["id"] = item.ID
		respItem["internal_name"] = item.InternalName
		respItem["name"] = item.Name
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceAdministrationDictionaryAttributesListGetDeviceAdminDictionariesAuthenticationItemsAllowedValues(items []isegosdk.ResponseDeviceAdministrationDictionaryAttributesListGetDeviceAdminDictionariesAuthenticationResponseAllowedValues) []map[string]interface{} {
	var respItems []map[string]interface{}
	for _, item := range items {
		respItem := make(map[string]interface{})
		respItem["is_default"] = item.IsDefault
		respItem["key"] = item.Key
		respItem["value"] = item.Value
	}
	return respItems

}
