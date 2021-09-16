package ciscoise

import (
	"context"
	"reflect"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAciSettings() *schema.Resource {
	return &schema.Resource{
		Description: `It manages read and update operations on ACISettings.
  
  This resource allows the client to update ACI settings.`,

		CreateContext: resourceAciSettingsCreate,
		ReadContext:   resourceAciSettingsRead,
		UpdateContext: resourceAciSettingsUpdate,
		DeleteContext: resourceAciSettingsDelete,
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
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"aci50": &schema.Schema{
							Description: `Enable 5.0 ACI Version`,
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"aci51": &schema.Schema{
							Description: `Enable 5.1 ACI Version`,
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"aciipaddress": &schema.Schema{
							Description: `ACI Domain manager Ip Address.`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"acipassword": &schema.Schema{
							Description: `ACI Domain manager Password.`,
							Type:        schema.TypeString,
							Optional:    true,
							Sensitive:   true,
						},
						"aciuser_name": &schema.Schema{
							Description: `ACI Domain manager Username.`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"admin_name": &schema.Schema{
							Description: `ACI Cluster Admin name`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"admin_password": &schema.Schema{
							Description: `ACI Cluster Admin password`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"all_sxp_domain": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
						"default_sgt_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"enable_aci": &schema.Schema{
							Description: `Enable ACI Integration`,
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"enable_data_plane": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
						"enable_elements_limit": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
						"id": &schema.Schema{
							Description: `Resource UUID value`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"ip_address_host_name": &schema.Schema{
							Description: `ACI Cluster IP Address / Host name`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"l3_route_network": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"max_num_iepg_from_aci": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"max_num_sgt_to_aci": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"specific_sxp_domain": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
						"specifix_sxp_domain_list": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"suffix_to_epg": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"suffix_to_sgt": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"tenant_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"untagged_packet_iepg_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourceAciSettingsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	resourceItem := *getResourceItem(d.Get("item"))
	resourceMap := make(map[string]string)
	// NOTE: Function does not perform create on ISE
	// TODO: Review if it has other unique values to use for comparisons
	resourceMap["id"] = interfaceToString(resourceItem["id"])
	d.SetId(joinResourceID(resourceMap))
	return diags
}

func resourceAciSettingsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetAciSettings")

		response1, _, err := client.AciSettings.GetAciSettings()

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetAciSettings", err,
				"Failure at GetAciSettings, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		vItem1 := flattenAciSettingsGetAciSettingsItem(response1.AciSettings)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAciSettings response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceAciSettingsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	var vvID string
	if d.HasChange("item") {
		log.Printf("[DEBUG] vvID %s", vvID)
		request1 := expandRequestAciSettingsUpdateAciSettingsByID(ctx, "item.0", d)
		log.Printf("[DEBUG] request1 => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.AciSettings.UpdateAciSettingsByID(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] restyResp1 => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateAciSettingsByID", err, restyResp1.String(),
					"Failure at UpdateAciSettingsByID, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateAciSettingsByID", err,
				"Failure at UpdateAciSettingsByID, unexpected response", ""))
			return diags
		}
	}

	return resourceAciSettingsRead(ctx, d, m)
}

func resourceAciSettingsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// NOTE: Function does not perform delete on ISE
	//       Returning empty diags to delete it on Terraform
	return diags
}
func expandRequestAciSettingsUpdateAciSettingsByID(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestAciSettingsUpdateAciSettingsByID {
	request := isegosdk.RequestAciSettingsUpdateAciSettingsByID{}
	request.AciSettings = expandRequestAciSettingsUpdateAciSettingsByIDAciSettings(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAciSettingsUpdateAciSettingsByIDAciSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestAciSettingsUpdateAciSettingsByIDAciSettings {
	request := isegosdk.RequestAciSettingsUpdateAciSettingsByIDAciSettings{}
	if v, ok := d.GetOkExists(key + ".id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".id"))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".enable_aci"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enable_aci"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enable_aci"))) {
		request.EnableAci = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".ip_address_host_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".ip_address_host_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".ip_address_host_name"))) {
		request.IPAddressHostName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".admin_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".admin_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".admin_name"))) {
		request.AdminName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".admin_password"); !isEmptyValue(reflect.ValueOf(d.Get(key+".admin_password"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".admin_password"))) {
		request.AdminPassword = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".aciipaddress"); !isEmptyValue(reflect.ValueOf(d.Get(key+".aciipaddress"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".aciipaddress"))) {
		request.Aciipaddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".aciuser_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".aciuser_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".aciuser_name"))) {
		request.AciuserName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".acipassword"); !isEmptyValue(reflect.ValueOf(d.Get(key+".acipassword"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".acipassword"))) {
		request.Acipassword = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".tenant_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".tenant_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".tenant_name"))) {
		request.TenantName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".l3_route_network"); !isEmptyValue(reflect.ValueOf(d.Get(key+".l3_route_network"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".l3_route_network"))) {
		request.L3RouteNetwork = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".suffix_to_epg"); !isEmptyValue(reflect.ValueOf(d.Get(key+".suffix_to_epg"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".suffix_to_epg"))) {
		request.SuffixToEpg = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".suffix_to_sgt"); !isEmptyValue(reflect.ValueOf(d.Get(key+".suffix_to_sgt"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".suffix_to_sgt"))) {
		request.SuffixToSgt = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".all_sxp_domain"); !isEmptyValue(reflect.ValueOf(d.Get(key+".all_sxp_domain"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".all_sxp_domain"))) {
		request.AllSxpDomain = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".specific_sxp_domain"); !isEmptyValue(reflect.ValueOf(d.Get(key+".specific_sxp_domain"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".specific_sxp_domain"))) {
		request.SpecificSxpDomain = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".specifix_sxp_domain_list"); !isEmptyValue(reflect.ValueOf(d.Get(key+".specifix_sxp_domain_list"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".specifix_sxp_domain_list"))) {
		request.SpecifixSxpDomainList = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(key + ".enable_data_plane"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enable_data_plane"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enable_data_plane"))) {
		request.EnableDataPlane = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".untagged_packet_iepg_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".untagged_packet_iepg_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".untagged_packet_iepg_name"))) {
		request.UntaggedPacketIepgName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".default_sgt_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".default_sgt_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".default_sgt_name"))) {
		request.DefaultSgtName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".enable_elements_limit"); !isEmptyValue(reflect.ValueOf(d.Get(key+".enable_elements_limit"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".enable_elements_limit"))) {
		request.EnableElementsLimit = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".max_num_iepg_from_aci"); !isEmptyValue(reflect.ValueOf(d.Get(key+".max_num_iepg_from_aci"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".max_num_iepg_from_aci"))) {
		request.MaxNumIepgFromAci = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".max_num_sgt_to_aci"); !isEmptyValue(reflect.ValueOf(d.Get(key+".max_num_sgt_to_aci"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".max_num_sgt_to_aci"))) {
		request.MaxNumSgtToAci = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".aci50"); !isEmptyValue(reflect.ValueOf(d.Get(key+".aci50"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".aci50"))) {
		request.Aci50 = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".aci51"); !isEmptyValue(reflect.ValueOf(d.Get(key+".aci51"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".aci51"))) {
		request.Aci51 = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
