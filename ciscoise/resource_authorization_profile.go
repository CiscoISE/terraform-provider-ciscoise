package ciscoise

import (
	"context"
	"fmt"
	"reflect"

	"ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAuthorizationProfile() *schema.Resource {
	return &schema.Resource{

		CreateContext: resourceAuthorizationProfileCreate,
		ReadContext:   resourceAuthorizationProfileRead,
		UpdateContext: resourceAuthorizationProfileUpdate,
		DeleteContext: resourceAuthorizationProfileDelete,
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

						"access_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"acl": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"advanced_attributes": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"left_hand_side_dictionary_attribue": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"advanced_attribute_value_type": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"attribute_name": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"dictionary_name": &schema.Schema{
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
									"right_hand_side_attribue_value": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"advanced_attribute_value_type": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"attribute_name": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"dictionary_name": &schema.Schema{
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
								},
							},
						},
						"agentless_posture": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"airespace_acl": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"airespace_ipv6_acl": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"asa_vpn": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"authz_profile_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"auto_smart_port": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"avc_profile": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dacl_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"easywired_session_candidate": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"interface_template": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv6_acl_filter": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv6_dacl_name": &schema.Schema{
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
						"mac_sec_policy": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"neat": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"profile_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"reauth": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"connectivity": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"timer": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"service_template": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"track_movement": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"vlan": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"name_id": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"tag_id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"voice_domain_permission": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"web_auth": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"web_redirection": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"web_redirection_type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"acl": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"display_certificates_renewal_messages": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"portal_name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"static_iphost_name_fqd_n": &schema.Schema{
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
	}
}

func resourceAuthorizationProfileCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("item"))
	request1 := expandRequestAuthorizationProfileCreateAuthorizationProfile(ctx, "item.0", d)
	log.Printf("[DEBUG] request1 => %v", responseInterfaceToString(*request1))

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vName, okName := resourceItem["name"]
	vvName := interfaceToString(vName)
	if okID && vvID != "" {
		getResponse1, _, err := client.AuthorizationProfile.GetAuthorizationProfileByID(vvID)
		if err == nil && getResponse1 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return diags
		}
	}
	if okName && vvName != "" {
		getResponse2, _, err := client.AuthorizationProfile.GetAuthorizationProfileByName(vvName)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return diags
		}
	}
	restyResp1, err := client.AuthorizationProfile.CreateAuthorizationProfile(request1)
	if err != nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateAuthorizationProfile", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateAuthorizationProfile", err))
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

func resourceAuthorizationProfileRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		log.Printf("[DEBUG] Selected method: GetAuthorizationProfileByName")
		vvName := vName

		response1, _, err := client.AuthorizationProfile.GetAuthorizationProfileByName(vvName)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetAuthorizationProfileByName", err,
				"Failure at GetAuthorizationProfileByName, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		vItemName1 := flattenAuthorizationProfileGetAuthorizationProfileByNameItemName(response1.AuthorizationProfile)
		if err := d.Set("item", vItemName1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAuthorizationProfileByName response",
				err))
			return diags
		}
		return diags

	}
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetAuthorizationProfileByID")
		vvID := vID

		response2, _, err := client.AuthorizationProfile.GetAuthorizationProfileByID(vvID)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetAuthorizationProfileByID", err,
				"Failure at GetAuthorizationProfileByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

		vItemID2 := flattenAuthorizationProfileGetAuthorizationProfileByIDItemID(response2.AuthorizationProfile)
		if err := d.Set("item", vItemID2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAuthorizationProfileByID response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceAuthorizationProfileUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		getResp, _, err := client.AuthorizationProfile.GetAuthorizationProfileByName(vvName)
		if err != nil || getResp == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetAuthorizationProfileByName", err,
				"Failure at GetAuthorizationProfileByName, unexpected response", ""))
			return diags
		}
		//Set value vvID = getResp.
		if getResp.AuthorizationProfile != nil {
			vvID = getResp.AuthorizationProfile.ID
		}
	}
	if d.HasChange("item") {
		log.Printf("[DEBUG] vvID %s", vvID)
		request1 := expandRequestAuthorizationProfileUpdateAuthorizationProfileByID(ctx, "item.0", d)
		log.Printf("[DEBUG] request1 => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.AuthorizationProfile.UpdateAuthorizationProfileByID(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] restyResp1 => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateAuthorizationProfileByID", err, restyResp1.String(),
					"Failure at UpdateAuthorizationProfileByID, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateAuthorizationProfileByID", err,
				"Failure at UpdateAuthorizationProfileByID, unexpected response", ""))
			return diags
		}
	}

	return resourceAuthorizationProfileRead(ctx, d, m)
}

func resourceAuthorizationProfileDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		getResp, _, err := client.AuthorizationProfile.GetAuthorizationProfileByID(vvID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	if selectedMethod == 2 {
		vvName = vName
		getResp, _, err := client.AuthorizationProfile.GetAuthorizationProfileByName(vvName)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
		//Set value vvID = getResp.
		if getResp.AuthorizationProfile != nil {
			vvID = getResp.AuthorizationProfile.ID
		}
	}
	restyResp1, err := client.AuthorizationProfile.DeleteAuthorizationProfileByID(vvID)
	if err != nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] restyResp1 => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteAuthorizationProfileByID", err, restyResp1.String(),
				"Failure at DeleteAuthorizationProfileByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteAuthorizationProfileByID", err,
			"Failure at DeleteAuthorizationProfileByID, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestAuthorizationProfileCreateAuthorizationProfile(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestAuthorizationProfileCreateAuthorizationProfile {
	request := isegosdk.RequestAuthorizationProfileCreateAuthorizationProfile{}
	request.AuthorizationProfile = expandRequestAuthorizationProfileCreateAuthorizationProfileAuthorizationProfile(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAuthorizationProfileCreateAuthorizationProfileAuthorizationProfile(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestAuthorizationProfileCreateAuthorizationProfileAuthorizationProfile {
	request := isegosdk.RequestAuthorizationProfileCreateAuthorizationProfileAuthorizationProfile{}
	if v, ok := d.GetOkExists(key + ".id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".id"))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".name"))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".advanced_attributes"); !isEmptyValue(reflect.ValueOf(d.Get(key+".advanced_attributes"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".advanced_attributes"))) {
		request.AdvancedAttributes = expandRequestAuthorizationProfileCreateAuthorizationProfileAuthorizationProfileAdvancedAttributesArray(ctx, key, d)
	}
	if v, ok := d.GetOkExists(key + ".access_type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".access_type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".access_type"))) {
		request.AccessType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".authz_profile_type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".authz_profile_type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".authz_profile_type"))) {
		request.AuthzProfileType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".vlan"); !isEmptyValue(reflect.ValueOf(d.Get(key+".vlan"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".vlan"))) {
		request.VLAN = expandRequestAuthorizationProfileCreateAuthorizationProfileAuthorizationProfileVLAN(ctx, key+".vlan.0", d)
	}
	if v, ok := d.GetOkExists(key + ".reauth"); !isEmptyValue(reflect.ValueOf(d.Get(key+".reauth"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".reauth"))) {
		request.Reauth = expandRequestAuthorizationProfileCreateAuthorizationProfileAuthorizationProfileReauth(ctx, key+".reauth.0", d)
	}
	if v, ok := d.GetOkExists(key + ".airespace_acl"); !isEmptyValue(reflect.ValueOf(d.Get(key+".airespace_acl"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".airespace_acl"))) {
		request.AirespaceACL = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".airespace_ipv6_acl"); !isEmptyValue(reflect.ValueOf(d.Get(key+".airespace_ipv6_acl"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".airespace_ipv6_acl"))) {
		request.AirespaceIPv6ACL = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".web_redirection"); !isEmptyValue(reflect.ValueOf(d.Get(key+".web_redirection"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".web_redirection"))) {
		request.WebRedirection = expandRequestAuthorizationProfileCreateAuthorizationProfileAuthorizationProfileWebRedirection(ctx, key+".web_redirection.0", d)
	}
	if v, ok := d.GetOkExists(key + ".acl"); !isEmptyValue(reflect.ValueOf(d.Get(key+".acl"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".acl"))) {
		request.ACL = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".track_movement"); !isEmptyValue(reflect.ValueOf(d.Get(key+".track_movement"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".track_movement"))) {
		request.TrackMovement = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".agentless_posture"); !isEmptyValue(reflect.ValueOf(d.Get(key+".agentless_posture"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".agentless_posture"))) {
		request.AgentlessPosture = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".service_template"); !isEmptyValue(reflect.ValueOf(d.Get(key+".service_template"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".service_template"))) {
		request.ServiceTemplate = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".easywired_session_candidate"); !isEmptyValue(reflect.ValueOf(d.Get(key+".easywired_session_candidate"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".easywired_session_candidate"))) {
		request.EasywiredSessionCandidate = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".dacl_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".dacl_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".dacl_name"))) {
		request.DaclName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".voice_domain_permission"); !isEmptyValue(reflect.ValueOf(d.Get(key+".voice_domain_permission"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".voice_domain_permission"))) {
		request.VoiceDomainPermission = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".neat"); !isEmptyValue(reflect.ValueOf(d.Get(key+".neat"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".neat"))) {
		request.Neat = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".web_auth"); !isEmptyValue(reflect.ValueOf(d.Get(key+".web_auth"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".web_auth"))) {
		request.WebAuth = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".auto_smart_port"); !isEmptyValue(reflect.ValueOf(d.Get(key+".auto_smart_port"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".auto_smart_port"))) {
		request.AutoSmartPort = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".interface_template"); !isEmptyValue(reflect.ValueOf(d.Get(key+".interface_template"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".interface_template"))) {
		request.InterfaceTemplate = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".ipv6_acl_filter"); !isEmptyValue(reflect.ValueOf(d.Get(key+".ipv6_acl_filter"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".ipv6_acl_filter"))) {
		request.IPv6ACLFilter = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".avc_profile"); !isEmptyValue(reflect.ValueOf(d.Get(key+".avc_profile"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".avc_profile"))) {
		request.AvcProfile = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".mac_sec_policy"); !isEmptyValue(reflect.ValueOf(d.Get(key+".mac_sec_policy"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".mac_sec_policy"))) {
		request.MacSecPolicy = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".asa_vpn"); !isEmptyValue(reflect.ValueOf(d.Get(key+".asa_vpn"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".asa_vpn"))) {
		request.AsaVpn = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".profile_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".profile_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".profile_name"))) {
		request.ProfileName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".ipv6_dacl_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".ipv6_dacl_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".ipv6_dacl_name"))) {
		request.IPv6DaclName = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAuthorizationProfileCreateAuthorizationProfileAuthorizationProfileAdvancedAttributesArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestAuthorizationProfileCreateAuthorizationProfileAuthorizationProfileAdvancedAttributes {
	request := []isegosdk.RequestAuthorizationProfileCreateAuthorizationProfileAuthorizationProfileAdvancedAttributes{}
	o := d.Get(key)
	if o != nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestAuthorizationProfileCreateAuthorizationProfileAuthorizationProfileAdvancedAttributes(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		request = append(request, *i)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAuthorizationProfileCreateAuthorizationProfileAuthorizationProfileAdvancedAttributes(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestAuthorizationProfileCreateAuthorizationProfileAuthorizationProfileAdvancedAttributes {
	request := isegosdk.RequestAuthorizationProfileCreateAuthorizationProfileAuthorizationProfileAdvancedAttributes{}
	if v, ok := d.GetOkExists(key + ".left_hand_side_dictionary_attribue"); !isEmptyValue(reflect.ValueOf(d.Get(key+".left_hand_side_dictionary_attribue"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".left_hand_side_dictionary_attribue"))) {
		request.LeftHandSideDictionaryAttribue = expandRequestAuthorizationProfileCreateAuthorizationProfileAuthorizationProfileAdvancedAttributesLeftHandSideDictionaryAttribue(ctx, key+".left_hand_side_dictionary_attribue.0", d)
	}
	if v, ok := d.GetOkExists(key + ".right_hand_side_attribue_value"); !isEmptyValue(reflect.ValueOf(d.Get(key+".right_hand_side_attribue_value"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".right_hand_side_attribue_value"))) {
		request.RightHandSideAttribueValue = expandRequestAuthorizationProfileCreateAuthorizationProfileAuthorizationProfileAdvancedAttributesRightHandSideAttribueValue(ctx, key+".right_hand_side_attribue_value.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAuthorizationProfileCreateAuthorizationProfileAuthorizationProfileAdvancedAttributesLeftHandSideDictionaryAttribue(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestAuthorizationProfileCreateAuthorizationProfileAuthorizationProfileAdvancedAttributesLeftHandSideDictionaryAttribue {
	request := isegosdk.RequestAuthorizationProfileCreateAuthorizationProfileAuthorizationProfileAdvancedAttributesLeftHandSideDictionaryAttribue{}
	if v, ok := d.GetOkExists(key + ".advanced_attribute_value_type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".advanced_attribute_value_type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".advanced_attribute_value_type"))) {
		request.AdvancedAttributeValueType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".dictionary_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".dictionary_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".dictionary_name"))) {
		request.DictionaryName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".attribute_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".attribute_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".attribute_name"))) {
		request.AttributeName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".value"); !isEmptyValue(reflect.ValueOf(d.Get(key+".value"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".value"))) {
		request.Value = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAuthorizationProfileCreateAuthorizationProfileAuthorizationProfileAdvancedAttributesRightHandSideAttribueValue(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestAuthorizationProfileCreateAuthorizationProfileAuthorizationProfileAdvancedAttributesRightHandSideAttribueValue {
	request := isegosdk.RequestAuthorizationProfileCreateAuthorizationProfileAuthorizationProfileAdvancedAttributesRightHandSideAttribueValue{}
	if v, ok := d.GetOkExists(key + ".advanced_attribute_value_type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".advanced_attribute_value_type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".advanced_attribute_value_type"))) {
		request.AdvancedAttributeValueType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".dictionary_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".dictionary_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".dictionary_name"))) {
		request.DictionaryName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".attribute_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".attribute_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".attribute_name"))) {
		request.AttributeName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".value"); !isEmptyValue(reflect.ValueOf(d.Get(key+".value"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".value"))) {
		request.Value = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAuthorizationProfileCreateAuthorizationProfileAuthorizationProfileVLAN(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestAuthorizationProfileCreateAuthorizationProfileAuthorizationProfileVLAN {
	request := isegosdk.RequestAuthorizationProfileCreateAuthorizationProfileAuthorizationProfileVLAN{}
	if v, ok := d.GetOkExists(key + ".name_id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".name_id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".name_id"))) {
		request.NameID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".tag_id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".tag_id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".tag_id"))) {
		request.TagID = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAuthorizationProfileCreateAuthorizationProfileAuthorizationProfileReauth(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestAuthorizationProfileCreateAuthorizationProfileAuthorizationProfileReauth {
	request := isegosdk.RequestAuthorizationProfileCreateAuthorizationProfileAuthorizationProfileReauth{}
	if v, ok := d.GetOkExists(key + ".timer"); !isEmptyValue(reflect.ValueOf(d.Get(key+".timer"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".timer"))) {
		request.Timer = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".connectivity"); !isEmptyValue(reflect.ValueOf(d.Get(key+".connectivity"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".connectivity"))) {
		request.Connectivity = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAuthorizationProfileCreateAuthorizationProfileAuthorizationProfileWebRedirection(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestAuthorizationProfileCreateAuthorizationProfileAuthorizationProfileWebRedirection {
	request := isegosdk.RequestAuthorizationProfileCreateAuthorizationProfileAuthorizationProfileWebRedirection{}
	if v, ok := d.GetOkExists(key + ".web_redirection_type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".web_redirection_type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".web_redirection_type"))) {
		request.WebRedirectionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".acl"); !isEmptyValue(reflect.ValueOf(d.Get(key+".acl"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".acl"))) {
		request.ACL = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".portal_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".portal_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".portal_name"))) {
		request.PortalName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".static_iphost_name_fqd_n"); !isEmptyValue(reflect.ValueOf(d.Get(key+".static_iphost_name_fqd_n"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".static_iphost_name_fqd_n"))) {
		request.StaticIPHostNameFQDN = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".display_certificates_renewal_messages"); !isEmptyValue(reflect.ValueOf(d.Get(key+".display_certificates_renewal_messages"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".display_certificates_renewal_messages"))) {
		request.DisplayCertificatesRenewalMessages = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAuthorizationProfileUpdateAuthorizationProfileByID(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestAuthorizationProfileUpdateAuthorizationProfileByID {
	request := isegosdk.RequestAuthorizationProfileUpdateAuthorizationProfileByID{}
	request.AuthorizationProfile = expandRequestAuthorizationProfileUpdateAuthorizationProfileByIDAuthorizationProfile(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAuthorizationProfileUpdateAuthorizationProfileByIDAuthorizationProfile(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestAuthorizationProfileUpdateAuthorizationProfileByIDAuthorizationProfile {
	request := isegosdk.RequestAuthorizationProfileUpdateAuthorizationProfileByIDAuthorizationProfile{}
	if v, ok := d.GetOkExists(key + ".id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".id"))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".name"))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".description"); !isEmptyValue(reflect.ValueOf(d.Get(key+".description"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".description"))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".advanced_attributes"); !isEmptyValue(reflect.ValueOf(d.Get(key+".advanced_attributes"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".advanced_attributes"))) {
		request.AdvancedAttributes = expandRequestAuthorizationProfileUpdateAuthorizationProfileByIDAuthorizationProfileAdvancedAttributesArray(ctx, key, d)
	}
	if v, ok := d.GetOkExists(key + ".access_type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".access_type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".access_type"))) {
		request.AccessType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".authz_profile_type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".authz_profile_type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".authz_profile_type"))) {
		request.AuthzProfileType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".vlan"); !isEmptyValue(reflect.ValueOf(d.Get(key+".vlan"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".vlan"))) {
		request.VLAN = expandRequestAuthorizationProfileUpdateAuthorizationProfileByIDAuthorizationProfileVLAN(ctx, key+".vlan.0", d)
	}
	if v, ok := d.GetOkExists(key + ".reauth"); !isEmptyValue(reflect.ValueOf(d.Get(key+".reauth"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".reauth"))) {
		request.Reauth = expandRequestAuthorizationProfileUpdateAuthorizationProfileByIDAuthorizationProfileReauth(ctx, key+".reauth.0", d)
	}
	if v, ok := d.GetOkExists(key + ".airespace_acl"); !isEmptyValue(reflect.ValueOf(d.Get(key+".airespace_acl"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".airespace_acl"))) {
		request.AirespaceACL = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".airespace_ipv6_acl"); !isEmptyValue(reflect.ValueOf(d.Get(key+".airespace_ipv6_acl"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".airespace_ipv6_acl"))) {
		request.AirespaceIPv6ACL = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".web_redirection"); !isEmptyValue(reflect.ValueOf(d.Get(key+".web_redirection"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".web_redirection"))) {
		request.WebRedirection = expandRequestAuthorizationProfileUpdateAuthorizationProfileByIDAuthorizationProfileWebRedirection(ctx, key+".web_redirection.0", d)
	}
	if v, ok := d.GetOkExists(key + ".acl"); !isEmptyValue(reflect.ValueOf(d.Get(key+".acl"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".acl"))) {
		request.ACL = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".track_movement"); !isEmptyValue(reflect.ValueOf(d.Get(key+".track_movement"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".track_movement"))) {
		request.TrackMovement = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".agentless_posture"); !isEmptyValue(reflect.ValueOf(d.Get(key+".agentless_posture"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".agentless_posture"))) {
		request.AgentlessPosture = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".service_template"); !isEmptyValue(reflect.ValueOf(d.Get(key+".service_template"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".service_template"))) {
		request.ServiceTemplate = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".easywired_session_candidate"); !isEmptyValue(reflect.ValueOf(d.Get(key+".easywired_session_candidate"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".easywired_session_candidate"))) {
		request.EasywiredSessionCandidate = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".dacl_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".dacl_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".dacl_name"))) {
		request.DaclName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".voice_domain_permission"); !isEmptyValue(reflect.ValueOf(d.Get(key+".voice_domain_permission"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".voice_domain_permission"))) {
		request.VoiceDomainPermission = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".neat"); !isEmptyValue(reflect.ValueOf(d.Get(key+".neat"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".neat"))) {
		request.Neat = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".web_auth"); !isEmptyValue(reflect.ValueOf(d.Get(key+".web_auth"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".web_auth"))) {
		request.WebAuth = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".auto_smart_port"); !isEmptyValue(reflect.ValueOf(d.Get(key+".auto_smart_port"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".auto_smart_port"))) {
		request.AutoSmartPort = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".interface_template"); !isEmptyValue(reflect.ValueOf(d.Get(key+".interface_template"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".interface_template"))) {
		request.InterfaceTemplate = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".ipv6_acl_filter"); !isEmptyValue(reflect.ValueOf(d.Get(key+".ipv6_acl_filter"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".ipv6_acl_filter"))) {
		request.IPv6ACLFilter = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".avc_profile"); !isEmptyValue(reflect.ValueOf(d.Get(key+".avc_profile"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".avc_profile"))) {
		request.AvcProfile = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".mac_sec_policy"); !isEmptyValue(reflect.ValueOf(d.Get(key+".mac_sec_policy"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".mac_sec_policy"))) {
		request.MacSecPolicy = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".asa_vpn"); !isEmptyValue(reflect.ValueOf(d.Get(key+".asa_vpn"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".asa_vpn"))) {
		request.AsaVpn = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".profile_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".profile_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".profile_name"))) {
		request.ProfileName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".ipv6_dacl_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".ipv6_dacl_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".ipv6_dacl_name"))) {
		request.IPv6DaclName = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAuthorizationProfileUpdateAuthorizationProfileByIDAuthorizationProfileAdvancedAttributesArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestAuthorizationProfileUpdateAuthorizationProfileByIDAuthorizationProfileAdvancedAttributes {
	request := []isegosdk.RequestAuthorizationProfileUpdateAuthorizationProfileByIDAuthorizationProfileAdvancedAttributes{}
	o := d.Get(key)
	if o != nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestAuthorizationProfileUpdateAuthorizationProfileByIDAuthorizationProfileAdvancedAttributes(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		request = append(request, *i)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAuthorizationProfileUpdateAuthorizationProfileByIDAuthorizationProfileAdvancedAttributes(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestAuthorizationProfileUpdateAuthorizationProfileByIDAuthorizationProfileAdvancedAttributes {
	request := isegosdk.RequestAuthorizationProfileUpdateAuthorizationProfileByIDAuthorizationProfileAdvancedAttributes{}
	if v, ok := d.GetOkExists(key + ".left_hand_side_dictionary_attribue"); !isEmptyValue(reflect.ValueOf(d.Get(key+".left_hand_side_dictionary_attribue"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".left_hand_side_dictionary_attribue"))) {
		request.LeftHandSideDictionaryAttribue = expandRequestAuthorizationProfileUpdateAuthorizationProfileByIDAuthorizationProfileAdvancedAttributesLeftHandSideDictionaryAttribue(ctx, key+".left_hand_side_dictionary_attribue.0", d)
	}
	if v, ok := d.GetOkExists(key + ".right_hand_side_attribue_value"); !isEmptyValue(reflect.ValueOf(d.Get(key+".right_hand_side_attribue_value"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".right_hand_side_attribue_value"))) {
		request.RightHandSideAttribueValue = expandRequestAuthorizationProfileUpdateAuthorizationProfileByIDAuthorizationProfileAdvancedAttributesRightHandSideAttribueValue(ctx, key+".right_hand_side_attribue_value.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAuthorizationProfileUpdateAuthorizationProfileByIDAuthorizationProfileAdvancedAttributesLeftHandSideDictionaryAttribue(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestAuthorizationProfileUpdateAuthorizationProfileByIDAuthorizationProfileAdvancedAttributesLeftHandSideDictionaryAttribue {
	request := isegosdk.RequestAuthorizationProfileUpdateAuthorizationProfileByIDAuthorizationProfileAdvancedAttributesLeftHandSideDictionaryAttribue{}
	if v, ok := d.GetOkExists(key + ".advanced_attribute_value_type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".advanced_attribute_value_type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".advanced_attribute_value_type"))) {
		request.AdvancedAttributeValueType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".dictionary_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".dictionary_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".dictionary_name"))) {
		request.DictionaryName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".attribute_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".attribute_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".attribute_name"))) {
		request.AttributeName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".value"); !isEmptyValue(reflect.ValueOf(d.Get(key+".value"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".value"))) {
		request.Value = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAuthorizationProfileUpdateAuthorizationProfileByIDAuthorizationProfileAdvancedAttributesRightHandSideAttribueValue(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestAuthorizationProfileUpdateAuthorizationProfileByIDAuthorizationProfileAdvancedAttributesRightHandSideAttribueValue {
	request := isegosdk.RequestAuthorizationProfileUpdateAuthorizationProfileByIDAuthorizationProfileAdvancedAttributesRightHandSideAttribueValue{}
	if v, ok := d.GetOkExists(key + ".advanced_attribute_value_type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".advanced_attribute_value_type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".advanced_attribute_value_type"))) {
		request.AdvancedAttributeValueType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".dictionary_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".dictionary_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".dictionary_name"))) {
		request.DictionaryName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".attribute_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".attribute_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".attribute_name"))) {
		request.AttributeName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".value"); !isEmptyValue(reflect.ValueOf(d.Get(key+".value"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".value"))) {
		request.Value = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAuthorizationProfileUpdateAuthorizationProfileByIDAuthorizationProfileVLAN(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestAuthorizationProfileUpdateAuthorizationProfileByIDAuthorizationProfileVLAN {
	request := isegosdk.RequestAuthorizationProfileUpdateAuthorizationProfileByIDAuthorizationProfileVLAN{}
	if v, ok := d.GetOkExists(key + ".name_id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".name_id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".name_id"))) {
		request.NameID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".tag_id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".tag_id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".tag_id"))) {
		request.TagID = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAuthorizationProfileUpdateAuthorizationProfileByIDAuthorizationProfileReauth(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestAuthorizationProfileUpdateAuthorizationProfileByIDAuthorizationProfileReauth {
	request := isegosdk.RequestAuthorizationProfileUpdateAuthorizationProfileByIDAuthorizationProfileReauth{}
	if v, ok := d.GetOkExists(key + ".timer"); !isEmptyValue(reflect.ValueOf(d.Get(key+".timer"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".timer"))) {
		request.Timer = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".connectivity"); !isEmptyValue(reflect.ValueOf(d.Get(key+".connectivity"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".connectivity"))) {
		request.Connectivity = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAuthorizationProfileUpdateAuthorizationProfileByIDAuthorizationProfileWebRedirection(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestAuthorizationProfileUpdateAuthorizationProfileByIDAuthorizationProfileWebRedirection {
	request := isegosdk.RequestAuthorizationProfileUpdateAuthorizationProfileByIDAuthorizationProfileWebRedirection{}
	if v, ok := d.GetOkExists(key + ".web_redirection_type"); !isEmptyValue(reflect.ValueOf(d.Get(key+".web_redirection_type"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".web_redirection_type"))) {
		request.WebRedirectionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".acl"); !isEmptyValue(reflect.ValueOf(d.Get(key+".acl"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".acl"))) {
		request.ACL = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".portal_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".portal_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".portal_name"))) {
		request.PortalName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".static_iphost_name_fqd_n"); !isEmptyValue(reflect.ValueOf(d.Get(key+".static_iphost_name_fqd_n"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".static_iphost_name_fqd_n"))) {
		request.StaticIPHostNameFQDN = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".display_certificates_renewal_messages"); !isEmptyValue(reflect.ValueOf(d.Get(key+".display_certificates_renewal_messages"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".display_certificates_renewal_messages"))) {
		request.DisplayCertificatesRenewalMessages = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
