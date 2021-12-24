package ciscoise

import (
	"context"
	"reflect"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEndpoint() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on endpoint.

- This resource allows the client to update an endpoint.

- This resource deletes an endpoint.

- This resource creates an endpoint.
`,

		CreateContext: resourceEndpointCreate,
		ReadContext:   resourceEndpointRead,
		UpdateContext: resourceEndpointUpdate,
		DeleteContext: resourceEndpointDelete,
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

						"custom_attributes": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"custom_attributes": &schema.Schema{
										Description: `Key value map`,
										Type:        schema.TypeMap,
										Computed:    true,
									},
								},
							},
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"group_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"identity_store": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"identity_store_id": &schema.Schema{
							Type:     schema.TypeString,
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
						"mac": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"mdm_attributes": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"mdm_compliance_status": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"mdm_encrypted": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"mdm_enrolled": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"mdm_ime_i": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"mdm_jail_broken": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"mdm_manufacturer": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"mdm_model": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"mdm_os": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"mdm_phone_number": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"mdm_pinlock": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"mdm_reachable": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"mdm_serial": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"mdm_server_name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"portal_user": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"profile_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"static_group_assignment": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"static_profile_assignment": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"custom_attributes": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"custom_attributes": &schema.Schema{
										Description: `Key value map`,
										Type:        schema.TypeMap,
										Optional:    true,
									},
								},
							},
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"group_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"identity_store": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"identity_store_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"mac": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"mdm_attributes": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"mdm_compliance_status": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
									},
									"mdm_encrypted": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
									},
									"mdm_enrolled": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
									},
									"mdm_ime_i": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"mdm_jail_broken": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
									},
									"mdm_manufacturer": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"mdm_model": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"mdm_os": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"mdm_phone_number": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"mdm_pinlock": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
									},
									"mdm_reachable": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
									},
									"mdm_serial": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"mdm_server_name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"portal_user": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"profile_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"static_group_assignment": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"static_profile_assignment": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
					},
				},
			},
		},
	}
}

func resourceEndpointCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning Endpoint Create")
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestEndpointCreateEndpoint(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vName, okName := resourceItem["name"]
	vvName := interfaceToString(vName)
	if okID && vvID != "" {
		getResponse1, _, err := client.Endpoint.GetEndpointByID(vvID)
		if err == nil && getResponse1 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return resourceEndpointRead(ctx, d, m)
		}
	}
	if okName && vvName != "" {
		getResponse2, _, err := client.Endpoint.GetEndpointByName(vvName)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return resourceEndpointRead(ctx, d, m)
		}
	}
	restyResp1, err := client.Endpoint.CreateEndpoint(request1)
	if err != nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateEndpoint", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateEndpoint", err))
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
	return resourceEndpointRead(ctx, d, m)
}

func resourceEndpointRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning Endpoint Read for id=[%s]", d.Id())
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
		log.Printf("[DEBUG] Selected method: GetEndpointByName")
		vvName := vName

		response1, restyResp1, err := client.Endpoint.GetEndpointByName(vvName)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItemName1 := flattenEndpointGetEndpointByNameItemName(response1.ERSEndPoint)
		if err := d.Set("item", vItemName1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetEndpointByName response",
				err))
			return diags
		}
		return diags

	}
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetEndpointByID")
		vvID := vID

		response2, restyResp2, err := client.Endpoint.GetEndpointByID(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItemID2 := flattenEndpointGetEndpointByIDItemID(response2.ERSEndPoint)
		if err := d.Set("item", vItemID2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetEndpointByID response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceEndpointUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning Endpoint Update for id=[%s]", d.Id())
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
		getResp, _, err := client.Endpoint.GetEndpointByName(vvName)
		if err != nil || getResp == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetEndpointByName", err,
				"Failure at GetEndpointByName, unexpected response", ""))
			return diags
		}
		//Set value vvID = getResp.
		if getResp.ERSEndPoint != nil {
			vvID = getResp.ERSEndPoint.ID
		}
	}
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] ID used for update operation %s", vvID)
		request1 := expandRequestEndpointUpdateEndpointByID(ctx, "parameters.0", d)
		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
		response1, restyResp1, err := client.Endpoint.UpdateEndpointByID(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateEndpointByID", err, restyResp1.String(),
					"Failure at UpdateEndpointByID, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateEndpointByID", err,
				"Failure at UpdateEndpointByID, unexpected response", ""))
			return diags
		}
	}

	return resourceEndpointRead(ctx, d, m)
}

func resourceEndpointDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning Endpoint Delete for id=[%s]", d.Id())
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
		getResp, _, err := client.Endpoint.GetEndpointByID(vvID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	if selectedMethod == 2 {
		vvName = vName
		getResp, _, err := client.Endpoint.GetEndpointByName(vvName)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
		//Set value vvID = getResp.
		if getResp.ERSEndPoint != nil {
			vvID = getResp.ERSEndPoint.ID
		}
	}
	restyResp1, err := client.Endpoint.DeleteEndpointByID(vvID)
	if err != nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteEndpointByID", err, restyResp1.String(),
				"Failure at DeleteEndpointByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteEndpointByID", err,
			"Failure at DeleteEndpointByID, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestEndpointCreateEndpoint(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestEndpointCreateEndpoint {
	request := isegosdk.RequestEndpointCreateEndpoint{}
	request.ERSEndPoint = expandRequestEndpointCreateEndpointERSEndPoint(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEndpointCreateEndpointERSEndPoint(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestEndpointCreateEndpointERSEndPoint {
	request := isegosdk.RequestEndpointCreateEndpointERSEndPoint{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mac")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mac")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mac")))) {
		request.Mac = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".profile_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".profile_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".profile_id")))) {
		request.ProfileID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".static_profile_assignment")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".static_profile_assignment")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".static_profile_assignment")))) {
		request.StaticProfileAssignment = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".group_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".group_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".group_id")))) {
		request.GroupID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".static_group_assignment")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".static_group_assignment")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".static_group_assignment")))) {
		request.StaticGroupAssignment = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".portal_user")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".portal_user")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".portal_user")))) {
		request.PortalUser = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".identity_store")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".identity_store")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".identity_store")))) {
		request.IDentityStore = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".identity_store_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".identity_store_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".identity_store_id")))) {
		request.IDentityStoreID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mdm_attributes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mdm_attributes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mdm_attributes")))) {
		request.MdmAttributes = expandRequestEndpointCreateEndpointERSEndPointMdmAttributes(ctx, key+".mdm_attributes.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".custom_attributes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".custom_attributes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".custom_attributes")))) {
		request.CustomAttributes = expandRequestEndpointCreateEndpointERSEndPointCustomAttributes(ctx, key+".custom_attributes", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEndpointCreateEndpointERSEndPointMdmAttributes(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestEndpointCreateEndpointERSEndPointMdmAttributes {
	request := isegosdk.RequestEndpointCreateEndpointERSEndPointMdmAttributes{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mdm_server_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mdm_server_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mdm_server_name")))) {
		request.MdmServerName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mdm_reachable")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mdm_reachable")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mdm_reachable")))) {
		request.MdmReachable = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mdm_enrolled")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mdm_enrolled")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mdm_enrolled")))) {
		request.MdmEnrolled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mdm_compliance_status")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mdm_compliance_status")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mdm_compliance_status")))) {
		request.MdmComplianceStatus = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mdm_os")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mdm_os")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mdm_os")))) {
		request.MdmOS = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mdm_manufacturer")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mdm_manufacturer")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mdm_manufacturer")))) {
		request.MdmManufacturer = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mdm_model")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mdm_model")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mdm_model")))) {
		request.MdmModel = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mdm_serial")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mdm_serial")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mdm_serial")))) {
		request.MdmSerial = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mdm_encrypted")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mdm_encrypted")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mdm_encrypted")))) {
		request.MdmEncrypted = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mdm_pinlock")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mdm_pinlock")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mdm_pinlock")))) {
		request.MdmPinlock = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mdm_jail_broken")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mdm_jail_broken")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mdm_jail_broken")))) {
		request.MdmJailBroken = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mdm_ime_i")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mdm_ime_i")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mdm_ime_i")))) {
		request.MdmIMEI = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mdm_phone_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mdm_phone_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mdm_phone_number")))) {
		request.MdmPhoneNumber = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEndpointCreateEndpointERSEndPointCustomAttributes(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestEndpointCreateEndpointERSEndPointCustomAttributes {
	request := isegosdk.RequestEndpointCreateEndpointERSEndPointCustomAttributes{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".custom_attributes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".custom_attributes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".custom_attributes")))) {
		request.CustomAttributes = expandRequestEndpointCreateEndpointERSEndPointCustomAttributesCustomAttributes(ctx, key+".custom_attributes", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEndpointCreateEndpointERSEndPointCustomAttributesCustomAttributes(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestEndpointCreateEndpointERSEndPointCustomAttributesCustomAttributes {
	var request isegosdk.RequestEndpointCreateEndpointERSEndPointCustomAttributesCustomAttributes
	v := d.Get(fixKeyAccess(key))
	log.Printf("[DEBUG] customAttributes => %s => %v", key, v)
	request = v.(map[string]interface{})
	return &request
}

func expandRequestEndpointUpdateEndpointByID(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestEndpointUpdateEndpointByID {
	request := isegosdk.RequestEndpointUpdateEndpointByID{}
	request.ERSEndPoint = expandRequestEndpointUpdateEndpointByIDERSEndPoint(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEndpointUpdateEndpointByIDERSEndPoint(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestEndpointUpdateEndpointByIDERSEndPoint {
	request := isegosdk.RequestEndpointUpdateEndpointByIDERSEndPoint{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mac")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mac")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mac")))) {
		request.Mac = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".profile_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".profile_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".profile_id")))) {
		request.ProfileID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".static_profile_assignment")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".static_profile_assignment")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".static_profile_assignment")))) {
		request.StaticProfileAssignment = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".group_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".group_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".group_id")))) {
		request.GroupID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".static_group_assignment")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".static_group_assignment")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".static_group_assignment")))) {
		request.StaticGroupAssignment = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".portal_user")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".portal_user")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".portal_user")))) {
		request.PortalUser = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".identity_store")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".identity_store")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".identity_store")))) {
		request.IDentityStore = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".identity_store_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".identity_store_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".identity_store_id")))) {
		request.IDentityStoreID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mdm_attributes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mdm_attributes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mdm_attributes")))) {
		request.MdmAttributes = expandRequestEndpointUpdateEndpointByIDERSEndPointMdmAttributes(ctx, key+".mdm_attributes.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".custom_attributes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".custom_attributes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".custom_attributes")))) {
		request.CustomAttributes = expandRequestEndpointUpdateEndpointByIDERSEndPointCustomAttributes(ctx, key+".custom_attributes", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEndpointUpdateEndpointByIDERSEndPointMdmAttributes(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestEndpointUpdateEndpointByIDERSEndPointMdmAttributes {
	request := isegosdk.RequestEndpointUpdateEndpointByIDERSEndPointMdmAttributes{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mdm_server_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mdm_server_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mdm_server_name")))) {
		request.MdmServerName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mdm_reachable")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mdm_reachable")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mdm_reachable")))) {
		request.MdmReachable = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mdm_enrolled")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mdm_enrolled")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mdm_enrolled")))) {
		request.MdmEnrolled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mdm_compliance_status")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mdm_compliance_status")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mdm_compliance_status")))) {
		request.MdmComplianceStatus = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mdm_os")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mdm_os")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mdm_os")))) {
		request.MdmOS = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mdm_manufacturer")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mdm_manufacturer")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mdm_manufacturer")))) {
		request.MdmManufacturer = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mdm_model")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mdm_model")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mdm_model")))) {
		request.MdmModel = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mdm_serial")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mdm_serial")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mdm_serial")))) {
		request.MdmSerial = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mdm_encrypted")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mdm_encrypted")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mdm_encrypted")))) {
		request.MdmEncrypted = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mdm_pinlock")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mdm_pinlock")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mdm_pinlock")))) {
		request.MdmPinlock = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mdm_jail_broken")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mdm_jail_broken")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mdm_jail_broken")))) {
		request.MdmJailBroken = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mdm_ime_i")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mdm_ime_i")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mdm_ime_i")))) {
		request.MdmIMEI = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mdm_phone_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mdm_phone_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mdm_phone_number")))) {
		request.MdmPhoneNumber = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEndpointUpdateEndpointByIDERSEndPointCustomAttributes(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestEndpointUpdateEndpointByIDERSEndPointCustomAttributes {
	request := isegosdk.RequestEndpointUpdateEndpointByIDERSEndPointCustomAttributes{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".custom_attributes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".custom_attributes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".custom_attributes")))) {
		request.CustomAttributes = expandRequestEndpointUpdateEndpointByIDERSEndPointCustomAttributesCustomAttributes(ctx, key+".custom_attributes", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEndpointUpdateEndpointByIDERSEndPointCustomAttributesCustomAttributes(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestEndpointUpdateEndpointByIDERSEndPointCustomAttributesCustomAttributes {
	var request isegosdk.RequestEndpointUpdateEndpointByIDERSEndPointCustomAttributesCustomAttributes
	v := d.Get(fixKeyAccess(key))
	request = v.(map[string]interface{})
	return &request
}
