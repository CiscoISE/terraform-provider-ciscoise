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

func resourceNetworkAccessNetworkCondition() *schema.Resource {
	return &schema.Resource{

		CreateContext: resourceNetworkAccessNetworkConditionCreate,
		ReadContext:   resourceNetworkAccessNetworkConditionRead,
		UpdateContext: resourceNetworkAccessNetworkConditionUpdate,
		DeleteContext: resourceNetworkAccessNetworkConditionDelete,
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

func resourceNetworkAccessNetworkConditionCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("item"))
	request1 := expandRequestNetworkAccessNetworkConditionCreateNetworkAccessNetworkCondition(ctx, "item.0", d)
	log.Printf("[DEBUG] request1 => %v", responseInterfaceToString(*request1))

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vName, _ := resourceItem["name"]
	vvName := interfaceToString(vName)
	if okID && vvID != "" {
		getResponse2, _, err := client.NetworkAccessNetworkConditions.GetNetworkAccessNetworkConditionByID(vvID)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return diags
		}
	} else {
		response2, _, err := client.NetworkAccessNetworkConditions.GetNetworkAccessNetworkConditions()
		if response2 != nil && err == nil {
			items2 := getAllItemsNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditions(m, response2)
			item2, err := searchNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditions(m, items2, vvName, vvID)
			if err == nil && item2 != nil {
				resourceMap := make(map[string]string)
				resourceMap["id"] = vvID
				resourceMap["name"] = vvName
				d.SetId(joinResourceID(resourceMap))
				return diags
			}
		}
	}
	resp1, restyResp1, err := client.NetworkAccessNetworkConditions.CreateNetworkAccessNetworkCondition(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateNetworkAccessNetworkCondition", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateNetworkAccessNetworkCondition", err))
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

func resourceNetworkAccessNetworkConditionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		log.Printf("[DEBUG] Selected method: GetNetworkAccessNetworkConditions")

		response1, _, err := client.NetworkAccessNetworkConditions.GetNetworkAccessNetworkConditions()

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetNetworkAccessNetworkConditions", err,
				"Failure at GetNetworkAccessNetworkConditions, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		items1 := getAllItemsNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditions(m, response1)
		item1, err := searchNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditions(m, items1, vvName, vvID)
		if err != nil || item1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when searching item from GetNetworkAccessNetworkConditions response", err,
				"Failure when searching item from GetNetworkAccessNetworkConditions, unexpected response", ""))
			return diags
		}
		if err := d.Set("item", item1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNetworkAccessNetworkConditions search response",
				err))
			return diags
		}

	}
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkAccessNetworkConditionByID")

		response2, _, err := client.NetworkAccessNetworkConditions.GetNetworkAccessNetworkConditionByID(vvID)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetNetworkAccessNetworkConditionByID", err,
				"Failure at GetNetworkAccessNetworkConditionByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

		vItem2 := flattenNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditionByIDItem(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNetworkAccessNetworkConditionByID response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceNetworkAccessNetworkConditionUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		getResp1, _, err := client.NetworkAccessNetworkConditions.GetNetworkAccessNetworkConditions()
		if err == nil && getResp1 != nil {
			items1 := getAllItemsNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditions(m, getResp1)
			item1, err := searchNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditions(m, items1, vvName, vvID)
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
		request1 := expandRequestNetworkAccessNetworkConditionUpdateNetworkAccessNetworkConditionByID(ctx, "item.0", d)
		log.Printf("[DEBUG] request1 => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.NetworkAccessNetworkConditions.UpdateNetworkAccessNetworkConditionByID(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] restyResp1 => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateNetworkAccessNetworkConditionByID", err, restyResp1.String(),
					"Failure at UpdateNetworkAccessNetworkConditionByID, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateNetworkAccessNetworkConditionByID", err,
				"Failure at UpdateNetworkAccessNetworkConditionByID, unexpected response", ""))
			return diags
		}
	}

	return resourceNetworkAccessNetworkConditionRead(ctx, d, m)
}

func resourceNetworkAccessNetworkConditionDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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

		getResp1, _, err := client.NetworkAccessNetworkConditions.GetNetworkAccessNetworkConditions()
		if err != nil || getResp1 == nil {
			// Assume that element it is already gone
			return diags
		}
		items1 := getAllItemsNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditions(m, getResp1)
		item1, err := searchNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditions(m, items1, vvName, vvID)
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
		getResp, _, err := client.NetworkAccessNetworkConditions.GetNetworkAccessNetworkConditionByID(vvID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	response1, restyResp1, err := client.NetworkAccessNetworkConditions.DeleteNetworkAccessNetworkConditionByID(vvID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] restyResp1 => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteNetworkAccessNetworkConditionByID", err, restyResp1.String(),
				"Failure at DeleteNetworkAccessNetworkConditionByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteNetworkAccessNetworkConditionByID", err,
			"Failure at DeleteNetworkAccessNetworkConditionByID, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestNetworkAccessNetworkConditionCreateNetworkAccessNetworkCondition(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessNetworkConditionsCreateNetworkAccessNetworkCondition {
	request := isegosdk.RequestNetworkAccessNetworkConditionsCreateNetworkAccessNetworkCondition{}
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
		request.Link = expandRequestNetworkAccessNetworkConditionCreateNetworkAccessNetworkConditionLink(ctx, key+".link.0", d)
	}
	if v, ok := d.GetOkExists(key + ".name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".name"))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".conditions"); !isEmptyValue(reflect.ValueOf(d.Get(key+".conditions"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".conditions"))) {
		request.Conditions = expandRequestNetworkAccessNetworkConditionCreateNetworkAccessNetworkConditionConditionsArray(ctx, key, d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkAccessNetworkConditionCreateNetworkAccessNetworkConditionLink(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessNetworkConditionsCreateNetworkAccessNetworkConditionLink {
	request := isegosdk.RequestNetworkAccessNetworkConditionsCreateNetworkAccessNetworkConditionLink{}
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

func expandRequestNetworkAccessNetworkConditionCreateNetworkAccessNetworkConditionConditionsArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestNetworkAccessNetworkConditionsCreateNetworkAccessNetworkConditionConditions {
	request := []isegosdk.RequestNetworkAccessNetworkConditionsCreateNetworkAccessNetworkConditionConditions{}
	o := d.Get(key)
	if o != nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestNetworkAccessNetworkConditionCreateNetworkAccessNetworkConditionConditions(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		request = append(request, *i)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkAccessNetworkConditionCreateNetworkAccessNetworkConditionConditions(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessNetworkConditionsCreateNetworkAccessNetworkConditionConditions {
	request := isegosdk.RequestNetworkAccessNetworkConditionsCreateNetworkAccessNetworkConditionConditions{}
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

func expandRequestNetworkAccessNetworkConditionUpdateNetworkAccessNetworkConditionByID(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessNetworkConditionsUpdateNetworkAccessNetworkConditionByID {
	request := isegosdk.RequestNetworkAccessNetworkConditionsUpdateNetworkAccessNetworkConditionByID{}
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
		request.Link = expandRequestNetworkAccessNetworkConditionUpdateNetworkAccessNetworkConditionByIDLink(ctx, key+".link.0", d)
	}
	if v, ok := d.GetOkExists(key + ".name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".name"))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".conditions"); !isEmptyValue(reflect.ValueOf(d.Get(key+".conditions"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".conditions"))) {
		request.Conditions = expandRequestNetworkAccessNetworkConditionUpdateNetworkAccessNetworkConditionByIDConditionsArray(ctx, key, d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkAccessNetworkConditionUpdateNetworkAccessNetworkConditionByIDLink(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessNetworkConditionsUpdateNetworkAccessNetworkConditionByIDLink {
	request := isegosdk.RequestNetworkAccessNetworkConditionsUpdateNetworkAccessNetworkConditionByIDLink{}
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

func expandRequestNetworkAccessNetworkConditionUpdateNetworkAccessNetworkConditionByIDConditionsArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestNetworkAccessNetworkConditionsUpdateNetworkAccessNetworkConditionByIDConditions {
	request := []isegosdk.RequestNetworkAccessNetworkConditionsUpdateNetworkAccessNetworkConditionByIDConditions{}
	o := d.Get(key)
	if o != nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestNetworkAccessNetworkConditionUpdateNetworkAccessNetworkConditionByIDConditions(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		request = append(request, *i)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkAccessNetworkConditionUpdateNetworkAccessNetworkConditionByIDConditions(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessNetworkConditionsUpdateNetworkAccessNetworkConditionByIDConditions {
	request := isegosdk.RequestNetworkAccessNetworkConditionsUpdateNetworkAccessNetworkConditionByIDConditions{}
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

func getAllItemsNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditions(m interface{}, response *isegosdk.ResponseNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditions) []isegosdk.ResponseNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditionsResponse {
	var respItems []isegosdk.ResponseNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditionsResponse
	for response.Response != nil && len(*response.Response) > 0 {
		respItems = append(respItems, *response.Response...)
	}
	return respItems
}

func searchNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditions(m interface{}, items []isegosdk.ResponseNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditionsResponse, name string, id string) (*isegosdk.ResponseNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditionByIDResponse, error) {
	client := m.(*isegosdk.Client)
	var err error
	var foundItem *isegosdk.ResponseNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditionByIDResponse
	for _, item := range items {
		if id != "" && item.ID == id {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditionByID
			getItem, _, err = client.NetworkAccessNetworkConditions.GetNetworkAccessNetworkConditionByID(id)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetNetworkAccessNetworkConditionByID")
			}
			foundItem = getItem.Response
			return foundItem, err
		} else if name != "" && item.Name == name {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseNetworkAccessNetworkConditionsGetNetworkAccessNetworkConditionByID
			getItem, _, err = client.NetworkAccessNetworkConditions.GetNetworkAccessNetworkConditionByID(item.ID)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetNetworkAccessNetworkConditionByID")
			}
			foundItem = getItem.Response
			return foundItem, err
		}
	}
	return foundItem, err
}
