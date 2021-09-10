package ciscoise

import (
	"context"
	"fmt"
	"reflect"

	"ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDeviceAdministrationGlobalExceptionRules() *schema.Resource {
	return &schema.Resource{

		CreateContext: resourceDeviceAdministrationGlobalExceptionRulesCreate,
		ReadContext:   resourceDeviceAdministrationGlobalExceptionRulesRead,
		UpdateContext: resourceDeviceAdministrationGlobalExceptionRulesUpdate,
		DeleteContext: resourceDeviceAdministrationGlobalExceptionRulesDelete,
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
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"href": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"link": &schema.Schema{
							Type:             schema.TypeList,
							DiffSuppressFunc: diffSuppressAlways(),
							Computed:         true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"href": &schema.Schema{
										Type:             schema.TypeString,
										DiffSuppressFunc: diffSuppressAlways(),
										Computed:         true,
									},
									"rel": &schema.Schema{
										Type:             schema.TypeString,
										DiffSuppressFunc: diffSuppressAlways(),
										Computed:         true,
									},
									"type": &schema.Schema{
										Type:             schema.TypeString,
										DiffSuppressFunc: diffSuppressAlways(),
										Computed:         true,
									},
								},
							},
						},
						"profile": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"rel": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourceDeviceAdministrationGlobalExceptionRulesCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("item"))
	request1 := expandRequestDeviceAdministrationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalException(ctx, "item.0", d)
	log.Printf("[DEBUG] request1 => %v", responseInterfaceToString(*request1))

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

	if okID && vvID != "" {
		getResponse2, _, err := client.DeviceAdministrationAuthorizationGlobalExceptionRules.GetDeviceAdminPolicySetGlobalExceptionByRuleID(vvID)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return diags
		}
	} else {
		response2, _, err := client.DeviceAdministrationAuthorizationGlobalExceptionRules.GetDeviceAdminPolicySetGlobalExceptionRules()
		if response2 != nil && err == nil {
			items2 := getAllItemsDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionRules(m, response2)
			item2, err := searchDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionRules(m, items2, vvName, vvID)
			if err == nil && item2 != nil {
				resourceMap := make(map[string]string)
				resourceMap["id"] = vvID
				resourceMap["name"] = vvName
				d.SetId(joinResourceID(resourceMap))
				return diags
			}
		}
	}
	resp1, restyResp1, err := client.DeviceAdministrationAuthorizationGlobalExceptionRules.CreateDeviceAdminPolicySetGlobalException(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateDeviceAdminPolicySetGlobalException", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateDeviceAdminPolicySetGlobalException", err))
		return diags
	}
	if vvID != resp1.Response.Rule.ID {
		vvID = resp1.Response.Rule.ID
	}
	if vvName != resp1.Response.Rule.Name {
		vvName = resp1.Response.Rule.Name
	}
	resourceMap := make(map[string]string)
	resourceMap["id"] = vvID
	resourceMap["name"] = vvName
	d.SetId(joinResourceID(resourceMap))
	return diags
}

func resourceDeviceAdministrationGlobalExceptionRulesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
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
	vvName := vName
	method1 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetDeviceAdminPolicySetGlobalExceptionRules")

		response1, _, err := client.DeviceAdministrationAuthorizationGlobalExceptionRules.GetDeviceAdminPolicySetGlobalExceptionRules()

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDeviceAdminPolicySetGlobalExceptionRules", err,
				"Failure at GetDeviceAdminPolicySetGlobalExceptionRules, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		items1 := getAllItemsDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionRules(m, response1)
		item1, err := searchDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionRules(m, items1, vvName, vvID)
		if err != nil || item1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when searching item from GetDeviceAdminPolicySetGlobalExceptionRules response", err,
				"Failure when searching item from GetDeviceAdminPolicySetGlobalExceptionRules, unexpected response", ""))
			return diags
		}
		if err := d.Set("item", item1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceAdminPolicySetGlobalExceptionRules search response",
				err))
			return diags
		}

	}
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDeviceAdminPolicySetGlobalExceptionByRuleID")

		response2, _, err := client.DeviceAdministrationAuthorizationGlobalExceptionRules.GetDeviceAdminPolicySetGlobalExceptionByRuleID(vvID)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDeviceAdminPolicySetGlobalExceptionByRuleID", err,
				"Failure at GetDeviceAdminPolicySetGlobalExceptionByRuleID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

		vItem2 := flattenDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionByRuleIDItem(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceAdminPolicySetGlobalExceptionByRuleID response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceDeviceAdministrationGlobalExceptionRulesUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
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
	vvName := vName
	method1 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	// NOTE: Consider adding getAllItems and search function to get missing params
	if selectedMethod == 2 {
		getResp1, _, err := client.DeviceAdministrationAuthorizationGlobalExceptionRules.GetDeviceAdminPolicySetGlobalExceptionRules()
		if err == nil && getResp1 != nil {
			items1 := getAllItemsDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionRules(m, getResp1)
			item1, err := searchDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionRules(m, items1, vvName, vvID)
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
		log.Printf("[DEBUG] vvID %s", vvID)
		request1 := expandRequestDeviceAdministrationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleID(ctx, "item.0", d)
		log.Printf("[DEBUG] request1 => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.DeviceAdministrationAuthorizationGlobalExceptionRules.UpdateDeviceAdminPolicySetGlobalExceptionByRuleID(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] restyResp1 => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateDeviceAdminPolicySetGlobalExceptionByRuleID", err, restyResp1.String(),
					"Failure at UpdateDeviceAdminPolicySetGlobalExceptionByRuleID, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateDeviceAdminPolicySetGlobalExceptionByRuleID", err,
				"Failure at UpdateDeviceAdminPolicySetGlobalExceptionByRuleID, unexpected response", ""))
			return diags
		}
	}

	return resourceDeviceAdministrationGlobalExceptionRulesRead(ctx, d, m)
}

func resourceDeviceAdministrationGlobalExceptionRulesDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
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
	vvName := vName
	method1 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	// REVIEW: Add getAllItems and search function to get missing params
	if selectedMethod == 2 {

		getResp1, _, err := client.DeviceAdministrationAuthorizationGlobalExceptionRules.GetDeviceAdminPolicySetGlobalExceptionRules()
		if err != nil || getResp1 == nil {
			// Assume that element it is already gone
			return diags
		}
		items1 := getAllItemsDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionRules(m, getResp1)
		item1, err := searchDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionRules(m, items1, vvName, vvID)
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
		getResp, _, err := client.DeviceAdministrationAuthorizationGlobalExceptionRules.GetDeviceAdminPolicySetGlobalExceptionByRuleID(vvID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	response1, restyResp1, err := client.DeviceAdministrationAuthorizationGlobalExceptionRules.DeleteDeviceAdminPolicySetGlobalExceptionByRuleID(vvID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] restyResp1 => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteDeviceAdminPolicySetGlobalExceptionByRuleID", err, restyResp1.String(),
				"Failure at DeleteDeviceAdminPolicySetGlobalExceptionByRuleID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteDeviceAdminPolicySetGlobalExceptionByRuleID", err,
			"Failure at DeleteDeviceAdminPolicySetGlobalExceptionByRuleID, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestDeviceAdministrationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalException(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalException {
	request := isegosdk.RequestDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalException{}
	if v, ok := d.GetOkExists(key + ".commands"); !isEmptyValue(reflect.ValueOf(d.Get(key+".commands"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".commands"))) {
		request.Commands = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(key + ".link"); !isEmptyValue(reflect.ValueOf(d.Get(key+".link"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".link"))) {
		request.Link = expandRequestDeviceAdministrationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionLink(ctx, key+".link.0", d)
	}
	if v, ok := d.GetOkExists(key + ".profile"); !isEmptyValue(reflect.ValueOf(d.Get(key+".profile"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".profile"))) {
		request.Profile = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".rule"); !isEmptyValue(reflect.ValueOf(d.Get(key+".rule"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".rule"))) {
		request.Rule = expandRequestDeviceAdministrationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionRule(ctx, key+".rule.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceAdministrationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionLink(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionLink {
	request := isegosdk.RequestDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionLink{}
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

func expandRequestDeviceAdministrationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionRule(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionRule {
	request := isegosdk.RequestDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionRule{}
	if v, ok := d.GetOkExists(key + ".condition"); !isEmptyValue(reflect.ValueOf(d.Get(key+".condition"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".condition"))) {
		request.Condition = expandRequestDeviceAdministrationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionRuleCondition(ctx, key+".condition.0", d)
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

func expandRequestDeviceAdministrationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionRuleCondition(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionRuleCondition {
	request := isegosdk.RequestDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionRuleCondition{}
	if v, ok := d.GetOkExists(key + ".condition_type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".condition_type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".condition_type"))) {
		request.ConditionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".is_negate"); !isEmptyValue(reflect.ValueOf(d.Get(key+".is_negate"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".is_negate"))) {
		request.IsNegate = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".link"); !isEmptyValue(reflect.ValueOf(d.Get(key+".link"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".link"))) {
		request.Link = expandRequestDeviceAdministrationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionRuleConditionLink(ctx, key+".link.0", d)
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
		request.Children = expandRequestDeviceAdministrationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionRuleConditionChildrenArray(ctx, key, d)
	}
	if v, ok := d.GetOkExists(key + ".dates_range"); !isEmptyValue(reflect.ValueOf(d.Get(key+".dates_range"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".dates_range"))) {
		request.DatesRange = expandRequestDeviceAdministrationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionRuleConditionDatesRange(ctx, key+".dates_range.0", d)
	}
	if v, ok := d.GetOkExists(key + ".dates_range_exception"); !isEmptyValue(reflect.ValueOf(d.Get(key+".dates_range_exception"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".dates_range_exception"))) {
		request.DatesRangeException = expandRequestDeviceAdministrationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionRuleConditionDatesRangeException(ctx, key+".dates_range_exception.0", d)
	}
	if v, ok := d.GetOkExists(key + ".hours_range"); !isEmptyValue(reflect.ValueOf(d.Get(key+".hours_range"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".hours_range"))) {
		request.HoursRange = expandRequestDeviceAdministrationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionRuleConditionHoursRange(ctx, key+".hours_range.0", d)
	}
	if v, ok := d.GetOkExists(key + ".hours_range_exception"); !isEmptyValue(reflect.ValueOf(d.Get(key+".hours_range_exception"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".hours_range_exception"))) {
		request.HoursRangeException = expandRequestDeviceAdministrationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionRuleConditionHoursRangeException(ctx, key+".hours_range_exception.0", d)
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

func expandRequestDeviceAdministrationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionRuleConditionLink(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionRuleConditionLink {
	request := isegosdk.RequestDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionRuleConditionLink{}
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

func expandRequestDeviceAdministrationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionRuleConditionChildrenArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionRuleConditionChildren {
	request := []isegosdk.RequestDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionRuleConditionChildren{}
	o := d.Get(key)
	if o != nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestDeviceAdministrationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionRuleConditionChildren(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		request = append(request, *i)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceAdministrationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionRuleConditionChildren(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionRuleConditionChildren {
	request := isegosdk.RequestDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionRuleConditionChildren{}
	if v, ok := d.GetOkExists(key + ".condition_type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".condition_type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".condition_type"))) {
		request.ConditionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".is_negate"); !isEmptyValue(reflect.ValueOf(d.Get(key+".is_negate"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".is_negate"))) {
		request.IsNegate = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".link"); !isEmptyValue(reflect.ValueOf(d.Get(key+".link"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".link"))) {
		request.Link = expandRequestDeviceAdministrationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionRuleConditionChildrenLink(ctx, key+".link.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceAdministrationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionRuleConditionChildrenLink(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionRuleConditionChildrenLink {
	request := isegosdk.RequestDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionRuleConditionChildrenLink{}
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

func expandRequestDeviceAdministrationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionRuleConditionDatesRange(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionRuleConditionDatesRange {
	request := isegosdk.RequestDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionRuleConditionDatesRange{}
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

func expandRequestDeviceAdministrationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionRuleConditionDatesRangeException(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionRuleConditionDatesRangeException {
	request := isegosdk.RequestDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionRuleConditionDatesRangeException{}
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

func expandRequestDeviceAdministrationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionRuleConditionHoursRange(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionRuleConditionHoursRange {
	request := isegosdk.RequestDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionRuleConditionHoursRange{}
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

func expandRequestDeviceAdministrationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionRuleConditionHoursRangeException(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionRuleConditionHoursRangeException {
	request := isegosdk.RequestDeviceAdministrationAuthorizationGlobalExceptionRulesCreateDeviceAdminPolicySetGlobalExceptionRuleConditionHoursRangeException{}
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

func expandRequestDeviceAdministrationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleID(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleID {
	request := isegosdk.RequestDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleID{}
	if v, ok := d.GetOkExists(key + ".commands"); !isEmptyValue(reflect.ValueOf(d.Get(key+".commands"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".commands"))) {
		request.Commands = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(key + ".link"); !isEmptyValue(reflect.ValueOf(d.Get(key+".link"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".link"))) {
		request.Link = expandRequestDeviceAdministrationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDLink(ctx, key+".link.0", d)
	}
	if v, ok := d.GetOkExists(key + ".profile"); !isEmptyValue(reflect.ValueOf(d.Get(key+".profile"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".profile"))) {
		request.Profile = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".rule"); !isEmptyValue(reflect.ValueOf(d.Get(key+".rule"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".rule"))) {
		request.Rule = expandRequestDeviceAdministrationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDRule(ctx, key+".rule.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceAdministrationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDLink(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDLink {
	request := isegosdk.RequestDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDLink{}
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

func expandRequestDeviceAdministrationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDRule(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDRule {
	request := isegosdk.RequestDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDRule{}
	if v, ok := d.GetOkExists(key + ".condition"); !isEmptyValue(reflect.ValueOf(d.Get(key+".condition"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".condition"))) {
		request.Condition = expandRequestDeviceAdministrationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDRuleCondition(ctx, key+".condition.0", d)
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

func expandRequestDeviceAdministrationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDRuleCondition(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDRuleCondition {
	request := isegosdk.RequestDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDRuleCondition{}
	if v, ok := d.GetOkExists(key + ".condition_type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".condition_type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".condition_type"))) {
		request.ConditionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".is_negate"); !isEmptyValue(reflect.ValueOf(d.Get(key+".is_negate"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".is_negate"))) {
		request.IsNegate = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".link"); !isEmptyValue(reflect.ValueOf(d.Get(key+".link"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".link"))) {
		request.Link = expandRequestDeviceAdministrationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDRuleConditionLink(ctx, key+".link.0", d)
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
		request.Children = expandRequestDeviceAdministrationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDRuleConditionChildrenArray(ctx, key, d)
	}
	if v, ok := d.GetOkExists(key + ".dates_range"); !isEmptyValue(reflect.ValueOf(d.Get(key+".dates_range"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".dates_range"))) {
		request.DatesRange = expandRequestDeviceAdministrationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDRuleConditionDatesRange(ctx, key+".dates_range.0", d)
	}
	if v, ok := d.GetOkExists(key + ".dates_range_exception"); !isEmptyValue(reflect.ValueOf(d.Get(key+".dates_range_exception"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".dates_range_exception"))) {
		request.DatesRangeException = expandRequestDeviceAdministrationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDRuleConditionDatesRangeException(ctx, key+".dates_range_exception.0", d)
	}
	if v, ok := d.GetOkExists(key + ".hours_range"); !isEmptyValue(reflect.ValueOf(d.Get(key+".hours_range"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".hours_range"))) {
		request.HoursRange = expandRequestDeviceAdministrationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDRuleConditionHoursRange(ctx, key+".hours_range.0", d)
	}
	if v, ok := d.GetOkExists(key + ".hours_range_exception"); !isEmptyValue(reflect.ValueOf(d.Get(key+".hours_range_exception"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".hours_range_exception"))) {
		request.HoursRangeException = expandRequestDeviceAdministrationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDRuleConditionHoursRangeException(ctx, key+".hours_range_exception.0", d)
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

func expandRequestDeviceAdministrationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDRuleConditionLink(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDRuleConditionLink {
	request := isegosdk.RequestDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDRuleConditionLink{}
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

func expandRequestDeviceAdministrationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDRuleConditionChildrenArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDRuleConditionChildren {
	request := []isegosdk.RequestDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDRuleConditionChildren{}
	o := d.Get(key)
	if o != nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestDeviceAdministrationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDRuleConditionChildren(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		request = append(request, *i)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceAdministrationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDRuleConditionChildren(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDRuleConditionChildren {
	request := isegosdk.RequestDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDRuleConditionChildren{}
	if v, ok := d.GetOkExists(key + ".condition_type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".condition_type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".condition_type"))) {
		request.ConditionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".is_negate"); !isEmptyValue(reflect.ValueOf(d.Get(key+".is_negate"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".is_negate"))) {
		request.IsNegate = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".link"); !isEmptyValue(reflect.ValueOf(d.Get(key+".link"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".link"))) {
		request.Link = expandRequestDeviceAdministrationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDRuleConditionChildrenLink(ctx, key+".link.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceAdministrationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDRuleConditionChildrenLink(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDRuleConditionChildrenLink {
	request := isegosdk.RequestDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDRuleConditionChildrenLink{}
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

func expandRequestDeviceAdministrationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDRuleConditionDatesRange(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDRuleConditionDatesRange {
	request := isegosdk.RequestDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDRuleConditionDatesRange{}
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

func expandRequestDeviceAdministrationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDRuleConditionDatesRangeException(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDRuleConditionDatesRangeException {
	request := isegosdk.RequestDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDRuleConditionDatesRangeException{}
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

func expandRequestDeviceAdministrationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDRuleConditionHoursRange(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDRuleConditionHoursRange {
	request := isegosdk.RequestDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDRuleConditionHoursRange{}
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

func expandRequestDeviceAdministrationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDRuleConditionHoursRangeException(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDRuleConditionHoursRangeException {
	request := isegosdk.RequestDeviceAdministrationAuthorizationGlobalExceptionRulesUpdateDeviceAdminPolicySetGlobalExceptionByRuleIDRuleConditionHoursRangeException{}
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

func getAllItemsDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionRules(m interface{}, response *isegosdk.ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionRules) []isegosdk.ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionRulesResponse {
	var respItems []isegosdk.ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionRulesResponse
	for response.Response != nil && len(*response.Response) > 0 {
		respItems = append(respItems, *response.Response...)
	}
	return respItems
}

func searchDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionRules(m interface{}, items []isegosdk.ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionRulesResponse, name string, id string) (*isegosdk.ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionByRuleIDResponse, error) {
	client := m.(*isegosdk.Client)
	var err error
	var foundItem *isegosdk.ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionByRuleIDResponse
	for _, item := range items {
		if id != "" && item.Rule != nil && item.Rule.ID == id {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionByRuleID
			getItem, _, err = client.DeviceAdministrationAuthorizationGlobalExceptionRules.GetDeviceAdminPolicySetGlobalExceptionByRuleID(id)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetDeviceAdminPolicySetGlobalExceptionByRuleID")
			}
			foundItem = getItem.Response
			return foundItem, err
		} else if name != "" && item.Rule != nil && item.Rule.Name == name {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseDeviceAdministrationAuthorizationGlobalExceptionRulesGetDeviceAdminPolicySetGlobalExceptionByRuleID
			getItem, _, err = client.DeviceAdministrationAuthorizationGlobalExceptionRules.GetDeviceAdminPolicySetGlobalExceptionByRuleID(item.Rule.ID)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetDeviceAdminPolicySetGlobalExceptionByRuleID")
			}
			foundItem = getItem.Response
			return foundItem, err
		}
	}
	return foundItem, err
}
