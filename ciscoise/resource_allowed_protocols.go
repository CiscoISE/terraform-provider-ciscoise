package ciscoise

import (
	"context"
	"reflect"

	"ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAllowedProtocols() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAllowedProtocolsCreate,
		ReadContext:   resourceAllowedProtocolsRead,
		UpdateContext: resourceAllowedProtocolsUpdate,
		DeleteContext: resourceAllowedProtocolsDelete,
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
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"eap_tls": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"allow_eap_tls_auth_of_expired_certs": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"eap_tls_enable_stateless_session_resume": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"eap_tls_session_ticket_ttl": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"eap_tls_session_ticket_ttl_units": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"eap_tls_session_ticket_precentage": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"peap": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"allow_peap_eap_ms_chap_v2": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"allow_peap_eap_ms_chap_v2_pwd_change": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"allow_peap_eap_ms_chap_v2_pwd_change_retries": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"allow_peap_eap_gtc": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"allow_peap_eap_gtc_pwd_change": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"allow_peap_eap_gtc_pwd_change_retries": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"allow_peap_eap_tls": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"allow_peap_eap_tls_auth_of_expired_certs": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"require_cryptobinding": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"allow_peap_v0": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"eap_fast": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"allow_eap_fast_eap_ms_chap_v2": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"allow_eap_fast_eap_ms_chap_v2_pwd_change": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"allow_eap_fast_eap_ms_chap_v2_pwd_change_retries": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"allow_eap_fast_eap_gtc": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"allow_eap_fast_eap_gtc_pwd_change": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"allow_eap_fast_eap_gtc_pwd_change_retries": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"allow_eap_fast_eap_tls": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"allow_eap_fast_eap_tls_auth_of_expired_certs": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"eap_fast_use_pacs": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"eap_fast_use_pacs_tunnel_pac_ttl": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"eap_fast_use_pacs_tunnel_pac_ttl_units": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"eap_fast_use_pacs_use_proactive_pac_update_precentage": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"eap_fast_use_pacs_allow_anonym_provisioning": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"eap_fast_use_pacs_allow_authen_provisioning": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"eap_fast_use_pacs_return_access_accept_after_authenticated_provisioning": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"eap_fast_use_pacs_accept_client_cert": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"eap_fast_use_pacs_machine_pac_ttl": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"eap_fast_use_pacs_machine_pac_ttl_units": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"eap_fast_use_pacs_allow_machine_authentication": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"eap_fast_use_pacs_stateless_session_resume": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"eap_fast_use_pacs_authorization_pac_ttl": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"eap_fast_use_pacs_authorization_pac_ttl_units": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"eap_fast_dont_use_pacs_accept_client_cert": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"eap_fast_dont_use_pacs_allow_machine_authentication": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"eap_fast_enable_eap_chaining": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"eap_ttls": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"eap_ttls_pap_ascii": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"eap_ttls_chap": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"eap_ttls_ms_chap_v1": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"eap_ttls_ms_chap_v2": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"eap_ttls_eap_md5": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"eap_ttls_eap_ms_chap_v2": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"eap_ttls_eap_ms_chap_v2_pwd_change": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"eap_ttls_eap_ms_chap_v2_pwd_change_retries": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"teap": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"allow_teap_eap_ms_chap_v2": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"allow_teap_eap_ms_chap_v2_pwd_change": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"allow_teap_eap_ms_chap_v2_pwd_change_retries": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"allow_teap_eap_tls": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"allow_teap_eap_tls_auth_of_expired_certs": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"accept_client_cert_during_tunnel_est": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"enable_eap_chaining": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"allow_downgrade_msk": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"process_host_lookup": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"allow_pap_ascii": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"allow_chap": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"allow_ms_chap_v1": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"allow_ms_chap_v2": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"allow_eap_md5": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"allow_leap": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"allow_eap_tls": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"allow_eap_ttls": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"allow_eap_fast": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"allow_peap": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"allow_teap": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"allow_preferred_eap_protocol": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"preferred_eap_protocol": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"eap_tls_l_bit": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"allow_weak_ciphers_for_eap": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"require_message_auth": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"link": &schema.Schema{
							Type:             schema.TypeList,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: diffSuppressAlways(),
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"rel": &schema.Schema{
										Type:             schema.TypeString,
										Optional:         true,
										Computed:         true,
										DiffSuppressFunc: diffSuppressAlways(),
									},
									"href": &schema.Schema{
										Type:             schema.TypeString,
										Optional:         true,
										Computed:         true,
										DiffSuppressFunc: diffSuppressAlways(),
									},
									"type": &schema.Schema{
										Type:             schema.TypeString,
										Optional:         true,
										Computed:         true,
										DiffSuppressFunc: diffSuppressAlways(),
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

func resourceAllowedProtocolsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("item"))
	request1 := expandRequestAllowedProtocolsCreateAllowedProtocol("item.0", d)
	log.Printf("[DEBUG] request1 => %v", responseInterfaceToString(*request1))

	restyResp1, err := client.AllowedProtocols.CreateAllowedProtocol(request1)
	if err != nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateAllowedProtocol", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateAllowedProtocol", err))
		return diags
	}

	resourceMap := make(map[string]string)
	resourceMap["id"] = interfaceToString(resourceItem["id"])
	resourceMap["name"] = interfaceToString(resourceItem["name"])

	d.SetId(joinResourceID(resourceMap))

	return diags
}

func resourceAllowedProtocolsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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

	if selectedMethod == 1 {
		vvID := vID
		response1, _, err := client.AllowedProtocols.GetAllowedProtocolByID(vvID)
		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetAllowedProtocolByID", err,
				"Failure at GetAllowedProtocolByID, unexpected response", ""))
			return diags
		}

		item := flattenAllowedProtocolsGetAllowedProtocolByIDItemID(response1.AllowedProtocols)
		if err := d.Set("item", item); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAllowedProtocolByID", err))
			return diags
		}

		return diags
	}

	if selectedMethod == 2 {
		vvName := vName
		response2, _, err := client.AllowedProtocols.GetAllowedProtocolByName(vvName)
		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetAllowedProtocolByName", err,
				"Failure at GetAllowedProtocolByName, unexpected response", ""))
			return diags
		}

		item := flattenAllowedProtocolsGetAllowedProtocolByNameItemName(response2.AllowedProtocols)
		if err := d.Set("item", item); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAllowedProtocolByID", err))
			return diags
		}

		return diags
	}

	return diags
}

func resourceAllowedProtocolsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
	if selectedMethod == 1 {
		vvID = vID
	}
	if selectedMethod == 2 {
		vvName := vName
		getResponse2, _, err := client.AllowedProtocols.GetAllowedProtocolByName(vvName)
		if err != nil || getResponse2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetAllowedProtocolByName", err,
				"Failure at GetAllowedProtocolByName, unexpected response", ""))
			return diags
		}
		if getResponse2.AllowedProtocols != nil {
			vvID = getResponse2.AllowedProtocols.ID
		}
	}
	if d.HasChange("item") {
		log.Printf("[DEBUG] vvID %s", vvID)

		request1 := expandRequestAllowedProtocolsUpdateAllowedProtocolByID("item.0", d)
		log.Printf("[DEBUG] request1 => %v", responseInterfaceToString(*request1))

		response1, restyResp1, err := client.AllowedProtocols.UpdateAllowedProtocolByID(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] restyResp1 => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateAllowedProtocolByID", err, restyResp1.String(),
					"Failure at UpdateAllowedProtocolByID, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateAllowedProtocolByID", err,
				"Failure at UpdateAllowedProtocolByID, unexpected response", ""))
			return diags
		}

		d.Set("last_updated", getUnixTimeString())
	}

	return resourceAllowedProtocolsRead(ctx, d, m)
}

func resourceAllowedProtocolsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
	if selectedMethod == 1 {
		vvID = vID
		getResponse1, _, err := client.AllowedProtocols.GetAllowedProtocolByID(vvID)
		if err != nil || getResponse1 == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	if selectedMethod == 2 {
		vvName := vName
		getResponse2, _, err := client.AllowedProtocols.GetAllowedProtocolByName(vvName)
		if err != nil || getResponse2 == nil {
			// Assume that element it is already gone
			return diags
		}
		if getResponse2.AllowedProtocols != nil {
			vvID = getResponse2.AllowedProtocols.ID
		}
	}

	restyResp1, err := client.AllowedProtocols.DeleteAllowedProtocolByID(vvID)
	if err != nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] restyResp1 => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteAllowedProtocolByID", err, restyResp1.String(),
				"Failure at DeleteAllowedProtocolByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteAllowedProtocolByID", err,
			"Failure at DeleteAllowedProtocolByID, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}

func expandRequestAllowedProtocolsCreateAllowedProtocol(key string, d *schema.ResourceData) *isegosdk.RequestAllowedProtocolsCreateAllowedProtocol {
	request := isegosdk.RequestAllowedProtocolsCreateAllowedProtocol{}
	request.AllowedProtocols = expandRequestAllowedProtocolsCreateAllowedProtocolAllowedProtocols(key, d)
	return &request
}

func expandRequestAllowedProtocolsCreateAllowedProtocolAllowedProtocols(key string, d *schema.ResourceData) *isegosdk.RequestAllowedProtocolsCreateAllowedProtocolAllowedProtocols {
	request := isegosdk.RequestAllowedProtocolsCreateAllowedProtocolAllowedProtocols{}
	if v, ok := d.GetOkExists(key + ".name"); ok || true {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".description"); ok || true {
		request.Description = interfaceToString(v)
	}
	if _, ok := d.GetOk(key + ".eap_tls"); ok {
		request.EapTls = expandRequestAllowedProtocolsCreateAllowedProtocolAllowedProtocolsEapTls(key+".eap_tls.0", d)
	}
	if _, ok := d.GetOk(key + ".peap"); ok {
		request.Peap = expandRequestAllowedProtocolsCreateAllowedProtocolAllowedProtocolsPeap(key+".peap.0", d)
	}
	if _, ok := d.GetOk(key + ".eap_fast"); ok {
		request.EapFast = expandRequestAllowedProtocolsCreateAllowedProtocolAllowedProtocolsEapFast(key+".eap_fast.0", d)
	}
	if _, ok := d.GetOk(key + ".eap_ttls"); ok {
		request.EapTtls = expandRequestAllowedProtocolsCreateAllowedProtocolAllowedProtocolsEapTtls(key+".eap_ttls.0", d)
	}
	if _, ok := d.GetOk(key + ".teap"); ok {
		request.Teap = expandRequestAllowedProtocolsCreateAllowedProtocolAllowedProtocolsTeap(key+".teap.0", d)
	}
	if v, ok := d.GetOkExists(key + ".process_host_lookup"); ok || true {
		request.ProcessHostLookup = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".allow_pap_ascii"); ok || true {
		request.AllowPapAscii = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".allow_chap"); ok || true {
		request.AllowChap = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".allow_ms_chap_v1"); ok || true {
		request.AllowMsChapV1 = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".allow_ms_chap_v2"); ok || true {
		request.AllowMsChapV2 = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".allow_eap_md5"); ok || true {
		request.AllowEapMd5 = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".allow_leap"); ok || true {
		request.AllowLeap = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".allow_eap_tls"); ok || true {
		request.AllowEapTls = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".allow_eap_ttls"); ok || true {
		request.AllowEapTtls = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".allow_eap_fast"); ok || true {
		request.AllowEapFast = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".allow_peap"); ok || true {
		request.AllowPeap = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".allow_teap"); ok || true {
		request.AllowTeap = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".allow_preferred_eap_protocol"); ok || true {
		request.AllowPreferredEapProtocol = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".preferred_eap_protocol"); ok || true {
		request.PreferredEapProtocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".eap_tls_l_bit"); ok || true {
		request.EapTlsLBit = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".allow_weak_ciphers_for_eap"); ok || true {
		request.AllowWeakCiphersForEap = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".require_message_auth"); ok || true {
		request.RequireMessageAuth = interfaceToBoolPtr(v)
	}
	return &request
}

func expandRequestAllowedProtocolsCreateAllowedProtocolAllowedProtocolsEapTls(key string, d *schema.ResourceData) *isegosdk.RequestAllowedProtocolsCreateAllowedProtocolAllowedProtocolsEapTls {
	request := isegosdk.RequestAllowedProtocolsCreateAllowedProtocolAllowedProtocolsEapTls{}
	if v, ok := d.GetOkExists(key + ".allow_eap_tls_auth_of_expired_certs"); ok {
		log.Printf("[DEBUG] eap_tls.allow_eap_tls_auth_of_expired_certs => %v %v", v, ok)
		request.AllowEapTlsAuthOfExpiredCerts = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".eap_tls_enable_stateless_session_resume"); ok {
		log.Printf("[DEBUG] eap_tls.eap_tls_enable_stateless_session_resume => %v %v", v, ok)
		request.EapTlsEnableStatelessSessionResume = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".eap_tls_session_ticket_ttl"); ok {
		log.Printf("[DEBUG] eap_tls.eap_tls_session_ticket_ttl => %v %v", v, ok)
		request.EapTlsSessionTicketTtl = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".eap_tls_session_ticket_ttl_units"); ok {
		log.Printf("[DEBUG] eap_tls.eap_tls_session_ticket_ttl_units => %v %v", v, ok)
		request.EapTlsSessionTicketTtlUnits = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".eap_tls_session_ticket_precentage"); ok {
		log.Printf("[DEBUG] eap_tls.eap_tls_session_ticket_precentage => %v %v", v, ok)
		request.EapTlsSessionTicketPrecentage = interfaceToIntPtr(v)
	}
	return &request
}

func expandRequestAllowedProtocolsCreateAllowedProtocolAllowedProtocolsPeap(key string, d *schema.ResourceData) *isegosdk.RequestAllowedProtocolsCreateAllowedProtocolAllowedProtocolsPeap {
	request := isegosdk.RequestAllowedProtocolsCreateAllowedProtocolAllowedProtocolsPeap{}
	if v, ok := d.GetOkExists(key + ".allow_peap_eap_ms_chap_v2"); ok {
		request.AllowPeapEapMsChapV2 = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".allow_peap_eap_ms_chap_v2_pwd_change"); ok {
		request.AllowPeapEapMsChapV2PwdChange = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".allow_peap_eap_ms_chap_v2_pwd_change_retries"); ok {
		request.AllowPeapEapMsChapV2PwdChangeRetries = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".allow_peap_eap_gtc"); ok {
		request.AllowPeapEapGtc = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".allow_peap_eap_gtc_pwd_change"); ok {
		request.AllowPeapEapGtcPwdChange = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".allow_peap_eap_gtc_pwd_change_retries"); ok {
		request.AllowPeapEapGtcPwdChangeRetries = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".allow_peap_eap_tls"); ok {
		request.AllowPeapEapTls = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".allow_peap_eap_tls_auth_of_expired_certs"); ok {
		request.AllowPeapEapTlsAuthOfExpiredCerts = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".require_cryptobinding"); ok {
		request.RequireCryptobinding = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".allow_peap_v0"); ok {
		request.AllowPeapV0 = interfaceToBoolPtr(v)
	}
	return &request
}

func expandRequestAllowedProtocolsCreateAllowedProtocolAllowedProtocolsEapFast(key string, d *schema.ResourceData) *isegosdk.RequestAllowedProtocolsCreateAllowedProtocolAllowedProtocolsEapFast {
	request := isegosdk.RequestAllowedProtocolsCreateAllowedProtocolAllowedProtocolsEapFast{}
	if v, ok := d.GetOkExists(key + ".allow_eap_fast_eap_ms_chap_v2"); ok {
		request.AllowEapFastEapMsChapV2 = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".allow_eap_fast_eap_ms_chap_v2_pwd_change"); ok {
		request.AllowEapFastEapMsChapV2PwdChange = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".allow_eap_fast_eap_ms_chap_v2_pwd_change_retries"); ok {
		request.AllowEapFastEapMsChapV2PwdChangeRetries = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".allow_eap_fast_eap_gtc"); ok {
		request.AllowEapFastEapGtc = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".allow_eap_fast_eap_gtc_pwd_change"); ok {
		request.AllowEapFastEapGtcPwdChange = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".allow_eap_fast_eap_gtc_pwd_change_retries"); ok {
		request.AllowEapFastEapGtcPwdChangeRetries = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".allow_eap_fast_eap_tls"); ok {
		request.AllowEapFastEapTls = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".allow_eap_fast_eap_tls_auth_of_expired_certs"); ok {
		request.AllowEapFastEapTlsAuthOfExpiredCerts = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".eap_fast_use_pacs"); ok {
		request.EapFastUsePacs = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".eap_fast_use_pacs_tunnel_pac_ttl"); ok {
		request.EapFastUsePacsTunnelPacTtl = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".eap_fast_use_pacs_tunnel_pac_ttl_units"); ok {
		request.EapFastUsePacsTunnelPacTtlUnits = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".eap_fast_use_pacs_use_proactive_pac_update_precentage"); ok {
		request.EapFastUsePacsUseProactivePacUpdatePrecentage = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".eap_fast_use_pacs_allow_anonym_provisioning"); ok {
		request.EapFastUsePacsAllowAnonymProvisioning = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".eap_fast_use_pacs_allow_authen_provisioning"); ok {
		request.EapFastUsePacsAllowAuthenProvisioning = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".eap_fast_use_pacs_return_access_accept_after_authenticated_provisioning"); ok {
		request.EapFastUsePacsReturnAccessAcceptAfterAuthenticatedProvisioning = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".eap_fast_use_pacs_accept_client_cert"); ok {
		request.EapFastUsePacsAcceptClientCert = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".eap_fast_use_pacs_machine_pac_ttl"); ok {
		request.EapFastUsePacsMachinePacTtl = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".eap_fast_use_pacs_machine_pac_ttl_units"); ok {
		request.EapFastUsePacsMachinePacTtlUnits = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".eap_fast_use_pacs_allow_machine_authentication"); ok {
		request.EapFastUsePacsAllowMachineAuthentication = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".eap_fast_use_pacs_stateless_session_resume"); ok {
		request.EapFastUsePacsStatelessSessionResume = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".eap_fast_use_pacs_authorization_pac_ttl"); ok {
		request.EapFastUsePacsAuthorizationPacTtl = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".eap_fast_use_pacs_authorization_pac_ttl_units"); ok {
		request.EapFastUsePacsAuthorizationPacTtlUnits = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".eap_fast_dont_use_pacs_accept_client_cert"); ok {
		request.EapFastDontUsePacsAcceptClientCert = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".eap_fast_dont_use_pacs_allow_machine_authentication"); ok {
		request.EapFastDontUsePacsAllowMachineAuthentication = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".eap_fast_enable_eap_chaining"); ok {
		request.EapFastEnableEApChaining = interfaceToBoolPtr(v)
	}
	return &request
}

func expandRequestAllowedProtocolsCreateAllowedProtocolAllowedProtocolsEapTtls(key string, d *schema.ResourceData) *isegosdk.RequestAllowedProtocolsCreateAllowedProtocolAllowedProtocolsEapTtls {
	request := isegosdk.RequestAllowedProtocolsCreateAllowedProtocolAllowedProtocolsEapTtls{}
	if v, ok := d.GetOkExists(key + ".eap_ttls_pap_ascii"); ok {
		request.EapTtlsPapAscii = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".eap_ttls_chap"); ok {
		request.EapTtlsChap = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".eap_ttls_ms_chap_v1"); ok {
		request.EapTtlsMsChapV1 = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".eap_ttls_ms_chap_v2"); ok {
		request.EapTtlsMsChapV2 = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".eap_ttls_eap_md5"); ok {
		request.EapTtlsEapMd5 = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".eap_ttls_eap_ms_chap_v2"); ok {
		request.EapTtlsEapMsChapV2 = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".eap_ttls_eap_ms_chap_v2_pwd_change"); ok {
		request.EapTtlsEapMsChapV2PwdChange = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".eap_ttls_eap_ms_chap_v2_pwd_change_retries"); ok {
		request.EapTtlsEapMsChapV2PwdChangeRetries = interfaceToIntPtr(v)
	}
	return &request
}

func expandRequestAllowedProtocolsCreateAllowedProtocolAllowedProtocolsTeap(key string, d *schema.ResourceData) *isegosdk.RequestAllowedProtocolsCreateAllowedProtocolAllowedProtocolsTeap {
	request := isegosdk.RequestAllowedProtocolsCreateAllowedProtocolAllowedProtocolsTeap{}
	if v, ok := d.GetOkExists(key + ".allow_teap_eap_ms_chap_v2"); ok {
		request.AllowTeapEapMsChapV2 = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".allow_teap_eap_ms_chap_v2_pwd_change"); ok {
		request.AllowTeapEapMsChapV2PwdChange = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".allow_teap_eap_ms_chap_v2_pwd_change_retries"); ok {
		request.AllowTeapEapMsChapV2PwdChangeRetries = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".allow_teap_eap_tls"); ok {
		request.AllowTeapEapTls = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".allow_teap_eap_tls_auth_of_expired_certs"); ok {
		request.AllowTeapEapTlsAuthOfExpiredCerts = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".accept_client_cert_during_tunnel_est"); ok {
		request.AcceptClientCertDuringTunnelEst = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".enable_eap_chaining"); ok {
		request.EnableEapChaining = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".allow_downgrade_msk"); ok {
		request.AllowDowngradeMsk = interfaceToBoolPtr(v)
	}
	return &request
}

func expandRequestAllowedProtocolsUpdateAllowedProtocolByID(key string, d *schema.ResourceData) *isegosdk.RequestAllowedProtocolsUpdateAllowedProtocolByID {
	request := isegosdk.RequestAllowedProtocolsUpdateAllowedProtocolByID{}
	request.AllowedProtocols = expandRequestAllowedProtocolsUpdateAllowedProtocolByIDAllowedProtocols(key, d)
	return &request
}

func expandRequestAllowedProtocolsUpdateAllowedProtocolByIDAllowedProtocols(key string, d *schema.ResourceData) *isegosdk.RequestAllowedProtocolsUpdateAllowedProtocolByIDAllowedProtocols {
	request := isegosdk.RequestAllowedProtocolsUpdateAllowedProtocolByIDAllowedProtocols{}
	if oldResource, newResource := d.GetChange(key + ".name"); reflect.DeepEqual(oldResource, newResource) == false || true {
		request.Name = interfaceToString(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".description"); reflect.DeepEqual(oldResource, newResource) == false || true {
		request.Description = interfaceToString(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".eap_tls"); reflect.DeepEqual(oldResource, newResource) == false {
		request.EapTls = expandRequestAllowedProtocolsUpdateAllowedProtocolByIDAllowedProtocolsEapTls(key+".eap_tls.0", d)
	}
	if oldResource, newResource := d.GetChange(key + ".peap"); reflect.DeepEqual(oldResource, newResource) == false {
		request.Peap = expandRequestAllowedProtocolsUpdateAllowedProtocolByIDAllowedProtocolsPeap(key+".peap.0", d)
	}
	if oldResource, newResource := d.GetChange(key + ".eap_fast"); reflect.DeepEqual(oldResource, newResource) == false {
		request.EapFast = expandRequestAllowedProtocolsUpdateAllowedProtocolByIDAllowedProtocolsEapFast(key+".eap_fast.0", d)
	}
	if oldResource, newResource := d.GetChange(key + ".eap_ttls"); reflect.DeepEqual(oldResource, newResource) == false {
		request.EapTtls = expandRequestAllowedProtocolsUpdateAllowedProtocolByIDAllowedProtocolsEapTtls(key+".eap_ttls.0", d)
	}
	if oldResource, newResource := d.GetChange(key + ".teap"); reflect.DeepEqual(oldResource, newResource) == false {
		request.Teap = expandRequestAllowedProtocolsUpdateAllowedProtocolByIDAllowedProtocolsTeap(key+".teap.0", d)
	}
	if oldResource, newResource := d.GetChange(key + ".process_host_lookup"); reflect.DeepEqual(oldResource, newResource) == false || true {
		request.ProcessHostLookup = interfaceToBoolPtr(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".allow_pap_ascii"); reflect.DeepEqual(oldResource, newResource) == false || true {
		request.AllowPapAscii = interfaceToBoolPtr(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".allow_chap"); reflect.DeepEqual(oldResource, newResource) == false || true {
		request.AllowChap = interfaceToBoolPtr(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".allow_ms_chap_v1"); reflect.DeepEqual(oldResource, newResource) == false || true {
		request.AllowMsChapV1 = interfaceToBoolPtr(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".allow_ms_chap_v2"); reflect.DeepEqual(oldResource, newResource) == false || true {
		request.AllowMsChapV2 = interfaceToBoolPtr(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".allow_eap_md5"); reflect.DeepEqual(oldResource, newResource) == false || true {
		request.AllowEapMd5 = interfaceToBoolPtr(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".allow_leap"); reflect.DeepEqual(oldResource, newResource) == false || true {
		request.AllowLeap = interfaceToBoolPtr(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".allow_eap_tls"); reflect.DeepEqual(oldResource, newResource) == false || true {
		request.AllowEapTls = interfaceToBoolPtr(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".allow_eap_ttls"); reflect.DeepEqual(oldResource, newResource) == false || true {
		request.AllowEapTtls = interfaceToBoolPtr(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".allow_eap_fast"); reflect.DeepEqual(oldResource, newResource) == false || true {
		request.AllowEapFast = interfaceToBoolPtr(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".allow_peap"); reflect.DeepEqual(oldResource, newResource) == false || true {
		request.AllowPeap = interfaceToBoolPtr(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".allow_teap"); reflect.DeepEqual(oldResource, newResource) == false || true {
		request.AllowTeap = interfaceToBoolPtr(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".allow_preferred_eap_protocol"); reflect.DeepEqual(oldResource, newResource) == false || true {
		request.AllowPreferredEapProtocol = interfaceToBoolPtr(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".preferred_eap_protocol"); reflect.DeepEqual(oldResource, newResource) == false || true {
		request.PreferredEapProtocol = interfaceToString(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".eap_tls_l_bit"); reflect.DeepEqual(oldResource, newResource) == false || true {
		request.EapTlsLBit = interfaceToBoolPtr(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".allow_weak_ciphers_for_eap"); reflect.DeepEqual(oldResource, newResource) == false || true {
		request.AllowWeakCiphersForEap = interfaceToBoolPtr(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".require_message_auth"); reflect.DeepEqual(oldResource, newResource) == false || true {
		request.RequireMessageAuth = interfaceToBoolPtr(newResource)
	}
	return &request
}

func expandRequestAllowedProtocolsUpdateAllowedProtocolByIDAllowedProtocolsEapTls(key string, d *schema.ResourceData) *isegosdk.RequestAllowedProtocolsUpdateAllowedProtocolByIDAllowedProtocolsEapTls {
	request := isegosdk.RequestAllowedProtocolsUpdateAllowedProtocolByIDAllowedProtocolsEapTls{}
	if oldResource, newResource := d.GetChange(key + ".allow_eap_tls_auth_of_expired_certs"); reflect.DeepEqual(oldResource, newResource) == false || true {
		request.AllowEapTlsAuthOfExpiredCerts = interfaceToBoolPtr(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".eap_tls_enable_stateless_session_resume"); reflect.DeepEqual(oldResource, newResource) == false || true {
		request.EapTlsEnableStatelessSessionResume = interfaceToBoolPtr(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".eap_tls_session_ticket_ttl"); reflect.DeepEqual(oldResource, newResource) == false {
		request.EapTlsSessionTicketTtl = interfaceToIntPtr(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".eap_tls_session_ticket_ttl_units"); reflect.DeepEqual(oldResource, newResource) == false || true {
		request.EapTlsSessionTicketTtlUnits = interfaceToString(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".eap_tls_session_ticket_precentage"); reflect.DeepEqual(oldResource, newResource) == false {
		request.EapTlsSessionTicketPrecentage = interfaceToIntPtr(newResource)
	}
	return &request
}

func expandRequestAllowedProtocolsUpdateAllowedProtocolByIDAllowedProtocolsPeap(key string, d *schema.ResourceData) *isegosdk.RequestAllowedProtocolsUpdateAllowedProtocolByIDAllowedProtocolsPeap {
	request := isegosdk.RequestAllowedProtocolsUpdateAllowedProtocolByIDAllowedProtocolsPeap{}
	if oldResource, newResource := d.GetChange(key + ".allow_peap_eap_ms_chap_v2"); reflect.DeepEqual(oldResource, newResource) == false || true {
		request.AllowPeapEapMsChapV2 = interfaceToBoolPtr(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".allow_peap_eap_ms_chap_v2_pwd_change"); reflect.DeepEqual(oldResource, newResource) == false || true {
		request.AllowPeapEapMsChapV2PwdChange = interfaceToBoolPtr(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".allow_peap_eap_ms_chap_v2_pwd_change_retries"); reflect.DeepEqual(oldResource, newResource) == false {
		request.AllowPeapEapMsChapV2PwdChangeRetries = interfaceToIntPtr(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".allow_peap_eap_gtc"); reflect.DeepEqual(oldResource, newResource) == false || true {
		request.AllowPeapEapGtc = interfaceToBoolPtr(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".allow_peap_eap_gtc_pwd_change"); reflect.DeepEqual(oldResource, newResource) == false || true {
		request.AllowPeapEapGtcPwdChange = interfaceToBoolPtr(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".allow_peap_eap_gtc_pwd_change_retries"); reflect.DeepEqual(oldResource, newResource) == false {
		request.AllowPeapEapGtcPwdChangeRetries = interfaceToIntPtr(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".allow_peap_eap_tls"); reflect.DeepEqual(oldResource, newResource) == false || true {
		request.AllowPeapEapTls = interfaceToBoolPtr(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".allow_peap_eap_tls_auth_of_expired_certs"); reflect.DeepEqual(oldResource, newResource) == false || true {
		request.AllowPeapEapTlsAuthOfExpiredCerts = interfaceToBoolPtr(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".require_cryptobinding"); reflect.DeepEqual(oldResource, newResource) == false || true {
		request.RequireCryptobinding = interfaceToBoolPtr(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".allow_peap_v0"); reflect.DeepEqual(oldResource, newResource) == false || true {
		request.AllowPeapV0 = interfaceToBoolPtr(newResource)
	}
	return &request
}

func expandRequestAllowedProtocolsUpdateAllowedProtocolByIDAllowedProtocolsEapFast(key string, d *schema.ResourceData) *isegosdk.RequestAllowedProtocolsUpdateAllowedProtocolByIDAllowedProtocolsEapFast {
	request := isegosdk.RequestAllowedProtocolsUpdateAllowedProtocolByIDAllowedProtocolsEapFast{}
	if oldResource, newResource := d.GetChange(key + ".allow_eap_fast_eap_ms_chap_v2"); reflect.DeepEqual(oldResource, newResource) == false || true {
		request.AllowEapFastEapMsChapV2 = interfaceToBoolPtr(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".allow_eap_fast_eap_ms_chap_v2_pwd_change"); reflect.DeepEqual(oldResource, newResource) == false || true {
		request.AllowEapFastEapMsChapV2PwdChange = interfaceToBoolPtr(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".allow_eap_fast_eap_ms_chap_v2_pwd_change_retries"); reflect.DeepEqual(oldResource, newResource) == false {
		request.AllowEapFastEapMsChapV2PwdChangeRetries = interfaceToIntPtr(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".allow_eap_fast_eap_gtc"); reflect.DeepEqual(oldResource, newResource) == false || true {
		request.AllowEapFastEapGtc = interfaceToBoolPtr(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".allow_eap_fast_eap_gtc_pwd_change"); reflect.DeepEqual(oldResource, newResource) == false || true {
		request.AllowEapFastEapGtcPwdChange = interfaceToBoolPtr(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".allow_eap_fast_eap_gtc_pwd_change_retries"); reflect.DeepEqual(oldResource, newResource) == false {
		request.AllowEapFastEapGtcPwdChangeRetries = interfaceToIntPtr(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".allow_eap_fast_eap_tls"); reflect.DeepEqual(oldResource, newResource) == false || true {
		request.AllowEapFastEapTls = interfaceToBoolPtr(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".allow_eap_fast_eap_tls_auth_of_expired_certs"); reflect.DeepEqual(oldResource, newResource) == false || true {
		request.AllowEapFastEapTlsAuthOfExpiredCerts = interfaceToBoolPtr(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".eap_fast_use_pacs"); reflect.DeepEqual(oldResource, newResource) == false || true {
		request.EapFastUsePacs = interfaceToBoolPtr(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".eap_fast_use_pacs_tunnel_pac_ttl"); reflect.DeepEqual(oldResource, newResource) == false {
		request.EapFastUsePacsTunnelPacTtl = interfaceToIntPtr(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".eap_fast_use_pacs_tunnel_pac_ttl_units"); reflect.DeepEqual(oldResource, newResource) == false || true {
		request.EapFastUsePacsTunnelPacTtlUnits = interfaceToString(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".eap_fast_use_pacs_use_proactive_pac_update_precentage"); reflect.DeepEqual(oldResource, newResource) == false {
		request.EapFastUsePacsUseProactivePacUpdatePrecentage = interfaceToIntPtr(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".eap_fast_use_pacs_allow_anonym_provisioning"); reflect.DeepEqual(oldResource, newResource) == false || true {
		request.EapFastUsePacsAllowAnonymProvisioning = interfaceToBoolPtr(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".eap_fast_use_pacs_allow_authen_provisioning"); reflect.DeepEqual(oldResource, newResource) == false || true {
		request.EapFastUsePacsAllowAuthenProvisioning = interfaceToBoolPtr(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".eap_fast_use_pacs_return_access_accept_after_authenticated_provisioning"); reflect.DeepEqual(oldResource, newResource) == false || true {
		request.EapFastUsePacsReturnAccessAcceptAfterAuthenticatedProvisioning = interfaceToBoolPtr(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".eap_fast_use_pacs_accept_client_cert"); reflect.DeepEqual(oldResource, newResource) == false || true {
		request.EapFastUsePacsAcceptClientCert = interfaceToBoolPtr(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".eap_fast_use_pacs_machine_pac_ttl"); reflect.DeepEqual(oldResource, newResource) == false {
		request.EapFastUsePacsMachinePacTtl = interfaceToIntPtr(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".eap_fast_use_pacs_machine_pac_ttl_units"); reflect.DeepEqual(oldResource, newResource) == false || true {
		request.EapFastUsePacsMachinePacTtlUnits = interfaceToString(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".eap_fast_use_pacs_allow_machine_authentication"); reflect.DeepEqual(oldResource, newResource) == false || true {
		request.EapFastUsePacsAllowMachineAuthentication = interfaceToBoolPtr(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".eap_fast_use_pacs_stateless_session_resume"); reflect.DeepEqual(oldResource, newResource) == false || true {
		request.EapFastUsePacsStatelessSessionResume = interfaceToBoolPtr(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".eap_fast_use_pacs_authorization_pac_ttl"); reflect.DeepEqual(oldResource, newResource) == false {
		request.EapFastUsePacsAuthorizationPacTtl = interfaceToIntPtr(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".eap_fast_use_pacs_authorization_pac_ttl_units"); reflect.DeepEqual(oldResource, newResource) == false || true {
		request.EapFastUsePacsAuthorizationPacTtlUnits = interfaceToString(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".eap_fast_dont_use_pacs_accept_client_cert"); reflect.DeepEqual(oldResource, newResource) == false || true {
		request.EapFastDontUsePacsAcceptClientCert = interfaceToBoolPtr(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".eap_fast_dont_use_pacs_allow_machine_authentication"); reflect.DeepEqual(oldResource, newResource) == false || true {
		request.EapFastDontUsePacsAllowMachineAuthentication = interfaceToBoolPtr(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".eap_fast_enable_eap_chaining"); reflect.DeepEqual(oldResource, newResource) == false || true {
		request.EapFastEnableEApChaining = interfaceToBoolPtr(newResource)
	}
	return &request
}

func expandRequestAllowedProtocolsUpdateAllowedProtocolByIDAllowedProtocolsEapTtls(key string, d *schema.ResourceData) *isegosdk.RequestAllowedProtocolsUpdateAllowedProtocolByIDAllowedProtocolsEapTtls {
	request := isegosdk.RequestAllowedProtocolsUpdateAllowedProtocolByIDAllowedProtocolsEapTtls{}
	if oldResource, newResource := d.GetChange(key + ".eap_ttls_pap_ascii"); reflect.DeepEqual(oldResource, newResource) == false || true {
		request.EapTtlsPapAscii = interfaceToBoolPtr(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".eap_ttls_chap"); reflect.DeepEqual(oldResource, newResource) == false || true {
		request.EapTtlsChap = interfaceToBoolPtr(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".eap_ttls_ms_chap_v1"); reflect.DeepEqual(oldResource, newResource) == false || true {
		request.EapTtlsMsChapV1 = interfaceToBoolPtr(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".eap_ttls_ms_chap_v2"); reflect.DeepEqual(oldResource, newResource) == false || true {
		request.EapTtlsMsChapV2 = interfaceToBoolPtr(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".eap_ttls_eap_md5"); reflect.DeepEqual(oldResource, newResource) == false || true {
		request.EapTtlsEapMd5 = interfaceToBoolPtr(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".eap_ttls_eap_ms_chap_v2"); reflect.DeepEqual(oldResource, newResource) == false || true {
		request.EapTtlsEapMsChapV2 = interfaceToBoolPtr(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".eap_ttls_eap_ms_chap_v2_pwd_change"); reflect.DeepEqual(oldResource, newResource) == false || true {
		request.EapTtlsEapMsChapV2PwdChange = interfaceToBoolPtr(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".eap_ttls_eap_ms_chap_v2_pwd_change_retries"); reflect.DeepEqual(oldResource, newResource) == false {
		request.EapTtlsEapMsChapV2PwdChangeRetries = interfaceToIntPtr(newResource)
	}
	return &request
}

func expandRequestAllowedProtocolsUpdateAllowedProtocolByIDAllowedProtocolsTeap(key string, d *schema.ResourceData) *isegosdk.RequestAllowedProtocolsUpdateAllowedProtocolByIDAllowedProtocolsTeap {
	request := isegosdk.RequestAllowedProtocolsUpdateAllowedProtocolByIDAllowedProtocolsTeap{}
	if oldResource, newResource := d.GetChange(key + ".allow_teap_eap_ms_chap_v2"); reflect.DeepEqual(oldResource, newResource) == false || true {
		request.AllowTeapEapMsChapV2 = interfaceToBoolPtr(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".allow_teap_eap_ms_chap_v2_pwd_change"); reflect.DeepEqual(oldResource, newResource) == false || true {
		request.AllowTeapEapMsChapV2PwdChange = interfaceToBoolPtr(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".allow_teap_eap_ms_chap_v2_pwd_change_retries"); reflect.DeepEqual(oldResource, newResource) == false {
		request.AllowTeapEapMsChapV2PwdChangeRetries = interfaceToIntPtr(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".allow_teap_eap_tls"); reflect.DeepEqual(oldResource, newResource) == false || true {
		request.AllowTeapEapTls = interfaceToBoolPtr(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".allow_teap_eap_tls_auth_of_expired_certs"); reflect.DeepEqual(oldResource, newResource) == false || true {
		request.AllowTeapEapTlsAuthOfExpiredCerts = interfaceToBoolPtr(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".accept_client_cert_during_tunnel_est"); reflect.DeepEqual(oldResource, newResource) == false || true {
		request.AcceptClientCertDuringTunnelEst = interfaceToBoolPtr(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".enable_eap_chaining"); reflect.DeepEqual(oldResource, newResource) == false || true {
		request.EnableEapChaining = interfaceToBoolPtr(newResource)
	}
	if oldResource, newResource := d.GetChange(key + ".allow_downgrade_msk"); reflect.DeepEqual(oldResource, newResource) == false || true {
		request.AllowDowngradeMsk = interfaceToBoolPtr(newResource)
	}
	return &request
}
