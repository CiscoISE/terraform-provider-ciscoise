package ciscoise

import (
	"context"
	"reflect"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceExternalRadiusServer() *schema.Resource {
	return &schema.Resource{

		CreateContext: resourceExternalRadiusServerCreate,
		ReadContext:   resourceExternalRadiusServerRead,
		UpdateContext: resourceExternalRadiusServerUpdate,
		DeleteContext: resourceExternalRadiusServerDelete,
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

						"accounting_port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"authentication_port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"authenticator_key": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"enable_key_wrap": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"encryption_key": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"host_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"key_input_format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
							Optional: true,
							Computed: true,
						},
						"proxy_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"retries": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"shared_secret": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceExternalRadiusServerCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("item"))
	request1 := expandRequestExternalRadiusServerCreateExternalRadiusServer(ctx, "item.0", d)
	log.Printf("[DEBUG] request1 => %v", responseInterfaceToString(*request1))

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vName, okName := resourceItem["name"]
	vvName := interfaceToString(vName)
	if okID && vvID != "" {
		getResponse1, _, err := client.ExternalRadiusServer.GetExternalRadiusServerByID(vvID)
		if err == nil && getResponse1 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return diags
		}
	}
	if okName && vvName != "" {
		getResponse2, _, err := client.ExternalRadiusServer.GetExternalRadiusServerByName(vvName)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return diags
		}
	}
	restyResp1, err := client.ExternalRadiusServer.CreateExternalRadiusServer(request1)
	if err != nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateExternalRadiusServer", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateExternalRadiusServer", err))
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

func resourceExternalRadiusServerRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		log.Printf("[DEBUG] Selected method: GetExternalRadiusServerByName")
		vvName := vName

		response1, _, err := client.ExternalRadiusServer.GetExternalRadiusServerByName(vvName)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetExternalRadiusServerByName", err,
				"Failure at GetExternalRadiusServerByName, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		vItemName1 := flattenExternalRadiusServerGetExternalRadiusServerByNameItemName(response1.ExternalRadiusServer)
		if err := d.Set("item", vItemName1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetExternalRadiusServerByName response",
				err))
			return diags
		}
		return diags

	}
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetExternalRadiusServerByID")
		vvID := vID

		response2, _, err := client.ExternalRadiusServer.GetExternalRadiusServerByID(vvID)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetExternalRadiusServerByID", err,
				"Failure at GetExternalRadiusServerByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

		vItemID2 := flattenExternalRadiusServerGetExternalRadiusServerByIDItemID(response2.ExternalRadiusServer)
		if err := d.Set("item", vItemID2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetExternalRadiusServerByID response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceExternalRadiusServerUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		getResp, _, err := client.ExternalRadiusServer.GetExternalRadiusServerByName(vvName)
		if err != nil || getResp == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetExternalRadiusServerByName", err,
				"Failure at GetExternalRadiusServerByName, unexpected response", ""))
			return diags
		}
		//Set value vvID = getResp.
		if getResp.ExternalRadiusServer != nil {
			vvID = getResp.ExternalRadiusServer.ID
		}
	}
	if d.HasChange("item") {
		log.Printf("[DEBUG] vvID %s", vvID)
		request1 := expandRequestExternalRadiusServerUpdateExternalRadiusServerByID(ctx, "item.0", d)
		log.Printf("[DEBUG] request1 => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.ExternalRadiusServer.UpdateExternalRadiusServerByID(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] restyResp1 => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateExternalRadiusServerByID", err, restyResp1.String(),
					"Failure at UpdateExternalRadiusServerByID, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateExternalRadiusServerByID", err,
				"Failure at UpdateExternalRadiusServerByID, unexpected response", ""))
			return diags
		}
	}

	return resourceExternalRadiusServerRead(ctx, d, m)
}

func resourceExternalRadiusServerDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		getResp, _, err := client.ExternalRadiusServer.GetExternalRadiusServerByID(vvID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	if selectedMethod == 2 {
		vvName = vName
		getResp, _, err := client.ExternalRadiusServer.GetExternalRadiusServerByName(vvName)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
		//Set value vvID = getResp.
		if getResp.ExternalRadiusServer != nil {
			vvID = getResp.ExternalRadiusServer.ID
		}
	}
	restyResp1, err := client.ExternalRadiusServer.DeleteExternalRadiusServerByID(vvID)
	if err != nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] restyResp1 => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteExternalRadiusServerByID", err, restyResp1.String(),
				"Failure at DeleteExternalRadiusServerByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteExternalRadiusServerByID", err,
			"Failure at DeleteExternalRadiusServerByID, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestExternalRadiusServerCreateExternalRadiusServer(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestExternalRadiusServerCreateExternalRadiusServer {
	request := isegosdk.RequestExternalRadiusServerCreateExternalRadiusServer{}
	request.ExternalRadiusServer = expandRequestExternalRadiusServerCreateExternalRadiusServerExternalRadiusServer(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestExternalRadiusServerCreateExternalRadiusServerExternalRadiusServer(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestExternalRadiusServerCreateExternalRadiusServerExternalRadiusServer {
	request := isegosdk.RequestExternalRadiusServerCreateExternalRadiusServerExternalRadiusServer{}
	if v, ok := d.GetOkExists(key + ".name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".name"))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".host_ip"); !isEmptyValue(reflect.ValueOf(d.Get(key+".host_ip"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".host_ip"))) {
		request.HostIP = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".shared_secret"); !isEmptyValue(reflect.ValueOf(d.Get(key+".shared_secret"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".shared_secret"))) {
		request.SharedSecret = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".enable_key_wrap"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enable_key_wrap"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enable_key_wrap"))) {
		request.EnableKeyWrap = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".encryption_key"); !isEmptyValue(reflect.ValueOf(d.Get(key+".encryption_key"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".encryption_key"))) {
		request.EncryptionKey = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".authenticator_key"); !isEmptyValue(reflect.ValueOf(d.Get(key+".authenticator_key"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".authenticator_key"))) {
		request.AuthenticatorKey = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".key_input_format"); !isEmptyValue(reflect.ValueOf(d.Get(key+".key_input_format"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".key_input_format"))) {
		request.KeyInputFormat = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".authentication_port"); !isEmptyValue(reflect.ValueOf(d.Get(key+".authentication_port"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".authentication_port"))) {
		request.AuthenticationPort = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".accounting_port"); !isEmptyValue(reflect.ValueOf(d.Get(key+".accounting_port"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".accounting_port"))) {
		request.AccountingPort = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".timeout"); !isEmptyValue(reflect.ValueOf(d.Get(key+".timeout"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".timeout"))) {
		request.Timeout = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".retries"); !isEmptyValue(reflect.ValueOf(d.Get(key+".retries"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".retries"))) {
		request.Retries = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".proxy_timeout"); !isEmptyValue(reflect.ValueOf(d.Get(key+".proxy_timeout"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".proxy_timeout"))) {
		request.ProxyTimeout = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestExternalRadiusServerUpdateExternalRadiusServerByID(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestExternalRadiusServerUpdateExternalRadiusServerByID {
	request := isegosdk.RequestExternalRadiusServerUpdateExternalRadiusServerByID{}
	request.ExternalRadiusServer = expandRequestExternalRadiusServerUpdateExternalRadiusServerByIDExternalRadiusServer(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestExternalRadiusServerUpdateExternalRadiusServerByIDExternalRadiusServer(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestExternalRadiusServerUpdateExternalRadiusServerByIDExternalRadiusServer {
	request := isegosdk.RequestExternalRadiusServerUpdateExternalRadiusServerByIDExternalRadiusServer{}
	if v, ok := d.GetOkExists(key + ".id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".id"))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".name"))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".host_ip"); !isEmptyValue(reflect.ValueOf(d.Get(key+".host_ip"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".host_ip"))) {
		request.HostIP = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".shared_secret"); !isEmptyValue(reflect.ValueOf(d.Get(key+".shared_secret"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".shared_secret"))) {
		request.SharedSecret = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".enable_key_wrap"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enable_key_wrap"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enable_key_wrap"))) {
		request.EnableKeyWrap = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".encryption_key"); !isEmptyValue(reflect.ValueOf(d.Get(key+".encryption_key"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".encryption_key"))) {
		request.EncryptionKey = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".authenticator_key"); !isEmptyValue(reflect.ValueOf(d.Get(key+".authenticator_key"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".authenticator_key"))) {
		request.AuthenticatorKey = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".key_input_format"); !isEmptyValue(reflect.ValueOf(d.Get(key+".key_input_format"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".key_input_format"))) {
		request.KeyInputFormat = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".authentication_port"); !isEmptyValue(reflect.ValueOf(d.Get(key+".authentication_port"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".authentication_port"))) {
		request.AuthenticationPort = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".accounting_port"); !isEmptyValue(reflect.ValueOf(d.Get(key+".accounting_port"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".accounting_port"))) {
		request.AccountingPort = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".timeout"); !isEmptyValue(reflect.ValueOf(d.Get(key+".timeout"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".timeout"))) {
		request.Timeout = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".retries"); !isEmptyValue(reflect.ValueOf(d.Get(key+".retries"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".retries"))) {
		request.Retries = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".proxy_timeout"); !isEmptyValue(reflect.ValueOf(d.Get(key+".proxy_timeout"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".proxy_timeout"))) {
		request.ProxyTimeout = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}