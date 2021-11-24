package ciscoise

import (
	"context"
	"fmt"
	"reflect"

	"log"

	isegosdk "ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkDevice() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on NetworkDevice.

- This resource allows the client to update a network device by name.

- This resource deletes a network device by name.

- This resource allows the client to update a network device by ID.

- This resource deletes a network device by ID.

- This resource creates a network device.
`,

		CreateContext: resourceNetworkDeviceCreate,
		ReadContext:   resourceNetworkDeviceRead,
		UpdateContext: resourceNetworkDeviceUpdate,
		DeleteContext: resourceNetworkDeviceDelete,
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

						"network_device_group_list": &schema.Schema{
							Description: `List of Network Device Group names for this node`,
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"network_device_iplist": &schema.Schema{
							Description: `List of IP Subnets for this node`,
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"get_ipaddress_exclude": &schema.Schema{
										Description: `It can be either single IP address or IP range address`,
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"ipaddress": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"mask": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"authentication_settings": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"dtls_required": &schema.Schema{
										Description:  `This value enforces use of dtls`,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										Computed:     true,
									},
									"enable_key_wrap": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										Computed:     true,
									},
									"enable_multi_secret": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"enabled": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										Computed:     true,
									},
									"key_encryption_key": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"key_input_format": &schema.Schema{
										Description: `Allowed values:
- ASCII,
- HEXADECIMAL`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"message_authenticator_code_key": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"network_protocol": &schema.Schema{
										Description: `Allowed values:
- RADIUS,
- TACACS_PLUS`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"radius_shared_secret": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"second_radius_shared_secret": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"coa_port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dtls_dns_name": &schema.Schema{
							Description: `This value is used to verify the client identity contained in the X.509 RADIUS/DTLS client certificate`,
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
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
						"model_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"profile_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"snmpsettings": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"link_trap_query": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										Computed:     true,
									},
									"mac_trap_query": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										Computed:     true,
									},
									"originating_policy_services_node": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"polling_interval": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"ro_community": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"version": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"software_version": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tacacs_settings": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"connect_mode_options": &schema.Schema{
										Description: `Allowed values:
- OFF,
- ON_LEGACY,
- ON_DRAFT_COMPLIANT`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"shared_secret": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"trustsecsettings": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"device_authentication_settings": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"sga_device_id": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"sga_device_password": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
											},
										},
									},
									"device_configuration_deployment": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"enable_mode_password": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"exec_mode_password": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"exec_mode_username": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"include_when_deploying_sgt_updates": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
													Computed:     true,
												},
											},
										},
									},
									"push_id_support": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										Computed:     true,
									},
									"sga_notification_and_updates": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"coa_source_host": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"downlaod_environment_data_every_x_seconds": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"downlaod_peer_authorization_policy_every_x_seconds": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"download_sga_cllists_every_x_seconds": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"other_sga_devices_to_trust_this_device": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
													Computed:     true,
												},
												"re_authentication_every_x_seconds": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"send_configuration_to_device": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
													Computed:     true,
												},
												"send_configuration_to_device_using": &schema.Schema{
													Description: `Allowed values:
- ENABLE_USING_COA,
- ENABLE_USING_CLI,
- DISABLE_ALL`,
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

func resourceNetworkDeviceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("item"))
	request1 := expandRequestNetworkDeviceCreateNetworkDevice(ctx, "item.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vName, okName := resourceItem["name"]
	vvName := interfaceToString(vName)
	if okID && vvID != "" {
		getResponse1, _, err := client.NetworkDevice.GetNetworkDeviceByID(vvID)
		if err == nil && getResponse1 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return diags
		}
	}
	if okName && vvName != "" {
		getResponse2, _, err := client.NetworkDevice.GetNetworkDeviceByName(vvName)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return diags
		}
	}
	restyResp1, err := client.NetworkDevice.CreateNetworkDevice(request1)
	if err != nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateNetworkDevice", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateNetworkDevice", err))
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

func resourceNetworkDeviceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		log.Printf("[DEBUG] Selected method: GetNetworkDeviceByName")
		vvName := vName

		response1, _, err := client.NetworkDevice.GetNetworkDeviceByName(vvName)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetNetworkDeviceByName", err,
				"Failure at GetNetworkDeviceByName, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItemName1 := flattenNetworkDeviceGetNetworkDeviceByNameItemName(response1.NetworkDevice)
		if err := d.Set("item", vItemName1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNetworkDeviceByName response",
				err))
			return diags
		}
		return diags

	}
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkDeviceByID")
		vvID := vID

		response2, _, err := client.NetworkDevice.GetNetworkDeviceByID(vvID)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetNetworkDeviceByID", err,
				"Failure at GetNetworkDeviceByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItemID2 := flattenNetworkDeviceGetNetworkDeviceByIDItemID(response2.NetworkDevice)
		if err := d.Set("item", vItemID2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNetworkDeviceByID response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceNetworkDeviceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		getResp, _, err := client.NetworkDevice.GetNetworkDeviceByName(vvName)
		if err != nil || getResp == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetNetworkDeviceByName", err,
				"Failure at GetNetworkDeviceByName, unexpected response", ""))
			return diags
		}
		//Set value vvID = getResp.
		if getResp.NetworkDevice != nil {
			vvID = getResp.NetworkDevice.ID
		}
	}
	if d.HasChange("item") {
		log.Printf("[DEBUG] ID used for update operation %s", vvID)
		request1 := expandRequestNetworkDeviceUpdateNetworkDeviceByID(ctx, "item.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.NetworkDevice.UpdateNetworkDeviceByID(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateNetworkDeviceByID", err, restyResp1.String(),
					"Failure at UpdateNetworkDeviceByID, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateNetworkDeviceByID", err,
				"Failure at UpdateNetworkDeviceByID, unexpected response", ""))
			return diags
		}
	}

	return resourceNetworkDeviceRead(ctx, d, m)
}

func resourceNetworkDeviceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		getResp, _, err := client.NetworkDevice.GetNetworkDeviceByID(vvID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	if selectedMethod == 2 {
		vvName = vName
		getResp, _, err := client.NetworkDevice.GetNetworkDeviceByName(vvName)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
		//Set value vvID = getResp.
		if getResp.NetworkDevice != nil {
			vvID = getResp.NetworkDevice.ID
		}
	}
	restyResp1, err := client.NetworkDevice.DeleteNetworkDeviceByID(vvID)
	if err != nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteNetworkDeviceByID", err, restyResp1.String(),
				"Failure at DeleteNetworkDeviceByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteNetworkDeviceByID", err,
			"Failure at DeleteNetworkDeviceByID, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}

func expandRequestNetworkDeviceCreateNetworkDevice(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkDeviceCreateNetworkDevice {
	request := isegosdk.RequestNetworkDeviceCreateNetworkDevice{}
	request.NetworkDevice = expandRequestNetworkDeviceCreateNetworkDeviceNetworkDevice(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkDeviceCreateNetworkDeviceNetworkDevice(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkDeviceCreateNetworkDeviceNetworkDevice {
	request := isegosdk.RequestNetworkDeviceCreateNetworkDeviceNetworkDevice{}
	if v, ok := d.GetOkExists(key + ".name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".name"))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".authentication_settings"); !isEmptyValue(reflect.ValueOf(d.Get(key+".authentication_settings"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".authentication_settings"))) {
		request.AuthenticationSettings = expandRequestNetworkDeviceCreateNetworkDeviceNetworkDeviceAuthenticationSettings(ctx, key+".authentication_settings.0", d)
	}
	if v, ok := d.GetOkExists(key + ".snmpsettings"); !isEmptyValue(reflect.ValueOf(d.Get(key+".snmpsettings"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".snmpsettings"))) {
		request.SNMPsettings = expandRequestNetworkDeviceCreateNetworkDeviceNetworkDeviceSNMPsettings(ctx, key+".snmpsettings.0", d)
	}
	if v, ok := d.GetOkExists(key + ".trustsecsettings"); !isEmptyValue(reflect.ValueOf(d.Get(key+".trustsecsettings"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".trustsecsettings"))) {
		request.Trustsecsettings = expandRequestNetworkDeviceCreateNetworkDeviceNetworkDeviceTrustsecsettings(ctx, key+".trustsecsettings.0", d)
	}
	if v, ok := d.GetOkExists(key + ".tacacs_settings"); !isEmptyValue(reflect.ValueOf(d.Get(key+".tacacs_settings"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".tacacs_settings"))) {
		request.TacacsSettings = expandRequestNetworkDeviceCreateNetworkDeviceNetworkDeviceTacacsSettings(ctx, key+".tacacs_settings.0", d)
	}
	if v, ok := d.GetOkExists(key + ".profile_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".profile_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".profile_name"))) {
		request.ProfileName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".coa_port"); !isEmptyValue(reflect.ValueOf(d.Get(key+".coa_port"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".coa_port"))) {
		request.CoaPort = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".dtls_dns_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".dtls_dns_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".dtls_dns_name"))) {
		request.DtlsDNSName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".model_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".model_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".model_name"))) {
		request.ModelName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".software_version"); !isEmptyValue(reflect.ValueOf(d.Get(key+".software_version"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".software_version"))) {
		request.SoftwareVersion = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".network_device_iplist"); !isEmptyValue(reflect.ValueOf(d.Get(key+".network_device_iplist"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".network_device_iplist"))) {
		request.NetworkDeviceIPList = expandRequestNetworkDeviceCreateNetworkDeviceNetworkDeviceNetworkDeviceIPListArray(ctx, key+".network_device_iplist", d)
	}
	if v, ok := d.GetOkExists(key + ".network_device_group_list"); !isEmptyValue(reflect.ValueOf(d.Get(key+".network_device_group_list"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".network_device_group_list"))) {
		request.NetworkDeviceGroupList = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkDeviceCreateNetworkDeviceNetworkDeviceAuthenticationSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkDeviceCreateNetworkDeviceNetworkDeviceAuthenticationSettings {
	request := isegosdk.RequestNetworkDeviceCreateNetworkDeviceNetworkDeviceAuthenticationSettings{}
	if v, ok := d.GetOkExists(key + ".network_protocol"); !isEmptyValue(reflect.ValueOf(d.Get(key+".network_protocol"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".network_protocol"))) {
		request.NetworkProtocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".second_radius_shared_secret"); !isEmptyValue(reflect.ValueOf(d.Get(key+".second_radius_shared_secret"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".second_radius_shared_secret"))) {
		request.SecondRadiusSharedSecret = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".radius_shared_secret"); !isEmptyValue(reflect.ValueOf(d.Get(key+".radius_shared_secret"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".radius_shared_secret"))) {
		request.RadiusSharedSecret = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".enable_key_wrap"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enable_key_wrap"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enable_key_wrap"))) {
		request.EnableKeyWrap = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".enabled"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enabled"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enabled"))) {
		request.Enabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".dtls_required"); !isEmptyValue(reflect.ValueOf(d.Get(key+".dtls_required"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".dtls_required"))) {
		request.DtlsRequired = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".enable_multi_secret"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enable_multi_secret"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enable_multi_secret"))) {
		request.EnableMultiSecret = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".key_encryption_key"); !isEmptyValue(reflect.ValueOf(d.Get(key+".key_encryption_key"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".key_encryption_key"))) {
		request.KeyEncryptionKey = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".message_authenticator_code_key"); !isEmptyValue(reflect.ValueOf(d.Get(key+".message_authenticator_code_key"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".message_authenticator_code_key"))) {
		request.MessageAuthenticatorCodeKey = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".key_input_format"); !isEmptyValue(reflect.ValueOf(d.Get(key+".key_input_format"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".key_input_format"))) {
		request.KeyInputFormat = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkDeviceCreateNetworkDeviceNetworkDeviceSNMPsettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkDeviceCreateNetworkDeviceNetworkDeviceSNMPsettings {
	request := isegosdk.RequestNetworkDeviceCreateNetworkDeviceNetworkDeviceSNMPsettings{}
	if v, ok := d.GetOkExists(key + ".version"); !isEmptyValue(reflect.ValueOf(d.Get(key+".version"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".version"))) {
		request.Version = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".ro_community"); !isEmptyValue(reflect.ValueOf(d.Get(key+".ro_community"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".ro_community"))) {
		request.RoCommunity = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".polling_interval"); !isEmptyValue(reflect.ValueOf(d.Get(key+".polling_interval"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".polling_interval"))) {
		request.PollingInterval = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".link_trap_query"); !isEmptyValue(reflect.ValueOf(d.Get(key+".link_trap_query"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".link_trap_query"))) {
		request.LinkTrapQuery = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".mac_trap_query"); !isEmptyValue(reflect.ValueOf(d.Get(key+".mac_trap_query"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".mac_trap_query"))) {
		request.MacTrapQuery = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".originating_policy_services_node"); !isEmptyValue(reflect.ValueOf(d.Get(key+".originating_policy_services_node"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".originating_policy_services_node"))) {
		request.OriginatingPolicyServicesNode = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkDeviceCreateNetworkDeviceNetworkDeviceTrustsecsettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkDeviceCreateNetworkDeviceNetworkDeviceTrustsecsettings {
	request := isegosdk.RequestNetworkDeviceCreateNetworkDeviceNetworkDeviceTrustsecsettings{}
	if v, ok := d.GetOkExists(key + ".device_authentication_settings"); !isEmptyValue(reflect.ValueOf(d.Get(key+".device_authentication_settings"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".device_authentication_settings"))) {
		request.DeviceAuthenticationSettings = expandRequestNetworkDeviceCreateNetworkDeviceNetworkDeviceTrustsecsettingsDeviceAuthenticationSettings(ctx, key+".device_authentication_settings.0", d)
	}
	if v, ok := d.GetOkExists(key + ".sga_notification_and_updates"); !isEmptyValue(reflect.ValueOf(d.Get(key+".sga_notification_and_updates"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".sga_notification_and_updates"))) {
		request.SgaNotificationAndUpdates = expandRequestNetworkDeviceCreateNetworkDeviceNetworkDeviceTrustsecsettingsSgaNotificationAndUpdates(ctx, key+".sga_notification_and_updates.0", d)
	}
	if v, ok := d.GetOkExists(key + ".device_configuration_deployment"); !isEmptyValue(reflect.ValueOf(d.Get(key+".device_configuration_deployment"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".device_configuration_deployment"))) {
		request.DeviceConfigurationDeployment = expandRequestNetworkDeviceCreateNetworkDeviceNetworkDeviceTrustsecsettingsDeviceConfigurationDeployment(ctx, key+".device_configuration_deployment.0", d)
	}
	if v, ok := d.GetOkExists(key + ".push_id_support"); !isEmptyValue(reflect.ValueOf(d.Get(key+".push_id_support"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".push_id_support"))) {
		request.PushIDSupport = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkDeviceCreateNetworkDeviceNetworkDeviceTrustsecsettingsDeviceAuthenticationSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkDeviceCreateNetworkDeviceNetworkDeviceTrustsecsettingsDeviceAuthenticationSettings {
	request := isegosdk.RequestNetworkDeviceCreateNetworkDeviceNetworkDeviceTrustsecsettingsDeviceAuthenticationSettings{}
	if v, ok := d.GetOkExists(key + ".sga_device_id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".sga_device_id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".sga_device_id"))) {
		request.SgaDeviceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".sga_device_password"); !isEmptyValue(reflect.ValueOf(d.Get(key+".sga_device_password"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".sga_device_password"))) {
		request.SgaDevicePassword = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkDeviceCreateNetworkDeviceNetworkDeviceTrustsecsettingsSgaNotificationAndUpdates(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkDeviceCreateNetworkDeviceNetworkDeviceTrustsecsettingsSgaNotificationAndUpdates {
	request := isegosdk.RequestNetworkDeviceCreateNetworkDeviceNetworkDeviceTrustsecsettingsSgaNotificationAndUpdates{}
	if v, ok := d.GetOkExists(key + ".downlaod_environment_data_every_x_seconds"); !isEmptyValue(reflect.ValueOf(d.Get(key+".downlaod_environment_data_every_x_seconds"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".downlaod_environment_data_every_x_seconds"))) {
		request.DownlaodEnvironmentDataEveryXSeconds = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".downlaod_peer_authorization_policy_every_x_seconds"); !isEmptyValue(reflect.ValueOf(d.Get(key+".downlaod_peer_authorization_policy_every_x_seconds"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".downlaod_peer_authorization_policy_every_x_seconds"))) {
		request.DownlaodPeerAuthorizationPolicyEveryXSeconds = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".re_authentication_every_x_seconds"); !isEmptyValue(reflect.ValueOf(d.Get(key+".re_authentication_every_x_seconds"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".re_authentication_every_x_seconds"))) {
		request.ReAuthenticationEveryXSeconds = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".download_sga_cllists_every_x_seconds"); !isEmptyValue(reflect.ValueOf(d.Get(key+".download_sga_cllists_every_x_seconds"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".download_sga_cllists_every_x_seconds"))) {
		request.DownloadSgACLListsEveryXSeconds = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".other_sga_devices_to_trust_this_device"); !isEmptyValue(reflect.ValueOf(d.Get(key+".other_sga_devices_to_trust_this_device"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".other_sga_devices_to_trust_this_device"))) {
		request.OtherSgADevicesToTrustThisDevice = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".send_configuration_to_device"); !isEmptyValue(reflect.ValueOf(d.Get(key+".send_configuration_to_device"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".send_configuration_to_device"))) {
		request.SendConfigurationToDevice = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".send_configuration_to_device_using"); !isEmptyValue(reflect.ValueOf(d.Get(key+".send_configuration_to_device_using"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".send_configuration_to_device_using"))) {
		request.SendConfigurationToDeviceUsing = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".coa_source_host"); !isEmptyValue(reflect.ValueOf(d.Get(key+".coa_source_host"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".coa_source_host"))) {
		request.CoaSourceHost = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkDeviceCreateNetworkDeviceNetworkDeviceTrustsecsettingsDeviceConfigurationDeployment(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkDeviceCreateNetworkDeviceNetworkDeviceTrustsecsettingsDeviceConfigurationDeployment {
	request := isegosdk.RequestNetworkDeviceCreateNetworkDeviceNetworkDeviceTrustsecsettingsDeviceConfigurationDeployment{}
	if v, ok := d.GetOkExists(key + ".include_when_deploying_sgt_updates"); !isEmptyValue(reflect.ValueOf(d.Get(key+".include_when_deploying_sgt_updates"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".include_when_deploying_sgt_updates"))) {
		request.IncludeWhenDeployingSgtUpdates = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".enable_mode_password"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enable_mode_password"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enable_mode_password"))) {
		request.EnableModePassword = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".exec_mode_password"); !isEmptyValue(reflect.ValueOf(d.Get(key+".exec_mode_password"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".exec_mode_password"))) {
		request.ExecModePassword = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".exec_mode_username"); !isEmptyValue(reflect.ValueOf(d.Get(key+".exec_mode_username"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".exec_mode_username"))) {
		request.ExecModeUsername = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkDeviceCreateNetworkDeviceNetworkDeviceTacacsSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkDeviceCreateNetworkDeviceNetworkDeviceTacacsSettings {
	request := isegosdk.RequestNetworkDeviceCreateNetworkDeviceNetworkDeviceTacacsSettings{}
	if v, ok := d.GetOkExists(key + ".shared_secret"); !isEmptyValue(reflect.ValueOf(d.Get(key+".shared_secret"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".shared_secret"))) {
		request.SharedSecret = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".connect_mode_options"); !isEmptyValue(reflect.ValueOf(d.Get(key+".connect_mode_options"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".connect_mode_options"))) {
		request.ConnectModeOptions = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkDeviceCreateNetworkDeviceNetworkDeviceNetworkDeviceIPListArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestNetworkDeviceCreateNetworkDeviceNetworkDeviceNetworkDeviceIPList {
	request := []isegosdk.RequestNetworkDeviceCreateNetworkDeviceNetworkDeviceNetworkDeviceIPList{}
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestNetworkDeviceCreateNetworkDeviceNetworkDeviceNetworkDeviceIPList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkDeviceCreateNetworkDeviceNetworkDeviceNetworkDeviceIPList(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkDeviceCreateNetworkDeviceNetworkDeviceNetworkDeviceIPList {
	request := isegosdk.RequestNetworkDeviceCreateNetworkDeviceNetworkDeviceNetworkDeviceIPList{}
	if v, ok := d.GetOkExists(key + ".ipaddress"); !isEmptyValue(reflect.ValueOf(d.Get(key+".ipaddress"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".ipaddress"))) {
		request.IPaddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".mask"); !isEmptyValue(reflect.ValueOf(d.Get(key+".mask"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".mask"))) {
		request.Mask = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".get_ipaddress_exclude"); !isEmptyValue(reflect.ValueOf(d.Get(key+".get_ipaddress_exclude"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".get_ipaddress_exclude"))) {
		request.GetIPaddressExclude = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkDeviceUpdateNetworkDeviceByID(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkDeviceUpdateNetworkDeviceByID {
	request := isegosdk.RequestNetworkDeviceUpdateNetworkDeviceByID{}
	request.NetworkDevice = expandRequestNetworkDeviceUpdateNetworkDeviceByIDNetworkDevice(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkDeviceUpdateNetworkDeviceByIDNetworkDevice(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkDeviceUpdateNetworkDeviceByIDNetworkDevice {
	request := isegosdk.RequestNetworkDeviceUpdateNetworkDeviceByIDNetworkDevice{}
	if v, ok := d.GetOkExists(key + ".id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".id"))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".name"))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".authentication_settings"); !isEmptyValue(reflect.ValueOf(d.Get(key+".authentication_settings"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".authentication_settings"))) {
		request.AuthenticationSettings = expandRequestNetworkDeviceUpdateNetworkDeviceByIDNetworkDeviceAuthenticationSettings(ctx, key+".authentication_settings.0", d)
	}
	if v, ok := d.GetOkExists(key + ".snmpsettings"); !isEmptyValue(reflect.ValueOf(d.Get(key+".snmpsettings"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".snmpsettings"))) {
		request.SNMPsettings = expandRequestNetworkDeviceUpdateNetworkDeviceByIDNetworkDeviceSNMPsettings(ctx, key+".snmpsettings.0", d)
	}
	if v, ok := d.GetOkExists(key + ".trustsecsettings"); !isEmptyValue(reflect.ValueOf(d.Get(key+".trustsecsettings"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".trustsecsettings"))) {
		request.Trustsecsettings = expandRequestNetworkDeviceUpdateNetworkDeviceByIDNetworkDeviceTrustsecsettings(ctx, key+".trustsecsettings.0", d)
	}
	if v, ok := d.GetOkExists(key + ".tacacs_settings"); !isEmptyValue(reflect.ValueOf(d.Get(key+".tacacs_settings"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".tacacs_settings"))) {
		request.TacacsSettings = expandRequestNetworkDeviceUpdateNetworkDeviceByIDNetworkDeviceTacacsSettings(ctx, key+".tacacs_settings.0", d)
	}
	if v, ok := d.GetOkExists(key + ".profile_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".profile_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".profile_name"))) {
		request.ProfileName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".coa_port"); !isEmptyValue(reflect.ValueOf(d.Get(key+".coa_port"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".coa_port"))) {
		request.CoaPort = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".dtls_dns_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".dtls_dns_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".dtls_dns_name"))) {
		request.DtlsDNSName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".model_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".model_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".model_name"))) {
		request.ModelName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".software_version"); !isEmptyValue(reflect.ValueOf(d.Get(key+".software_version"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".software_version"))) {
		request.SoftwareVersion = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".network_device_iplist"); !isEmptyValue(reflect.ValueOf(d.Get(key+".network_device_iplist"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".network_device_iplist"))) {
		request.NetworkDeviceIPList = expandRequestNetworkDeviceUpdateNetworkDeviceByIDNetworkDeviceNetworkDeviceIPListArray(ctx, key+".network_device_iplist", d)
	}
	if v, ok := d.GetOkExists(key + ".network_device_group_list"); !isEmptyValue(reflect.ValueOf(d.Get(key+".network_device_group_list"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".network_device_group_list"))) {
		request.NetworkDeviceGroupList = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkDeviceUpdateNetworkDeviceByIDNetworkDeviceAuthenticationSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkDeviceUpdateNetworkDeviceByIDNetworkDeviceAuthenticationSettings {
	request := isegosdk.RequestNetworkDeviceUpdateNetworkDeviceByIDNetworkDeviceAuthenticationSettings{}
	if v, ok := d.GetOkExists(key + ".network_protocol"); !isEmptyValue(reflect.ValueOf(d.Get(key+".network_protocol"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".network_protocol"))) {
		request.NetworkProtocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".second_radius_shared_secret"); !isEmptyValue(reflect.ValueOf(d.Get(key+".second_radius_shared_secret"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".second_radius_shared_secret"))) {
		request.SecondRadiusSharedSecret = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".radius_shared_secret"); !isEmptyValue(reflect.ValueOf(d.Get(key+".radius_shared_secret"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".radius_shared_secret"))) {
		request.RadiusSharedSecret = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".enable_key_wrap"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enable_key_wrap"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enable_key_wrap"))) {
		request.EnableKeyWrap = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".enabled"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enabled"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enabled"))) {
		request.Enabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".dtls_required"); !isEmptyValue(reflect.ValueOf(d.Get(key+".dtls_required"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".dtls_required"))) {
		request.DtlsRequired = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".enable_multi_secret"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enable_multi_secret"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enable_multi_secret"))) {
		request.EnableMultiSecret = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".key_encryption_key"); !isEmptyValue(reflect.ValueOf(d.Get(key+".key_encryption_key"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".key_encryption_key"))) {
		request.KeyEncryptionKey = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".message_authenticator_code_key"); !isEmptyValue(reflect.ValueOf(d.Get(key+".message_authenticator_code_key"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".message_authenticator_code_key"))) {
		request.MessageAuthenticatorCodeKey = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".key_input_format"); !isEmptyValue(reflect.ValueOf(d.Get(key+".key_input_format"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".key_input_format"))) {
		request.KeyInputFormat = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkDeviceUpdateNetworkDeviceByIDNetworkDeviceSNMPsettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkDeviceUpdateNetworkDeviceByIDNetworkDeviceSNMPsettings {
	request := isegosdk.RequestNetworkDeviceUpdateNetworkDeviceByIDNetworkDeviceSNMPsettings{}
	if v, ok := d.GetOkExists(key + ".version"); !isEmptyValue(reflect.ValueOf(d.Get(key+".version"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".version"))) {
		request.Version = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".ro_community"); !isEmptyValue(reflect.ValueOf(d.Get(key+".ro_community"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".ro_community"))) {
		request.RoCommunity = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".polling_interval"); !isEmptyValue(reflect.ValueOf(d.Get(key+".polling_interval"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".polling_interval"))) {
		request.PollingInterval = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".link_trap_query"); !isEmptyValue(reflect.ValueOf(d.Get(key+".link_trap_query"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".link_trap_query"))) {
		request.LinkTrapQuery = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".mac_trap_query"); !isEmptyValue(reflect.ValueOf(d.Get(key+".mac_trap_query"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".mac_trap_query"))) {
		request.MacTrapQuery = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".originating_policy_services_node"); !isEmptyValue(reflect.ValueOf(d.Get(key+".originating_policy_services_node"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".originating_policy_services_node"))) {
		request.OriginatingPolicyServicesNode = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkDeviceUpdateNetworkDeviceByIDNetworkDeviceTrustsecsettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkDeviceUpdateNetworkDeviceByIDNetworkDeviceTrustsecsettings {
	request := isegosdk.RequestNetworkDeviceUpdateNetworkDeviceByIDNetworkDeviceTrustsecsettings{}
	if v, ok := d.GetOkExists(key + ".device_authentication_settings"); !isEmptyValue(reflect.ValueOf(d.Get(key+".device_authentication_settings"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".device_authentication_settings"))) {
		request.DeviceAuthenticationSettings = expandRequestNetworkDeviceUpdateNetworkDeviceByIDNetworkDeviceTrustsecsettingsDeviceAuthenticationSettings(ctx, key+".device_authentication_settings.0", d)
	}
	if v, ok := d.GetOkExists(key + ".sga_notification_and_updates"); !isEmptyValue(reflect.ValueOf(d.Get(key+".sga_notification_and_updates"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".sga_notification_and_updates"))) {
		request.SgaNotificationAndUpdates = expandRequestNetworkDeviceUpdateNetworkDeviceByIDNetworkDeviceTrustsecsettingsSgaNotificationAndUpdates(ctx, key+".sga_notification_and_updates.0", d)
	}
	if v, ok := d.GetOkExists(key + ".device_configuration_deployment"); !isEmptyValue(reflect.ValueOf(d.Get(key+".device_configuration_deployment"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".device_configuration_deployment"))) {
		request.DeviceConfigurationDeployment = expandRequestNetworkDeviceUpdateNetworkDeviceByIDNetworkDeviceTrustsecsettingsDeviceConfigurationDeployment(ctx, key+".device_configuration_deployment.0", d)
	}
	if v, ok := d.GetOkExists(key + ".push_id_support"); !isEmptyValue(reflect.ValueOf(d.Get(key+".push_id_support"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".push_id_support"))) {
		request.PushIDSupport = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkDeviceUpdateNetworkDeviceByIDNetworkDeviceTrustsecsettingsDeviceAuthenticationSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkDeviceUpdateNetworkDeviceByIDNetworkDeviceTrustsecsettingsDeviceAuthenticationSettings {
	request := isegosdk.RequestNetworkDeviceUpdateNetworkDeviceByIDNetworkDeviceTrustsecsettingsDeviceAuthenticationSettings{}
	if v, ok := d.GetOkExists(key + ".sga_device_id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".sga_device_id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".sga_device_id"))) {
		request.SgaDeviceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".sga_device_password"); !isEmptyValue(reflect.ValueOf(d.Get(key+".sga_device_password"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".sga_device_password"))) {
		request.SgaDevicePassword = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkDeviceUpdateNetworkDeviceByIDNetworkDeviceTrustsecsettingsSgaNotificationAndUpdates(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkDeviceUpdateNetworkDeviceByIDNetworkDeviceTrustsecsettingsSgaNotificationAndUpdates {
	request := isegosdk.RequestNetworkDeviceUpdateNetworkDeviceByIDNetworkDeviceTrustsecsettingsSgaNotificationAndUpdates{}
	if v, ok := d.GetOkExists(key + ".downlaod_environment_data_every_x_seconds"); !isEmptyValue(reflect.ValueOf(d.Get(key+".downlaod_environment_data_every_x_seconds"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".downlaod_environment_data_every_x_seconds"))) {
		request.DownlaodEnvironmentDataEveryXSeconds = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".downlaod_peer_authorization_policy_every_x_seconds"); !isEmptyValue(reflect.ValueOf(d.Get(key+".downlaod_peer_authorization_policy_every_x_seconds"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".downlaod_peer_authorization_policy_every_x_seconds"))) {
		request.DownlaodPeerAuthorizationPolicyEveryXSeconds = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".re_authentication_every_x_seconds"); !isEmptyValue(reflect.ValueOf(d.Get(key+".re_authentication_every_x_seconds"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".re_authentication_every_x_seconds"))) {
		request.ReAuthenticationEveryXSeconds = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".download_sga_cllists_every_x_seconds"); !isEmptyValue(reflect.ValueOf(d.Get(key+".download_sga_cllists_every_x_seconds"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".download_sga_cllists_every_x_seconds"))) {
		request.DownloadSgACLListsEveryXSeconds = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".other_sga_devices_to_trust_this_device"); !isEmptyValue(reflect.ValueOf(d.Get(key+".other_sga_devices_to_trust_this_device"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".other_sga_devices_to_trust_this_device"))) {
		request.OtherSgADevicesToTrustThisDevice = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".send_configuration_to_device"); !isEmptyValue(reflect.ValueOf(d.Get(key+".send_configuration_to_device"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".send_configuration_to_device"))) {
		request.SendConfigurationToDevice = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".send_configuration_to_device_using"); !isEmptyValue(reflect.ValueOf(d.Get(key+".send_configuration_to_device_using"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".send_configuration_to_device_using"))) {
		request.SendConfigurationToDeviceUsing = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".coa_source_host"); !isEmptyValue(reflect.ValueOf(d.Get(key+".coa_source_host"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".coa_source_host"))) {
		request.CoaSourceHost = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkDeviceUpdateNetworkDeviceByIDNetworkDeviceTrustsecsettingsDeviceConfigurationDeployment(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkDeviceUpdateNetworkDeviceByIDNetworkDeviceTrustsecsettingsDeviceConfigurationDeployment {
	request := isegosdk.RequestNetworkDeviceUpdateNetworkDeviceByIDNetworkDeviceTrustsecsettingsDeviceConfigurationDeployment{}
	if v, ok := d.GetOkExists(key + ".include_when_deploying_sgt_updates"); !isEmptyValue(reflect.ValueOf(d.Get(key+".include_when_deploying_sgt_updates"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".include_when_deploying_sgt_updates"))) {
		request.IncludeWhenDeployingSgtUpdates = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".enable_mode_password"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enable_mode_password"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enable_mode_password"))) {
		request.EnableModePassword = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".exec_mode_password"); !isEmptyValue(reflect.ValueOf(d.Get(key+".exec_mode_password"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".exec_mode_password"))) {
		request.ExecModePassword = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".exec_mode_username"); !isEmptyValue(reflect.ValueOf(d.Get(key+".exec_mode_username"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".exec_mode_username"))) {
		request.ExecModeUsername = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkDeviceUpdateNetworkDeviceByIDNetworkDeviceTacacsSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkDeviceUpdateNetworkDeviceByIDNetworkDeviceTacacsSettings {
	request := isegosdk.RequestNetworkDeviceUpdateNetworkDeviceByIDNetworkDeviceTacacsSettings{}
	if v, ok := d.GetOkExists(key + ".shared_secret"); !isEmptyValue(reflect.ValueOf(d.Get(key+".shared_secret"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".shared_secret"))) {
		request.SharedSecret = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".connect_mode_options"); !isEmptyValue(reflect.ValueOf(d.Get(key+".connect_mode_options"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".connect_mode_options"))) {
		request.ConnectModeOptions = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkDeviceUpdateNetworkDeviceByIDNetworkDeviceNetworkDeviceIPListArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestNetworkDeviceUpdateNetworkDeviceByIDNetworkDeviceNetworkDeviceIPList {
	request := []isegosdk.RequestNetworkDeviceUpdateNetworkDeviceByIDNetworkDeviceNetworkDeviceIPList{}
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestNetworkDeviceUpdateNetworkDeviceByIDNetworkDeviceNetworkDeviceIPList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkDeviceUpdateNetworkDeviceByIDNetworkDeviceNetworkDeviceIPList(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkDeviceUpdateNetworkDeviceByIDNetworkDeviceNetworkDeviceIPList {
	request := isegosdk.RequestNetworkDeviceUpdateNetworkDeviceByIDNetworkDeviceNetworkDeviceIPList{}
	if v, ok := d.GetOkExists(key + ".ipaddress"); !isEmptyValue(reflect.ValueOf(d.Get(key+".ipaddress"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".ipaddress"))) {
		request.IPaddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".mask"); !isEmptyValue(reflect.ValueOf(d.Get(key+".mask"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".mask"))) {
		request.Mask = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".get_ipaddress_exclude"); !isEmptyValue(reflect.ValueOf(d.Get(key+".get_ipaddress_exclude"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".get_ipaddress_exclude"))) {
		request.GetIPaddressExclude = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
