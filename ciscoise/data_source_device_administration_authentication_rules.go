package ciscoise

import (
	"context"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDeviceAdministrationAuthenticationRules() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDeviceAdministrationAuthenticationRulesRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"policy_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"identity_source_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"identity_source_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"if_auth_fail": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"if_process_fail": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"if_user_not_found": &schema.Schema{
							Type:     schema.TypeString,
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
						"rule": &schema.Schema{
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
									"hit_counts": &schema.Schema{
										Type:     schema.TypeInt,
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
									"rank": &schema.Schema{
										Type:     schema.TypeInt,
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
				},
			},
			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"identity_source_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"identity_source_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"if_auth_fail": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"if_process_fail": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"if_user_not_found": &schema.Schema{
							Type:     schema.TypeString,
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
						"rule": &schema.Schema{
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
									"hit_counts": &schema.Schema{
										Type:     schema.TypeInt,
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
									"rank": &schema.Schema{
										Type:     schema.TypeInt,
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
				},
			},
		},
	}
}

func dataSourceDeviceAdministrationAuthenticationRulesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vPolicyID, okPolicyID := d.GetOk("policy_id")
	vID, okID := d.GetOk("id")

	method1 := []bool{okPolicyID}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okPolicyID, okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetDeviceAdminAuthenticationRules")
		vvPolicyID := vPolicyID.(string)

		response1, _, err := client.DeviceAdministrationAuthenticationRules.GetDeviceAdminAuthenticationRules(vvPolicyID)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDeviceAdminAuthenticationRules", err,
				"Failure at GetDeviceAdminAuthenticationRules, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		vItems1 := flattenDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRulesItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceAdminAuthenticationRules response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: GetDeviceAdminAuthenticationRuleByID")
		vvPolicyID := vPolicyID.(string)
		vvID := vID.(string)

		response2, _, err := client.DeviceAdministrationAuthenticationRules.GetDeviceAdminAuthenticationRuleByID(vvPolicyID, vvID)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDeviceAdminAuthenticationRuleByID", err,
				"Failure at GetDeviceAdminAuthenticationRuleByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

		vItem2 := flattenDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRuleByIDItem(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceAdminAuthenticationRuleByID response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRulesItems(items *[]isegosdk.ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRulesResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["identity_source_id"] = item.IDentitySourceID
		respItem["identity_source_name"] = item.IDentitySourceName
		respItem["if_auth_fail"] = item.IfAuthFail
		respItem["if_process_fail"] = item.IfProcessFail
		respItem["if_user_not_found"] = item.IfUserNotFound
		respItem["link"] = flattenDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRulesItemsLink(item.Link)
		respItem["rule"] = flattenDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRulesItemsRule(item.Rule)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRulesItemsLink(item *isegosdk.ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRulesResponseLink) []map[string]interface{} {
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

func flattenDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRulesItemsRule(item *isegosdk.ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRulesResponseRule) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["condition"] = flattenDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRulesItemsRuleCondition(item.Condition)
	respItem["default"] = item.Default
	respItem["hit_counts"] = item.HitCounts
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["rank"] = item.Rank
	respItem["state"] = item.State

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRulesItemsRuleCondition(item *isegosdk.ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRulesResponseRuleCondition) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["condition_type"] = item.ConditionType
	respItem["is_negate"] = item.IsNegate
	respItem["link"] = flattenDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRulesItemsRuleConditionLink(item.Link)
	respItem["description"] = item.Description
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["attribute_name"] = item.AttributeName
	respItem["attribute_id"] = item.AttributeID
	respItem["attribute_value"] = item.AttributeValue
	respItem["dictionary_name"] = item.DictionaryName
	respItem["dictionary_value"] = item.DictionaryValue
	respItem["operator"] = item.Operator
	respItem["children"] = flattenDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRulesItemsRuleConditionChildren(item.Children)
	respItem["dates_range"] = flattenDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRulesItemsRuleConditionDatesRange(item.DatesRange)
	respItem["dates_range_exception"] = flattenDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRulesItemsRuleConditionDatesRangeException(item.DatesRangeException)
	respItem["hours_range"] = flattenDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRulesItemsRuleConditionHoursRange(item.HoursRange)
	respItem["hours_range_exception"] = flattenDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRulesItemsRuleConditionHoursRangeException(item.HoursRangeException)
	respItem["week_days"] = item.WeekDays
	respItem["week_days_exception"] = item.WeekDaysException

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRulesItemsRuleConditionLink(item *isegosdk.ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRulesResponseRuleConditionLink) []map[string]interface{} {
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

func flattenDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRulesItemsRuleConditionChildren(items *[]isegosdk.ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRulesResponseRuleConditionChildren) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["condition_type"] = item.ConditionType
		respItem["is_negate"] = item.IsNegate
		respItem["link"] = flattenDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRulesItemsRuleConditionChildrenLink(item.Link)
	}
	return respItems

}

func flattenDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRulesItemsRuleConditionChildrenLink(item *isegosdk.ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRulesResponseRuleConditionChildrenLink) []map[string]interface{} {
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

func flattenDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRulesItemsRuleConditionDatesRange(item *isegosdk.ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRulesResponseRuleConditionDatesRange) []map[string]interface{} {
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

func flattenDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRulesItemsRuleConditionDatesRangeException(item *isegosdk.ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRulesResponseRuleConditionDatesRangeException) []map[string]interface{} {
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

func flattenDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRulesItemsRuleConditionHoursRange(item *isegosdk.ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRulesResponseRuleConditionHoursRange) []map[string]interface{} {
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

func flattenDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRulesItemsRuleConditionHoursRangeException(item *isegosdk.ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRulesResponseRuleConditionHoursRangeException) []map[string]interface{} {
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

func flattenDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRuleByIDItem(item *isegosdk.ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRuleByIDResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["identity_source_id"] = item.IDentitySourceID
	respItem["identity_source_name"] = item.IDentitySourceName
	respItem["if_auth_fail"] = item.IfAuthFail
	respItem["if_process_fail"] = item.IfProcessFail
	respItem["if_user_not_found"] = item.IfUserNotFound
	respItem["link"] = flattenDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRuleByIDItemLink(item.Link)
	respItem["rule"] = flattenDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRuleByIDItemRule(item.Rule)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRuleByIDItemLink(item *isegosdk.ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRuleByIDResponseLink) []map[string]interface{} {
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

func flattenDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRuleByIDItemRule(item *isegosdk.ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRuleByIDResponseRule) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["condition"] = flattenDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRuleByIDItemRuleCondition(item.Condition)
	respItem["default"] = item.Default
	respItem["hit_counts"] = item.HitCounts
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["rank"] = item.Rank
	respItem["state"] = item.State

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRuleByIDItemRuleCondition(item *isegosdk.ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRuleByIDResponseRuleCondition) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["condition_type"] = item.ConditionType
	respItem["is_negate"] = item.IsNegate
	respItem["link"] = flattenDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRuleByIDItemRuleConditionLink(item.Link)
	respItem["description"] = item.Description
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["attribute_name"] = item.AttributeName
	respItem["attribute_id"] = item.AttributeID
	respItem["attribute_value"] = item.AttributeValue
	respItem["dictionary_name"] = item.DictionaryName
	respItem["dictionary_value"] = item.DictionaryValue
	respItem["operator"] = item.Operator
	respItem["children"] = flattenDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRuleByIDItemRuleConditionChildren(item.Children)
	respItem["dates_range"] = flattenDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRuleByIDItemRuleConditionDatesRange(item.DatesRange)
	respItem["dates_range_exception"] = flattenDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRuleByIDItemRuleConditionDatesRangeException(item.DatesRangeException)
	respItem["hours_range"] = flattenDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRuleByIDItemRuleConditionHoursRange(item.HoursRange)
	respItem["hours_range_exception"] = flattenDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRuleByIDItemRuleConditionHoursRangeException(item.HoursRangeException)
	respItem["week_days"] = item.WeekDays
	respItem["week_days_exception"] = item.WeekDaysException

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRuleByIDItemRuleConditionLink(item *isegosdk.ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRuleByIDResponseRuleConditionLink) []map[string]interface{} {
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

func flattenDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRuleByIDItemRuleConditionChildren(items *[]isegosdk.ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRuleByIDResponseRuleConditionChildren) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["condition_type"] = item.ConditionType
		respItem["is_negate"] = item.IsNegate
		respItem["link"] = flattenDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRuleByIDItemRuleConditionChildrenLink(item.Link)
	}
	return respItems

}

func flattenDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRuleByIDItemRuleConditionChildrenLink(item *isegosdk.ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRuleByIDResponseRuleConditionChildrenLink) []map[string]interface{} {
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

func flattenDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRuleByIDItemRuleConditionDatesRange(item *isegosdk.ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRuleByIDResponseRuleConditionDatesRange) []map[string]interface{} {
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

func flattenDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRuleByIDItemRuleConditionDatesRangeException(item *isegosdk.ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRuleByIDResponseRuleConditionDatesRangeException) []map[string]interface{} {
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

func flattenDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRuleByIDItemRuleConditionHoursRange(item *isegosdk.ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRuleByIDResponseRuleConditionHoursRange) []map[string]interface{} {
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

func flattenDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRuleByIDItemRuleConditionHoursRangeException(item *isegosdk.ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRuleByIDResponseRuleConditionHoursRangeException) []map[string]interface{} {
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
