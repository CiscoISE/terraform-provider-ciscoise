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

func resourceDeviceAdministrationTimeDateConditions() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on Device Administration - Time/Date Conditions.

- Device Admin Creates time/date condition.

- Device Admin Update network condition.

- Device Admin Delete Time/Date condition.
`,

		CreateContext: resourceDeviceAdministrationTimeDateConditionsCreate,
		ReadContext:   resourceDeviceAdministrationTimeDateConditionsRead,
		UpdateContext: resourceDeviceAdministrationTimeDateConditionsUpdate,
		DeleteContext: resourceDeviceAdministrationTimeDateConditionsDelete,
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
				Optional: true,
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
		},
	}
}

func resourceDeviceAdministrationTimeDateConditionsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeCondition(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vName, _ := resourceItem["name"]
	vvName := interfaceToString(vName)
	if okID && vvID != "" {
		getResponse2, _, err := client.DeviceAdministrationTimeDateConditions.GetDeviceAdminTimeConditionByID(vvID)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return resourceDeviceAdministrationTimeDateConditionsRead(ctx, d, m)
		}
	} else {
		response2, _, err := client.DeviceAdministrationTimeDateConditions.GetDeviceAdminTimeConditions()
		if response2 != nil && err == nil {
			items2 := getAllItemsDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditions(m, response2)
			item2, err := searchDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditions(m, items2, vvName, vvID)
			if err == nil && item2 != nil {
				resourceMap := make(map[string]string)
				resourceMap["id"] = vvID
				resourceMap["name"] = vvName
				d.SetId(joinResourceID(resourceMap))
				return resourceDeviceAdministrationTimeDateConditionsRead(ctx, d, m)
			}
		}
	}
	resp1, restyResp1, err := client.DeviceAdministrationTimeDateConditions.CreateDeviceAdminTimeCondition(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateDeviceAdminTimeCondition", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateDeviceAdminTimeCondition", err))
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
	return resourceDeviceAdministrationTimeDateConditionsRead(ctx, d, m)
}

func resourceDeviceAdministrationTimeDateConditionsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

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
		log.Printf("[DEBUG] Selected method: GetDeviceAdminTimeConditions")

		response1, restyResp1, err := client.DeviceAdministrationTimeDateConditions.GetDeviceAdminTimeConditions()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		items1 := getAllItemsDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditions(m, response1)
		item1, err := searchDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditions(m, items1, vvName, vvID)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}
		vItem1 := flattenDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionByIDItem(item1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceAdminTimeConditions search response",
				err))
			return diags
		}

	}
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDeviceAdminTimeConditionByID")

		response2, restyResp2, err := client.DeviceAdministrationTimeDateConditions.GetDeviceAdminTimeConditionByID(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionByIDItem(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceAdminTimeConditionByID response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceDeviceAdministrationTimeDateConditionsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

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
	// NOTE: Consider adding getAllItems and search function to get missing params
	if selectedMethod == 2 {
		getResp1, _, err := client.DeviceAdministrationTimeDateConditions.GetDeviceAdminTimeConditions()
		if err == nil && getResp1 != nil {
			items1 := getAllItemsDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditions(m, getResp1)
			item1, err := searchDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditions(m, items1, vvName, vvID)
			if err == nil && item1 != nil {
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
		request1 := expandRequestDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByID(ctx, "parameters.0", d)
		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
		response1, restyResp1, err := client.DeviceAdministrationTimeDateConditions.UpdateDeviceAdminTimeConditionByID(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateDeviceAdminTimeConditionByID", err, restyResp1.String(),
					"Failure at UpdateDeviceAdminTimeConditionByID, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateDeviceAdminTimeConditionByID", err,
				"Failure at UpdateDeviceAdminTimeConditionByID, unexpected response", ""))
			return diags
		}
	}

	return resourceDeviceAdministrationTimeDateConditionsRead(ctx, d, m)
}

func resourceDeviceAdministrationTimeDateConditionsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

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

		getResp1, _, err := client.DeviceAdministrationTimeDateConditions.GetDeviceAdminTimeConditions()
		if err != nil || getResp1 == nil {
			// Assume that element it is already gone
			return diags
		}
		items1 := getAllItemsDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditions(m, getResp1)
		item1, err := searchDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditions(m, items1, vvName, vvID)
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
		getResp, _, err := client.DeviceAdministrationTimeDateConditions.GetDeviceAdminTimeConditionByID(vvID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	response1, restyResp1, err := client.DeviceAdministrationTimeDateConditions.DeleteDeviceAdminTimeConditionByID(vvID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteDeviceAdminTimeConditionByID", err, restyResp1.String(),
				"Failure at DeleteDeviceAdminTimeConditionByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteDeviceAdminTimeConditionByID", err,
			"Failure at DeleteDeviceAdminTimeConditionByID, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeCondition(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeCondition {
	request := isegosdk.RequestDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeCondition{}
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
		request.Children = expandRequestDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionChildrenArray(ctx, key+".children", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dates_range")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dates_range")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dates_range")))) {
		request.DatesRange = expandRequestDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionDatesRange(ctx, key+".dates_range.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dates_range_exception")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dates_range_exception")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dates_range_exception")))) {
		request.DatesRangeException = expandRequestDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionDatesRangeException(ctx, key+".dates_range_exception.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".hours_range")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".hours_range")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".hours_range")))) {
		request.HoursRange = expandRequestDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionHoursRange(ctx, key+".hours_range.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".hours_range_exception")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".hours_range_exception")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".hours_range_exception")))) {
		request.HoursRangeException = expandRequestDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionHoursRangeException(ctx, key+".hours_range_exception.0", d)
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

func expandRequestDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionLink(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionLink {
	request := isegosdk.RequestDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionLink{}
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

func expandRequestDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionChildrenArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionChildren {
	request := []isegosdk.RequestDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionChildren{}
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
		i := expandRequestDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionChildren(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionChildren(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionChildren {
	request := isegosdk.RequestDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionChildren{}
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

func expandRequestDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionChildrenLink(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionChildrenLink {
	request := isegosdk.RequestDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionChildrenLink{}
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

func expandRequestDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionDatesRange(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionDatesRange {
	request := isegosdk.RequestDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionDatesRange{}
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

func expandRequestDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionDatesRangeException(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionDatesRangeException {
	request := isegosdk.RequestDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionDatesRangeException{}
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

func expandRequestDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionHoursRange(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionHoursRange {
	request := isegosdk.RequestDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionHoursRange{}
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

func expandRequestDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionHoursRangeException(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionHoursRangeException {
	request := isegosdk.RequestDeviceAdministrationTimeDateConditionsCreateDeviceAdminTimeConditionHoursRangeException{}
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

func expandRequestDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByID(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByID {
	request := isegosdk.RequestDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByID{}
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
		request.Children = expandRequestDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDChildrenArray(ctx, key+".children", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dates_range")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dates_range")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dates_range")))) {
		request.DatesRange = expandRequestDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDDatesRange(ctx, key+".dates_range.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dates_range_exception")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dates_range_exception")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dates_range_exception")))) {
		request.DatesRangeException = expandRequestDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDDatesRangeException(ctx, key+".dates_range_exception.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".hours_range")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".hours_range")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".hours_range")))) {
		request.HoursRange = expandRequestDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDHoursRange(ctx, key+".hours_range.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".hours_range_exception")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".hours_range_exception")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".hours_range_exception")))) {
		request.HoursRangeException = expandRequestDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDHoursRangeException(ctx, key+".hours_range_exception.0", d)
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

func expandRequestDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDLink(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDLink {
	request := isegosdk.RequestDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDLink{}
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

func expandRequestDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDChildrenArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDChildren {
	request := []isegosdk.RequestDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDChildren{}
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
		i := expandRequestDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDChildren(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDChildren(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDChildren {
	request := isegosdk.RequestDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDChildren{}
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

func expandRequestDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDChildrenLink(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDChildrenLink {
	request := isegosdk.RequestDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDChildrenLink{}
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

func expandRequestDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDDatesRange(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDDatesRange {
	request := isegosdk.RequestDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDDatesRange{}
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

func expandRequestDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDDatesRangeException(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDDatesRangeException {
	request := isegosdk.RequestDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDDatesRangeException{}
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

func expandRequestDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDHoursRange(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDHoursRange {
	request := isegosdk.RequestDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDHoursRange{}
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

func expandRequestDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDHoursRangeException(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDHoursRangeException {
	request := isegosdk.RequestDeviceAdministrationTimeDateConditionsUpdateDeviceAdminTimeConditionByIDHoursRangeException{}
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

func getAllItemsDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditions(m interface{}, response *isegosdk.ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditions) []isegosdk.ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionsResponse {
	var respItems []isegosdk.ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionsResponse
	if response.Response != nil && len(*response.Response) > 0 {
		respItems = append(respItems, *response.Response...)
	}
	return respItems
}

func searchDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditions(m interface{}, items []isegosdk.ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionsResponse, name string, id string) (*isegosdk.ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionByIDResponse, error) {
	client := m.(*isegosdk.Client)
	var err error
	var foundItem *isegosdk.ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionByIDResponse
	for _, item := range items {
		if id != "" && item.ID == id {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionByID
			getItem, _, err = client.DeviceAdministrationTimeDateConditions.GetDeviceAdminTimeConditionByID(id)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetDeviceAdminTimeConditionByID")
			}
			foundItem = getItem.Response
			return foundItem, err
		} else if name != "" && item.Name == name {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseDeviceAdministrationTimeDateConditionsGetDeviceAdminTimeConditionByID
			getItem, _, err = client.DeviceAdministrationTimeDateConditions.GetDeviceAdminTimeConditionByID(item.ID)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetDeviceAdminTimeConditionByID")
			}
			foundItem = getItem.Response
			return foundItem, err
		}
	}
	return foundItem, err
}
