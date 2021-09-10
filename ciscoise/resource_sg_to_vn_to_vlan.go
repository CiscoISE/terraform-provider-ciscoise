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

func resourceSgToVnToVLAN() *schema.Resource {
	return &schema.Resource{

		CreateContext: resourceSgToVnToVLANCreate,
		ReadContext:   resourceSgToVnToVLANRead,
		UpdateContext: resourceSgToVnToVLANUpdate,
		DeleteContext: resourceSgToVnToVLANDelete,
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

						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
							Optional: true,
							Computed: true,
						},
						"sgt_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"virtualnetworklist": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"default_virtual_network": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"description": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"id": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"vlans": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"data": &schema.Schema{
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
												},
												"default_vlan": &schema.Schema{
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
												},
												"description": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"id": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"max_value": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"name": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
											},
										},
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

func resourceSgToVnToVLANCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("item"))
	request1 := expandRequestSgToVnToVLANCreateSecurityGroupsToVnToVLAN(ctx, "item.0", d)
	log.Printf("[DEBUG] request1 => %v", responseInterfaceToString(*request1))

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vName, _ := resourceItem["name"]
	vvName := interfaceToString(vName)

	if okID && vvID != "" {
		getResponse2, _, err := client.SecurityGroupToVirtualNetwork.GetSecurityGroupsToVnToVLANByID(vvID)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return diags
		}
	} else {
		queryParams2 := isegosdk.GetSecurityGroupsToVnToVLANQueryParams{}

		response2, _, err := client.SecurityGroupToVirtualNetwork.GetSecurityGroupsToVnToVLAN(&queryParams2)
		if response2 != nil && err == nil {
			items2 := getAllItemsSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLAN(m, response2, &queryParams2)
			item2, err := searchSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLAN(m, items2, vvName, vvID)
			if err == nil && item2 != nil {
				resourceMap := make(map[string]string)
				resourceMap["id"] = vvID
				resourceMap["name"] = vvName
				d.SetId(joinResourceID(resourceMap))
				return diags
			}
		}
	}
	restyResp1, err := client.SecurityGroupToVirtualNetwork.CreateSecurityGroupsToVnToVLAN(request1)
	if err != nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateSecurityGroupsToVnToVLAN", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateSecurityGroupsToVnToVLAN", err))
		return diags
	}
	headers := restyResp1.Header()
	if locationHeader, ok := headers["Location"]; ok && len(locationHeader) > 0 {
		vvID = getLocationID(locationHeader[0])
	}
	resourceMap := make(map[string]string)
	resourceMap["id"] = vvID
	resourceMap["name"] = vvName
	d.SetId(joinResourceID(resourceMap))
	return diags
}

func resourceSgToVnToVLANRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		vvName := vName
		vvID := vID
		log.Printf("[DEBUG] Selected method: GetSecurityGroupsToVnToVLAN")
		queryParams1 := isegosdk.GetSecurityGroupsToVnToVLANQueryParams{}

		response1, _, err := client.SecurityGroupToVirtualNetwork.GetSecurityGroupsToVnToVLAN(&queryParams1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetSecurityGroupsToVnToVLAN", err,
				"Failure at GetSecurityGroupsToVnToVLAN, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		items1 := getAllItemsSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLAN(m, response1, &queryParams1)
		item1, err := searchSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLAN(m, items1, vvName, vvID)
		if err != nil || item1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when searching item from GetSecurityGroupsToVnToVLAN response", err,
				"Failure when searching item from GetSecurityGroupsToVnToVLAN, unexpected response", ""))
			return diags
		}
		if err := d.Set("item", item1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSecurityGroupsToVnToVLAN search response",
				err))
			return diags
		}

	}
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetSecurityGroupsToVnToVLANByID")
		vvID := vID

		response2, _, err := client.SecurityGroupToVirtualNetwork.GetSecurityGroupsToVnToVLANByID(vvID)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetSecurityGroupsToVnToVLANByID", err,
				"Failure at GetSecurityGroupsToVnToVLANByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

		vItem2 := flattenSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLANByIDItem(response2.SgtVnVLANContainer)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSecurityGroupsToVnToVLANByID response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceSgToVnToVLANUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
	// NOTE: Consider adding getAllItems and search function to get missing params
	if selectedMethod == 2 {
		queryParams1 := isegosdk.GetSecurityGroupsToVnToVLANQueryParams{}

		getResp1, _, err := client.SecurityGroupToVirtualNetwork.GetSecurityGroupsToVnToVLAN(&queryParams1)
		if err == nil && getResp1 != nil {
			items1 := getAllItemsSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLAN(m, getResp1, &queryParams1)
			item1, err := searchSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLAN(m, items1, vName, vID)
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
		request1 := expandRequestSgToVnToVLANUpdateSecurityGroupsToVnToVLANByID(ctx, "item.0", d)
		log.Printf("[DEBUG] request1 => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.SecurityGroupToVirtualNetwork.UpdateSecurityGroupsToVnToVLANByID(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] restyResp1 => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateSecurityGroupsToVnToVLANByID", err, restyResp1.String(),
					"Failure at UpdateSecurityGroupsToVnToVLANByID, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateSecurityGroupsToVnToVLANByID", err,
				"Failure at UpdateSecurityGroupsToVnToVLANByID, unexpected response", ""))
			return diags
		}
	}

	return resourceSgToVnToVLANRead(ctx, d, m)
}

func resourceSgToVnToVLANDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
	// REVIEW: Add getAllItems and search function to get missing params
	if selectedMethod == 2 {
		queryParams1 := isegosdk.GetSecurityGroupsToVnToVLANQueryParams{}

		getResp1, _, err := client.SecurityGroupToVirtualNetwork.GetSecurityGroupsToVnToVLAN(&queryParams1)
		if err != nil || getResp1 == nil {
			// Assume that element it is already gone
			return diags
		}
		items1 := getAllItemsSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLAN(m, getResp1, &queryParams1)
		item1, err := searchSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLAN(m, items1, vName, vID)
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
		getResp, _, err := client.SecurityGroupToVirtualNetwork.GetSecurityGroupsToVnToVLANByID(vvID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	restyResp1, err := client.SecurityGroupToVirtualNetwork.DeleteSecurityGroupsToVnToVLANByID(vvID)
	if err != nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] restyResp1 => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteSecurityGroupsToVnToVLANByID", err, restyResp1.String(),
				"Failure at DeleteSecurityGroupsToVnToVLANByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteSecurityGroupsToVnToVLANByID", err,
			"Failure at DeleteSecurityGroupsToVnToVLANByID, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestSgToVnToVLANCreateSecurityGroupsToVnToVLAN(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSecurityGroupToVirtualNetworkCreateSecurityGroupsToVnToVLAN {
	request := isegosdk.RequestSecurityGroupToVirtualNetworkCreateSecurityGroupsToVnToVLAN{}
	request.SgtVnVLANContainer = expandRequestSgToVnToVLANCreateSecurityGroupsToVnToVLANSgtVnVLANContainer(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSgToVnToVLANCreateSecurityGroupsToVnToVLANSgtVnVLANContainer(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSecurityGroupToVirtualNetworkCreateSecurityGroupsToVnToVLANSgtVnVLANContainer {
	request := isegosdk.RequestSecurityGroupToVirtualNetworkCreateSecurityGroupsToVnToVLANSgtVnVLANContainer{}
	if v, ok := d.GetOkExists(key + ".id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".id"))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".name"))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".sgt_id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".sgt_id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".sgt_id"))) {
		request.SgtID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".virtualnetworklist"); !isEmptyValue(reflect.ValueOf(d.Get(key+".virtualnetworklist"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".virtualnetworklist"))) {
		request.Virtualnetworklist = expandRequestSgToVnToVLANCreateSecurityGroupsToVnToVLANSgtVnVLANContainerVirtualnetworklistArray(ctx, key, d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSgToVnToVLANCreateSecurityGroupsToVnToVLANSgtVnVLANContainerVirtualnetworklistArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestSecurityGroupToVirtualNetworkCreateSecurityGroupsToVnToVLANSgtVnVLANContainerVirtualnetworklist {
	request := []isegosdk.RequestSecurityGroupToVirtualNetworkCreateSecurityGroupsToVnToVLANSgtVnVLANContainerVirtualnetworklist{}
	o := d.Get(key)
	if o != nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestSgToVnToVLANCreateSecurityGroupsToVnToVLANSgtVnVLANContainerVirtualnetworklist(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		request = append(request, *i)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSgToVnToVLANCreateSecurityGroupsToVnToVLANSgtVnVLANContainerVirtualnetworklist(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSecurityGroupToVirtualNetworkCreateSecurityGroupsToVnToVLANSgtVnVLANContainerVirtualnetworklist {
	request := isegosdk.RequestSecurityGroupToVirtualNetworkCreateSecurityGroupsToVnToVLANSgtVnVLANContainerVirtualnetworklist{}
	if v, ok := d.GetOkExists(key + ".id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".id"))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".name"))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".default_virtual_network"); !isEmptyValue(reflect.ValueOf(d.Get(key+".default_virtual_network"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".default_virtual_network"))) {
		request.DefaultVirtualNetwork = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".vlans"); !isEmptyValue(reflect.ValueOf(d.Get(key+".vlans"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".vlans"))) {
		request.VLANs = expandRequestSgToVnToVLANCreateSecurityGroupsToVnToVLANSgtVnVLANContainerVirtualnetworklistVLANsArray(ctx, key, d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSgToVnToVLANCreateSecurityGroupsToVnToVLANSgtVnVLANContainerVirtualnetworklistVLANsArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestSecurityGroupToVirtualNetworkCreateSecurityGroupsToVnToVLANSgtVnVLANContainerVirtualnetworklistVLANs {
	request := []isegosdk.RequestSecurityGroupToVirtualNetworkCreateSecurityGroupsToVnToVLANSgtVnVLANContainerVirtualnetworklistVLANs{}
	o := d.Get(key)
	if o != nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestSgToVnToVLANCreateSecurityGroupsToVnToVLANSgtVnVLANContainerVirtualnetworklistVLANs(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		request = append(request, *i)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSgToVnToVLANCreateSecurityGroupsToVnToVLANSgtVnVLANContainerVirtualnetworklistVLANs(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSecurityGroupToVirtualNetworkCreateSecurityGroupsToVnToVLANSgtVnVLANContainerVirtualnetworklistVLANs {
	request := isegosdk.RequestSecurityGroupToVirtualNetworkCreateSecurityGroupsToVnToVLANSgtVnVLANContainerVirtualnetworklistVLANs{}
	if v, ok := d.GetOkExists(key + ".id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".id"))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".name"))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".default_vlan"); !isEmptyValue(reflect.ValueOf(d.Get(key+".default_vlan"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".default_vlan"))) {
		request.DefaultVLAN = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".max_value"); !isEmptyValue(reflect.ValueOf(d.Get(key+".max_value"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".max_value"))) {
		request.MaxValue = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".data"); !isEmptyValue(reflect.ValueOf(d.Get(key+".data"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".data"))) {
		request.Data = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSgToVnToVLANUpdateSecurityGroupsToVnToVLANByID(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSecurityGroupToVirtualNetworkUpdateSecurityGroupsToVnToVLANByID {
	request := isegosdk.RequestSecurityGroupToVirtualNetworkUpdateSecurityGroupsToVnToVLANByID{}
	request.SgtVnVLANContainer = expandRequestSgToVnToVLANUpdateSecurityGroupsToVnToVLANByIDSgtVnVLANContainer(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSgToVnToVLANUpdateSecurityGroupsToVnToVLANByIDSgtVnVLANContainer(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSecurityGroupToVirtualNetworkUpdateSecurityGroupsToVnToVLANByIDSgtVnVLANContainer {
	request := isegosdk.RequestSecurityGroupToVirtualNetworkUpdateSecurityGroupsToVnToVLANByIDSgtVnVLANContainer{}
	if v, ok := d.GetOkExists(key + ".id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".id"))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".name"))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".sgt_id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".sgt_id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".sgt_id"))) {
		request.SgtID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".virtualnetworklist"); !isEmptyValue(reflect.ValueOf(d.Get(key+".virtualnetworklist"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".virtualnetworklist"))) {
		request.Virtualnetworklist = expandRequestSgToVnToVLANUpdateSecurityGroupsToVnToVLANByIDSgtVnVLANContainerVirtualnetworklistArray(ctx, key, d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSgToVnToVLANUpdateSecurityGroupsToVnToVLANByIDSgtVnVLANContainerVirtualnetworklistArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestSecurityGroupToVirtualNetworkUpdateSecurityGroupsToVnToVLANByIDSgtVnVLANContainerVirtualnetworklist {
	request := []isegosdk.RequestSecurityGroupToVirtualNetworkUpdateSecurityGroupsToVnToVLANByIDSgtVnVLANContainerVirtualnetworklist{}
	o := d.Get(key)
	if o != nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestSgToVnToVLANUpdateSecurityGroupsToVnToVLANByIDSgtVnVLANContainerVirtualnetworklist(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		request = append(request, *i)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSgToVnToVLANUpdateSecurityGroupsToVnToVLANByIDSgtVnVLANContainerVirtualnetworklist(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSecurityGroupToVirtualNetworkUpdateSecurityGroupsToVnToVLANByIDSgtVnVLANContainerVirtualnetworklist {
	request := isegosdk.RequestSecurityGroupToVirtualNetworkUpdateSecurityGroupsToVnToVLANByIDSgtVnVLANContainerVirtualnetworklist{}
	if v, ok := d.GetOkExists(key + ".id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".id"))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".name"))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".default_virtual_network"); !isEmptyValue(reflect.ValueOf(d.Get(key+".default_virtual_network"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".default_virtual_network"))) {
		request.DefaultVirtualNetwork = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".vlans"); !isEmptyValue(reflect.ValueOf(d.Get(key+".vlans"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".vlans"))) {
		request.VLANs = expandRequestSgToVnToVLANUpdateSecurityGroupsToVnToVLANByIDSgtVnVLANContainerVirtualnetworklistVLANsArray(ctx, key, d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSgToVnToVLANUpdateSecurityGroupsToVnToVLANByIDSgtVnVLANContainerVirtualnetworklistVLANsArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestSecurityGroupToVirtualNetworkUpdateSecurityGroupsToVnToVLANByIDSgtVnVLANContainerVirtualnetworklistVLANs {
	request := []isegosdk.RequestSecurityGroupToVirtualNetworkUpdateSecurityGroupsToVnToVLANByIDSgtVnVLANContainerVirtualnetworklistVLANs{}
	o := d.Get(key)
	if o != nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestSgToVnToVLANUpdateSecurityGroupsToVnToVLANByIDSgtVnVLANContainerVirtualnetworklistVLANs(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		request = append(request, *i)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSgToVnToVLANUpdateSecurityGroupsToVnToVLANByIDSgtVnVLANContainerVirtualnetworklistVLANs(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSecurityGroupToVirtualNetworkUpdateSecurityGroupsToVnToVLANByIDSgtVnVLANContainerVirtualnetworklistVLANs {
	request := isegosdk.RequestSecurityGroupToVirtualNetworkUpdateSecurityGroupsToVnToVLANByIDSgtVnVLANContainerVirtualnetworklistVLANs{}
	if v, ok := d.GetOkExists(key + ".id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".id"))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".name"))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".default_vlan"); !isEmptyValue(reflect.ValueOf(d.Get(key+".default_vlan"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".default_vlan"))) {
		request.DefaultVLAN = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".max_value"); !isEmptyValue(reflect.ValueOf(d.Get(key+".max_value"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".max_value"))) {
		request.MaxValue = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".data"); !isEmptyValue(reflect.ValueOf(d.Get(key+".data"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".data"))) {
		request.Data = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func getAllItemsSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLAN(m interface{}, response *isegosdk.ResponseSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLAN, queryParams *isegosdk.GetSecurityGroupsToVnToVLANQueryParams) []isegosdk.ResponseSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLANSearchResultResources {
	client := m.(*isegosdk.Client)
	var respItems []isegosdk.ResponseSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLANSearchResultResources
	for response.SearchResult != nil && response.SearchResult.Resources != nil && len(*response.SearchResult.Resources) > 0 {
		respItems = append(respItems, *response.SearchResult.Resources...)
		if response.SearchResult.NextPage != nil && response.SearchResult.NextPage.Rel == "next" {
			href := response.SearchResult.NextPage.Href
			page, size, err := getNextPageAndSizeParams(href)
			if err != nil {
				break
			}
			if queryParams != nil {
				queryParams.Page = page
				queryParams.Size = size
			}
			response, _, err = client.SecurityGroupToVirtualNetwork.GetSecurityGroupsToVnToVLAN(queryParams)
			if err != nil {
				break
			}
			// All is good, continue to the next page
			continue
		}
		// Does not have next page finish iteration
		break
	}
	return respItems
}

func searchSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLAN(m interface{}, items []isegosdk.ResponseSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLANSearchResultResources, name string, id string) (*isegosdk.ResponseSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLANByIDSgtVnVLANContainer, error) {
	client := m.(*isegosdk.Client)
	var err error
	var foundItem *isegosdk.ResponseSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLANByIDSgtVnVLANContainer
	for _, item := range items {
		if id != "" && item.ID == id {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLANByID
			getItem, _, err = client.SecurityGroupToVirtualNetwork.GetSecurityGroupsToVnToVLANByID(id)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetSecurityGroupsToVnToVLANByID")
			}
			foundItem = getItem.SgtVnVLANContainer
			return foundItem, err
		} else if name != "" && item.Name == name {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseSecurityGroupToVirtualNetworkGetSecurityGroupsToVnToVLANByID
			getItem, _, err = client.SecurityGroupToVirtualNetwork.GetSecurityGroupsToVnToVLANByID(item.ID)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetSecurityGroupsToVnToVLANByID")
			}
			foundItem = getItem.SgtVnVLANContainer
			return foundItem, err
		}
	}
	return foundItem, err
}
