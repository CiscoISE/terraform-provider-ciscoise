package ciscoise

import (
	"context"

	"log"

	isegosdk "ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDeviceAdministrationAuthenticationRules() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Device Administration - Authentication Rules.

- Device Admin Get authentication rules.

- Device Admin Get rule attributes.
`,

		ReadContext: dataSourceDeviceAdministrationAuthenticationRulesRead,
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

						"identity_source_id": &schema.Schema{
							Description: `Identity source id from the identity stores`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"identity_source_name": &schema.Schema{
							Description: `Identity source name from the identity stores`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"if_auth_fail": &schema.Schema{
							Description: `Action to perform when authentication fails such as Bad credentials, disabled user and so on`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"if_process_fail": &schema.Schema{
							Description: `Action to perform when ISE is uanble to access the identity database`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"if_user_not_found": &schema.Schema{
							Description: `Action to perform when user is not found in any of identity stores`,
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

						"identity_source_id": &schema.Schema{
							Description: `Identity source id from the identity stores`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"identity_source_name": &schema.Schema{
							Description: `Identity source name from the identity stores`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"if_auth_fail": &schema.Schema{
							Description: `Action to perform when authentication fails such as Bad credentials, disabled user and so on`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"if_process_fail": &schema.Schema{
							Description: `Action to perform when ISE is uanble to access the identity database`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"if_user_not_found": &schema.Schema{
							Description: `Action to perform when user is not found in any of identity stores`,
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

func dataSourceDeviceAdministrationAuthenticationRulesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vPolicyID, okPolicyID := d.GetOk("policy_id")
	vID, okID := d.GetOk("id")

	method1 := []bool{okPolicyID, okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okPolicyID}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: GetDeviceAdminAuthenticationRules")
		vvPolicyID := vPolicyID.(string)

		response1, restyResp1, err := client.DeviceAdministrationAuthenticationRules.GetDeviceAdminAuthenticationRules(vvPolicyID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDeviceAdminAuthenticationRules", err,
				"Failure at GetDeviceAdminAuthenticationRules, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

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
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetDeviceAdminAuthenticationRuleByID")
		vvPolicyID := vPolicyID.(string)
		vvID := vID.(string)

		response2, _, err := client.DeviceAdministrationAuthenticationRules.GetDeviceAdminAuthenticationRuleByID(vvPolicyID, vvID)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDeviceAdminAuthenticationRuleByID", err,
				"Failure at GetDeviceAdminAuthenticationRuleByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

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

func flattenDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRulesItemsRuleCondition(item *isegosdk.ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRulesResponseRuleCondition) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["condition_type"] = item.ConditionType
	respItem["is_negate"] = boolPtrToString(item.IsNegate)
	respItem["link"] = flattenDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRulesItemsRuleConditionLink(item.Link)
	respItem["description"] = item.Description
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["attribute_name"] = item.AttributeName
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
		respItem["is_negate"] = boolPtrToString(item.IsNegate)
		respItem["link"] = flattenDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRulesItemsRuleConditionChildrenLink(item.Link)
		respItems = append(respItems, respItem)
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

func flattenDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRuleByIDItemRuleCondition(item *isegosdk.ResponseDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRuleByIDResponseRuleCondition) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["condition_type"] = item.ConditionType
	respItem["is_negate"] = boolPtrToString(item.IsNegate)
	respItem["link"] = flattenDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRuleByIDItemRuleConditionLink(item.Link)
	respItem["description"] = item.Description
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["attribute_name"] = item.AttributeName
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
		respItem["is_negate"] = boolPtrToString(item.IsNegate)
		respItem["link"] = flattenDeviceAdministrationAuthenticationRulesGetDeviceAdminAuthenticationRuleByIDItemRuleConditionChildrenLink(item.Link)
		respItems = append(respItems, respItem)
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
