package ciscoise

import (
	"context"
	"reflect"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTacacsServerSequence() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on TacacsServerSequence.

- This resource allows the client to update a TACACS server sequence.

- This resource deletes a TACACS server sequence.

- This resource creates a TACACS server sequence.
`,

		CreateContext: resourceTacacsServerSequenceCreate,
		ReadContext:   resourceTacacsServerSequenceRead,
		UpdateContext: resourceTacacsServerSequenceUpdate,
		DeleteContext: resourceTacacsServerSequenceDelete,
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
						"local_accounting": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"prefix_delimiter": &schema.Schema{
							Description: `The delimiter that will be used for prefix strip`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"prefix_strip": &schema.Schema{
							Description: `Define if a delimiter will be used for prefix strip`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"remote_accounting": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"server_list": &schema.Schema{
							Description: `The names of Tacacs external servers separated by commas.
The order of the names in the string is the order of servers that will be used during authentication`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"suffix_delimiter": &schema.Schema{
							Description: `The delimiter that will be used for suffix strip`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"suffix_strip": &schema.Schema{
							Description: `Define if a delimiter will be used for suffix strip`,
							Type:        schema.TypeString,
							Computed:    true,
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
						"local_accounting": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"prefix_delimiter": &schema.Schema{
							Description: `The delimiter that will be used for prefix strip`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"prefix_strip": &schema.Schema{
							Description:  `Define if a delimiter will be used for prefix strip`,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"remote_accounting": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"server_list": &schema.Schema{
							Description: `The names of Tacacs external servers separated by commas.
The order of the names in the string is the order of servers that will be used during authentication`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"suffix_delimiter": &schema.Schema{
							Description: `The delimiter that will be used for suffix strip`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"suffix_strip": &schema.Schema{
							Description:  `Define if a delimiter will be used for suffix strip`,
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

func resourceTacacsServerSequenceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestTacacsServerSequenceCreateTacacsServerSequence(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vName, okName := resourceItem["name"]
	vvName := interfaceToString(vName)
	if okID && vvID != "" {
		getResponse1, _, err := client.TacacsServerSequence.GetTacacsServerSequenceByID(vvID)
		if err == nil && getResponse1 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return resourceTacacsServerSequenceRead(ctx, d, m)
		}
	}
	if okName && vvName != "" {
		getResponse2, _, err := client.TacacsServerSequence.GetTacacsServerSequenceByName(vvName)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return resourceTacacsServerSequenceRead(ctx, d, m)
		}
	}
	restyResp1, err := client.TacacsServerSequence.CreateTacacsServerSequence(request1)
	if err != nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateTacacsServerSequence", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateTacacsServerSequence", err))
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
	return resourceTacacsServerSequenceRead(ctx, d, m)
}

func resourceTacacsServerSequenceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		log.Printf("[DEBUG] Selected method: GetTacacsServerSequenceByName")
		vvName := vName

		response1, restyResp1, err := client.TacacsServerSequence.GetTacacsServerSequenceByName(vvName)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItemName1 := flattenTacacsServerSequenceGetTacacsServerSequenceByNameItemName(response1.TacacsServerSequence)
		if err := d.Set("item", vItemName1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTacacsServerSequenceByName response",
				err))
			return diags
		}
		return diags

	}
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetTacacsServerSequenceByID")
		vvID := vID

		response2, restyResp2, err := client.TacacsServerSequence.GetTacacsServerSequenceByID(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItemID2 := flattenTacacsServerSequenceGetTacacsServerSequenceByIDItemID(response2.TacacsServerSequence)
		if err := d.Set("item", vItemID2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTacacsServerSequenceByID response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceTacacsServerSequenceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		getResp, _, err := client.TacacsServerSequence.GetTacacsServerSequenceByName(vvName)
		if err != nil || getResp == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetTacacsServerSequenceByName", err,
				"Failure at GetTacacsServerSequenceByName, unexpected response", ""))
			return diags
		}
		//Set value vvID = getResp.
		if getResp.TacacsServerSequence != nil {
			vvID = getResp.TacacsServerSequence.ID
		}
	}
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] ID used for update operation %s", vvID)
		request1 := expandRequestTacacsServerSequenceUpdateTacacsServerSequenceByID(ctx, "parameters.0", d)
		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
		response1, restyResp1, err := client.TacacsServerSequence.UpdateTacacsServerSequenceByID(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateTacacsServerSequenceByID", err, restyResp1.String(),
					"Failure at UpdateTacacsServerSequenceByID, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateTacacsServerSequenceByID", err,
				"Failure at UpdateTacacsServerSequenceByID, unexpected response", ""))
			return diags
		}
	}

	return resourceTacacsServerSequenceRead(ctx, d, m)
}

func resourceTacacsServerSequenceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		getResp, _, err := client.TacacsServerSequence.GetTacacsServerSequenceByID(vvID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	if selectedMethod == 2 {
		vvName = vName
		getResp, _, err := client.TacacsServerSequence.GetTacacsServerSequenceByName(vvName)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
		//Set value vvID = getResp.
		if getResp.TacacsServerSequence != nil {
			vvID = getResp.TacacsServerSequence.ID
		}
	}
	restyResp1, err := client.TacacsServerSequence.DeleteTacacsServerSequenceByID(vvID)
	if err != nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteTacacsServerSequenceByID", err, restyResp1.String(),
				"Failure at DeleteTacacsServerSequenceByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteTacacsServerSequenceByID", err,
			"Failure at DeleteTacacsServerSequenceByID, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestTacacsServerSequenceCreateTacacsServerSequence(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestTacacsServerSequenceCreateTacacsServerSequence {
	request := isegosdk.RequestTacacsServerSequenceCreateTacacsServerSequence{}
	request.TacacsServerSequence = expandRequestTacacsServerSequenceCreateTacacsServerSequenceTacacsServerSequence(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestTacacsServerSequenceCreateTacacsServerSequenceTacacsServerSequence(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestTacacsServerSequenceCreateTacacsServerSequenceTacacsServerSequence {
	request := isegosdk.RequestTacacsServerSequenceCreateTacacsServerSequenceTacacsServerSequence{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".server_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".server_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".server_list")))) {
		request.ServerList = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".local_accounting")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".local_accounting")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".local_accounting")))) {
		request.LocalAccounting = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".remote_accounting")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".remote_accounting")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".remote_accounting")))) {
		request.RemoteAccounting = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".prefix_strip")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".prefix_strip")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".prefix_strip")))) {
		request.PrefixStrip = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".prefix_delimiter")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".prefix_delimiter")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".prefix_delimiter")))) {
		request.PrefixDelimiter = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".suffix_strip")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".suffix_strip")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".suffix_strip")))) {
		request.SuffixStrip = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".suffix_delimiter")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".suffix_delimiter")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".suffix_delimiter")))) {
		request.SuffixDelimiter = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestTacacsServerSequenceUpdateTacacsServerSequenceByID(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestTacacsServerSequenceUpdateTacacsServerSequenceByID {
	request := isegosdk.RequestTacacsServerSequenceUpdateTacacsServerSequenceByID{}
	request.TacacsServerSequence = expandRequestTacacsServerSequenceUpdateTacacsServerSequenceByIDTacacsServerSequence(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestTacacsServerSequenceUpdateTacacsServerSequenceByIDTacacsServerSequence(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestTacacsServerSequenceUpdateTacacsServerSequenceByIDTacacsServerSequence {
	request := isegosdk.RequestTacacsServerSequenceUpdateTacacsServerSequenceByIDTacacsServerSequence{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".server_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".server_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".server_list")))) {
		request.ServerList = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".local_accounting")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".local_accounting")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".local_accounting")))) {
		request.LocalAccounting = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".remote_accounting")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".remote_accounting")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".remote_accounting")))) {
		request.RemoteAccounting = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".prefix_strip")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".prefix_strip")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".prefix_strip")))) {
		request.PrefixStrip = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".prefix_delimiter")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".prefix_delimiter")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".prefix_delimiter")))) {
		request.PrefixDelimiter = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".suffix_strip")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".suffix_strip")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".suffix_strip")))) {
		request.SuffixStrip = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".suffix_delimiter")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".suffix_delimiter")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".suffix_delimiter")))) {
		request.SuffixDelimiter = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
