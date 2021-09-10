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

func resourceDeviceAdministrationConditions() *schema.Resource {
	return &schema.Resource{

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
				Optional: true,
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
						"href": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"operator": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"rel": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
		},
	}
}

func resourceDeviceAdministrationConditionsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("item"))
	request1 := expandRequestDeviceAdministrationConditionsCreateDeviceAdminCondition(ctx, "item.0", d)
	log.Printf("[DEBUG] request1 => %v", responseInterfaceToString(*request1))

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
			return diags
		}
	}
	if okName && vvName != "" {
		getResponse2, _, err := client.DeviceAdministrationConditions.GetDeviceAdminConditionByName(vvName)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return diags
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
	return diags
}

func resourceDeviceAdministrationConditionsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vName, okName := resourceMap["name"]
	vID, okID := resourceMap["id"]

	method1 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetDeviceAdminConditionByName")
		vvName := vName

		response1, _, err := client.DeviceAdministrationConditions.GetDeviceAdminConditionByName(vvName)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDeviceAdminConditionByName", err,
				"Failure at GetDeviceAdminConditionByName, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

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

		response2, _, err := client.DeviceAdministrationConditions.GetDeviceAdminConditionByID(vvID)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDeviceAdminConditionByID", err,
				"Failure at GetDeviceAdminConditionByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

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
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vName, okName := resourceMap["name"]
	vID, okID := resourceMap["id"]

	method1 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

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
	if d.HasChange("item") {
		log.Printf("[DEBUG] vvID %s", vvID)
		request1 := expandRequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByID(ctx, "item.0", d)
		log.Printf("[DEBUG] request1 => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.DeviceAdministrationConditions.UpdateDeviceAdminConditionByID(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] restyResp1 => %v", restyResp1.String())
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
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vName, okName := resourceMap["name"]
	vID, okID := resourceMap["id"]

	method1 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

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
			log.Printf("[DEBUG] restyResp1 => %v", restyResp1.String())
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
	if v, ok := d.GetOkExists(key + ".condition_type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".condition_type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".condition_type"))) {
		request.ConditionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".is_negate"); !isEmptyValue(reflect.ValueOf(d.Get(key+".is_negate"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".is_negate"))) {
		request.IsNegate = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".link"); !isEmptyValue(reflect.ValueOf(d.Get(key+".link"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".link"))) {
		request.Link = expandRequestDeviceAdministrationConditionsCreateDeviceAdminConditionLink(ctx, key+".link.0", d)
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
		request.Children = expandRequestDeviceAdministrationConditionsCreateDeviceAdminConditionChildrenArray(ctx, key, d)
	}
	if v, ok := d.GetOkExists(key + ".dates_range"); !isEmptyValue(reflect.ValueOf(d.Get(key+".dates_range"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".dates_range"))) {
		request.DatesRange = expandRequestDeviceAdministrationConditionsCreateDeviceAdminConditionDatesRange(ctx, key+".dates_range.0", d)
	}
	if v, ok := d.GetOkExists(key + ".dates_range_exception"); !isEmptyValue(reflect.ValueOf(d.Get(key+".dates_range_exception"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".dates_range_exception"))) {
		request.DatesRangeException = expandRequestDeviceAdministrationConditionsCreateDeviceAdminConditionDatesRangeException(ctx, key+".dates_range_exception.0", d)
	}
	if v, ok := d.GetOkExists(key + ".hours_range"); !isEmptyValue(reflect.ValueOf(d.Get(key+".hours_range"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".hours_range"))) {
		request.HoursRange = expandRequestDeviceAdministrationConditionsCreateDeviceAdminConditionHoursRange(ctx, key+".hours_range.0", d)
	}
	if v, ok := d.GetOkExists(key + ".hours_range_exception"); !isEmptyValue(reflect.ValueOf(d.Get(key+".hours_range_exception"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".hours_range_exception"))) {
		request.HoursRangeException = expandRequestDeviceAdministrationConditionsCreateDeviceAdminConditionHoursRangeException(ctx, key+".hours_range_exception.0", d)
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

func expandRequestDeviceAdministrationConditionsCreateDeviceAdminConditionLink(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationConditionsCreateDeviceAdminConditionLink {
	request := isegosdk.RequestDeviceAdministrationConditionsCreateDeviceAdminConditionLink{}
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

func expandRequestDeviceAdministrationConditionsCreateDeviceAdminConditionChildrenArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestDeviceAdministrationConditionsCreateDeviceAdminConditionChildren {
	request := []isegosdk.RequestDeviceAdministrationConditionsCreateDeviceAdminConditionChildren{}
	o := d.Get(key)
	if o != nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestDeviceAdministrationConditionsCreateDeviceAdminConditionChildren(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		request = append(request, *i)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceAdministrationConditionsCreateDeviceAdminConditionChildren(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationConditionsCreateDeviceAdminConditionChildren {
	request := isegosdk.RequestDeviceAdministrationConditionsCreateDeviceAdminConditionChildren{}
	if v, ok := d.GetOkExists(key + ".condition_type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".condition_type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".condition_type"))) {
		request.ConditionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".is_negate"); !isEmptyValue(reflect.ValueOf(d.Get(key+".is_negate"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".is_negate"))) {
		request.IsNegate = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".link"); !isEmptyValue(reflect.ValueOf(d.Get(key+".link"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".link"))) {
		request.Link = expandRequestDeviceAdministrationConditionsCreateDeviceAdminConditionChildrenLink(ctx, key+".link.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceAdministrationConditionsCreateDeviceAdminConditionChildrenLink(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationConditionsCreateDeviceAdminConditionChildrenLink {
	request := isegosdk.RequestDeviceAdministrationConditionsCreateDeviceAdminConditionChildrenLink{}
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

func expandRequestDeviceAdministrationConditionsCreateDeviceAdminConditionDatesRange(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationConditionsCreateDeviceAdminConditionDatesRange {
	request := isegosdk.RequestDeviceAdministrationConditionsCreateDeviceAdminConditionDatesRange{}
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

func expandRequestDeviceAdministrationConditionsCreateDeviceAdminConditionDatesRangeException(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationConditionsCreateDeviceAdminConditionDatesRangeException {
	request := isegosdk.RequestDeviceAdministrationConditionsCreateDeviceAdminConditionDatesRangeException{}
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

func expandRequestDeviceAdministrationConditionsCreateDeviceAdminConditionHoursRange(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationConditionsCreateDeviceAdminConditionHoursRange {
	request := isegosdk.RequestDeviceAdministrationConditionsCreateDeviceAdminConditionHoursRange{}
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

func expandRequestDeviceAdministrationConditionsCreateDeviceAdminConditionHoursRangeException(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationConditionsCreateDeviceAdminConditionHoursRangeException {
	request := isegosdk.RequestDeviceAdministrationConditionsCreateDeviceAdminConditionHoursRangeException{}
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

func expandRequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByID(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByID {
	request := isegosdk.RequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByID{}
	if v, ok := d.GetOkExists(key + ".condition_type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".condition_type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".condition_type"))) {
		request.ConditionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".is_negate"); !isEmptyValue(reflect.ValueOf(d.Get(key+".is_negate"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".is_negate"))) {
		request.IsNegate = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".link"); !isEmptyValue(reflect.ValueOf(d.Get(key+".link"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".link"))) {
		request.Link = expandRequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDLink(ctx, key+".link.0", d)
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
		request.Children = expandRequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDChildrenArray(ctx, key, d)
	}
	if v, ok := d.GetOkExists(key + ".dates_range"); !isEmptyValue(reflect.ValueOf(d.Get(key+".dates_range"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".dates_range"))) {
		request.DatesRange = expandRequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDDatesRange(ctx, key+".dates_range.0", d)
	}
	if v, ok := d.GetOkExists(key + ".dates_range_exception"); !isEmptyValue(reflect.ValueOf(d.Get(key+".dates_range_exception"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".dates_range_exception"))) {
		request.DatesRangeException = expandRequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDDatesRangeException(ctx, key+".dates_range_exception.0", d)
	}
	if v, ok := d.GetOkExists(key + ".hours_range"); !isEmptyValue(reflect.ValueOf(d.Get(key+".hours_range"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".hours_range"))) {
		request.HoursRange = expandRequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDHoursRange(ctx, key+".hours_range.0", d)
	}
	if v, ok := d.GetOkExists(key + ".hours_range_exception"); !isEmptyValue(reflect.ValueOf(d.Get(key+".hours_range_exception"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".hours_range_exception"))) {
		request.HoursRangeException = expandRequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDHoursRangeException(ctx, key+".hours_range_exception.0", d)
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

func expandRequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDLink(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDLink {
	request := isegosdk.RequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDLink{}
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

func expandRequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDChildrenArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDChildren {
	request := []isegosdk.RequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDChildren{}
	o := d.Get(key)
	if o != nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDChildren(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		request = append(request, *i)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDChildren(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDChildren {
	request := isegosdk.RequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDChildren{}
	if v, ok := d.GetOkExists(key + ".condition_type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".condition_type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".condition_type"))) {
		request.ConditionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".is_negate"); !isEmptyValue(reflect.ValueOf(d.Get(key+".is_negate"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".is_negate"))) {
		request.IsNegate = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".link"); !isEmptyValue(reflect.ValueOf(d.Get(key+".link"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".link"))) {
		request.Link = expandRequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDChildrenLink(ctx, key+".link.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDChildrenLink(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDChildrenLink {
	request := isegosdk.RequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDChildrenLink{}
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

func expandRequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDDatesRange(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDDatesRange {
	request := isegosdk.RequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDDatesRange{}
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

func expandRequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDDatesRangeException(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDDatesRangeException {
	request := isegosdk.RequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDDatesRangeException{}
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

func expandRequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDHoursRange(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDHoursRange {
	request := isegosdk.RequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDHoursRange{}
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

func expandRequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDHoursRangeException(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDHoursRangeException {
	request := isegosdk.RequestDeviceAdministrationConditionsUpdateDeviceAdminConditionByIDHoursRangeException{}
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
