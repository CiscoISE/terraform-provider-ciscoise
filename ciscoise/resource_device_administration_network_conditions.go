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

func resourceDeviceAdministrationNetworkConditions() *schema.Resource {
	return &schema.Resource{

		CreateContext: resourceDeviceAdministrationNetworkConditionsCreate,
		ReadContext:   resourceDeviceAdministrationNetworkConditionsRead,
		UpdateContext: resourceDeviceAdministrationNetworkConditionsUpdate,
		DeleteContext: resourceDeviceAdministrationNetworkConditionsDelete,
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

						"condition_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"conditions": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"cli_dnis_list": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"device_group_list": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"device_list": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"ip_addr_list": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"mac_addr_list": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"href": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
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
						"rel": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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

func resourceDeviceAdministrationNetworkConditionsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("item"))
	request1 := expandRequestDeviceAdministrationNetworkConditionsCreateDeviceAdminNetworkCondition(ctx, "item.0", d)
	log.Printf("[DEBUG] request1 => %v", responseInterfaceToString(*request1))

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vName, _ := resourceItem["name"]
	vvName := interfaceToString(vName)
	if okID && vvID != "" {
		getResponse2, _, err := client.DeviceAdministrationNetworkConditions.GetDeviceAdminNetworkConditionByID(vvID)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return diags
		}
	} else {
		response2, _, err := client.DeviceAdministrationNetworkConditions.GetDeviceAdminNetworkConditions()
		if response2 != nil && err == nil {
			items2 := getAllItemsDeviceAdministrationNetworkConditionsGetDeviceAdminNetworkConditions(m, response2)
			item2, err := searchDeviceAdministrationNetworkConditionsGetDeviceAdminNetworkConditions(m, items2, vvName, vvID)
			if err == nil && item2 != nil {
				resourceMap := make(map[string]string)
				resourceMap["id"] = vvID
				resourceMap["name"] = vvName
				d.SetId(joinResourceID(resourceMap))
				return diags
			}
		}
	}
	resp1, restyResp1, err := client.DeviceAdministrationNetworkConditions.CreateDeviceAdminNetworkCondition(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateDeviceAdminNetworkCondition", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateDeviceAdminNetworkCondition", err))
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

func resourceDeviceAdministrationNetworkConditionsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID, okID := resourceMap["id"]
	vName, okName := resourceMap["name"]
	vvID := vID
	vvName := vName
	method1 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetDeviceAdminNetworkConditions")

		response1, _, err := client.DeviceAdministrationNetworkConditions.GetDeviceAdminNetworkConditions()

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDeviceAdminNetworkConditions", err,
				"Failure at GetDeviceAdminNetworkConditions, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		items1 := getAllItemsDeviceAdministrationNetworkConditionsGetDeviceAdminNetworkConditions(m, response1)
		item1, err := searchDeviceAdministrationNetworkConditionsGetDeviceAdminNetworkConditions(m, items1, vvName, vvID)
		if err != nil || item1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when searching item from GetDeviceAdminNetworkConditions response", err,
				"Failure when searching item from GetDeviceAdminNetworkConditions, unexpected response", ""))
			return diags
		}
		if err := d.Set("item", item1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceAdminNetworkConditions search response",
				err))
			return diags
		}

	}
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDeviceAdminNetworkConditionByID")

		response2, _, err := client.DeviceAdministrationNetworkConditions.GetDeviceAdminNetworkConditionByID(vvID)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDeviceAdminNetworkConditionByID", err,
				"Failure at GetDeviceAdminNetworkConditionByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

		vItem2 := flattenDeviceAdministrationNetworkConditionsGetDeviceAdminNetworkConditionByIDItem(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceAdminNetworkConditionByID response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceDeviceAdministrationNetworkConditionsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID, okID := resourceMap["id"]
	vName, okName := resourceMap["name"]
	vvID := vID
	vvName := vName
	method1 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	// NOTE: Consider adding getAllItems and search function to get missing params
	if selectedMethod == 2 {
		getResp1, _, err := client.DeviceAdministrationNetworkConditions.GetDeviceAdminNetworkConditions()
		if err == nil && getResp1 != nil {
			items1 := getAllItemsDeviceAdministrationNetworkConditionsGetDeviceAdminNetworkConditions(m, getResp1)
			item1, err := searchDeviceAdministrationNetworkConditionsGetDeviceAdminNetworkConditions(m, items1, vvName, vvID)
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
	if d.HasChange("item") {
		log.Printf("[DEBUG] vvID %s", vvID)
		request1 := expandRequestDeviceAdministrationNetworkConditionsUpdateDeviceAdminNetworkConditionByID(ctx, "item.0", d)
		log.Printf("[DEBUG] request1 => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.DeviceAdministrationNetworkConditions.UpdateDeviceAdminNetworkConditionByID(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] restyResp1 => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateDeviceAdminNetworkConditionByID", err, restyResp1.String(),
					"Failure at UpdateDeviceAdminNetworkConditionByID, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateDeviceAdminNetworkConditionByID", err,
				"Failure at UpdateDeviceAdminNetworkConditionByID, unexpected response", ""))
			return diags
		}
	}

	return resourceDeviceAdministrationNetworkConditionsRead(ctx, d, m)
}

func resourceDeviceAdministrationNetworkConditionsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID, okID := resourceMap["id"]
	vName, okName := resourceMap["name"]
	vvID := vID
	vvName := vName
	method1 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	// REVIEW: Add getAllItems and search function to get missing params
	if selectedMethod == 2 {

		getResp1, _, err := client.DeviceAdministrationNetworkConditions.GetDeviceAdminNetworkConditions()
		if err != nil || getResp1 == nil {
			// Assume that element it is already gone
			return diags
		}
		items1 := getAllItemsDeviceAdministrationNetworkConditionsGetDeviceAdminNetworkConditions(m, getResp1)
		item1, err := searchDeviceAdministrationNetworkConditionsGetDeviceAdminNetworkConditions(m, items1, vvName, vvID)
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
		getResp, _, err := client.DeviceAdministrationNetworkConditions.GetDeviceAdminNetworkConditionByID(vvID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	response1, restyResp1, err := client.DeviceAdministrationNetworkConditions.DeleteDeviceAdminNetworkConditionByID(vvID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] restyResp1 => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteDeviceAdminNetworkConditionByID", err, restyResp1.String(),
				"Failure at DeleteDeviceAdminNetworkConditionByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteDeviceAdminNetworkConditionByID", err,
			"Failure at DeleteDeviceAdminNetworkConditionByID, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestDeviceAdministrationNetworkConditionsCreateDeviceAdminNetworkCondition(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationNetworkConditionsCreateDeviceAdminNetworkCondition {
	request := isegosdk.RequestDeviceAdministrationNetworkConditionsCreateDeviceAdminNetworkCondition{}
	if v, ok := d.GetOkExists(key + ".condition_type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".condition_type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".condition_type"))) {
		request.ConditionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".id"))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".link"); !isEmptyValue(reflect.ValueOf(d.Get(key+".link"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".link"))) {
		request.Link = expandRequestDeviceAdministrationNetworkConditionsCreateDeviceAdminNetworkConditionLink(ctx, key+".link.0", d)
	}
	if v, ok := d.GetOkExists(key + ".name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".name"))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".conditions"); !isEmptyValue(reflect.ValueOf(d.Get(key+".conditions"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".conditions"))) {
		request.Conditions = expandRequestDeviceAdministrationNetworkConditionsCreateDeviceAdminNetworkConditionConditionsArray(ctx, key, d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceAdministrationNetworkConditionsCreateDeviceAdminNetworkConditionLink(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationNetworkConditionsCreateDeviceAdminNetworkConditionLink {
	request := isegosdk.RequestDeviceAdministrationNetworkConditionsCreateDeviceAdminNetworkConditionLink{}
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

func expandRequestDeviceAdministrationNetworkConditionsCreateDeviceAdminNetworkConditionConditionsArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestDeviceAdministrationNetworkConditionsCreateDeviceAdminNetworkConditionConditions {
	request := []isegosdk.RequestDeviceAdministrationNetworkConditionsCreateDeviceAdminNetworkConditionConditions{}
	o := d.Get(key)
	if o != nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestDeviceAdministrationNetworkConditionsCreateDeviceAdminNetworkConditionConditions(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		request = append(request, *i)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceAdministrationNetworkConditionsCreateDeviceAdminNetworkConditionConditions(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationNetworkConditionsCreateDeviceAdminNetworkConditionConditions {
	request := isegosdk.RequestDeviceAdministrationNetworkConditionsCreateDeviceAdminNetworkConditionConditions{}
	if v, ok := d.GetOkExists(key + ".cli_dnis_list"); !isEmptyValue(reflect.ValueOf(d.Get(key+".cli_dnis_list"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".cli_dnis_list"))) {
		request.CliDnisList = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(key + ".ip_addr_list"); !isEmptyValue(reflect.ValueOf(d.Get(key+".ip_addr_list"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".ip_addr_list"))) {
		request.IPAddrList = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(key + ".mac_addr_list"); !isEmptyValue(reflect.ValueOf(d.Get(key+".mac_addr_list"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".mac_addr_list"))) {
		request.MacAddrList = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(key + ".device_group_list"); !isEmptyValue(reflect.ValueOf(d.Get(key+".device_group_list"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".device_group_list"))) {
		request.DeviceGroupList = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(key + ".device_list"); !isEmptyValue(reflect.ValueOf(d.Get(key+".device_list"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".device_list"))) {
		request.DeviceList = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceAdministrationNetworkConditionsUpdateDeviceAdminNetworkConditionByID(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationNetworkConditionsUpdateDeviceAdminNetworkConditionByID {
	request := isegosdk.RequestDeviceAdministrationNetworkConditionsUpdateDeviceAdminNetworkConditionByID{}
	if v, ok := d.GetOkExists(key + ".condition_type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".condition_type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".condition_type"))) {
		request.ConditionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".id"))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".link"); !isEmptyValue(reflect.ValueOf(d.Get(key+".link"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".link"))) {
		request.Link = expandRequestDeviceAdministrationNetworkConditionsUpdateDeviceAdminNetworkConditionByIDLink(ctx, key+".link.0", d)
	}
	if v, ok := d.GetOkExists(key + ".name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".name"))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".conditions"); !isEmptyValue(reflect.ValueOf(d.Get(key+".conditions"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".conditions"))) {
		request.Conditions = expandRequestDeviceAdministrationNetworkConditionsUpdateDeviceAdminNetworkConditionByIDConditionsArray(ctx, key, d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceAdministrationNetworkConditionsUpdateDeviceAdminNetworkConditionByIDLink(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationNetworkConditionsUpdateDeviceAdminNetworkConditionByIDLink {
	request := isegosdk.RequestDeviceAdministrationNetworkConditionsUpdateDeviceAdminNetworkConditionByIDLink{}
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

func expandRequestDeviceAdministrationNetworkConditionsUpdateDeviceAdminNetworkConditionByIDConditionsArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestDeviceAdministrationNetworkConditionsUpdateDeviceAdminNetworkConditionByIDConditions {
	request := []isegosdk.RequestDeviceAdministrationNetworkConditionsUpdateDeviceAdminNetworkConditionByIDConditions{}
	o := d.Get(key)
	if o != nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestDeviceAdministrationNetworkConditionsUpdateDeviceAdminNetworkConditionByIDConditions(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		request = append(request, *i)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceAdministrationNetworkConditionsUpdateDeviceAdminNetworkConditionByIDConditions(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestDeviceAdministrationNetworkConditionsUpdateDeviceAdminNetworkConditionByIDConditions {
	request := isegosdk.RequestDeviceAdministrationNetworkConditionsUpdateDeviceAdminNetworkConditionByIDConditions{}
	if v, ok := d.GetOkExists(key + ".cli_dnis_list"); !isEmptyValue(reflect.ValueOf(d.Get(key+".cli_dnis_list"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".cli_dnis_list"))) {
		request.CliDnisList = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(key + ".ip_addr_list"); !isEmptyValue(reflect.ValueOf(d.Get(key+".ip_addr_list"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".ip_addr_list"))) {
		request.IPAddrList = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(key + ".mac_addr_list"); !isEmptyValue(reflect.ValueOf(d.Get(key+".mac_addr_list"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".mac_addr_list"))) {
		request.MacAddrList = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(key + ".device_group_list"); !isEmptyValue(reflect.ValueOf(d.Get(key+".device_group_list"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".device_group_list"))) {
		request.DeviceGroupList = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(key + ".device_list"); !isEmptyValue(reflect.ValueOf(d.Get(key+".device_list"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".device_list"))) {
		request.DeviceList = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func getAllItemsDeviceAdministrationNetworkConditionsGetDeviceAdminNetworkConditions(m interface{}, response *isegosdk.ResponseDeviceAdministrationNetworkConditionsGetDeviceAdminNetworkConditions) []isegosdk.ResponseDeviceAdministrationNetworkConditionsGetDeviceAdminNetworkConditionsResponse {
	var respItems []isegosdk.ResponseDeviceAdministrationNetworkConditionsGetDeviceAdminNetworkConditionsResponse
	for response.Response != nil && len(*response.Response) > 0 {
		respItems = append(respItems, *response.Response...)
	}
	return respItems
}

func searchDeviceAdministrationNetworkConditionsGetDeviceAdminNetworkConditions(m interface{}, items []isegosdk.ResponseDeviceAdministrationNetworkConditionsGetDeviceAdminNetworkConditionsResponse, name string, id string) (*isegosdk.ResponseDeviceAdministrationNetworkConditionsGetDeviceAdminNetworkConditionByIDResponse, error) {
	client := m.(*isegosdk.Client)
	var err error
	var foundItem *isegosdk.ResponseDeviceAdministrationNetworkConditionsGetDeviceAdminNetworkConditionByIDResponse
	for _, item := range items {
		if id != "" && item.ID == id {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseDeviceAdministrationNetworkConditionsGetDeviceAdminNetworkConditionByID
			getItem, _, err = client.DeviceAdministrationNetworkConditions.GetDeviceAdminNetworkConditionByID(id)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetDeviceAdminNetworkConditionByID")
			}
			foundItem = getItem.Response
			return foundItem, err
		} else if name != "" && item.Name == name {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseDeviceAdministrationNetworkConditionsGetDeviceAdminNetworkConditionByID
			getItem, _, err = client.DeviceAdministrationNetworkConditions.GetDeviceAdminNetworkConditionByID(item.ID)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetDeviceAdminNetworkConditionByID")
			}
			foundItem = getItem.Response
			return foundItem, err
		}
	}
	return foundItem, err
}
