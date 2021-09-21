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

func resourceNetworkAccessAuthenticationRules() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on Network Access - Authentication Rules.

- Network Access Create authentication rule:



 Rule must include name and condition.


 Condition has hierarchical structure which define a set of conditions for which authentication policy rule could be
match.


 Condition can be either reference to a stored Library condition, using model
ConditionReference


or dynamically built conditions which are not stored in the conditions Library, using models
ConditionAttributes, ConditionAndBlock, ConditionOrBlock
.



- Network Access Update rule.

- Network Access Delete rule.
`,

		CreateContext: resourceNetworkAccessAuthenticationRulesCreate,
		ReadContext:   resourceNetworkAccessAuthenticationRulesRead,
		UpdateContext: resourceNetworkAccessAuthenticationRulesUpdate,
		DeleteContext: resourceNetworkAccessAuthenticationRulesDelete,
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

						"id": &schema.Schema{
							Description: `id path parameter. Rule id`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"identity_source_id": &schema.Schema{
							Description: `Identity source id from the identity stores`,
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"identity_source_name": &schema.Schema{
							Description: `Identity source name from the identity stores`,
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"if_auth_fail": &schema.Schema{
							Description: `Action to perform when authentication fails such as Bad credentials, disabled user and so on`,
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"if_process_fail": &schema.Schema{
							Description: `Action to perform when ISE is uanble to access the identity database`,
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"if_user_not_found": &schema.Schema{
							Description: `Action to perform when user is not found in any of identity stores`,
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
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

func resourceNetworkAccessAuthenticationRulesCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("item"))
	request1 := expandRequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRule(ctx, "item.0", d)
	log.Printf("[DEBUG] request1 => %v", responseInterfaceToString(*request1))

	vPolicyID, okPolicyID := resourceItem["policy_id"]
	vvPolicyID := interfaceToString(vPolicyID)
	vID, okID := resourceItem["id"]
	var vvName string
	if !okID || vID == "" {
		if _, ok := d.GetOk("item.0.rule"); ok {
			if v, ok2 := d.GetOkExists("item.0.rule.0.id"); ok2 {
				vID = interfaceToString(v)
				okID = ok2
			}
		}
	}
	vvID := interfaceToString(vID)
	if _, ok := d.GetOk("item.0.rule"); ok {
		if v, ok2 := d.GetOkExists("item.0.rule.0.name"); ok2 {
			vvName = interfaceToString(v)
		}
	}

	if okPolicyID && vvPolicyID != "" && okID && vvID != "" {
		getResponse2, _, err := client.NetworkAccessAuthenticationRules.GetNetworkAccessAuthenticationRuleByID(vvPolicyID, vvID)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["policy_id"] = vvPolicyID
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return diags
		}
	} else {
		response2, _, err := client.NetworkAccessAuthenticationRules.GetNetworkAccessAuthenticationRules(vvPolicyID)
		if response2 != nil && err == nil {
			items2 := getAllItemsNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRules(m, response2, vvPolicyID)
			item2, err := searchNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRules(m, items2, vvName, vvID, vvPolicyID)
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
	resp1, restyResp1, err := client.NetworkAccessAuthenticationRules.CreateNetworkAccessAuthenticationRule(vvPolicyID, request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateNetworkAccessAuthenticationRule", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateNetworkAccessAuthenticationRule", err))
		return diags
	}
	if resp1.Response != nil && resp1.Response.Rule != nil && vvID != resp1.Response.Rule.ID {
		vvID = resp1.Response.Rule.ID
	}
	if resp1.Response != nil && resp1.Response.Rule != nil && vvName != resp1.Response.Rule.Name {
		vvName = resp1.Response.Rule.Name
	}
	resourceMap := make(map[string]string)
	resourceMap["policy_id"] = vvPolicyID
	resourceMap["id"] = vvID
	resourceMap["name"] = vvName
	d.SetId(joinResourceID(resourceMap))
	return diags
}

func resourceNetworkAccessAuthenticationRulesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vPolicyID, okPolicyID := resourceMap["policy_id"]
	vID, okID := resourceMap["id"]
	vName, okName := resourceMap["name"]

	if !okID || vID == "" {
		if _, ok := d.GetOk("item.0.rule"); ok {
			if v, ok2 := d.GetOkExists("item.0.rule.0.id"); ok2 {
				vID = interfaceToString(v)
			}
		}
	}
	if _, ok := d.GetOk("item.0.rule"); ok {
		if v, ok2 := d.GetOkExists("item.0.rule.0.name"); ok2 {
			vName = interfaceToString(v)
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
		log.Printf("[DEBUG] Selected method: GetNetworkAccessAuthenticationRules")
		response1, _, err := client.NetworkAccessAuthenticationRules.GetNetworkAccessAuthenticationRules(vvPolicyID)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetNetworkAccessAuthenticationRules", err,
				"Failure at GetNetworkAccessAuthenticationRules, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		items1 := getAllItemsNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRules(m, response1, vvPolicyID)
		item1, err := searchNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRules(m, items1, vvName, vvID, vvPolicyID)
		if err != nil || item1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when searching item from GetNetworkAccessAuthenticationRules response", err,
				"Failure when searching item from GetNetworkAccessAuthenticationRules, unexpected response", ""))
			return diags
		}
		if err := d.Set("item", item1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNetworkAccessAuthenticationRules search response",
				err))
			return diags
		}

	}
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkAccessAuthenticationRuleByID")
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
		return diags

	}
	return diags
}

func resourceNetworkAccessAuthenticationRulesUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vPolicyID, okPolicyID := resourceMap["policy_id"]
	vID, okID := resourceMap["id"]
	vName, okName := resourceMap["name"]

	if !okID || vID == "" {
		if _, ok := d.GetOk("item.0.rule"); ok {
			if v, ok2 := d.GetOkExists("item.0.rule.0.id"); ok2 {
				vID = interfaceToString(v)
			}
		}
	}
	if _, ok := d.GetOk("item.0.rule"); ok {
		if v, ok2 := d.GetOkExists("item.0.rule.0.name"); ok2 {
			vName = interfaceToString(v)
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
		getResp1, _, err := client.NetworkAccessAuthenticationRules.GetNetworkAccessAuthenticationRules(vvPolicyID)
		if err == nil && getResp1 != nil {
			items1 := getAllItemsNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRules(m, getResp1, vvPolicyID)
			item1, err := searchNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRules(m, items1, vvName, vvID, vvPolicyID)
			if err == nil && item1 != nil {
				if vID != item1.Rule.ID {
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
		log.Printf("[DEBUG] vvID %s", vvID)
		request1 := expandRequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByID(ctx, "item.0", d)
		log.Printf("[DEBUG] request1 => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.NetworkAccessAuthenticationRules.UpdateNetworkAccessAuthenticationRuleByID(vvPolicyID, vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] restyResp1 => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateNetworkAccessAuthenticationRuleByID", err, restyResp1.String(),
					"Failure at UpdateNetworkAccessAuthenticationRuleByID, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateNetworkAccessAuthenticationRuleByID", err,
				"Failure at UpdateNetworkAccessAuthenticationRuleByID, unexpected response", ""))
			return diags
		}
	}

	return resourceNetworkAccessAuthenticationRulesRead(ctx, d, m)
}

func resourceNetworkAccessAuthenticationRulesDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vPolicyID, okPolicyID := resourceMap["policy_id"]
	vID, okID := resourceMap["id"]
	vName, okName := resourceMap["name"]

	if !okID || vID == "" {
		if _, ok := d.GetOk("item.0.rule"); ok {
			if v, ok2 := d.GetOkExists("item.0.rule.0.id"); ok2 {
				vID = interfaceToString(v)
			}
		}
	}
	if _, ok := d.GetOk("item.0.rule"); ok {
		if v, ok2 := d.GetOkExists("item.0.rule.0.name"); ok2 {
			vName = interfaceToString(v)
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

		getResp1, _, err := client.NetworkAccessAuthenticationRules.GetNetworkAccessAuthenticationRules(vvPolicyID)
		if err != nil || getResp1 == nil {
			// Assume that element it is already gone
			return diags
		}
		items1 := getAllItemsNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRules(m, getResp1, vvPolicyID)
		item1, err := searchNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRules(m, items1, vvName, vvID, vvPolicyID)
		if err != nil || item1 == nil {
			// Assume that element it is already gone
			return diags
		}
		if vID != item1.Rule.ID {
			vvID = item1.Rule.ID
		} else {
			vvID = vID
		}
	}
	if selectedMethod == 1 {
		vvID = vID
		getResp, _, err := client.NetworkAccessAuthenticationRules.GetNetworkAccessAuthenticationRuleByID(vvPolicyID, vvID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	response1, restyResp1, err := client.NetworkAccessAuthenticationRules.DeleteNetworkAccessAuthenticationRuleByID(vvPolicyID, vvID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] restyResp1 => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteNetworkAccessAuthenticationRuleByID", err, restyResp1.String(),
				"Failure at DeleteNetworkAccessAuthenticationRuleByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteNetworkAccessAuthenticationRuleByID", err,
			"Failure at DeleteNetworkAccessAuthenticationRuleByID, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRule(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRule {
	request := isegosdk.RequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRule{}
	if v, ok := d.GetOkExists(key + ".identity_source_id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".identity_source_id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".identity_source_id"))) {
		request.IDentitySourceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".identity_source_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".identity_source_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".identity_source_name"))) {
		request.IDentitySourceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".if_auth_fail"); !isEmptyValue(reflect.ValueOf(d.Get(key+".if_auth_fail"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".if_auth_fail"))) {
		request.IfAuthFail = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".if_process_fail"); !isEmptyValue(reflect.ValueOf(d.Get(key+".if_process_fail"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".if_process_fail"))) {
		request.IfProcessFail = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".if_user_not_found"); !isEmptyValue(reflect.ValueOf(d.Get(key+".if_user_not_found"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".if_user_not_found"))) {
		request.IfUserNotFound = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".link"); !isEmptyValue(reflect.ValueOf(d.Get(key+".link"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".link"))) {
		request.Link = expandRequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleLink(ctx, key+".link.0", d)
	}
	if v, ok := d.GetOkExists(key + ".rule"); !isEmptyValue(reflect.ValueOf(d.Get(key+".rule"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".rule"))) {
		request.Rule = expandRequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleRule(ctx, key+".rule.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleLink(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleLink {
	request := isegosdk.RequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleLink{}
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

func expandRequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleRule(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleRule {
	request := isegosdk.RequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleRule{}
	if v, ok := d.GetOkExists(key + ".condition"); !isEmptyValue(reflect.ValueOf(d.Get(key+".condition"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".condition"))) {
		request.Condition = expandRequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleRuleCondition(ctx, key+".condition.0", d)
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

func expandRequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleRuleCondition(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleRuleCondition {
	request := isegosdk.RequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleRuleCondition{}
	if v, ok := d.GetOkExists(key + ".condition_type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".condition_type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".condition_type"))) {
		request.ConditionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".is_negate"); !isEmptyValue(reflect.ValueOf(d.Get(key+".is_negate"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".is_negate"))) {
		request.IsNegate = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".link"); !isEmptyValue(reflect.ValueOf(d.Get(key+".link"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".link"))) {
		request.Link = expandRequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleRuleConditionLink(ctx, key+".link.0", d)
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
		request.Children = expandRequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleRuleConditionChildrenArray(ctx, key, d)
	}
	if v, ok := d.GetOkExists(key + ".dates_range"); !isEmptyValue(reflect.ValueOf(d.Get(key+".dates_range"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".dates_range"))) {
		request.DatesRange = expandRequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleRuleConditionDatesRange(ctx, key+".dates_range.0", d)
	}
	if v, ok := d.GetOkExists(key + ".dates_range_exception"); !isEmptyValue(reflect.ValueOf(d.Get(key+".dates_range_exception"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".dates_range_exception"))) {
		request.DatesRangeException = expandRequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleRuleConditionDatesRangeException(ctx, key+".dates_range_exception.0", d)
	}
	if v, ok := d.GetOkExists(key + ".hours_range"); !isEmptyValue(reflect.ValueOf(d.Get(key+".hours_range"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".hours_range"))) {
		request.HoursRange = expandRequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleRuleConditionHoursRange(ctx, key+".hours_range.0", d)
	}
	if v, ok := d.GetOkExists(key + ".hours_range_exception"); !isEmptyValue(reflect.ValueOf(d.Get(key+".hours_range_exception"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".hours_range_exception"))) {
		request.HoursRangeException = expandRequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleRuleConditionHoursRangeException(ctx, key+".hours_range_exception.0", d)
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

func expandRequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleRuleConditionLink(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleRuleConditionLink {
	request := isegosdk.RequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleRuleConditionLink{}
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

func expandRequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleRuleConditionChildrenArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleRuleConditionChildren {
	request := []isegosdk.RequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleRuleConditionChildren{}
	o := d.Get(key)
	if o != nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleRuleConditionChildren(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		request = append(request, *i)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleRuleConditionChildren(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleRuleConditionChildren {
	request := isegosdk.RequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleRuleConditionChildren{}
	if v, ok := d.GetOkExists(key + ".condition_type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".condition_type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".condition_type"))) {
		request.ConditionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".is_negate"); !isEmptyValue(reflect.ValueOf(d.Get(key+".is_negate"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".is_negate"))) {
		request.IsNegate = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".link"); !isEmptyValue(reflect.ValueOf(d.Get(key+".link"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".link"))) {
		request.Link = expandRequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleRuleConditionChildrenLink(ctx, key+".link.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleRuleConditionChildrenLink(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleRuleConditionChildrenLink {
	request := isegosdk.RequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleRuleConditionChildrenLink{}
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

func expandRequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleRuleConditionDatesRange(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleRuleConditionDatesRange {
	request := isegosdk.RequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleRuleConditionDatesRange{}
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

func expandRequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleRuleConditionDatesRangeException(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleRuleConditionDatesRangeException {
	request := isegosdk.RequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleRuleConditionDatesRangeException{}
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

func expandRequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleRuleConditionHoursRange(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleRuleConditionHoursRange {
	request := isegosdk.RequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleRuleConditionHoursRange{}
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

func expandRequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleRuleConditionHoursRangeException(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleRuleConditionHoursRangeException {
	request := isegosdk.RequestNetworkAccessAuthenticationRulesCreateNetworkAccessAuthenticationRuleRuleConditionHoursRangeException{}
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

func expandRequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByID(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByID {
	request := isegosdk.RequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByID{}
	if v, ok := d.GetOkExists(key + ".identity_source_id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".identity_source_id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".identity_source_id"))) {
		request.IDentitySourceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".identity_source_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".identity_source_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".identity_source_name"))) {
		request.IDentitySourceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".if_auth_fail"); !isEmptyValue(reflect.ValueOf(d.Get(key+".if_auth_fail"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".if_auth_fail"))) {
		request.IfAuthFail = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".if_process_fail"); !isEmptyValue(reflect.ValueOf(d.Get(key+".if_process_fail"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".if_process_fail"))) {
		request.IfProcessFail = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".if_user_not_found"); !isEmptyValue(reflect.ValueOf(d.Get(key+".if_user_not_found"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".if_user_not_found"))) {
		request.IfUserNotFound = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".link"); !isEmptyValue(reflect.ValueOf(d.Get(key+".link"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".link"))) {
		request.Link = expandRequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDLink(ctx, key+".link.0", d)
	}
	if v, ok := d.GetOkExists(key + ".rule"); !isEmptyValue(reflect.ValueOf(d.Get(key+".rule"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".rule"))) {
		request.Rule = expandRequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDRule(ctx, key+".rule.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDLink(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDLink {
	request := isegosdk.RequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDLink{}
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

func expandRequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDRule(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDRule {
	request := isegosdk.RequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDRule{}
	if v, ok := d.GetOkExists(key + ".condition"); !isEmptyValue(reflect.ValueOf(d.Get(key+".condition"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".condition"))) {
		request.Condition = expandRequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDRuleCondition(ctx, key+".condition.0", d)
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

func expandRequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDRuleCondition(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDRuleCondition {
	request := isegosdk.RequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDRuleCondition{}
	if v, ok := d.GetOkExists(key + ".condition_type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".condition_type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".condition_type"))) {
		request.ConditionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".is_negate"); !isEmptyValue(reflect.ValueOf(d.Get(key+".is_negate"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".is_negate"))) {
		request.IsNegate = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".link"); !isEmptyValue(reflect.ValueOf(d.Get(key+".link"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".link"))) {
		request.Link = expandRequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDRuleConditionLink(ctx, key+".link.0", d)
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
		request.Children = expandRequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDRuleConditionChildrenArray(ctx, key, d)
	}
	if v, ok := d.GetOkExists(key + ".dates_range"); !isEmptyValue(reflect.ValueOf(d.Get(key+".dates_range"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".dates_range"))) {
		request.DatesRange = expandRequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDRuleConditionDatesRange(ctx, key+".dates_range.0", d)
	}
	if v, ok := d.GetOkExists(key + ".dates_range_exception"); !isEmptyValue(reflect.ValueOf(d.Get(key+".dates_range_exception"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".dates_range_exception"))) {
		request.DatesRangeException = expandRequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDRuleConditionDatesRangeException(ctx, key+".dates_range_exception.0", d)
	}
	if v, ok := d.GetOkExists(key + ".hours_range"); !isEmptyValue(reflect.ValueOf(d.Get(key+".hours_range"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".hours_range"))) {
		request.HoursRange = expandRequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDRuleConditionHoursRange(ctx, key+".hours_range.0", d)
	}
	if v, ok := d.GetOkExists(key + ".hours_range_exception"); !isEmptyValue(reflect.ValueOf(d.Get(key+".hours_range_exception"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".hours_range_exception"))) {
		request.HoursRangeException = expandRequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDRuleConditionHoursRangeException(ctx, key+".hours_range_exception.0", d)
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

func expandRequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDRuleConditionLink(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDRuleConditionLink {
	request := isegosdk.RequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDRuleConditionLink{}
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

func expandRequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDRuleConditionChildrenArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDRuleConditionChildren {
	request := []isegosdk.RequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDRuleConditionChildren{}
	o := d.Get(key)
	if o != nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDRuleConditionChildren(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		request = append(request, *i)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDRuleConditionChildren(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDRuleConditionChildren {
	request := isegosdk.RequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDRuleConditionChildren{}
	if v, ok := d.GetOkExists(key + ".condition_type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".condition_type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".condition_type"))) {
		request.ConditionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".is_negate"); !isEmptyValue(reflect.ValueOf(d.Get(key+".is_negate"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".is_negate"))) {
		request.IsNegate = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".link"); !isEmptyValue(reflect.ValueOf(d.Get(key+".link"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".link"))) {
		request.Link = expandRequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDRuleConditionChildrenLink(ctx, key+".link.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDRuleConditionChildrenLink(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDRuleConditionChildrenLink {
	request := isegosdk.RequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDRuleConditionChildrenLink{}
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

func expandRequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDRuleConditionDatesRange(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDRuleConditionDatesRange {
	request := isegosdk.RequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDRuleConditionDatesRange{}
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

func expandRequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDRuleConditionDatesRangeException(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDRuleConditionDatesRangeException {
	request := isegosdk.RequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDRuleConditionDatesRangeException{}
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

func expandRequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDRuleConditionHoursRange(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDRuleConditionHoursRange {
	request := isegosdk.RequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDRuleConditionHoursRange{}
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

func expandRequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDRuleConditionHoursRangeException(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDRuleConditionHoursRangeException {
	request := isegosdk.RequestNetworkAccessAuthenticationRulesUpdateNetworkAccessAuthenticationRuleByIDRuleConditionHoursRangeException{}
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

func getAllItemsNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRules(m interface{}, response *isegosdk.ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRules, policyTypeID string) []isegosdk.ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRulesResponse {
	var respItems []isegosdk.ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRulesResponse
	for response.Response != nil && len(*response.Response) > 0 {
		respItems = append(respItems, *response.Response...)
	}
	return respItems
}

func searchNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRules(m interface{}, items []isegosdk.ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRulesResponse, name string, id string, policyID string) (*isegosdk.ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRuleByIDResponse, error) {
	client := m.(*isegosdk.Client)
	var err error
	var foundItem *isegosdk.ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRuleByIDResponse
	for _, item := range items {
		if id != "" && item.Rule != nil && item.Rule.ID == id {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRuleByID
			getItem, _, err = client.NetworkAccessAuthenticationRules.GetNetworkAccessAuthenticationRuleByID(policyID, id)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetNetworkAccessAuthenticationRuleByID")
			}
			foundItem = getItem.Response
			return foundItem, err
		} else if name != "" && item.Rule != nil && item.Rule.Name == name {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseNetworkAccessAuthenticationRulesGetNetworkAccessAuthenticationRuleByID
			getItem, _, err = client.NetworkAccessAuthenticationRules.GetNetworkAccessAuthenticationRuleByID(policyID, item.Rule.ID)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetNetworkAccessAuthenticationRuleByID")
			}
			foundItem = getItem.Response
			return foundItem, err
		}
	}
	return foundItem, err
}
