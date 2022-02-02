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

func resourceDeviceAdministrationConditions() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on Device Administration - Conditions.

- Device Admin Creates a library condition:

 Library Condition has hierarchical structure which define a set of condition for which authentication and authorization
policy rules could be match.

 Condition can be compose from single dictionary attribute name and value using model
LibraryConditionAttributes
 , or from combination of dictionary attributes with logical operation of AND/OR between them, using models:
LibraryConditionAndBlock
 or
LibraryConditionOrBlock
.

 When using AND/OR blocks, the condition will include inner layers inside these blocks, these layers are built using the
inner condition models:
ConditionAttributes
,
ConditionAndBlock
,
ConditionOrBlock
, that represent dynamically built Conditions which are not stored in the conditions Library, or using
ConditionReference
, which includes an ID to existing stored condition in the library.

 The LibraryCondition models can only be used in the outer-most layer (root of the condition) and must always include
the condition name.

 When using one of the 3 inner condition models (
ConditionAttributes, ConditionAndBlock, ConditionOrBlock
), condition name cannot be included in the request, since these will not be stored in the conditions library, and used
only as inner members of the root condition.

 When using
ConditionReference
 model in inner layers, the condition name is not required.

 ConditionReference objects can also include a reference ID to a condition of type
TimeAndDate
.



- Device Admin Update library condition using condition name.

- NDevice Admin Delete a library condition using condition Name.

- Device Admin Update library condition.

- Device Admin Delete a library condition.
`,

		CreateContext: resourceDeviceAdministrationConditionsCreate,
		ReadContext:   resourceDeviceAdministrationConditionsRead,
		UpdateContext: resourceDeviceAdministrationConditionsUpdate,
		DeleteContext: resourceDeviceAdministrationConditionsDelete,
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
				Required: true,
				MaxItems: 1,
				MinItems: 1,
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

func resourceDeviceAdministrationConditionsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning DeviceAdministrationConditions create")
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestDeviceAdministrationConditionsCreateDeviceAdminCondition(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vName, okName := resourceItem["name"]
	vvName := interfaceToString(vName)
	if okID && vvID != "" {
		getResponse1, _, err := client.DeviceAdministrationConditions.GetDeviceAdminConditionByID(vvID)
		if err == nil && getResponse1 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return resourceDeviceAdministrationConditionsRead(ctx, d, m)
		}
	}
	if okName && vvName != "" {
		getResponse2, _, err := client.DeviceAdministrationConditions.GetDeviceAdminConditionByName(vvName)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return resourceDeviceAdministrationConditionsRead(ctx, d, m)
		}
	}
	resp1, restyResp1, err := client.DeviceAdministrationConditions.CreateDeviceAdminCondition(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateDeviceAdminCondition", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateDeviceAdminCondition", err))
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
	return resourceDeviceAdministrationConditionsRead(ctx, d, m)
}

func resourceDeviceAdministrationConditionsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning DeviceAdministrationConditions read for id=[%s]", d.Id())
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vName, okName := resourceMap["name"]
	vID, okID := resourceMap["id"]

	method1 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetDeviceAdminConditionByName")
		vvName := vName

		response1, restyResp1, err := client.DeviceAdministrationConditions.GetDeviceAdminConditionByName(vvName)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItemName1 := flattenDeviceAdministrationConditionsGetDeviceAdminConditionByNameItemName(response1.Response)
		if err := d.Set("item", vItemName1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceAdminConditionByName response",
				err))
			return diags
		}
		return diags

	}
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDeviceAdminConditionByID")
		vvID := vID

		response2, restyResp2, err := client.DeviceAdministrationConditions.GetDeviceAdminConditionByID(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItemID2 := flattenDeviceAdministrationConditionsGetDeviceAdminConditionByIDItemID(response2.Response)
		if err := d.Set("item", vItemID2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceAdminConditionByID response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceDeviceAdministrationConditionsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning DeviceAdministrationConditions update for id=[%s]", d.Id())
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vName, okName := resourceMap["name"]
	vID, okID := resourceMap["id"]

	method1 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	var vvID string
	var vvName string
	if selectedMethod == 1 {
		vvID = vID
	}
	if selectedMethod == 2 {
		vvName = vName
		getResp, _, err := client.DeviceAdministrationConditions.GetDeviceAdminConditionByName(vvName)
		if err != nil || getResp == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDeviceAdminConditionByName", err,
				"Failure at GetDeviceAdminConditionByName, unexpected response", ""))
			return diags
		}
		//Set value vvID = getResp.
		if getResp.Response != nil {
			vvID = getResp.Response.ID
		}
	}
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] ID used for update operation %s", vvID)
		request1 := expandRequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByID(ctx, "parameters.0", d)
		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
		response1, restyResp1, err := client.DeviceAdministrationConditions.UpdateDeviceAdminConditionByID(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateDeviceAdminConditionByID", err, restyResp1.String(),
					"Failure at UpdateDeviceAdminConditionByID, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateDeviceAdminConditionByID", err,
				"Failure at UpdateDeviceAdminConditionByID, unexpected response", ""))
			return diags
		}
	}

	return resourceDeviceAdministrationConditionsRead(ctx, d, m)
}

func resourceDeviceAdministrationConditionsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning DeviceAdministrationConditions delete for id=[%s]", d.Id())
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vName, okName := resourceMap["name"]
	vID, okID := resourceMap["id"]

	method1 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	var vvID string
	var vvName string
	if selectedMethod == 1 {
		vvID = vID
		getResp, _, err := client.DeviceAdministrationConditions.GetDeviceAdminConditionByID(vvID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	if selectedMethod == 2 {
		vvName = vName
		getResp, _, err := client.DeviceAdministrationConditions.GetDeviceAdminConditionByName(vvName)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
		//Set value vvID = getResp.
		if getResp.Response != nil {
			vvID = getResp.Response.ID
		}
	}
	response1, restyResp1, err := client.DeviceAdministrationConditions.DeleteDeviceAdminConditionByID(vvID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteDeviceAdminConditionByID", err, restyResp1.String(),
				"Failure at DeleteDeviceAdminConditionByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteDeviceAdminConditionByID", err,
			"Failure at DeleteDeviceAdminConditionByID, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestDeviceAdministrationConditionsCreateDeviceAdminCondition(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationConditionsCreateDeviceAdminCondition {
	request := isegosdk.RequestDeviceAdministrationConditionsCreateDeviceAdminCondition{}
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
		request.Children = expandRequestDeviceAdministrationConditionsCreateDeviceAdminConditionChildrenArray(ctx, key+".children", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dates_range")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dates_range")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dates_range")))) {
		request.DatesRange = expandRequestDeviceAdministrationConditionsCreateDeviceAdminConditionDatesRange(ctx, key+".dates_range.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dates_range_exception")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dates_range_exception")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dates_range_exception")))) {
		request.DatesRangeException = expandRequestDeviceAdministrationConditionsCreateDeviceAdminConditionDatesRangeException(ctx, key+".dates_range_exception.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".hours_range")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".hours_range")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".hours_range")))) {
		request.HoursRange = expandRequestDeviceAdministrationConditionsCreateDeviceAdminConditionHoursRange(ctx, key+".hours_range.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".hours_range_exception")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".hours_range_exception")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".hours_range_exception")))) {
		request.HoursRangeException = expandRequestDeviceAdministrationConditionsCreateDeviceAdminConditionHoursRangeException(ctx, key+".hours_range_exception.0", d)
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

func expandRequestDeviceAdministrationConditionsCreateDeviceAdminConditionLink(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationConditionsCreateDeviceAdminConditionLink {
	request := isegosdk.RequestDeviceAdministrationConditionsCreateDeviceAdminConditionLink{}
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

func expandRequestDeviceAdministrationConditionsCreateDeviceAdminConditionChildrenArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestDeviceAdministrationConditionsCreateDeviceAdminConditionChildren {
	request := []isegosdk.RequestDeviceAdministrationConditionsCreateDeviceAdminConditionChildren{}
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
		i := expandRequestDeviceAdministrationConditionsCreateDeviceAdminConditionChildren(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceAdministrationConditionsCreateDeviceAdminConditionChildren(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationConditionsCreateDeviceAdminConditionChildren {
	request := isegosdk.RequestDeviceAdministrationConditionsCreateDeviceAdminConditionChildren{}
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

func expandRequestDeviceAdministrationConditionsCreateDeviceAdminConditionChildrenLink(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationConditionsCreateDeviceAdminConditionChildrenLink {
	request := isegosdk.RequestDeviceAdministrationConditionsCreateDeviceAdminConditionChildrenLink{}
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

func expandRequestDeviceAdministrationConditionsCreateDeviceAdminConditionDatesRange(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationConditionsCreateDeviceAdminConditionDatesRange {
	request := isegosdk.RequestDeviceAdministrationConditionsCreateDeviceAdminConditionDatesRange{}
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

func expandRequestDeviceAdministrationConditionsCreateDeviceAdminConditionDatesRangeException(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationConditionsCreateDeviceAdminConditionDatesRangeException {
	request := isegosdk.RequestDeviceAdministrationConditionsCreateDeviceAdminConditionDatesRangeException{}
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

func expandRequestDeviceAdministrationConditionsCreateDeviceAdminConditionHoursRange(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationConditionsCreateDeviceAdminConditionHoursRange {
	request := isegosdk.RequestDeviceAdministrationConditionsCreateDeviceAdminConditionHoursRange{}
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

func expandRequestDeviceAdministrationConditionsCreateDeviceAdminConditionHoursRangeException(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationConditionsCreateDeviceAdminConditionHoursRangeException {
	request := isegosdk.RequestDeviceAdministrationConditionsCreateDeviceAdminConditionHoursRangeException{}
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

func expandRequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByID(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByID {
	request := isegosdk.RequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByID{}
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
		request.Children = expandRequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDChildrenArray(ctx, key+".children", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dates_range")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dates_range")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dates_range")))) {
		request.DatesRange = expandRequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDDatesRange(ctx, key+".dates_range.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dates_range_exception")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dates_range_exception")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dates_range_exception")))) {
		request.DatesRangeException = expandRequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDDatesRangeException(ctx, key+".dates_range_exception.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".hours_range")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".hours_range")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".hours_range")))) {
		request.HoursRange = expandRequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDHoursRange(ctx, key+".hours_range.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".hours_range_exception")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".hours_range_exception")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".hours_range_exception")))) {
		request.HoursRangeException = expandRequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDHoursRangeException(ctx, key+".hours_range_exception.0", d)
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

func expandRequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDLink(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDLink {
	request := isegosdk.RequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDLink{}
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

func expandRequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDChildrenArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDChildren {
	request := []isegosdk.RequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDChildren{}
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
		i := expandRequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDChildren(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDChildren(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDChildren {
	request := isegosdk.RequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDChildren{}
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

func expandRequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDChildrenLink(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDChildrenLink {
	request := isegosdk.RequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDChildrenLink{}
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

func expandRequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDDatesRange(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDDatesRange {
	request := isegosdk.RequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDDatesRange{}
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

func expandRequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDDatesRangeException(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDDatesRangeException {
	request := isegosdk.RequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDDatesRangeException{}
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

func expandRequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDHoursRange(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDHoursRange {
	request := isegosdk.RequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDHoursRange{}
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

func expandRequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDHoursRangeException(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDHoursRangeException {
	request := isegosdk.RequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDHoursRangeException{}
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
