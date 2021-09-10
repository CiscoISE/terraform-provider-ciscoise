package ciscoise

import (
	"context"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAuthorizationProfile() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceAuthorizationProfileRead,
		Schema: map[string]*schema.Schema{
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
			"item_id": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"access_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"acl": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"advanced_attributes": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"left_hand_side_dictionary_attribue": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"advanced_attribute_value_type": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"attribute_name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"dictionary_name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"value": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"right_hand_side_attribue_value": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"advanced_attribute_value_type": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"attribute_name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"dictionary_name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"value": &schema.Schema{
													Type:     schema.TypeString,
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
							Computed: true,
						},
						"airespace_acl": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"airespace_ipv6_acl": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"asa_vpn": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"authz_profile_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"auto_smart_port": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"avc_profile": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"dacl_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"easywired_session_candidate": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"interface_template": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ipv6_acl_filter": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ipv6_dacl_name": &schema.Schema{
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
						"mac_sec_policy": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"neat": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"profile_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"reauth": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"connectivity": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"timer": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"service_template": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"track_movement": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"vlan": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"name_id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"tag_id": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"voice_domain_permission": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"web_auth": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"web_redirection": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"web_redirection_type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"acl": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"display_certificates_renewal_messages": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"portal_name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"static_iphost_name_fqd_n": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
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

						"access_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"acl": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"advanced_attributes": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"left_hand_side_dictionary_attribue": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"advanced_attribute_value_type": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"attribute_name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"dictionary_name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"value": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"right_hand_side_attribue_value": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"advanced_attribute_value_type": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"attribute_name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"dictionary_name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"value": &schema.Schema{
													Type:     schema.TypeString,
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
							Computed: true,
						},
						"airespace_acl": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"airespace_ipv6_acl": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"asa_vpn": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"authz_profile_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"auto_smart_port": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"avc_profile": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"dacl_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"easywired_session_candidate": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"interface_template": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ipv6_acl_filter": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ipv6_dacl_name": &schema.Schema{
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
						"mac_sec_policy": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"neat": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"profile_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"reauth": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"connectivity": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"timer": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"service_template": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"track_movement": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"vlan": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"name_id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"tag_id": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"voice_domain_permission": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"web_auth": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"web_redirection": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"web_redirection_type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"acl": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"display_certificates_renewal_messages": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"portal_name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"static_iphost_name_fqd_n": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
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

func dataSourceAuthorizationProfileRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vPage, okPage := d.GetOk("page")
	vSize, okSize := d.GetOk("size")
	vName, okName := d.GetOk("name")
	vID, okID := d.GetOk("id")

	method1 := []bool{okPage, okSize}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)
	method3 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 3 %v", method3)

	selectedMethod := pickMethod([][]bool{method1, method2, method3})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetAuthorizationProfiles")
		queryParams1 := isegosdk.GetAuthorizationProfilesQueryParams{}

		if okPage {
			queryParams1.Page = vPage.(int)
		}
		if okSize {
			queryParams1.Size = vSize.(int)
		}

		response1, _, err := client.AuthorizationProfile.GetAuthorizationProfiles(&queryParams1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetAuthorizationProfiles", err,
				"Failure at GetAuthorizationProfiles, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		var items1 []isegosdk.ResponseAuthorizationProfileGetAuthorizationProfilesSearchResultResources
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
				response1, _, err = client.AuthorizationProfile.GetAuthorizationProfiles(&queryParams1)
				if err != nil {
					break
				}
				// All is good, continue to the next page
				continue
			}
			// Does not have next page finish iteration
			break
		}
		vItems1 := flattenAuthorizationProfileGetAuthorizationProfilesItems(&items1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAuthorizationProfiles response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: GetAuthorizationProfileByName")
		vvName := vName.(string)

		response2, _, err := client.AuthorizationProfile.GetAuthorizationProfileByName(vvName)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetAuthorizationProfileByName", err,
				"Failure at GetAuthorizationProfileByName, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

		vItemName2 := flattenAuthorizationProfileGetAuthorizationProfileByNameItemName(response2.AuthorizationProfile)
		if err := d.Set("item_name", vItemName2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAuthorizationProfileByName response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 3 {
		log.Printf("[DEBUG] Selected method 3: GetAuthorizationProfileByID")
		vvID := vID.(string)

		response3, _, err := client.AuthorizationProfile.GetAuthorizationProfileByID(vvID)

		if err != nil || response3 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetAuthorizationProfileByID", err,
				"Failure at GetAuthorizationProfileByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response3)

		vItemID3 := flattenAuthorizationProfileGetAuthorizationProfileByIDItemID(response3.AuthorizationProfile)
		if err := d.Set("item_id", vItemID3); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAuthorizationProfileByID response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenAuthorizationProfileGetAuthorizationProfilesItems(items *[]isegosdk.ResponseAuthorizationProfileGetAuthorizationProfilesSearchResultResources) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItem["description"] = item.Description
		respItem["link"] = flattenAuthorizationProfileGetAuthorizationProfilesItemsLink(item.Link)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenAuthorizationProfileGetAuthorizationProfilesItemsLink(item *isegosdk.ResponseAuthorizationProfileGetAuthorizationProfilesSearchResultResourcesLink) []map[string]interface{} {
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

func flattenAuthorizationProfileGetAuthorizationProfileByNameItemName(item *isegosdk.ResponseAuthorizationProfileGetAuthorizationProfileByNameAuthorizationProfile) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["description"] = item.Description
	respItem["advanced_attributes"] = flattenAuthorizationProfileGetAuthorizationProfileByNameItemNameAdvancedAttributes(item.AdvancedAttributes)
	respItem["access_type"] = item.AccessType
	respItem["authz_profile_type"] = item.AuthzProfileType
	respItem["vlan"] = flattenAuthorizationProfileGetAuthorizationProfileByNameItemNameVLAN(item.VLAN)
	respItem["reauth"] = flattenAuthorizationProfileGetAuthorizationProfileByNameItemNameReauth(item.Reauth)
	respItem["airespace_acl"] = item.AirespaceACL
	respItem["airespace_ipv6_acl"] = item.AirespaceIPv6ACL
	respItem["web_redirection"] = flattenAuthorizationProfileGetAuthorizationProfileByNameItemNameWebRedirection(item.WebRedirection)
	respItem["acl"] = item.ACL
	respItem["track_movement"] = item.TrackMovement
	respItem["agentless_posture"] = item.AgentlessPosture
	respItem["service_template"] = item.ServiceTemplate
	respItem["easywired_session_candidate"] = item.EasywiredSessionCandidate
	respItem["dacl_name"] = item.DaclName
	respItem["voice_domain_permission"] = item.VoiceDomainPermission
	respItem["neat"] = item.Neat
	respItem["web_auth"] = item.WebAuth
	respItem["auto_smart_port"] = item.AutoSmartPort
	respItem["interface_template"] = item.InterfaceTemplate
	respItem["ipv6_acl_filter"] = item.IPv6ACLFilter
	respItem["avc_profile"] = item.AvcProfile
	respItem["mac_sec_policy"] = item.MacSecPolicy
	respItem["asa_vpn"] = item.AsaVpn
	respItem["profile_name"] = item.ProfileName
	respItem["ipv6_dacl_name"] = item.IPv6DaclName
	respItem["link"] = flattenAuthorizationProfileGetAuthorizationProfileByNameItemNameLink(item.Link)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenAuthorizationProfileGetAuthorizationProfileByNameItemNameAdvancedAttributes(items *[]isegosdk.ResponseAuthorizationProfileGetAuthorizationProfileByNameAuthorizationProfileAdvancedAttributes) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["left_hand_side_dictionary_attribue"] = flattenAuthorizationProfileGetAuthorizationProfileByNameItemNameAdvancedAttributesLeftHandSideDictionaryAttribue(item.LeftHandSideDictionaryAttribue)
		respItem["right_hand_side_attribue_value"] = flattenAuthorizationProfileGetAuthorizationProfileByNameItemNameAdvancedAttributesRightHandSideAttribueValue(item.RightHandSideAttribueValue)
	}
	return respItems

}

func flattenAuthorizationProfileGetAuthorizationProfileByNameItemNameAdvancedAttributesLeftHandSideDictionaryAttribue(item *isegosdk.ResponseAuthorizationProfileGetAuthorizationProfileByNameAuthorizationProfileAdvancedAttributesLeftHandSideDictionaryAttribue) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["advanced_attribute_value_type"] = item.AdvancedAttributeValueType
	respItem["dictionary_name"] = item.DictionaryName
	respItem["attribute_name"] = item.AttributeName
	respItem["value"] = item.Value

	return []map[string]interface{}{
		respItem,
	}

}

func flattenAuthorizationProfileGetAuthorizationProfileByNameItemNameAdvancedAttributesRightHandSideAttribueValue(item *isegosdk.ResponseAuthorizationProfileGetAuthorizationProfileByNameAuthorizationProfileAdvancedAttributesRightHandSideAttribueValue) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["advanced_attribute_value_type"] = item.AdvancedAttributeValueType
	respItem["dictionary_name"] = item.DictionaryName
	respItem["attribute_name"] = item.AttributeName
	respItem["value"] = item.Value

	return []map[string]interface{}{
		respItem,
	}

}

func flattenAuthorizationProfileGetAuthorizationProfileByNameItemNameVLAN(item *isegosdk.ResponseAuthorizationProfileGetAuthorizationProfileByNameAuthorizationProfileVLAN) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["name_id"] = item.NameID
	respItem["tag_id"] = item.TagID

	return []map[string]interface{}{
		respItem,
	}

}

func flattenAuthorizationProfileGetAuthorizationProfileByNameItemNameReauth(item *isegosdk.ResponseAuthorizationProfileGetAuthorizationProfileByNameAuthorizationProfileReauth) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["timer"] = item.Timer
	respItem["connectivity"] = item.Connectivity

	return []map[string]interface{}{
		respItem,
	}

}

func flattenAuthorizationProfileGetAuthorizationProfileByNameItemNameWebRedirection(item *isegosdk.ResponseAuthorizationProfileGetAuthorizationProfileByNameAuthorizationProfileWebRedirection) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["web_redirection_type"] = item.WebRedirectionType
	respItem["acl"] = item.ACL
	respItem["portal_name"] = item.PortalName
	respItem["static_iphost_name_fqd_n"] = item.StaticIPHostNameFQDN
	respItem["display_certificates_renewal_messages"] = item.DisplayCertificatesRenewalMessages

	return []map[string]interface{}{
		respItem,
	}

}

func flattenAuthorizationProfileGetAuthorizationProfileByNameItemNameLink(item *isegosdk.ResponseAuthorizationProfileGetAuthorizationProfileByNameAuthorizationProfileLink) []map[string]interface{} {
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

func flattenAuthorizationProfileGetAuthorizationProfileByIDItemID(item *isegosdk.ResponseAuthorizationProfileGetAuthorizationProfileByIDAuthorizationProfile) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["description"] = item.Description
	respItem["advanced_attributes"] = flattenAuthorizationProfileGetAuthorizationProfileByIDItemIDAdvancedAttributes(item.AdvancedAttributes)
	respItem["access_type"] = item.AccessType
	respItem["authz_profile_type"] = item.AuthzProfileType
	respItem["vlan"] = flattenAuthorizationProfileGetAuthorizationProfileByIDItemIDVLAN(item.VLAN)
	respItem["reauth"] = flattenAuthorizationProfileGetAuthorizationProfileByIDItemIDReauth(item.Reauth)
	respItem["airespace_acl"] = item.AirespaceACL
	respItem["airespace_ipv6_acl"] = item.AirespaceIPv6ACL
	respItem["web_redirection"] = flattenAuthorizationProfileGetAuthorizationProfileByIDItemIDWebRedirection(item.WebRedirection)
	respItem["acl"] = item.ACL
	respItem["track_movement"] = item.TrackMovement
	respItem["agentless_posture"] = item.AgentlessPosture
	respItem["service_template"] = item.ServiceTemplate
	respItem["easywired_session_candidate"] = item.EasywiredSessionCandidate
	respItem["dacl_name"] = item.DaclName
	respItem["voice_domain_permission"] = item.VoiceDomainPermission
	respItem["neat"] = item.Neat
	respItem["web_auth"] = item.WebAuth
	respItem["auto_smart_port"] = item.AutoSmartPort
	respItem["interface_template"] = item.InterfaceTemplate
	respItem["ipv6_acl_filter"] = item.IPv6ACLFilter
	respItem["avc_profile"] = item.AvcProfile
	respItem["mac_sec_policy"] = item.MacSecPolicy
	respItem["asa_vpn"] = item.AsaVpn
	respItem["profile_name"] = item.ProfileName
	respItem["ipv6_dacl_name"] = item.IPv6DaclName
	respItem["link"] = flattenAuthorizationProfileGetAuthorizationProfileByIDItemIDLink(item.Link)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenAuthorizationProfileGetAuthorizationProfileByIDItemIDAdvancedAttributes(items *[]isegosdk.ResponseAuthorizationProfileGetAuthorizationProfileByIDAuthorizationProfileAdvancedAttributes) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["left_hand_side_dictionary_attribue"] = flattenAuthorizationProfileGetAuthorizationProfileByIDItemIDAdvancedAttributesLeftHandSideDictionaryAttribue(item.LeftHandSideDictionaryAttribue)
		respItem["right_hand_side_attribue_value"] = flattenAuthorizationProfileGetAuthorizationProfileByIDItemIDAdvancedAttributesRightHandSideAttribueValue(item.RightHandSideAttribueValue)
	}
	return respItems

}

func flattenAuthorizationProfileGetAuthorizationProfileByIDItemIDAdvancedAttributesLeftHandSideDictionaryAttribue(item *isegosdk.ResponseAuthorizationProfileGetAuthorizationProfileByIDAuthorizationProfileAdvancedAttributesLeftHandSideDictionaryAttribue) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["advanced_attribute_value_type"] = item.AdvancedAttributeValueType
	respItem["dictionary_name"] = item.DictionaryName
	respItem["attribute_name"] = item.AttributeName
	respItem["value"] = item.Value

	return []map[string]interface{}{
		respItem,
	}

}

func flattenAuthorizationProfileGetAuthorizationProfileByIDItemIDAdvancedAttributesRightHandSideAttribueValue(item *isegosdk.ResponseAuthorizationProfileGetAuthorizationProfileByIDAuthorizationProfileAdvancedAttributesRightHandSideAttribueValue) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["advanced_attribute_value_type"] = item.AdvancedAttributeValueType
	respItem["dictionary_name"] = item.DictionaryName
	respItem["attribute_name"] = item.AttributeName
	respItem["value"] = item.Value

	return []map[string]interface{}{
		respItem,
	}

}

func flattenAuthorizationProfileGetAuthorizationProfileByIDItemIDVLAN(item *isegosdk.ResponseAuthorizationProfileGetAuthorizationProfileByIDAuthorizationProfileVLAN) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["name_id"] = item.NameID
	respItem["tag_id"] = item.TagID

	return []map[string]interface{}{
		respItem,
	}

}

func flattenAuthorizationProfileGetAuthorizationProfileByIDItemIDReauth(item *isegosdk.ResponseAuthorizationProfileGetAuthorizationProfileByIDAuthorizationProfileReauth) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["timer"] = item.Timer
	respItem["connectivity"] = item.Connectivity

	return []map[string]interface{}{
		respItem,
	}

}

func flattenAuthorizationProfileGetAuthorizationProfileByIDItemIDWebRedirection(item *isegosdk.ResponseAuthorizationProfileGetAuthorizationProfileByIDAuthorizationProfileWebRedirection) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["web_redirection_type"] = item.WebRedirectionType
	respItem["acl"] = item.ACL
	respItem["portal_name"] = item.PortalName
	respItem["static_iphost_name_fqd_n"] = item.StaticIPHostNameFQDN
	respItem["display_certificates_renewal_messages"] = item.DisplayCertificatesRenewalMessages

	return []map[string]interface{}{
		respItem,
	}

}

func flattenAuthorizationProfileGetAuthorizationProfileByIDItemIDLink(item *isegosdk.ResponseAuthorizationProfileGetAuthorizationProfileByIDAuthorizationProfileLink) []map[string]interface{} {
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
