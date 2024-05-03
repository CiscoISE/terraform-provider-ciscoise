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

func resourceIPsec() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read and update operations on Native IPsec.

- Updates the configuration of existing IPsec connection.
 The following parameters are present in the PUT request body:




PARAMETER

DESCRIPTION

EXAMPLE





id
*required

ID of the existing IPsec configuration.

"id": "7c9484cf-0ebc-47ad-a9ef-bc12729ed73b"



iface
*required

Ethernet port used for establishing connection

"iface": "0"



psk
*required

Pre-shared key used for establishing connection.

"psk": "psk12345"



authType
*required

Pre-shared key used for establishing connection.

"auth_type": "psk"



configureVti

Used For VTI Configurations

"configure_vti": "false"



remotePeerInternalIp

VTI Internal IP of the NAD

"remote_peer_internal_ip": "1.2.3.1"



localInternalIp

IP address assigned to the VTI interface so this would be the internal ip

"local_internal_ip": "1.1.3.1"



certId
*required

ID of the certificate for establishing connection.

"cert_id": "21323243545433"



phaseOneEncryptionAlgo
*required

Phase-one encryption algorithm used for establishing connection.

"phase_one_encryption_algo": "aes"



phaseTwoEncryptionAlgo
*required

Phase-two encryption algorithm used for establishing connection.

"phase_two_encryption_algo": "aes"



espAhProtocol
*required

Encryption protocol used for establishing connection.

"esp_ah_protocol": "ah"



phaseOneHashAlgo
*required

Phase-one hashing algorithm used for establishing connection.

"phase_one_hash_algo": "sha"



phaseTwoHashAlgo
*required

Phase-two hashing algorithm used for establishing connection.

"phase_two_hash_algo": "sha"



phaseOneDHGroup
*required

Phase-one DH group used for establishing connection.

"phase_one_dhgroup": "GROUP1"



phaseTwoDHGroup

Phase-two DH group used for establishing connection.

"phase_two_dhgroup": "GROUP1"



phaseOneLifeTime

DH Phase-one connection lifetime.

"phase_one_life_time": 14400



phaseTwoLifeTime

DH Phase-two connection lifetime.

"phase_two_life_time": 14400



ikeVersion
*required

IKE version.

"ike_version": "1"



ikeReAuthTime

IKE re-authentication time.

"ike_re_auth_time": 86400



nadIp
*required

NAD IP for establishing connection.

"nad_ip": "1.1.1.1"



modeOption
*required

The Mode type used for establishing the connection.

"mode_option": "tunnel"




NOTE:

psk
field is mandatory if authType=psk
certId
field is mandatory if authType=x509

If FIPS mode is on.:


Cannot choose DES or 3DES for Phase-one and Phase-two Encryption algorithms.

PSK length must be 14 characters or more.

DH Groups 1, 2, and 5 cannot be chosen for Phase-one and Phase-two fields.




- Creates an IPsec connection.
 The following parameters are present in the POST request body:




PARAMETER

DESCRIPTION

EXAMPLE





hostName
*required

Hostname of the node for which IPsec should be enabled

"host_name": "ise-host1"



iface
*required

Ethernet port used for establishing connection

"iface": "0"



psk
*required

Pre-shared key used for establishing connection.

"psk": "psk12345"



authType
*required

Pre-shared key used for establishing connection.

"auth_type": "psk"



configureVti

Used For VTI Configurations

"configure_vti": "false"



remotePeerInternalIp

VTI Internal IP of the NAD

"remote_peer_internal_ip": "1.2.3.1"



localInternalIp

IP address assigned to the VTI interface so this would be the internal ip

"local_internal_ip": "1.1.3.1"



certId
*required

ID of the certificate for establishing connection.

"cert_id": "21323243545433"



phaseOneEncryptionAlgo
*required

Phase-one encryption algorithm used for establishing connection.

"phase_one_encryption_algo": "aes"



phaseTwoEncryptionAlgo
*required

Phase-two encryption algorithm used for establishing connection.

"phase_two_encryption_algo": "aes"



espAhProtocol
*required

Encryption protocol used for establishing connection.

"esp_ah_protocol": "ah"



phaseOneHashAlgo
*required

Phase-one hashing algorithm used for establishing connection.

"phase_one_hash_algo": "sha"



phaseTwoHashAlgo
*required

Phase-two hashing algorithm used for establishing connection.

"phase_two_hash_algo": "sha"



phaseOneDHGroup
*required

Phase-one DH group used for establishing connection.

"phase_one_dhgroup": "GROUP1"



phaseTwoDHGroup

Phase-two DH group used for establishing connection.

"phase_two_dhgroup": "GROUP1"



phaseOneLifeTime

DH Phase-one connection lifetime.

"phase_one_life_time": 14400



phaseTwoLifeTime

DH Phase-two connection lifetime.

"phase_two_life_time": 14400



ikeVersion
*required

IKE version.

"ike_version": "1"



ikeReAuthTime

IKE re-authentication time.

"ike_re_auth_time": 86400



nadIp
*required

NAD IP for establishing the connection.

"nad_ip": "1.1.1.1"



modeOption
*required

The Mode type used for establishing the connection.

"mode_option": "tunnel"




NOTE:

psk
field is mandatory if authType=psk
certId
field is mandatory if authType=x509

If FIPS mode is on.:


Cannot choose DES or 3DES for Phase-one and Phase-two Encryption algorithms.

PSK length must be 14 characters or more.

DH Groups 1, 2, and 5 cannot be chosen for Phase-one and Phase-two fields.



`,

		CreateContext: resourceIPsecCreate,
		ReadContext:   resourceIPsecRead,
		UpdateContext: resourceIPsecUpdate,
		DeleteContext: resourceIPsec2Delete,
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
						"auth_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"cert_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"configure_vti": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"create_time": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"esp_ah_protocol": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"host_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"iface": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ike_re_auth_time": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"ike_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"link": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"href": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"rel": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"type": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"local_internal_ip": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"mode_option": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"nad_ip": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"phase_one_dhgroup": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"phase_one_encryption_algo": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"phase_one_hash_algo": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"phase_one_life_time": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"phase_two_dhgroup": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"phase_two_encryption_algo": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"phase_two_hash_algo": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"phase_two_life_time": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"psk": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"remote_peer_internal_ip": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"update_time": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"auth_type": &schema.Schema{
							Description:      `Authentication type for establishing connection`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"cert_id": &schema.Schema{
							Description:      `ID of the certificate for establishing connection`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"configure_vti": &schema.Schema{
							Description:      `Authentication type for establishing connection`,
							Type:             schema.TypeString,
							ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:         true,
							DiffSuppressFunc: diffSupressBool(),
							Computed:         true,
						},
						"esp_ah_protocol": &schema.Schema{
							Description:      `Encryption protocol used for establishing connection`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"host_name": &schema.Schema{
							Description:      `Hostname of the node`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"iface": &schema.Schema{
							Description:      `Ethernet port of the node`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"ike_re_auth_time": &schema.Schema{
							Description:      `IKE re-authentication time`,
							Type:             schema.TypeInt,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"ike_version": &schema.Schema{
							Description:      `IKE version`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"local_internal_ip": &schema.Schema{
							Description:      `Local Tunnel IP address`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"mode_option": &schema.Schema{
							Description:      `The Mode type used for establishing the connection`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"nad_ip": &schema.Schema{
							Description:      `NAD IP address for establishing connection`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"phase_one_dhgroup": &schema.Schema{
							Description:      `Phase-one DH group used for establishing connection`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"phase_one_encryption_algo": &schema.Schema{
							Description:      `Phase-one encryption algorithm used for establishing connection`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"phase_one_hash_algo": &schema.Schema{
							Description:      `Phase-one hashing algorithm used for establishing connection`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"phase_one_life_time": &schema.Schema{
							Description:      `Phase-one connection lifetime`,
							Type:             schema.TypeInt,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"phase_two_dhgroup": &schema.Schema{
							Description:      `Phase-two DH group used for establishing connection`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"phase_two_encryption_algo": &schema.Schema{
							Description:      `Phase-two encryption algorithm used for establishing connection`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"phase_two_hash_algo": &schema.Schema{
							Description:      `Phase-two hashing algorithm used for establishing connection`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"phase_two_life_time": &schema.Schema{
							Description:      `Phase-two connection lifetime`,
							Type:             schema.TypeInt,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"psk": &schema.Schema{
							Description:      `Pre-shared key used for establishing connection`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"remote_peer_internal_ip": &schema.Schema{
							Description:      `Remote Tunnel IP address`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"response": &schema.Schema{
							Type:     schema.TypeMap,
							Computed: true,
						},
						"version": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceIPsecCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client
	isEnableAutoImport := m.(ClientConfig).EnableAutoImport
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestIPsecCreateIPsecConnection(ctx, "parameters.0", d)

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	vHostName, okHostName := resourceItem["host_name"]
	vvHostName := interfaceToString(vHostName)
	vNadIP, okNadIP := resourceItem["nad_ip"]
	vvNadIP := interfaceToString(vNadIP)
	if isEnableAutoImport {
		if okHostName && vvHostName != "" && okNadIP && vvNadIP != "" {
			getResponse2, _, err := client.NativeIPsec.GetIPsecNode(vvHostName, vvNadIP)
			if err == nil && getResponse2 != nil {
				resourceMap := make(map[string]string)
				resourceMap["host_name"] = vvHostName
				resourceMap["nad_ip"] = vvNadIP
				d.SetId(joinResourceID(resourceMap))
				return resourceIPsecRead(ctx, d, m)
			}
		} else {
			queryParams2 := isegosdk.GetIPsecEnabledNodesQueryParams{}

			response2, _, err := client.NativeIPsec.GetIPsecEnabledNodes(&queryParams2)
			if response2 != nil && err == nil {
				items2 := getAllItemsNativeIPsecGetIPsecEnabledNodes(m, response2, &queryParams2)
				item2, err := searchNativeIPsecGetIPsecEnabledNodes(m, items2, vvHostName, vvNadIP)
				if err == nil && item2 != nil {
					resourceMap := make(map[string]string)
					resourceMap["host_name"] = vvHostName
					resourceMap["nad_ip"] = vvNadIP
					d.SetId(joinResourceID(resourceMap))
					return resourceIPsecRead(ctx, d, m)
				}
			}
		}
	}
	resp1, restyResp1, err := client.NativeIPsec.CreateIPsecConnection(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateIPsecConnection", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateIPsecConnection", err))
		return diags
	}
	if vvHostName != resp1.Response.HostName {
		vvHostName = resp1.Response.HostName
	}
	if vvNadIP != resp1.Response.NadIP {
		vvNadIP = resp1.Response.NadIP
	}
	resourceMap := make(map[string]string)
	resourceMap["host_name"] = vvHostName
	resourceMap["nad_ip"] = vvNadIP
	d.SetId(joinResourceID(resourceMap))
	return resourceIPsecRead(ctx, d, m)
}

func resourceIPsecRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vHostName, okHostName := resourceMap["host_name"]
	vvHostName := vHostName
	vNadIP, okNadIP := resourceMap["nad_ip"]
	vvNadIP := vNadIP

	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okHostName, okNadIP}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetIPsecEnabledNodes")
		queryParams1 := isegosdk.GetIPsecEnabledNodesQueryParams{}

		response1, restyResp1, err := client.NativeIPsec.GetIPsecEnabledNodes(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		items1 := getAllItemsNativeIPsecGetIPsecEnabledNodes(m, response1, &queryParams1)
		item1, err := searchNativeIPsecGetIPsecEnabledNodes(m, items1, vvHostName, vvNadIP)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}
		// Review flatten function used
		vItem1 := flattenNativeIPsecGetIPsecNodeItemResponse(item1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetIPsecEnabledNodes search response",
				err))
			return diags
		}
		if err := d.Set("parameters", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetIPsecEnabledNodes search response",
				err))
			return diags
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetIPsecNode")

		response2, restyResp2, err := client.NativeIPsec.GetIPsecNode(vvHostName, vvNadIP)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenNativeIPsecGetIPsecNodeItemResponse(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetIPsecNode response",
				err))
			return diags
		}
		if err := d.Set("parameters", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetIPsecNode response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceIPsecUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vHostName, _ := resourceMap["host_name"]
	vvHostName := vHostName
	vNadIP, _ := resourceMap["nad_ip"]
	vvNadIP := vNadIP

	if d.HasChange("parameters") {

		log.Printf("[DEBUG] Name used for update operation %s, %s", vvHostName, vvNadIP)

		request1 := expandRequestIPsecUpdateIPsecConnectionConfig(ctx, "parameters.0", d)

		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

		response1, restyResp1, err := client.NativeIPsec.UpdateIPsecConnectionConfig(request1)

		if err != nil || response1 == nil {

			if restyResp1 != nil {

				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())

				diags = append(diags, diagErrorWithAltAndResponse(

					"Failure when executing UpdateIPsecConnectionConfig", err, restyResp1.String(),

					"Failure at UpdateIPsecConnectionConfig, unexpected response", ""))

				return diags

			}

			diags = append(diags, diagErrorWithAlt(

				"Failure when executing UpdateIPsecConnectionConfig", err,

				"Failure at UpdateIPsecConnectionConfig, unexpected response", ""))

			return diags

		}

	}

	return resourceIPsecRead(ctx, d, m)
}

func resourceIPsec2Delete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// NOTE: Unable to delete IPsec on Cisco ISE
	//       Returning empty diags to delete it on Terraform
	return diags
}
func expandRequestIPsecCreateIPsecConnection(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNativeIPsecCreateIPsecConnection {
	request := isegosdk.RequestNativeIPsecCreateIPsecConnection{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".auth_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".auth_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".auth_type")))) {
		request.AuthType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".cert_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".cert_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".cert_id")))) {
		request.CertID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".configure_vti")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".configure_vti")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".configure_vti")))) {
		request.ConfigureVti = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".esp_ah_protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".esp_ah_protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".esp_ah_protocol")))) {
		request.EspAhProtocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".host_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".host_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".host_name")))) {
		request.HostName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".iface")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".iface")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".iface")))) {
		request.Iface = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ike_re_auth_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ike_re_auth_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ike_re_auth_time")))) {
		request.IkeReAuthTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ike_version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ike_version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ike_version")))) {
		request.IkeVersion = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".local_internal_ip")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".local_internal_ip")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".local_internal_ip")))) {
		request.LocalInternalIP = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mode_option")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mode_option")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mode_option")))) {
		request.ModeOption = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".nad_ip")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".nad_ip")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".nad_ip")))) {
		request.NadIP = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".phase_one_dhgroup")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".phase_one_dhgroup")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".phase_one_dhgroup")))) {
		request.PhaseOneDHGroup = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".phase_one_encryption_algo")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".phase_one_encryption_algo")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".phase_one_encryption_algo")))) {
		request.PhaseOneEncryptionAlgo = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".phase_one_hash_algo")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".phase_one_hash_algo")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".phase_one_hash_algo")))) {
		request.PhaseOneHashAlgo = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".phase_one_life_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".phase_one_life_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".phase_one_life_time")))) {
		request.PhaseOneLifeTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".phase_two_dhgroup")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".phase_two_dhgroup")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".phase_two_dhgroup")))) {
		request.PhaseTwoDHGroup = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".phase_two_encryption_algo")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".phase_two_encryption_algo")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".phase_two_encryption_algo")))) {
		request.PhaseTwoEncryptionAlgo = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".phase_two_hash_algo")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".phase_two_hash_algo")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".phase_two_hash_algo")))) {
		request.PhaseTwoHashAlgo = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".phase_two_life_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".phase_two_life_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".phase_two_life_time")))) {
		request.PhaseTwoLifeTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".psk")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".psk")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".psk")))) {
		request.Psk = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".remote_peer_internal_ip")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".remote_peer_internal_ip")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".remote_peer_internal_ip")))) {
		request.RemotePeerInternalIP = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestIPsecUpdateIPsecConnectionConfig(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNativeIPsecUpdateIPsecConnectionConfig {
	request := isegosdk.RequestNativeIPsecUpdateIPsecConnectionConfig{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".auth_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".auth_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".auth_type")))) {
		request.AuthType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".cert_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".cert_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".cert_id")))) {
		request.CertID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".configure_vti")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".configure_vti")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".configure_vti")))) {
		request.ConfigureVti = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".esp_ah_protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".esp_ah_protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".esp_ah_protocol")))) {
		request.EspAhProtocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".host_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".host_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".host_name")))) {
		request.HostName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".iface")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".iface")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".iface")))) {
		request.Iface = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ike_re_auth_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ike_re_auth_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ike_re_auth_time")))) {
		request.IkeReAuthTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ike_version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ike_version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ike_version")))) {
		request.IkeVersion = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".local_internal_ip")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".local_internal_ip")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".local_internal_ip")))) {
		request.LocalInternalIP = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mode_option")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mode_option")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mode_option")))) {
		request.ModeOption = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".nad_ip")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".nad_ip")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".nad_ip")))) {
		request.NadIP = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".phase_one_dhgroup")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".phase_one_dhgroup")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".phase_one_dhgroup")))) {
		request.PhaseOneDHGroup = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".phase_one_encryption_algo")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".phase_one_encryption_algo")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".phase_one_encryption_algo")))) {
		request.PhaseOneEncryptionAlgo = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".phase_one_hash_algo")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".phase_one_hash_algo")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".phase_one_hash_algo")))) {
		request.PhaseOneHashAlgo = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".phase_one_life_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".phase_one_life_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".phase_one_life_time")))) {
		request.PhaseOneLifeTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".phase_two_dhgroup")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".phase_two_dhgroup")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".phase_two_dhgroup")))) {
		request.PhaseTwoDHGroup = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".phase_two_encryption_algo")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".phase_two_encryption_algo")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".phase_two_encryption_algo")))) {
		request.PhaseTwoEncryptionAlgo = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".phase_two_hash_algo")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".phase_two_hash_algo")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".phase_two_hash_algo")))) {
		request.PhaseTwoHashAlgo = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".phase_two_life_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".phase_two_life_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".phase_two_life_time")))) {
		request.PhaseTwoLifeTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".psk")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".psk")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".psk")))) {
		request.Psk = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".remote_peer_internal_ip")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".remote_peer_internal_ip")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".remote_peer_internal_ip")))) {
		request.RemotePeerInternalIP = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func getAllItemsNativeIPsecGetIPsecEnabledNodes(m interface{}, response *isegosdk.ResponseNativeIPsecGetIPsecEnabledNodes, queryParams *isegosdk.GetIPsecEnabledNodesQueryParams) []isegosdk.ResponseNativeIPsecGetIPsecEnabledNodesResponse {
	client := m.(*isegosdk.Client)
	var respItems []isegosdk.ResponseNativeIPsecGetIPsecEnabledNodesResponse
	for response.Response != nil && len(*response.Response) > 0 {
		respItems = append(respItems, *response.Response...)
		if response.NextPage != nil && response.NextPage.Rel == "next" {
			href := response.NextPage.Href
			page, size, err := getNextPageAndSizeParams(href)
			if err != nil {
				break
			}
			if queryParams != nil {
				queryParams.Page = page
				queryParams.Size = size
			}
			response, _, err = client.NativeIPsec.GetIPsecEnabledNodes(queryParams)
			if err != nil {
				break
			}
			// All is good, continue to the next page
			continue
		}
		// Does not have next page finish iteration
		break
	}
	return respItems
}

func searchNativeIPsecGetIPsecEnabledNodes(m interface{}, items []isegosdk.ResponseNativeIPsecGetIPsecEnabledNodesResponse, hostName string, nadIP string) (*isegosdk.ResponseNativeIPsecGetIPsecNodeResponse, error) {
	client := m.(*isegosdk.Client)
	var err error
	var foundItem *isegosdk.ResponseNativeIPsecGetIPsecNodeResponse
	for _, item := range items {
		if hostName != "" && nadIP != "" && item.HostName == hostName && item.NadIP == nadIP {
			var getItem *isegosdk.ResponseNativeIPsecGetIPsecNode
			getItem, _, err = client.NativeIPsec.GetIPsecNode(hostName, nadIP)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetIPsecNode")
			}
			foundItem = getItem.Response
			return foundItem, err
		}
	}
	return foundItem, err
}
