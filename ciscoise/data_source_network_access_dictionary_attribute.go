package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkAccessDictionaryAttribute() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Network Access - Dictionary Attribute.

- Returns a list of Dictionary Attributes for an existing Dictionary.

- Get a Dictionary Attribute.
`,

		ReadContext: dataSourceNetworkAccessDictionaryAttributeRead,
		Schema: map[string]*schema.Schema{
			"dictionary_name": &schema.Schema{
				Description: `dictionaryName path parameter. the name of the dictionary the dictionary attribute belongs to`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"name": &schema.Schema{
				Description: `name path parameter. the dictionary attribute name`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"item": &schema.Schema{
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
										Type:        schema.TypeString,
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
										Type:        schema.TypeString,
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

func dataSourceNetworkAccessDictionaryAttributeRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics
	vDictionaryName, okDictionaryName := d.GetOk("dictionary_name")
	vName, okName := d.GetOk("name")

	method1 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okDictionaryName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetNetworkAccessDictionaryAttributesByDictionaryName")
		vvDictionaryName := vDictionaryName.(string)

		response1, restyResp1, err := client.NetworkAccessDictionaryAttribute.GetNetworkAccessDictionaryAttributesByDictionaryName(vvDictionaryName)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetNetworkAccessDictionaryAttributesByDictionaryName", err,
				"Failure at GetNetworkAccessDictionaryAttributesByDictionaryName, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

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
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkAccessDictionaryAttributeByName")
		vvName := vName.(string)
		vvDictionaryName := vDictionaryName.(string)

		response2, restyResp2, err := client.NetworkAccessDictionaryAttribute.GetNetworkAccessDictionaryAttributeByName(vvName, vvDictionaryName)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetNetworkAccessDictionaryAttributeByName", err,
				"Failure at GetNetworkAccessDictionaryAttributeByName, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

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
		respItem["is_default"] = boolPtrToString(item.IsDefault)
		respItem["key"] = item.Key
		respItem["value"] = item.Value
		respItems = append(respItems, respItem)
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
		respItem["is_default"] = boolPtrToString(item.IsDefault)
		respItem["key"] = item.Key
		respItem["value"] = item.Value
		respItems = append(respItems, respItem)
	}
	return respItems
}
