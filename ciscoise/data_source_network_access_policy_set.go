package ciscoise

import (
	"context"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkAccessPolicySet() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceNetworkAccessPolicySetRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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

									"attribute_id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"attribute_name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"attribute_value": &schema.Schema{
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
									"condition_type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
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
									"description": &schema.Schema{
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
									"id": &schema.Schema{
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
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"operator": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
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

									"attribute_id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"attribute_name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"attribute_value": &schema.Schema{
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
									"condition_type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
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
									"description": &schema.Schema{
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
									"id": &schema.Schema{
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
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"operator": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
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

func dataSourceNetworkAccessPolicySetRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vID, okID := d.GetOk("id")

	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetNetworkAccessPolicySets")

		response1, _, err := client.NetworkAccessPolicySet.GetNetworkAccessPolicySets()

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetNetworkAccessPolicySets", err,
				"Failure at GetNetworkAccessPolicySets, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		vItems1 := flattenNetworkAccessPolicySetGetNetworkAccessPolicySetsItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNetworkAccessPolicySets response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: GetNetworkAccessPolicySetByID")
		vvID := vID.(string)

		response2, _, err := client.NetworkAccessPolicySet.GetNetworkAccessPolicySetByID(vvID)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetNetworkAccessPolicySetByID", err,
				"Failure at GetNetworkAccessPolicySetByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

		vItem2 := flattenNetworkAccessPolicySetGetNetworkAccessPolicySetByIDItem(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNetworkAccessPolicySetByID response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenNetworkAccessPolicySetGetNetworkAccessPolicySetsItems(items *[]isegosdk.ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetsResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["condition"] = flattenNetworkAccessPolicySetGetNetworkAccessPolicySetsItemsCondition(item.Condition)
		respItem["default"] = item.Default
		respItem["description"] = item.Description
		respItem["hit_counts"] = item.HitCounts
		respItem["id"] = item.ID
		respItem["is_proxy"] = item.IsProxy
		respItem["link"] = flattenNetworkAccessPolicySetGetNetworkAccessPolicySetsItemsLink(item.Link)
		respItem["name"] = item.Name
		respItem["rank"] = item.Rank
		respItem["service_name"] = item.ServiceName
		respItem["state"] = item.State
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenNetworkAccessPolicySetGetNetworkAccessPolicySetsItemsCondition(item *isegosdk.ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetsResponseCondition) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["condition_type"] = item.ConditionType
	respItem["is_negate"] = item.IsNegate
	respItem["link"] = flattenNetworkAccessPolicySetGetNetworkAccessPolicySetsItemsConditionLink(item.Link)
	respItem["description"] = item.Description
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["attribute_name"] = item.AttributeName
	respItem["attribute_id"] = item.AttributeID
	respItem["attribute_value"] = item.AttributeValue
	respItem["dictionary_name"] = item.DictionaryName
	respItem["dictionary_value"] = item.DictionaryValue
	respItem["operator"] = item.Operator
	respItem["children"] = flattenNetworkAccessPolicySetGetNetworkAccessPolicySetsItemsConditionChildren(item.Children)
	respItem["dates_range"] = flattenNetworkAccessPolicySetGetNetworkAccessPolicySetsItemsConditionDatesRange(item.DatesRange)
	respItem["dates_range_exception"] = flattenNetworkAccessPolicySetGetNetworkAccessPolicySetsItemsConditionDatesRangeException(item.DatesRangeException)
	respItem["hours_range"] = flattenNetworkAccessPolicySetGetNetworkAccessPolicySetsItemsConditionHoursRange(item.HoursRange)
	respItem["hours_range_exception"] = flattenNetworkAccessPolicySetGetNetworkAccessPolicySetsItemsConditionHoursRangeException(item.HoursRangeException)
	respItem["week_days"] = item.WeekDays
	respItem["week_days_exception"] = item.WeekDaysException

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNetworkAccessPolicySetGetNetworkAccessPolicySetsItemsConditionLink(item *isegosdk.ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetsResponseConditionLink) []map[string]interface{} {
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

func flattenNetworkAccessPolicySetGetNetworkAccessPolicySetsItemsConditionChildren(items *[]isegosdk.ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetsResponseConditionChildren) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["condition_type"] = item.ConditionType
		respItem["is_negate"] = item.IsNegate
		respItem["link"] = flattenNetworkAccessPolicySetGetNetworkAccessPolicySetsItemsConditionChildrenLink(item.Link)
	}
	return respItems

}

func flattenNetworkAccessPolicySetGetNetworkAccessPolicySetsItemsConditionChildrenLink(item *isegosdk.ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetsResponseConditionChildrenLink) []map[string]interface{} {
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

func flattenNetworkAccessPolicySetGetNetworkAccessPolicySetsItemsConditionDatesRange(item *isegosdk.ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetsResponseConditionDatesRange) []map[string]interface{} {
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

func flattenNetworkAccessPolicySetGetNetworkAccessPolicySetsItemsConditionDatesRangeException(item *isegosdk.ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetsResponseConditionDatesRangeException) []map[string]interface{} {
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

func flattenNetworkAccessPolicySetGetNetworkAccessPolicySetsItemsConditionHoursRange(item *isegosdk.ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetsResponseConditionHoursRange) []map[string]interface{} {
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

func flattenNetworkAccessPolicySetGetNetworkAccessPolicySetsItemsConditionHoursRangeException(item *isegosdk.ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetsResponseConditionHoursRangeException) []map[string]interface{} {
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

func flattenNetworkAccessPolicySetGetNetworkAccessPolicySetsItemsLink(item *isegosdk.ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetsResponseLink) []map[string]interface{} {
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

func flattenNetworkAccessPolicySetGetNetworkAccessPolicySetByIDItem(item *isegosdk.ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetByIDResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["condition"] = flattenNetworkAccessPolicySetGetNetworkAccessPolicySetByIDItemCondition(item.Condition)
	respItem["default"] = item.Default
	respItem["description"] = item.Description
	respItem["hit_counts"] = item.HitCounts
	respItem["id"] = item.ID
	respItem["is_proxy"] = item.IsProxy
	respItem["link"] = flattenNetworkAccessPolicySetGetNetworkAccessPolicySetByIDItemLink(item.Link)
	respItem["name"] = item.Name
	respItem["rank"] = item.Rank
	respItem["service_name"] = item.ServiceName
	respItem["state"] = item.State
	return []map[string]interface{}{
		respItem,
	}
}

func flattenNetworkAccessPolicySetGetNetworkAccessPolicySetByIDItemCondition(item *isegosdk.ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetByIDResponseCondition) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["condition_type"] = item.ConditionType
	respItem["is_negate"] = item.IsNegate
	respItem["link"] = flattenNetworkAccessPolicySetGetNetworkAccessPolicySetByIDItemConditionLink(item.Link)
	respItem["description"] = item.Description
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["attribute_name"] = item.AttributeName
	respItem["attribute_id"] = item.AttributeID
	respItem["attribute_value"] = item.AttributeValue
	respItem["dictionary_name"] = item.DictionaryName
	respItem["dictionary_value"] = item.DictionaryValue
	respItem["operator"] = item.Operator
	respItem["children"] = flattenNetworkAccessPolicySetGetNetworkAccessPolicySetByIDItemConditionChildren(item.Children)
	respItem["dates_range"] = flattenNetworkAccessPolicySetGetNetworkAccessPolicySetByIDItemConditionDatesRange(item.DatesRange)
	respItem["dates_range_exception"] = flattenNetworkAccessPolicySetGetNetworkAccessPolicySetByIDItemConditionDatesRangeException(item.DatesRangeException)
	respItem["hours_range"] = flattenNetworkAccessPolicySetGetNetworkAccessPolicySetByIDItemConditionHoursRange(item.HoursRange)
	respItem["hours_range_exception"] = flattenNetworkAccessPolicySetGetNetworkAccessPolicySetByIDItemConditionHoursRangeException(item.HoursRangeException)
	respItem["week_days"] = item.WeekDays
	respItem["week_days_exception"] = item.WeekDaysException

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNetworkAccessPolicySetGetNetworkAccessPolicySetByIDItemConditionLink(item *isegosdk.ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetByIDResponseConditionLink) []map[string]interface{} {
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

func flattenNetworkAccessPolicySetGetNetworkAccessPolicySetByIDItemConditionChildren(items *[]isegosdk.ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetByIDResponseConditionChildren) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["condition_type"] = item.ConditionType
		respItem["is_negate"] = item.IsNegate
		respItem["link"] = flattenNetworkAccessPolicySetGetNetworkAccessPolicySetByIDItemConditionChildrenLink(item.Link)
	}
	return respItems

}

func flattenNetworkAccessPolicySetGetNetworkAccessPolicySetByIDItemConditionChildrenLink(item *isegosdk.ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetByIDResponseConditionChildrenLink) []map[string]interface{} {
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

func flattenNetworkAccessPolicySetGetNetworkAccessPolicySetByIDItemConditionDatesRange(item *isegosdk.ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetByIDResponseConditionDatesRange) []map[string]interface{} {
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

func flattenNetworkAccessPolicySetGetNetworkAccessPolicySetByIDItemConditionDatesRangeException(item *isegosdk.ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetByIDResponseConditionDatesRangeException) []map[string]interface{} {
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

func flattenNetworkAccessPolicySetGetNetworkAccessPolicySetByIDItemConditionHoursRange(item *isegosdk.ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetByIDResponseConditionHoursRange) []map[string]interface{} {
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

func flattenNetworkAccessPolicySetGetNetworkAccessPolicySetByIDItemConditionHoursRangeException(item *isegosdk.ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetByIDResponseConditionHoursRangeException) []map[string]interface{} {
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

func flattenNetworkAccessPolicySetGetNetworkAccessPolicySetByIDItemLink(item *isegosdk.ResponseNetworkAccessPolicySetGetNetworkAccessPolicySetByIDResponseLink) []map[string]interface{} {
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
