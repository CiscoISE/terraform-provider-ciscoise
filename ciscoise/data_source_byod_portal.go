package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/kuba-mazurkiewicz/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceByodPortal() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on BYODPortal.

- This data source allows the client to get a BYOD portal by ID.

- This data source allows the client to get all the BYOD portals.

Filter:

[name, description]

Sorting:

[name, description]
`,

		ReadContext: dataSourceByodPortalRead,
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
				Description: `id path parameter. Portal id`,
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
							Description: `Defines all of the Portal Customizations available for a BYOD`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"global_customizations": &schema.Schema{
										Description: `Represent the portal Global customizations`,
										Type:        schema.TypeList,
										Computed:    true,
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
										Description: `This property is supported only for Read operation and it allows to show the customizations in English.
Other languages are not supported`,
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
										Description: `Represent the entire page customization as a giant dictionary`,
										Type:        schema.TypeList,
										Computed:    true,
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
										Description: `Defines the configuration for portal theme`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"id": &schema.Schema{
													Description: `The unique internal identifier of the portal theme`,
													Type:        schema.TypeString,
													Computed:    true,
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
							Description: `Resource UUID, mandatory for update`,
							Type:        schema.TypeString,
							Computed:    true,
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
							Description: `Resource Name`,
							Type:        schema.TypeString,
							Computed:    true,
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
							Description: `Defines all of the settings groups available for a BYOD`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"byod_settings": &schema.Schema{
										Description: `Configuration of BYOD Device Welcome, Registration and Success steps`,
										Type:        schema.TypeList,
										Computed:    true,
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

															"redirect_url": &schema.Schema{
																Description: `Target URL for redirection, used when successRedirect = URL`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"success_redirect": &schema.Schema{
																Description: `After an Authentication Success where should device be redirected. Allowed values:`,
																Type:        schema.TypeString,
																Computed:    true,
															},
														},
													},
												},
												"byod_welcome_settings": &schema.Schema{
													Description: `Configuration of BYOD endpoint welcome step configuration`,
													Type:        schema.TypeList,
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"aup_display": &schema.Schema{
																Description: `How the AUP should be displayed, either on page or as a link.
Only valid if includeAup = true.
Allowed values:
- ONPAGE,
- ASLINK`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"enable_byo_d": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"enable_guest_access": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"include_aup": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"require_aup_acceptance": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"require_mdm": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"require_scrolling": &schema.Schema{
																Description: `Require BYOD devices to scroll down to the bottom of the AUP, Only valid if includeAup = true`,
																Type:        schema.TypeString,
																Computed:    true,
															},
														},
													},
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
												"always_used_language": &schema.Schema{
													Description: `Used when displayLang = ALWAYSUSE`,
													Type:        schema.TypeString,
													Computed:    true,
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
												"endpoint_identity_group": &schema.Schema{
													Description: `Unique Id of the endpoint identity group where user's devices will be added. Used only in Hotspot Portal`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"fallback_language": &schema.Schema{
													Description: `Used when displayLang = USEBROWSERLOCALE`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"https_port": &schema.Schema{
													Description: `The port number that the allowed interfaces will listen on. Range from 8000 to 8999`,
													Type:        schema.TypeInt,
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
													Description: `Specifies how empty fields are handled on the Support Information Page.
Allowed values:
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

func dataSourceByodPortalRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

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
		log.Printf("[DEBUG] Selected method: GetByodPortal")
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

		response1, restyResp1, err := client.ByodPortal.GetByodPortal(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetByodPortal", err,
				"Failure at GetByodPortal, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

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
		log.Printf("[DEBUG] Selected method: GetByodPortalByID")
		vvID := vID.(string)

		response2, restyResp2, err := client.ByodPortal.GetByodPortalByID(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetByodPortalByID", err,
				"Failure at GetByodPortalByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

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
	respItem["enable_byo_d"] = boolPtrToString(item.EnableByod)
	respItem["enable_guest_access"] = boolPtrToString(item.EnableGuestAccess)
	respItem["require_mdm"] = boolPtrToString(item.RequireMdm)
	respItem["include_aup"] = boolPtrToString(item.IncludeAup)
	respItem["aup_display"] = item.AupDisplay
	respItem["require_aup_acceptance"] = boolPtrToString(item.RequireAupAcceptance)
	respItem["require_scrolling"] = boolPtrToString(item.RequireScrolling)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenByodPortalGetByodPortalByIDItemSettingsByodSettingsByodRegistrationSettings(item *isegosdk.ResponseByodPortalGetByodPortalByIDByodPortalSettingsByodSettingsByodRegistrationSettings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["show_device_id"] = boolPtrToString(item.ShowDeviceID)
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
		respItems = append(respItems, respItem)
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
