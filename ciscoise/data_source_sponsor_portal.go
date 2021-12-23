package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSponsorPortal() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on SponsorPortal.

- This data source allows the client to get a sponsor portal by ID.

- This data source allows the client to get all the sponsor portals.

Filter:

[name, description]

Sorting:

[name, description]
`,

		ReadContext: dataSourceSponsorPortalRead,
		Schema: map[string]*schema.Schema{
			"filter": &schema.Schema{
				Description: `filter query parameter. 

**Simple filtering** should be available through the filter query string parameter. The structure of a filter is
a triplet of field operator and value separated with dots. More than one filter can be sent. The logical operator
common to ALL filter criteria will be by default AND, and can be changed by using the "filterType=or" query
string parameter. Each resource Data model description should specify if an attribute is a filtered field.



              Operator    | Description 

              ------------|----------------

              EQ          | Equals 

              NEQ         | Not Equals 

              GT          | Greater Than 

              LT          | Less Then 

              STARTSW     | Starts With 

              NSTARTSW    | Not Starts With 

              ENDSW       | Ends With 

              NENDSW      | Not Ends With 

              CONTAINS	  | Contains 

              NCONTAINS	  | Not Contains 

`,
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"filter_type": &schema.Schema{
				Description: `filterType query parameter. The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the parameter`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"id": &schema.Schema{
				Description: `id path parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"page": &schema.Schema{
				Description: `page query parameter. Page number`,
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"size": &schema.Schema{
				Description: `size query parameter. Number of objects returned per page`,
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"sortasc": &schema.Schema{
				Description: `sortasc query parameter. sort asc`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"sortdsc": &schema.Schema{
				Description: `sortdsc query parameter. sort desc`,
				Type:        schema.TypeString,
				Optional:    true,
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
																Description: `Represented as base 64 encoded string of the image byte array`,
																Type:        schema.TypeString,
																Computed:    true,
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
																Description: `Represented as base 64 encoded string of the image byte array`,
																Type:        schema.TypeString,
																Computed:    true,
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
																Description: `Represented as base 64 encoded string of the image byte array`,
																Type:        schema.TypeString,
																Computed:    true,
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
																Description: `Represented as base 64 encoded string of the image byte array`,
																Type:        schema.TypeString,
																Computed:    true,
															},
														},
													},
												},
											},
										},
									},
									"language": &schema.Schema{
										Description: `This property is supported only for Read operation and it allows to show the customizations in English. Other languages are not supported`,
										Type:        schema.TypeList,
										Computed:    true,
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
													Description: `The Dictionary will be exposed here as key value pair`,
													Type:        schema.TypeList,
													Computed:    true,
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
													Description: `The system- or user-assigned name of the portal theme`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"theme_data": &schema.Schema{
													Description: `A CSS file, represented as a Base64-encoded byte array`,
													Type:        schema.TypeString,
													Computed:    true,
												},
											},
										},
									},
									"portal_tweak_settings": &schema.Schema{
										Description: `The Tweak Settings are a customization of the Portal Theme that has been selected for the portal.
When the Portal Theme selection is changed, the Tweak Settings are overwritten to match the values in the theme.
The Tweak Settings can subsequently be changed by the user`,
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"banner_color": &schema.Schema{
													Description: `Hex value of color`,
													Type:        schema.TypeString,
													Computed:    true,
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
							Description: `URL to bring up a test page for this portal`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"portal_type": &schema.Schema{
							Description: `Allowed values:
- BYOD,
- HOTSPOTGUEST,
- MYDEVICE,
- SELFREGGUEST,
- SPONSOR,
- SPONSOREDGUEST`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"settings": &schema.Schema{
							Description: `Defines all of the settings groups available for a portal`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"aup_settings": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"display_frequency": &schema.Schema{
													Description: `How the AUP should be displayed, either on page or as a link. Only valid if includeAup = true. Allowed Values:
- FIRSTLOGIN,
- EVERYLOGIN,
- RECURRING`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"display_frequency_interval_days": &schema.Schema{
													Description: `Number of days between AUP confirmations (when displayFrequency = recurring)`,
													Type:        schema.TypeInt,
													Computed:    true,
												},
												"include_aup": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"require_scrolling": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"login_page_settings": &schema.Schema{
										Description: `Portal Login Page settings groups follow`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"aup_display": &schema.Schema{
													Description: `How the AUP should be displayed, either on page or as a link.
Only valid if includeAup = true.
Allowed values:
-  ONPAGE,
- ASLINK`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"include_aup": &schema.Schema{
													Description: `Include an Acceptable Use Policy (AUP) that should be displayed during login`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"max_failed_attempts_before_rate_limit": &schema.Schema{
													Description: `Maximum failed login attempts before rate limiting`,
													Type:        schema.TypeInt,
													Computed:    true,
												},
												"require_aup_acceptance": &schema.Schema{
													Description: `Require the portal user to accept the AUP.
Only valid if includeAup = true`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"require_aup_scrolling": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"social_configs": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"time_between_logins_during_rate_limit": &schema.Schema{
													Description: `Time between login attempts when rate limiting`,
													Type:        schema.TypeInt,
													Computed:    true,
												},
											},
										},
									},
									"portal_settings": &schema.Schema{
										Description: `The port, interface, certificate, and other basic settings of a portal`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"allowed_interfaces": &schema.Schema{
													Description: `Interfaces that the portal will be reachable on.
Allowed values:
- eth0,
- eth1,
- eth2,
- eth3,
- eth4,
- eth5,
- bond0,
- bond1,
- bond2`,
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"authentication_method": &schema.Schema{
													Description: `Unique Id of the identity source sequence`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"available_ssids": &schema.Schema{
													Description: `Names of the SSIDs available for assignment to guest users by sponsors`,
													Type:        schema.TypeList,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"certificate_group_tag": &schema.Schema{
													Description: `Logical name of the x.509 server certificate that will be used for the portal`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"display_lang": &schema.Schema{
													Description: `Allowed values:
- USEBROWSERLOCALE,
- ALWAYSUSE`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"fallback_language": &schema.Schema{
													Description: `Used when displayLang = USEBROWSERLOCALE`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"fqdn": &schema.Schema{
													Description: `The fully-qualified domain name (FQDN) that end-users will use to access this portal. Used only in Sponsor portal`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"https_port": &schema.Schema{
													Description: `The port number that the allowed interfaces will listen on.
Range from 8000 to 8999`,
													Type:     schema.TypeInt,
													Computed: true,
												},
												"idle_timeout": &schema.Schema{
													Type:     schema.TypeInt,
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
													Description: `Include a Post-Login Banner page`,
													Type:        schema.TypeString,
													Computed:    true,
												},
											},
										},
									},
									"sponsor_change_password_settings": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"allow_sponsor_to_change_pwd": &schema.Schema{
													Description: `Allow sponsors to change their own passwords`,
													Type:        schema.TypeString,
													Computed:    true,
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
													Description: `The default value displayed for an empty field.
Only valid when emptyFieldDisplay = DISPLAYWITHDEFAULTVALUE`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"empty_field_display": &schema.Schema{
													Description: `Specifies how empty fields are handled on the Support Information Page. Allowed values:
- HIDE,
- DISPLAYWITHNOVALUE,
- DISPLAYWITHDEFAULTVALUE`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"include_browser_user_agent": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"include_failure_code": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"include_ip_address": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"include_mac_addr": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"include_policy_server": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"include_support_info_page": &schema.Schema{
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

func dataSourceSponsorPortalRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		log.Printf("[DEBUG] Selected method: GetSponsorPortal")
		queryParams1 := isegosdk.GetSponsorPortalQueryParams{}

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

		response1, restyResp1, err := client.SponsorPortal.GetSponsorPortal(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetSponsorPortal", err,
				"Failure at GetSponsorPortal, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		var items1 []isegosdk.ResponseSponsorPortalGetSponsorPortalSearchResultResources
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
				response1, _, err = client.SponsorPortal.GetSponsorPortal(&queryParams1)
				if err != nil {
					break
				}
				// All is good, continue to the next page
				continue
			}
			// Does not have next page finish iteration
			break
		}
		vItems1 := flattenSponsorPortalGetSponsorPortalItems(&items1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSponsorPortal response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetSponsorPortalByID")
		vvID := vID.(string)

		response2, restyResp2, err := client.SponsorPortal.GetSponsorPortalByID(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetSponsorPortalByID", err,
				"Failure at GetSponsorPortalByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenSponsorPortalGetSponsorPortalByIDItem(response2.SponsorPortal)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSponsorPortalByID response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSponsorPortalGetSponsorPortalItems(items *[]isegosdk.ResponseSponsorPortalGetSponsorPortalSearchResultResources) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItem["description"] = item.Description
		respItem["link"] = flattenSponsorPortalGetSponsorPortalItemsLink(item.Link)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSponsorPortalGetSponsorPortalItemsLink(item *isegosdk.ResponseSponsorPortalGetSponsorPortalSearchResultResourcesLink) []map[string]interface{} {
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

func flattenSponsorPortalGetSponsorPortalByIDItem(item *isegosdk.ResponseSponsorPortalGetSponsorPortalByIDSponsorPortal) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["description"] = item.Description
	respItem["portal_type"] = item.PortalType
	respItem["portal_test_url"] = item.PortalTestURL
	respItem["settings"] = flattenSponsorPortalGetSponsorPortalByIDItemSettings(item.Settings)
	respItem["customizations"] = flattenSponsorPortalGetSponsorPortalByIDItemCustomizations(item.Customizations)
	respItem["link"] = flattenSponsorPortalGetSponsorPortalByIDItemLink(item.Link)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenSponsorPortalGetSponsorPortalByIDItemSettings(item *isegosdk.ResponseSponsorPortalGetSponsorPortalByIDSponsorPortalSettings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["portal_settings"] = flattenSponsorPortalGetSponsorPortalByIDItemSettingsPortalSettings(item.PortalSettings)
	respItem["login_page_settings"] = flattenSponsorPortalGetSponsorPortalByIDItemSettingsLoginPageSettings(item.LoginPageSettings)
	respItem["aup_settings"] = flattenSponsorPortalGetSponsorPortalByIDItemSettingsAupSettings(item.AupSettings)
	respItem["sponsor_change_password_settings"] = flattenSponsorPortalGetSponsorPortalByIDItemSettingsSponsorChangePasswordSettings(item.SponsorChangePasswordSettings)
	respItem["post_login_banner_settings"] = flattenSponsorPortalGetSponsorPortalByIDItemSettingsPostLoginBannerSettings(item.PostLoginBannerSettings)
	respItem["post_access_banner_settings"] = flattenSponsorPortalGetSponsorPortalByIDItemSettingsPostAccessBannerSettings(item.PostAccessBannerSettings)
	respItem["support_info_settings"] = flattenSponsorPortalGetSponsorPortalByIDItemSettingsSupportInfoSettings(item.SupportInfoSettings)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSponsorPortalGetSponsorPortalByIDItemSettingsPortalSettings(item *isegosdk.ResponseSponsorPortalGetSponsorPortalByIDSponsorPortalSettingsPortalSettings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["https_port"] = item.HTTPSPort
	respItem["allowed_interfaces"] = item.AllowedInterfaces
	respItem["certificate_group_tag"] = item.CertificateGroupTag
	respItem["fqdn"] = item.Fqdn
	respItem["authentication_method"] = item.AuthenticationMethod
	respItem["idle_timeout"] = item.IDleTimeout
	respItem["display_lang"] = item.DisplayLang
	respItem["fallback_language"] = item.FallbackLanguage
	respItem["available_ssids"] = item.AvailableSSIDs

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSponsorPortalGetSponsorPortalByIDItemSettingsLoginPageSettings(item *isegosdk.ResponseSponsorPortalGetSponsorPortalByIDSponsorPortalSettingsLoginPageSettings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["max_failed_attempts_before_rate_limit"] = item.MaxFailedAttemptsBeforeRateLimit
	respItem["time_between_logins_during_rate_limit"] = item.TimeBetweenLoginsDuringRateLimit
	respItem["include_aup"] = boolPtrToString(item.IncludeAup)
	respItem["aup_display"] = item.AupDisplay
	respItem["require_aup_acceptance"] = boolPtrToString(item.RequireAupAcceptance)
	respItem["require_aup_scrolling"] = boolPtrToString(item.RequireAupScrolling)
	respItem["social_configs"] = responseInterfaceToSliceString(item.SocialConfigs)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSponsorPortalGetSponsorPortalByIDItemSettingsAupSettings(item *isegosdk.ResponseSponsorPortalGetSponsorPortalByIDSponsorPortalSettingsAupSettings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["include_aup"] = boolPtrToString(item.IncludeAup)
	respItem["require_scrolling"] = boolPtrToString(item.RequireScrolling)
	respItem["display_frequency"] = item.DisplayFrequency
	respItem["display_frequency_interval_days"] = item.DisplayFrequencyIntervalDays

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSponsorPortalGetSponsorPortalByIDItemSettingsSponsorChangePasswordSettings(item *isegosdk.ResponseSponsorPortalGetSponsorPortalByIDSponsorPortalSettingsSponsorChangePasswordSettings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["allow_sponsor_to_change_pwd"] = boolPtrToString(item.AllowSponsorToChangePwd)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSponsorPortalGetSponsorPortalByIDItemSettingsPostLoginBannerSettings(item *isegosdk.ResponseSponsorPortalGetSponsorPortalByIDSponsorPortalSettingsPostLoginBannerSettings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["include_post_access_banner"] = boolPtrToString(item.IncludePostAccessBanner)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSponsorPortalGetSponsorPortalByIDItemSettingsPostAccessBannerSettings(item *isegosdk.ResponseSponsorPortalGetSponsorPortalByIDSponsorPortalSettingsPostAccessBannerSettings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["include_post_access_banner"] = boolPtrToString(item.IncludePostAccessBanner)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSponsorPortalGetSponsorPortalByIDItemSettingsSupportInfoSettings(item *isegosdk.ResponseSponsorPortalGetSponsorPortalByIDSponsorPortalSettingsSupportInfoSettings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["include_support_info_page"] = boolPtrToString(item.IncludeSupportInfoPage)
	respItem["include_mac_addr"] = boolPtrToString(item.IncludeMacAddr)
	respItem["include_ip_address"] = boolPtrToString(item.IncludeIPAddress)
	respItem["include_browser_user_agent"] = boolPtrToString(item.IncludeBrowserUserAgent)
	respItem["include_policy_server"] = boolPtrToString(item.IncludePolicyServer)
	respItem["include_failure_code"] = boolPtrToString(item.IncludeFailureCode)
	respItem["empty_field_display"] = item.EmptyFieldDisplay
	respItem["default_empty_field_value"] = item.DefaultEmptyFieldValue

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSponsorPortalGetSponsorPortalByIDItemCustomizations(item *isegosdk.ResponseSponsorPortalGetSponsorPortalByIDSponsorPortalCustomizations) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["portal_theme"] = flattenSponsorPortalGetSponsorPortalByIDItemCustomizationsPortalTheme(item.PortalTheme)
	respItem["portal_tweak_settings"] = flattenSponsorPortalGetSponsorPortalByIDItemCustomizationsPortalTweakSettings(item.PortalTweakSettings)
	respItem["language"] = flattenSponsorPortalGetSponsorPortalByIDItemCustomizationsLanguage(item.Language)
	respItem["global_customizations"] = flattenSponsorPortalGetSponsorPortalByIDItemCustomizationsGlobalCustomizations(item.GlobalCustomizations)
	respItem["page_customizations"] = flattenSponsorPortalGetSponsorPortalByIDItemCustomizationsPageCustomizations(item.PageCustomizations)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSponsorPortalGetSponsorPortalByIDItemCustomizationsPortalTheme(item *isegosdk.ResponseSponsorPortalGetSponsorPortalByIDSponsorPortalCustomizationsPortalTheme) []map[string]interface{} {
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

func flattenSponsorPortalGetSponsorPortalByIDItemCustomizationsPortalTweakSettings(item *isegosdk.ResponseSponsorPortalGetSponsorPortalByIDSponsorPortalCustomizationsPortalTweakSettings) []map[string]interface{} {
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

func flattenSponsorPortalGetSponsorPortalByIDItemCustomizationsLanguage(item *isegosdk.ResponseSponsorPortalGetSponsorPortalByIDSponsorPortalCustomizationsLanguage) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["view_language"] = item.ViewLanguage

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSponsorPortalGetSponsorPortalByIDItemCustomizationsGlobalCustomizations(item *isegosdk.ResponseSponsorPortalGetSponsorPortalByIDSponsorPortalCustomizationsGlobalCustomizations) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["mobile_logo_image"] = flattenSponsorPortalGetSponsorPortalByIDItemCustomizationsGlobalCustomizationsMobileLogoImage(item.MobileLogoImage)
	respItem["desktop_logo_image"] = flattenSponsorPortalGetSponsorPortalByIDItemCustomizationsGlobalCustomizationsDesktopLogoImage(item.DesktopLogoImage)
	respItem["banner_image"] = flattenSponsorPortalGetSponsorPortalByIDItemCustomizationsGlobalCustomizationsBannerImage(item.BannerImage)
	respItem["background_image"] = flattenSponsorPortalGetSponsorPortalByIDItemCustomizationsGlobalCustomizationsBackgroundImage(item.BackgroundImage)
	respItem["banner_title"] = item.BannerTitle
	respItem["contact_text"] = item.ContactText
	respItem["footer_element"] = item.FooterElement

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSponsorPortalGetSponsorPortalByIDItemCustomizationsGlobalCustomizationsMobileLogoImage(item *isegosdk.ResponseSponsorPortalGetSponsorPortalByIDSponsorPortalCustomizationsGlobalCustomizationsMobileLogoImage) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["data"] = item.Data

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSponsorPortalGetSponsorPortalByIDItemCustomizationsGlobalCustomizationsDesktopLogoImage(item *isegosdk.ResponseSponsorPortalGetSponsorPortalByIDSponsorPortalCustomizationsGlobalCustomizationsDesktopLogoImage) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["data"] = item.Data

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSponsorPortalGetSponsorPortalByIDItemCustomizationsGlobalCustomizationsBannerImage(item *isegosdk.ResponseSponsorPortalGetSponsorPortalByIDSponsorPortalCustomizationsGlobalCustomizationsBannerImage) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["data"] = item.Data

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSponsorPortalGetSponsorPortalByIDItemCustomizationsGlobalCustomizationsBackgroundImage(item *isegosdk.ResponseSponsorPortalGetSponsorPortalByIDSponsorPortalCustomizationsGlobalCustomizationsBackgroundImage) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["data"] = item.Data

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSponsorPortalGetSponsorPortalByIDItemCustomizationsPageCustomizations(item *isegosdk.ResponseSponsorPortalGetSponsorPortalByIDSponsorPortalCustomizationsPageCustomizations) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["data"] = flattenSponsorPortalGetSponsorPortalByIDItemCustomizationsPageCustomizationsData(item.Data)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSponsorPortalGetSponsorPortalByIDItemCustomizationsPageCustomizationsData(items *[]isegosdk.ResponseSponsorPortalGetSponsorPortalByIDSponsorPortalCustomizationsPageCustomizationsData) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["key"] = item.Key
		respItem["value"] = item.Value
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSponsorPortalGetSponsorPortalByIDItemLink(item *isegosdk.ResponseSponsorPortalGetSponsorPortalByIDSponsorPortalLink) []map[string]interface{} {
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
