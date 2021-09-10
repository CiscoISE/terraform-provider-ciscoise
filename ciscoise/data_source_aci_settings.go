package ciscoise

import (
	"context"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAciSettings() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceAciSettingsRead,
		Schema: map[string]*schema.Schema{
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"aci50": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"aci51": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"aciipaddress": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"acipassword": &schema.Schema{
							Type:      schema.TypeString,
							Sensitive: true,
							Computed:  true,
						},
						"aciuser_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"admin_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"admin_password": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"all_sxp_domain": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"default_sgt_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"enable_aci": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"enable_data_plane": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"enable_elements_limit": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip_address_host_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
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
							Type:     schema.TypeBool,
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
		},
	}
}

func dataSourceAciSettingsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetAciSettings")

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
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenAciSettingsGetAciSettingsItem(item *isegosdk.ResponseAciSettingsGetAciSettingsAciSettings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["enable_aci"] = item.EnableAci
	respItem["ip_address_host_name"] = item.IPAddressHostName
	respItem["admin_name"] = item.AdminName
	respItem["admin_password"] = item.AdminPassword
	respItem["aciipaddress"] = item.Aciipaddress
	respItem["aciuser_name"] = item.AciuserName
	respItem["acipassword"] = item.Acipassword
	respItem["tenant_name"] = item.TenantName
	respItem["l3_route_network"] = item.L3RouteNetwork
	respItem["suffix_to_epg"] = item.SuffixToEpg
	respItem["suffix_to_sgt"] = item.SuffixToSgt
	respItem["all_sxp_domain"] = item.AllSxpDomain
	respItem["specific_sxp_domain"] = item.SpecificSxpDomain
	respItem["specifix_sxp_domain_list"] = item.SpecifixSxpDomainList
	respItem["enable_data_plane"] = item.EnableDataPlane
	respItem["untagged_packet_iepg_name"] = item.UntaggedPacketIepgName
	respItem["default_sgt_name"] = item.DefaultSgtName
	respItem["enable_elements_limit"] = item.EnableElementsLimit
	respItem["max_num_iepg_from_aci"] = item.MaxNumIepgFromAci
	respItem["max_num_sgt_to_aci"] = item.MaxNumSgtToAci
	respItem["aci50"] = item.Aci50
	respItem["aci51"] = item.Aci51
	return []map[string]interface{}{
		respItem,
	}
}
