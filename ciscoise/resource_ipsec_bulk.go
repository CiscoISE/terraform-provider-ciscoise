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

// resourceAction
func resourceIPsecBulk() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Native IPsec.

- Create, update, disable, enable and remove IPsec connections in bulk
`,

		CreateContext: resourceIPsecBulkCreate,
		ReadContext:   resourceIPsecBulkRead,
		DeleteContext: resourceIPsecBulkDelete,
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

						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth_type": &schema.Schema{
							Description: `Authentication type for establishing connection`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
						},
						"cert_id": &schema.Schema{
							Description: `ID of the certificate for establishing connection`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
						},
						"configure_vti": &schema.Schema{
							Description: `Authentication type for establishing connection`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
						"esp_ah_protocol": &schema.Schema{
							Description: `Encryption protocol used for establishing connection`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
						},
						"host_name": &schema.Schema{
							Description: `Hostname of the node`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
						},
						"iface": &schema.Schema{
							Description: `Ethernet port of the node`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
						},
						"ike_re_auth_time": &schema.Schema{
							Description: `IKE re-authentication time`,
							Type:        schema.TypeInt,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
						},
						"ike_version": &schema.Schema{
							Description: `IKE version`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
						},
						"local_internal_ip": &schema.Schema{
							Description: `Local Tunnel IP address`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
						},
						"mode_option": &schema.Schema{
							Description: `The Mode type used for establishing the connection`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
						},
						"nad_ip": &schema.Schema{
							Description: `NAD IP address for establishing connection`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
						},
						"phase_one_dhgroup": &schema.Schema{
							Description: `Phase-one DH group used for establishing connection`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
						},
						"phase_one_encryption_algo": &schema.Schema{
							Description: `Phase-one encryption algorithm used for establishing connection`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
						},
						"phase_one_hash_algo": &schema.Schema{
							Description: `Phase-one hashing algorithm used for establishing connection`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
						},
						"phase_one_life_time": &schema.Schema{
							Description: `Phase-one connection lifetime`,
							Type:        schema.TypeInt,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
						},
						"phase_two_dhgroup": &schema.Schema{
							Description: `Phase-two DH group used for establishing connection`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
						},
						"phase_two_encryption_algo": &schema.Schema{
							Description: `Phase-two encryption algorithm used for establishing connection`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
						},
						"phase_two_hash_algo": &schema.Schema{
							Description: `Phase-two hashing algorithm used for establishing connection`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
						},
						"phase_two_life_time": &schema.Schema{
							Description: `Phase-two connection lifetime`,
							Type:        schema.TypeInt,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
						},
						"psk": &schema.Schema{
							Description: `Pre-shared key used for establishing connection`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
						},
						"remote_peer_internal_ip": &schema.Schema{
							Description: `Remote Tunnel IP address`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func resourceIPsecBulkCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	request1 := expandRequestIPsecBulkBulkIPSecOperation(ctx, "parameters.0", d)

	response1, restyResp1, err := client.NativeIPsec.BulkIPSecOperation(request1)

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		d.SetId("")
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	vItem1 := flattenNativeIPsecBulkIPSecOperationItem(response1)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting BulkIPSecOperation response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags
}

func expandRequestIPsecBulkBulkIPSecOperation(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNativeIPsecBulkIPSecOperation {
	request := isegosdk.RequestNativeIPsecBulkIPSecOperation{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".item_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".item_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".item_list")))) {
		request.ItemList = expandRequestIPsecBulkBulkIPSecOperationItemListArray(ctx, key+".item_list", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".operation")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".operation")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".operation")))) {
		request.Operation = interfaceToString(v)
	}
	return &request
}

func expandRequestIPsecBulkBulkIPSecOperationItemListArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestNativeIPsecBulkIPSecOperationItemList {
	request := []isegosdk.RequestNativeIPsecBulkIPSecOperationItemList{}
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
		i := expandRequestIPsecBulkBulkIPSecOperationItemList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestIPsecBulkBulkIPSecOperationItemList(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNativeIPsecBulkIPSecOperationItemList {
	request := isegosdk.RequestNativeIPsecBulkIPSecOperationItemList{}
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
	return &request
}

func flattenNativeIPsecBulkIPSecOperationItem(item *isegosdk.ResponseNativeIPsecBulkIPSecOperation) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	return []map[string]interface{}{
		respItem,
	}
}

func resourceIPsecBulkRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*isegosdk.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceIPsecBulkUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceIPsecBulkRead(ctx, d, m)
}

func resourceIPsecBulkDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	return diags
}
