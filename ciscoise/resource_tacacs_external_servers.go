package ciscoise

import (
	"context"
	"reflect"

	"ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTacacsExternalServers() *schema.Resource {
	return &schema.Resource{

		CreateContext: resourceTacacsExternalServersCreate,
		ReadContext:   resourceTacacsExternalServersRead,
		UpdateContext: resourceTacacsExternalServersUpdate,
		DeleteContext: resourceTacacsExternalServersDelete,
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

						"connection_port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"description": &schema.Schema{
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
						"shared_secret": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"single_connect": &schema.Schema{
							Type:     schema.TypeBool,
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

func resourceTacacsExternalServersCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("item"))
	request1 := expandRequestTacacsExternalServersCreateTacacsExternalServers(ctx, "item.0", d)
	log.Printf("[DEBUG] request1 => %v", responseInterfaceToString(*request1))

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vName, okName := resourceItem["name"]
	vvName := interfaceToString(vName)
	if okID && vvID != "" {
		getResponse1, _, err := client.TacacsExternalServers.GetTacacsExternalServersByID(vvID)
		if err == nil && getResponse1 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return diags
		}
	}
	if okName && vvName != "" {
		getResponse2, _, err := client.TacacsExternalServers.GetTacacsExternalServersByName(vvName)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return diags
		}
	}
	restyResp1, err := client.TacacsExternalServers.CreateTacacsExternalServers(request1)
	if err != nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateTacacsExternalServers", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateTacacsExternalServers", err))
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

func resourceTacacsExternalServersRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vName, okName := resourceMap["name"]
	vID, okID := resourceMap["id"]

	method1 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetTacacsExternalServersByName")
		vvName := vName

		response1, _, err := client.TacacsExternalServers.GetTacacsExternalServersByName(vvName)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetTacacsExternalServersByName", err,
				"Failure at GetTacacsExternalServersByName, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		vItemName1 := flattenTacacsExternalServersGetTacacsExternalServersByNameItemName(response1.TacacsExternalServer)
		if err := d.Set("item", vItemName1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTacacsExternalServersByName response",
				err))
			return diags
		}
		return diags

	}
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetTacacsExternalServersByID")
		vvID := vID

		response2, _, err := client.TacacsExternalServers.GetTacacsExternalServersByID(vvID)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetTacacsExternalServersByID", err,
				"Failure at GetTacacsExternalServersByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

		vItemID2 := flattenTacacsExternalServersGetTacacsExternalServersByIDItemID(response2.TacacsExternalServer)
		if err := d.Set("item", vItemID2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTacacsExternalServersByID response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceTacacsExternalServersUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vName, okName := resourceMap["name"]
	vID, okID := resourceMap["id"]

	method1 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	var vvID string
	var vvName string
	if selectedMethod == 1 {
		vvID = vID
	}
	if selectedMethod == 2 {
		vvName = vName
		getResp, _, err := client.TacacsExternalServers.GetTacacsExternalServersByName(vvName)
		if err != nil || getResp == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetTacacsExternalServersByName", err,
				"Failure at GetTacacsExternalServersByName, unexpected response", ""))
			return diags
		}
		//Set value vvID = getResp.
		if getResp.TacacsExternalServer != nil {
			vvID = getResp.TacacsExternalServer.ID
		}
	}
	if d.HasChange("item") {
		log.Printf("[DEBUG] vvID %s", vvID)
		request1 := expandRequestTacacsExternalServersUpdateTacacsExternalServersByID(ctx, "item.0", d)
		log.Printf("[DEBUG] request1 => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.TacacsExternalServers.UpdateTacacsExternalServersByID(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] restyResp1 => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateTacacsExternalServersByID", err, restyResp1.String(),
					"Failure at UpdateTacacsExternalServersByID, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateTacacsExternalServersByID", err,
				"Failure at UpdateTacacsExternalServersByID, unexpected response", ""))
			return diags
		}
	}

	return resourceTacacsExternalServersRead(ctx, d, m)
}

func resourceTacacsExternalServersDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vName, okName := resourceMap["name"]
	vID, okID := resourceMap["id"]

	method1 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	var vvID string
	var vvName string
	if selectedMethod == 1 {
		vvID = vID
		getResp, _, err := client.TacacsExternalServers.GetTacacsExternalServersByID(vvID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	if selectedMethod == 2 {
		vvName = vName
		getResp, _, err := client.TacacsExternalServers.GetTacacsExternalServersByName(vvName)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
		//Set value vvID = getResp.
		if getResp.TacacsExternalServer != nil {
			vvID = getResp.TacacsExternalServer.ID
		}
	}
	restyResp1, err := client.TacacsExternalServers.DeleteTacacsExternalServersByID(vvID)
	if err != nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] restyResp1 => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteTacacsExternalServersByID", err, restyResp1.String(),
				"Failure at DeleteTacacsExternalServersByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteTacacsExternalServersByID", err,
			"Failure at DeleteTacacsExternalServersByID, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestTacacsExternalServersCreateTacacsExternalServers(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestTacacsExternalServersCreateTacacsExternalServers {
	request := isegosdk.RequestTacacsExternalServersCreateTacacsExternalServers{}
	request.TacacsExternalServer = expandRequestTacacsExternalServersCreateTacacsExternalServersTacacsExternalServer(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestTacacsExternalServersCreateTacacsExternalServersTacacsExternalServer(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestTacacsExternalServersCreateTacacsExternalServersTacacsExternalServer {
	request := isegosdk.RequestTacacsExternalServersCreateTacacsExternalServersTacacsExternalServer{}
	if v, ok := d.GetOkExists(key + ".name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".name"))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".host_ip"); !isEmptyValue(reflect.ValueOf(d.Get(key+".host_ip"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".host_ip"))) {
		request.HostIP = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".connection_port"); !isEmptyValue(reflect.ValueOf(d.Get(key+".connection_port"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".connection_port"))) {
		request.ConnectionPort = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".single_connect"); !isEmptyValue(reflect.ValueOf(d.Get(key+".single_connect"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".single_connect"))) {
		request.SingleConnect = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".shared_secret"); !isEmptyValue(reflect.ValueOf(d.Get(key+".shared_secret"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".shared_secret"))) {
		request.SharedSecret = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".timeout"); !isEmptyValue(reflect.ValueOf(d.Get(key+".timeout"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".timeout"))) {
		request.Timeout = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestTacacsExternalServersUpdateTacacsExternalServersByID(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestTacacsExternalServersUpdateTacacsExternalServersByID {
	request := isegosdk.RequestTacacsExternalServersUpdateTacacsExternalServersByID{}
	request.TacacsExternalServer = expandRequestTacacsExternalServersUpdateTacacsExternalServersByIDTacacsExternalServer(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestTacacsExternalServersUpdateTacacsExternalServersByIDTacacsExternalServer(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestTacacsExternalServersUpdateTacacsExternalServersByIDTacacsExternalServer {
	request := isegosdk.RequestTacacsExternalServersUpdateTacacsExternalServersByIDTacacsExternalServer{}
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
	if v, ok := d.GetOkExists(key + ".connection_port"); !isEmptyValue(reflect.ValueOf(d.Get(key+".connection_port"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".connection_port"))) {
		request.ConnectionPort = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".single_connect"); !isEmptyValue(reflect.ValueOf(d.Get(key+".single_connect"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".single_connect"))) {
		request.SingleConnect = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".shared_secret"); !isEmptyValue(reflect.ValueOf(d.Get(key+".shared_secret"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".shared_secret"))) {
		request.SharedSecret = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".timeout"); !isEmptyValue(reflect.ValueOf(d.Get(key+".timeout"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".timeout"))) {
		request.Timeout = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
