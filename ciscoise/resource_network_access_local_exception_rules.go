package ciscoise

import (
	"context"
	"fmt"
	"reflect"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkAccessLocalExceptionRules() *schema.Resource {
	return &schema.Resource{

		CreateContext: resourceNetworkAccessLocalExceptionRulesCreate,
		ReadContext:   resourceNetworkAccessLocalExceptionRulesRead,
		UpdateContext: resourceNetworkAccessLocalExceptionRulesUpdate,
		DeleteContext: resourceNetworkAccessLocalExceptionRulesDelete,
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
						"policy_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"profile": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
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
						"security_group": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
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

func resourceNetworkAccessLocalExceptionRulesCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("item"))
	request1 := expandRequestNetworkAccessLocalExceptionRulesCreateNetworkAccessLocalExceptionRule(ctx, "item.0", d)
	log.Printf("[DEBUG] request1 => %v", responseInterfaceToString(*request1))

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
		getResponse2, _, err := client.NetworkAccessAuthorizationExceptionRules.GetNetworkAccessLocalExceptionRuleByID(vvPolicyID, vvID)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["policy_id"] = vvPolicyID
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return diags
		}
	} else {
		response2, _, err := client.NetworkAccessAuthorizationExceptionRules.GetNetworkAccessLocalExceptionRules(vvPolicyID)
		if response2 != nil && err == nil {
			items2 := getAllItemsNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRules(m, response2, vvPolicyID)
			item2, err := searchNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRules(m, items2, vvName, vvID, vvPolicyID)
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
	resp1, restyResp1, err := client.NetworkAccessAuthorizationExceptionRules.CreateNetworkAccessLocalExceptionRule(vvPolicyID, request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateNetworkAccessLocalExceptionRule", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateNetworkAccessLocalExceptionRule", err))
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

func resourceNetworkAccessLocalExceptionRulesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okPolicyID, okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetNetworkAccessLocalExceptionRules")

		response1, _, err := client.NetworkAccessAuthorizationExceptionRules.GetNetworkAccessLocalExceptionRules(vvPolicyID)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetNetworkAccessLocalExceptionRules", err,
				"Failure at GetNetworkAccessLocalExceptionRules, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		items1 := getAllItemsNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRules(m, response1, vvPolicyID)
		item1, err := searchNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRules(m, items1, vvName, vvID, vvPolicyID)
		if err != nil || item1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when searching item from GetNetworkAccessLocalExceptionRules response", err,
				"Failure when searching item from GetNetworkAccessLocalExceptionRules, unexpected response", ""))
			return diags
		}
		if err := d.Set("item", item1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNetworkAccessLocalExceptionRules search response",
				err))
			return diags
		}

	}
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkAccessLocalExceptionRuleByID")

		response2, _, err := client.NetworkAccessAuthorizationExceptionRules.GetNetworkAccessLocalExceptionRuleByID(vvPolicyID, vvID)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetNetworkAccessLocalExceptionRuleByID", err,
				"Failure at GetNetworkAccessLocalExceptionRuleByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

		vItem2 := flattenNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRuleByIDItem(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNetworkAccessLocalExceptionRuleByID response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceNetworkAccessLocalExceptionRulesUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okPolicyID, okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	// NOTE: Consider adding getAllItems and search function to get missing params
	if selectedMethod == 2 {
		getResp1, _, err := client.NetworkAccessAuthorizationExceptionRules.GetNetworkAccessLocalExceptionRules(vvPolicyID)
		if err == nil && getResp1 != nil {
			items1 := getAllItemsNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRules(m, getResp1, vvPolicyID)
			item1, err := searchNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRules(m, items1, vvName, vvID, vvPolicyID)
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
		request1 := expandRequestNetworkAccessLocalExceptionRulesUpdateNetworkAccessLocalExceptionRuleByID(ctx, "item.0", d)
		log.Printf("[DEBUG] request1 => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.NetworkAccessAuthorizationExceptionRules.UpdateNetworkAccessLocalExceptionRuleByID(vvPolicyID, vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] restyResp1 => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateNetworkAccessLocalExceptionRuleByID", err, restyResp1.String(),
					"Failure at UpdateNetworkAccessLocalExceptionRuleByID, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateNetworkAccessLocalExceptionRuleByID", err,
				"Failure at UpdateNetworkAccessLocalExceptionRuleByID, unexpected response", ""))
			return diags
		}
	}

	return resourceNetworkAccessLocalExceptionRulesRead(ctx, d, m)
}

func resourceNetworkAccessLocalExceptionRulesDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okPolicyID, okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	// REVIEW: Add getAllItems and search function to get missing params
	if selectedMethod == 2 {

		getResp1, _, err := client.NetworkAccessAuthorizationExceptionRules.GetNetworkAccessLocalExceptionRules(vvPolicyID)
		if err != nil || getResp1 == nil {
			// Assume that element it is already gone
			return diags
		}
		items1 := getAllItemsNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRules(m, getResp1, vvPolicyID)
		item1, err := searchNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRules(m, items1, vvName, vvID, vvPolicyID)
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
		getResp, _, err := client.NetworkAccessAuthorizationExceptionRules.GetNetworkAccessLocalExceptionRuleByID(vvPolicyID, vvID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	response1, restyResp1, err := client.NetworkAccessAuthorizationExceptionRules.DeleteNetworkAccessLocalExceptionRuleByID(vvPolicyID, vvID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] restyResp1 => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteNetworkAccessLocalExceptionRuleByID", err, restyResp1.String(),
				"Failure at DeleteNetworkAccessLocalExceptionRuleByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteNetworkAccessLocalExceptionRuleByID", err,
			"Failure at DeleteNetworkAccessLocalExceptionRuleByID, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestNetworkAccessLocalExceptionRulesCreateNetworkAccessLocalExceptionRule(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRule {
	request := isegosdk.RequestNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRule{}
	if v, ok := d.GetOkExists(key + ".link"); !isEmptyValue(reflect.ValueOf(d.Get(key+".link"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".link"))) {
		request.Link = expandRequestNetworkAccessLocalExceptionRulesCreateNetworkAccessLocalExceptionRuleLink(ctx, key+".link.0", d)
	}
	if v, ok := d.GetOkExists(key + ".profile"); !isEmptyValue(reflect.ValueOf(d.Get(key+".profile"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".profile"))) {
		request.Profile = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(key + ".rule"); !isEmptyValue(reflect.ValueOf(d.Get(key+".rule"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".rule"))) {
		request.Rule = expandRequestNetworkAccessLocalExceptionRulesCreateNetworkAccessLocalExceptionRuleRule(ctx, key+".rule.0", d)
	}
	if v, ok := d.GetOkExists(key + ".security_group"); !isEmptyValue(reflect.ValueOf(d.Get(key+".security_group"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".security_group"))) {
		request.SecurityGroup = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkAccessLocalExceptionRulesCreateNetworkAccessLocalExceptionRuleLink(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRuleLink {
	request := isegosdk.RequestNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRuleLink{}
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

func expandRequestNetworkAccessLocalExceptionRulesCreateNetworkAccessLocalExceptionRuleRule(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRuleRule {
	request := isegosdk.RequestNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRuleRule{}
	if v, ok := d.GetOkExists(key + ".condition"); !isEmptyValue(reflect.ValueOf(d.Get(key+".condition"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".condition"))) {
		request.Condition = expandRequestNetworkAccessLocalExceptionRulesCreateNetworkAccessLocalExceptionRuleRuleCondition(ctx, key+".condition.0", d)
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

func expandRequestNetworkAccessLocalExceptionRulesCreateNetworkAccessLocalExceptionRuleRuleCondition(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRuleRuleCondition {
	request := isegosdk.RequestNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRuleRuleCondition{}
	if v, ok := d.GetOkExists(key + ".condition_type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".condition_type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".condition_type"))) {
		request.ConditionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".is_negate"); !isEmptyValue(reflect.ValueOf(d.Get(key+".is_negate"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".is_negate"))) {
		request.IsNegate = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".link"); !isEmptyValue(reflect.ValueOf(d.Get(key+".link"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".link"))) {
		request.Link = expandRequestNetworkAccessLocalExceptionRulesCreateNetworkAccessLocalExceptionRuleRuleConditionLink(ctx, key+".link.0", d)
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
		request.Children = expandRequestNetworkAccessLocalExceptionRulesCreateNetworkAccessLocalExceptionRuleRuleConditionChildrenArray(ctx, key, d)
	}
	if v, ok := d.GetOkExists(key + ".dates_range"); !isEmptyValue(reflect.ValueOf(d.Get(key+".dates_range"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".dates_range"))) {
		request.DatesRange = expandRequestNetworkAccessLocalExceptionRulesCreateNetworkAccessLocalExceptionRuleRuleConditionDatesRange(ctx, key+".dates_range.0", d)
	}
	if v, ok := d.GetOkExists(key + ".dates_range_exception"); !isEmptyValue(reflect.ValueOf(d.Get(key+".dates_range_exception"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".dates_range_exception"))) {
		request.DatesRangeException = expandRequestNetworkAccessLocalExceptionRulesCreateNetworkAccessLocalExceptionRuleRuleConditionDatesRangeException(ctx, key+".dates_range_exception.0", d)
	}
	if v, ok := d.GetOkExists(key + ".hours_range"); !isEmptyValue(reflect.ValueOf(d.Get(key+".hours_range"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".hours_range"))) {
		request.HoursRange = expandRequestNetworkAccessLocalExceptionRulesCreateNetworkAccessLocalExceptionRuleRuleConditionHoursRange(ctx, key+".hours_range.0", d)
	}
	if v, ok := d.GetOkExists(key + ".hours_range_exception"); !isEmptyValue(reflect.ValueOf(d.Get(key+".hours_range_exception"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".hours_range_exception"))) {
		request.HoursRangeException = expandRequestNetworkAccessLocalExceptionRulesCreateNetworkAccessLocalExceptionRuleRuleConditionHoursRangeException(ctx, key+".hours_range_exception.0", d)
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

func expandRequestNetworkAccessLocalExceptionRulesCreateNetworkAccessLocalExceptionRuleRuleConditionLink(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRuleRuleConditionLink {
	request := isegosdk.RequestNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRuleRuleConditionLink{}
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

func expandRequestNetworkAccessLocalExceptionRulesCreateNetworkAccessLocalExceptionRuleRuleConditionChildrenArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRuleRuleConditionChildren {
	request := []isegosdk.RequestNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRuleRuleConditionChildren{}
	o := d.Get(key)
	if o != nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestNetworkAccessLocalExceptionRulesCreateNetworkAccessLocalExceptionRuleRuleConditionChildren(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		request = append(request, *i)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkAccessLocalExceptionRulesCreateNetworkAccessLocalExceptionRuleRuleConditionChildren(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRuleRuleConditionChildren {
	request := isegosdk.RequestNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRuleRuleConditionChildren{}
	if v, ok := d.GetOkExists(key + ".condition_type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".condition_type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".condition_type"))) {
		request.ConditionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".is_negate"); !isEmptyValue(reflect.ValueOf(d.Get(key+".is_negate"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".is_negate"))) {
		request.IsNegate = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".link"); !isEmptyValue(reflect.ValueOf(d.Get(key+".link"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".link"))) {
		request.Link = expandRequestNetworkAccessLocalExceptionRulesCreateNetworkAccessLocalExceptionRuleRuleConditionChildrenLink(ctx, key+".link.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkAccessLocalExceptionRulesCreateNetworkAccessLocalExceptionRuleRuleConditionChildrenLink(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRuleRuleConditionChildrenLink {
	request := isegosdk.RequestNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRuleRuleConditionChildrenLink{}
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

func expandRequestNetworkAccessLocalExceptionRulesCreateNetworkAccessLocalExceptionRuleRuleConditionDatesRange(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRuleRuleConditionDatesRange {
	request := isegosdk.RequestNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRuleRuleConditionDatesRange{}
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

func expandRequestNetworkAccessLocalExceptionRulesCreateNetworkAccessLocalExceptionRuleRuleConditionDatesRangeException(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRuleRuleConditionDatesRangeException {
	request := isegosdk.RequestNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRuleRuleConditionDatesRangeException{}
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

func expandRequestNetworkAccessLocalExceptionRulesCreateNetworkAccessLocalExceptionRuleRuleConditionHoursRange(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRuleRuleConditionHoursRange {
	request := isegosdk.RequestNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRuleRuleConditionHoursRange{}
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

func expandRequestNetworkAccessLocalExceptionRulesCreateNetworkAccessLocalExceptionRuleRuleConditionHoursRangeException(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRuleRuleConditionHoursRangeException {
	request := isegosdk.RequestNetworkAccessAuthorizationExceptionRulesCreateNetworkAccessLocalExceptionRuleRuleConditionHoursRangeException{}
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

func expandRequestNetworkAccessLocalExceptionRulesUpdateNetworkAccessLocalExceptionRuleByID(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByID {
	request := isegosdk.RequestNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByID{}
	if v, ok := d.GetOkExists(key + ".link"); !isEmptyValue(reflect.ValueOf(d.Get(key+".link"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".link"))) {
		request.Link = expandRequestNetworkAccessLocalExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDLink(ctx, key+".link.0", d)
	}
	if v, ok := d.GetOkExists(key + ".profile"); !isEmptyValue(reflect.ValueOf(d.Get(key+".profile"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".profile"))) {
		request.Profile = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(key + ".rule"); !isEmptyValue(reflect.ValueOf(d.Get(key+".rule"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".rule"))) {
		request.Rule = expandRequestNetworkAccessLocalExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDRule(ctx, key+".rule.0", d)
	}
	if v, ok := d.GetOkExists(key + ".security_group"); !isEmptyValue(reflect.ValueOf(d.Get(key+".security_group"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".security_group"))) {
		request.SecurityGroup = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkAccessLocalExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDLink(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDLink {
	request := isegosdk.RequestNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDLink{}
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

func expandRequestNetworkAccessLocalExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDRule(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDRule {
	request := isegosdk.RequestNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDRule{}
	if v, ok := d.GetOkExists(key + ".condition"); !isEmptyValue(reflect.ValueOf(d.Get(key+".condition"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".condition"))) {
		request.Condition = expandRequestNetworkAccessLocalExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDRuleCondition(ctx, key+".condition.0", d)
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

func expandRequestNetworkAccessLocalExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDRuleCondition(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDRuleCondition {
	request := isegosdk.RequestNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDRuleCondition{}
	if v, ok := d.GetOkExists(key + ".condition_type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".condition_type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".condition_type"))) {
		request.ConditionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".is_negate"); !isEmptyValue(reflect.ValueOf(d.Get(key+".is_negate"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".is_negate"))) {
		request.IsNegate = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".link"); !isEmptyValue(reflect.ValueOf(d.Get(key+".link"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".link"))) {
		request.Link = expandRequestNetworkAccessLocalExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDRuleConditionLink(ctx, key+".link.0", d)
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
		request.Children = expandRequestNetworkAccessLocalExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDRuleConditionChildrenArray(ctx, key, d)
	}
	if v, ok := d.GetOkExists(key + ".dates_range"); !isEmptyValue(reflect.ValueOf(d.Get(key+".dates_range"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".dates_range"))) {
		request.DatesRange = expandRequestNetworkAccessLocalExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDRuleConditionDatesRange(ctx, key+".dates_range.0", d)
	}
	if v, ok := d.GetOkExists(key + ".dates_range_exception"); !isEmptyValue(reflect.ValueOf(d.Get(key+".dates_range_exception"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".dates_range_exception"))) {
		request.DatesRangeException = expandRequestNetworkAccessLocalExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDRuleConditionDatesRangeException(ctx, key+".dates_range_exception.0", d)
	}
	if v, ok := d.GetOkExists(key + ".hours_range"); !isEmptyValue(reflect.ValueOf(d.Get(key+".hours_range"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".hours_range"))) {
		request.HoursRange = expandRequestNetworkAccessLocalExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDRuleConditionHoursRange(ctx, key+".hours_range.0", d)
	}
	if v, ok := d.GetOkExists(key + ".hours_range_exception"); !isEmptyValue(reflect.ValueOf(d.Get(key+".hours_range_exception"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".hours_range_exception"))) {
		request.HoursRangeException = expandRequestNetworkAccessLocalExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDRuleConditionHoursRangeException(ctx, key+".hours_range_exception.0", d)
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

func expandRequestNetworkAccessLocalExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDRuleConditionLink(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDRuleConditionLink {
	request := isegosdk.RequestNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDRuleConditionLink{}
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

func expandRequestNetworkAccessLocalExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDRuleConditionChildrenArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDRuleConditionChildren {
	request := []isegosdk.RequestNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDRuleConditionChildren{}
	o := d.Get(key)
	if o != nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestNetworkAccessLocalExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDRuleConditionChildren(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		request = append(request, *i)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkAccessLocalExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDRuleConditionChildren(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDRuleConditionChildren {
	request := isegosdk.RequestNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDRuleConditionChildren{}
	if v, ok := d.GetOkExists(key + ".condition_type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".condition_type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".condition_type"))) {
		request.ConditionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".is_negate"); !isEmptyValue(reflect.ValueOf(d.Get(key+".is_negate"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".is_negate"))) {
		request.IsNegate = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".link"); !isEmptyValue(reflect.ValueOf(d.Get(key+".link"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".link"))) {
		request.Link = expandRequestNetworkAccessLocalExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDRuleConditionChildrenLink(ctx, key+".link.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkAccessLocalExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDRuleConditionChildrenLink(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDRuleConditionChildrenLink {
	request := isegosdk.RequestNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDRuleConditionChildrenLink{}
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

func expandRequestNetworkAccessLocalExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDRuleConditionDatesRange(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDRuleConditionDatesRange {
	request := isegosdk.RequestNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDRuleConditionDatesRange{}
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

func expandRequestNetworkAccessLocalExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDRuleConditionDatesRangeException(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDRuleConditionDatesRangeException {
	request := isegosdk.RequestNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDRuleConditionDatesRangeException{}
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

func expandRequestNetworkAccessLocalExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDRuleConditionHoursRange(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDRuleConditionHoursRange {
	request := isegosdk.RequestNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDRuleConditionHoursRange{}
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

func expandRequestNetworkAccessLocalExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDRuleConditionHoursRangeException(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDRuleConditionHoursRangeException {
	request := isegosdk.RequestNetworkAccessAuthorizationExceptionRulesUpdateNetworkAccessLocalExceptionRuleByIDRuleConditionHoursRangeException{}
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

func getAllItemsNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRules(m interface{}, response *isegosdk.ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRules, policyTypeID string) []isegosdk.ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRulesResponse {
	var respItems []isegosdk.ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRulesResponse
	for response.Response != nil && len(*response.Response) > 0 {
		respItems = append(respItems, *response.Response...)
	}
	return respItems
}

func searchNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRules(m interface{}, items []isegosdk.ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRulesResponse, name string, id string, policyID string) (*isegosdk.ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRuleByIDResponse, error) {
	client := m.(*isegosdk.Client)
	var err error
	var foundItem *isegosdk.ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRuleByIDResponse
	for _, item := range items {
		if id != "" && item.Rule != nil && item.Rule.ID == id {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRuleByID
			getItem, _, err = client.NetworkAccessAuthorizationExceptionRules.GetNetworkAccessLocalExceptionRuleByID(policyID, id)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetNetworkAccessLocalExceptionRuleByID")
			}
			foundItem = getItem.Response
			return foundItem, err
		} else if name != "" && item.Rule != nil && item.Rule.Name == name {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseNetworkAccessAuthorizationExceptionRulesGetNetworkAccessLocalExceptionRuleByID
			getItem, _, err = client.NetworkAccessAuthorizationExceptionRules.GetNetworkAccessLocalExceptionRuleByID(policyID, item.Rule.ID)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetNetworkAccessLocalExceptionRuleByID")
			}
			foundItem = getItem.Response
			return foundItem, err
		}
	}
	return foundItem, err
}
