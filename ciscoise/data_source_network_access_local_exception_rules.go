package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkAccessLocalExceptionRules() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Network Access - Authorization Exception Rules.

- Network Access Get local exception rules.

- Network Access Get local exception rule attributes.
`,

		ReadContext: dataSourceNetworkAccessLocalExceptionRulesRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. Rule id`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"policy_id": &schema.Schema{
				Description: `policyId path parameter. Policy id`,
				Type:        schema.TypeString,
				Optional:    true,
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
							Description: `The authorization profile/s`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"rule": &schema.Schema{
							Description: `Common attributes in rule authentication/authorization`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"condition": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"attribute_name": &schema.Schema{
													Description: `Dictionary attribute name`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"attribute_value": &schema.Schema{
													Description: `<ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"children": &schema.Schema{
													Description: `In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition`,
													Type:        schema.TypeList,
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"condition_type": &schema.Schema{
																Description: `<ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"is_negate": &schema.Schema{
																Description: `Indicates whereas this condition is in negate mode`,
																Type:        schema.TypeString,
																Computed:    true,
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
													Description: `<ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"dates_range": &schema.Schema{
													Description: `<p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>`,
													Type:        schema.TypeList,
													Computed:    true,
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
													Description: `<p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>`,
													Type:        schema.TypeList,
													Computed:    true,
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
													Description: `Condition description`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"dictionary_name": &schema.Schema{
													Description: `Dictionary name`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"dictionary_value": &schema.Schema{
													Description: `Dictionary value`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"hours_range": &schema.Schema{
													Description: `<p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>`,
													Type:        schema.TypeList,
													Computed:    true,
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
													Description: `<p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>`,
													Type:        schema.TypeList,
													Computed:    true,
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
													Description: `Indicates whereas this condition is in negate mode`,
													Type:        schema.TypeString,
													Computed:    true,
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
													Description: `Condition name`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"operator": &schema.Schema{
													Description: `Equality operator`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"week_days": &schema.Schema{
													Description: `<p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>`,
													Type:        schema.TypeList,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"week_days_exception": &schema.Schema{
													Description: `<p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>`,
													Type:        schema.TypeList,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
									"default": &schema.Schema{
										Description: `Indicates if this rule is the default one`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"hit_counts": &schema.Schema{
										Description: `The amount of times the rule was matched`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
									"id": &schema.Schema{
										Description: `The identifier of the rule`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"name": &schema.Schema{
										Description: `Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"rank": &schema.Schema{
										Description: `The rank(priority) in relation to other rules. Lower rank is higher priority.`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
									"state": &schema.Schema{
										Description: `The state that the rule is in. A disabled rule cannot be matched.`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
						"security_group": &schema.Schema{
							Description: `Security group used in authorization policies`,
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
							Description: `The authorization profile/s`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"rule": &schema.Schema{
							Description: `Common attributes in rule authentication/authorization`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"condition": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"attribute_name": &schema.Schema{
													Description: `Dictionary attribute name`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"attribute_value": &schema.Schema{
													Description: `<ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"children": &schema.Schema{
													Description: `In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition`,
													Type:        schema.TypeList,
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"condition_type": &schema.Schema{
																Description: `<ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"is_negate": &schema.Schema{
																Description: `Indicates whereas this condition is in negate mode`,
																Type:        schema.TypeString,
																Computed:    true,
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
													Description: `<ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"dates_range": &schema.Schema{
													Description: `<p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>`,
													Type:        schema.TypeList,
													Computed:    true,
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
													Description: `<p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>`,
													Type:        schema.TypeList,
													Computed:    true,
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
													Description: `Condition description`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"dictionary_name": &schema.Schema{
													Description: `Dictionary name`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"dictionary_value": &schema.Schema{
													Description: `Dictionary value`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"hours_range": &schema.Schema{
													Description: `<p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>`,
													Type:        schema.TypeList,
													Computed:    true,
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
													Description: `<p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>`,
													Type:        schema.TypeList,
													Computed:    true,
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
													Description: `Indicates whereas this condition is in negate mode`,
													Type:        schema.TypeString,
													Computed:    true,
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
													Description: `Condition name`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"operator": &schema.Schema{
													Description: `Equality operator`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"week_days": &schema.Schema{
													Description: `<p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>`,
													Type:        schema.TypeList,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"week_days_exception": &schema.Schema{
													Description: `<p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>`,
													Type:        schema.TypeList,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
									"default": &schema.Schema{
										Description: `Indicates if this rule is the default one`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"hit_counts": &schema.Schema{
										Description: `The amount of times the rule was matched`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
									"id": &schema.Schema{
										Description: `The identifier of the rule`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"name": &schema.Schema{
										Description: `Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"rank": &schema.Schema{
										Description: `The rank(priority) in relation to other rules. Lower rank is higher priority.`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
									"state": &schema.Schema{
										Description: `The state that the rule is in. A disabled rule cannot be matched.`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
						"security_group": &schema.Schema{
							Description: `Security group used in authorization policies`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceNetworkAccessLocalExceptionRulesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		log.Printf("[DEBUG] Selected method 1: GetNetworkAccessLocalExceptionRules")
		vvPolicyID := vPolicyID.(string)

		response1, restyResp1, err := client.NetworkAccessAuthorizationExceptionRules.GetNetworkAccessLocalExceptionRules(vvPolicyID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetNetworkAccessLocalExceptionRules", err,
				"Failure at GetNetworkAccessLocalExceptionRules, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRulesItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNetworkAccessLocalExceptionRules response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: GetNetworkAccessLocalExceptionRuleByID")
		vvPolicyID := vPolicyID.(string)
		vvID := vID.(string)

		response2, restyResp2, err := client.NetworkAccessAuthorizationExceptionRules.GetNetworkAccessLocalExceptionRuleByID(vvPolicyID, vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetNetworkAccessLocalExceptionRuleByID", err,
				"Failure at GetNetworkAccessLocalExceptionRuleByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRuleByIDItem(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNetworkAccessLocalExceptionRuleByID response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRulesItems(items *[]isegosdk.ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRulesResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["link"] = flattenNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRulesItemsLink(item.Link)
		respItem["profile"] = item.Profile
		respItem["rule"] = flattenNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRulesItemsRule(item.Rule)
		respItem["security_group"] = item.SecurityGroup
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRulesItemsLink(item *isegosdk.ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRulesResponseLink) []map[string]interface{} {
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

func flattenNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRulesItemsRule(item *isegosdk.ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRulesResponseRule) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["condition"] = flattenNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRulesItemsRuleCondition(item.Condition)
	respItem["default"] = boolPtrToString(item.Default)
	respItem["hit_counts"] = item.HitCounts
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["rank"] = item.Rank
	respItem["state"] = item.State

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRulesItemsRuleCondition(item *isegosdk.ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRulesResponseRuleCondition) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["condition_type"] = item.ConditionType
	respItem["is_negate"] = boolPtrToString(item.IsNegate)
	respItem["link"] = flattenNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRulesItemsRuleConditionLink(item.Link)
	respItem["description"] = item.Description
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["attribute_name"] = item.AttributeName
	respItem["attribute_value"] = item.AttributeValue
	respItem["dictionary_name"] = item.DictionaryName
	respItem["dictionary_value"] = item.DictionaryValue
	respItem["operator"] = item.Operator
	respItem["children"] = flattenNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRulesItemsRuleConditionChildren(item.Children)
	respItem["dates_range"] = flattenNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRulesItemsRuleConditionDatesRange(item.DatesRange)
	respItem["dates_range_exception"] = flattenNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRulesItemsRuleConditionDatesRangeException(item.DatesRangeException)
	respItem["hours_range"] = flattenNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRulesItemsRuleConditionHoursRange(item.HoursRange)
	respItem["hours_range_exception"] = flattenNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRulesItemsRuleConditionHoursRangeException(item.HoursRangeException)
	respItem["week_days"] = item.WeekDays
	respItem["week_days_exception"] = item.WeekDaysException

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRulesItemsRuleConditionLink(item *isegosdk.ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRulesResponseRuleConditionLink) []map[string]interface{} {
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

func flattenNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRulesItemsRuleConditionChildren(items *[]isegosdk.ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRulesResponseRuleConditionChildren) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["condition_type"] = item.ConditionType
		respItem["is_negate"] = boolPtrToString(item.IsNegate)
		respItem["link"] = flattenNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRulesItemsRuleConditionChildrenLink(item.Link)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRulesItemsRuleConditionChildrenLink(item *isegosdk.ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRulesResponseRuleConditionChildrenLink) []map[string]interface{} {
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

func flattenNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRulesItemsRuleConditionDatesRange(item *isegosdk.ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRulesResponseRuleConditionDatesRange) []map[string]interface{} {
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

func flattenNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRulesItemsRuleConditionDatesRangeException(item *isegosdk.ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRulesResponseRuleConditionDatesRangeException) []map[string]interface{} {
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

func flattenNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRulesItemsRuleConditionHoursRange(item *isegosdk.ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRulesResponseRuleConditionHoursRange) []map[string]interface{} {
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

func flattenNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRulesItemsRuleConditionHoursRangeException(item *isegosdk.ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRulesResponseRuleConditionHoursRangeException) []map[string]interface{} {
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

func flattenNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRuleByIDItem(item *isegosdk.ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRuleByIDResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["link"] = flattenNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRuleByIDItemLink(item.Link)
	respItem["profile"] = item.Profile
	respItem["rule"] = flattenNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRuleByIDItemRule(item.Rule)
	respItem["security_group"] = item.SecurityGroup
	return []map[string]interface{}{
		respItem,
	}
}

func flattenNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRuleByIDItemLink(item *isegosdk.ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRuleByIDResponseLink) []map[string]interface{} {
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

func flattenNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRuleByIDItemRule(item *isegosdk.ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRuleByIDResponseRule) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["condition"] = flattenNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRuleByIDItemRuleCondition(item.Condition)
	respItem["default"] = boolPtrToString(item.Default)
	respItem["hit_counts"] = item.HitCounts
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["rank"] = item.Rank
	respItem["state"] = item.State

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRuleByIDItemRuleCondition(item *isegosdk.ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRuleByIDResponseRuleCondition) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["condition_type"] = item.ConditionType
	respItem["is_negate"] = boolPtrToString(item.IsNegate)
	respItem["link"] = flattenNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRuleByIDItemRuleConditionLink(item.Link)
	respItem["description"] = item.Description
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["attribute_name"] = item.AttributeName
	respItem["attribute_value"] = item.AttributeValue
	respItem["dictionary_name"] = item.DictionaryName
	respItem["dictionary_value"] = item.DictionaryValue
	respItem["operator"] = item.Operator
	respItem["children"] = flattenNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRuleByIDItemRuleConditionChildren(item.Children)
	respItem["dates_range"] = flattenNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRuleByIDItemRuleConditionDatesRange(item.DatesRange)
	respItem["dates_range_exception"] = flattenNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRuleByIDItemRuleConditionDatesRangeException(item.DatesRangeException)
	respItem["hours_range"] = flattenNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRuleByIDItemRuleConditionHoursRange(item.HoursRange)
	respItem["hours_range_exception"] = flattenNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRuleByIDItemRuleConditionHoursRangeException(item.HoursRangeException)
	respItem["week_days"] = item.WeekDays
	respItem["week_days_exception"] = item.WeekDaysException

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRuleByIDItemRuleConditionLink(item *isegosdk.ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRuleByIDResponseRuleConditionLink) []map[string]interface{} {
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

func flattenNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRuleByIDItemRuleConditionChildren(items *[]isegosdk.ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRuleByIDResponseRuleConditionChildren) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["condition_type"] = item.ConditionType
		respItem["is_negate"] = boolPtrToString(item.IsNegate)
		respItem["link"] = flattenNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRuleByIDItemRuleConditionChildrenLink(item.Link)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRuleByIDItemRuleConditionChildrenLink(item *isegosdk.ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRuleByIDResponseRuleConditionChildrenLink) []map[string]interface{} {
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

func flattenNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRuleByIDItemRuleConditionDatesRange(item *isegosdk.ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRuleByIDResponseRuleConditionDatesRange) []map[string]interface{} {
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

func flattenNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRuleByIDItemRuleConditionDatesRangeException(item *isegosdk.ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRuleByIDResponseRuleConditionDatesRangeException) []map[string]interface{} {
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

func flattenNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRuleByIDItemRuleConditionHoursRange(item *isegosdk.ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRuleByIDResponseRuleConditionHoursRange) []map[string]interface{} {
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

func flattenNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRuleByIDItemRuleConditionHoursRangeException(item *isegosdk.ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRuleByIDResponseRuleConditionHoursRangeException) []map[string]interface{} {
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
