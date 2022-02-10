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

func resourceTrustsecVn() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on virtualNetwork.

- Create Virtual Network

- Update Virtual Network

- Delete Virtual Network
`,

		CreateContext: resourceTrustsecVnCreate,
		ReadContext:   resourceTrustsecVnRead,
		UpdateContext: resourceTrustsecVnUpdate,
		DeleteContext: resourceTrustsecVnDelete,
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

						"additional_attributes": &schema.Schema{
							Description: `JSON String of additional attributes for the Virtual Network`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"id": &schema.Schema{
							Description: `Identifier of the Virtual Network`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"last_update": &schema.Schema{
							Description: `Timestamp for the last update of the Virtual Network`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"name": &schema.Schema{
							Description: `Name of the Virtual Network`,
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

						"additional_attributes": &schema.Schema{
							Description: `JSON String of additional attributes for the Virtual Network`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"id": &schema.Schema{
							Description: `Identifier of the Virtual Network`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"last_update": &schema.Schema{
							Description: `Timestamp for the last update of the Virtual Network`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"name": &schema.Schema{
							Description: `Name of the Virtual Network`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func resourceTrustsecVnCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning TrustsecVn create")
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestTrustsecVnCreateVirtualNetwork(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vName, _ := resourceItem["name"]
	vvName := interfaceToString(vName)
	if okID && vvID != "" {
		getResponse2, _, err := client.VirtualNetwork.GetVirtualNetworkByID(vvID)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return resourceTrustsecVnRead(ctx, d, m)
		}
	} else {
		queryParams2 := isegosdk.GetVirtualNetworksQueryParams{}

		response2, _, err := client.VirtualNetwork.GetVirtualNetworks(&queryParams2)
		if response2 != nil && err == nil {
			items2 := getAllItemsVirtualNetworkGetVirtualNetworks(m, response2, &queryParams2)
			item2, err := searchVirtualNetworkGetVirtualNetworks(m, items2, vvName, vvID)
			if err == nil && item2 != nil {
				resourceMap := make(map[string]string)
				resourceMap["id"] = vvID
				resourceMap["name"] = vvName
				d.SetId(joinResourceID(resourceMap))
				return resourceTrustsecVnRead(ctx, d, m)
			}
		}
	}
	resp1, restyResp1, err := client.VirtualNetwork.CreateVirtualNetwork(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateVirtualNetwork", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateVirtualNetwork", err))
		return diags
	}
	resourceMap := make(map[string]string)
	resourceMap["id"] = vvID
	resourceMap["name"] = vvName
	d.SetId(joinResourceID(resourceMap))
	return resourceTrustsecVnRead(ctx, d, m)
}

func resourceTrustsecVnRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning TrustsecVn read for id=[%s]", d.Id())
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID, okID := resourceMap["id"]
	vName, okName := resourceMap["name"]
	method1 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	vvID := vID
	vvName := vName
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetVirtualNetworks")
		queryParams1 := isegosdk.GetVirtualNetworksQueryParams{}

		response1, restyResp1, err := client.VirtualNetwork.GetVirtualNetworks(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		items1 := getAllItemsVirtualNetworkGetVirtualNetworks(m, response1, nil)
		item1, err := searchVirtualNetworkGetVirtualNetworks(m, items1, vvName, vvID)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}
		vItem1 := flattenVirtualNetworkGetVirtualNetworkByIDItem(item1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetVirtualNetworks search response",
				err))
			return diags
		}

	}
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetVirtualNetworkByID")
		response2, restyResp2, err := client.VirtualNetwork.GetVirtualNetworkByID(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenVirtualNetworkGetVirtualNetworkByIDItem(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetVirtualNetworkByID response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceTrustsecVnUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning TrustsecVn update for id=[%s]", d.Id())
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID, okID := resourceMap["id"]
	vName, okName := resourceMap["name"]

	method1 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	vvID := vID
	// NOTE: Added getAllItems and search function to get missing params
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetSgVnMappings")
		queryParams1 := isegosdk.GetVirtualNetworksQueryParams{}
		getResp1, _, err := client.VirtualNetwork.GetVirtualNetworks(&queryParams1)

		if err == nil && getResp1 != nil {
			items1 := getAllItemsVirtualNetworkGetVirtualNetworks(m, getResp1, &queryParams1)
			item1, err := searchVirtualNetworkGetVirtualNetworks(m, items1, vName, vID)
			if err == nil && item1 != nil {
				if len(*item1) > 0 {
					vvID = (*item1)[0].ID
				}
			}
		}
	}
	if selectedMethod == 1 {
		vvID = vID
	}
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] ID used for update operation %s", vvID)
		request1 := expandRequestTrustsecVnUpdateVirtualNetworkByID(ctx, "parameters.0", d)
		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
		response1, restyResp1, err := client.VirtualNetwork.UpdateVirtualNetworkByID(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateVirtualNetworkByID", err, restyResp1.String(),
					"Failure at UpdateVirtualNetworkByID, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateVirtualNetworkByID", err,
				"Failure at UpdateVirtualNetworkByID, unexpected response", ""))
			return diags
		}
		d.Set("last_updated", getUnixTimeString())
	}

	return resourceTrustsecVnRead(ctx, d, m)
}

func resourceTrustsecVnDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning TrustsecVn delete for id=[%s]", d.Id())
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID, okID := resourceMap["id"]
	vName, okName := resourceMap["name"]

	method1 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	vvID := vID
	// REVIEW: Add getAllItems and search function to get missing params
	if selectedMethod == 2 {
		queryParams1 := isegosdk.GetVirtualNetworksQueryParams{}

		getResp1, _, err := client.VirtualNetwork.GetVirtualNetworks(&queryParams1)
		if err != nil || getResp1 == nil {
			// Assume that element it is already gone
			return diags
		}
		items1 := getAllItemsVirtualNetworkGetVirtualNetworks(m, getResp1, &queryParams1)
		item1, err := searchVirtualNetworkGetVirtualNetworks(m, items1, vName, vID)
		if err != nil || item1 == nil {
			// Assume that element it is already gone
			return diags
		}
		if item1 != nil && len(*item1) > 0 {
			if vvID != (*item1)[0].ID {
				vvID = (*item1)[0].ID
			} else {
				vvID = vID
			}
		} else {
			vvID = vID
		}
	}
	if selectedMethod == 1 {
		vvID = vID
		getResp, _, err := client.VirtualNetwork.GetVirtualNetworkByID(vvID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	response1, restyResp1, err := client.VirtualNetwork.DeleteVirtualNetworkByID(vvID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteVirtualNetworkByID", err, restyResp1.String(),
				"Failure at DeleteVirtualNetworkByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteVirtualNetworkByID", err,
			"Failure at DeleteVirtualNetworkByID, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestTrustsecVnCreateVirtualNetwork(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestVirtualNetworkCreateVirtualNetwork {
	request := isegosdk.RequestVirtualNetworkCreateVirtualNetwork{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".additional_attributes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".additional_attributes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".additional_attributes")))) {
		request.AdditionalAttributes = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".last_update")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".last_update")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".last_update")))) {
		request.LastUpdate = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestTrustsecVnUpdateVirtualNetworkByID(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestVirtualNetworkUpdateVirtualNetworkByID {
	request := isegosdk.RequestVirtualNetworkUpdateVirtualNetworkByID{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".additional_attributes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".additional_attributes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".additional_attributes")))) {
		request.AdditionalAttributes = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".last_update")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".last_update")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".last_update")))) {
		request.LastUpdate = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func getAllItemsVirtualNetworkGetVirtualNetworks(m interface{}, response *isegosdk.ResponseVirtualNetworkGetVirtualNetworks, queryParams *isegosdk.GetVirtualNetworksQueryParams) []isegosdk.ResponseVirtualNetworkGetVirtualNetworksResponse {
	var respItems []isegosdk.ResponseVirtualNetworkGetVirtualNetworksResponse
	if response.Response != nil && len(*response.Response) > 0 {
		respItems = append(respItems, *response.Response...)
	}
	return respItems
}

func searchVirtualNetworkGetVirtualNetworks(m interface{}, items []isegosdk.ResponseVirtualNetworkGetVirtualNetworksResponse, name string, id string) (*[]isegosdk.ResponseVirtualNetworkGetVirtualNetworkByIDResponse, error) {
	client := m.(*isegosdk.Client)
	var err error
	var foundItem *[]isegosdk.ResponseVirtualNetworkGetVirtualNetworkByIDResponse
	for _, item := range items {
		if id != "" && item.ID == id {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseVirtualNetworkGetVirtualNetworkByID
			getItem, _, err = client.VirtualNetwork.GetVirtualNetworkByID(id)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetVirtualNetworkByID")
			}
			foundItem = getItem.Response
			return foundItem, err
		} else if name != "" && item.Name == name {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseVirtualNetworkGetVirtualNetworkByID
			getItem, _, err = client.VirtualNetwork.GetVirtualNetworkByID(id)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetVirtualNetworkByID")
			}
			foundItem = getItem.Response
			return foundItem, err
		}
	}
	return foundItem, err
}
