package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDeviceAdministrationGlobalExceptionRules() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Device Administration - Authorization Global Exception Rules.

- Device Admin Get global execption rules.

- Device Admin Get global exception rule attribute
`,

		ReadContext: dataSourceDeviceAdministrationGlobalExceptionRulesRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. Rule id`,
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
					},
				},
			},
		},
	}
}

func dataSourceDeviceAdministrationGlobalExceptionRulesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vID, okID := d.GetOk("id")

	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetDeviceAdminPolicySetGlobalExceptionRules")

		response1, _, err := client.DeviceAdministrationAuthorizationGlobalExceptionRules.GetDeviceAdminPolicySetGlobalExceptionRules()

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDeviceAdminPolicySetGlobalExceptionRules", err,
				"Failure at GetDeviceAdminPolicySetGlobalExceptionRules, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionRulesItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceAdminPolicySetGlobalExceptionRules response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: GetDeviceAdminPolicySetGlobalExceptionByRuleID")
		vvID := vID.(string)

		response2, _, err := client.DeviceAdministrationAuthorizationGlobalExceptionRules.GetDeviceAdminPolicySetGlobalExceptionByRuleID(vvID)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDeviceAdminPolicySetGlobalExceptionByRuleID", err,
				"Failure at GetDeviceAdminPolicySetGlobalExceptionByRuleID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionByRuleIDItem(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceAdminPolicySetGlobalExceptionByRuleID response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionRulesItems(items *[]isegosdk.ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionRulesResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["commands"] = item.Commands
		respItem["link"] = flattenDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionRulesItemsLink(item.Link)
		respItem["profile"] = item.Profile
		respItem["rule"] = flattenDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionRulesItemsRule(item.Rule)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionRulesItemsLink(item *isegosdk.ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionRulesResponseLink) []map[string]interface{} {
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

func flattenDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionRulesItemsRule(item *isegosdk.ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionRulesResponseRule) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["condition"] = flattenDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionRulesItemsRuleCondition(item.Condition)
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

func flattenDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionRulesItemsRuleCondition(item *isegosdk.ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionRulesResponseRuleCondition) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["condition_type"] = item.ConditionType
	respItem["is_negate"] = boolPtrToString(item.IsNegate)
	respItem["link"] = flattenDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionRulesItemsRuleConditionLink(item.Link)
	respItem["description"] = item.Description
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["attribute_name"] = item.AttributeName
	respItem["attribute_id"] = item.AttributeID
	respItem["attribute_value"] = item.AttributeValue
	respItem["dictionary_name"] = item.DictionaryName
	respItem["dictionary_value"] = item.DictionaryValue
	respItem["operator"] = item.Operator
	respItem["children"] = flattenDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionRulesItemsRuleConditionChildren(item.Children)
	respItem["dates_range"] = flattenDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionRulesItemsRuleConditionDatesRange(item.DatesRange)
	respItem["dates_range_exception"] = flattenDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionRulesItemsRuleConditionDatesRangeException(item.DatesRangeException)
	respItem["hours_range"] = flattenDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionRulesItemsRuleConditionHoursRange(item.HoursRange)
	respItem["hours_range_exception"] = flattenDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionRulesItemsRuleConditionHoursRangeException(item.HoursRangeException)
	respItem["week_days"] = item.WeekDays
	respItem["week_days_exception"] = item.WeekDaysException

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionRulesItemsRuleConditionLink(item *isegosdk.ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionRulesResponseRuleConditionLink) []map[string]interface{} {
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

func flattenDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionRulesItemsRuleConditionChildren(items *[]isegosdk.ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionRulesResponseRuleConditionChildren) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["condition_type"] = item.ConditionType
		respItem["is_negate"] = boolPtrToString(item.IsNegate)
		respItem["link"] = flattenDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionRulesItemsRuleConditionChildrenLink(item.Link)
		respItems = append(respItems, respItem)
	}
	return respItems

}

func flattenDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionRulesItemsRuleConditionChildrenLink(item *isegosdk.ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionRulesResponseRuleConditionChildrenLink) []map[string]interface{} {
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

func flattenDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionRulesItemsRuleConditionDatesRange(item *isegosdk.ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionRulesResponseRuleConditionDatesRange) []map[string]interface{} {
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

func flattenDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionRulesItemsRuleConditionDatesRangeException(item *isegosdk.ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionRulesResponseRuleConditionDatesRangeException) []map[string]interface{} {
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

func flattenDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionRulesItemsRuleConditionHoursRange(item *isegosdk.ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionRulesResponseRuleConditionHoursRange) []map[string]interface{} {
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

func flattenDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionRulesItemsRuleConditionHoursRangeException(item *isegosdk.ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionRulesResponseRuleConditionHoursRangeException) []map[string]interface{} {
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

func flattenDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionByRuleIDItem(item *isegosdk.ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionByRuleIDResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["commands"] = item.Commands
	respItem["link"] = flattenDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionByRuleIDItemLink(item.Link)
	respItem["profile"] = item.Profile
	respItem["rule"] = flattenDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionByRuleIDItemRule(item.Rule)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionByRuleIDItemLink(item *isegosdk.ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionByRuleIDResponseLink) []map[string]interface{} {
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

func flattenDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionByRuleIDItemRule(item *isegosdk.ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionByRuleIDResponseRule) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["condition"] = flattenDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionByRuleIDItemRuleCondition(item.Condition)
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

func flattenDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionByRuleIDItemRuleCondition(item *isegosdk.ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionByRuleIDResponseRuleCondition) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["condition_type"] = item.ConditionType
	respItem["is_negate"] = boolPtrToString(item.IsNegate)
	respItem["link"] = flattenDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionByRuleIDItemRuleConditionLink(item.Link)
	respItem["description"] = item.Description
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["attribute_name"] = item.AttributeName
	respItem["attribute_id"] = item.AttributeID
	respItem["attribute_value"] = item.AttributeValue
	respItem["dictionary_name"] = item.DictionaryName
	respItem["dictionary_value"] = item.DictionaryValue
	respItem["operator"] = item.Operator
	respItem["children"] = flattenDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionByRuleIDItemRuleConditionChildren(item.Children)
	respItem["dates_range"] = flattenDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionByRuleIDItemRuleConditionDatesRange(item.DatesRange)
	respItem["dates_range_exception"] = flattenDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionByRuleIDItemRuleConditionDatesRangeException(item.DatesRangeException)
	respItem["hours_range"] = flattenDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionByRuleIDItemRuleConditionHoursRange(item.HoursRange)
	respItem["hours_range_exception"] = flattenDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionByRuleIDItemRuleConditionHoursRangeException(item.HoursRangeException)
	respItem["week_days"] = item.WeekDays
	respItem["week_days_exception"] = item.WeekDaysException

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionByRuleIDItemRuleConditionLink(item *isegosdk.ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionByRuleIDResponseRuleConditionLink) []map[string]interface{} {
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

func flattenDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionByRuleIDItemRuleConditionChildren(items *[]isegosdk.ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionByRuleIDResponseRuleConditionChildren) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["condition_type"] = item.ConditionType
		respItem["is_negate"] = boolPtrToString(item.IsNegate)
		respItem["link"] = flattenDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionByRuleIDItemRuleConditionChildrenLink(item.Link)
		respItems = append(respItems, respItem)
	}
	return respItems

}

func flattenDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionByRuleIDItemRuleConditionChildrenLink(item *isegosdk.ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionByRuleIDResponseRuleConditionChildrenLink) []map[string]interface{} {
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

func flattenDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionByRuleIDItemRuleConditionDatesRange(item *isegosdk.ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionByRuleIDResponseRuleConditionDatesRange) []map[string]interface{} {
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

func flattenDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionByRuleIDItemRuleConditionDatesRangeException(item *isegosdk.ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionByRuleIDResponseRuleConditionDatesRangeException) []map[string]interface{} {
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

func flattenDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionByRuleIDItemRuleConditionHoursRange(item *isegosdk.ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionByRuleIDResponseRuleConditionHoursRange) []map[string]interface{} {
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

func flattenDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionByRuleIDItemRuleConditionHoursRangeException(item *isegosdk.ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionByRuleIDResponseRuleConditionHoursRangeException) []map[string]interface{} {
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
