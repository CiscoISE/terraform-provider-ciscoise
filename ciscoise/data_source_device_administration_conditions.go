package ciscoise

import (
	"context"

	"ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDeviceAdministrationConditions() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDeviceAdministrationConditionsRead,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"item_name": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"condition_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_negate": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"link": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"href": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"rel": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"attribute_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"attribute_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"attribute_value": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"dictionary_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"dictionary_value": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"operator": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"children": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"condition_type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_negate": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"link": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"href": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"rel": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"type": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
						"dates_range": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"end_date": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"start_date": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"dates_range_exception": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"end_date": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"start_date": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"hours_range": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"end_time": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"start_time": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"hours_range_exception": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"end_time": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"start_time": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"week_days": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"week_days_exception": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"item_id": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"condition_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_negate": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"link": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"href": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"rel": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"attribute_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"attribute_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"attribute_value": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"dictionary_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"dictionary_value": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"operator": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"children": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"condition_type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_negate": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"link": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"href": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"rel": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"type": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
						"dates_range": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"end_date": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"start_date": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"dates_range_exception": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"end_date": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"start_date": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"hours_range": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"end_time": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"start_time": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"hours_range_exception": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"end_time": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"start_time": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"week_days": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"week_days_exception": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
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

						"condition_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_negate": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"link": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"href": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"rel": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"attribute_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"attribute_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"attribute_value": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"dictionary_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"dictionary_value": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"operator": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"children": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"condition_type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_negate": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"link": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"href": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"rel": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"type": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
						"dates_range": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"end_date": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"start_date": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"dates_range_exception": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"end_date": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"start_date": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"hours_range": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"end_time": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"start_time": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"hours_range_exception": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"end_time": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"start_time": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"week_days": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"week_days_exception": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceDeviceAdministrationConditionsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vName, okName := d.GetOk("name")
	vID, okID := d.GetOk("id")

	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)
	method3 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 3 %q", method3)

	selectedMethod := pickMethod([][]bool{method1, method2, method3})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetDeviceAdminConditions")

		response1, _, err := client.DeviceAdministrationConditions.GetDeviceAdminConditions()

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDeviceAdminConditions", err,
				"Failure at GetDeviceAdminConditions, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		vItems1 := flattenDeviceAdministrationConditionsGetDeviceAdminConditionsItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceAdminConditions response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: GetDeviceAdminConditionByName")
		vvName := vName.(string)

		response2, _, err := client.DeviceAdministrationConditions.GetDeviceAdminConditionByName(vvName)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDeviceAdminConditionByName", err,
				"Failure at GetDeviceAdminConditionByName, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

		vItemName2 := flattenDeviceAdministrationConditionsGetDeviceAdminConditionByNameItemName(response2.Response)
		if err := d.Set("item_name", vItemName2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceAdminConditionByName response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 3 {
		log.Printf("[DEBUG] Selected method 3: GetDeviceAdminConditionByID")
		vvID := vID.(string)

		response3, _, err := client.DeviceAdministrationConditions.GetDeviceAdminConditionByID(vvID)

		if err != nil || response3 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDeviceAdminConditionByID", err,
				"Failure at GetDeviceAdminConditionByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response3)

		vItemID3 := flattenDeviceAdministrationConditionsGetDeviceAdminConditionByIDItemID(response3.Response)
		if err := d.Set("item_id", vItemID3); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceAdminConditionByID response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDeviceAdministrationConditionsGetDeviceAdminConditionsItems(items *[]isegosdk.ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["condition_type"] = item.ConditionType
		respItem["is_negate"] = item.IsNegate
		respItem["link"] = flattenDeviceAdministrationConditionsGetDeviceAdminConditionsItemsLink(item.Link)
		respItem["description"] = item.Description
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItem["attribute_name"] = item.AttributeName
		respItem["attribute_id"] = item.AttributeID
		respItem["attribute_value"] = item.AttributeValue
		respItem["dictionary_name"] = item.DictionaryName
		respItem["dictionary_value"] = item.DictionaryValue
		respItem["operator"] = item.Operator
		respItem["children"] = flattenDeviceAdministrationConditionsGetDeviceAdminConditionsItemsChildren(item.Children)
		respItem["dates_range"] = flattenDeviceAdministrationConditionsGetDeviceAdminConditionsItemsDatesRange(item.DatesRange)
		respItem["dates_range_exception"] = flattenDeviceAdministrationConditionsGetDeviceAdminConditionsItemsDatesRangeException(item.DatesRangeException)
		respItem["hours_range"] = flattenDeviceAdministrationConditionsGetDeviceAdminConditionsItemsHoursRange(item.HoursRange)
		respItem["hours_range_exception"] = flattenDeviceAdministrationConditionsGetDeviceAdminConditionsItemsHoursRangeException(item.HoursRangeException)
		respItem["week_days"] = item.WeekDays
		respItem["week_days_exception"] = item.WeekDaysException
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceAdministrationConditionsGetDeviceAdminConditionsItemsLink(item *isegosdk.ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsResponseLink) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["href"] = item.Href
	respItem["rel"] = item.Rel
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceAdministrationConditionsGetDeviceAdminConditionsItemsChildren(items *[]isegosdk.ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsResponseChildren) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["condition_type"] = item.ConditionType
		respItem["is_negate"] = item.IsNegate
		respItem["link"] = flattenDeviceAdministrationConditionsGetDeviceAdminConditionsItemsChildrenLink(item.Link)
	}
	return respItems

}

func flattenDeviceAdministrationConditionsGetDeviceAdminConditionsItemsChildrenLink(item *isegosdk.ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsResponseChildrenLink) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["href"] = item.Href
	respItem["rel"] = item.Rel
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceAdministrationConditionsGetDeviceAdminConditionsItemsDatesRange(item *isegosdk.ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsResponseDatesRange) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["end_date"] = item.EndDate
	respItem["start_date"] = item.StartDate

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceAdministrationConditionsGetDeviceAdminConditionsItemsDatesRangeException(item *isegosdk.ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsResponseDatesRangeException) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["end_date"] = item.EndDate
	respItem["start_date"] = item.StartDate

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceAdministrationConditionsGetDeviceAdminConditionsItemsHoursRange(item *isegosdk.ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsResponseHoursRange) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["end_time"] = item.EndTime
	respItem["start_time"] = item.StartTime

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceAdministrationConditionsGetDeviceAdminConditionsItemsHoursRangeException(item *isegosdk.ResponseDeviceAdministrationConditionsGetDeviceAdminConditionsResponseHoursRangeException) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["end_time"] = item.EndTime
	respItem["start_time"] = item.StartTime

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceAdministrationConditionsGetDeviceAdminConditionByNameItemName(item *isegosdk.ResponseDeviceAdministrationConditionsGetDeviceAdminConditionByNameResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["condition_type"] = item.ConditionType
	respItem["is_negate"] = item.IsNegate
	respItem["link"] = flattenDeviceAdministrationConditionsGetDeviceAdminConditionByNameItemNameLink(item.Link)
	respItem["description"] = item.Description
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["attribute_name"] = item.AttributeName
	respItem["attribute_id"] = item.AttributeID
	respItem["attribute_value"] = item.AttributeValue
	respItem["dictionary_name"] = item.DictionaryName
	respItem["dictionary_value"] = item.DictionaryValue
	respItem["operator"] = item.Operator
	respItem["children"] = flattenDeviceAdministrationConditionsGetDeviceAdminConditionByNameItemNameChildren(item.Children)
	respItem["dates_range"] = flattenDeviceAdministrationConditionsGetDeviceAdminConditionByNameItemNameDatesRange(item.DatesRange)
	respItem["dates_range_exception"] = flattenDeviceAdministrationConditionsGetDeviceAdminConditionByNameItemNameDatesRangeException(item.DatesRangeException)
	respItem["hours_range"] = flattenDeviceAdministrationConditionsGetDeviceAdminConditionByNameItemNameHoursRange(item.HoursRange)
	respItem["hours_range_exception"] = flattenDeviceAdministrationConditionsGetDeviceAdminConditionByNameItemNameHoursRangeException(item.HoursRangeException)
	respItem["week_days"] = item.WeekDays
	respItem["week_days_exception"] = item.WeekDaysException
	return []map[string]interface{}{
		respItem,
	}
}

func flattenDeviceAdministrationConditionsGetDeviceAdminConditionByNameItemNameLink(item *isegosdk.ResponseDeviceAdministrationConditionsGetDeviceAdminConditionByNameResponseLink) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["href"] = item.Href
	respItem["rel"] = item.Rel
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceAdministrationConditionsGetDeviceAdminConditionByNameItemNameChildren(items *[]isegosdk.ResponseDeviceAdministrationConditionsGetDeviceAdminConditionByNameResponseChildren) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["condition_type"] = item.ConditionType
		respItem["is_negate"] = item.IsNegate
		respItem["link"] = flattenDeviceAdministrationConditionsGetDeviceAdminConditionByNameItemNameChildrenLink(item.Link)
	}
	return respItems

}

func flattenDeviceAdministrationConditionsGetDeviceAdminConditionByNameItemNameChildrenLink(item *isegosdk.ResponseDeviceAdministrationConditionsGetDeviceAdminConditionByNameResponseChildrenLink) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["href"] = item.Href
	respItem["rel"] = item.Rel
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceAdministrationConditionsGetDeviceAdminConditionByNameItemNameDatesRange(item *isegosdk.ResponseDeviceAdministrationConditionsGetDeviceAdminConditionByNameResponseDatesRange) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["end_date"] = item.EndDate
	respItem["start_date"] = item.StartDate

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceAdministrationConditionsGetDeviceAdminConditionByNameItemNameDatesRangeException(item *isegosdk.ResponseDeviceAdministrationConditionsGetDeviceAdminConditionByNameResponseDatesRangeException) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["end_date"] = item.EndDate
	respItem["start_date"] = item.StartDate

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceAdministrationConditionsGetDeviceAdminConditionByNameItemNameHoursRange(item *isegosdk.ResponseDeviceAdministrationConditionsGetDeviceAdminConditionByNameResponseHoursRange) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["end_time"] = item.EndTime
	respItem["start_time"] = item.StartTime

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceAdministrationConditionsGetDeviceAdminConditionByNameItemNameHoursRangeException(item *isegosdk.ResponseDeviceAdministrationConditionsGetDeviceAdminConditionByNameResponseHoursRangeException) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["end_time"] = item.EndTime
	respItem["start_time"] = item.StartTime

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceAdministrationConditionsGetDeviceAdminConditionByIDItemID(item *isegosdk.ResponseDeviceAdministrationConditionsGetDeviceAdminConditionByIDResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["condition_type"] = item.ConditionType
	respItem["is_negate"] = item.IsNegate
	respItem["link"] = flattenDeviceAdministrationConditionsGetDeviceAdminConditionByIDItemIDLink(item.Link)
	respItem["description"] = item.Description
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["attribute_name"] = item.AttributeName
	respItem["attribute_id"] = item.AttributeID
	respItem["attribute_value"] = item.AttributeValue
	respItem["dictionary_name"] = item.DictionaryName
	respItem["dictionary_value"] = item.DictionaryValue
	respItem["operator"] = item.Operator
	respItem["children"] = flattenDeviceAdministrationConditionsGetDeviceAdminConditionByIDItemIDChildren(item.Children)
	respItem["dates_range"] = flattenDeviceAdministrationConditionsGetDeviceAdminConditionByIDItemIDDatesRange(item.DatesRange)
	respItem["dates_range_exception"] = flattenDeviceAdministrationConditionsGetDeviceAdminConditionByIDItemIDDatesRangeException(item.DatesRangeException)
	respItem["hours_range"] = flattenDeviceAdministrationConditionsGetDeviceAdminConditionByIDItemIDHoursRange(item.HoursRange)
	respItem["hours_range_exception"] = flattenDeviceAdministrationConditionsGetDeviceAdminConditionByIDItemIDHoursRangeException(item.HoursRangeException)
	respItem["week_days"] = item.WeekDays
	respItem["week_days_exception"] = item.WeekDaysException
	return []map[string]interface{}{
		respItem,
	}
}

func flattenDeviceAdministrationConditionsGetDeviceAdminConditionByIDItemIDLink(item *isegosdk.ResponseDeviceAdministrationConditionsGetDeviceAdminConditionByIDResponseLink) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["href"] = item.Href
	respItem["rel"] = item.Rel
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceAdministrationConditionsGetDeviceAdminConditionByIDItemIDChildren(items *[]isegosdk.ResponseDeviceAdministrationConditionsGetDeviceAdminConditionByIDResponseChildren) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["condition_type"] = item.ConditionType
		respItem["is_negate"] = item.IsNegate
		respItem["link"] = flattenDeviceAdministrationConditionsGetDeviceAdminConditionByIDItemIDChildrenLink(item.Link)
	}
	return respItems

}

func flattenDeviceAdministrationConditionsGetDeviceAdminConditionByIDItemIDChildrenLink(item *isegosdk.ResponseDeviceAdministrationConditionsGetDeviceAdminConditionByIDResponseChildrenLink) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["href"] = item.Href
	respItem["rel"] = item.Rel
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceAdministrationConditionsGetDeviceAdminConditionByIDItemIDDatesRange(item *isegosdk.ResponseDeviceAdministrationConditionsGetDeviceAdminConditionByIDResponseDatesRange) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["end_date"] = item.EndDate
	respItem["start_date"] = item.StartDate

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceAdministrationConditionsGetDeviceAdminConditionByIDItemIDDatesRangeException(item *isegosdk.ResponseDeviceAdministrationConditionsGetDeviceAdminConditionByIDResponseDatesRangeException) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["end_date"] = item.EndDate
	respItem["start_date"] = item.StartDate

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceAdministrationConditionsGetDeviceAdminConditionByIDItemIDHoursRange(item *isegosdk.ResponseDeviceAdministrationConditionsGetDeviceAdminConditionByIDResponseHoursRange) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["end_time"] = item.EndTime
	respItem["start_time"] = item.StartTime

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceAdministrationConditionsGetDeviceAdminConditionByIDItemIDHoursRangeException(item *isegosdk.ResponseDeviceAdministrationConditionsGetDeviceAdminConditionByIDResponseHoursRangeException) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["end_time"] = item.EndTime
	respItem["start_time"] = item.StartTime

	return []map[string]interface{}{
		respItem,
	}

}
