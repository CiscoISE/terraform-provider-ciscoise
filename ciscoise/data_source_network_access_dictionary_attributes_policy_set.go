package ciscoise

import (
	"context"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkAccessDictionaryAttributesPolicySet() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Network Access - Dictionary Attributes List.

Network Access Returns list of dictionary attributes for policyset.`,

		ReadContext: dataSourceNetworkAccessDictionaryAttributesPolicySetRead,
		Schema: map[string]*schema.Schema{
			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"allowed_values": &schema.Schema{
							Description: `all of the allowed values for the dictionary attribute`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"is_default": &schema.Schema{
										Description: `true if this key value is the default between the allowed values of the dictionary attribute`,
										Type:        schema.TypeBool,
										Computed:    true,
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
							Description: `the data type for the dictionary attribute`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"description": &schema.Schema{
							Description: `The description of the Dictionary attribute`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"dictionary_name": &schema.Schema{
							Description: `the name of the dictionary which the dictionary attribute belongs to`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"direction_type": &schema.Schema{
							Description: `the direction for the useage of the dictionary attribute`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"id": &schema.Schema{
							Description: `Identifier for the dictionary attribute`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"internal_name": &schema.Schema{
							Description: `the internal name of the dictionary attribute`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"name": &schema.Schema{
							Description: `The dictionary attribute's name`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceNetworkAccessDictionaryAttributesPolicySetRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetNetworkAccessDictionariesPolicySet")

		response1, _, err := client.NetworkAccessDictionaryAttributesList.GetNetworkAccessDictionariesPolicySet()

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetNetworkAccessDictionariesPolicySet", err,
				"Failure at GetNetworkAccessDictionariesPolicySet, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		vItems1 := flattenNetworkAccessDictionaryAttributesListGetNetworkAccessDictionariesPolicySetItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNetworkAccessDictionariesPolicySet response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenNetworkAccessDictionaryAttributesListGetNetworkAccessDictionariesPolicySetItems(items *[]isegosdk.ResponseNetworkAccessDictionaryAttributesListGetNetworkAccessDictionariesPolicySetResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["allowed_values"] = flattenNetworkAccessDictionaryAttributesListGetNetworkAccessDictionariesPolicySetItemsAllowedValues(item.AllowedValues)
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

func flattenNetworkAccessDictionaryAttributesListGetNetworkAccessDictionariesPolicySetItemsAllowedValues(items *[]isegosdk.ResponseNetworkAccessDictionaryAttributesListGetNetworkAccessDictionariesPolicySetResponseAllowedValues) []map[string]interface{} {
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
