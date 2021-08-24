package ciscoise

import (
	"context"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceMyDevicePortal() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceMyDevicePortalRead,
		Schema: map[string]*schema.Schema{
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
			"items": &schema.Schema{
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
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"link": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"rel": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"href": &schema.Schema{
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
					},
				},
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
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"portal_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"portal_test_url": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"settings": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"portal_settings": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"https_port": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"allowed_interfaces": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"certificate_group_tag": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"endpoint_identity_group": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"display_lang": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"fallback_language": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"always_used_language": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"login_page_settings": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"max_failed_attempts_before_rate_limit": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"time_between_logins_during_rate_limit": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"include_aup": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"aup_display": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"require_aup_acceptance": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"require_scrolling": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"social_configs": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
									"aup_settings": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"display_frequency_interval_days": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"display_frequency": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"include_aup": &schema.Schema{
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
									"employee_change_password_settings": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"allow_employee_to_change_pwd": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
											},
										},
									},
									"post_login_banner_settings": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"include_post_access_banner": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
											},
										},
									},
									"post_access_banner_settings": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"include_post_access_banner": &schema.Schema{
													Type:     schema.TypeBool,
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

												"include_support_info_page": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"include_mac_addr": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"include_ip_address": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"include_browser_user_agent": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"include_policy_server": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"include_failure_code": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"empty_field_display": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"default_empty_field_value": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
						"customizations": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

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
									"global_customizations": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

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
												"banner_title": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"contact_text": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"footer_element": &schema.Schema{
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
								},
							},
						},
						"link": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"rel": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"href": &schema.Schema{
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
					},
				},
			},
		},
	}
}

func dataSourceMyDevicePortalRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetMyDevicePortal")
		queryParams1 := isegosdk.GetMyDevicePortalQueryParams{}

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

		response1, _, err := client.MyDevicePortal.GetMyDevicePortal(&queryParams1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetMyDevicePortal", err,
				"Failure at GetMyDevicePortal, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		var items1 []isegosdk.ResponseMyDevicePortalGetMyDevicePortalSearchResultResources
		for len(response1.SearchResult.Resources) > 0 {
			items1 = append(items1, response1.SearchResult.Resources...)
			if response1.SearchResult.NextPage.Rel == "next" {
				href := response1.SearchResult.NextPage.Href
				page, size, err := getNextPageAndSizeParams(href)
				if err != nil {
					break
				}
				queryParams1.Page = page
				queryParams1.Size = size
				response1, _, err = client.MyDevicePortal.GetMyDevicePortal(&queryParams1)
				if err != nil {
					break
				}
				// All is good, continue to the next page
				continue
			}
			// Does not have next page finish iteration
			break
		}
		vItems1 := flattenMyDevicePortalGetMyDevicePortalItems(&items1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetMyDevicePortal response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: GetMyDevicePortalByID")
		vvID := vID.(string)

		response2, _, err := client.MyDevicePortal.GetMyDevicePortalByID(vvID)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetMyDevicePortalByID", err,
				"Failure at GetMyDevicePortalByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

		vItem2 := flattenMyDevicePortalGetMyDevicePortalByIDItem(&response2.MyDevicePortal)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetMyDevicePortalByID response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenMyDevicePortalGetMyDevicePortalItems(items *[]isegosdk.ResponseMyDevicePortalGetMyDevicePortalSearchResultResources) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItem["description"] = item.Description
		respItem["link"] = flattenMyDevicePortalGetMyDevicePortalItemsLink(item.Link)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenMyDevicePortalGetMyDevicePortalItemsLink(item isegosdk.ResponseMyDevicePortalGetMyDevicePortalSearchResultResourcesLink) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["rel"] = item.Rel
	respItem["href"] = item.Href
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}

func flattenMyDevicePortalGetMyDevicePortalByIDItem(item *isegosdk.ResponseMyDevicePortalGetMyDevicePortalByIDMyDevicePortal) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["description"] = item.Description
	respItem["portal_type"] = item.PortalType
	respItem["portal_test_url"] = item.PortalTestURL
	respItem["settings"] = flattenMyDevicePortalGetMyDevicePortalByIDItemSettings(item.Settings)
	respItem["customizations"] = flattenMyDevicePortalGetMyDevicePortalByIDItemCustomizations(item.Customizations)
	respItem["link"] = flattenMyDevicePortalGetMyDevicePortalByIDItemLink(item.Link)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenMyDevicePortalGetMyDevicePortalByIDItemSettings(item isegosdk.ResponseMyDevicePortalGetMyDevicePortalByIDMyDevicePortalSettings) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["portal_settings"] = flattenMyDevicePortalGetMyDevicePortalByIDItemSettingsPortalSettings(item.PortalSettings)
	respItem["login_page_settings"] = flattenMyDevicePortalGetMyDevicePortalByIDItemSettingsLoginPageSettings(item.LoginPageSettings)
	respItem["aup_settings"] = flattenMyDevicePortalGetMyDevicePortalByIDItemSettingsAupSettings(item.AupSettings)
	respItem["employee_change_password_settings"] = flattenMyDevicePortalGetMyDevicePortalByIDItemSettingsEmployeeChangePasswordSettings(item.EmployeeChangePasswordSettings)
	respItem["post_login_banner_settings"] = flattenMyDevicePortalGetMyDevicePortalByIDItemSettingsPostLoginBannerSettings(item.PostLoginBannerSettings)
	respItem["post_access_banner_settings"] = flattenMyDevicePortalGetMyDevicePortalByIDItemSettingsPostAccessBannerSettings(item.PostAccessBannerSettings)
	respItem["support_info_settings"] = flattenMyDevicePortalGetMyDevicePortalByIDItemSettingsSupportInfoSettings(item.SupportInfoSettings)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenMyDevicePortalGetMyDevicePortalByIDItemSettingsPortalSettings(item isegosdk.ResponseMyDevicePortalGetMyDevicePortalByIDMyDevicePortalSettingsPortalSettings) []map[string]interface{} {
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

func flattenMyDevicePortalGetMyDevicePortalByIDItemSettingsLoginPageSettings(item isegosdk.ResponseMyDevicePortalGetMyDevicePortalByIDMyDevicePortalSettingsLoginPageSettings) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["max_failed_attempts_before_rate_limit"] = item.MaxFailedAttemptsBeforeRateLimit
	respItem["time_between_logins_during_rate_limit"] = item.TimeBetweenLoginsDuringRateLimit
	respItem["include_aup"] = item.IncludeAup
	respItem["aup_display"] = item.AupDisplay
	respItem["require_aup_acceptance"] = item.RequireAupAcceptance
	respItem["require_scrolling"] = item.RequireScrolling
	respItem["social_configs"] = responseInterfaceToSliceString(item.SocialConfigs)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenMyDevicePortalGetMyDevicePortalByIDItemSettingsAupSettings(item isegosdk.ResponseMyDevicePortalGetMyDevicePortalByIDMyDevicePortalSettingsAupSettings) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["display_frequency_interval_days"] = item.DisplayFrequencyIntervalDays
	respItem["display_frequency"] = item.DisplayFrequency
	respItem["include_aup"] = item.IncludeAup
	respItem["require_scrolling"] = item.RequireScrolling

	return []map[string]interface{}{
		respItem,
	}

}

func flattenMyDevicePortalGetMyDevicePortalByIDItemSettingsEmployeeChangePasswordSettings(item isegosdk.ResponseMyDevicePortalGetMyDevicePortalByIDMyDevicePortalSettingsEmployeeChangePasswordSettings) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["allow_employee_to_change_pwd"] = item.AllowEmployeeToChangePwd

	return []map[string]interface{}{
		respItem,
	}

}

func flattenMyDevicePortalGetMyDevicePortalByIDItemSettingsPostLoginBannerSettings(item isegosdk.ResponseMyDevicePortalGetMyDevicePortalByIDMyDevicePortalSettingsPostLoginBannerSettings) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["include_post_access_banner"] = item.IncludePostAccessBanner

	return []map[string]interface{}{
		respItem,
	}

}

func flattenMyDevicePortalGetMyDevicePortalByIDItemSettingsPostAccessBannerSettings(item isegosdk.ResponseMyDevicePortalGetMyDevicePortalByIDMyDevicePortalSettingsPostAccessBannerSettings) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["include_post_access_banner"] = item.IncludePostAccessBanner

	return []map[string]interface{}{
		respItem,
	}

}

func flattenMyDevicePortalGetMyDevicePortalByIDItemSettingsSupportInfoSettings(item isegosdk.ResponseMyDevicePortalGetMyDevicePortalByIDMyDevicePortalSettingsSupportInfoSettings) []map[string]interface{} {
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

func flattenMyDevicePortalGetMyDevicePortalByIDItemCustomizations(item isegosdk.ResponseMyDevicePortalGetMyDevicePortalByIDMyDevicePortalCustomizations) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["portal_theme"] = flattenMyDevicePortalGetMyDevicePortalByIDItemCustomizationsPortalTheme(item.PortalTheme)
	respItem["portal_tweak_settings"] = flattenMyDevicePortalGetMyDevicePortalByIDItemCustomizationsPortalTweakSettings(item.PortalTweakSettings)
	respItem["language"] = flattenMyDevicePortalGetMyDevicePortalByIDItemCustomizationsLanguage(item.Language)
	respItem["global_customizations"] = flattenMyDevicePortalGetMyDevicePortalByIDItemCustomizationsGlobalCustomizations(item.GlobalCustomizations)
	respItem["page_customizations"] = flattenMyDevicePortalGetMyDevicePortalByIDItemCustomizationsPageCustomizations(item.PageCustomizations)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenMyDevicePortalGetMyDevicePortalByIDItemCustomizationsPortalTheme(item isegosdk.ResponseMyDevicePortalGetMyDevicePortalByIDMyDevicePortalCustomizationsPortalTheme) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["theme_data"] = item.ThemeData

	return []map[string]interface{}{
		respItem,
	}

}

func flattenMyDevicePortalGetMyDevicePortalByIDItemCustomizationsPortalTweakSettings(item isegosdk.ResponseMyDevicePortalGetMyDevicePortalByIDMyDevicePortalCustomizationsPortalTweakSettings) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["banner_color"] = item.BannerColor
	respItem["banner_text_color"] = item.BannerTextColor
	respItem["page_background_color"] = item.PageBackgroundColor
	respItem["page_label_and_text_color"] = item.PageLabelAndTextColor

	return []map[string]interface{}{
		respItem,
	}

}

func flattenMyDevicePortalGetMyDevicePortalByIDItemCustomizationsLanguage(item isegosdk.ResponseMyDevicePortalGetMyDevicePortalByIDMyDevicePortalCustomizationsLanguage) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["view_language"] = item.ViewLanguage

	return []map[string]interface{}{
		respItem,
	}

}

func flattenMyDevicePortalGetMyDevicePortalByIDItemCustomizationsGlobalCustomizations(item isegosdk.ResponseMyDevicePortalGetMyDevicePortalByIDMyDevicePortalCustomizationsGlobalCustomizations) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["mobile_logo_image"] = flattenMyDevicePortalGetMyDevicePortalByIDItemCustomizationsGlobalCustomizationsMobileLogoImage(item.MobileLogoImage)
	respItem["desktop_logo_image"] = flattenMyDevicePortalGetMyDevicePortalByIDItemCustomizationsGlobalCustomizationsDesktopLogoImage(item.DesktopLogoImage)
	respItem["banner_image"] = flattenMyDevicePortalGetMyDevicePortalByIDItemCustomizationsGlobalCustomizationsBannerImage(item.BannerImage)
	respItem["background_image"] = flattenMyDevicePortalGetMyDevicePortalByIDItemCustomizationsGlobalCustomizationsBackgroundImage(item.BackgroundImage)
	respItem["banner_title"] = item.BannerTitle
	respItem["contact_text"] = item.ContactText
	respItem["footer_element"] = item.FooterElement

	return []map[string]interface{}{
		respItem,
	}

}

func flattenMyDevicePortalGetMyDevicePortalByIDItemCustomizationsGlobalCustomizationsMobileLogoImage(item isegosdk.ResponseMyDevicePortalGetMyDevicePortalByIDMyDevicePortalCustomizationsGlobalCustomizationsMobileLogoImage) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["data"] = item.Data

	return []map[string]interface{}{
		respItem,
	}

}

func flattenMyDevicePortalGetMyDevicePortalByIDItemCustomizationsGlobalCustomizationsDesktopLogoImage(item isegosdk.ResponseMyDevicePortalGetMyDevicePortalByIDMyDevicePortalCustomizationsGlobalCustomizationsDesktopLogoImage) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["data"] = item.Data

	return []map[string]interface{}{
		respItem,
	}

}

func flattenMyDevicePortalGetMyDevicePortalByIDItemCustomizationsGlobalCustomizationsBannerImage(item isegosdk.ResponseMyDevicePortalGetMyDevicePortalByIDMyDevicePortalCustomizationsGlobalCustomizationsBannerImage) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["data"] = item.Data

	return []map[string]interface{}{
		respItem,
	}

}

func flattenMyDevicePortalGetMyDevicePortalByIDItemCustomizationsGlobalCustomizationsBackgroundImage(item isegosdk.ResponseMyDevicePortalGetMyDevicePortalByIDMyDevicePortalCustomizationsGlobalCustomizationsBackgroundImage) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["data"] = item.Data

	return []map[string]interface{}{
		respItem,
	}

}

func flattenMyDevicePortalGetMyDevicePortalByIDItemCustomizationsPageCustomizations(item isegosdk.ResponseMyDevicePortalGetMyDevicePortalByIDMyDevicePortalCustomizationsPageCustomizations) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["data"] = flattenMyDevicePortalGetMyDevicePortalByIDItemCustomizationsPageCustomizationsData(item.Data)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenMyDevicePortalGetMyDevicePortalByIDItemCustomizationsPageCustomizationsData(items []isegosdk.ResponseMyDevicePortalGetMyDevicePortalByIDMyDevicePortalCustomizationsPageCustomizationsData) []map[string]interface{} {
	var respItems []map[string]interface{}
	for _, item := range items {
		respItem := make(map[string]interface{})
		respItem["key"] = item.Key
		respItem["value"] = item.Value
	}
	return respItems

}

func flattenMyDevicePortalGetMyDevicePortalByIDItemLink(item isegosdk.ResponseMyDevicePortalGetMyDevicePortalByIDMyDevicePortalLink) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["rel"] = item.Rel
	respItem["href"] = item.Href
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}
