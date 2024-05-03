package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePxGridDirect() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on pxGrid Direct.

- pxGrid Direct Get ALL connectorConfig information

- pxGrid Direct Get connectorConfig information based on ConnectorName
`,

		ReadContext: dataSourcePxGridDirectRead,
		Schema: map[string]*schema.Schema{
			"connector_name": &schema.Schema{
				Description: `connectorName path parameter. update or delete or retrieve the connector config.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"connector": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"additional_properties": {
										Type:     schema.TypeMap,
										Computed: true,
									},
									"attributes": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"attribute_mapping": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"dictionary_attribute": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"include_in_dictionary": {
																Type:     schema.TypeBool,
																Computed: true,
															},
															"json_attribute": {
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"correlation_identifier": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"top_level_object": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"unique_identifier": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"version_identifier": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"connector_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"connector_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"deltasync_schedule": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"interval": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"interval_unit": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"start_date": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"enabled": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"fullsync_schedule": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"interval": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"interval_unit": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"start_date": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"protocol": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"skip_certificate_validations": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"url": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"authentication_type": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"bulk_url": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"incremental_url": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"password": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"user_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"connector": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"additional_properties": {
										Type:     schema.TypeMap,
										Computed: true,
									},
									"attributes": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"attribute_mapping": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"dictionary_attribute": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"include_in_dictionary": {
																Type:     schema.TypeBool,
																Computed: true,
															},
															"json_attribute": {
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"correlation_identifier": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"top_level_object": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"unique_identifier": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"version_identifier": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"connector_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"connector_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"deltasync_schedule": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"interval": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"interval_unit": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"start_date": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"enabled": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"fullsync_schedule": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"interval": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"interval_unit": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"start_date": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"protocol": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"skip_certificate_validations": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"url": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"authentication_type": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"bulk_url": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"incremental_url": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"password": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"user_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourcePxGridDirectRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics
	vConnectorName, okConnectorName := d.GetOk("connector_name")

	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okConnectorName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetConnectorConfig")

		response1, restyResp1, err := client.PxGridDirect.GetConnectorConfig()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetConnectorConfig", err,
				"Failure at GetConnectorConfig, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenPxGridDirectGetConnectorConfigItemsResponse(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetConnectorConfig response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetConnectorConfigByConnectorName")
		vvConnectorName := vConnectorName.(string)

		response2, restyResp2, err := client.PxGridDirect.GetConnectorConfigByConnectorName(vvConnectorName)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetConnectorConfigByConnectorName", err,
				"Failure at GetConnectorConfigByConnectorName, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenPxGridDirectGetConnectorConfigByConnectorNameItemResponse(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetConnectorConfigByConnectorName response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenPxGridDirectGetConnectorConfigItemsResponse(items *[]isegosdk.ResponsePxGridDirectGetConnectorConfigResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["connector"] = flattenPxGridDirectGetConnectorConfigItemsResponseConnector(item.Connector)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPxGridDirectGetConnectorConfigItemsResponseConnector(item *isegosdk.ResponsePxGridDirectGetConnectorConfigResponseConnector) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["additional_properties"] = flattenPxGridDirectGetConnectorConfigItemsResponseConnectorAdditionalProperties(item.AdditionalProperties)
	respItem["attributes"] = flattenPxGridDirectGetConnectorConfigItemsResponseConnectorAttributes(item.Attributes)
	respItem["connector_name"] = item.ConnectorName
	respItem["connector_type"] = item.ConnectorType
	respItem["deltasync_schedule"] = flattenPxGridDirectGetConnectorConfigItemsResponseConnectorDeltasyncSchedule(item.DeltasyncSchedule)
	respItem["description"] = item.Description
	respItem["enabled"] = boolPtrToString(item.Enabled)
	respItem["fullsync_schedule"] = flattenPxGridDirectGetConnectorConfigItemsResponseConnectorFullsyncSchedule(item.FullsyncSchedule)
	respItem["protocol"] = item.Protocol
	respItem["skip_certificate_validations"] = boolPtrToString(item.SkipCertificateValidations)
	respItem["url"] = flattenPxGridDirectGetConnectorConfigItemsResponseConnectorURL(item.URL)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPxGridDirectGetConnectorConfigItemsResponseConnectorAdditionalProperties(item *isegosdk.ResponsePxGridDirectGetConnectorConfigResponseConnectorAdditionalProperties) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenPxGridDirectGetConnectorConfigItemsResponseConnectorAttributes(item *isegosdk.ResponsePxGridDirectGetConnectorConfigResponseConnectorAttributes) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["attribute_mapping"] = flattenPxGridDirectGetConnectorConfigItemsResponseConnectorAttributesAttributeMapping(item.AttributeMapping)
	respItem["correlation_identifier"] = item.CorrelationIDentifier
	respItem["top_level_object"] = item.TopLevelObject
	respItem["unique_identifier"] = item.UniqueIDentifier
	respItem["version_identifier"] = item.VersionIDentifier

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPxGridDirectGetConnectorConfigItemsResponseConnectorAttributesAttributeMapping(items *[]isegosdk.ResponsePxGridDirectGetConnectorConfigResponseConnectorAttributesAttributeMapping) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["dictionary_attribute"] = item.DictionaryAttribute
		respItem["include_in_dictionary"] = boolPtrToString(item.IncludeInDictionary)
		respItem["json_attribute"] = item.JSONAttribute
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPxGridDirectGetConnectorConfigItemsResponseConnectorDeltasyncSchedule(item *isegosdk.ResponsePxGridDirectGetConnectorConfigResponseConnectorDeltasyncSchedule) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["interval"] = item.Interval
	respItem["interval_unit"] = item.IntervalUnit
	respItem["start_date"] = item.StartDate

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPxGridDirectGetConnectorConfigItemsResponseConnectorFullsyncSchedule(item *isegosdk.ResponsePxGridDirectGetConnectorConfigResponseConnectorFullsyncSchedule) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["interval"] = item.Interval
	respItem["interval_unit"] = item.IntervalUnit
	respItem["start_date"] = item.StartDate

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPxGridDirectGetConnectorConfigItemsResponseConnectorURL(item *isegosdk.ResponsePxGridDirectGetConnectorConfigResponseConnectorURL) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["authentication_type"] = item.AuthenticationType
	respItem["bulk_url"] = item.BulkURL
	respItem["incremental_url"] = item.IncrementalURL
	respItem["password"] = item.Password
	respItem["user_name"] = item.UserName

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPxGridDirectGetConnectorConfigByConnectorNameItemResponse(item *isegosdk.ResponsePxGridDirectGetConnectorConfigByConnectorNameResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["connector"] = flattenPxGridDirectGetConnectorConfigByConnectorNameItemResponseConnector(item.Connector)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPxGridDirectGetConnectorConfigByConnectorNameItemResponseConnector(item *isegosdk.ResponsePxGridDirectGetConnectorConfigByConnectorNameResponseConnector) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["additional_properties"] = flattenPxGridDirectGetConnectorConfigByConnectorNameItemResponseConnectorAdditionalProperties(item.AdditionalProperties)
	respItem["attributes"] = flattenPxGridDirectGetConnectorConfigByConnectorNameItemResponseConnectorAttributes(item.Attributes)
	respItem["connector_name"] = item.ConnectorName
	respItem["connector_type"] = item.ConnectorType
	respItem["deltasync_schedule"] = flattenPxGridDirectGetConnectorConfigByConnectorNameItemResponseConnectorDeltasyncSchedule(item.DeltasyncSchedule)
	respItem["description"] = item.Description
	respItem["enabled"] = boolPtrToString(item.Enabled)
	respItem["fullsync_schedule"] = flattenPxGridDirectGetConnectorConfigByConnectorNameItemResponseConnectorFullsyncSchedule(item.FullsyncSchedule)
	respItem["protocol"] = item.Protocol
	respItem["skip_certificate_validations"] = boolPtrToString(item.SkipCertificateValidations)
	respItem["url"] = flattenPxGridDirectGetConnectorConfigByConnectorNameItemResponseConnectorURL(item.URL)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPxGridDirectGetConnectorConfigByConnectorNameItemResponseConnectorAdditionalProperties(item *isegosdk.ResponsePxGridDirectGetConnectorConfigByConnectorNameResponseConnectorAdditionalProperties) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenPxGridDirectGetConnectorConfigByConnectorNameItemResponseConnectorAttributes(item *isegosdk.ResponsePxGridDirectGetConnectorConfigByConnectorNameResponseConnectorAttributes) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["attribute_mapping"] = flattenPxGridDirectGetConnectorConfigByConnectorNameItemResponseConnectorAttributesAttributeMapping(item.AttributeMapping)
	respItem["correlation_identifier"] = item.CorrelationIDentifier
	respItem["top_level_object"] = item.TopLevelObject
	respItem["unique_identifier"] = item.UniqueIDentifier
	respItem["version_identifier"] = item.VersionIDentifier

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPxGridDirectGetConnectorConfigByConnectorNameItemResponseConnectorAttributesAttributeMapping(items *[]isegosdk.ResponsePxGridDirectGetConnectorConfigByConnectorNameResponseConnectorAttributesAttributeMapping) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["dictionary_attribute"] = item.DictionaryAttribute
		respItem["include_in_dictionary"] = boolPtrToString(item.IncludeInDictionary)
		respItem["json_attribute"] = item.JSONAttribute
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPxGridDirectGetConnectorConfigByConnectorNameItemResponseConnectorDeltasyncSchedule(item *isegosdk.ResponsePxGridDirectGetConnectorConfigByConnectorNameResponseConnectorDeltasyncSchedule) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["interval"] = item.Interval
	respItem["interval_unit"] = item.IntervalUnit
	respItem["start_date"] = item.StartDate

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPxGridDirectGetConnectorConfigByConnectorNameItemResponseConnectorFullsyncSchedule(item *isegosdk.ResponsePxGridDirectGetConnectorConfigByConnectorNameResponseConnectorFullsyncSchedule) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["interval"] = item.Interval
	respItem["interval_unit"] = item.IntervalUnit
	respItem["start_date"] = item.StartDate

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPxGridDirectGetConnectorConfigByConnectorNameItemResponseConnectorURL(item *isegosdk.ResponsePxGridDirectGetConnectorConfigByConnectorNameResponseConnectorURL) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["authentication_type"] = item.AuthenticationType
	respItem["bulk_url"] = item.BulkURL
	respItem["incremental_url"] = item.IncrementalURL
	respItem["password"] = item.Password
	respItem["user_name"] = item.UserName

	return []map[string]interface{}{
		respItem,
	}

}
