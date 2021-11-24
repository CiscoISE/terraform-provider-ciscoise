package ciscoise

import (
	"context"

	"log"

	isegosdk "ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAciSettings() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on ACISettings.

- This data source allows the client to get ACI Settings.
`,

		ReadContext: dataSourceAciSettingsRead,
		Schema: map[string]*schema.Schema{
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
		},
	}
}

func dataSourceAciSettingsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetAciSettings")

		response1, restyResp1, err := client.AciSettings.GetAciSettings()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetAciSettings", err,
				"Failure at GetAciSettings, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

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
	respItem["enable_aci"] = boolPtrToString(item.EnableAci)
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
	respItem["all_sxp_domain"] = boolPtrToString(item.AllSxpDomain)
	respItem["specific_sxp_domain"] = boolPtrToString(item.SpecificSxpDomain)
	respItem["specifix_sxp_domain_list"] = item.SpecifixSxpDomainList
	respItem["enable_data_plane"] = boolPtrToString(item.EnableDataPlane)
	respItem["untagged_packet_iepg_name"] = item.UntaggedPacketIepgName
	respItem["default_sgt_name"] = item.DefaultSgtName
	respItem["enable_elements_limit"] = boolPtrToString(item.EnableElementsLimit)
	respItem["max_num_iepg_from_aci"] = item.MaxNumIepgFromAci
	respItem["max_num_sgt_to_aci"] = item.MaxNumSgtToAci
	respItem["aci50"] = boolPtrToString(item.Aci50)
	respItem["aci51"] = boolPtrToString(item.Aci51)
	return []map[string]interface{}{
		respItem,
	}
}
