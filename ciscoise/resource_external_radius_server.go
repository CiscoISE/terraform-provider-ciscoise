package ciscoise

import (
	"context"
	"reflect"

	"log"

	isegosdk "github.com/kuba-mazurkiewicz/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceExternalRadiusServer() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on ExternalRADIUSServer.

- This resource allows the client to update an external RADIUS server.

- This resource deletes an external RADIUS server.

- This resource creates an external RADIUS server.
`,

		CreateContext: resourceExternalRadiusServerCreate,
		ReadContext:   resourceExternalRadiusServerRead,
		UpdateContext: resourceExternalRadiusServerUpdate,
		DeleteContext: resourceExternalRadiusServerDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Description: `Unix timestamp records the last time that the resource was updated.`,
				Type:        schema.TypeString,
				Computed:    true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"accounting_port": &schema.Schema{
							Description: `Valid Range 1 to 65535`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"authentication_port": &schema.Schema{
							Description: `Valid Range 1 to 65535`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"authenticator_key": &schema.Schema{
							Description: `The authenticatorKey is required only if enableKeyWrap is true, otherwise it must be ignored or empty.
The maximum length is 20 ASCII characters or 40 HEXADECIMAL characters (depend on selection in field 'keyInputFormat')`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"enable_key_wrap": &schema.Schema{
							Description: `KeyWrap may only be enabled if it is supported on the device.
When running in FIPS mode this option should be enabled for such devices`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"encryption_key": &schema.Schema{
							Description: `The encryptionKey is required only if enableKeyWrap is true, otherwise it must be ignored or empty.
The maximum length is 16 ASCII characters or 32 HEXADECIMAL characters (depend on selection in field 'keyInputFormat')`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"host_ip": &schema.Schema{
							Description: `The IP of the host - must be a valid IPV4 address`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"key_input_format": &schema.Schema{
							Description: `Specifies the format of the input for fields 'encryptionKey' and 'authenticatorKey'.
Allowed Values:
- ASCII
- HEXADECIMAL`,
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
							Description: `Resource Name. Allowed charactera are alphanumeric and _ (underscore).`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"proxy_timeout": &schema.Schema{
							Description: `Valid Range 1 to 600`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"retries": &schema.Schema{
							Description: `Valid Range 1 to 9`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"shared_secret": &schema.Schema{
							Description: `Shared secret maximum length is 128 characters`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"timeout": &schema.Schema{
							Description: `Valid Range 1 to 120`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"accounting_port": &schema.Schema{
							Description:      `Valid Range 1 to 65535`,
							Type:             schema.TypeInt,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"authentication_port": &schema.Schema{
							Description:      `Valid Range 1 to 65535`,
							Type:             schema.TypeInt,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"authenticator_key": &schema.Schema{
							Description: `The authenticatorKey is required only if enableKeyWrap is true, otherwise it must be ignored or empty.
		The maximum length is 20 ASCII characters or 40 HEXADECIMAL characters (depend on selection in field 'keyInputFormat')`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"description": &schema.Schema{
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"enable_key_wrap": &schema.Schema{
							Description: `KeyWrap may only be enabled if it is supported on the device.
		When running in FIPS mode this option should be enabled for such devices`,
							Type:             schema.TypeString,
							ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:         true,
							DiffSuppressFunc: diffSupressBool(),
							Computed:         true,
						},
						"encryption_key": &schema.Schema{
							Description: `The encryptionKey is required only if enableKeyWrap is true, otherwise it must be ignored or empty.
		The maximum length is 16 ASCII characters or 32 HEXADECIMAL characters (depend on selection in field 'keyInputFormat')`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"host_ip": &schema.Schema{
							Description:      `The IP of the host - must be a valid IPV4 address`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"id": &schema.Schema{
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"key_input_format": &schema.Schema{
							Description: `Specifies the format of the input for fields 'encryptionKey' and 'authenticatorKey'.
		Allowed Values:
		- ASCII
		- HEXADECIMAL`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
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
							Description:      `Resource Name. Allowed charactera are alphanumeric and _ (underscore).`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"proxy_timeout": &schema.Schema{
							Description:      `Valid Range 1 to 600`,
							Type:             schema.TypeInt,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"retries": &schema.Schema{
							Description:      `Valid Range 1 to 9`,
							Type:             schema.TypeInt,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"shared_secret": &schema.Schema{
							Description:      `Shared secret maximum length is 128 characters`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"timeout": &schema.Schema{
							Description:      `Valid Range 1 to 120`,
							Type:             schema.TypeInt,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
					},
				},
			},
		},
	}
}

func resourceExternalRadiusServerCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning ExternalRadiusServer create")
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client
	isEnableAutoImport := clientConfig.EnableAutoImport

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestExternalRadiusServerCreateExternalRadiusServer(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vName, okName := resourceItem["name"]
	vvName := interfaceToString(vName)
	if isEnableAutoImport {
		if okID && vvID != "" {
			getResponse1, _, err := client.ExternalRadiusServer.GetExternalRadiusServerByID(vvID)
			if err == nil && getResponse1 != nil {
				resourceMap := make(map[string]string)
				resourceMap["id"] = vvID
				resourceMap["name"] = vvName
				d.SetId(joinResourceID(resourceMap))
				return resourceExternalRadiusServerRead(ctx, d, m)
			}
		}
		if okName && vvName != "" {
			getResponse2, _, err := client.ExternalRadiusServer.GetExternalRadiusServerByName(vvName)
			if err == nil && getResponse2 != nil {
				resourceMap := make(map[string]string)
				resourceMap["id"] = getResponse2.ExternalRadiusServer.ID
				resourceMap["name"] = vvName
				d.SetId(joinResourceID(resourceMap))
				return resourceExternalRadiusServerRead(ctx, d, m)
			}
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
	return resourceExternalRadiusServerRead(ctx, d, m)
}

func resourceExternalRadiusServerRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning ExternalRadiusServer read for id=[%s]", d.Id())
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

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

		response1, restyResp1, err := client.ExternalRadiusServer.GetExternalRadiusServerByName(vvName)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItemName1 := flattenExternalRadiusServerGetExternalRadiusServerByNameItemName(response1.ExternalRadiusServer)
		if err := d.Set("item", vItemName1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetExternalRadiusServerByName response",
				err))
			return diags
		}
		if err := d.Set("parameters", vItemName1); err != nil {
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

		response2, restyResp2, err := client.ExternalRadiusServer.GetExternalRadiusServerByID(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItemID2 := flattenExternalRadiusServerGetExternalRadiusServerByIDItemID(response2.ExternalRadiusServer)
		if err := d.Set("item", vItemID2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetExternalRadiusServerByID response",
				err))
			return diags
		}
		if err := d.Set("parameters", vItemID2); err != nil {
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
	log.Printf("[DEBUG] Beginning ExternalRadiusServer update for id=[%s]", d.Id())
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

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
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] ID used for update operation %s", vvID)
		request1 := expandRequestExternalRadiusServerUpdateExternalRadiusServerByID(ctx, "parameters.0", d)
		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
		response1, restyResp1, err := client.ExternalRadiusServer.UpdateExternalRadiusServerByID(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
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
		_ = d.Set("last_updated", getUnixTimeString())
	}

	return resourceExternalRadiusServerRead(ctx, d, m)
}

func resourceExternalRadiusServerDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning ExternalRadiusServer delete for id=[%s]", d.Id())
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

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
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
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
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".host_ip")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".host_ip")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".host_ip")))) {
		request.HostIP = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".shared_secret")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".shared_secret")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".shared_secret")))) {
		request.SharedSecret = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_key_wrap")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_key_wrap")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_key_wrap")))) {
		request.EnableKeyWrap = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".encryption_key")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".encryption_key")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".encryption_key")))) {
		request.EncryptionKey = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".authenticator_key")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".authenticator_key")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".authenticator_key")))) {
		request.AuthenticatorKey = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".key_input_format")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".key_input_format")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".key_input_format")))) {
		request.KeyInputFormat = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".authentication_port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".authentication_port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".authentication_port")))) {
		request.AuthenticationPort = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".accounting_port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".accounting_port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".accounting_port")))) {
		request.AccountingPort = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".timeout")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".timeout")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".timeout")))) {
		request.Timeout = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".retries")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".retries")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".retries")))) {
		request.Retries = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".proxy_timeout")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".proxy_timeout")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".proxy_timeout")))) {
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
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".host_ip")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".host_ip")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".host_ip")))) {
		request.HostIP = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".shared_secret")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".shared_secret")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".shared_secret")))) {
		request.SharedSecret = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_key_wrap")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_key_wrap")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_key_wrap")))) {
		request.EnableKeyWrap = interfaceToBoolPtr(v)
	}
	vEnableKeyWrap, okEnableKeyWrap := d.GetOk(fixKeyAccess(key + ".enable_key_wrap"))
	vvEnableKeyWrap := interfaceToBoolPtr(vEnableKeyWrap)
	if okEnableKeyWrap && vvEnableKeyWrap != nil && *vvEnableKeyWrap {
		if v, ok := d.GetOkExists(fixKeyAccess(key + ".encryption_key")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".encryption_key")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".encryption_key")))) {
			request.EncryptionKey = interfaceToString(v)
		}
		if v, ok := d.GetOkExists(fixKeyAccess(key + ".authenticator_key")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".authenticator_key")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".authenticator_key")))) {
			request.AuthenticatorKey = interfaceToString(v)
		}
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".key_input_format")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".key_input_format")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".key_input_format")))) {
		request.KeyInputFormat = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".authentication_port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".authentication_port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".authentication_port")))) {
		request.AuthenticationPort = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".accounting_port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".accounting_port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".accounting_port")))) {
		request.AccountingPort = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".timeout")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".timeout")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".timeout")))) {
		request.Timeout = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".retries")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".retries")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".retries")))) {
		request.Retries = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".proxy_timeout")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".proxy_timeout")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".proxy_timeout")))) {
		request.ProxyTimeout = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
