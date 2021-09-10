package ciscoise

import (
	"context"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkAccessAuthenticationRules() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceNetworkAccessAuthenticationRulesRead,
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

func dataSourceNetworkAccessAuthenticationRulesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vPolicyID, okPolicyID := d.GetOk("policy_id")
	vID, okID := d.GetOk("id")

	method1 := []bool{okPolicyID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okPolicyID, okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetNetworkAccessAuthenticationRules")
		vvPolicyID := vPolicyID.(string)

		response1, _, err := client.NetworkAccessAuthenticationRules.GetNetworkAccessAuthenticationRules(vvPolicyID)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetNetworkAccessAuthenticationRules", err,
				"Failure at GetNetworkAccessAuthenticationRules, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		vItems1 := flattenNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRulesItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNetworkAccessAuthenticationRules response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: GetNetworkAccessAuthenticationRuleByID")
		vvPolicyID := vPolicyID.(string)
		vvID := vID.(string)

		response2, _, err := client.NetworkAccessAuthenticationRules.GetNetworkAccessAuthenticationRuleByID(vvPolicyID, vvID)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetNetworkAccessAuthenticationRuleByID", err,
				"Failure at GetNetworkAccessAuthenticationRuleByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

		vItem2 := flattenNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRuleByIDItem(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNetworkAccessAuthenticationRuleByID response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRulesItems(items *[]isegosdk.ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRulesResponse) []map[string]interface{} {
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
		respItem["link"] = flattenNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRulesItemsLink(item.Link)
		respItem["rule"] = flattenNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRulesItemsRule(item.Rule)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRulesItemsLink(item *isegosdk.ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRulesResponseLink) []map[string]interface{} {
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

func flattenNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRulesItemsRule(item *isegosdk.ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRulesResponseRule) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["condition"] = flattenNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRulesItemsRuleCondition(item.Condition)
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

func flattenNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRulesItemsRuleCondition(item *isegosdk.ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRulesResponseRuleCondition) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["condition_type"] = item.ConditionType
	respItem["is_negate"] = item.IsNegate
	respItem["link"] = flattenNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRulesItemsRuleConditionLink(item.Link)
	respItem["description"] = item.Description
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["attribute_name"] = item.AttributeName
	respItem["attribute_id"] = item.AttributeID
	respItem["attribute_value"] = item.AttributeValue
	respItem["dictionary_name"] = item.DictionaryName
	respItem["dictionary_value"] = item.DictionaryValue
	respItem["operator"] = item.Operator
	respItem["children"] = flattenNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRulesItemsRuleConditionChildren(item.Children)
	respItem["dates_range"] = flattenNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRulesItemsRuleConditionDatesRange(item.DatesRange)
	respItem["dates_range_exception"] = flattenNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRulesItemsRuleConditionDatesRangeException(item.DatesRangeException)
	respItem["hours_range"] = flattenNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRulesItemsRuleConditionHoursRange(item.HoursRange)
	respItem["hours_range_exception"] = flattenNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRulesItemsRuleConditionHoursRangeException(item.HoursRangeException)
	respItem["week_days"] = item.WeekDays
	respItem["week_days_exception"] = item.WeekDaysException

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRulesItemsRuleConditionLink(item *isegosdk.ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRulesResponseRuleConditionLink) []map[string]interface{} {
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

func flattenNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRulesItemsRuleConditionChildren(items *[]isegosdk.ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRulesResponseRuleConditionChildren) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["condition_type"] = item.ConditionType
		respItem["is_negate"] = item.IsNegate
		respItem["link"] = flattenNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRulesItemsRuleConditionChildrenLink(item.Link)
	}
	return respItems

}

func flattenNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRulesItemsRuleConditionChildrenLink(item *isegosdk.ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRulesResponseRuleConditionChildrenLink) []map[string]interface{} {
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

func flattenNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRulesItemsRuleConditionDatesRange(item *isegosdk.ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRulesResponseRuleConditionDatesRange) []map[string]interface{} {
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

func flattenNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRulesItemsRuleConditionDatesRangeException(item *isegosdk.ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRulesResponseRuleConditionDatesRangeException) []map[string]interface{} {
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

func flattenNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRulesItemsRuleConditionHoursRange(item *isegosdk.ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRulesResponseRuleConditionHoursRange) []map[string]interface{} {
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

func flattenNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRulesItemsRuleConditionHoursRangeException(item *isegosdk.ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRulesResponseRuleConditionHoursRangeException) []map[string]interface{} {
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

func flattenNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRuleByIDItem(item *isegosdk.ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRuleByIDResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["identity_source_id"] = item.IDentitySourceID
	respItem["identity_source_name"] = item.IDentitySourceName
	respItem["if_auth_fail"] = item.IfAuthFail
	respItem["if_process_fail"] = item.IfProcessFail
	respItem["if_user_not_found"] = item.IfUserNotFound
	respItem["link"] = flattenNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRuleByIDItemLink(item.Link)
	respItem["rule"] = flattenNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRuleByIDItemRule(item.Rule)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRuleByIDItemLink(item *isegosdk.ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRuleByIDResponseLink) []map[string]interface{} {
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

func flattenNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRuleByIDItemRule(item *isegosdk.ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRuleByIDResponseRule) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["condition"] = flattenNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRuleByIDItemRuleCondition(item.Condition)
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

func flattenNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRuleByIDItemRuleCondition(item *isegosdk.ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRuleByIDResponseRuleCondition) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["condition_type"] = item.ConditionType
	respItem["is_negate"] = item.IsNegate
	respItem["link"] = flattenNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRuleByIDItemRuleConditionLink(item.Link)
	respItem["description"] = item.Description
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["attribute_name"] = item.AttributeName
	respItem["attribute_id"] = item.AttributeID
	respItem["attribute_value"] = item.AttributeValue
	respItem["dictionary_name"] = item.DictionaryName
	respItem["dictionary_value"] = item.DictionaryValue
	respItem["operator"] = item.Operator
	respItem["children"] = flattenNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRuleByIDItemRuleConditionChildren(item.Children)
	respItem["dates_range"] = flattenNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRuleByIDItemRuleConditionDatesRange(item.DatesRange)
	respItem["dates_range_exception"] = flattenNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRuleByIDItemRuleConditionDatesRangeException(item.DatesRangeException)
	respItem["hours_range"] = flattenNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRuleByIDItemRuleConditionHoursRange(item.HoursRange)
	respItem["hours_range_exception"] = flattenNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRuleByIDItemRuleConditionHoursRangeException(item.HoursRangeException)
	respItem["week_days"] = item.WeekDays
	respItem["week_days_exception"] = item.WeekDaysException

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRuleByIDItemRuleConditionLink(item *isegosdk.ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRuleByIDResponseRuleConditionLink) []map[string]interface{} {
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

func flattenNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRuleByIDItemRuleConditionChildren(items *[]isegosdk.ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRuleByIDResponseRuleConditionChildren) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["condition_type"] = item.ConditionType
		respItem["is_negate"] = item.IsNegate
		respItem["link"] = flattenNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRuleByIDItemRuleConditionChildrenLink(item.Link)
	}
	return respItems

}

func flattenNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRuleByIDItemRuleConditionChildrenLink(item *isegosdk.ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRuleByIDResponseRuleConditionChildrenLink) []map[string]interface{} {
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

func flattenNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRuleByIDItemRuleConditionDatesRange(item *isegosdk.ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRuleByIDResponseRuleConditionDatesRange) []map[string]interface{} {
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

func flattenNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRuleByIDItemRuleConditionDatesRangeException(item *isegosdk.ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRuleByIDResponseRuleConditionDatesRangeException) []map[string]interface{} {
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

func flattenNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRuleByIDItemRuleConditionHoursRange(item *isegosdk.ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRuleByIDResponseRuleConditionHoursRange) []map[string]interface{} {
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

func flattenNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRuleByIDItemRuleConditionHoursRangeException(item *isegosdk.ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRuleByIDResponseRuleConditionHoursRangeException) []map[string]interface{} {
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
