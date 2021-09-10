package ciscoise

import (
	"context"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceByodPortal() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceByodPortalRead,
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
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"customizations": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"global_customizations": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"background_image": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"data": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"banner_image": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"data": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"banner_title": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"contact_text": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"desktop_logo_image": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"data": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"footer_element": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"mobile_logo_image": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"data": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
											},
										},
									},
									"language": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"view_language": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"page_customizations": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"data": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"key": &schema.Schema{
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
									"portal_theme": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"id": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"theme_data": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"portal_tweak_settings": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"banner_color": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"banner_text_color": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"page_background_color": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"page_label_and_text_color": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
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
						"portal_test_url": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"portal_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"settings": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"byod_settings": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"byod_registration_settings": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"end_point_identity_group_id": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"show_device_id": &schema.Schema{
																Type:     schema.TypeBool,
																Computed: true,
															},
														},
													},
												},
												"byod_registration_success_settings": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"redirect_url": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"success_redirect": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"byod_welcome_settings": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"aup_display": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"enable_byo_d": &schema.Schema{
																Type:     schema.TypeBool,
																Computed: true,
															},
															"enable_guest_access": &schema.Schema{
																Type:     schema.TypeBool,
																Computed: true,
															},
															"include_aup": &schema.Schema{
																Type:     schema.TypeBool,
																Computed: true,
															},
															"require_aup_acceptance": &schema.Schema{
																Type:     schema.TypeBool,
																Computed: true,
															},
															"require_mdm": &schema.Schema{
																Type:     schema.TypeBool,
																Computed: true,
															},
															"require_scrolling": &schema.Schema{
																Type:     schema.TypeBool,
																Computed: true,
															},
														},
													},
												},
											},
										},
									},
									"portal_settings": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"allowed_interfaces": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"always_used_language": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"certificate_group_tag": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"display_lang": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"endpoint_identity_group": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"fallback_language": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"https_port": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},
									"support_info_settings": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"default_empty_field_value": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"empty_field_display": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"include_browser_user_agent": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"include_failure_code": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"include_ip_address": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"include_mac_addr": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"include_policy_server": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"include_support_info_page": &schema.Schema{
													Type:     schema.TypeBool,
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

func dataSourceByodPortalRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vPage, okPage := d.GetOk("page")
	vSize, okSize := d.GetOk("size")
	vSortasc, okSortasc := d.GetOk("sortasc")
	vSortdsc, okSortdsc := d.GetOk("sortdsc")
	vFilter, okFilter := d.GetOk("filter")
	vFilterType, okFilterType := d.GetOk("filter_type")
	vID, okID := d.GetOk("id")

	method1 := []bool{okPage, okSize, okSortasc, okSortdsc, okFilter, okFilterType}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetByodPortal")
		queryParams1 := isegosdk.GetByodPortalQueryParams{}

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

		response1, _, err := client.ByodPortal.GetByodPortal(&queryParams1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetByodPortal", err,
				"Failure at GetByodPortal, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		var items1 []isegosdk.ResponseByodPortalGetByodPortalSearchResultResources
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
				response1, _, err = client.ByodPortal.GetByodPortal(&queryParams1)
				if err != nil {
					break
				}
				// All is good, continue to the next page
				continue
			}
			// Does not have next page finish iteration
			break
		}
		vItems1 := flattenByodPortalGetByodPortalItems(&items1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetByodPortal response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: GetByodPortalByID")
		vvID := vID.(string)

		response2, _, err := client.ByodPortal.GetByodPortalByID(vvID)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetByodPortalByID", err,
				"Failure at GetByodPortalByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

		vItem2 := flattenByodPortalGetByodPortalByIDItem(response2.ByodPortal)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetByodPortalByID response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenByodPortalGetByodPortalItems(items *[]isegosdk.ResponseByodPortalGetByodPortalSearchResultResources) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItem["description"] = item.Description
		respItem["link"] = flattenByodPortalGetByodPortalItemsLink(item.Link)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenByodPortalGetByodPortalItemsLink(item *isegosdk.ResponseByodPortalGetByodPortalSearchResultResourcesLink) []map[string]interface{} {
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

func flattenByodPortalGetByodPortalByIDItem(item *isegosdk.ResponseByodPortalGetByodPortalByIDByodPortal) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["description"] = item.Description
	respItem["portal_type"] = item.PortalType
	respItem["portal_test_url"] = item.PortalTestURL
	respItem["settings"] = flattenByodPortalGetByodPortalByIDItemSettings(item.Settings)
	respItem["customizations"] = flattenByodPortalGetByodPortalByIDItemCustomizations(item.Customizations)
	respItem["link"] = flattenByodPortalGetByodPortalByIDItemLink(item.Link)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenByodPortalGetByodPortalByIDItemSettings(item *isegosdk.ResponseByodPortalGetByodPortalByIDByodPortalSettings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["portal_settings"] = flattenByodPortalGetByodPortalByIDItemSettingsPortalSettings(item.PortalSettings)
	respItem["byod_settings"] = flattenByodPortalGetByodPortalByIDItemSettingsByodSettings(item.ByodSettings)
	respItem["support_info_settings"] = flattenByodPortalGetByodPortalByIDItemSettingsSupportInfoSettings(item.SupportInfoSettings)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenByodPortalGetByodPortalByIDItemSettingsPortalSettings(item *isegosdk.ResponseByodPortalGetByodPortalByIDByodPortalSettingsPortalSettings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["https_port"] = item.HTTPSPort
	respItem["allowed_interfaces"] = item.AllowedInterfaces
	respItem["certificate_group_tag"] = item.CertificateGroupTag
	respItem["endpoint_identity_group"] = item.EndpointIDentityGroup
	respItem["display_lang"] = item.DisplayLang
	respItem["fallback_language"] = item.FallbackLanguage
	respItem["always_used_language"] = item.AlwaysUsedLanguage

	return []map[string]interface{}{
		respItem,
	}

}

func flattenByodPortalGetByodPortalByIDItemSettingsByodSettings(item *isegosdk.ResponseByodPortalGetByodPortalByIDByodPortalSettingsByodSettings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["byod_welcome_settings"] = flattenByodPortalGetByodPortalByIDItemSettingsByodSettingsByodWelcomeSettings(item.ByodWelcomeSettings)
	respItem["byod_registration_settings"] = flattenByodPortalGetByodPortalByIDItemSettingsByodSettingsByodRegistrationSettings(item.ByodRegistrationSettings)
	respItem["byod_registration_success_settings"] = flattenByodPortalGetByodPortalByIDItemSettingsByodSettingsByodRegistrationSuccessSettings(item.ByodRegistrationSuccessSettings)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenByodPortalGetByodPortalByIDItemSettingsByodSettingsByodWelcomeSettings(item *isegosdk.ResponseByodPortalGetByodPortalByIDByodPortalSettingsByodSettingsByodWelcomeSettings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["enable_byo_d"] = item.EnableByod
	respItem["enable_guest_access"] = item.EnableGuestAccess
	respItem["require_mdm"] = item.RequireMdm
	respItem["include_aup"] = item.IncludeAup
	respItem["aup_display"] = item.AupDisplay
	respItem["require_aup_acceptance"] = item.RequireAupAcceptance
	respItem["require_scrolling"] = item.RequireScrolling

	return []map[string]interface{}{
		respItem,
	}

}

func flattenByodPortalGetByodPortalByIDItemSettingsByodSettingsByodRegistrationSettings(item *isegosdk.ResponseByodPortalGetByodPortalByIDByodPortalSettingsByodSettingsByodRegistrationSettings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["show_device_id"] = item.ShowDeviceID
	respItem["end_point_identity_group_id"] = item.EndPointIDentityGroupID

	return []map[string]interface{}{
		respItem,
	}

}

func flattenByodPortalGetByodPortalByIDItemSettingsByodSettingsByodRegistrationSuccessSettings(item *isegosdk.ResponseByodPortalGetByodPortalByIDByodPortalSettingsByodSettingsByodRegistrationSuccessSettings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["success_redirect"] = item.SuccessRedirect
	respItem["redirect_url"] = item.RedirectURL

	return []map[string]interface{}{
		respItem,
	}

}

func flattenByodPortalGetByodPortalByIDItemSettingsSupportInfoSettings(item *isegosdk.ResponseByodPortalGetByodPortalByIDByodPortalSettingsSupportInfoSettings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["include_support_info_page"] = item.IncludeSupportInfoPage
	respItem["include_mac_addr"] = item.IncludeMacAddr
	respItem["include_ip_address"] = item.IncludeIPAddress
	respItem["include_browser_user_agent"] = item.IncludeBrowserUserAgent
	respItem["include_policy_server"] = item.IncludePolicyServer
	respItem["include_failure_code"] = item.IncludeFailureCode
	respItem["empty_field_display"] = item.EmptyFieldDisplay
	respItem["default_empty_field_value"] = item.DefaultEmptyFieldValue

	return []map[string]interface{}{
		respItem,
	}

}

func flattenByodPortalGetByodPortalByIDItemCustomizations(item *isegosdk.ResponseByodPortalGetByodPortalByIDByodPortalCustomizations) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["portal_theme"] = flattenByodPortalGetByodPortalByIDItemCustomizationsPortalTheme(item.PortalTheme)
	respItem["portal_tweak_settings"] = flattenByodPortalGetByodPortalByIDItemCustomizationsPortalTweakSettings(item.PortalTweakSettings)
	respItem["language"] = flattenByodPortalGetByodPortalByIDItemCustomizationsLanguage(item.Language)
	respItem["global_customizations"] = flattenByodPortalGetByodPortalByIDItemCustomizationsGlobalCustomizations(item.GlobalCustomizations)
	respItem["page_customizations"] = flattenByodPortalGetByodPortalByIDItemCustomizationsPageCustomizations(item.PageCustomizations)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenByodPortalGetByodPortalByIDItemCustomizationsPortalTheme(item *isegosdk.ResponseByodPortalGetByodPortalByIDByodPortalCustomizationsPortalTheme) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["theme_data"] = item.ThemeData

	return []map[string]interface{}{
		respItem,
	}

}

func flattenByodPortalGetByodPortalByIDItemCustomizationsPortalTweakSettings(item *isegosdk.ResponseByodPortalGetByodPortalByIDByodPortalCustomizationsPortalTweakSettings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["banner_color"] = item.BannerColor
	respItem["banner_text_color"] = item.BannerTextColor
	respItem["page_background_color"] = item.PageBackgroundColor
	respItem["page_label_and_text_color"] = item.PageLabelAndTextColor

	return []map[string]interface{}{
		respItem,
	}

}

func flattenByodPortalGetByodPortalByIDItemCustomizationsLanguage(item *isegosdk.ResponseByodPortalGetByodPortalByIDByodPortalCustomizationsLanguage) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["view_language"] = item.ViewLanguage

	return []map[string]interface{}{
		respItem,
	}

}

func flattenByodPortalGetByodPortalByIDItemCustomizationsGlobalCustomizations(item *isegosdk.ResponseByodPortalGetByodPortalByIDByodPortalCustomizationsGlobalCustomizations) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["mobile_logo_image"] = flattenByodPortalGetByodPortalByIDItemCustomizationsGlobalCustomizationsMobileLogoImage(item.MobileLogoImage)
	respItem["desktop_logo_image"] = flattenByodPortalGetByodPortalByIDItemCustomizationsGlobalCustomizationsDesktopLogoImage(item.DesktopLogoImage)
	respItem["banner_image"] = flattenByodPortalGetByodPortalByIDItemCustomizationsGlobalCustomizationsBannerImage(item.BannerImage)
	respItem["background_image"] = flattenByodPortalGetByodPortalByIDItemCustomizationsGlobalCustomizationsBackgroundImage(item.BackgroundImage)
	respItem["banner_title"] = item.BannerTitle
	respItem["contact_text"] = item.ContactText
	respItem["footer_element"] = item.FooterElement

	return []map[string]interface{}{
		respItem,
	}

}

func flattenByodPortalGetByodPortalByIDItemCustomizationsGlobalCustomizationsMobileLogoImage(item *isegosdk.ResponseByodPortalGetByodPortalByIDByodPortalCustomizationsGlobalCustomizationsMobileLogoImage) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["data"] = item.Data

	return []map[string]interface{}{
		respItem,
	}

}

func flattenByodPortalGetByodPortalByIDItemCustomizationsGlobalCustomizationsDesktopLogoImage(item *isegosdk.ResponseByodPortalGetByodPortalByIDByodPortalCustomizationsGlobalCustomizationsDesktopLogoImage) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["data"] = item.Data

	return []map[string]interface{}{
		respItem,
	}

}

func flattenByodPortalGetByodPortalByIDItemCustomizationsGlobalCustomizationsBannerImage(item *isegosdk.ResponseByodPortalGetByodPortalByIDByodPortalCustomizationsGlobalCustomizationsBannerImage) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["data"] = item.Data

	return []map[string]interface{}{
		respItem,
	}

}

func flattenByodPortalGetByodPortalByIDItemCustomizationsGlobalCustomizationsBackgroundImage(item *isegosdk.ResponseByodPortalGetByodPortalByIDByodPortalCustomizationsGlobalCustomizationsBackgroundImage) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["data"] = item.Data

	return []map[string]interface{}{
		respItem,
	}

}

func flattenByodPortalGetByodPortalByIDItemCustomizationsPageCustomizations(item *isegosdk.ResponseByodPortalGetByodPortalByIDByodPortalCustomizationsPageCustomizations) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["data"] = flattenByodPortalGetByodPortalByIDItemCustomizationsPageCustomizationsData(item.Data)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenByodPortalGetByodPortalByIDItemCustomizationsPageCustomizationsData(items *[]isegosdk.ResponseByodPortalGetByodPortalByIDByodPortalCustomizationsPageCustomizationsData) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["key"] = item.Key
		respItem["value"] = item.Value
	}
	return respItems

}

func flattenByodPortalGetByodPortalByIDItemLink(item *isegosdk.ResponseByodPortalGetByodPortalByIDByodPortalLink) []map[string]interface{} {
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
