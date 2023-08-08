package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSponsoredGuestPortal() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on SponsoredGuestPortal.

- This data source allows the client to get a sponsored guest portal by ID.

- This data source allows the client to get all the sponsored guest portals.

Filter:

[name, description]

Sorting:

[name, description]
`,

		ReadContext: dataSourceSponsoredGuestPortalRead,
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
							Description: `Defines all of the Portal Customizations available`,
							Type:        schema.TypeList,
							Computed:    true,
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
												"require_aup_scrolling": &schema.Schema{
													Description: `Require the portal user to scroll to the end of the AUP. Only valid if requireAupAcceptance = true`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"require_scrolling": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"skip_aup_for_employees": &schema.Schema{
													Description: `Only valid if requireAupAcceptance = true`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"use_diff_aup_for_employees": &schema.Schema{
													Description: `Only valid if requireAupAcceptance = true`,
													Type:        schema.TypeString,
													Computed:    true,
												},
											},
										},
									},
									"auth_success_settings": &schema.Schema{
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
													Description: `After an Authentication Success where should device be redirected.
Allowed values:
- AUTHSUCCESSPAGE,
- ORIGINATINGURL,
- URL`,
													Type:     schema.TypeString,
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

												"byod_registration_settings": &schema.Schema{
													Description: `Configuration of BYOD endpoint Registration step configuration`,
													Type:        schema.TypeList,
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"end_point_identity_group_id": &schema.Schema{
																Description: `Identity group id for which endpoint belongs`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"show_device_id": &schema.Schema{
																Description: `Display Device ID field during registration`,
																Type:        schema.TypeString,
																Computed:    true,
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
																Description: `After an Authentication Success where should device be redirected. Allowed values:
- AUTHSUCCESSPAGE,
- ORIGINATINGURL,
- URL`,
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
																Description: `Require BYOD devices to scroll down to the bottom of the AUP.
Only valid if includeAup = true`,
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
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
													Description: `Allow guest to change their own passwords`,
													Type:        schema.TypeString,
													Computed:    true,
												},
											},
										},
									},
									"guest_device_registration_settings": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"allow_guests_to_register_devices": &schema.Schema{
													Description: `Allow guests to register devices`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"auto_register_guest_devices": &schema.Schema{
													Description: `Automatically register guest devices`,
													Type:        schema.TypeString,
													Computed:    true,
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

												"access_code": &schema.Schema{
													Description: `Access code that must be entered by the portal user (only valid if requireAccessCode = true)`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"allow_alternate_guest_portal": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"allow_forgot_password": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"allow_guest_to_change_password": &schema.Schema{
													Description: `Require the portal user to enter an access code`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"allow_guest_to_create_accounts": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
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
												"require_access_code": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"require_aup_acceptance": &schema.Schema{
													Description: `Require the portal user to accept the AUP.
Only valid if includeAup = true`,
													Type:     schema.TypeString,
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
												"always_used_language": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"assigned_guest_type_for_employee": &schema.Schema{
													Description: `Unique Id of a guest type.
Employees using this portal as a guest inherit login options from the guest type`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"authentication_method": &schema.Schema{
													Description: `Unique Id of the identity source sequence`,
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
												"fallback_language": &schema.Schema{
													Description: `Used when displayLang = USEBROWSERLOCALE`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"https_port": &schema.Schema{
													Description: `The port number that the allowed interfaces will listen on.
Range from 8000 to 8999`,
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

func dataSourceSponsoredGuestPortalRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		log.Printf("[DEBUG] Selected method: GetSponsoredGuestPortals")
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

		response1, restyResp1, err := client.SponsoredGuestPortal.GetSponsoredGuestPortals(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetSponsoredGuestPortals", err,
				"Failure at GetSponsoredGuestPortals, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		var items1 []isegosdk.ResponseSponsoredGuestPortalGetSponsoredGuestPortalsSearchResultResources
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
		log.Printf("[DEBUG] Selected method: GetSponsoredGuestPortalByID")
		vvID := vID.(string)

		response2, restyResp2, err := client.SponsoredGuestPortal.GetSponsoredGuestPortalByID(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetSponsoredGuestPortalByID", err,
				"Failure at GetSponsoredGuestPortalByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItem(response2.SponsoredGuestPortal)
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

func flattenSponsoredGuestPortalGetSponsoredGuestPortalsItemsLink(item *isegosdk.ResponseSponsoredGuestPortalGetSponsoredGuestPortalsSearchResultResourcesLink) []map[string]interface{} {
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

func flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemSettings(item *isegosdk.ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalSettings) []map[string]interface{} {
	if item == nil {
		return nil
	}
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

func flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemSettingsPortalSettings(item *isegosdk.ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalSettingsPortalSettings) []map[string]interface{} {
	if item == nil {
		return nil
	}
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

func flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemSettingsLoginPageSettings(item *isegosdk.ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalSettingsLoginPageSettings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["require_access_code"] = boolPtrToString(item.RequireAccessCode)
	respItem["max_failed_attempts_before_rate_limit"] = item.MaxFailedAttemptsBeforeRateLimit
	respItem["time_between_logins_during_rate_limit"] = item.TimeBetweenLoginsDuringRateLimit
	respItem["include_aup"] = boolPtrToString(item.IncludeAup)
	respItem["aup_display"] = item.AupDisplay
	respItem["require_aup_acceptance"] = boolPtrToString(item.RequireAupAcceptance)
	respItem["access_code"] = item.AccessCode
	respItem["allow_guest_to_create_accounts"] = boolPtrToString(item.AllowGuestToCreateAccounts)
	respItem["allow_forgot_password"] = boolPtrToString(item.AllowForgotPassword)
	respItem["allow_guest_to_change_password"] = boolPtrToString(item.AllowGuestToChangePassword)
	respItem["allow_alternate_guest_portal"] = boolPtrToString(item.AllowAlternateGuestPortal)
	respItem["social_configs"] = flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemSettingsLoginPageSettingsSocialConfigs(item.SocialConfigs)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemSettingsLoginPageSettingsSocialConfigs(items *[]isegosdk.ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalSettingsLoginPageSettingsSocialConfigs) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["social_media_type"] = item.SocialMediaType
		respItem["social_media_value"] = item.SocialMediaValue
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemSettingsAupSettings(item *isegosdk.ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalSettingsAupSettings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["include_aup"] = boolPtrToString(item.IncludeAup)
	respItem["require_aup_scrolling"] = boolPtrToString(item.RequireAupScrolling)
	respItem["use_diff_aup_for_employees"] = boolPtrToString(item.UseDiffAupForEmployees)
	respItem["skip_aup_for_employees"] = boolPtrToString(item.SkipAupForEmployees)
	respItem["display_frequency_interval_days"] = item.DisplayFrequencyIntervalDays
	respItem["require_scrolling"] = boolPtrToString(item.RequireScrolling)
	respItem["display_frequency"] = item.DisplayFrequency

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemSettingsGuestChangePasswordSettings(item *isegosdk.ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalSettingsGuestChangePasswordSettings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["allow_change_passwd_at_first_login"] = boolPtrToString(item.AllowChangePasswdAtFirstLogin)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemSettingsGuestDeviceRegistrationSettings(item *isegosdk.ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalSettingsGuestDeviceRegistrationSettings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["auto_register_guest_devices"] = boolPtrToString(item.AutoRegisterGuestDevices)
	respItem["allow_guests_to_register_devices"] = boolPtrToString(item.AllowGuestsToRegisterDevices)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemSettingsByodSettings(item *isegosdk.ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalSettingsByodSettings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["byod_welcome_settings"] = flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemSettingsByodSettingsByodWelcomeSettings(item.ByodWelcomeSettings)
	respItem["byod_registration_settings"] = flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemSettingsByodSettingsByodRegistrationSettings(item.ByodRegistrationSettings)
	respItem["byod_registration_success_settings"] = flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemSettingsByodSettingsByodRegistrationSuccessSettings(item.ByodRegistrationSuccessSettings)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemSettingsByodSettingsByodWelcomeSettings(item *isegosdk.ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalSettingsByodSettingsByodWelcomeSettings) []map[string]interface{} {
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

func flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemSettingsByodSettingsByodRegistrationSettings(item *isegosdk.ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalSettingsByodSettingsByodRegistrationSettings) []map[string]interface{} {
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

func flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemSettingsByodSettingsByodRegistrationSuccessSettings(item *isegosdk.ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalSettingsByodSettingsByodRegistrationSuccessSettings) []map[string]interface{} {
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

func flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemSettingsPostAccessBannerSettings(item *isegosdk.ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalSettingsPostAccessBannerSettings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["include_post_access_banner"] = boolPtrToString(item.IncludePostAccessBanner)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemSettingsAuthSuccessSettings(item *isegosdk.ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalSettingsAuthSuccessSettings) []map[string]interface{} {
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

func flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemSettingsPostLoginBannerSettings(item *isegosdk.ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalSettingsPostLoginBannerSettings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["include_post_access_banner"] = boolPtrToString(item.IncludePostAccessBanner)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemSettingsSupportInfoSettings(item *isegosdk.ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalSettingsSupportInfoSettings) []map[string]interface{} {
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

func flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemCustomizations(item *isegosdk.ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalCustomizations) []map[string]interface{} {
	if item == nil {
		return nil
	}
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

func flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemCustomizationsPortalTheme(item *isegosdk.ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsPortalTheme) []map[string]interface{} {
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

func flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemCustomizationsPortalTweakSettings(item *isegosdk.ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsPortalTweakSettings) []map[string]interface{} {
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

func flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemCustomizationsLanguage(item *isegosdk.ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsLanguage) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["view_language"] = item.ViewLanguage

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemCustomizationsGlobalCustomizations(item *isegosdk.ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsGlobalCustomizations) []map[string]interface{} {
	if item == nil {
		return nil
	}
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

func flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemCustomizationsGlobalCustomizationsMobileLogoImage(item *isegosdk.ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsGlobalCustomizationsMobileLogoImage) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["data"] = item.Data

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemCustomizationsGlobalCustomizationsDesktopLogoImage(item *isegosdk.ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsGlobalCustomizationsDesktopLogoImage) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["data"] = item.Data

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemCustomizationsGlobalCustomizationsBannerImage(item *isegosdk.ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsGlobalCustomizationsBannerImage) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["data"] = item.Data

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemCustomizationsGlobalCustomizationsBackgroundImage(item *isegosdk.ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsGlobalCustomizationsBackgroundImage) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["data"] = item.Data

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemCustomizationsPageCustomizations(item *isegosdk.ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsPageCustomizations) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["data"] = flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemCustomizationsPageCustomizationsData(item.Data)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemCustomizationsPageCustomizationsData(items *[]isegosdk.ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsPageCustomizationsData) []map[string]interface{} {
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

func flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItemLink(item *isegosdk.ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortalLink) []map[string]interface{} {
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
