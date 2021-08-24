package ciscoise

import (
	"context"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDeviceAdministrationPolicySet() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDeviceAdministrationPolicySetRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"condition": &schema.Schema{
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
						"default": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"hit_counts": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_proxy": &schema.Schema{
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
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"rank": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"service_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"state": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"condition": &schema.Schema{
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
						"default": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"hit_counts": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_proxy": &schema.Schema{
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
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"rank": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"service_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"state": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceDeviceAdministrationPolicySetRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vID, okID := d.GetOk("id")

	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetDeviceAdminPolicySets")

		response1, _, err := client.DeviceAdministrationPolicySet.GetDeviceAdminPolicySets()

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDeviceAdminPolicySets", err,
				"Failure at GetDeviceAdminPolicySets, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		vItems1 := flattenDeviceAdministrationPolicySetGetDeviceAdminPolicySetsItems(&response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceAdminPolicySets response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: GetDeviceAdminPolicySetByID")
		vvID := vID.(string)

		response2, _, err := client.DeviceAdministrationPolicySet.GetDeviceAdminPolicySetByID(vvID)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDeviceAdminPolicySetByID", err,
				"Failure at GetDeviceAdminPolicySetByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

		vItem2 := flattenDeviceAdministrationPolicySetGetDeviceAdminPolicySetByIDItem(&response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceAdminPolicySetByID response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDeviceAdministrationPolicySetGetDeviceAdminPolicySetsItems(items *[]isegosdk.ResponseDeviceAdministrationPolicySetGetDeviceAdminPolicySetsResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["condition"] = flattenDeviceAdministrationPolicySetGetDeviceAdminPolicySetsItemsCondition(item.Condition)
		respItem["default"] = item.Default
		respItem["description"] = item.Description
		respItem["hit_counts"] = item.HitCounts
		respItem["id"] = item.ID
		respItem["is_proxy"] = item.IsProxy
		respItem["link"] = flattenDeviceAdministrationPolicySetGetDeviceAdminPolicySetsItemsLink(item.Link)
		respItem["name"] = item.Name
		respItem["rank"] = item.Rank
		respItem["service_name"] = item.ServiceName
		respItem["state"] = item.State
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceAdministrationPolicySetGetDeviceAdminPolicySetsItemsCondition(item isegosdk.ResponseDeviceAdministrationPolicySetGetDeviceAdminPolicySetsResponseCondition) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["condition_type"] = item.ConditionType
	respItem["is_negate"] = item.IsNegate
	respItem["link"] = flattenDeviceAdministrationPolicySetGetDeviceAdminPolicySetsItemsConditionLink(item.Link)
	respItem["description"] = item.Description
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["attribute_name"] = item.AttributeName
	respItem["attribute_id"] = item.AttributeID
	respItem["attribute_value"] = item.AttributeValue
	respItem["dictionary_name"] = item.DictionaryName
	respItem["dictionary_value"] = item.DictionaryValue
	respItem["operator"] = item.Operator
	respItem["children"] = flattenDeviceAdministrationPolicySetGetDeviceAdminPolicySetsItemsConditionChildren(item.Children)
	respItem["dates_range"] = flattenDeviceAdministrationPolicySetGetDeviceAdminPolicySetsItemsConditionDatesRange(item.DatesRange)
	respItem["dates_range_exception"] = flattenDeviceAdministrationPolicySetGetDeviceAdminPolicySetsItemsConditionDatesRangeException(item.DatesRangeException)
	respItem["hours_range"] = flattenDeviceAdministrationPolicySetGetDeviceAdminPolicySetsItemsConditionHoursRange(item.HoursRange)
	respItem["hours_range_exception"] = flattenDeviceAdministrationPolicySetGetDeviceAdminPolicySetsItemsConditionHoursRangeException(item.HoursRangeException)
	respItem["week_days"] = item.WeekDays
	respItem["week_days_exception"] = item.WeekDaysException

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceAdministrationPolicySetGetDeviceAdminPolicySetsItemsConditionLink(item isegosdk.ResponseDeviceAdministrationPolicySetGetDeviceAdminPolicySetsResponseConditionLink) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["href"] = item.Href
	respItem["rel"] = item.Rel
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceAdministrationPolicySetGetDeviceAdminPolicySetsItemsConditionChildren(items []isegosdk.ResponseDeviceAdministrationPolicySetGetDeviceAdminPolicySetsResponseConditionChildren) []map[string]interface{} {
	var respItems []map[string]interface{}
	for _, item := range items {
		respItem := make(map[string]interface{})
		respItem["condition_type"] = item.ConditionType
		respItem["is_negate"] = item.IsNegate
		respItem["link"] = flattenDeviceAdministrationPolicySetGetDeviceAdminPolicySetsItemsConditionChildrenLink(item.Link)
	}
	return respItems

}

func flattenDeviceAdministrationPolicySetGetDeviceAdminPolicySetsItemsConditionChildrenLink(item isegosdk.ResponseDeviceAdministrationPolicySetGetDeviceAdminPolicySetsResponseConditionChildrenLink) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["href"] = item.Href
	respItem["rel"] = item.Rel
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceAdministrationPolicySetGetDeviceAdminPolicySetsItemsConditionDatesRange(item isegosdk.ResponseDeviceAdministrationPolicySetGetDeviceAdminPolicySetsResponseConditionDatesRange) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["end_date"] = item.EndDate
	respItem["start_date"] = item.StartDate

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceAdministrationPolicySetGetDeviceAdminPolicySetsItemsConditionDatesRangeException(item isegosdk.ResponseDeviceAdministrationPolicySetGetDeviceAdminPolicySetsResponseConditionDatesRangeException) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["end_date"] = item.EndDate
	respItem["start_date"] = item.StartDate

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceAdministrationPolicySetGetDeviceAdminPolicySetsItemsConditionHoursRange(item isegosdk.ResponseDeviceAdministrationPolicySetGetDeviceAdminPolicySetsResponseConditionHoursRange) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["end_time"] = item.EndTime
	respItem["start_time"] = item.StartTime

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceAdministrationPolicySetGetDeviceAdminPolicySetsItemsConditionHoursRangeException(item isegosdk.ResponseDeviceAdministrationPolicySetGetDeviceAdminPolicySetsResponseConditionHoursRangeException) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["end_time"] = item.EndTime
	respItem["start_time"] = item.StartTime

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceAdministrationPolicySetGetDeviceAdminPolicySetsItemsLink(item isegosdk.ResponseDeviceAdministrationPolicySetGetDeviceAdminPolicySetsResponseLink) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["href"] = item.Href
	respItem["rel"] = item.Rel
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceAdministrationPolicySetGetDeviceAdminPolicySetByIDItem(item *isegosdk.ResponseDeviceAdministrationPolicySetGetDeviceAdminPolicySetByIDResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["condition"] = flattenDeviceAdministrationPolicySetGetDeviceAdminPolicySetByIDItemCondition(item.Condition)
	respItem["default"] = item.Default
	respItem["description"] = item.Description
	respItem["hit_counts"] = item.HitCounts
	respItem["id"] = item.ID
	respItem["is_proxy"] = item.IsProxy
	respItem["link"] = flattenDeviceAdministrationPolicySetGetDeviceAdminPolicySetByIDItemLink(item.Link)
	respItem["name"] = item.Name
	respItem["rank"] = item.Rank
	respItem["service_name"] = item.ServiceName
	respItem["state"] = item.State
	return []map[string]interface{}{
		respItem,
	}
}

func flattenDeviceAdministrationPolicySetGetDeviceAdminPolicySetByIDItemCondition(item isegosdk.ResponseDeviceAdministrationPolicySetGetDeviceAdminPolicySetByIDResponseCondition) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["condition_type"] = item.ConditionType
	respItem["is_negate"] = item.IsNegate
	respItem["link"] = flattenDeviceAdministrationPolicySetGetDeviceAdminPolicySetByIDItemConditionLink(item.Link)
	respItem["description"] = item.Description
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["attribute_name"] = item.AttributeName
	respItem["attribute_id"] = item.AttributeID
	respItem["attribute_value"] = item.AttributeValue
	respItem["dictionary_name"] = item.DictionaryName
	respItem["dictionary_value"] = item.DictionaryValue
	respItem["operator"] = item.Operator
	respItem["children"] = flattenDeviceAdministrationPolicySetGetDeviceAdminPolicySetByIDItemConditionChildren(item.Children)
	respItem["dates_range"] = flattenDeviceAdministrationPolicySetGetDeviceAdminPolicySetByIDItemConditionDatesRange(item.DatesRange)
	respItem["dates_range_exception"] = flattenDeviceAdministrationPolicySetGetDeviceAdminPolicySetByIDItemConditionDatesRangeException(item.DatesRangeException)
	respItem["hours_range"] = flattenDeviceAdministrationPolicySetGetDeviceAdminPolicySetByIDItemConditionHoursRange(item.HoursRange)
	respItem["hours_range_exception"] = flattenDeviceAdministrationPolicySetGetDeviceAdminPolicySetByIDItemConditionHoursRangeException(item.HoursRangeException)
	respItem["week_days"] = item.WeekDays
	respItem["week_days_exception"] = item.WeekDaysException

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceAdministrationPolicySetGetDeviceAdminPolicySetByIDItemConditionLink(item isegosdk.ResponseDeviceAdministrationPolicySetGetDeviceAdminPolicySetByIDResponseConditionLink) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["href"] = item.Href
	respItem["rel"] = item.Rel
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceAdministrationPolicySetGetDeviceAdminPolicySetByIDItemConditionChildren(items []isegosdk.ResponseDeviceAdministrationPolicySetGetDeviceAdminPolicySetByIDResponseConditionChildren) []map[string]interface{} {
	var respItems []map[string]interface{}
	for _, item := range items {
		respItem := make(map[string]interface{})
		respItem["condition_type"] = item.ConditionType
		respItem["is_negate"] = item.IsNegate
		respItem["link"] = flattenDeviceAdministrationPolicySetGetDeviceAdminPolicySetByIDItemConditionChildrenLink(item.Link)
	}
	return respItems

}

func flattenDeviceAdministrationPolicySetGetDeviceAdminPolicySetByIDItemConditionChildrenLink(item isegosdk.ResponseDeviceAdministrationPolicySetGetDeviceAdminPolicySetByIDResponseConditionChildrenLink) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["href"] = item.Href
	respItem["rel"] = item.Rel
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceAdministrationPolicySetGetDeviceAdminPolicySetByIDItemConditionDatesRange(item isegosdk.ResponseDeviceAdministrationPolicySetGetDeviceAdminPolicySetByIDResponseConditionDatesRange) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["end_date"] = item.EndDate
	respItem["start_date"] = item.StartDate

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceAdministrationPolicySetGetDeviceAdminPolicySetByIDItemConditionDatesRangeException(item isegosdk.ResponseDeviceAdministrationPolicySetGetDeviceAdminPolicySetByIDResponseConditionDatesRangeException) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["end_date"] = item.EndDate
	respItem["start_date"] = item.StartDate

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceAdministrationPolicySetGetDeviceAdminPolicySetByIDItemConditionHoursRange(item isegosdk.ResponseDeviceAdministrationPolicySetGetDeviceAdminPolicySetByIDResponseConditionHoursRange) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["end_time"] = item.EndTime
	respItem["start_time"] = item.StartTime

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceAdministrationPolicySetGetDeviceAdminPolicySetByIDItemConditionHoursRangeException(item isegosdk.ResponseDeviceAdministrationPolicySetGetDeviceAdminPolicySetByIDResponseConditionHoursRangeException) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["end_time"] = item.EndTime
	respItem["start_time"] = item.StartTime

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceAdministrationPolicySetGetDeviceAdminPolicySetByIDItemLink(item isegosdk.ResponseDeviceAdministrationPolicySetGetDeviceAdminPolicySetByIDResponseLink) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["href"] = item.Href
	respItem["rel"] = item.Rel
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}
