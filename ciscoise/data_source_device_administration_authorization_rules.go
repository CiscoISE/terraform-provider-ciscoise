package ciscoise

import (
	"context"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDeviceAdministrationAuthorizationRules() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Device Administration - Authorization Rules.

- Device Admin Get authorization rules.
- Device Admin Get authorization rule attributes.`,

		ReadContext: dataSourceDeviceAdministrationAuthorizationRulesRead,
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

						"commands": &schema.Schema{
							Description: `Command sets enforce the specified list of commands that can be executed by a device administrator`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
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
						"profile": &schema.Schema{
							Description: `Device admin profiles control the initial login session of the device administrator`,
							Type:        schema.TypeString,
							Computed:    true,
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

												"attribute_id": &schema.Schema{
													Description: `Dictionary attribute id (Optional), used for additional verification`,
													Type:        schema.TypeString,
													Computed:    true,
												},
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
																Type:        schema.TypeBool,
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
													Description: `<p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>`,
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
													Description: `<p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>`,
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
													Description: `<p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>`,
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
													Description: `<p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>`,
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
													Type:        schema.TypeBool,
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
										Type:        schema.TypeBool,
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
					},
				},
			},
			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"commands": &schema.Schema{
							Description: `Command sets enforce the specified list of commands that can be executed by a device administrator`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
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
						"profile": &schema.Schema{
							Description: `Device admin profiles control the initial login session of the device administrator`,
							Type:        schema.TypeString,
							Computed:    true,
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

												"attribute_id": &schema.Schema{
													Description: `Dictionary attribute id (Optional), used for additional verification`,
													Type:        schema.TypeString,
													Computed:    true,
												},
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
																Type:        schema.TypeBool,
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
													Description: `<p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>`,
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
													Description: `<p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>`,
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
													Description: `<p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>`,
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
													Description: `<p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>`,
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
													Type:        schema.TypeBool,
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
										Type:        schema.TypeBool,
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
					},
				},
			},
		},
	}
}

func dataSourceDeviceAdministrationAuthorizationRulesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		log.Printf("[DEBUG] Selected method 1: GetDeviceAdminAuthorizationRules")
		vvPolicyID := vPolicyID.(string)

		response1, _, err := client.DeviceAdministrationAuthorizationRules.GetDeviceAdminAuthorizationRules(vvPolicyID)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDeviceAdminAuthorizationRules", err,
				"Failure at GetDeviceAdminAuthorizationRules, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		vItems1 := flattenDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRulesItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceAdminAuthorizationRules response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: GetDeviceAdminAuthorizationRuleByID")
		vvPolicyID := vPolicyID.(string)
		vvID := vID.(string)

		response2, _, err := client.DeviceAdministrationAuthorizationRules.GetDeviceAdminAuthorizationRuleByID(vvPolicyID, vvID)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDeviceAdminAuthorizationRuleByID", err,
				"Failure at GetDeviceAdminAuthorizationRuleByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

		vItem2 := flattenDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRuleByIDItem(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceAdminAuthorizationRuleByID response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRulesItems(items *[]isegosdk.ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRulesResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["commands"] = item.Commands
		respItem["link"] = flattenDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRulesItemsLink(item.Link)
		respItem["profile"] = item.Profile
		respItem["rule"] = flattenDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRulesItemsRule(item.Rule)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRulesItemsLink(item *isegosdk.ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRulesResponseLink) []map[string]interface{} {
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

func flattenDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRulesItemsRule(item *isegosdk.ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRulesResponseRule) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["condition"] = flattenDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRulesItemsRuleCondition(item.Condition)
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

func flattenDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRulesItemsRuleCondition(item *isegosdk.ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRulesResponseRuleCondition) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["condition_type"] = item.ConditionType
	respItem["is_negate"] = item.IsNegate
	respItem["link"] = flattenDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRulesItemsRuleConditionLink(item.Link)
	respItem["description"] = item.Description
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["attribute_name"] = item.AttributeName
	respItem["attribute_id"] = item.AttributeID
	respItem["attribute_value"] = item.AttributeValue
	respItem["dictionary_name"] = item.DictionaryName
	respItem["dictionary_value"] = item.DictionaryValue
	respItem["operator"] = item.Operator
	respItem["children"] = flattenDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRulesItemsRuleConditionChildren(item.Children)
	respItem["dates_range"] = flattenDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRulesItemsRuleConditionDatesRange(item.DatesRange)
	respItem["dates_range_exception"] = flattenDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRulesItemsRuleConditionDatesRangeException(item.DatesRangeException)
	respItem["hours_range"] = flattenDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRulesItemsRuleConditionHoursRange(item.HoursRange)
	respItem["hours_range_exception"] = flattenDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRulesItemsRuleConditionHoursRangeException(item.HoursRangeException)
	respItem["week_days"] = item.WeekDays
	respItem["week_days_exception"] = item.WeekDaysException

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRulesItemsRuleConditionLink(item *isegosdk.ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRulesResponseRuleConditionLink) []map[string]interface{} {
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

func flattenDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRulesItemsRuleConditionChildren(items *[]isegosdk.ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRulesResponseRuleConditionChildren) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["condition_type"] = item.ConditionType
		respItem["is_negate"] = item.IsNegate
		respItem["link"] = flattenDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRulesItemsRuleConditionChildrenLink(item.Link)
	}
	return respItems

}

func flattenDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRulesItemsRuleConditionChildrenLink(item *isegosdk.ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRulesResponseRuleConditionChildrenLink) []map[string]interface{} {
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

func flattenDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRulesItemsRuleConditionDatesRange(item *isegosdk.ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRulesResponseRuleConditionDatesRange) []map[string]interface{} {
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

func flattenDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRulesItemsRuleConditionDatesRangeException(item *isegosdk.ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRulesResponseRuleConditionDatesRangeException) []map[string]interface{} {
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

func flattenDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRulesItemsRuleConditionHoursRange(item *isegosdk.ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRulesResponseRuleConditionHoursRange) []map[string]interface{} {
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

func flattenDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRulesItemsRuleConditionHoursRangeException(item *isegosdk.ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRulesResponseRuleConditionHoursRangeException) []map[string]interface{} {
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

func flattenDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRuleByIDItem(item *isegosdk.ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRuleByIDResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["commands"] = item.Commands
	respItem["link"] = flattenDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRuleByIDItemLink(item.Link)
	respItem["profile"] = item.Profile
	respItem["rule"] = flattenDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRuleByIDItemRule(item.Rule)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRuleByIDItemLink(item *isegosdk.ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRuleByIDResponseLink) []map[string]interface{} {
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

func flattenDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRuleByIDItemRule(item *isegosdk.ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRuleByIDResponseRule) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["condition"] = flattenDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRuleByIDItemRuleCondition(item.Condition)
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

func flattenDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRuleByIDItemRuleCondition(item *isegosdk.ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRuleByIDResponseRuleCondition) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["condition_type"] = item.ConditionType
	respItem["is_negate"] = item.IsNegate
	respItem["link"] = flattenDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRuleByIDItemRuleConditionLink(item.Link)
	respItem["description"] = item.Description
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["attribute_name"] = item.AttributeName
	respItem["attribute_id"] = item.AttributeID
	respItem["attribute_value"] = item.AttributeValue
	respItem["dictionary_name"] = item.DictionaryName
	respItem["dictionary_value"] = item.DictionaryValue
	respItem["operator"] = item.Operator
	respItem["children"] = flattenDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRuleByIDItemRuleConditionChildren(item.Children)
	respItem["dates_range"] = flattenDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRuleByIDItemRuleConditionDatesRange(item.DatesRange)
	respItem["dates_range_exception"] = flattenDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRuleByIDItemRuleConditionDatesRangeException(item.DatesRangeException)
	respItem["hours_range"] = flattenDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRuleByIDItemRuleConditionHoursRange(item.HoursRange)
	respItem["hours_range_exception"] = flattenDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRuleByIDItemRuleConditionHoursRangeException(item.HoursRangeException)
	respItem["week_days"] = item.WeekDays
	respItem["week_days_exception"] = item.WeekDaysException

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRuleByIDItemRuleConditionLink(item *isegosdk.ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRuleByIDResponseRuleConditionLink) []map[string]interface{} {
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

func flattenDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRuleByIDItemRuleConditionChildren(items *[]isegosdk.ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRuleByIDResponseRuleConditionChildren) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["condition_type"] = item.ConditionType
		respItem["is_negate"] = item.IsNegate
		respItem["link"] = flattenDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRuleByIDItemRuleConditionChildrenLink(item.Link)
	}
	return respItems

}

func flattenDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRuleByIDItemRuleConditionChildrenLink(item *isegosdk.ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRuleByIDResponseRuleConditionChildrenLink) []map[string]interface{} {
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

func flattenDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRuleByIDItemRuleConditionDatesRange(item *isegosdk.ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRuleByIDResponseRuleConditionDatesRange) []map[string]interface{} {
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

func flattenDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRuleByIDItemRuleConditionDatesRangeException(item *isegosdk.ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRuleByIDResponseRuleConditionDatesRangeException) []map[string]interface{} {
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

func flattenDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRuleByIDItemRuleConditionHoursRange(item *isegosdk.ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRuleByIDResponseRuleConditionHoursRange) []map[string]interface{} {
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

func flattenDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRuleByIDItemRuleConditionHoursRangeException(item *isegosdk.ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRuleByIDResponseRuleConditionHoursRangeException) []map[string]interface{} {
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
