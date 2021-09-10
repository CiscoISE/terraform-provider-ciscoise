package ciscoise

import (
	"context"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkDevice() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceNetworkDeviceRead,
		Schema: map[string]*schema.Schema{
			"filter": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"filter_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"page": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"sortasc": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sortdsc": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"item_id": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"network_device_group_list": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"network_device_iplist": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"get_ipaddress_exclude": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"ipaddress": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"mask": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"authentication_settings": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"dtls_required": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"enable_key_wrap": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"enable_multi_secret": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"enabled": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"key_encryption_key": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"key_input_format": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"message_authenticator_code_key": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"network_protocol": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"radius_shared_secret": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"second_radius_shared_secret": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"coa_port": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"dtls_dns_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
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
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"profile_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"snmpsettings": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"link_trap_query": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"mac_trap_query": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"originating_policy_services_node": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"polling_interval": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"ro_community": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"version": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"software_version": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"tacacs_settings": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"connect_mode_options": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"shared_secret": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"trustsecsettings": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"device_authentication_settings": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"sga_device_id": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"sga_device_password": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"device_configuration_deployment": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"enable_mode_password": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"exec_mode_password": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"exec_mode_username": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"include_when_deploying_sgt_updates": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
											},
										},
									},
									"push_id_support": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"sga_notification_and_updates": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"coa_source_host": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"downlaod_environment_data_every_x_seconds": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"downlaod_peer_authorization_policy_every_x_seconds": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"download_sga_cllists_every_x_seconds": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"other_sga_devices_to_trust_this_device": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"re_authentication_every_x_seconds": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"send_configuration_to_device": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"send_configuration_to_device_using": &schema.Schema{
													Type:     schema.TypeString,
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
			"item_name": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"network_device_group_list": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"network_device_iplist": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"get_ipaddress_exclude": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"ipaddress": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"mask": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"authentication_settings": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"dtls_required": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"enable_key_wrap": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"enable_multi_secret": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"enabled": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"key_encryption_key": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"key_input_format": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"message_authenticator_code_key": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"network_protocol": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"radius_shared_secret": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"second_radius_shared_secret": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"coa_port": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"dtls_dns_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
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
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"profile_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"snmpsettings": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"link_trap_query": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"mac_trap_query": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"originating_policy_services_node": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"polling_interval": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"ro_community": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"version": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"software_version": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"tacacs_settings": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"connect_mode_options": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"shared_secret": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"trustsecsettings": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"device_authentication_settings": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"sga_device_id": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"sga_device_password": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"device_configuration_deployment": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"enable_mode_password": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"exec_mode_password": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"exec_mode_username": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"include_when_deploying_sgt_updates": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
											},
										},
									},
									"push_id_support": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"sga_notification_and_updates": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"coa_source_host": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"downlaod_environment_data_every_x_seconds": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"downlaod_peer_authorization_policy_every_x_seconds": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"download_sga_cllists_every_x_seconds": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"other_sga_devices_to_trust_this_device": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"re_authentication_every_x_seconds": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"send_configuration_to_device": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"send_configuration_to_device_using": &schema.Schema{
													Type:     schema.TypeString,
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
			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
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
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceNetworkDeviceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vPage, okPage := d.GetOk("page")
	vSize, okSize := d.GetOk("size")
	vSortasc, okSortasc := d.GetOk("sortasc")
	vSortdsc, okSortdsc := d.GetOk("sortdsc")
	vFilter, okFilter := d.GetOk("filter")
	vFilterType, okFilterType := d.GetOk("filter_type")
	vName, okName := d.GetOk("name")
	vID, okID := d.GetOk("id")

	method1 := []bool{okPage, okSize, okSortasc, okSortdsc, okFilter, okFilterType}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)
	method3 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 3 %v", method3)

	selectedMethod := pickMethod([][]bool{method1, method2, method3})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetNetworkDevice")
		queryParams1 := isegosdk.GetNetworkDeviceQueryParams{}

		if okPage {
			queryParams1.Page = vPage.(int)
		}
		if okSize {
			queryParams1.Size = vSize.(int)
		}
		if okSortasc {
			queryParams1.Sortasc = vSortasc.(string)
		}
		if okSortdsc {
			queryParams1.Sortdsc = vSortdsc.(string)
		}
		if okFilter {
			queryParams1.Filter = interfaceToSliceString(vFilter)
		}
		if okFilterType {
			queryParams1.FilterType = vFilterType.(string)
		}

		response1, _, err := client.NetworkDevice.GetNetworkDevice(&queryParams1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetNetworkDevice", err,
				"Failure at GetNetworkDevice, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		var items1 []isegosdk.ResponseNetworkDeviceGetNetworkDeviceSearchResultResources
		for response1.SearchResult != nil && response1.SearchResult.Resources != nil && len(*response1.SearchResult.Resources) > 0 {
			items1 = append(items1, *response1.SearchResult.Resources...)
			if response1.SearchResult.NextPage != nil && response1.SearchResult.NextPage.Rel == "next" {
				href := response1.SearchResult.NextPage.Href
				page, size, err := getNextPageAndSizeParams(href)
				if err != nil {
					break
				}
				queryParams1.Page = page
				queryParams1.Size = size
				response1, _, err = client.NetworkDevice.GetNetworkDevice(&queryParams1)
				if err != nil {
					break
				}
				// All is good, continue to the next page
				continue
			}
			// Does not have next page finish iteration
			break
		}
		vItems1 := flattenNetworkDeviceGetNetworkDeviceItems(&items1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNetworkDevice response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: GetNetworkDeviceByName")
		vvName := vName.(string)

		response2, _, err := client.NetworkDevice.GetNetworkDeviceByName(vvName)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetNetworkDeviceByName", err,
				"Failure at GetNetworkDeviceByName, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

		vItemName2 := flattenNetworkDeviceGetNetworkDeviceByNameItemName(response2.NetworkDevice)
		if err := d.Set("item_name", vItemName2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNetworkDeviceByName response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 3 {
		log.Printf("[DEBUG] Selected method 3: GetNetworkDeviceByID")
		vvID := vID.(string)

		response3, _, err := client.NetworkDevice.GetNetworkDeviceByID(vvID)

		if err != nil || response3 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetNetworkDeviceByID", err,
				"Failure at GetNetworkDeviceByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response3)

		vItemID3 := flattenNetworkDeviceGetNetworkDeviceByIDItemID(response3.NetworkDevice)
		if err := d.Set("item_id", vItemID3); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNetworkDeviceByID response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenNetworkDeviceGetNetworkDeviceItems(items *[]isegosdk.ResponseNetworkDeviceGetNetworkDeviceSearchResultResources) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItem["description"] = item.Description
		respItem["link"] = flattenNetworkDeviceGetNetworkDeviceItemsLink(item.Link)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenNetworkDeviceGetNetworkDeviceItemsLink(item *isegosdk.ResponseNetworkDeviceGetNetworkDeviceSearchResultResourcesLink) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["rel"] = item.Rel
	respItem["href"] = item.Href
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNetworkDeviceGetNetworkDeviceByNameItemName(item *isegosdk.ResponseNetworkDeviceGetNetworkDeviceByNameNetworkDevice) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["description"] = item.Description
	respItem["authentication_settings"] = flattenNetworkDeviceGetNetworkDeviceByNameItemNameAuthenticationSettings(item.AuthenticationSettings)
	respItem["snmpsettings"] = flattenNetworkDeviceGetNetworkDeviceByNameItemNameSNMPsettings(item.SNMPsettings)
	respItem["trustsecsettings"] = flattenNetworkDeviceGetNetworkDeviceByNameItemNameTrustsecsettings(item.Trustsecsettings)
	respItem["tacacs_settings"] = flattenNetworkDeviceGetNetworkDeviceByNameItemNameTacacsSettings(item.TacacsSettings)
	respItem["profile_name"] = item.ProfileName
	respItem["model_name"] = item.ModelName
	respItem["software_version"] = item.SoftwareVersion
	respItem["coa_port"] = item.CoaPort
	respItem["dtls_dns_name"] = item.DtlsDNSName
	respItem["network_device_iplist"] = flattenNetworkDeviceGetNetworkDeviceByNameItemNameNetworkDeviceIPList(item.NetworkDeviceIPList)
	respItem["network_device_group_list"] = item.NetworkDeviceGroupList
	respItem["link"] = flattenNetworkDeviceGetNetworkDeviceByNameItemNameLink(item.Link)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenNetworkDeviceGetNetworkDeviceByNameItemNameAuthenticationSettings(item *isegosdk.ResponseNetworkDeviceGetNetworkDeviceByNameNetworkDeviceAuthenticationSettings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["network_protocol"] = item.NetworkProtocol
	respItem["second_radius_shared_secret"] = item.SecondRadiusSharedSecret
	respItem["radius_shared_secret"] = item.RadiusSharedSecret
	respItem["enable_key_wrap"] = item.EnableKeyWrap
	respItem["enabled"] = item.Enabled
	respItem["dtls_required"] = item.DtlsRequired
	respItem["enable_multi_secret"] = item.EnableMultiSecret
	respItem["key_encryption_key"] = item.KeyEncryptionKey
	respItem["message_authenticator_code_key"] = item.MessageAuthenticatorCodeKey
	respItem["key_input_format"] = item.KeyInputFormat

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNetworkDeviceGetNetworkDeviceByNameItemNameSNMPsettings(item *isegosdk.ResponseNetworkDeviceGetNetworkDeviceByNameNetworkDeviceSNMPsettings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["version"] = item.Version
	respItem["ro_community"] = item.RoCommunity
	respItem["polling_interval"] = item.PollingInterval
	respItem["link_trap_query"] = item.LinkTrapQuery
	respItem["mac_trap_query"] = item.MacTrapQuery
	respItem["originating_policy_services_node"] = item.OriginatingPolicyServicesNode

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNetworkDeviceGetNetworkDeviceByNameItemNameTrustsecsettings(item *isegosdk.ResponseNetworkDeviceGetNetworkDeviceByNameNetworkDeviceTrustsecsettings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["device_authentication_settings"] = flattenNetworkDeviceGetNetworkDeviceByNameItemNameTrustsecsettingsDeviceAuthenticationSettings(item.DeviceAuthenticationSettings)
	respItem["sga_notification_and_updates"] = flattenNetworkDeviceGetNetworkDeviceByNameItemNameTrustsecsettingsSgaNotificationAndUpdates(item.SgaNotificationAndUpdates)
	respItem["device_configuration_deployment"] = flattenNetworkDeviceGetNetworkDeviceByNameItemNameTrustsecsettingsDeviceConfigurationDeployment(item.DeviceConfigurationDeployment)
	respItem["push_id_support"] = item.PushIDSupport

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNetworkDeviceGetNetworkDeviceByNameItemNameTrustsecsettingsDeviceAuthenticationSettings(item *isegosdk.ResponseNetworkDeviceGetNetworkDeviceByNameNetworkDeviceTrustsecsettingsDeviceAuthenticationSettings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["sga_device_id"] = item.SgaDeviceID
	respItem["sga_device_password"] = item.SgaDevicePassword

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNetworkDeviceGetNetworkDeviceByNameItemNameTrustsecsettingsSgaNotificationAndUpdates(item *isegosdk.ResponseNetworkDeviceGetNetworkDeviceByNameNetworkDeviceTrustsecsettingsSgaNotificationAndUpdates) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["downlaod_environment_data_every_x_seconds"] = item.DownlaodEnvironmentDataEveryXSeconds
	respItem["downlaod_peer_authorization_policy_every_x_seconds"] = item.DownlaodPeerAuthorizationPolicyEveryXSeconds
	respItem["re_authentication_every_x_seconds"] = item.ReAuthenticationEveryXSeconds
	respItem["download_sga_cllists_every_x_seconds"] = item.DownloadSgACLListsEveryXSeconds
	respItem["other_sga_devices_to_trust_this_device"] = item.OtherSgADevicesToTrustThisDevice
	respItem["send_configuration_to_device"] = item.SendConfigurationToDevice
	respItem["send_configuration_to_device_using"] = item.SendConfigurationToDeviceUsing
	respItem["coa_source_host"] = item.CoaSourceHost

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNetworkDeviceGetNetworkDeviceByNameItemNameTrustsecsettingsDeviceConfigurationDeployment(item *isegosdk.ResponseNetworkDeviceGetNetworkDeviceByNameNetworkDeviceTrustsecsettingsDeviceConfigurationDeployment) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["include_when_deploying_sgt_updates"] = item.IncludeWhenDeployingSgtUpdates
	respItem["enable_mode_password"] = item.EnableModePassword
	respItem["exec_mode_password"] = item.ExecModePassword
	respItem["exec_mode_username"] = item.ExecModeUsername

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNetworkDeviceGetNetworkDeviceByNameItemNameTacacsSettings(item *isegosdk.ResponseNetworkDeviceGetNetworkDeviceByNameNetworkDeviceTacacsSettings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["shared_secret"] = item.SharedSecret
	respItem["connect_mode_options"] = item.ConnectModeOptions

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNetworkDeviceGetNetworkDeviceByNameItemNameNetworkDeviceIPList(items *[]isegosdk.ResponseNetworkDeviceGetNetworkDeviceByNameNetworkDeviceNetworkDeviceIPList) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ipaddress"] = item.IPaddress
		respItem["mask"] = item.Mask
		respItem["get_ipaddress_exclude"] = item.GetIPaddressExclude
	}
	return respItems

}

func flattenNetworkDeviceGetNetworkDeviceByNameItemNameLink(item *isegosdk.ResponseNetworkDeviceGetNetworkDeviceByNameNetworkDeviceLink) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["rel"] = item.Rel
	respItem["href"] = item.Href
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNetworkDeviceGetNetworkDeviceByIDItemID(item *isegosdk.ResponseNetworkDeviceGetNetworkDeviceByIDNetworkDevice) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["description"] = item.Description
	respItem["authentication_settings"] = flattenNetworkDeviceGetNetworkDeviceByIDItemIDAuthenticationSettings(item.AuthenticationSettings)
	respItem["snmpsettings"] = flattenNetworkDeviceGetNetworkDeviceByIDItemIDSNMPsettings(item.SNMPsettings)
	respItem["trustsecsettings"] = flattenNetworkDeviceGetNetworkDeviceByIDItemIDTrustsecsettings(item.Trustsecsettings)
	respItem["tacacs_settings"] = flattenNetworkDeviceGetNetworkDeviceByIDItemIDTacacsSettings(item.TacacsSettings)
	respItem["profile_name"] = item.ProfileName
	respItem["coa_port"] = item.CoaPort
	respItem["dtls_dns_name"] = item.DtlsDNSName
	respItem["model_name"] = item.ModelName
	respItem["software_version"] = item.SoftwareVersion
	respItem["network_device_iplist"] = flattenNetworkDeviceGetNetworkDeviceByIDItemIDNetworkDeviceIPList(item.NetworkDeviceIPList)
	respItem["network_device_group_list"] = item.NetworkDeviceGroupList
	respItem["link"] = flattenNetworkDeviceGetNetworkDeviceByIDItemIDLink(item.Link)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenNetworkDeviceGetNetworkDeviceByIDItemIDAuthenticationSettings(item *isegosdk.ResponseNetworkDeviceGetNetworkDeviceByIDNetworkDeviceAuthenticationSettings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["network_protocol"] = item.NetworkProtocol
	respItem["second_radius_shared_secret"] = item.SecondRadiusSharedSecret
	respItem["radius_shared_secret"] = item.RadiusSharedSecret
	respItem["enable_key_wrap"] = item.EnableKeyWrap
	respItem["enabled"] = item.Enabled
	respItem["dtls_required"] = item.DtlsRequired
	respItem["enable_multi_secret"] = item.EnableMultiSecret
	respItem["key_encryption_key"] = item.KeyEncryptionKey
	respItem["message_authenticator_code_key"] = item.MessageAuthenticatorCodeKey
	respItem["key_input_format"] = item.KeyInputFormat

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNetworkDeviceGetNetworkDeviceByIDItemIDSNMPsettings(item *isegosdk.ResponseNetworkDeviceGetNetworkDeviceByIDNetworkDeviceSNMPsettings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["version"] = item.Version
	respItem["ro_community"] = item.RoCommunity
	respItem["polling_interval"] = item.PollingInterval
	respItem["link_trap_query"] = item.LinkTrapQuery
	respItem["mac_trap_query"] = item.MacTrapQuery
	respItem["originating_policy_services_node"] = item.OriginatingPolicyServicesNode

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNetworkDeviceGetNetworkDeviceByIDItemIDTrustsecsettings(item *isegosdk.ResponseNetworkDeviceGetNetworkDeviceByIDNetworkDeviceTrustsecsettings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["device_authentication_settings"] = flattenNetworkDeviceGetNetworkDeviceByIDItemIDTrustsecsettingsDeviceAuthenticationSettings(item.DeviceAuthenticationSettings)
	respItem["sga_notification_and_updates"] = flattenNetworkDeviceGetNetworkDeviceByIDItemIDTrustsecsettingsSgaNotificationAndUpdates(item.SgaNotificationAndUpdates)
	respItem["device_configuration_deployment"] = flattenNetworkDeviceGetNetworkDeviceByIDItemIDTrustsecsettingsDeviceConfigurationDeployment(item.DeviceConfigurationDeployment)
	respItem["push_id_support"] = item.PushIDSupport

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNetworkDeviceGetNetworkDeviceByIDItemIDTrustsecsettingsDeviceAuthenticationSettings(item *isegosdk.ResponseNetworkDeviceGetNetworkDeviceByIDNetworkDeviceTrustsecsettingsDeviceAuthenticationSettings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["sga_device_id"] = item.SgaDeviceID
	respItem["sga_device_password"] = item.SgaDevicePassword

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNetworkDeviceGetNetworkDeviceByIDItemIDTrustsecsettingsSgaNotificationAndUpdates(item *isegosdk.ResponseNetworkDeviceGetNetworkDeviceByIDNetworkDeviceTrustsecsettingsSgaNotificationAndUpdates) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["downlaod_environment_data_every_x_seconds"] = item.DownlaodEnvironmentDataEveryXSeconds
	respItem["downlaod_peer_authorization_policy_every_x_seconds"] = item.DownlaodPeerAuthorizationPolicyEveryXSeconds
	respItem["re_authentication_every_x_seconds"] = item.ReAuthenticationEveryXSeconds
	respItem["download_sga_cllists_every_x_seconds"] = item.DownloadSgACLListsEveryXSeconds
	respItem["other_sga_devices_to_trust_this_device"] = item.OtherSgADevicesToTrustThisDevice
	respItem["send_configuration_to_device"] = item.SendConfigurationToDevice
	respItem["send_configuration_to_device_using"] = item.SendConfigurationToDeviceUsing
	respItem["coa_source_host"] = item.CoaSourceHost

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNetworkDeviceGetNetworkDeviceByIDItemIDTrustsecsettingsDeviceConfigurationDeployment(item *isegosdk.ResponseNetworkDeviceGetNetworkDeviceByIDNetworkDeviceTrustsecsettingsDeviceConfigurationDeployment) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["include_when_deploying_sgt_updates"] = item.IncludeWhenDeployingSgtUpdates
	respItem["enable_mode_password"] = item.EnableModePassword
	respItem["exec_mode_password"] = item.ExecModePassword
	respItem["exec_mode_username"] = item.ExecModeUsername

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNetworkDeviceGetNetworkDeviceByIDItemIDTacacsSettings(item *isegosdk.ResponseNetworkDeviceGetNetworkDeviceByIDNetworkDeviceTacacsSettings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["shared_secret"] = item.SharedSecret
	respItem["connect_mode_options"] = item.ConnectModeOptions

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNetworkDeviceGetNetworkDeviceByIDItemIDNetworkDeviceIPList(items *[]isegosdk.ResponseNetworkDeviceGetNetworkDeviceByIDNetworkDeviceNetworkDeviceIPList) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ipaddress"] = item.IPaddress
		respItem["mask"] = item.Mask
		respItem["get_ipaddress_exclude"] = item.GetIPaddressExclude
	}
	return respItems

}

func flattenNetworkDeviceGetNetworkDeviceByIDItemIDLink(item *isegosdk.ResponseNetworkDeviceGetNetworkDeviceByIDNetworkDeviceLink) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["rel"] = item.Rel
	respItem["href"] = item.Href
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}
