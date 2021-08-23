package ciscoise

import (
	"context"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSponsoredGuestPortal() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSponsoredGuestPortalRead,
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
												"authentication_method": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"assigned_guest_type_for_employee": &schema.Schema{
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

												"require_access_code": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
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
												"access_code": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"allow_guest_to_create_accounts": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"allow_forgot_password": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"allow_guest_to_change_password": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"allow_alternate_guest_portal": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"social_configs": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"social_media_type": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"social_media_value": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
														},
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

												"include_aup": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"require_aup_scrolling": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"use_diff_aup_for_employees": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"skip_aup_for_employees": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"display_frequency_interval_days": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"require_scrolling": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"display_frequency": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"guest_change_password_settings": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"allow_change_passwd_at_first_login": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
											},
										},
									},
									"guest_device_registration_settings": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"auto_register_guest_devices": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"allow_guests_to_register_devices": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
											},
										},
									},
									"byod_settings": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"byod_welcome_settings": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"enable_byo_d": &schema.Schema{
																Type:     schema.TypeBool,
																Computed: true,
															},
															"enable_guest_access": &schema.Schema{
																Type:     schema.TypeBool,
																Computed: true,
															},
															"require_mdm": &schema.Schema{
																Type:     schema.TypeBool,
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
														},
													},
												},
												"byod_registration_settings": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"show_device_id": &schema.Schema{
																Type:     schema.TypeBool,
																Computed: true,
															},
															"end_point_identity_group_id": &schema.Schema{
																Type:     schema.TypeString,
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

															"success_redirect": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"redirect_url": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
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
									"auth_success_settings": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"success_redirect": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"redirect_url": &schema.Schema{
													Type:     schema.TypeString,
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

func dataSourceSponsoredGuestPortalRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		log.Printf("[DEBUG] Selected method 1: GetSponsoredGuestPortals")
		queryParams1 := isegosdk.GetSponsoredGuestPortalsQueryParams{}

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

		response1, _, err := client.SponsoredGuestPortal.GetSponsoredGuestPortals(&queryParams1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetSponsoredGuestPortals", err,
				"Failure at GetSponsoredGuestPortals, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		var items1 []isegosdk.ResponseSponsoredGuestPortalGetSponsoredGuestPortalsSearchResultResources
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
				response1, _, err = client.SponsoredGuestPortal.GetSponsoredGuestPortals(&queryParams1)
				if err != nil {
					break
				}
				// All is good, continue to the next page
				continue
			}
			// Does not have next page finish iteration
			break
		}
		vItems1 := flattenSponsoredGuestPortalGetSponsoredGuestPortalsItems(&items1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSponsoredGuestPortals response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: GetSponsoredGuestPortalByID")
		vvID := vID.(string)

		response2, _, err := client.SponsoredGuestPortal.GetSponsoredGuestPortalByID(vvID)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetSponsoredGuestPortalByID", err,
				"Failure at GetSponsoredGuestPortalByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

		vItem2 := flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItem(&response2.SponsoredGuestPortal)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSponsoredGuestPortalByID response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSponsoredGuestPortalGetSponsoredGuestPortalsItems(items *[]isegosdk.ResponseSponsoredGuestPortalGetSponsoredGuestPortalsSearchResultResources) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItem["description"] = item.Description
		respItem["link"] = flattenSponsoredGuestPortalGetSponsoredGuestPortalsItemsLink(item.Link)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSponsoredGuestPortalGetSponsoredGuestPortalsItemsLink(item isegosdk.ResponseSponsoredGuestPortalGetSponsoredGuestPortalsSearchResultResourcesLink) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["rel"] = item.Rel
	respItem["href"] = item.Href
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItem(item *isegosdk.ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortal) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["description"] = item.Description
	respItem["portal_type"] = item.PortalType
	respItem["portal_test_url"] = item.PortalTestURL
	respItem["settings"] = flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemSettings(item.Settings)
	respItem["customizations"] = flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemCustomizations(item.Customizations)
	respItem["link"] = flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemLink(item.Link)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemSettings(item isegosdk.ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalSettings) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["portal_settings"] = flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemSettingsPortalSettings(item.PortalSettings)
	respItem["login_page_settings"] = flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemSettingsLoginPageSettings(item.LoginPageSettings)
	respItem["aup_settings"] = flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemSettingsAupSettings(item.AupSettings)
	respItem["guest_change_password_settings"] = flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemSettingsGuestChangePasswordSettings(item.GuestChangePasswordSettings)
	respItem["guest_device_registration_settings"] = flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemSettingsGuestDeviceRegistrationSettings(item.GuestDeviceRegistrationSettings)
	respItem["byod_settings"] = flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemSettingsByodSettings(item.ByodSettings)
	respItem["post_access_banner_settings"] = flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemSettingsPostAccessBannerSettings(item.PostAccessBannerSettings)
	respItem["auth_success_settings"] = flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemSettingsAuthSuccessSettings(item.AuthSuccessSettings)
	respItem["post_login_banner_settings"] = flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemSettingsPostLoginBannerSettings(item.PostLoginBannerSettings)
	respItem["support_info_settings"] = flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemSettingsSupportInfoSettings(item.SupportInfoSettings)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemSettingsPortalSettings(item isegosdk.ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalSettingsPortalSettings) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["https_port"] = item.HTTPSPort
	respItem["allowed_interfaces"] = item.AllowedInterfaces
	respItem["certificate_group_tag"] = item.CertificateGroupTag
	respItem["authentication_method"] = item.AuthenticationMethod
	respItem["assigned_guest_type_for_employee"] = item.AssignedGuestTypeForEmployee
	respItem["display_lang"] = item.DisplayLang
	respItem["fallback_language"] = item.FallbackLanguage
	respItem["always_used_language"] = item.AlwaysUsedLanguage

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemSettingsLoginPageSettings(item isegosdk.ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalSettingsLoginPageSettings) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["require_access_code"] = item.RequireAccessCode
	respItem["max_failed_attempts_before_rate_limit"] = item.MaxFailedAttemptsBeforeRateLimit
	respItem["time_between_logins_during_rate_limit"] = item.TimeBetweenLoginsDuringRateLimit
	respItem["include_aup"] = item.IncludeAup
	respItem["aup_display"] = item.AupDisplay
	respItem["require_aup_acceptance"] = item.RequireAupAcceptance
	respItem["access_code"] = item.AccessCode
	respItem["allow_guest_to_create_accounts"] = item.AllowGuestToCreateAccounts
	respItem["allow_forgot_password"] = item.AllowForgotPassword
	respItem["allow_guest_to_change_password"] = item.AllowGuestToChangePassword
	respItem["allow_alternate_guest_portal"] = item.AllowAlternateGuestPortal
	respItem["social_configs"] = flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemSettingsLoginPageSettingsSocialConfigs(item.SocialConfigs)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemSettingsLoginPageSettingsSocialConfigs(items []isegosdk.ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalSettingsLoginPageSettingsSocialConfigs) []map[string]interface{} {
	var respItems []map[string]interface{}
	for _, item := range items {
		respItem := make(map[string]interface{})
		respItem["social_media_type"] = item.SocialMediaType
		respItem["social_media_value"] = item.SocialMediaValue
	}
	return respItems

}

func flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemSettingsAupSettings(item isegosdk.ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalSettingsAupSettings) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["include_aup"] = item.IncludeAup
	respItem["require_aup_scrolling"] = item.RequireAupScrolling
	respItem["use_diff_aup_for_employees"] = item.UseDiffAupForEmployees
	respItem["skip_aup_for_employees"] = item.SkipAupForEmployees
	respItem["display_frequency_interval_days"] = item.DisplayFrequencyIntervalDays
	respItem["require_scrolling"] = item.RequireScrolling
	respItem["display_frequency"] = item.DisplayFrequency

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemSettingsGuestChangePasswordSettings(item isegosdk.ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalSettingsGuestChangePasswordSettings) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["allow_change_passwd_at_first_login"] = item.AllowChangePasswdAtFirstLogin

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemSettingsGuestDeviceRegistrationSettings(item isegosdk.ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalSettingsGuestDeviceRegistrationSettings) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["auto_register_guest_devices"] = item.AutoRegisterGuestDevices
	respItem["allow_guests_to_register_devices"] = item.AllowGuestsToRegisterDevices

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemSettingsByodSettings(item isegosdk.ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalSettingsByodSettings) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["byod_welcome_settings"] = flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemSettingsByodSettingsByodWelcomeSettings(item.ByodWelcomeSettings)
	respItem["byod_registration_settings"] = flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemSettingsByodSettingsByodRegistrationSettings(item.ByodRegistrationSettings)
	respItem["byod_registration_success_settings"] = flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemSettingsByodSettingsByodRegistrationSuccessSettings(item.ByodRegistrationSuccessSettings)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemSettingsByodSettingsByodWelcomeSettings(item isegosdk.ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalSettingsByodSettingsByodWelcomeSettings) []map[string]interface{} {
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

func flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemSettingsByodSettingsByodRegistrationSettings(item isegosdk.ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalSettingsByodSettingsByodRegistrationSettings) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["show_device_id"] = item.ShowDeviceID
	respItem["end_point_identity_group_id"] = item.EndPointIDentityGroupID

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemSettingsByodSettingsByodRegistrationSuccessSettings(item isegosdk.ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalSettingsByodSettingsByodRegistrationSuccessSettings) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["success_redirect"] = item.SuccessRedirect
	respItem["redirect_url"] = item.RedirectURL

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemSettingsPostAccessBannerSettings(item isegosdk.ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalSettingsPostAccessBannerSettings) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["include_post_access_banner"] = item.IncludePostAccessBanner

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemSettingsAuthSuccessSettings(item isegosdk.ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalSettingsAuthSuccessSettings) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["success_redirect"] = item.SuccessRedirect
	respItem["redirect_url"] = item.RedirectURL

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemSettingsPostLoginBannerSettings(item isegosdk.ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalSettingsPostLoginBannerSettings) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["include_post_access_banner"] = item.IncludePostAccessBanner

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemSettingsSupportInfoSettings(item isegosdk.ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalSettingsSupportInfoSettings) []map[string]interface{} {
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

func flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemCustomizations(item isegosdk.ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalCustomizations) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["portal_theme"] = flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemCustomizationsPortalTheme(item.PortalTheme)
	respItem["portal_tweak_settings"] = flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemCustomizationsPortalTweakSettings(item.PortalTweakSettings)
	respItem["language"] = flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemCustomizationsLanguage(item.Language)
	respItem["global_customizations"] = flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemCustomizationsGlobalCustomizations(item.GlobalCustomizations)
	respItem["page_customizations"] = flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemCustomizationsPageCustomizations(item.PageCustomizations)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemCustomizationsPortalTheme(item isegosdk.ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsPortalTheme) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["theme_data"] = item.ThemeData

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemCustomizationsPortalTweakSettings(item isegosdk.ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsPortalTweakSettings) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["banner_color"] = item.BannerColor
	respItem["banner_text_color"] = item.BannerTextColor
	respItem["page_background_color"] = item.PageBackgroundColor
	respItem["page_label_and_text_color"] = item.PageLabelAndTextColor

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemCustomizationsLanguage(item isegosdk.ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsLanguage) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["view_language"] = item.ViewLanguage

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemCustomizationsGlobalCustomizations(item isegosdk.ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsGlobalCustomizations) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["mobile_logo_image"] = flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemCustomizationsGlobalCustomizationsMobileLogoImage(item.MobileLogoImage)
	respItem["desktop_logo_image"] = flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemCustomizationsGlobalCustomizationsDesktopLogoImage(item.DesktopLogoImage)
	respItem["banner_image"] = flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemCustomizationsGlobalCustomizationsBannerImage(item.BannerImage)
	respItem["background_image"] = flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemCustomizationsGlobalCustomizationsBackgroundImage(item.BackgroundImage)
	respItem["banner_title"] = item.BannerTitle
	respItem["contact_text"] = item.ContactText
	respItem["footer_element"] = item.FooterElement

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemCustomizationsGlobalCustomizationsMobileLogoImage(item isegosdk.ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsGlobalCustomizationsMobileLogoImage) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["data"] = item.Data

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemCustomizationsGlobalCustomizationsDesktopLogoImage(item isegosdk.ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsGlobalCustomizationsDesktopLogoImage) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["data"] = item.Data

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemCustomizationsGlobalCustomizationsBannerImage(item isegosdk.ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsGlobalCustomizationsBannerImage) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["data"] = item.Data

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemCustomizationsGlobalCustomizationsBackgroundImage(item isegosdk.ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsGlobalCustomizationsBackgroundImage) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["data"] = item.Data

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemCustomizationsPageCustomizations(item isegosdk.ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsPageCustomizations) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["data"] = flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemCustomizationsPageCustomizationsData(item.Data)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemCustomizationsPageCustomizationsData(items []isegosdk.ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsPageCustomizationsData) []map[string]interface{} {
	var respItems []map[string]interface{}
	for _, item := range items {
		respItem := make(map[string]interface{})
		respItem["key"] = item.Key
		respItem["value"] = item.Value
	}
	return respItems

}

func flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemLink(item isegosdk.ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalLink) []map[string]interface{} {
	respItem := make(map[string]interface{})
	respItem["rel"] = item.Rel
	respItem["href"] = item.Href
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}
