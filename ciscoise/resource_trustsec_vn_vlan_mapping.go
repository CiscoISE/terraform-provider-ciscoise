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

func resourceTrustsecVnVLANMapping() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on vnVlanMapping.

- Create VN-Vlan Mapping

- Update VN-Vlan Mapping

- Delete VN-Vlan Mapping
`,

		CreateContext: resourceTrustsecVnVLANMappingCreate,
		ReadContext:   resourceTrustsecVnVLANMappingRead,
		UpdateContext: resourceTrustsecVnVLANMappingUpdate,
		DeleteContext: resourceTrustsecVnVLANMappingDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Description: `Identifier of the VN-Vlan Mapping`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"is_data": &schema.Schema{
							Description: `Flag which indicates whether the Vlan is data or voice type`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"is_default_vlan": &schema.Schema{
							Description: `Flag which indicates if the Vlan is default`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"last_update": &schema.Schema{
							Description: `Timestamp for the last update of the VN-Vlan Mapping`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"max_value": &schema.Schema{
							Description: `Max value`,
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"name": &schema.Schema{
							Description: `Name of the Vlan`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"vn_id": &schema.Schema{
							Description: `Identifier for the associated Virtual Network which is required unless its name is provided`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"vn_name": &schema.Schema{
							Description: `Name of the associated Virtual Network to be used for identity if id is not provided`,
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Description: `Identifier of the VN-Vlan Mapping`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"is_data": &schema.Schema{
							Description: `Flag which indicates whether the Vlan is data or voice type`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"is_default_vlan": &schema.Schema{
							Description: `Flag which indicates if the Vlan is default`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"last_update": &schema.Schema{
							Description: `Timestamp for the last update of the VN-Vlan Mapping`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"max_value": &schema.Schema{
							Description: `Max value`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"name": &schema.Schema{
							Description: `Name of the Vlan`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"vn_id": &schema.Schema{
							Description: `Identifier for the associated Virtual Network which is required unless its name is provided`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"vn_name": &schema.Schema{
							Description: `Name of the associated Virtual Network to be used for identity if id is not provided`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func resourceTrustsecVnVLANMappingCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning TrustsecVnVLANMapping create")
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestTrustsecVnVLANMappingCreateVnVLANMapping(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vName, _ := resourceItem["name"]
	vvName := interfaceToString(vName)
	vVnID, _ := resourceItem["vn_id"]
	vvVnID := interfaceToString(vVnID)
	vVnName, _ := resourceItem["vn_name"]
	vvVnName := interfaceToString(vVnName)
	if okID && vvID != "" {
		getResponse2, _, err := client.VnVLANMapping.GetVnVLANMappingByID(vvID)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			resourceMap["vn_id"] = vvVnID
			resourceMap["vn_name"] = vvVnName
			d.SetId(joinResourceID(resourceMap))
			return resourceTrustsecVnVLANMappingRead(ctx, d, m)
		}
	} else {
		queryParams2 := isegosdk.GetVnVLANMappingsQueryParams{}

		response2, _, err := client.VnVLANMapping.GetVnVLANMappings(&queryParams2)
		if response2 != nil && err == nil {
			items2 := getAllItemsVnVLANMappingGetVnVLANMappings(m, response2, &queryParams2)
			item2, err := searchVnVLANMappingGetVnVLANMappings(m, items2, vvName, vvVnID, vvVnName, vvID)
			if err == nil && item2 != nil {
				resourceMap := make(map[string]string)
				resourceMap["id"] = vvID
				resourceMap["name"] = vvName
				resourceMap["vn_id"] = vvVnID
				resourceMap["vn_name"] = vvVnName
				d.SetId(joinResourceID(resourceMap))
				return resourceTrustsecVnVLANMappingRead(ctx, d, m)
			}
		}
	}
	resp1, restyResp1, err := client.VnVLANMapping.CreateVnVLANMapping(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateVnVLANMapping", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateVnVLANMapping", err))
		return diags
	}
	resourceMap := make(map[string]string)
	resourceMap["id"] = vvID
	resourceMap["name"] = vvName
	resourceMap["vn_id"] = vvVnID
	resourceMap["vn_name"] = vvVnName
	d.SetId(joinResourceID(resourceMap))
	return resourceTrustsecVnVLANMappingRead(ctx, d, m)
}

func resourceTrustsecVnVLANMappingRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning TrustsecVnVLANMapping read for id=[%s]", d.Id())
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID, okID := resourceMap["id"]
	vvID, _ := resourceMap["id"]
	vvName, _ := resourceMap["name"]
	vvVnID, _ := resourceMap["vn_id"]
	vvVnName, _ := resourceMap["vn_name"]

	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetVnVLANMappings")
		queryParams1 := isegosdk.GetVnVLANMappingsQueryParams{}

		response1, restyResp1, err := client.VnVLANMapping.GetVnVLANMappings(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		items1 := getAllItemsVnVLANMappingGetVnVLANMappings(m, response1, nil)
		item1, err := searchVnVLANMappingGetVnVLANMappings(m, items1, vvName, vvVnID, vvVnName, vvID)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}
		vItem1 := flattenVnVLANMappingGetVnVLANMappingByIDItem(item1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetVnVLANMappings search response",
				err))
			return diags
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetVnVLANMappingByID")
		vvID = vID

		response2, restyResp2, err := client.VnVLANMapping.GetVnVLANMappingByID(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenVnVLANMappingGetVnVLANMappingByIDItem(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetVnVLANMappingByID response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceTrustsecVnVLANMappingUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning TrustsecVnVLANMapping update for id=[%s]", d.Id())
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID, okID := resourceMap["id"]
	vvID, _ := resourceMap["id"]
	vvName, _ := resourceMap["name"]
	vvVnID, _ := resourceMap["vn_id"]
	vvVnName, _ := resourceMap["vn_name"]

	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	// NOTE: Added getAllItems and search function to get missing params
	if selectedMethod == 1 {
		queryParams1 := isegosdk.GetVnVLANMappingsQueryParams{}
		getResp1, _, err := client.VnVLANMapping.GetVnVLANMappings(&queryParams1)
		if err == nil && getResp1 != nil {
			items1 := getAllItemsVnVLANMappingGetVnVLANMappings(m, getResp1, &queryParams1)
			item1, err := searchVnVLANMappingGetVnVLANMappings(m, items1, vvName, vvVnID, vvVnName, vvID)
			vvID = vID
			if err == nil && item1 != nil {
				if len(*item1) > 0 && (*item1)[0].ID != "" {
					if vID != (*item1)[0].ID {
						vvID = (*item1)[0].ID
					}
				}
			}
		}
	}
	if selectedMethod == 2 {
		vvID = vID
	}
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] ID used for update operation %s", vvID)
		request1 := expandRequestTrustsecVnVLANMappingUpdateVnVLANMappingByID(ctx, "parameters.0", d)
		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
		response1, restyResp1, err := client.VnVLANMapping.UpdateVnVLANMappingByID(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateVnVLANMappingByID", err, restyResp1.String(),
					"Failure at UpdateVnVLANMappingByID, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateVnVLANMappingByID", err,
				"Failure at UpdateVnVLANMappingByID, unexpected response", ""))
			return diags
		}
	}

	return resourceTrustsecVnVLANMappingRead(ctx, d, m)
}

func resourceTrustsecVnVLANMappingDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning TrustsecVnVLANMapping delete for id=[%s]", d.Id())
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID, okID := resourceMap["id"]
	vvID, _ := resourceMap["id"]
	vvName, _ := resourceMap["name"]
	vvVnID, _ := resourceMap["vn_id"]
	vvVnName, _ := resourceMap["vn_name"]

	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	// REVIEW: Add getAllItems and search function to get missing params
	if selectedMethod == 1 {
		queryParams1 := isegosdk.GetVnVLANMappingsQueryParams{}

		getResp1, _, err := client.VnVLANMapping.GetVnVLANMappings(&queryParams1)
		if err != nil || getResp1 == nil {
			// Assume that element it is already gone
			return diags
		}
		items1 := getAllItemsVnVLANMappingGetVnVLANMappings(m, getResp1, &queryParams1)
		item1, err := searchVnVLANMappingGetVnVLANMappings(m, items1, vvName, vvVnID, vvVnName, vvID)
		if err != nil || item1 == nil {
			// Assume that element it is already gone
			return diags
		}
		vvID = vID
		if len(*item1) > 0 && (*item1)[0].ID != "" {
			if vID != (*item1)[0].ID {
				vvID = (*item1)[0].ID
			}
		}
	}
	if selectedMethod == 2 {
		vvID = vID
		getResp, _, err := client.VnVLANMapping.GetVnVLANMappingByID(vvID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	response1, restyResp1, err := client.VnVLANMapping.DeleteVnVLANMappingByID(vvID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteVnVLANMappingByID", err, restyResp1.String(),
				"Failure at DeleteVnVLANMappingByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteVnVLANMappingByID", err,
			"Failure at DeleteVnVLANMappingByID, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestTrustsecVnVLANMappingCreateVnVLANMapping(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestVnVLANMappingCreateVnVLANMapping {
	request := isegosdk.RequestVnVLANMappingCreateVnVLANMapping{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_data")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_data")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_data")))) {
		request.IsData = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_default_vlan")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_default_vlan")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_default_vlan")))) {
		request.IsDefaultVLAN = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".last_update")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".last_update")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".last_update")))) {
		request.LastUpdate = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".max_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".max_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".max_value")))) {
		request.MaxValue = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vn_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vn_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vn_id")))) {
		request.VnID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vn_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vn_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vn_name")))) {
		request.VnName = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestTrustsecVnVLANMappingUpdateVnVLANMappingByID(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestVnVLANMappingUpdateVnVLANMappingByID {
	request := isegosdk.RequestVnVLANMappingUpdateVnVLANMappingByID{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_data")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_data")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_data")))) {
		request.IsData = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_default_vlan")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_default_vlan")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_default_vlan")))) {
		request.IsDefaultVLAN = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".last_update")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".last_update")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".last_update")))) {
		request.LastUpdate = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".max_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".max_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".max_value")))) {
		request.MaxValue = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vn_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vn_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vn_id")))) {
		request.VnID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vn_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vn_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vn_name")))) {
		request.VnName = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func getAllItemsVnVLANMappingGetVnVLANMappings(m interface{}, response *isegosdk.ResponseVnVLANMappingGetVnVLANMappings, queryParams *isegosdk.GetVnVLANMappingsQueryParams) []isegosdk.ResponseVnVLANMappingGetVnVLANMappingsResponse {
	var respItems []isegosdk.ResponseVnVLANMappingGetVnVLANMappingsResponse
	if response.Response != nil && len(*response.Response) > 0 {
		respItems = append(respItems, *response.Response...)
	}
	return respItems
}

func searchVnVLANMappingGetVnVLANMappings(m interface{}, items []isegosdk.ResponseVnVLANMappingGetVnVLANMappingsResponse, name string, vnID string, vnName string, id string) (*[]isegosdk.ResponseVnVLANMappingGetVnVLANMappingByIDResponse, error) {
	client := m.(*isegosdk.Client)
	var err error
	var foundItem *[]isegosdk.ResponseVnVLANMappingGetVnVLANMappingByIDResponse
	for _, item := range items {
		if id != "" && item.ID == id {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseVnVLANMappingGetVnVLANMappingByID
			getItem, _, err = client.VnVLANMapping.GetVnVLANMappingByID(id)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetVnVLANMappingByID")
			}
			foundItem = getItem.Response
			return foundItem, err
		} else if name != "" && item.Name == name {
			hasVnID := (vnID != "" && vnID == item.VnID)
			hasVnName := (vnName != "" && vnName == item.VnName)
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseVnVLANMappingGetVnVLANMappingByID
			getItem, _, err = client.VnVLANMapping.GetVnVLANMappingByID(item.ID)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetVnVLANMappingByID")
			}
			if hasVnID || hasVnName {
				foundItem = getItem.Response
				return foundItem, err
			}
		}
	}
	return foundItem, err
}
