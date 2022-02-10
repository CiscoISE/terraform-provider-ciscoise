package ciscoise

import (
	"context"
	"reflect"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAciSettings() *schema.Resource {
	return &schema.Resource{
		Description: `It manages read and update operations on ACISettings.

- This resource allows the client to update ACI settings.
`,

		CreateContext: resourceAciSettingsCreate,
		ReadContext:   resourceAciSettingsRead,
		UpdateContext: resourceAciSettingsUpdate,
		DeleteContext: resourceAciSettingsDelete,
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

						"aci50": &schema.Schema{
							Description: `Enable 5.0 ACI Version`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"aci51": &schema.Schema{
							Description: `Enable 5.1 ACI Version`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"aciipaddress": &schema.Schema{
							Description: `ACI Domain manager Ip Address.`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"acipassword": &schema.Schema{
							Description: `ACI Domain manager Password.`,
							Type:        schema.TypeString,
							Sensitive:   true,
							Computed:    true,
						},
						"aciuser_name": &schema.Schema{
							Description: `ACI Domain manager Username.`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"admin_name": &schema.Schema{
							Description: `ACI Cluster Admin name`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"admin_password": &schema.Schema{
							Description: `ACI Cluster Admin password`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"all_sxp_domain": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"default_sgt_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"enable_aci": &schema.Schema{
							Description: `Enable ACI Integration`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"enable_data_plane": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"enable_elements_limit": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Description: `Resource UUID value`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"ip_address_host_name": &schema.Schema{
							Description: `ACI Cluster IP Address / Host name`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"l3_route_network": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"max_num_iepg_from_aci": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"max_num_sgt_to_aci": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"specific_sxp_domain": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"specifix_sxp_domain_list": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"suffix_to_epg": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"suffix_to_sgt": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"tenant_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"untagged_packet_iepg_name": &schema.Schema{
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
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"aci50": &schema.Schema{
							Description:  `Enable 5.0 ACI Version`,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"aci51": &schema.Schema{
							Description:  `Enable 5.1 ACI Version`,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
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
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"default_sgt_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"enable_aci": &schema.Schema{
							Description:  `Enable ACI Integration`,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"enable_data_plane": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"enable_elements_limit": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"id": &schema.Schema{
							Description: `Resource UUID value`,
							Type:        schema.TypeString,
							Required:    true,
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
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
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
	log.Printf("[DEBUG] Beginning AciSettings create")
	log.Printf("[DEBUG] Missing AciSettings create on Cisco ISE. It will only be create it on Terraform")
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	resourceMap := make(map[string]string)
	vvID := interfaceToString(resourceItem["id"])
	log.Printf("[DEBUG] ID used for update operation %s", vvID)
	request1 := expandRequestAciSettingsUpdateAciSettingsByID(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}
	response1, restyResp1, err := client.AciSettings.UpdateAciSettingsByID(vvID, request1)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
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

	resourceMap["id"] = interfaceToString(resourceItem["id"])
	d.SetId(joinResourceID(resourceMap))
	return resourceAciSettingsRead(ctx, d, m)
}

func resourceAciSettingsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning AciSettings read for id=[%s]", d.Id())
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetAciSettings")

		response1, restyResp1, err := client.AciSettings.GetAciSettings()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenAciSettingsGetAciSettingsItem(response1.AciSettings)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAciSettings response to item",
				err))
			return diags
		}
		if err := d.Set("parameters", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAciSettings response to parameters",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceAciSettingsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning AciSettings update for id=[%s]", d.Id())
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID, _ := resourceMap["id"]
	vvID := vID
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] ID used for update operation %s", vvID)
		request1 := expandRequestAciSettingsUpdateAciSettingsByID(ctx, "parameters.0", d)
		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
		response1, restyResp1, err := client.AciSettings.UpdateAciSettingsByID(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
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
		d.Set("last_updated", getUnixTimeString())
	}

	return resourceAciSettingsRead(ctx, d, m)
}

func resourceAciSettingsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning AciSettings delete for id=[%s]", d.Id())
	var diags diag.Diagnostics
	log.Printf("[DEBUG] Missing AciSettings delete on Cisco ISE. It will only be delete it on Terraform id=[%s]", d.Id())
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
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_aci")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_aci")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_aci")))) {
		request.EnableAci = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_address_host_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_address_host_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_address_host_name")))) {
		request.IPAddressHostName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".admin_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".admin_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".admin_name")))) {
		request.AdminName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".admin_password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".admin_password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".admin_password")))) {
		request.AdminPassword = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".aciipaddress")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".aciipaddress")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".aciipaddress")))) {
		request.Aciipaddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".aciuser_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".aciuser_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".aciuser_name")))) {
		request.AciuserName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".acipassword")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".acipassword")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".acipassword")))) {
		request.Acipassword = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tenant_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tenant_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tenant_name")))) {
		request.TenantName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".l3_route_network")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".l3_route_network")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".l3_route_network")))) {
		request.L3RouteNetwork = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".suffix_to_epg")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".suffix_to_epg")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".suffix_to_epg")))) {
		request.SuffixToEpg = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".suffix_to_sgt")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".suffix_to_sgt")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".suffix_to_sgt")))) {
		request.SuffixToSgt = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".all_sxp_domain")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".all_sxp_domain")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".all_sxp_domain")))) {
		request.AllSxpDomain = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".specific_sxp_domain")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".specific_sxp_domain")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".specific_sxp_domain")))) {
		request.SpecificSxpDomain = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".specifix_sxp_domain_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".specifix_sxp_domain_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".specifix_sxp_domain_list")))) {
		request.SpecifixSxpDomainList = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_data_plane")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_data_plane")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_data_plane")))) {
		request.EnableDataPlane = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".untagged_packet_iepg_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".untagged_packet_iepg_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".untagged_packet_iepg_name")))) {
		request.UntaggedPacketIepgName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".default_sgt_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".default_sgt_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".default_sgt_name")))) {
		request.DefaultSgtName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_elements_limit")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_elements_limit")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_elements_limit")))) {
		request.EnableElementsLimit = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".max_num_iepg_from_aci")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".max_num_iepg_from_aci")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".max_num_iepg_from_aci")))) {
		request.MaxNumIepgFromAci = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".max_num_sgt_to_aci")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".max_num_sgt_to_aci")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".max_num_sgt_to_aci")))) {
		request.MaxNumSgtToAci = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".aci50")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".aci50")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".aci50")))) {
		request.Aci50 = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".aci51")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".aci51")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".aci51")))) {
		request.Aci51 = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
