package ciscoise

import (
	"context"

	"ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkAccessDictionaryAttribute() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceNetworkAccessDictionaryAttributeRead,
		Schema: map[string]*schema.Schema{
			"dictionary_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"item": &schema.Schema{
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

func dataSourceNetworkAccessDictionaryAttributeRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vDictionaryName, okDictionaryName := d.GetOk("dictionary_name")
	vName, okName := d.GetOk("name")

	method1 := []bool{okDictionaryName}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okName, okDictionaryName}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetNetworkAccessDictionaryAttributesByDictionaryName")
		vvDictionaryName := vDictionaryName.(string)

		response1, _, err := client.NetworkAccessDictionaryAttribute.GetNetworkAccessDictionaryAttributesByDictionaryName(vvDictionaryName)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetNetworkAccessDictionaryAttributesByDictionaryName", err,
				"Failure at GetNetworkAccessDictionaryAttributesByDictionaryName, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		vItems1 := flattenNetworkAccessDictionaryAttributeGetNetworkAccessDictionaryAttributesByDictionaryNameItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNetworkAccessDictionaryAttributesByDictionaryName response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: GetNetworkAccessDictionaryAttributeByName")
		vvName := vName.(string)
		vvDictionaryName := vDictionaryName.(string)

		response2, _, err := client.NetworkAccessDictionaryAttribute.GetNetworkAccessDictionaryAttributeByName(vvName, vvDictionaryName)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetNetworkAccessDictionaryAttributeByName", err,
				"Failure at GetNetworkAccessDictionaryAttributeByName, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

		vItem2 := flattenNetworkAccessDictionaryAttributeGetNetworkAccessDictionaryAttributeByNameItem(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNetworkAccessDictionaryAttributeByName response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenNetworkAccessDictionaryAttributeGetNetworkAccessDictionaryAttributesByDictionaryNameItems(items *[]isegosdk.ResponseNetworkAccessDictionaryAttributeGetNetworkAccessDictionaryAttributesByDictionaryNameResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["allowed_values"] = flattenNetworkAccessDictionaryAttributeGetNetworkAccessDictionaryAttributesByDictionaryNameItemsAllowedValues(item.AllowedValues)
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

func flattenNetworkAccessDictionaryAttributeGetNetworkAccessDictionaryAttributesByDictionaryNameItemsAllowedValues(items *[]isegosdk.ResponseNetworkAccessDictionaryAttributeGetNetworkAccessDictionaryAttributesByDictionaryNameResponseAllowedValues) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["is_default"] = item.IsDefault
		respItem["key"] = item.Key
		respItem["value"] = item.Value
	}
	return respItems

}

func flattenNetworkAccessDictionaryAttributeGetNetworkAccessDictionaryAttributeByNameItem(item *isegosdk.ResponseNetworkAccessDictionaryAttributeGetNetworkAccessDictionaryAttributeByNameResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["allowed_values"] = flattenNetworkAccessDictionaryAttributeGetNetworkAccessDictionaryAttributeByNameItemAllowedValues(item.AllowedValues)
	respItem["data_type"] = item.DataType
	respItem["description"] = item.Description
	respItem["dictionary_name"] = item.DictionaryName
	respItem["direction_type"] = item.DirectionType
	respItem["id"] = item.ID
	respItem["internal_name"] = item.InternalName
	respItem["name"] = item.Name
	return []map[string]interface{}{
		respItem,
	}
}

func flattenNetworkAccessDictionaryAttributeGetNetworkAccessDictionaryAttributeByNameItemAllowedValues(items *[]isegosdk.ResponseNetworkAccessDictionaryAttributeGetNetworkAccessDictionaryAttributeByNameResponseAllowedValues) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["is_default"] = item.IsDefault
		respItem["key"] = item.Key
		respItem["value"] = item.Value
	}
	return respItems

}
