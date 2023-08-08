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

// resourceAction
func resourceDeviceAdministrationAuthorizationRulesUpdateUpdate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on Device Administration - Authorization Rules.

- Device Admin Update authorization rule.
`,

		CreateContext: resourceDeviceAdministrationAuthorizationRulesUpdateUpdateCreate,
		ReadContext:   resourceDeviceAdministrationAuthorizationRulesUpdateUpdateRead,
		DeleteContext: resourceDeviceAdministrationAuthorizationRulesUpdateUpdateDelete,
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
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"commands": &schema.Schema{
							Description:      `Command sets enforce the specified list of commands that can be executed by a device administrator`,
							Type:             schema.TypeList,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							ForceNew:         true,
							Computed:         true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"id": &schema.Schema{
							Description:      `id path parameter. Rule id`,
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: diffSupressOptional(),
							ForceNew:         true,
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
							Description:      `policyId path parameter. Policy id`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							ForceNew:         true,
							Computed:         true,
						},
						"profile": &schema.Schema{
							Description:      `Device admin profiles control the initial login session of the device administrator`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							ForceNew:         true,
							Computed:         true,
						},
						"rule": &schema.Schema{
							Description:      `Common attributes in rule authentication/authorization`,
							Type:             schema.TypeList,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							ForceNew:         true,
							Computed:         true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"condition": &schema.Schema{
										Type:             schema.TypeList,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										ForceNew:         true,
										Computed:         true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"attribute_name": &schema.Schema{
													Description:      `Atribute Name`,
													Type:             schema.TypeString,
													Optional:         true,
													DiffSuppressFunc: diffSupressOptional(),
													ForceNew:         true,
													Computed:         true,
												},
												"attribute_value": &schema.Schema{
													Description:      `Attibute Name`,
													Type:             schema.TypeString,
													Optional:         true,
													DiffSuppressFunc: diffSupressOptional(),
													ForceNew:         true,
													Computed:         true,
												},
												"children": &schema.Schema{
													Description:      `In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition`,
													Type:             schema.TypeList,
													Optional:         true,
													DiffSuppressFunc: diffSupressOptional(),
													ForceNew:         true,
													Computed:         true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"attribute_name": &schema.Schema{
																Description:      `Atribute Name`,
																Type:             schema.TypeString,
																Optional:         true,
																DiffSuppressFunc: diffSupressOptional(),
																ForceNew:         true,
																Computed:         true,
															},
															"attribute_value": &schema.Schema{
																Description:      `Attibute Name`,
																Type:             schema.TypeString,
																Optional:         true,
																DiffSuppressFunc: diffSupressOptional(),
																ForceNew:         true,
																Computed:         true,
															},
															"condition_type": &schema.Schema{
																Description:      `<ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>`,
																Type:             schema.TypeString,
																Optional:         true,
																DiffSuppressFunc: diffSupressOptional(),
																ForceNew:         true,
																Computed:         true,
															},
															"dictionary_name": &schema.Schema{
																Description:      `Dictionary Name`,
																Type:             schema.TypeString,
																Optional:         true,
																DiffSuppressFunc: diffSupressOptional(),
																ForceNew:         true,
																Computed:         true,
															},
															"is_negate": &schema.Schema{
																Description:      `Indicates whereas this condition is in negate mode`,
																Type:             schema.TypeString,
																ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:         true,
																DiffSuppressFunc: diffSupressBool(),
																Computed:         true,
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
															"operator": &schema.Schema{
																Description:      `Operator`,
																Type:             schema.TypeString,
																Optional:         true,
																DiffSuppressFunc: diffSupressOptional(),
																ForceNew:         true,
																Computed:         true,
															},
														},
													},
												},
												"condition_type": &schema.Schema{
													Description:      `<ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>`,
													Type:             schema.TypeString,
													Optional:         true,
													DiffSuppressFunc: diffSupressOptional(),
													ForceNew:         true,
													Computed:         true,
												},
												"dates_range": &schema.Schema{
													Description:      `<p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>`,
													Type:             schema.TypeList,
													Optional:         true,
													DiffSuppressFunc: diffSupressOptional(),
													ForceNew:         true,
													Computed:         true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"end_date": &schema.Schema{
																Type:             schema.TypeString,
																Optional:         true,
																DiffSuppressFunc: diffSupressOptional(),
																ForceNew:         true,
																Computed:         true,
															},
															"start_date": &schema.Schema{
																Type:             schema.TypeString,
																Optional:         true,
																DiffSuppressFunc: diffSupressOptional(),
																ForceNew:         true,
																Computed:         true,
															},
														},
													},
												},
												"dates_range_exception": &schema.Schema{
													Description:      `<p>Defines for which date/s TimeAndDate condition will be matched<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>`,
													Type:             schema.TypeList,
													Optional:         true,
													DiffSuppressFunc: diffSupressOptional(),
													ForceNew:         true,
													Computed:         true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"end_date": &schema.Schema{
																Type:             schema.TypeString,
																Optional:         true,
																DiffSuppressFunc: diffSupressOptional(),
																ForceNew:         true,
																Computed:         true,
															},
															"start_date": &schema.Schema{
																Type:             schema.TypeString,
																Optional:         true,
																DiffSuppressFunc: diffSupressOptional(),
																ForceNew:         true,
																Computed:         true,
															},
														},
													},
												},
												"description": &schema.Schema{
													Description:      `Condition description`,
													Type:             schema.TypeString,
													Optional:         true,
													DiffSuppressFunc: diffSupressOptional(),
													ForceNew:         true,
													Computed:         true,
												},
												"dictionary_name": &schema.Schema{
													Description:      `Dictionary Name`,
													Type:             schema.TypeString,
													Optional:         true,
													DiffSuppressFunc: diffSupressOptional(),
													ForceNew:         true,
													Computed:         true,
												},
												"dictionary_value": &schema.Schema{
													Description:      `Dictionary value`,
													Type:             schema.TypeString,
													Optional:         true,
													DiffSuppressFunc: diffSupressOptional(),
													ForceNew:         true,
													Computed:         true,
												},
												"hours_range": &schema.Schema{
													Description:      `<p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>`,
													Type:             schema.TypeList,
													Optional:         true,
													DiffSuppressFunc: diffSupressOptional(),
													ForceNew:         true,
													Computed:         true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"end_time": &schema.Schema{
																Type:             schema.TypeString,
																Optional:         true,
																DiffSuppressFunc: diffSupressOptional(),
																ForceNew:         true,
																Computed:         true,
															},
															"start_time": &schema.Schema{
																Type:             schema.TypeString,
																Optional:         true,
																DiffSuppressFunc: diffSupressOptional(),
																ForceNew:         true,
																Computed:         true,
															},
														},
													},
												},
												"hours_range_exception": &schema.Schema{
													Description:      `<p>Defines for which hours a TimeAndDate condition will be matched<br> Time format - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>`,
													Type:             schema.TypeList,
													Optional:         true,
													DiffSuppressFunc: diffSupressOptional(),
													ForceNew:         true,
													Computed:         true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"end_time": &schema.Schema{
																Type:             schema.TypeString,
																Optional:         true,
																DiffSuppressFunc: diffSupressOptional(),
																ForceNew:         true,
																Computed:         true,
															},
															"start_time": &schema.Schema{
																Type:             schema.TypeString,
																Optional:         true,
																DiffSuppressFunc: diffSupressOptional(),
																ForceNew:         true,
																Computed:         true,
															},
														},
													},
												},
												"id": &schema.Schema{
													Type:             schema.TypeString,
													Optional:         true,
													DiffSuppressFunc: diffSupressOptional(),
													ForceNew:         true,
													Computed:         true,
												},
												"is_negate": &schema.Schema{
													Description:      `Indicates whereas this condition is in negate mode`,
													Type:             schema.TypeString,
													ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:         true,
													DiffSuppressFunc: diffSupressBool(),
													Computed:         true,
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
													Description:      `Condition name`,
													Type:             schema.TypeString,
													Optional:         true,
													DiffSuppressFunc: diffSupressOptional(),
													ForceNew:         true,
													Computed:         true,
												},
												"operator": &schema.Schema{
													Description:      `Operator`,
													Type:             schema.TypeString,
													Optional:         true,
													DiffSuppressFunc: diffSupressOptional(),
													ForceNew:         true,
													Computed:         true,
												},
												"week_days": &schema.Schema{
													Description:      `<p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>`,
													Type:             schema.TypeList,
													Optional:         true,
													DiffSuppressFunc: diffSupressOptional(),
													ForceNew:         true,
													Computed:         true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"week_days_exception": &schema.Schema{
													Description:      `<p>Defines for which days this condition will NOT be matched<br> Days format - Arrays of WeekDay enums <br> Default - Not enabled</p>`,
													Type:             schema.TypeList,
													Optional:         true,
													DiffSuppressFunc: diffSupressOptional(),
													ForceNew:         true,
													Computed:         true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
									"default": &schema.Schema{
										Description:      `Indicates if this rule is the default one`,
										Type:             schema.TypeString,
										ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:         true,
										DiffSuppressFunc: diffSupressBool(),
										Computed:         true,
									},
									"hit_counts": &schema.Schema{
										Description:      `The amount of times the rule was matched`,
										Type:             schema.TypeInt,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										ForceNew:         true,
										Computed:         true,
									},
									"id": &schema.Schema{
										Description:      `The identifier of the rule`,
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										ForceNew:         true,
										Computed:         true,
									},
									"name": &schema.Schema{
										Description:      `Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]`,
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										ForceNew:         true,
										Computed:         true,
									},
									"rank": &schema.Schema{
										Description:      `The rank(priority) in relation to other rules. Lower rank is higher priority.`,
										Type:             schema.TypeInt,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										ForceNew:         true,
										Computed:         true,
									},
									"state": &schema.Schema{
										Description:      `The state that the rule is in. A disabled rule cannot be matched.`,
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										ForceNew:         true,
										Computed:         true,
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

func resourceDeviceAdministrationAuthorizationRulesUpdateUpdateCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics
	resourceItem := *getResourceItem(d.Get("parameters"))

	vPolicyID := resourceItem["policy_id"]

	vID := resourceItem["id"]

	vvPolicyID := vPolicyID.(string)
	vvID := vID.(string)
	request1 := expandRequestDeviceAdministrationAuthorizationRulesUpdateUpdateDeviceAdminAuthorizationRuleByID(ctx, "parameters.0", d)

	response1, restyResp1, err := client.DeviceAdministrationAuthorizationRules.UpdateDeviceAdminAuthorizationRuleByID(vvPolicyID, vvID, request1)

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			err = fmt.Errorf(restyResp1.String())
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing resourceDeviceAdministrationAuthorizationRulesUpdateUpdateCreate ", err,
			"Failure at resourceDeviceAdministrationAuthorizationRulesUpdateUpdateCreate, unexpected response", ""))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	vItem1 := flattenDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDItem(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting UpdateDeviceAdminAuthorizationRuleByID response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags
}

func resourceDeviceAdministrationAuthorizationRulesUpdateUpdateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*isegosdk.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceDeviceAdministrationAuthorizationRulesUpdateUpdateDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestDeviceAdministrationAuthorizationRulesUpdateUpdateDeviceAdminAuthorizationRuleByID(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByID {
	request := isegosdk.RequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByID{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".commands")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".commands")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".commands")))) {
		request.Commands = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".link")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".link")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".link")))) {
		request.Link = expandRequestDeviceAdministrationAuthorizationRulesUpdateUpdateDeviceAdminAuthorizationRuleByIDLink(ctx, key+".link.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".profile")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".profile")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".profile")))) {
		request.Profile = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rule")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rule")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rule")))) {
		request.Rule = expandRequestDeviceAdministrationAuthorizationRulesUpdateUpdateDeviceAdminAuthorizationRuleByIDRule(ctx, key+".rule.0", d)
	}
	return &request
}

func expandRequestDeviceAdministrationAuthorizationRulesUpdateUpdateDeviceAdminAuthorizationRuleByIDLink(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDLink {
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
	return &request
}

func expandRequestDeviceAdministrationAuthorizationRulesUpdateUpdateDeviceAdminAuthorizationRuleByIDRule(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRule {
	request := isegosdk.RequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRule{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".condition")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".condition")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".condition")))) {
		request.Condition = expandRequestDeviceAdministrationAuthorizationRulesUpdateUpdateDeviceAdminAuthorizationRuleByIDRuleCondition(ctx, key+".condition.0", d)
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
	return &request
}

func expandRequestDeviceAdministrationAuthorizationRulesUpdateUpdateDeviceAdminAuthorizationRuleByIDRuleCondition(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRuleCondition {
	request := isegosdk.RequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRuleCondition{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".condition_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".condition_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".condition_type")))) {
		request.ConditionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_negate")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_negate")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_negate")))) {
		request.IsNegate = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".link")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".link")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".link")))) {
		request.Link = expandRequestDeviceAdministrationAuthorizationRulesUpdateUpdateDeviceAdminAuthorizationRuleByIDRuleConditionLink(ctx, key+".link.0", d)
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
		request.Children = expandRequestDeviceAdministrationAuthorizationRulesUpdateUpdateDeviceAdminAuthorizationRuleByIDRuleConditionChildrenArray(ctx, key+".children", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dates_range")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dates_range")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dates_range")))) {
		request.DatesRange = expandRequestDeviceAdministrationAuthorizationRulesUpdateUpdateDeviceAdminAuthorizationRuleByIDRuleConditionDatesRange(ctx, key+".dates_range.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dates_range_exception")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dates_range_exception")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dates_range_exception")))) {
		request.DatesRangeException = expandRequestDeviceAdministrationAuthorizationRulesUpdateUpdateDeviceAdminAuthorizationRuleByIDRuleConditionDatesRangeException(ctx, key+".dates_range_exception.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".hours_range")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".hours_range")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".hours_range")))) {
		request.HoursRange = expandRequestDeviceAdministrationAuthorizationRulesUpdateUpdateDeviceAdminAuthorizationRuleByIDRuleConditionHoursRange(ctx, key+".hours_range.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".hours_range_exception")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".hours_range_exception")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".hours_range_exception")))) {
		request.HoursRangeException = expandRequestDeviceAdministrationAuthorizationRulesUpdateUpdateDeviceAdminAuthorizationRuleByIDRuleConditionHoursRangeException(ctx, key+".hours_range_exception.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".week_days")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".week_days")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".week_days")))) {
		request.WeekDays = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".week_days_exception")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".week_days_exception")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".week_days_exception")))) {
		request.WeekDaysException = interfaceToSliceString(v)
	}
	return &request
}

func expandRequestDeviceAdministrationAuthorizationRulesUpdateUpdateDeviceAdminAuthorizationRuleByIDRuleConditionLink(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRuleConditionLink {
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
	return &request
}

func expandRequestDeviceAdministrationAuthorizationRulesUpdateUpdateDeviceAdminAuthorizationRuleByIDRuleConditionChildrenArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRuleConditionChildren {
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
	for item_no := range objs {
		i := expandRequestDeviceAdministrationAuthorizationRulesUpdateUpdateDeviceAdminAuthorizationRuleByIDRuleConditionChildren(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestDeviceAdministrationAuthorizationRulesUpdateUpdateDeviceAdminAuthorizationRuleByIDRuleConditionChildren(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRuleConditionChildren {
	request := isegosdk.RequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRuleConditionChildren{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".condition_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".condition_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".condition_type")))) {
		request.ConditionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_negate")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_negate")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_negate")))) {
		request.IsNegate = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".link")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".link")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".link")))) {
		request.Link = expandRequestDeviceAdministrationAuthorizationRulesUpdateUpdateDeviceAdminAuthorizationRuleByIDRuleConditionChildrenLink(ctx, key+".link.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dictionary_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dictionary_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dictionary_name")))) {
		request.DictionaryName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".attribute_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".attribute_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".attribute_name")))) {
		request.AttributeName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".operator")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".operator")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".operator")))) {
		request.Operator = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".attribute_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".attribute_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".attribute_value")))) {
		request.AttributeValue = interfaceToString(v)
	}

	return &request
}

func expandRequestDeviceAdministrationAuthorizationRulesUpdateUpdateDeviceAdminAuthorizationRuleByIDRuleConditionChildrenLink(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRuleConditionChildrenLink {
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
	return &request
}

func expandRequestDeviceAdministrationAuthorizationRulesUpdateUpdateDeviceAdminAuthorizationRuleByIDRuleConditionDatesRange(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRuleConditionDatesRange {
	request := isegosdk.RequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRuleConditionDatesRange{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_date")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_date")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_date")))) {
		request.EndDate = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_date")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_date")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_date")))) {
		request.StartDate = interfaceToString(v)
	}
	return &request
}

func expandRequestDeviceAdministrationAuthorizationRulesUpdateUpdateDeviceAdminAuthorizationRuleByIDRuleConditionDatesRangeException(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRuleConditionDatesRangeException {
	request := isegosdk.RequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRuleConditionDatesRangeException{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_date")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_date")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_date")))) {
		request.EndDate = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_date")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_date")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_date")))) {
		request.StartDate = interfaceToString(v)
	}
	return &request
}

func expandRequestDeviceAdministrationAuthorizationRulesUpdateUpdateDeviceAdminAuthorizationRuleByIDRuleConditionHoursRange(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRuleConditionHoursRange {
	request := isegosdk.RequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRuleConditionHoursRange{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_time")))) {
		request.EndTime = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_time")))) {
		request.StartTime = interfaceToString(v)
	}
	return &request
}

func expandRequestDeviceAdministrationAuthorizationRulesUpdateUpdateDeviceAdminAuthorizationRuleByIDRuleConditionHoursRangeException(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRuleConditionHoursRangeException {
	request := isegosdk.RequestDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDRuleConditionHoursRangeException{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_time")))) {
		request.EndTime = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_time")))) {
		request.StartTime = interfaceToString(v)
	}
	return &request
}

func flattenDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDItem(item *isegosdk.ResponseDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["commands"] = item.Commands
	respItem["link"] = flattenDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDItemLink(item.Link)
	respItem["profile"] = item.Profile
	respItem["rule"] = flattenDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDItemRule(item.Rule)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDItemLink(item *isegosdk.ResponseDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDResponseLink) []map[string]interface{} {
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

func flattenDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDItemRule(item *isegosdk.ResponseDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDResponseRule) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["condition"] = flattenDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDItemRuleCondition(item.Condition)
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

func flattenDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDItemRuleCondition(item *isegosdk.ResponseDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDResponseRuleCondition) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["condition_type"] = item.ConditionType
	respItem["is_negate"] = boolPtrToString(item.IsNegate)
	respItem["link"] = flattenDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDItemRuleConditionLink(item.Link)
	respItem["description"] = item.Description
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["attribute_name"] = item.AttributeName
	respItem["attribute_value"] = item.AttributeValue
	respItem["dictionary_name"] = item.DictionaryName
	respItem["dictionary_value"] = item.DictionaryValue
	respItem["operator"] = item.Operator
	respItem["children"] = flattenDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDItemRuleConditionChildren(item.Children)
	respItem["dates_range"] = flattenDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDItemRuleConditionDatesRange(item.DatesRange)
	respItem["dates_range_exception"] = flattenDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDItemRuleConditionDatesRangeException(item.DatesRangeException)
	respItem["hours_range"] = flattenDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDItemRuleConditionHoursRange(item.HoursRange)
	respItem["hours_range_exception"] = flattenDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDItemRuleConditionHoursRangeException(item.HoursRangeException)
	respItem["week_days"] = item.WeekDays
	respItem["week_days_exception"] = item.WeekDaysException

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDItemRuleConditionLink(item *isegosdk.ResponseDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDResponseRuleConditionLink) []map[string]interface{} {
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

func flattenDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDItemRuleConditionChildren(items *[]isegosdk.ResponseDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDResponseRuleConditionChildren) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["condition_type"] = item.ConditionType
		respItem["is_negate"] = boolPtrToString(item.IsNegate)
		respItem["link"] = flattenDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDItemRuleConditionChildrenLink(item.Link)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDItemRuleConditionChildrenLink(item *isegosdk.ResponseDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDResponseRuleConditionChildrenLink) []map[string]interface{} {
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

func flattenDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDItemRuleConditionDatesRange(item *isegosdk.ResponseDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDResponseRuleConditionDatesRange) []map[string]interface{} {
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

func flattenDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDItemRuleConditionDatesRangeException(item *isegosdk.ResponseDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDResponseRuleConditionDatesRangeException) []map[string]interface{} {
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

func flattenDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDItemRuleConditionHoursRange(item *isegosdk.ResponseDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDResponseRuleConditionHoursRange) []map[string]interface{} {
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

func flattenDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDItemRuleConditionHoursRangeException(item *isegosdk.ResponseDeviceAdministrationAuthorizationRulesUpdateDeviceAdminAuthorizationRuleByIDResponseRuleConditionHoursRangeException) []map[string]interface{} {
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
