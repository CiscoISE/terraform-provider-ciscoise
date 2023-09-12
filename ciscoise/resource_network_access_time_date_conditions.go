package ciscoise

import (
	"context"
	"fmt"
	"reflect"

	"log"

	isegosdk "github.com/kuba-mazurkiewicz/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkAccessTimeDateConditions() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on Network Access - Time/Date Conditions.

- Network Access Creates time/date condition

- Network Access Update network condition

- Network Access Delete Time/Date condition.
`,

		CreateContext: resourceNetworkAccessTimeDateConditionsCreate,
		ReadContext:   resourceNetworkAccessTimeDateConditionsRead,
		UpdateContext: resourceNetworkAccessTimeDateConditionsUpdate,
		DeleteContext: resourceNetworkAccessTimeDateConditionsDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Description: `Unix timestamp records the last time that the resource was updated.`,
				Type:        schema.TypeString,
				Computed:    true,
			},
			"item": &schema.Schema{
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
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"children": &schema.Schema{
							Description:      `In case type is andBlock or orBlock addtional conditions will be aggregated under this logical (OR/AND) condition`,
							Type:             schema.TypeList,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"attribute_id": &schema.Schema{
										Description:      `Dictionary attribute id (Optional), used for additional verification`,
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
									"attribute_name": &schema.Schema{
										Description:      `Dictionary attribute name`,
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
									"attribute_value": &schema.Schema{
										Description:      `<ul><li>Attribute value for condition</li> <li>Value type is specified in dictionary object</li> <li>if multiple values allowed is specified in dictionary object</li></ul>`,
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
									"dictionary_name": &schema.Schema{
										Description:      `Dictionary name`,
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
									"dictionary_value": &schema.Schema{
										Description:      `Dictionary value`,
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
									"end_date": &schema.Schema{
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
									"name": &schema.Schema{
										Description:      `Dictionary attribute name`,
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
									"operator": &schema.Schema{
										Description:      `Equality operator`,
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
									"start_date": &schema.Schema{
										Description:      `<p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>`,
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
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
								},
							},
						},
						"condition_type": &schema.Schema{
							Description:      `<ul><li>Inidicates whether the record is the condition itself(data) or a logical(or,and) aggregation</li> <li>Data type enum(reference,single) indicates than "conditonId" OR "ConditionAttrs" fields should contain condition data but not both</li> <li>Logical aggreation(and,or) enum indicates that additional conditions are present under the children field</li></ul>`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"dates_range": &schema.Schema{
							Description:      `<p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>`,
							Type:             schema.TypeList,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"end_date": &schema.Schema{
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
									"start_date": &schema.Schema{
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
								},
							},
						},
						"dates_range_exception": &schema.Schema{
							Description:      `<p>Defines for which date/s TimeAndDate condition will be matched or NOT matched if used in exceptionDates prooperty<br> Options are - Date range, for specific date, the same date should be used for start/end date <br> Default - no specific dates<br> In order to reset the dates to have no specific dates Date format - yyyy-mm-dd (MM = month, dd = day, yyyy = year)</p>`,
							Type:             schema.TypeList,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"end_date": &schema.Schema{
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
									"start_date": &schema.Schema{
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
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
							Computed:         true,
						},
						"hours_range": &schema.Schema{
							Description:      `<p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>`,
							Type:             schema.TypeList,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"end_time": &schema.Schema{
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
									"start_time": &schema.Schema{
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
								},
							},
						},
						"hours_range_exception": &schema.Schema{
							Description:      `<p>Defines for which hours a TimeAndDate condition will be matched or not matched if used in exceptionHours property<br> Time foramt - hh:mm  ( h = hour , mm = minutes ) <br> Default - All Day </p>`,
							Type:             schema.TypeList,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"end_time": &schema.Schema{
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
									"start_time": &schema.Schema{
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
									},
								},
							},
						},
						"id": &schema.Schema{
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
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
							Computed:         true,
						},
						"week_days": &schema.Schema{
							Description:      `<p>Defines for which days this condition will be matched<br> Days format - Arrays of WeekDay enums <br> Default - List of All week days</p>`,
							Type:             schema.TypeList,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
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
							Computed:         true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

func resourceNetworkAccessTimeDateConditionsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning NetworkAccessTimeDateConditions create")
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	isEnableAutoImport := clientConfig.EnableAutoImport
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestNetworkAccessTimeDateConditionsCreateNetworkAccessTimeCondition(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vName, _ := resourceItem["name"]
	vvName := interfaceToString(vName)
	if isEnableAutoImport {
		if okID && vvID != "" {
			getResponse2, _, err := client.NetworkAccessTimeDateConditions.GetNetworkAccessTimeConditionByID(vvID)
			if err == nil && getResponse2 != nil {
				resourceMap := make(map[string]string)
				resourceMap["id"] = vvID
				resourceMap["name"] = vvName
				d.SetId(joinResourceID(resourceMap))
				return resourceNetworkAccessTimeDateConditionsRead(ctx, d, m)
			}
		} else {
			response2, _, err := client.NetworkAccessTimeDateConditions.GetNetworkAccessTimeConditions()
			if response2 != nil && err == nil {
				items2 := getAllItemsNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditions(m, response2)
				item2, err := searchNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditions(m, items2, vvName, vvID)
				if err == nil && item2 != nil {
					resourceMap := make(map[string]string)
					resourceMap["id"] = item2.ID
					resourceMap["name"] = vvName
					d.SetId(joinResourceID(resourceMap))
					return resourceNetworkAccessTimeDateConditionsRead(ctx, d, m)
				}
			}
		}
	}
	resp1, restyResp1, err := client.NetworkAccessTimeDateConditions.CreateNetworkAccessTimeCondition(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateNetworkAccessTimeCondition", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateNetworkAccessTimeCondition", err))
		return diags
	}
	if vvID != resp1.Response.ID {
		vvID = resp1.Response.ID
	}
	if vvName != resp1.Response.Name {
		vvName = resp1.Response.Name
	}
	resourceMap := make(map[string]string)
	resourceMap["id"] = vvID
	resourceMap["name"] = vvName
	d.SetId(joinResourceID(resourceMap))
	return resourceNetworkAccessTimeDateConditionsRead(ctx, d, m)
}

func resourceNetworkAccessTimeDateConditionsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning NetworkAccessTimeDateConditions read for id=[%s]", d.Id())
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID, okID := resourceMap["id"]
	vName, okName := resourceMap["name"]
	vvID := vID
	vvName := vName
	method1 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetNetworkAccessTimeConditions")

		response1, restyResp1, err := client.NetworkAccessTimeDateConditions.GetNetworkAccessTimeConditions()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		items1 := getAllItemsNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditions(m, response1)
		item1, err := searchNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditions(m, items1, vvName, vvID)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}
		vItem1 := flattenNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditionByIDItem(item1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNetworkAccessTimeConditions search response",
				err))
			return diags
		}
		if err := d.Set("parameters", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNetworkAccessTimeConditions search response",
				err))
			return diags
		}

	}
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkAccessTimeConditionByID")

		response2, restyResp2, err := client.NetworkAccessTimeDateConditions.GetNetworkAccessTimeConditionByID(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditionByIDItem(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNetworkAccessTimeConditionByID response",
				err))
			return diags
		}
		if err := d.Set("parameters", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNetworkAccessTimeConditionByID response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceNetworkAccessTimeDateConditionsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning NetworkAccessTimeDateConditions update for id=[%s]", d.Id())
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID, okID := resourceMap["id"]
	vName, okName := resourceMap["name"]
	vvID := vID
	vvName := vName
	method1 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	// NOTE: Added getAllItems and search function to get missing params
	if selectedMethod == 2 {
		getResp1, _, err := client.NetworkAccessTimeDateConditions.GetNetworkAccessTimeConditions()
		if err == nil && getResp1 == nil {
			items1 := getAllItemsNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditions(m, getResp1)
			item1, err := searchNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditions(m, items1, vvName, vvID)
			if err == nil && item1 == nil {
				if vID != item1.ID {
					vvID = item1.ID
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
		request1 := expandRequestNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByID(ctx, "parameters.0", d)
		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
		response1, restyResp1, err := client.NetworkAccessTimeDateConditions.UpdateNetworkAccessTimeConditionByID(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateNetworkAccessTimeConditionByID", err, restyResp1.String(),
					"Failure at UpdateNetworkAccessTimeConditionByID, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateNetworkAccessTimeConditionByID", err,
				"Failure at UpdateNetworkAccessTimeConditionByID, unexpected response", ""))
			return diags
		}
		_ = d.Set("last_updated", getUnixTimeString())
	}

	return resourceNetworkAccessTimeDateConditionsRead(ctx, d, m)
}

func resourceNetworkAccessTimeDateConditionsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning NetworkAccessTimeDateConditions delete for id=[%s]", d.Id())
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID, okID := resourceMap["id"]
	vName, okName := resourceMap["name"]
	vvID := vID
	vvName := vName
	method1 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	// REVIEW: Add getAllItems and search function to get missing params
	if selectedMethod == 2 {

		getResp1, _, err := client.NetworkAccessTimeDateConditions.GetNetworkAccessTimeConditions()
		if err != nil || getResp1 == nil {
			// Assume that element it is already gone
			return diags
		}
		items1 := getAllItemsNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditions(m, getResp1)
		item1, err := searchNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditions(m, items1, vvName, vvID)
		if err != nil || item1 == nil {
			// Assume that element it is already gone
			return diags
		}
		if vID != item1.ID {
			vvID = item1.ID
		} else {
			vvID = vID
		}
	}
	if selectedMethod == 1 {
		vvID = vID
		getResp, _, err := client.NetworkAccessTimeDateConditions.GetNetworkAccessTimeConditionByID(vvID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	response1, restyResp1, err := client.NetworkAccessTimeDateConditions.DeleteNetworkAccessTimeConditionByID(vvID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteNetworkAccessTimeConditionByID", err, restyResp1.String(),
				"Failure at DeleteNetworkAccessTimeConditionByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteNetworkAccessTimeConditionByID", err,
			"Failure at DeleteNetworkAccessTimeConditionByID, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestNetworkAccessTimeDateConditionsCreateNetworkAccessTimeCondition(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessTimeDateConditionsCreateNetworkAccessTimeCondition {
	request := isegosdk.RequestNetworkAccessTimeDateConditionsCreateNetworkAccessTimeCondition{}
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
		request.Children = expandRequestNetworkAccessTimeDateConditionsCreateNetworkAccessTimeConditionChildrenArray(ctx, key+".children", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dates_range")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dates_range")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dates_range")))) {
		request.DatesRange = expandRequestNetworkAccessTimeDateConditionsCreateNetworkAccessTimeConditionDatesRange(ctx, key+".dates_range.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dates_range_exception")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dates_range_exception")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dates_range_exception")))) {
		request.DatesRangeException = expandRequestNetworkAccessTimeDateConditionsCreateNetworkAccessTimeConditionDatesRangeException(ctx, key+".dates_range_exception.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".hours_range")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".hours_range")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".hours_range")))) {
		request.HoursRange = expandRequestNetworkAccessTimeDateConditionsCreateNetworkAccessTimeConditionHoursRange(ctx, key+".hours_range.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".hours_range_exception")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".hours_range_exception")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".hours_range_exception")))) {
		request.HoursRangeException = expandRequestNetworkAccessTimeDateConditionsCreateNetworkAccessTimeConditionHoursRangeException(ctx, key+".hours_range_exception.0", d)
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

func expandRequestNetworkAccessTimeDateConditionsCreateNetworkAccessTimeConditionLink(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessTimeDateConditionsCreateNetworkAccessTimeConditionLink {
	request := isegosdk.RequestNetworkAccessTimeDateConditionsCreateNetworkAccessTimeConditionLink{}
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

func expandRequestNetworkAccessTimeDateConditionsCreateNetworkAccessTimeConditionChildrenArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestNetworkAccessTimeDateConditionsCreateNetworkAccessTimeConditionChildren {
	request := []isegosdk.RequestNetworkAccessTimeDateConditionsCreateNetworkAccessTimeConditionChildren{}
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
		i := expandRequestNetworkAccessTimeDateConditionsCreateNetworkAccessTimeConditionChildren(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkAccessTimeDateConditionsCreateNetworkAccessTimeConditionChildren(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessTimeDateConditionsCreateNetworkAccessTimeConditionChildren {
	request := isegosdk.RequestNetworkAccessTimeDateConditionsCreateNetworkAccessTimeConditionChildren{}
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

func expandRequestNetworkAccessTimeDateConditionsCreateNetworkAccessTimeConditionChildrenLink(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessTimeDateConditionsCreateNetworkAccessTimeConditionChildrenLink {
	request := isegosdk.RequestNetworkAccessTimeDateConditionsCreateNetworkAccessTimeConditionChildrenLink{}
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

func expandRequestNetworkAccessTimeDateConditionsCreateNetworkAccessTimeConditionDatesRange(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessTimeDateConditionsCreateNetworkAccessTimeConditionDatesRange {
	request := isegosdk.RequestNetworkAccessTimeDateConditionsCreateNetworkAccessTimeConditionDatesRange{}
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

func expandRequestNetworkAccessTimeDateConditionsCreateNetworkAccessTimeConditionDatesRangeException(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessTimeDateConditionsCreateNetworkAccessTimeConditionDatesRangeException {
	request := isegosdk.RequestNetworkAccessTimeDateConditionsCreateNetworkAccessTimeConditionDatesRangeException{}
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

func expandRequestNetworkAccessTimeDateConditionsCreateNetworkAccessTimeConditionHoursRange(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessTimeDateConditionsCreateNetworkAccessTimeConditionHoursRange {
	request := isegosdk.RequestNetworkAccessTimeDateConditionsCreateNetworkAccessTimeConditionHoursRange{}
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

func expandRequestNetworkAccessTimeDateConditionsCreateNetworkAccessTimeConditionHoursRangeException(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessTimeDateConditionsCreateNetworkAccessTimeConditionHoursRangeException {
	request := isegosdk.RequestNetworkAccessTimeDateConditionsCreateNetworkAccessTimeConditionHoursRangeException{}
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

func expandRequestNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByID(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByID {
	request := isegosdk.RequestNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByID{}
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
		request.Children = expandRequestNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByIDChildrenArray(ctx, key+".children", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dates_range")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dates_range")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dates_range")))) {
		request.DatesRange = expandRequestNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByIDDatesRange(ctx, key+".dates_range.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dates_range_exception")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dates_range_exception")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dates_range_exception")))) {
		request.DatesRangeException = expandRequestNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByIDDatesRangeException(ctx, key+".dates_range_exception.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".hours_range")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".hours_range")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".hours_range")))) {
		request.HoursRange = expandRequestNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByIDHoursRange(ctx, key+".hours_range.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".hours_range_exception")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".hours_range_exception")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".hours_range_exception")))) {
		request.HoursRangeException = expandRequestNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByIDHoursRangeException(ctx, key+".hours_range_exception.0", d)
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

func expandRequestNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByIDLink(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByIDLink {
	request := isegosdk.RequestNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByIDLink{}
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

func expandRequestNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByIDChildrenArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByIDChildren {
	request := []isegosdk.RequestNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByIDChildren{}
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
		i := expandRequestNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByIDChildren(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByIDChildren(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByIDChildren {
	request := isegosdk.RequestNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByIDChildren{}
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

func expandRequestNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByIDChildrenLink(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByIDChildrenLink {
	request := isegosdk.RequestNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByIDChildrenLink{}
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

func expandRequestNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByIDDatesRange(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByIDDatesRange {
	request := isegosdk.RequestNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByIDDatesRange{}
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

func expandRequestNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByIDDatesRangeException(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByIDDatesRangeException {
	request := isegosdk.RequestNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByIDDatesRangeException{}
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

func expandRequestNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByIDHoursRange(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByIDHoursRange {
	request := isegosdk.RequestNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByIDHoursRange{}
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

func expandRequestNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByIDHoursRangeException(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByIDHoursRangeException {
	request := isegosdk.RequestNetworkAccessTimeDateConditionsUpdateNetworkAccessTimeConditionByIDHoursRangeException{}
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

func getAllItemsNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditions(m interface{}, response *isegosdk.ResponseNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditions) []isegosdk.ResponseNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditionsResponse {
	var respItems []isegosdk.ResponseNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditionsResponse
	if response.Response != nil && len(*response.Response) > 0 {
		respItems = append(respItems, *response.Response...)
	}
	return respItems
}

func searchNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditions(m interface{}, items []isegosdk.ResponseNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditionsResponse, name string, id string) (*isegosdk.ResponseNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditionByIDResponse, error) {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client
	var err error
	var foundItem *isegosdk.ResponseNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditionByIDResponse
	for _, item := range items {
		if id != "" && item.ID == id {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditionByID
			getItem, _, err = client.NetworkAccessTimeDateConditions.GetNetworkAccessTimeConditionByID(id)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetNetworkAccessTimeConditionByID")
			}
			foundItem = getItem.Response
			return foundItem, err
		} else if name != "" && item.Name == name {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseNetworkAccessTimeDateConditionsGetNetworkAccessTimeConditionByID
			getItem, _, err = client.NetworkAccessTimeDateConditions.GetNetworkAccessTimeConditionByID(item.ID)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetNetworkAccessTimeConditionByID")
			}
			foundItem = getItem.Response
			return foundItem, err
		}
	}
	return foundItem, err
}
