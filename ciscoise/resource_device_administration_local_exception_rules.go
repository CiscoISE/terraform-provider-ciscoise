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

func resourceDeviceAdministrationLocalExceptionRules() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on Device Administration - Authorization Exception
Rules.

- Device Admin Create local authorization exception rule.

- Device Admin Update local exception rule.

- Device Admin Delete local exception rule.
`,

		CreateContext: resourceDeviceAdministrationLocalExceptionRulesCreate,
		ReadContext:   resourceDeviceAdministrationLocalExceptionRulesRead,
		UpdateContext: resourceDeviceAdministrationLocalExceptionRulesUpdate,
		DeleteContext: resourceDeviceAdministrationLocalExceptionRulesDelete,
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
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"commands": &schema.Schema{
							Description: `Command sets enforce the specified list of commands that can be executed by a device administrator`,
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"id": &schema.Schema{
							Description: `id path parameter. Rule id`,
							Type:        schema.TypeString,
							Optional:    true,
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
						"policy_id": &schema.Schema{
							Description: `policyId path parameter. Policy id`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"profile": &schema.Schema{
							Description: `Device admin profiles control the initial login session of the device administrator`,
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"rule": &schema.Schema{
							Description: `Common attributes in rule authentication/authorization`,
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"condition": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"attribute_id": &schema.Schema{
													Description: `Dictionary attribute id (Optional), used for additional verification`,
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"attribute_name": &schema.Schema{
													Description: `Dictionary attribute name`,
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"attribute_value": &schema.Schema{
													Description: `<ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>`,
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"children": &schema.Schema{
													Description: `In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition`,
													Type:        schema.TypeList,
													Optional:    true,
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"condition_type": &schema.Schema{
																Description: `<ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>`,
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
															},
															"is_negate": &schema.Schema{
																Description: `Indicates whereas this condition is in negate mode`,
																// Type:        schema.TypeBool,
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
																Computed:     true,
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
													Optional:    true,
													Computed:    true,
												},
												"dates_range": &schema.Schema{
													Description: `<p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>`,
													Type:        schema.TypeList,
													Optional:    true,
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"end_date": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"start_date": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
														},
													},
												},
												"dates_range_exception": &schema.Schema{
													Description: `<p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>`,
													Type:        schema.TypeList,
													Optional:    true,
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"end_date": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"start_date": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
														},
													},
												},
												"description": &schema.Schema{
													Description: `Condition description`,
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"dictionary_name": &schema.Schema{
													Description: `Dictionary name`,
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"dictionary_value": &schema.Schema{
													Description: `Dictionary value`,
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"hours_range": &schema.Schema{
													Description: `<p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>`,
													Type:        schema.TypeList,
													Optional:    true,
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"end_time": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"start_time": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
														},
													},
												},
												"hours_range_exception": &schema.Schema{
													Description: `<p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>`,
													Type:        schema.TypeList,
													Optional:    true,
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"end_time": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"start_time": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
														},
													},
												},
												"id": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"is_negate": &schema.Schema{
													Description: `Indicates whereas this condition is in negate mode`,
													// Type:        schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
													Computed:     true,
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
													Optional:    true,
													Computed:    true,
												},
												"operator": &schema.Schema{
													Description: `Equality operator`,
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"week_days": &schema.Schema{
													Description: `<p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>`,
													Type:        schema.TypeList,
													Optional:    true,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"week_days_exception": &schema.Schema{
													Description: `<p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>`,
													Type:        schema.TypeList,
													Optional:    true,
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
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										Computed:     true,
									},
									"hit_counts": &schema.Schema{
										Description: `The amount of times the rule was matched`,
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
									},
									"id": &schema.Schema{
										Description: `The identifier of the rule`,
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"name": &schema.Schema{
										Description: `Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]`,
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"rank": &schema.Schema{
										Description: `The rank(priority) in relation to other rules. Lower rank is higher priority.`,
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
									},
									"state": &schema.Schema{
										Description: `The state that the rule is in. A disabled rule cannot be matched.`,
										Type:        schema.TypeString,
										Optional:    true,
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

func resourceDeviceAdministrationLocalExceptionRulesCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("item"))
	request1 := expandRequestDeviceAdministrationLocalExceptionRulesCreateDeviceAdminLocalExceptionRule(ctx, "item.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vPolicyID, okPolicyID := resourceItem["policy_id"]
	vvPolicyID := interfaceToString(vPolicyID)
	vID, okID := resourceItem["id"]
	var vvName string
	if !okID || vID == "" {
		if _, ok := d.GetOkExists("item.0.rule"); ok {
			if v, ok2 := d.GetOkExists("item.0.rule.0.id"); ok2 {
				vID = interfaceToString(v)
				okID = ok2
			}
		}
	}
	vvID := interfaceToString(vID)
	if _, ok := d.GetOkExists("item.0.rule"); ok {
		if v, ok2 := d.GetOkExists("item.0.rule.0.name"); ok2 {
			vvName = interfaceToString(v)
		}
	}

	if okPolicyID && vvPolicyID != "" && okID && vvID != "" {
		getResponse2, _, err := client.DeviceAdministrationAuthorizationExceptionRules.GetDeviceAdminLocalExceptionRuleByID(vvPolicyID, vvID)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["policy_id"] = vvPolicyID
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return diags
		}
	} else {
		response2, _, err := client.DeviceAdministrationAuthorizationExceptionRules.GetDeviceAdminLocalExceptionRules(vvPolicyID)
		if response2 != nil && err == nil {
			items2 := getAllItemsDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRules(m, response2, vvPolicyID)
			item2, err := searchDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRules(m, items2, vvName, vvID, vvPolicyID)
			if err == nil && item2 != nil {
				resourceMap := make(map[string]string)
				resourceMap["policy_id"] = vvPolicyID
				resourceMap["id"] = vvID
				resourceMap["name"] = vvName
				d.SetId(joinResourceID(resourceMap))
				return diags
			}
		}
	}
	resp1, restyResp1, err := client.DeviceAdministrationAuthorizationExceptionRules.CreateDeviceAdminLocalExceptionRule(vvPolicyID, request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateDeviceAdminLocalExceptionRule", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateDeviceAdminLocalExceptionRule", err))
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
	resourceMap["name"] = vvName
	d.SetId(joinResourceID(resourceMap))
	return diags
}

func resourceDeviceAdministrationLocalExceptionRulesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vPolicyID, okPolicyID := resourceMap["policy_id"]
	vID, okID := resourceMap["id"]
	vName, okName := resourceMap["name"]

	if !okID || vID == "" {
		if _, ok := d.GetOk("item.0.rule"); ok {
			if v, ok2 := d.GetOk("item.0.rule.0.id"); ok2 {
				vID = interfaceToString(v)
				okID = ok2
			}
		}
	}
	if _, ok := d.GetOk("item.0.rule"); ok {
		if v, ok2 := d.GetOk("item.0.rule.0.name"); ok2 {
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
		log.Printf("[DEBUG] Selected method: GetDeviceAdminLocalExceptionRules")

		response1, _, err := client.DeviceAdministrationAuthorizationExceptionRules.GetDeviceAdminLocalExceptionRules(vvPolicyID)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDeviceAdminLocalExceptionRules", err,
				"Failure at GetDeviceAdminLocalExceptionRules, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		items1 := getAllItemsDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRules(m, response1, vvPolicyID)
		item1, err := searchDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRules(m, items1, vvName, vvID, vvPolicyID)
		if err != nil || item1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when searching item from GetDeviceAdminLocalExceptionRules response", err,
				"Failure when searching item from GetDeviceAdminLocalExceptionRules, unexpected response", ""))
			return diags
		}
		vItem1 := flattenDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRuleByIDItem(item1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceAdminLocalExceptionRules search response",
				err))
			return diags
		}

	}
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDeviceAdminLocalExceptionRuleByID")

		response2, _, err := client.DeviceAdministrationAuthorizationExceptionRules.GetDeviceAdminLocalExceptionRuleByID(vvPolicyID, vvID)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDeviceAdminLocalExceptionRuleByID", err,
				"Failure at GetDeviceAdminLocalExceptionRuleByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRuleByIDItem(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceAdminLocalExceptionRuleByID response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceDeviceAdministrationLocalExceptionRulesUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vPolicyID, okPolicyID := resourceMap["policy_id"]
	vID, okID := resourceMap["id"]
	vName, okName := resourceMap["name"]

	if !okID || vID == "" {
		if _, ok := d.GetOk("item.0.rule"); ok {
			if v, ok2 := d.GetOk("item.0.rule.0.id"); ok2 {
				vID = interfaceToString(v)
				okID = ok2
			}
		}
	}
	if _, ok := d.GetOk("item.0.rule"); ok {
		if v, ok2 := d.GetOk("item.0.rule.0.name"); ok2 {
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
		getResp1, _, err := client.DeviceAdministrationAuthorizationExceptionRules.GetDeviceAdminLocalExceptionRules(vvPolicyID)
		if err == nil && getResp1 != nil {
			items1 := getAllItemsDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRules(m, getResp1, vvPolicyID)
			item1, err := searchDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRules(m, items1, vvName, vvID, vvPolicyID)
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
	if d.HasChange("item") {
		log.Printf("[DEBUG] ID used for update operation %s", vvID)
		request1 := expandRequestDeviceAdministrationLocalExceptionRulesUpdateDeviceAdminLocalExceptionRuleByID(ctx, "item.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.DeviceAdministrationAuthorizationExceptionRules.UpdateDeviceAdminLocalExceptionRuleByID(vvPolicyID, vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateDeviceAdminLocalExceptionRuleByID", err, restyResp1.String(),
					"Failure at UpdateDeviceAdminLocalExceptionRuleByID, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateDeviceAdminLocalExceptionRuleByID", err,
				"Failure at UpdateDeviceAdminLocalExceptionRuleByID, unexpected response", ""))
			return diags
		}
	}

	return resourceDeviceAdministrationLocalExceptionRulesRead(ctx, d, m)
}

func resourceDeviceAdministrationLocalExceptionRulesDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vPolicyID, okPolicyID := resourceMap["policy_id"]
	vID, okID := resourceMap["id"]
	vName, okName := resourceMap["name"]

	if !okID || vID == "" {
		if _, ok := d.GetOk("item.0.rule"); ok {
			if v, ok2 := d.GetOk("item.0.rule.0.id"); ok2 {
				vID = interfaceToString(v)
				okID = ok2
			}
		}
	}
	if _, ok := d.GetOk("item.0.rule"); ok {
		if v, ok2 := d.GetOk("item.0.rule.0.name"); ok2 {
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

		getResp1, _, err := client.DeviceAdministrationAuthorizationExceptionRules.GetDeviceAdminLocalExceptionRules(vvPolicyID)
		if err != nil || getResp1 == nil {
			// Assume that element it is already gone
			return diags
		}
		items1 := getAllItemsDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRules(m, getResp1, vvPolicyID)
		item1, err := searchDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRules(m, items1, vvName, vvID, vvPolicyID)
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
		getResp, _, err := client.DeviceAdministrationAuthorizationExceptionRules.GetDeviceAdminLocalExceptionRuleByID(vvPolicyID, vvID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	response1, restyResp1, err := client.DeviceAdministrationAuthorizationExceptionRules.DeleteDeviceAdminLocalExceptionRuleByID(vvPolicyID, vvID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteDeviceAdminLocalExceptionRuleByID", err, restyResp1.String(),
				"Failure at DeleteDeviceAdminLocalExceptionRuleByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteDeviceAdminLocalExceptionRuleByID", err,
			"Failure at DeleteDeviceAdminLocalExceptionRuleByID, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestDeviceAdministrationLocalExceptionRulesCreateDeviceAdminLocalExceptionRule(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRule {
	request := isegosdk.RequestDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRule{}
	if v, ok := d.GetOkExists(key + ".commands"); !isEmptyValue(reflect.ValueOf(d.Get(key+".commands"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".commands"))) {
		request.Commands = interfaceToSliceString(v)
	}

	if v, ok := d.GetOkExists(key + ".profile"); !isEmptyValue(reflect.ValueOf(d.Get(key+".profile"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".profile"))) {
		request.Profile = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".rule"); !isEmptyValue(reflect.ValueOf(d.Get(key+".rule"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".rule"))) {
		request.Rule = expandRequestDeviceAdministrationLocalExceptionRulesCreateDeviceAdminLocalExceptionRuleRule(ctx, key+".rule.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceAdministrationLocalExceptionRulesCreateDeviceAdminLocalExceptionRuleLink(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRuleLink {
	request := isegosdk.RequestDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRuleLink{}
	if v, ok := d.GetOkExists(key + ".href"); !isEmptyValue(reflect.ValueOf(d.Get(key+".href"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".href"))) {
		request.Href = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".rel"); !isEmptyValue(reflect.ValueOf(d.Get(key+".rel"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".rel"))) {
		request.Rel = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".type"))) {
		request.Type = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceAdministrationLocalExceptionRulesCreateDeviceAdminLocalExceptionRuleRule(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRuleRule {
	request := isegosdk.RequestDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRuleRule{}
	if v, ok := d.GetOkExists(key + ".condition"); !isEmptyValue(reflect.ValueOf(d.Get(key+".condition"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".condition"))) {
		request.Condition = expandRequestDeviceAdministrationLocalExceptionRulesCreateDeviceAdminLocalExceptionRuleRuleCondition(ctx, key+".condition.0", d)
	}
	if v, ok := d.GetOkExists(key + ".default"); !isEmptyValue(reflect.ValueOf(d.Get(key+".default"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".default"))) {
		request.Default = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".hit_counts"); !isEmptyValue(reflect.ValueOf(d.Get(key+".hit_counts"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".hit_counts"))) {
		request.HitCounts = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".id"))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".name"))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".rank"); !isEmptyValue(reflect.ValueOf(d.Get(key+".rank"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".rank"))) {
		request.Rank = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".state"); !isEmptyValue(reflect.ValueOf(d.Get(key+".state"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".state"))) {
		request.State = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceAdministrationLocalExceptionRulesCreateDeviceAdminLocalExceptionRuleRuleCondition(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRuleRuleCondition {
	request := isegosdk.RequestDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRuleRuleCondition{}
	if v, ok := d.GetOkExists(key + ".condition_type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".condition_type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".condition_type"))) {
		request.ConditionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".is_negate"); !isEmptyValue(reflect.ValueOf(d.Get(key+".is_negate"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".is_negate"))) {
		request.IsNegate = interfaceToBoolPtr(v)
	}

	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".id"))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".name"))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".attribute_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".attribute_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".attribute_name"))) {
		request.AttributeName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".attribute_id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".attribute_id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".attribute_id"))) {
		request.AttributeID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".attribute_value"); !isEmptyValue(reflect.ValueOf(d.Get(key+".attribute_value"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".attribute_value"))) {
		request.AttributeValue = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".dictionary_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".dictionary_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".dictionary_name"))) {
		request.DictionaryName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".dictionary_value"); !isEmptyValue(reflect.ValueOf(d.Get(key+".dictionary_value"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".dictionary_value"))) {
		request.DictionaryValue = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".operator"); !isEmptyValue(reflect.ValueOf(d.Get(key+".operator"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".operator"))) {
		request.Operator = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".children"); !isEmptyValue(reflect.ValueOf(d.Get(key+".children"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".children"))) {
		request.Children = expandRequestDeviceAdministrationLocalExceptionRulesCreateDeviceAdminLocalExceptionRuleRuleConditionChildrenArray(ctx, key+".children", d)
	}
	if v, ok := d.GetOkExists(key + ".dates_range"); !isEmptyValue(reflect.ValueOf(d.Get(key+".dates_range"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".dates_range"))) {
		request.DatesRange = expandRequestDeviceAdministrationLocalExceptionRulesCreateDeviceAdminLocalExceptionRuleRuleConditionDatesRange(ctx, key+".dates_range.0", d)
	}
	if v, ok := d.GetOkExists(key + ".dates_range_exception"); !isEmptyValue(reflect.ValueOf(d.Get(key+".dates_range_exception"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".dates_range_exception"))) {
		request.DatesRangeException = expandRequestDeviceAdministrationLocalExceptionRulesCreateDeviceAdminLocalExceptionRuleRuleConditionDatesRangeException(ctx, key+".dates_range_exception.0", d)
	}
	if v, ok := d.GetOkExists(key + ".hours_range"); !isEmptyValue(reflect.ValueOf(d.Get(key+".hours_range"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".hours_range"))) {
		request.HoursRange = expandRequestDeviceAdministrationLocalExceptionRulesCreateDeviceAdminLocalExceptionRuleRuleConditionHoursRange(ctx, key+".hours_range.0", d)
	}
	if v, ok := d.GetOkExists(key + ".hours_range_exception"); !isEmptyValue(reflect.ValueOf(d.Get(key+".hours_range_exception"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".hours_range_exception"))) {
		request.HoursRangeException = expandRequestDeviceAdministrationLocalExceptionRulesCreateDeviceAdminLocalExceptionRuleRuleConditionHoursRangeException(ctx, key+".hours_range_exception.0", d)
	}
	if v, ok := d.GetOkExists(key + ".week_days"); !isEmptyValue(reflect.ValueOf(d.Get(key+".week_days"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".week_days"))) {
		request.WeekDays = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(key + ".week_days_exception"); !isEmptyValue(reflect.ValueOf(d.Get(key+".week_days_exception"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".week_days_exception"))) {
		request.WeekDaysException = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceAdministrationLocalExceptionRulesCreateDeviceAdminLocalExceptionRuleRuleConditionLink(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRuleRuleConditionLink {
	request := isegosdk.RequestDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRuleRuleConditionLink{}
	if v, ok := d.GetOkExists(key + ".href"); !isEmptyValue(reflect.ValueOf(d.Get(key+".href"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".href"))) {
		request.Href = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".rel"); !isEmptyValue(reflect.ValueOf(d.Get(key+".rel"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".rel"))) {
		request.Rel = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".type"))) {
		request.Type = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceAdministrationLocalExceptionRulesCreateDeviceAdminLocalExceptionRuleRuleConditionChildrenArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRuleRuleConditionChildren {
	request := []isegosdk.RequestDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRuleRuleConditionChildren{}
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestDeviceAdministrationLocalExceptionRulesCreateDeviceAdminLocalExceptionRuleRuleConditionChildren(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		request = append(request, *i)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceAdministrationLocalExceptionRulesCreateDeviceAdminLocalExceptionRuleRuleConditionChildren(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRuleRuleConditionChildren {
	request := isegosdk.RequestDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRuleRuleConditionChildren{}
	if v, ok := d.GetOkExists(key + ".condition_type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".condition_type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".condition_type"))) {
		request.ConditionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".is_negate"); !isEmptyValue(reflect.ValueOf(d.Get(key+".is_negate"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".is_negate"))) {
		request.IsNegate = interfaceToBoolPtr(v)
	}

	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceAdministrationLocalExceptionRulesCreateDeviceAdminLocalExceptionRuleRuleConditionChildrenLink(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRuleRuleConditionChildrenLink {
	request := isegosdk.RequestDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRuleRuleConditionChildrenLink{}
	if v, ok := d.GetOkExists(key + ".href"); !isEmptyValue(reflect.ValueOf(d.Get(key+".href"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".href"))) {
		request.Href = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".rel"); !isEmptyValue(reflect.ValueOf(d.Get(key+".rel"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".rel"))) {
		request.Rel = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".type"))) {
		request.Type = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceAdministrationLocalExceptionRulesCreateDeviceAdminLocalExceptionRuleRuleConditionDatesRange(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRuleRuleConditionDatesRange {
	request := isegosdk.RequestDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRuleRuleConditionDatesRange{}
	if v, ok := d.GetOkExists(key + ".end_date"); !isEmptyValue(reflect.ValueOf(d.Get(key+".end_date"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".end_date"))) {
		request.EndDate = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".start_date"); !isEmptyValue(reflect.ValueOf(d.Get(key+".start_date"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".start_date"))) {
		request.StartDate = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceAdministrationLocalExceptionRulesCreateDeviceAdminLocalExceptionRuleRuleConditionDatesRangeException(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRuleRuleConditionDatesRangeException {
	request := isegosdk.RequestDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRuleRuleConditionDatesRangeException{}
	if v, ok := d.GetOkExists(key + ".end_date"); !isEmptyValue(reflect.ValueOf(d.Get(key+".end_date"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".end_date"))) {
		request.EndDate = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".start_date"); !isEmptyValue(reflect.ValueOf(d.Get(key+".start_date"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".start_date"))) {
		request.StartDate = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceAdministrationLocalExceptionRulesCreateDeviceAdminLocalExceptionRuleRuleConditionHoursRange(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRuleRuleConditionHoursRange {
	request := isegosdk.RequestDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRuleRuleConditionHoursRange{}
	if v, ok := d.GetOkExists(key + ".end_time"); !isEmptyValue(reflect.ValueOf(d.Get(key+".end_time"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".end_time"))) {
		request.EndTime = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".start_time"); !isEmptyValue(reflect.ValueOf(d.Get(key+".start_time"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".start_time"))) {
		request.StartTime = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceAdministrationLocalExceptionRulesCreateDeviceAdminLocalExceptionRuleRuleConditionHoursRangeException(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRuleRuleConditionHoursRangeException {
	request := isegosdk.RequestDeviceAdministrationAuthorizationExceptionRulesCreateDeviceAdminLocalExceptionRuleRuleConditionHoursRangeException{}
	if v, ok := d.GetOkExists(key + ".end_time"); !isEmptyValue(reflect.ValueOf(d.Get(key+".end_time"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".end_time"))) {
		request.EndTime = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".start_time"); !isEmptyValue(reflect.ValueOf(d.Get(key+".start_time"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".start_time"))) {
		request.StartTime = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceAdministrationLocalExceptionRulesUpdateDeviceAdminLocalExceptionRuleByID(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByID {
	request := isegosdk.RequestDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByID{}
	if v, ok := d.GetOkExists(key + ".commands"); !isEmptyValue(reflect.ValueOf(d.Get(key+".commands"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".commands"))) {
		request.Commands = interfaceToSliceString(v)
	}

	if v, ok := d.GetOkExists(key + ".profile"); !isEmptyValue(reflect.ValueOf(d.Get(key+".profile"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".profile"))) {
		request.Profile = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".rule"); !isEmptyValue(reflect.ValueOf(d.Get(key+".rule"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".rule"))) {
		request.Rule = expandRequestDeviceAdministrationLocalExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDRule(ctx, key+".rule.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceAdministrationLocalExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDLink(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDLink {
	request := isegosdk.RequestDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDLink{}
	if v, ok := d.GetOkExists(key + ".href"); !isEmptyValue(reflect.ValueOf(d.Get(key+".href"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".href"))) {
		request.Href = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".rel"); !isEmptyValue(reflect.ValueOf(d.Get(key+".rel"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".rel"))) {
		request.Rel = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".type"))) {
		request.Type = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceAdministrationLocalExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDRule(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDRule {
	request := isegosdk.RequestDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDRule{}
	if v, ok := d.GetOkExists(key + ".condition"); !isEmptyValue(reflect.ValueOf(d.Get(key+".condition"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".condition"))) {
		request.Condition = expandRequestDeviceAdministrationLocalExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDRuleCondition(ctx, key+".condition.0", d)
	}
	if v, ok := d.GetOkExists(key + ".default"); !isEmptyValue(reflect.ValueOf(d.Get(key+".default"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".default"))) {
		request.Default = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".hit_counts"); !isEmptyValue(reflect.ValueOf(d.Get(key+".hit_counts"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".hit_counts"))) {
		request.HitCounts = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".id"))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".name"))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".rank"); !isEmptyValue(reflect.ValueOf(d.Get(key+".rank"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".rank"))) {
		request.Rank = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".state"); !isEmptyValue(reflect.ValueOf(d.Get(key+".state"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".state"))) {
		request.State = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceAdministrationLocalExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDRuleCondition(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDRuleCondition {
	request := isegosdk.RequestDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDRuleCondition{}
	if v, ok := d.GetOkExists(key + ".condition_type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".condition_type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".condition_type"))) {
		request.ConditionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".is_negate"); !isEmptyValue(reflect.ValueOf(d.Get(key+".is_negate"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".is_negate"))) {
		request.IsNegate = interfaceToBoolPtr(v)
	}

	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".id"))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".name"))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".attribute_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".attribute_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".attribute_name"))) {
		request.AttributeName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".attribute_id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".attribute_id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".attribute_id"))) {
		request.AttributeID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".attribute_value"); !isEmptyValue(reflect.ValueOf(d.Get(key+".attribute_value"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".attribute_value"))) {
		request.AttributeValue = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".dictionary_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".dictionary_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".dictionary_name"))) {
		request.DictionaryName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".dictionary_value"); !isEmptyValue(reflect.ValueOf(d.Get(key+".dictionary_value"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".dictionary_value"))) {
		request.DictionaryValue = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".operator"); !isEmptyValue(reflect.ValueOf(d.Get(key+".operator"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".operator"))) {
		request.Operator = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".children"); !isEmptyValue(reflect.ValueOf(d.Get(key+".children"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".children"))) {
		request.Children = expandRequestDeviceAdministrationLocalExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDRuleConditionChildrenArray(ctx, key+".children", d)
	}
	if v, ok := d.GetOkExists(key + ".dates_range"); !isEmptyValue(reflect.ValueOf(d.Get(key+".dates_range"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".dates_range"))) {
		request.DatesRange = expandRequestDeviceAdministrationLocalExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDRuleConditionDatesRange(ctx, key+".dates_range.0", d)
	}
	if v, ok := d.GetOkExists(key + ".dates_range_exception"); !isEmptyValue(reflect.ValueOf(d.Get(key+".dates_range_exception"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".dates_range_exception"))) {
		request.DatesRangeException = expandRequestDeviceAdministrationLocalExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDRuleConditionDatesRangeException(ctx, key+".dates_range_exception.0", d)
	}
	if v, ok := d.GetOkExists(key + ".hours_range"); !isEmptyValue(reflect.ValueOf(d.Get(key+".hours_range"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".hours_range"))) {
		request.HoursRange = expandRequestDeviceAdministrationLocalExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDRuleConditionHoursRange(ctx, key+".hours_range.0", d)
	}
	if v, ok := d.GetOkExists(key + ".hours_range_exception"); !isEmptyValue(reflect.ValueOf(d.Get(key+".hours_range_exception"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".hours_range_exception"))) {
		request.HoursRangeException = expandRequestDeviceAdministrationLocalExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDRuleConditionHoursRangeException(ctx, key+".hours_range_exception.0", d)
	}
	if v, ok := d.GetOkExists(key + ".week_days"); !isEmptyValue(reflect.ValueOf(d.Get(key+".week_days"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".week_days"))) {
		request.WeekDays = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(key + ".week_days_exception"); !isEmptyValue(reflect.ValueOf(d.Get(key+".week_days_exception"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".week_days_exception"))) {
		request.WeekDaysException = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceAdministrationLocalExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDRuleConditionLink(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDRuleConditionLink {
	request := isegosdk.RequestDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDRuleConditionLink{}
	if v, ok := d.GetOkExists(key + ".href"); !isEmptyValue(reflect.ValueOf(d.Get(key+".href"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".href"))) {
		request.Href = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".rel"); !isEmptyValue(reflect.ValueOf(d.Get(key+".rel"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".rel"))) {
		request.Rel = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".type"))) {
		request.Type = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceAdministrationLocalExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDRuleConditionChildrenArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDRuleConditionChildren {
	request := []isegosdk.RequestDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDRuleConditionChildren{}
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestDeviceAdministrationLocalExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDRuleConditionChildren(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		request = append(request, *i)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceAdministrationLocalExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDRuleConditionChildren(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDRuleConditionChildren {
	request := isegosdk.RequestDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDRuleConditionChildren{}
	if v, ok := d.GetOkExists(key + ".condition_type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".condition_type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".condition_type"))) {
		request.ConditionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".is_negate"); !isEmptyValue(reflect.ValueOf(d.Get(key+".is_negate"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".is_negate"))) {
		request.IsNegate = interfaceToBoolPtr(v)
	}

	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceAdministrationLocalExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDRuleConditionChildrenLink(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDRuleConditionChildrenLink {
	request := isegosdk.RequestDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDRuleConditionChildrenLink{}
	if v, ok := d.GetOkExists(key + ".href"); !isEmptyValue(reflect.ValueOf(d.Get(key+".href"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".href"))) {
		request.Href = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".rel"); !isEmptyValue(reflect.ValueOf(d.Get(key+".rel"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".rel"))) {
		request.Rel = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".type"))) {
		request.Type = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceAdministrationLocalExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDRuleConditionDatesRange(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDRuleConditionDatesRange {
	request := isegosdk.RequestDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDRuleConditionDatesRange{}
	if v, ok := d.GetOkExists(key + ".end_date"); !isEmptyValue(reflect.ValueOf(d.Get(key+".end_date"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".end_date"))) {
		request.EndDate = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".start_date"); !isEmptyValue(reflect.ValueOf(d.Get(key+".start_date"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".start_date"))) {
		request.StartDate = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceAdministrationLocalExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDRuleConditionDatesRangeException(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDRuleConditionDatesRangeException {
	request := isegosdk.RequestDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDRuleConditionDatesRangeException{}
	if v, ok := d.GetOkExists(key + ".end_date"); !isEmptyValue(reflect.ValueOf(d.Get(key+".end_date"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".end_date"))) {
		request.EndDate = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".start_date"); !isEmptyValue(reflect.ValueOf(d.Get(key+".start_date"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".start_date"))) {
		request.StartDate = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceAdministrationLocalExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDRuleConditionHoursRange(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDRuleConditionHoursRange {
	request := isegosdk.RequestDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDRuleConditionHoursRange{}
	if v, ok := d.GetOkExists(key + ".end_time"); !isEmptyValue(reflect.ValueOf(d.Get(key+".end_time"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".end_time"))) {
		request.EndTime = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".start_time"); !isEmptyValue(reflect.ValueOf(d.Get(key+".start_time"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".start_time"))) {
		request.StartTime = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceAdministrationLocalExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDRuleConditionHoursRangeException(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDRuleConditionHoursRangeException {
	request := isegosdk.RequestDeviceAdministrationAuthorizationExceptionRulesUpdateDeviceAdminLocalExceptionRuleByIDRuleConditionHoursRangeException{}
	if v, ok := d.GetOkExists(key + ".end_time"); !isEmptyValue(reflect.ValueOf(d.Get(key+".end_time"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".end_time"))) {
		request.EndTime = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".start_time"); !isEmptyValue(reflect.ValueOf(d.Get(key+".start_time"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".start_time"))) {
		request.StartTime = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func getAllItemsDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRules(m interface{}, response *isegosdk.ResponseDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRules, policyTypeID string) []isegosdk.ResponseDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRulesResponse {
	var respItems []isegosdk.ResponseDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRulesResponse
	for response.Response != nil && len(*response.Response) > 0 {
		respItems = append(respItems, *response.Response...)
	}
	return respItems
}

func searchDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRules(m interface{}, items []isegosdk.ResponseDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRulesResponse, name string, id string, policyID string) (*isegosdk.ResponseDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRuleByIDResponse, error) {
	client := m.(*isegosdk.Client)
	var err error
	var foundItem *isegosdk.ResponseDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRuleByIDResponse
	for _, item := range items {
		if id != "" && item.Rule != nil && item.Rule.ID == id {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRuleByID
			getItem, _, err = client.DeviceAdministrationAuthorizationExceptionRules.GetDeviceAdminLocalExceptionRuleByID(policyID, id)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetDeviceAdminLocalExceptionRuleByID")
			}
			foundItem = getItem.Response
			return foundItem, err
		} else if name != "" && item.Rule != nil && item.Rule.Name == name {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseDeviceAdministrationAuthorizationExceptionRulesGetDeviceAdminLocalExceptionRuleByID
			getItem, _, err = client.DeviceAdministrationAuthorizationExceptionRules.GetDeviceAdminLocalExceptionRuleByID(policyID, item.Rule.ID)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetDeviceAdminLocalExceptionRuleByID")
			}
			foundItem = getItem.Response
			return foundItem, err
		}
	}
	return foundItem, err
}
