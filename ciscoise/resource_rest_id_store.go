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

func resourceRestIDStore() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on RESTIDStore.
  
  This resource allows the client to update a REST ID store by name.
  This resource deletes a REST ID store by name.
  This resource allows the client to update a REST ID store.
  This resource deletes a REST ID store.
  This resource creates a REST ID store.`,

		CreateContext: resourceRestIDStoreCreate,
		ReadContext:   resourceRestIDStoreRead,
		UpdateContext: resourceRestIDStoreUpdate,
		DeleteContext: resourceRestIDStoreDelete,
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
						"ers_rest_idstore_attributes": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"headers": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"key": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"value": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
											},
										},
									},
									"predefined": &schema.Schema{
										Description: `The cloud provider connected to of the RestIDStore.
  Options are:
  - Azure,
  - Okta,
  - None`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"root_url": &schema.Schema{
										Description: `url of the root of the RestIDStore`,
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"username_suffix": &schema.Schema{
										Description: `Suffix of the username domain`,
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
								},
							},
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
					},
				},
			},
		},
	}
}

func resourceRestIDStoreCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("item"))
	request1 := expandRequestRestIDStoreCreateRestIDStore(ctx, "item.0", d)
	log.Printf("[DEBUG] request1 => %v", responseInterfaceToString(*request1))

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vName, okName := resourceItem["name"]
	vvName := interfaceToString(vName)
	if okID && vvID != "" {
		getResponse1, _, err := client.RestidStore.GetRestIDStoreByID(vvID)
		if err == nil && getResponse1 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return diags
		}
	}
	if okName && vvName != "" {
		getResponse2, _, err := client.RestidStore.GetRestIDStoreByName(vvName)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return diags
		}
	}
	restyResp1, err := client.RestidStore.CreateRestIDStore(request1)
	if err != nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateRestIDStore", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateRestIDStore", err))
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

func resourceRestIDStoreRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		log.Printf("[DEBUG] Selected method: GetRestIDStoreByName")
		vvName := vName

		response1, _, err := client.RestidStore.GetRestIDStoreByName(vvName)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetRestIDStoreByName", err,
				"Failure at GetRestIDStoreByName, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		vItemName1 := flattenRestidStoreGetRestIDStoreByNameItemName(response1.ERSRestIDStore)
		if err := d.Set("item", vItemName1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetRestIDStoreByName response",
				err))
			return diags
		}
		return diags

	}
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetRestIDStoreByID")
		vvID := vID

		response2, _, err := client.RestidStore.GetRestIDStoreByID(vvID)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetRestIDStoreByID", err,
				"Failure at GetRestIDStoreByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

		vItemID2 := flattenRestidStoreGetRestIDStoreByIDItemID(response2.ERSRestIDStore)
		if err := d.Set("item", vItemID2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetRestIDStoreByID response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceRestIDStoreUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		getResp, _, err := client.RestidStore.GetRestIDStoreByName(vvName)
		if err != nil || getResp == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetRestIDStoreByName", err,
				"Failure at GetRestIDStoreByName, unexpected response", ""))
			return diags
		}
		//Set value vvID = getResp.
		if getResp.ERSRestIDStore != nil {
			vvID = getResp.ERSRestIDStore.ID
		}
	}
	if d.HasChange("item") {
		log.Printf("[DEBUG] vvID %s", vvID)
		request1 := expandRequestRestIDStoreUpdateRestIDStoreByID(ctx, "item.0", d)
		log.Printf("[DEBUG] request1 => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.RestidStore.UpdateRestIDStoreByID(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] restyResp1 => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateRestIDStoreByID", err, restyResp1.String(),
					"Failure at UpdateRestIDStoreByID, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateRestIDStoreByID", err,
				"Failure at UpdateRestIDStoreByID, unexpected response", ""))
			return diags
		}
	}

	return resourceRestIDStoreRead(ctx, d, m)
}

func resourceRestIDStoreDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		getResp, _, err := client.RestidStore.GetRestIDStoreByID(vvID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	if selectedMethod == 2 {
		vvName = vName
		getResp, _, err := client.RestidStore.GetRestIDStoreByName(vvName)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
		//Set value vvID = getResp.
		if getResp.ERSRestIDStore != nil {
			vvID = getResp.ERSRestIDStore.ID
		}
	}
	restyResp1, err := client.RestidStore.DeleteRestIDStoreByID(vvID)
	if err != nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] restyResp1 => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteRestIDStoreByID", err, restyResp1.String(),
				"Failure at DeleteRestIDStoreByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteRestIDStoreByID", err,
			"Failure at DeleteRestIDStoreByID, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestRestIDStoreCreateRestIDStore(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestRestidStoreCreateRestIDStore {
	request := isegosdk.RequestRestidStoreCreateRestIDStore{}
	request.ERSRestIDStore = expandRequestRestIDStoreCreateRestIDStoreERSRestIDStore(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestRestIDStoreCreateRestIDStoreERSRestIDStore(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestRestidStoreCreateRestIDStoreERSRestIDStore {
	request := isegosdk.RequestRestidStoreCreateRestIDStoreERSRestIDStore{}
	if v, ok := d.GetOkExists(key + ".name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".name"))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".ers_rest_idstore_attributes"); !isEmptyValue(reflect.ValueOf(d.Get(key+".ers_rest_idstore_attributes"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".ers_rest_idstore_attributes"))) {
		request.ErsRestIDStoreAttributes = expandRequestRestIDStoreCreateRestIDStoreERSRestIDStoreErsRestIDStoreAttributes(ctx, key+".ers_rest_idstore_attributes.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestRestIDStoreCreateRestIDStoreERSRestIDStoreErsRestIDStoreAttributes(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestRestidStoreCreateRestIDStoreERSRestIDStoreErsRestIDStoreAttributes {
	request := isegosdk.RequestRestidStoreCreateRestIDStoreERSRestIDStoreErsRestIDStoreAttributes{}
	if v, ok := d.GetOkExists(key + ".username_suffix"); !isEmptyValue(reflect.ValueOf(d.Get(key+".username_suffix"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".username_suffix"))) {
		request.UsernameSuffix = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".root_url"); !isEmptyValue(reflect.ValueOf(d.Get(key+".root_url"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".root_url"))) {
		request.RootURL = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".predefined"); !isEmptyValue(reflect.ValueOf(d.Get(key+".predefined"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".predefined"))) {
		request.Predefined = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".headers"); !isEmptyValue(reflect.ValueOf(d.Get(key+".headers"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".headers"))) {
		request.Headers = expandRequestRestIDStoreCreateRestIDStoreERSRestIDStoreErsRestIDStoreAttributesHeadersArray(ctx, key, d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestRestIDStoreCreateRestIDStoreERSRestIDStoreErsRestIDStoreAttributesHeadersArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestRestidStoreCreateRestIDStoreERSRestIDStoreErsRestIDStoreAttributesHeaders {
	request := []isegosdk.RequestRestidStoreCreateRestIDStoreERSRestIDStoreErsRestIDStoreAttributesHeaders{}
	o := d.Get(key)
	if o != nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestRestIDStoreCreateRestIDStoreERSRestIDStoreErsRestIDStoreAttributesHeaders(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		request = append(request, *i)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestRestIDStoreCreateRestIDStoreERSRestIDStoreErsRestIDStoreAttributesHeaders(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestRestidStoreCreateRestIDStoreERSRestIDStoreErsRestIDStoreAttributesHeaders {
	request := isegosdk.RequestRestidStoreCreateRestIDStoreERSRestIDStoreErsRestIDStoreAttributesHeaders{}
	if v, ok := d.GetOkExists(key + ".key"); !isEmptyValue(reflect.ValueOf(d.Get(key+".key"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".key"))) {
		request.Key = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".value"); !isEmptyValue(reflect.ValueOf(d.Get(key+".value"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".value"))) {
		request.Value = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestRestIDStoreUpdateRestIDStoreByID(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestRestidStoreUpdateRestIDStoreByID {
	request := isegosdk.RequestRestidStoreUpdateRestIDStoreByID{}
	request.ERSRestIDStore = expandRequestRestIDStoreUpdateRestIDStoreByIDERSRestIDStore(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestRestIDStoreUpdateRestIDStoreByIDERSRestIDStore(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestRestidStoreUpdateRestIDStoreByIDERSRestIDStore {
	request := isegosdk.RequestRestidStoreUpdateRestIDStoreByIDERSRestIDStore{}
	if v, ok := d.GetOkExists(key + ".id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".id"))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".name"))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".ers_rest_idstore_attributes"); !isEmptyValue(reflect.ValueOf(d.Get(key+".ers_rest_idstore_attributes"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".ers_rest_idstore_attributes"))) {
		request.ErsRestIDStoreAttributes = expandRequestRestIDStoreUpdateRestIDStoreByIDERSRestIDStoreErsRestIDStoreAttributes(ctx, key+".ers_rest_idstore_attributes.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestRestIDStoreUpdateRestIDStoreByIDERSRestIDStoreErsRestIDStoreAttributes(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestRestidStoreUpdateRestIDStoreByIDERSRestIDStoreErsRestIDStoreAttributes {
	request := isegosdk.RequestRestidStoreUpdateRestIDStoreByIDERSRestIDStoreErsRestIDStoreAttributes{}
	if v, ok := d.GetOkExists(key + ".username_suffix"); !isEmptyValue(reflect.ValueOf(d.Get(key+".username_suffix"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".username_suffix"))) {
		request.UsernameSuffix = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".root_url"); !isEmptyValue(reflect.ValueOf(d.Get(key+".root_url"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".root_url"))) {
		request.RootURL = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".predefined"); !isEmptyValue(reflect.ValueOf(d.Get(key+".predefined"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".predefined"))) {
		request.Predefined = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".headers"); !isEmptyValue(reflect.ValueOf(d.Get(key+".headers"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".headers"))) {
		request.Headers = expandRequestRestIDStoreUpdateRestIDStoreByIDERSRestIDStoreErsRestIDStoreAttributesHeadersArray(ctx, key, d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestRestIDStoreUpdateRestIDStoreByIDERSRestIDStoreErsRestIDStoreAttributesHeadersArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestRestidStoreUpdateRestIDStoreByIDERSRestIDStoreErsRestIDStoreAttributesHeaders {
	request := []isegosdk.RequestRestidStoreUpdateRestIDStoreByIDERSRestIDStoreErsRestIDStoreAttributesHeaders{}
	o := d.Get(key)
	if o != nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestRestIDStoreUpdateRestIDStoreByIDERSRestIDStoreErsRestIDStoreAttributesHeaders(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		request = append(request, *i)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestRestIDStoreUpdateRestIDStoreByIDERSRestIDStoreErsRestIDStoreAttributesHeaders(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestRestidStoreUpdateRestIDStoreByIDERSRestIDStoreErsRestIDStoreAttributesHeaders {
	request := isegosdk.RequestRestidStoreUpdateRestIDStoreByIDERSRestIDStoreErsRestIDStoreAttributesHeaders{}
	if v, ok := d.GetOkExists(key + ".key"); !isEmptyValue(reflect.ValueOf(d.Get(key+".key"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".key"))) {
		request.Key = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".value"); !isEmptyValue(reflect.ValueOf(d.Get(key+".value"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".value"))) {
		request.Value = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
