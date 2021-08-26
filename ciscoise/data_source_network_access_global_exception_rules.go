package ciscoise

import (
	"context"

	"ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkAccessGlobalExceptionRules() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceNetworkAccessGlobalExceptionRulesRead,
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
						"profile": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
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
						"security_group": &schema.Schema{
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
						"profile": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
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
						"security_group": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceNetworkAccessGlobalExceptionRulesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vID, okID := d.GetOk("id")

	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetNetworkAccessPolicySetGlobalExceptionRules")

		response1, _, err := client.NetworkAccessAuthorizationGlobalExceptionRules.GetNetworkAccessPolicySetGlobalExceptionRules()

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetNetworkAccessPolicySetGlobalExceptionRules", err,
				"Failure at GetNetworkAccessPolicySetGlobalExceptionRules, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		vItems1 := flattenNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRulesItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNetworkAccessPolicySetGlobalExceptionRules response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: GetNetworkAccessPolicySetGlobalExceptionRuleByID")
		vvID := vID.(string)

		response2, _, err := client.NetworkAccessAuthorizationGlobalExceptionRules.GetNetworkAccessPolicySetGlobalExceptionRuleByID(vvID)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetNetworkAccessPolicySetGlobalExceptionRuleByID", err,
				"Failure at GetNetworkAccessPolicySetGlobalExceptionRuleByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

		vItem2 := flattenNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRuleByIDItem(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNetworkAccessPolicySetGlobalExceptionRuleByID response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRulesItems(items *[]isegosdk.ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRulesResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["link"] = flattenNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRulesItemsLink(item.Link)
		respItem["profile"] = item.Profile
		respItem["rule"] = flattenNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRulesItemsRule(item.Rule)
		respItem["security_group"] = item.SecurityGroup
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRulesItemsLink(item *isegosdk.ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRulesResponseLink) []map[string]interface{} {
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

func flattenNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRulesItemsRule(item *isegosdk.ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRulesResponseRule) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["condition"] = flattenNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRulesItemsRuleCondition(item.Condition)
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

func flattenNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRulesItemsRuleCondition(item *isegosdk.ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRulesResponseRuleCondition) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["condition_type"] = item.ConditionType
	respItem["is_negate"] = item.IsNegate
	respItem["link"] = flattenNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRulesItemsRuleConditionLink(item.Link)
	respItem["description"] = item.Description
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["attribute_name"] = item.AttributeName
	respItem["attribute_id"] = item.AttributeID
	respItem["attribute_value"] = item.AttributeValue
	respItem["dictionary_name"] = item.DictionaryName
	respItem["dictionary_value"] = item.DictionaryValue
	respItem["operator"] = item.Operator
	respItem["children"] = flattenNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRulesItemsRuleConditionChildren(item.Children)
	respItem["dates_range"] = flattenNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRulesItemsRuleConditionDatesRange(item.DatesRange)
	respItem["dates_range_exception"] = flattenNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRulesItemsRuleConditionDatesRangeException(item.DatesRangeException)
	respItem["hours_range"] = flattenNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRulesItemsRuleConditionHoursRange(item.HoursRange)
	respItem["hours_range_exception"] = flattenNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRulesItemsRuleConditionHoursRangeException(item.HoursRangeException)
	respItem["week_days"] = item.WeekDays
	respItem["week_days_exception"] = item.WeekDaysException

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRulesItemsRuleConditionLink(item *isegosdk.ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRulesResponseRuleConditionLink) []map[string]interface{} {
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

func flattenNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRulesItemsRuleConditionChildren(items *[]isegosdk.ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRulesResponseRuleConditionChildren) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["condition_type"] = item.ConditionType
		respItem["is_negate"] = item.IsNegate
		respItem["link"] = flattenNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRulesItemsRuleConditionChildrenLink(item.Link)
	}
	return respItems

}

func flattenNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRulesItemsRuleConditionChildrenLink(item *isegosdk.ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRulesResponseRuleConditionChildrenLink) []map[string]interface{} {
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

func flattenNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRulesItemsRuleConditionDatesRange(item *isegosdk.ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRulesResponseRuleConditionDatesRange) []map[string]interface{} {
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

func flattenNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRulesItemsRuleConditionDatesRangeException(item *isegosdk.ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRulesResponseRuleConditionDatesRangeException) []map[string]interface{} {
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

func flattenNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRulesItemsRuleConditionHoursRange(item *isegosdk.ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRulesResponseRuleConditionHoursRange) []map[string]interface{} {
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

func flattenNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRulesItemsRuleConditionHoursRangeException(item *isegosdk.ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRulesResponseRuleConditionHoursRangeException) []map[string]interface{} {
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

func flattenNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRuleByIDItem(item *isegosdk.ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRuleByIDResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["link"] = flattenNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRuleByIDItemLink(item.Link)
	respItem["profile"] = item.Profile
	respItem["rule"] = flattenNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRuleByIDItemRule(item.Rule)
	respItem["security_group"] = item.SecurityGroup
	return []map[string]interface{}{
		respItem,
	}
}

func flattenNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRuleByIDItemLink(item *isegosdk.ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRuleByIDResponseLink) []map[string]interface{} {
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

func flattenNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRuleByIDItemRule(item *isegosdk.ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRuleByIDResponseRule) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["condition"] = flattenNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRuleByIDItemRuleCondition(item.Condition)
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

func flattenNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRuleByIDItemRuleCondition(item *isegosdk.ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRuleByIDResponseRuleCondition) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["condition_type"] = item.ConditionType
	respItem["is_negate"] = item.IsNegate
	respItem["link"] = flattenNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRuleByIDItemRuleConditionLink(item.Link)
	respItem["description"] = item.Description
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["attribute_name"] = item.AttributeName
	respItem["attribute_id"] = item.AttributeID
	respItem["attribute_value"] = item.AttributeValue
	respItem["dictionary_name"] = item.DictionaryName
	respItem["dictionary_value"] = item.DictionaryValue
	respItem["operator"] = item.Operator
	respItem["children"] = flattenNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRuleByIDItemRuleConditionChildren(item.Children)
	respItem["dates_range"] = flattenNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRuleByIDItemRuleConditionDatesRange(item.DatesRange)
	respItem["dates_range_exception"] = flattenNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRuleByIDItemRuleConditionDatesRangeException(item.DatesRangeException)
	respItem["hours_range"] = flattenNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRuleByIDItemRuleConditionHoursRange(item.HoursRange)
	respItem["hours_range_exception"] = flattenNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRuleByIDItemRuleConditionHoursRangeException(item.HoursRangeException)
	respItem["week_days"] = item.WeekDays
	respItem["week_days_exception"] = item.WeekDaysException

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRuleByIDItemRuleConditionLink(item *isegosdk.ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRuleByIDResponseRuleConditionLink) []map[string]interface{} {
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

func flattenNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRuleByIDItemRuleConditionChildren(items *[]isegosdk.ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRuleByIDResponseRuleConditionChildren) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["condition_type"] = item.ConditionType
		respItem["is_negate"] = item.IsNegate
		respItem["link"] = flattenNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRuleByIDItemRuleConditionChildrenLink(item.Link)
	}
	return respItems

}

func flattenNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRuleByIDItemRuleConditionChildrenLink(item *isegosdk.ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRuleByIDResponseRuleConditionChildrenLink) []map[string]interface{} {
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

func flattenNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRuleByIDItemRuleConditionDatesRange(item *isegosdk.ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRuleByIDResponseRuleConditionDatesRange) []map[string]interface{} {
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

func flattenNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRuleByIDItemRuleConditionDatesRangeException(item *isegosdk.ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRuleByIDResponseRuleConditionDatesRangeException) []map[string]interface{} {
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

func flattenNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRuleByIDItemRuleConditionHoursRange(item *isegosdk.ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRuleByIDResponseRuleConditionHoursRange) []map[string]interface{} {
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

func flattenNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRuleByIDItemRuleConditionHoursRangeException(item *isegosdk.ResponseNetworkAccessAuthorizationGlobalExceptionRulesGetNetworkAccessPolicySetGlobalExceptionRuleByIDResponseRuleConditionHoursRangeException) []map[string]interface{} {
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
