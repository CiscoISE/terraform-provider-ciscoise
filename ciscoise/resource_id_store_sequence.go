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

func resourceIDStoreSequence() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on IdentitySequence.

- This resource allows the client to update an identity sequence.
Partial update is not supported

- This resource deletes an identity sequence.

- This resource creates an identity sequence.
`,

		CreateContext: resourceIDStoreSequenceCreate,
		ReadContext:   resourceIDStoreSequenceRead,
		UpdateContext: resourceIDStoreSequenceUpdate,
		DeleteContext: resourceIDStoreSequenceDelete,
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

						"break_on_store_fail": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"certificate_authentication_profile": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"id_seq_item": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"idstore": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"order": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
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
						"parent": &schema.Schema{
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

						"break_on_store_fail": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"certificate_authentication_profile": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id_seq_item": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"idstore": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"order": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"parent": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourceIDStoreSequenceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestIDStoreSequenceCreateIDentitySequence(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vName, okName := resourceItem["name"]
	vvName := interfaceToString(vName)
	if okID && vvID != "" {
		getResponse1, _, err := client.IDentitySequence.GetIDentitySequenceByID(vvID)
		if err == nil && getResponse1 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return resourceIDStoreSequenceRead(ctx, d, m)
		}
	}
	if okName && vvName != "" {
		getResponse2, _, err := client.IDentitySequence.GetIDentitySequenceByName(vvName)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return resourceIDStoreSequenceRead(ctx, d, m)
		}
	}
	restyResp1, err := client.IDentitySequence.CreateIDentitySequence(request1)
	if err != nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateIDentitySequence", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateIDentitySequence", err))
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
	return resourceIDStoreSequenceRead(ctx, d, m)
}

func resourceIDStoreSequenceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		log.Printf("[DEBUG] Selected method: GetIDentitySequenceByName")
		vvName := vName

		response1, restyResp1, err := client.IDentitySequence.GetIDentitySequenceByName(vvName)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItemName1 := flattenIDentitySequenceGetIDentitySequenceByNameItemName(response1.IDStoreSequence)
		if err := d.Set("item", vItemName1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetIDentitySequenceByName response",
				err))
			return diags
		}
		return diags

	}
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetIDentitySequenceByID")
		vvID := vID

		response2, restyResp2, err := client.IDentitySequence.GetIDentitySequenceByID(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItemID2 := flattenIDentitySequenceGetIDentitySequenceByIDItemID(response2.IDStoreSequence)
		if err := d.Set("item", vItemID2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetIDentitySequenceByID response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceIDStoreSequenceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		getResp, _, err := client.IDentitySequence.GetIDentitySequenceByName(vvName)
		if err != nil || getResp == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetIDentitySequenceByName", err,
				"Failure at GetIDentitySequenceByName, unexpected response", ""))
			return diags
		}
		//Set value vvID = getResp.
		if getResp.IDStoreSequence != nil {
			vvID = getResp.IDStoreSequence.ID
		}
	}
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] ID used for update operation %s", vvID)
		request1 := expandRequestIDStoreSequenceUpdateIDentitySequenceByID(ctx, "parameters.0", d)
		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
		response1, restyResp1, err := client.IDentitySequence.UpdateIDentitySequenceByID(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateIDentitySequenceByID", err, restyResp1.String(),
					"Failure at UpdateIDentitySequenceByID, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateIDentitySequenceByID", err,
				"Failure at UpdateIDentitySequenceByID, unexpected response", ""))
			return diags
		}
	}

	return resourceIDStoreSequenceRead(ctx, d, m)
}

func resourceIDStoreSequenceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		getResp, _, err := client.IDentitySequence.GetIDentitySequenceByID(vvID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	if selectedMethod == 2 {
		vvName = vName
		getResp, _, err := client.IDentitySequence.GetIDentitySequenceByName(vvName)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
		//Set value vvID = getResp.
		if getResp.IDStoreSequence != nil {
			vvID = getResp.IDStoreSequence.ID
		}
	}
	restyResp1, err := client.IDentitySequence.DeleteIDentitySequenceByID(vvID)
	if err != nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteIDentitySequenceByID", err, restyResp1.String(),
				"Failure at DeleteIDentitySequenceByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteIDentitySequenceByID", err,
			"Failure at DeleteIDentitySequenceByID, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestIDStoreSequenceCreateIDentitySequence(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestIDentitySequenceCreateIDentitySequence {
	request := isegosdk.RequestIDentitySequenceCreateIDentitySequence{}
	request.IDStoreSequence = expandRequestIDStoreSequenceCreateIDentitySequenceIDStoreSequence(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestIDStoreSequenceCreateIDentitySequenceIDStoreSequence(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestIDentitySequenceCreateIDentitySequenceIDStoreSequence {
	request := isegosdk.RequestIDentitySequenceCreateIDentitySequenceIDStoreSequence{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".parent")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".parent")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".parent")))) {
		request.Parent = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id_seq_item")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id_seq_item")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id_seq_item")))) {
		request.IDSeqItem = expandRequestIDStoreSequenceCreateIDentitySequenceIDStoreSequenceIDSeqItemArray(ctx, key+".id_seq_item", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".certificate_authentication_profile")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".certificate_authentication_profile")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".certificate_authentication_profile")))) {
		request.CertificateAuthenticationProfile = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".break_on_store_fail")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".break_on_store_fail")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".break_on_store_fail")))) {
		request.BreakOnStoreFail = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestIDStoreSequenceCreateIDentitySequenceIDStoreSequenceIDSeqItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestIDentitySequenceCreateIDentitySequenceIDStoreSequenceIDSeqItem {
	request := []isegosdk.RequestIDentitySequenceCreateIDentitySequenceIDStoreSequenceIDSeqItem{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestIDStoreSequenceCreateIDentitySequenceIDStoreSequenceIDSeqItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestIDStoreSequenceCreateIDentitySequenceIDStoreSequenceIDSeqItem(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestIDentitySequenceCreateIDentitySequenceIDStoreSequenceIDSeqItem {
	request := isegosdk.RequestIDentitySequenceCreateIDentitySequenceIDStoreSequenceIDSeqItem{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".idstore")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".idstore")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".idstore")))) {
		request.IDstore = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".order")))) {
		request.Order = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestIDStoreSequenceUpdateIDentitySequenceByID(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestIDentitySequenceUpdateIDentitySequenceByID {
	request := isegosdk.RequestIDentitySequenceUpdateIDentitySequenceByID{}
	request.IDStoreSequence = expandRequestIDStoreSequenceUpdateIDentitySequenceByIDIDStoreSequence(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestIDStoreSequenceUpdateIDentitySequenceByIDIDStoreSequence(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestIDentitySequenceUpdateIDentitySequenceByIDIDStoreSequence {
	request := isegosdk.RequestIDentitySequenceUpdateIDentitySequenceByIDIDStoreSequence{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".parent")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".parent")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".parent")))) {
		request.Parent = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id_seq_item")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id_seq_item")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id_seq_item")))) {
		request.IDSeqItem = expandRequestIDStoreSequenceUpdateIDentitySequenceByIDIDStoreSequenceIDSeqItemArray(ctx, key+".id_seq_item", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".certificate_authentication_profile")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".certificate_authentication_profile")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".certificate_authentication_profile")))) {
		request.CertificateAuthenticationProfile = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".break_on_store_fail")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".break_on_store_fail")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".break_on_store_fail")))) {
		request.BreakOnStoreFail = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestIDStoreSequenceUpdateIDentitySequenceByIDIDStoreSequenceIDSeqItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestIDentitySequenceUpdateIDentitySequenceByIDIDStoreSequenceIDSeqItem {
	request := []isegosdk.RequestIDentitySequenceUpdateIDentitySequenceByIDIDStoreSequenceIDSeqItem{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestIDStoreSequenceUpdateIDentitySequenceByIDIDStoreSequenceIDSeqItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestIDStoreSequenceUpdateIDentitySequenceByIDIDStoreSequenceIDSeqItem(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestIDentitySequenceUpdateIDentitySequenceByIDIDStoreSequenceIDSeqItem {
	request := isegosdk.RequestIDentitySequenceUpdateIDentitySequenceByIDIDStoreSequenceIDSeqItem{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".idstore")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".idstore")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".idstore")))) {
		request.IDstore = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".order")))) {
		request.Order = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
