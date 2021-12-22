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

func resourceTacacsProfile() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on TACACSProfile.

- This resource allows the client to update a TACACS profile.

- This resource deletes a TACACS profile.

- This resource creates a TACACS profile.
`,

		CreateContext: resourceTacacsProfileCreate,
		ReadContext:   resourceTacacsProfileRead,
		UpdateContext: resourceTacacsProfileUpdate,
		DeleteContext: resourceTacacsProfileDelete,
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

						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
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
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"session_attributes": &schema.Schema{
							Description: `Holds list of session attributes. View type for GUI is Shell by default`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"session_attribute_list": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"type": &schema.Schema{
													Description: `Allowed values: MANDATORY, OPTIONAL`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"value": &schema.Schema{
													Type:     schema.TypeString,
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
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"session_attributes": &schema.Schema{
							Description: `Holds list of session attributes. View type for GUI is Shell by default`,
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"session_attribute_list": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"name": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"type": &schema.Schema{
													Description: `Allowed values: MANDATORY, OPTIONAL`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"value": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
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

func resourceTacacsProfileCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestTacacsProfileCreateTacacsProfile(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vName, okName := resourceItem["name"]
	vvName := interfaceToString(vName)
	if okID && vvID != "" {
		getResponse1, _, err := client.TacacsProfile.GetTacacsProfileByID(vvID)
		if err == nil && getResponse1 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return resourceTacacsProfileRead(ctx, d, m)
		}
	}
	if okName && vvName != "" {
		getResponse2, _, err := client.TacacsProfile.GetTacacsProfileByName(vvName)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return resourceTacacsProfileRead(ctx, d, m)
		}
	}
	restyResp1, err := client.TacacsProfile.CreateTacacsProfile(request1)
	if err != nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateTacacsProfile", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateTacacsProfile", err))
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
	return resourceTacacsProfileRead(ctx, d, m)
}

func resourceTacacsProfileRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		log.Printf("[DEBUG] Selected method: GetTacacsProfileByName")
		vvName := vName

		response1, restyResp1, err := client.TacacsProfile.GetTacacsProfileByName(vvName)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetTacacsProfileByName", err,
				"Failure at GetTacacsProfileByName, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItemName1 := flattenTacacsProfileGetTacacsProfileByNameItemName(response1.TacacsProfile)
		if err := d.Set("item", vItemName1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTacacsProfileByName response",
				err))
			return diags
		}
		return diags

	}
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetTacacsProfileByID")
		vvID := vID

		response2, restyResp2, err := client.TacacsProfile.GetTacacsProfileByID(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetTacacsProfileByID", err,
				"Failure at GetTacacsProfileByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItemID2 := flattenTacacsProfileGetTacacsProfileByIDItemID(response2.TacacsProfile)
		if err := d.Set("item", vItemID2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTacacsProfileByID response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceTacacsProfileUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		getResp, _, err := client.TacacsProfile.GetTacacsProfileByName(vvName)
		if err != nil || getResp == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetTacacsProfileByName", err,
				"Failure at GetTacacsProfileByName, unexpected response", ""))
			return diags
		}
		//Set value vvID = getResp.
		if getResp.TacacsProfile != nil {
			vvID = getResp.TacacsProfile.ID
		}
	}
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] ID used for update operation %s", vvID)
		request1 := expandRequestTacacsProfileUpdateTacacsProfileByID(ctx, "parameters.0", d)
		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
		response1, restyResp1, err := client.TacacsProfile.UpdateTacacsProfileByID(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateTacacsProfileByID", err, restyResp1.String(),
					"Failure at UpdateTacacsProfileByID, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateTacacsProfileByID", err,
				"Failure at UpdateTacacsProfileByID, unexpected response", ""))
			return diags
		}
	}

	return resourceTacacsProfileRead(ctx, d, m)
}

func resourceTacacsProfileDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		getResp, _, err := client.TacacsProfile.GetTacacsProfileByID(vvID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	if selectedMethod == 2 {
		vvName = vName
		getResp, _, err := client.TacacsProfile.GetTacacsProfileByName(vvName)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
		//Set value vvID = getResp.
		if getResp.TacacsProfile != nil {
			vvID = getResp.TacacsProfile.ID
		}
	}
	restyResp1, err := client.TacacsProfile.DeleteTacacsProfileByID(vvID)
	if err != nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteTacacsProfileByID", err, restyResp1.String(),
				"Failure at DeleteTacacsProfileByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteTacacsProfileByID", err,
			"Failure at DeleteTacacsProfileByID, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestTacacsProfileCreateTacacsProfile(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestTacacsProfileCreateTacacsProfile {
	request := isegosdk.RequestTacacsProfileCreateTacacsProfile{}
	request.TacacsProfile = expandRequestTacacsProfileCreateTacacsProfileTacacsProfile(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestTacacsProfileCreateTacacsProfileTacacsProfile(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestTacacsProfileCreateTacacsProfileTacacsProfile {
	request := isegosdk.RequestTacacsProfileCreateTacacsProfileTacacsProfile{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".session_attributes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".session_attributes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".session_attributes")))) {
		request.SessionAttributes = expandRequestTacacsProfileCreateTacacsProfileTacacsProfileSessionAttributes(ctx, key+".session_attributes.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestTacacsProfileCreateTacacsProfileTacacsProfileSessionAttributes(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestTacacsProfileCreateTacacsProfileTacacsProfileSessionAttributes {
	request := isegosdk.RequestTacacsProfileCreateTacacsProfileTacacsProfileSessionAttributes{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".session_attribute_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".session_attribute_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".session_attribute_list")))) {
		request.SessionAttributeList = expandRequestTacacsProfileCreateTacacsProfileTacacsProfileSessionAttributesSessionAttributeListArray(ctx, key+".session_attribute_list", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestTacacsProfileCreateTacacsProfileTacacsProfileSessionAttributesSessionAttributeListArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestTacacsProfileCreateTacacsProfileTacacsProfileSessionAttributesSessionAttributeList {
	request := []isegosdk.RequestTacacsProfileCreateTacacsProfileTacacsProfileSessionAttributesSessionAttributeList{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestTacacsProfileCreateTacacsProfileTacacsProfileSessionAttributesSessionAttributeList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestTacacsProfileCreateTacacsProfileTacacsProfileSessionAttributesSessionAttributeList(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestTacacsProfileCreateTacacsProfileTacacsProfileSessionAttributesSessionAttributeList {
	request := isegosdk.RequestTacacsProfileCreateTacacsProfileTacacsProfileSessionAttributesSessionAttributeList{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".value")))) {
		request.Value = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestTacacsProfileUpdateTacacsProfileByID(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestTacacsProfileUpdateTacacsProfileByID {
	request := isegosdk.RequestTacacsProfileUpdateTacacsProfileByID{}
	request.TacacsProfile = expandRequestTacacsProfileUpdateTacacsProfileByIDTacacsProfile(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestTacacsProfileUpdateTacacsProfileByIDTacacsProfile(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestTacacsProfileUpdateTacacsProfileByIDTacacsProfile {
	request := isegosdk.RequestTacacsProfileUpdateTacacsProfileByIDTacacsProfile{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".session_attributes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".session_attributes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".session_attributes")))) {
		request.SessionAttributes = expandRequestTacacsProfileUpdateTacacsProfileByIDTacacsProfileSessionAttributes(ctx, key+".session_attributes.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestTacacsProfileUpdateTacacsProfileByIDTacacsProfileSessionAttributes(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestTacacsProfileUpdateTacacsProfileByIDTacacsProfileSessionAttributes {
	request := isegosdk.RequestTacacsProfileUpdateTacacsProfileByIDTacacsProfileSessionAttributes{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".session_attribute_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".session_attribute_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".session_attribute_list")))) {
		request.SessionAttributeList = expandRequestTacacsProfileUpdateTacacsProfileByIDTacacsProfileSessionAttributesSessionAttributeListArray(ctx, key+".session_attribute_list", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestTacacsProfileUpdateTacacsProfileByIDTacacsProfileSessionAttributesSessionAttributeListArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestTacacsProfileUpdateTacacsProfileByIDTacacsProfileSessionAttributesSessionAttributeList {
	request := []isegosdk.RequestTacacsProfileUpdateTacacsProfileByIDTacacsProfileSessionAttributesSessionAttributeList{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestTacacsProfileUpdateTacacsProfileByIDTacacsProfileSessionAttributesSessionAttributeList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestTacacsProfileUpdateTacacsProfileByIDTacacsProfileSessionAttributesSessionAttributeList(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestTacacsProfileUpdateTacacsProfileByIDTacacsProfileSessionAttributesSessionAttributeList {
	request := isegosdk.RequestTacacsProfileUpdateTacacsProfileByIDTacacsProfileSessionAttributesSessionAttributeList{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".value")))) {
		request.Value = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
