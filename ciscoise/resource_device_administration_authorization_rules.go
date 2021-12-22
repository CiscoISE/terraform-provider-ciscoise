package ciscoise

import (
	"context"
	"fmt"
	"reflect"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDeviceAdministrationAuthorizationRules() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on Device Administration - Authorization Rules.

- Device Admin Create authorization rule:

 Rule must include name and condition.

 Condition has hierarchical structure which define a set of conditions for which authoriztion policy rule could be
match.

 Condition can be either reference to a stored Library condition, using model
ConditionReference

 or dynamically built conditions which are not stored in the conditions Library, using models
ConditionAttributes, ConditionAndBlock, ConditionOrBlock
.


- Device Admin Update authorization rule.

- Device Admin Delete authorization rule.
`,

		CreateContext: resourceDeviceAdministrationAuthorizationRulesCreate,
		ReadContext:   resourceDeviceAdministrationAuthorizationRulesRead,
		UpdateContext: resourceDeviceAdministrationAuthorizationRulesUpdate,
		DeleteContext: resourceDeviceAdministrationAuthorizationRulesDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
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
						"policy_id": &schema.Schema{
							Description: `policyId path parameter. Policy id`,
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
					},
				},
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"commands": &schema.Schema{
							Description: `Command sets enforce the specified list of commands that can be executed by a device administrator`,
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"id": &schema.Schema{
							Description: `id path parameter. Rule id`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"profile": &schema.Schema{
							Description: `Device admin profiles control the initial login session of the device administrator`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"policy_id": &schema.Schema{
							Description: `policyId path parameter. Policy id`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"rule": &schema.Schema{
							Description: `Common attributes in rule authentication/authorization`,
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"condition": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"attribute_name": &schema.Schema{
													Description: `Dictionary attribute name`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"attribute_value": &schema.Schema{
													Description: `<ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"children": &schema.Schema{
													Description: `In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition`,
													Type:        schema.TypeList,
													Optional:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"condition_type": &schema.Schema{
																Description: `<ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>`,
																Type:        schema.TypeString,
																Optional:    true,
															},
															"is_negate": &schema.Schema{
																Description:  `Indicates whereas this condition is in negate mode`,
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
															},
														},
													},
												},
												"condition_type": &schema.Schema{
													Description: `<ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"dates_range": &schema.Schema{
													Description: `<p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>`,
													Type:        schema.TypeList,
													Optional:    true,
													MaxItems:    1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"end_date": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
															"start_date": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},
												"dates_range_exception": &schema.Schema{
													Description: `<p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>`,
													Type:        schema.TypeList,
													Optional:    true,
													MaxItems:    1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"end_date": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
															"start_date": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},
												"description": &schema.Schema{
													Description: `Condition description`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"dictionary_name": &schema.Schema{
													Description: `Dictionary name`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"dictionary_value": &schema.Schema{
													Description: `Dictionary value`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"hours_range": &schema.Schema{
													Description: `<p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>`,
													Type:        schema.TypeList,
													Optional:    true,
													MaxItems:    1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"end_time": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
															"start_time": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},
												"hours_range_exception": &schema.Schema{
													Description: `<p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>`,
													Type:        schema.TypeList,
													Optional:    true,
													MaxItems:    1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"end_time": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
															"start_time": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},
												"id": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"is_negate": &schema.Schema{
													Description:  `Indicates whereas this condition is in negate mode`,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},

												"name": &schema.Schema{
													Description: `Condition name`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"operator": &schema.Schema{
													Description: `Equality operator`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"week_days": &schema.Schema{
													Description: `<p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>`,
													Type:        schema.TypeList,
													Optional:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"week_days_exception": &schema.Schema{
													Description: `<p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>`,
													Type:        schema.TypeList,
													Optional:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
									"default": &schema.Schema{
										Description:  `Indicates if this rule is the default one`,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
									},
									"hit_counts": &schema.Schema{
										Description: `The amount of times the rule was matched`,
										Type:        schema.TypeInt,
										Optional:    true,
									},
									"id": &schema.Schema{
										Description: `The identifier of the rule`,
										Type:        schema.TypeString,
										Optional:    true,
									},
									"name": &schema.Schema{
										Description: `Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]`,
										Type:        schema.TypeString,
										Optional:    true,
									},
									"rank": &schema.Schema{
										Description: `The rank(priority) in relation to other rules. Lower rank is higher priority.`,
										Type:        schema.TypeInt,
										Optional:    true,
									},
									"state": &schema.Schema{
										Description: `The state that the rule is in. A disabled rule cannot be matched.`,
										Type:        schema.TypeString,
										Optional:    true,
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

func resourceDeviceAdministrationAuthorizationRulesCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRule(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	vPolicyID, okPolicyID := resourceItem["policy_id"]
	vvPolicyID := interfaceToString(vPolicyID)
	vID, okID := resourceItem["id"]
	var vvName string
	if !okID || vID == "" {
		if _, ok := d.GetOk("parameters.0.rule"); ok {
			if v, ok2 := d.GetOkExists("parameters.0.rule.0.id"); ok2 {
				vID = interfaceToString(v)
				okID = ok2
			}
		}
	}
	vvID := interfaceToString(vID)
	if _, ok := d.GetOk("parameters.0.rule"); ok {
		if v, ok2 := d.GetOkExists("parameters.0.rule.0.name"); ok2 {
			vvName = interfaceToString(v)
		}
	}
	if okPolicyID && vvPolicyID != "" && okID && vvID != "" {
		getResponse2, _, err := client.DeviceAdministrationAuthorizationRules.GetDeviceAdminAuthorizationRuleByID(vvPolicyID, vvID)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["policy_id"] = vvPolicyID
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return resourceDeviceAdministrationAuthorizationRulesRead(ctx, d, m)
		}
	} else {
		response2, _, err := client.DeviceAdministrationAuthorizationRules.GetDeviceAdminAuthorizationRules(vvPolicyID)
		if response2 != nil && err == nil {
			items2 := getAllItemsDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRules(m, response2, vvPolicyID)
			item2, err := searchDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRules(m, items2, vvName, vvID, vvPolicyID)
			if err == nil && item2 != nil {
				resourceMap := make(map[string]string)
				resourceMap["policy_id"] = vvPolicyID
				resourceMap["id"] = vvID
				resourceMap["name"] = vvName
				d.SetId(joinResourceID(resourceMap))
				return resourceDeviceAdministrationAuthorizationRulesRead(ctx, d, m)
			}
		}
	}
	resp1, restyResp1, err := client.DeviceAdministrationAuthorizationRules.CreateDeviceAdminAuthorizationRule(vvPolicyID, request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateDeviceAdminAuthorizationRule", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateDeviceAdminAuthorizationRule", err))
		return diags
	}
	if vvID != resp1.Response.Rule.ID {
		vvID = resp1.Response.Rule.ID
	}
	if vvName != resp1.Response.Rule.Name {
		vvName = resp1.Response.Rule.Name
	}
	resourceMap := make(map[string]string)
	resourceMap["policy_id"] = vvPolicyID
	resourceMap["id"] = vvID
	d.SetId(joinResourceID(resourceMap))
	return resourceDeviceAdministrationAuthorizationRulesRead(ctx, d, m)
}

func resourceDeviceAdministrationAuthorizationRulesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vPolicyID, okPolicyID := resourceMap["policy_id"]
	vID, okID := resourceMap["id"]
	vName, okName := resourceMap["name"]

	if !okID || vID == "" {
		if _, ok := d.GetOk("parameters.0.rule"); ok {
			if v, ok2 := d.GetOk("parameters.0.rule.0.id"); ok2 {
				vID = interfaceToString(v)
				okID = ok2
			}
		}
	}
	if _, ok := d.GetOk("parameters.0.rule"); ok {
		if v, ok2 := d.GetOk("parameters.0.rule.0.name"); ok2 {
			vName = interfaceToString(v)
			okName = ok2
		}
	}
	vvID := vID
	vvPolicyID := vPolicyID
	vvName := vName
	method1 := []bool{okPolicyID, okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okPolicyID, okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetDeviceAdminAuthorizationRules")

		response1, restyResp1, err := client.DeviceAdministrationAuthorizationRules.GetDeviceAdminAuthorizationRules(vvPolicyID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDeviceAdminAuthorizationRules", err,
				"Failure at GetDeviceAdminAuthorizationRules, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		items1 := getAllItemsDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRules(m, response1, vvPolicyID)
		item1, err := searchDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRules(m, items1, vvName, vvID, vvPolicyID)
		if err != nil || item1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when searching item from GetDeviceAdminAuthorizationRules response", err,
				"Failure when searching item from GetDeviceAdminAuthorizationRules, unexpected response", ""))
			return diags
		}
		vItem1 := flattenDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRuleByIDItem(item1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceAdminAuthorizationRules search response",
				err))
			return diags
		}

	}
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDeviceAdminAuthorizationRuleByID")
		vvPolicyID := vPolicyID
		vvID := vID

		response2, restyResp2, err := client.DeviceAdministrationAuthorizationRules.GetDeviceAdminAuthorizationRuleByID(vvPolicyID, vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDeviceAdminAuthorizationRuleByID", err,
				"Failure at GetDeviceAdminAuthorizationRuleByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRuleByIDItem(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceAdminAuthorizationRuleByID response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceDeviceAdministrationAuthorizationRulesUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vPolicyID, okPolicyID := resourceMap["policy_id"]
	vID, okID := resourceMap["id"]
	vName, okName := resourceMap["name"]

	if !okID || vID == "" {
		if _, ok := d.GetOk("parameters.0.rule"); ok {
			if v, ok2 := d.GetOk("parameters.0.rule.0.id"); ok2 {
				vID = interfaceToString(v)
				okID = ok2
			}
		}
	}
	if _, ok := d.GetOk("parameters.0.rule"); ok {
		if v, ok2 := d.GetOk("parameters.0.rule.0.name"); ok2 {
			vName = interfaceToString(v)
			okName = ok2
		}
	}
	vvID := vID
	vvPolicyID := vPolicyID
	vvName := vName
	method1 := []bool{okPolicyID, okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okPolicyID, okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	// NOTE: Consider adding getAllItems and search function to get missing params
	if selectedMethod == 2 {
		getResp1, _, err := client.DeviceAdministrationAuthorizationRules.GetDeviceAdminAuthorizationRules(vvPolicyID)
		if err == nil && getResp1 != nil {
			items1 := getAllItemsDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRules(m, getResp1, vvPolicyID)
			item1, err := searchDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRules(m, items1, vvName, vvID, vvPolicyID)
			if err == nil && item1 != nil {
				if item1.Rule != nil && vID != item1.Rule.ID {
					vvID = item1.Rule.ID
				} else {
					vvID = vID
				}
			}
		}
	}
	if selectedMethod == 1 {
		vvID = vID
	}
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] ID used for update operation %s", vvID)
		request1 := expandRequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByID(ctx, "parameters.0", d)
		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
		response1, restyResp1, err := client.DeviceAdministrationAuthorizationRules.UpdateDeviceAdminAuthorizationRuleByID(vvPolicyID, vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateDeviceAdminAuthorizationRuleByID", err, restyResp1.String(),
					"Failure at UpdateDeviceAdminAuthorizationRuleByID, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateDeviceAdminAuthorizationRuleByID", err,
				"Failure at UpdateDeviceAdminAuthorizationRuleByID, unexpected response", ""))
			return diags
		}
	}

	return resourceDeviceAdministrationAuthorizationRulesRead(ctx, d, m)
}

func resourceDeviceAdministrationAuthorizationRulesDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vPolicyID, okPolicyID := resourceMap["policy_id"]
	vID, okID := resourceMap["id"]
	vName, okName := resourceMap["name"]

	if !okID || vID == "" {
		if _, ok := d.GetOk("parameters.0.rule"); ok {
			if v, ok2 := d.GetOk("parameters.0.rule.0.id"); ok2 {
				vID = interfaceToString(v)
				okID = ok2
			}
		}
	}
	if _, ok := d.GetOk("parameters.0.rule"); ok {
		if v, ok2 := d.GetOk("parameters.0.rule.0.name"); ok2 {
			vName = interfaceToString(v)
			okName = ok2
		}
	}
	vvID := vID
	vvPolicyID := vPolicyID
	vvName := vName
	method1 := []bool{okPolicyID, okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okPolicyID, okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	// REVIEW: Add getAllItems and search function to get missing params
	if selectedMethod == 2 {

		getResp1, _, err := client.DeviceAdministrationAuthorizationRules.GetDeviceAdminAuthorizationRules(vvPolicyID)
		if err != nil || getResp1 == nil {
			// Assume that element it is already gone
			return diags
		}
		items1 := getAllItemsDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRules(m, getResp1, vvPolicyID)
		item1, err := searchDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRules(m, items1, vvName, vvID, vvPolicyID)
		if err != nil || item1 == nil {
			// Assume that element it is already gone
			return diags
		}
		if item1.Rule != nil && vID != item1.Rule.ID {
			vvID = item1.Rule.ID
		} else {
			vvID = vID
		}
	}
	if selectedMethod == 1 {
		vvID = vID
		getResp, _, err := client.DeviceAdministrationAuthorizationRules.GetDeviceAdminAuthorizationRuleByID(vvPolicyID, vvID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	response1, restyResp1, err := client.DeviceAdministrationAuthorizationRules.DeleteDeviceAdminAuthorizationRuleByID(vvPolicyID, vvID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteDeviceAdminAuthorizationRuleByID", err, restyResp1.String(),
				"Failure at DeleteDeviceAdminAuthorizationRuleByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteDeviceAdminAuthorizationRuleByID", err,
			"Failure at DeleteDeviceAdminAuthorizationRuleByID, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRule(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRule {
	request := isegosdk.RequestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRule{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".commands")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".commands")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".commands")))) {
		request.Commands = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".profile")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".profile")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".profile")))) {
		request.Profile = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rule")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rule")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rule")))) {
		request.Rule = expandRequestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleRule(ctx, key+".rule.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleLink(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleLink {
	request := isegosdk.RequestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleLink{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".href")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".href")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".href")))) {
		request.Href = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rel")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rel")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rel")))) {
		request.Rel = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleRule(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleRule {
	request := isegosdk.RequestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleRule{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".condition")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".condition")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".condition")))) {
		request.Condition = expandRequestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleRuleCondition(ctx, key+".condition.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".default")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".default")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".default")))) {
		request.Default = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".hit_counts")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".hit_counts")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".hit_counts")))) {
		request.HitCounts = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rank")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rank")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rank")))) {
		request.Rank = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".state")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".state")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".state")))) {
		request.State = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleRuleCondition(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleRuleCondition {
	request := isegosdk.RequestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleRuleCondition{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".condition_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".condition_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".condition_type")))) {
		request.ConditionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_negate")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_negate")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_negate")))) {
		request.IsNegate = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".attribute_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".attribute_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".attribute_name")))) {
		request.AttributeName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".attribute_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".attribute_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".attribute_value")))) {
		request.AttributeValue = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dictionary_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dictionary_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dictionary_name")))) {
		request.DictionaryName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dictionary_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dictionary_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dictionary_value")))) {
		request.DictionaryValue = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".operator")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".operator")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".operator")))) {
		request.Operator = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".children")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".children")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".children")))) {
		request.Children = expandRequestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleRuleConditionChildrenArray(ctx, key+".children", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dates_range")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dates_range")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dates_range")))) {
		request.DatesRange = expandRequestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleRuleConditionDatesRange(ctx, key+".dates_range.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dates_range_exception")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dates_range_exception")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dates_range_exception")))) {
		request.DatesRangeException = expandRequestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleRuleConditionDatesRangeException(ctx, key+".dates_range_exception.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".hours_range")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".hours_range")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".hours_range")))) {
		request.HoursRange = expandRequestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleRuleConditionHoursRange(ctx, key+".hours_range.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".hours_range_exception")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".hours_range_exception")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".hours_range_exception")))) {
		request.HoursRangeException = expandRequestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleRuleConditionHoursRangeException(ctx, key+".hours_range_exception.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".week_days")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".week_days")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".week_days")))) {
		request.WeekDays = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".week_days_exception")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".week_days_exception")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".week_days_exception")))) {
		request.WeekDaysException = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleRuleConditionLink(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleRuleConditionLink {
	request := isegosdk.RequestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleRuleConditionLink{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".href")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".href")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".href")))) {
		request.Href = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rel")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rel")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rel")))) {
		request.Rel = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleRuleConditionChildrenArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleRuleConditionChildren {
	request := []isegosdk.RequestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleRuleConditionChildren{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleRuleConditionChildren(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleRuleConditionChildren(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleRuleConditionChildren {
	request := isegosdk.RequestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleRuleConditionChildren{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".condition_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".condition_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".condition_type")))) {
		request.ConditionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_negate")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_negate")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_negate")))) {
		request.IsNegate = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleRuleConditionChildrenLink(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleRuleConditionChildrenLink {
	request := isegosdk.RequestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleRuleConditionChildrenLink{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".href")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".href")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".href")))) {
		request.Href = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rel")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rel")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rel")))) {
		request.Rel = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleRuleConditionDatesRange(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleRuleConditionDatesRange {
	request := isegosdk.RequestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleRuleConditionDatesRange{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_date")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_date")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_date")))) {
		request.EndDate = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_date")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_date")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_date")))) {
		request.StartDate = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleRuleConditionDatesRangeException(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleRuleConditionDatesRangeException {
	request := isegosdk.RequestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleRuleConditionDatesRangeException{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_date")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_date")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_date")))) {
		request.EndDate = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_date")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_date")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_date")))) {
		request.StartDate = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleRuleConditionHoursRange(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleRuleConditionHoursRange {
	request := isegosdk.RequestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleRuleConditionHoursRange{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_time")))) {
		request.EndTime = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_time")))) {
		request.StartTime = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleRuleConditionHoursRangeException(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleRuleConditionHoursRangeException {
	request := isegosdk.RequestDeviceAdministrationAuthorizationRulesCreateDeviceAdminAuthorizationRuleRuleConditionHoursRangeException{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_time")))) {
		request.EndTime = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_time")))) {
		request.StartTime = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByID(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByID {
	request := isegosdk.RequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByID{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".commands")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".commands")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".commands")))) {
		request.Commands = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".profile")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".profile")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".profile")))) {
		request.Profile = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rule")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rule")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rule")))) {
		request.Rule = expandRequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRule(ctx, key+".rule.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDLink(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDLink {
	request := isegosdk.RequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDLink{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".href")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".href")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".href")))) {
		request.Href = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rel")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rel")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rel")))) {
		request.Rel = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRule(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRule {
	request := isegosdk.RequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRule{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".condition")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".condition")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".condition")))) {
		request.Condition = expandRequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRuleCondition(ctx, key+".condition.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".default")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".default")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".default")))) {
		request.Default = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".hit_counts")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".hit_counts")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".hit_counts")))) {
		request.HitCounts = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rank")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rank")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rank")))) {
		request.Rank = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".state")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".state")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".state")))) {
		request.State = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRuleCondition(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRuleCondition {
	request := isegosdk.RequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRuleCondition{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".condition_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".condition_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".condition_type")))) {
		request.ConditionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_negate")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_negate")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_negate")))) {
		request.IsNegate = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".attribute_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".attribute_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".attribute_name")))) {
		request.AttributeName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".attribute_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".attribute_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".attribute_value")))) {
		request.AttributeValue = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dictionary_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dictionary_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dictionary_name")))) {
		request.DictionaryName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dictionary_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dictionary_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dictionary_value")))) {
		request.DictionaryValue = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".operator")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".operator")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".operator")))) {
		request.Operator = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".children")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".children")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".children")))) {
		request.Children = expandRequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRuleConditionChildrenArray(ctx, key+".children", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dates_range")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dates_range")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dates_range")))) {
		request.DatesRange = expandRequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRuleConditionDatesRange(ctx, key+".dates_range.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dates_range_exception")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dates_range_exception")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dates_range_exception")))) {
		request.DatesRangeException = expandRequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRuleConditionDatesRangeException(ctx, key+".dates_range_exception.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".hours_range")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".hours_range")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".hours_range")))) {
		request.HoursRange = expandRequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRuleConditionHoursRange(ctx, key+".hours_range.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".hours_range_exception")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".hours_range_exception")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".hours_range_exception")))) {
		request.HoursRangeException = expandRequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRuleConditionHoursRangeException(ctx, key+".hours_range_exception.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".week_days")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".week_days")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".week_days")))) {
		request.WeekDays = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".week_days_exception")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".week_days_exception")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".week_days_exception")))) {
		request.WeekDaysException = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRuleConditionLink(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRuleConditionLink {
	request := isegosdk.RequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRuleConditionLink{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".href")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".href")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".href")))) {
		request.Href = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rel")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rel")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rel")))) {
		request.Rel = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRuleConditionChildrenArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRuleConditionChildren {
	request := []isegosdk.RequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRuleConditionChildren{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRuleConditionChildren(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRuleConditionChildren(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRuleConditionChildren {
	request := isegosdk.RequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRuleConditionChildren{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".condition_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".condition_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".condition_type")))) {
		request.ConditionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_negate")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_negate")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_negate")))) {
		request.IsNegate = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRuleConditionChildrenLink(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRuleConditionChildrenLink {
	request := isegosdk.RequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRuleConditionChildrenLink{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".href")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".href")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".href")))) {
		request.Href = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rel")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rel")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rel")))) {
		request.Rel = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRuleConditionDatesRange(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRuleConditionDatesRange {
	request := isegosdk.RequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRuleConditionDatesRange{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_date")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_date")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_date")))) {
		request.EndDate = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_date")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_date")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_date")))) {
		request.StartDate = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRuleConditionDatesRangeException(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRuleConditionDatesRangeException {
	request := isegosdk.RequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRuleConditionDatesRangeException{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_date")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_date")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_date")))) {
		request.EndDate = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_date")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_date")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_date")))) {
		request.StartDate = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRuleConditionHoursRange(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRuleConditionHoursRange {
	request := isegosdk.RequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRuleConditionHoursRange{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_time")))) {
		request.EndTime = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_time")))) {
		request.StartTime = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRuleConditionHoursRangeException(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRuleConditionHoursRangeException {
	request := isegosdk.RequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRuleConditionHoursRangeException{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_time")))) {
		request.EndTime = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_time")))) {
		request.StartTime = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func getAllItemsDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRules(m interface{}, response *isegosdk.ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRules, policyTypeID string) []isegosdk.ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRulesResponse {
	var respItems []isegosdk.ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRulesResponse
	if response.Response != nil && len(*response.Response) > 0 {
		respItems = append(respItems, *response.Response...)
	}
	return respItems
}

func searchDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRules(m interface{}, items []isegosdk.ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRulesResponse, name string, id string, policyID string) (*isegosdk.ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRuleByIDResponse, error) {
	client := m.(*isegosdk.Client)
	var err error
	var foundItem *isegosdk.ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRuleByIDResponse
	for _, item := range items {
		if id != "" && item.Rule != nil && item.Rule.ID == id {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRuleByID
			getItem, _, err = client.DeviceAdministrationAuthorizationRules.GetDeviceAdminAuthorizationRuleByID(policyID, id)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetDeviceAdminAuthorizationRuleByID")
			}
			foundItem = getItem.Response
			return foundItem, err
		} else if name != "" && item.Rule != nil && item.Rule.Name == name {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseDeviceAdministrationAuthorizationRulesGetDeviceAdminAuthorizationRuleByID
			getItem, _, err = client.DeviceAdministrationAuthorizationRules.GetDeviceAdminAuthorizationRuleByID(policyID, item.Rule.ID)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetDeviceAdminAuthorizationRuleByID")
			}
			foundItem = getItem.Response
			return foundItem, err
		}
	}
	return foundItem, err
}
