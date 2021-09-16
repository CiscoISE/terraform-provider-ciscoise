package ciscoise

import (
	"context"
	"reflect"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkDeviceGroup() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on NetworkDeviceGroup.
  
  This resource allows the client to update a network device group.
  This resource deletes a network device group.
  This resource creates a network device group.`,

		CreateContext: resourceNetworkDeviceGroupCreate,
		ReadContext:   resourceNetworkDeviceGroupRead,
		UpdateContext: resourceNetworkDeviceGroupUpdate,
		DeleteContext: resourceNetworkDeviceGroupDelete,
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
						"othername": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceNetworkDeviceGroupCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("item"))
	request1 := expandRequestNetworkDeviceGroupCreateNetworkDeviceGroup(ctx, "item.0", d)
	log.Printf("[DEBUG] request1 => %v", responseInterfaceToString(*request1))

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vName, okName := resourceItem["name"]
	vvName := interfaceToString(vName)
	if okID && vvID != "" {
		getResponse1, _, err := client.NetworkDeviceGroup.GetNetworkDeviceGroupByID(vvID)
		if err == nil && getResponse1 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return diags
		}
	}
	if okName && vvName != "" {
		getResponse2, _, err := client.NetworkDeviceGroup.GetNetworkDeviceGroupByName(replaceAllStr(vvName, "#", ":")) // WARNING: (:) colon is used as a seperator instead of (#) in the NDG name.
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return diags
		}
	}
	restyResp1, err := client.NetworkDeviceGroup.CreateNetworkDeviceGroup(request1)
	if err != nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateNetworkDeviceGroup", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateNetworkDeviceGroup", err))
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

func resourceNetworkDeviceGroupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		log.Printf("[DEBUG] Selected method: GetNetworkDeviceGroupByName")
		vvName := vName

		response1, _, err := client.NetworkDeviceGroup.GetNetworkDeviceGroupByName(replaceAllStr(vvName, "#", ":")) // WARNING: (:) colon is used as a seperator instead of (#) in the NDG name.

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetNetworkDeviceGroupByName", err,
				"Failure at GetNetworkDeviceGroupByName, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		vItemName1 := flattenNetworkDeviceGroupGetNetworkDeviceGroupByNameItemName(response1.NetworkDeviceGroup)
		if err := d.Set("item", vItemName1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNetworkDeviceGroupByName response",
				err))
			return diags
		}
		return diags

	}
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkDeviceGroupByID")
		vvID := vID

		response2, _, err := client.NetworkDeviceGroup.GetNetworkDeviceGroupByID(vvID)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetNetworkDeviceGroupByID", err,
				"Failure at GetNetworkDeviceGroupByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

		vItemID2 := flattenNetworkDeviceGroupGetNetworkDeviceGroupByIDItemID(response2.NetworkDeviceGroup)
		if err := d.Set("item", vItemID2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNetworkDeviceGroupByID response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceNetworkDeviceGroupUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		getResp, _, err := client.NetworkDeviceGroup.GetNetworkDeviceGroupByName(replaceAllStr(vvName, "#", ":")) // WARNING: (:) colon is used as a seperator instead of (#) in the NDG name.
		if err != nil || getResp == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetNetworkDeviceGroupByName", err,
				"Failure at GetNetworkDeviceGroupByName, unexpected response", ""))
			return diags
		}
		//Set value vvID = getResp.
		if getResp.NetworkDeviceGroup != nil {
			vvID = getResp.NetworkDeviceGroup.ID
		}
	}
	if d.HasChange("item") {
		log.Printf("[DEBUG] vvID %s", vvID)
		request1 := expandRequestNetworkDeviceGroupUpdateNetworkDeviceGroupByID(ctx, "item.0", d)
		log.Printf("[DEBUG] request1 => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.NetworkDeviceGroup.UpdateNetworkDeviceGroupByID(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] restyResp1 => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateNetworkDeviceGroupByID", err, restyResp1.String(),
					"Failure at UpdateNetworkDeviceGroupByID, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateNetworkDeviceGroupByID", err,
				"Failure at UpdateNetworkDeviceGroupByID, unexpected response", ""))
			return diags
		}
	}

	return resourceNetworkDeviceGroupRead(ctx, d, m)
}

func resourceNetworkDeviceGroupDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		getResp, _, err := client.NetworkDeviceGroup.GetNetworkDeviceGroupByID(vvID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	if selectedMethod == 2 {
		vvName = vName
		getResp, _, err := client.NetworkDeviceGroup.GetNetworkDeviceGroupByName(replaceAllStr(vvName, "#", ":")) // WARNING: (:) colon is used as a seperator instead of (#) in the NDG name.
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
		//Set value vvID = getResp.
		if getResp.NetworkDeviceGroup != nil {
			vvID = getResp.NetworkDeviceGroup.ID
		}
	}
	restyResp1, err := client.NetworkDeviceGroup.DeleteNetworkDeviceGroupByID(vvID)
	if err != nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] restyResp1 => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteNetworkDeviceGroupByID", err, restyResp1.String(),
				"Failure at DeleteNetworkDeviceGroupByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteNetworkDeviceGroupByID", err,
			"Failure at DeleteNetworkDeviceGroupByID, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestNetworkDeviceGroupCreateNetworkDeviceGroup(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkDeviceGroupCreateNetworkDeviceGroup {
	request := isegosdk.RequestNetworkDeviceGroupCreateNetworkDeviceGroup{}
	request.NetworkDeviceGroup = expandRequestNetworkDeviceGroupCreateNetworkDeviceGroupNetworkDeviceGroup(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkDeviceGroupCreateNetworkDeviceGroupNetworkDeviceGroup(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkDeviceGroupCreateNetworkDeviceGroupNetworkDeviceGroup {
	request := isegosdk.RequestNetworkDeviceGroupCreateNetworkDeviceGroupNetworkDeviceGroup{}
	if v, ok := d.GetOkExists(key + ".name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".name"))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".othername"); !isEmptyValue(reflect.ValueOf(d.Get(key+".othername"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".othername"))) {
		request.Othername = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkDeviceGroupUpdateNetworkDeviceGroupByID(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkDeviceGroupUpdateNetworkDeviceGroupByID {
	request := isegosdk.RequestNetworkDeviceGroupUpdateNetworkDeviceGroupByID{}
	request.NetworkDeviceGroup = expandRequestNetworkDeviceGroupUpdateNetworkDeviceGroupByIDNetworkDeviceGroup(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkDeviceGroupUpdateNetworkDeviceGroupByIDNetworkDeviceGroup(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkDeviceGroupUpdateNetworkDeviceGroupByIDNetworkDeviceGroup {
	request := isegosdk.RequestNetworkDeviceGroupUpdateNetworkDeviceGroupByIDNetworkDeviceGroup{}
	if v, ok := d.GetOkExists(key + ".id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".id"))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".name"))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".othername"); !isEmptyValue(reflect.ValueOf(d.Get(key+".othername"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".othername"))) {
		request.Othername = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
