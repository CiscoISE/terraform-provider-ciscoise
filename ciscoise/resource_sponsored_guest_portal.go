package ciscoise

import (
	"context"
	"fmt"
	"reflect"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSponsoredGuestPortal() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on SponsoredGuestPortal.

- This resource allows the client to update a sponsored guest portal by ID.

- This resource deletes a sponsored guest portal by ID.

- This resource creates a sponsored guest portal.
`,

		CreateContext: resourceSponsoredGuestPortalCreate,
		ReadContext:   resourceSponsoredGuestPortalRead,
		UpdateContext: resourceSponsoredGuestPortalUpdate,
		DeleteContext: resourceSponsoredGuestPortalDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
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
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"customizations": &schema.Schema{
							Description: `Defines all of the Portal Customizations available`,
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"global_customizations": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"background_image": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"data": &schema.Schema{
																Description: `Represented as base 64 encoded string of the image byte array`,
																Type:        schema.TypeString,
																Optional:    true,
															},
														},
													},
												},
												"banner_image": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"data": &schema.Schema{
																Description: `Represented as base 64 encoded string of the image byte array`,
																Type:        schema.TypeString,
																Optional:    true,
															},
														},
													},
												},
												"banner_title": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"contact_text": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"desktop_logo_image": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"data": &schema.Schema{
																Description: `Represented as base 64 encoded string of the image byte array`,
																Type:        schema.TypeString,
																Optional:    true,
															},
														},
													},
												},
												"footer_element": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"mobile_logo_image": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"data": &schema.Schema{
																Description: `Represented as base 64 encoded string of the image byte array`,
																Type:        schema.TypeString,
																Optional:    true,
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
										Optional:    true,
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"view_language": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"page_customizations": &schema.Schema{
										Description: `Represent the entire page customization as a giant dictionary`,
										Type:        schema.TypeList,
										Optional:    true,
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"data": &schema.Schema{
													Description: `The Dictionary will be exposed here as key value pair`,
													Type:        schema.TypeList,
													Optional:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"key": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
															"value": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},
											},
										},
									},
									"portal_theme": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"id": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"name": &schema.Schema{
													Description: `The system- or user-assigned name of the portal theme`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"theme_data": &schema.Schema{
													Description: `A CSS file, represented as a Base64-encoded byte array`,
													Type:        schema.TypeString,
													Optional:    true,
												},
											},
										},
									},
									"portal_tweak_settings": &schema.Schema{
										Description: `The Tweak Settings are a customization of the Portal Theme that has been selected for the portal.
When the Portal Theme selection is changed, the Tweak Settings are overwritten to match the values in the theme.
The Tweak Settings can subsequently be changed by the user`,
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"banner_color": &schema.Schema{
													Description: `Hex value of color`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"banner_text_color": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"page_background_color": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"page_label_and_text_color": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},
						"description": &schema.Schema{
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
						"portal_test_url": &schema.Schema{
							Description: `URL to bring up a test page for this portal`,
							Type:        schema.TypeString,
							Optional:    true,
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
							Optional: true,
						},
						"settings": &schema.Schema{
							Description: `Defines all of the settings groups available for a portal`,
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"aup_settings": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"display_frequency": &schema.Schema{
													Description: `How the AUP should be displayed, either on page or as a link. Only valid if includeAup = true. Allowed Values:
- FIRSTLOGIN,
- EVERYLOGIN,
- RECURRING`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"display_frequency_interval_days": &schema.Schema{
													Description: `Number of days between AUP confirmations (when displayFrequency = recurring)`,
													Type:        schema.TypeInt,
													Optional:    true,
												},
												"include_aup": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"require_aup_scrolling": &schema.Schema{
													Description:  `Require the portal user to scroll to the end of the AUP. Only valid if requireAupAcceptance = true`,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"require_scrolling": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"skip_aup_for_employees": &schema.Schema{
													Description:  `Only valid if requireAupAcceptance = true`,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"use_diff_aup_for_employees": &schema.Schema{
													Description:  `Only valid if requireAupAcceptance = true`,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
											},
										},
									},
									"auth_success_settings": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"redirect_url": &schema.Schema{
													Description: `Target URL for redirection, used when successRedirect = URL`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"success_redirect": &schema.Schema{
													Description: `After an Authentication Success where should device be redirected.
Allowed values:
- AUTHSUCCESSPAGE,
- ORIGINATINGURL,
- URL`,
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"byod_settings": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"byod_registration_settings": &schema.Schema{
													Description: `Configuration of BYOD endpoint Registration step configuration`,
													Type:        schema.TypeList,
													Optional:    true,
													MaxItems:    1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"end_point_identity_group_id": &schema.Schema{
																Description: `Identity group id for which endpoint belongs`,
																Type:        schema.TypeString,
																Optional:    true,
															},
															"show_device_id": &schema.Schema{
																Description:  `Display Device ID field during registration`,
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
															},
														},
													},
												},
												"byod_registration_success_settings": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"redirect_url": &schema.Schema{
																Description: `Target URL for redirection, used when successRedirect = URL`,
																Type:        schema.TypeString,
																Optional:    true,
															},
															"success_redirect": &schema.Schema{
																Description: `After an Authentication Success where should device be redirected. Allowed values:
- AUTHSUCCESSPAGE,
- ORIGINATINGURL,
- URL`,
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},
												"byod_welcome_settings": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"aup_display": &schema.Schema{
																Description: `How the AUP should be displayed, either on page or as a link.
Only valid if includeAup = true.
Allowed values:
- ONPAGE,
- ASLINK`,
																Type:     schema.TypeString,
																Optional: true,
															},
															"enable_byo_d": &schema.Schema{
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
															},
															"enable_guest_access": &schema.Schema{
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
															},
															"include_aup": &schema.Schema{
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
															},
															"require_aup_acceptance": &schema.Schema{
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
															},
															"require_mdm": &schema.Schema{
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
															},
															"require_scrolling": &schema.Schema{
																Description: `Require BYOD devices to scroll down to the bottom of the AUP.
Only valid if includeAup = true`,
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
															},
														},
													},
												},
											},
										},
									},
									"guest_change_password_settings": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"allow_change_passwd_at_first_login": &schema.Schema{
													Description:  `Allow guest to change their own passwords`,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
											},
										},
									},
									"guest_device_registration_settings": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"allow_guests_to_register_devices": &schema.Schema{
													Description:  `Allow guests to register devices`,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"auto_register_guest_devices": &schema.Schema{
													Description:  `Automatically register guest devices`,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
											},
										},
									},
									"login_page_settings": &schema.Schema{
										Description: `Portal Login Page settings groups follow`,
										Type:        schema.TypeList,
										Optional:    true,
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"access_code": &schema.Schema{
													Description: `Access code that must be entered by the portal user (only valid if requireAccessCode = true)`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"allow_alternate_guest_portal": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"allow_forgot_password": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"allow_guest_to_change_password": &schema.Schema{
													Description:  `Require the portal user to enter an access code`,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"allow_guest_to_create_accounts": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"aup_display": &schema.Schema{
													Description: `How the AUP should be displayed, either on page or as a link.
Only valid if includeAup = true.
Allowed values:
-  ONPAGE,
- ASLINK`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"include_aup": &schema.Schema{
													Description:  `Include an Acceptable Use Policy (AUP) that should be displayed during login`,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"max_failed_attempts_before_rate_limit": &schema.Schema{
													Description: `Maximum failed login attempts before rate limiting`,
													Type:        schema.TypeInt,
													Optional:    true,
												},
												"require_access_code": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"require_aup_acceptance": &schema.Schema{
													Description: `Require the portal user to accept the AUP.
Only valid if includeAup = true`,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"social_configs": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"social_media_type": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
															"social_media_value": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},
												"time_between_logins_during_rate_limit": &schema.Schema{
													Description: `Time between login attempts when rate limiting`,
													Type:        schema.TypeInt,
													Optional:    true,
												},
											},
										},
									},
									"portal_settings": &schema.Schema{
										Description: `The port, interface, certificate, and other basic settings of a portal`,
										Type:        schema.TypeList,
										Optional:    true,
										MaxItems:    1,
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
													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"always_used_language": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"assigned_guest_type_for_employee": &schema.Schema{
													Description: `Unique Id of a guest type.
Employees using this portal as a guest inherit login options from the guest type`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"authentication_method": &schema.Schema{
													Description: `Unique Id of the identity source sequence`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"certificate_group_tag": &schema.Schema{
													Description: `Logical name of the x.509 server certificate that will be used for the portal`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"display_lang": &schema.Schema{
													Description: `Allowed values:
- USEBROWSERLOCALE,
- ALWAYSUSE`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"fallback_language": &schema.Schema{
													Description: `Used when displayLang = USEBROWSERLOCALE`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"https_port": &schema.Schema{
													Description: `The port number that the allowed interfaces will listen on.
Range from 8000 to 8999`,
													Type:     schema.TypeInt,
													Optional: true,
												},
											},
										},
									},
									"post_access_banner_settings": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"include_post_access_banner": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
											},
										},
									},
									"post_login_banner_settings": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"include_post_access_banner": &schema.Schema{
													Description:  `Include a Post-Login Banner page`,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
											},
										},
									},
									"support_info_settings": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"default_empty_field_value": &schema.Schema{
													Description: `The default value displayed for an empty field.
Only valid when emptyFieldDisplay = DISPLAYWITHDEFAULTVALUE`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"empty_field_display": &schema.Schema{
													Description: `Specifies how empty fields are handled on the Support Information Page. Allowed values:
- HIDE,
- DISPLAYWITHNOVALUE,
- DISPLAYWITHDEFAULTVALUE`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"include_browser_user_agent": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"include_failure_code": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"include_ip_address": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"include_mac_addr": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"include_policy_server": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"include_support_info_page": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
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
		},
	}
}

func resourceSponsoredGuestPortalCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning SponsoredGuestPortal create")
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestSponsoredGuestPortalCreateSponsoredGuestPortal(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vName, _ := resourceItem["name"]
	vvName := interfaceToString(vName)
	if okID && vvID != "" {
		getResponse2, _, err := client.SponsoredGuestPortal.GetSponsoredGuestPortalByID(vvID)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return resourceSponsoredGuestPortalRead(ctx, d, m)
		}
	} else {
		queryParams2 := isegosdk.GetSponsoredGuestPortalsQueryParams{}

		response2, _, err := client.SponsoredGuestPortal.GetSponsoredGuestPortals(&queryParams2)
		if response2 != nil && err == nil {
			items2 := getAllItemsSponsoredGuestPortalGetSponsoredGuestPortals(m, response2, &queryParams2)
			item2, err := searchSponsoredGuestPortalGetSponsoredGuestPortals(m, items2, vvName, vvID)
			if err == nil && item2 != nil {
				resourceMap := make(map[string]string)
				resourceMap["id"] = vvID
				resourceMap["name"] = vvName
				d.SetId(joinResourceID(resourceMap))
				return resourceSponsoredGuestPortalRead(ctx, d, m)
			}
		}
	}
	restyResp1, err := client.SponsoredGuestPortal.CreateSponsoredGuestPortal(request1)
	if err != nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateSponsoredGuestPortal", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateSponsoredGuestPortal", err))
		return diags
	}
	headers := restyResp1.Header()
	if locationHeader, ok := headers["Location"]; ok && len(locationHeader) > 0 {
		vvID = getLocationID(locationHeader[0])
	}
	resourceMap := make(map[string]string)
	resourceMap["id"] = vvID
	resourceMap["name"] = vvName
	d.SetId(joinResourceID(resourceMap))
	return resourceSponsoredGuestPortalRead(ctx, d, m)
}

func resourceSponsoredGuestPortalRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning SponsoredGuestPortal read for id=[%s]", d.Id())
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID, okID := resourceMap["id"]
	vName, okName := resourceMap["name"]

	method1 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 2 {
		vvID := vID
		vvName := vName
		log.Printf("[DEBUG] Selected method: GetSponsoredGuestPortals")
		queryParams1 := isegosdk.GetSponsoredGuestPortalsQueryParams{}

		response1, restyResp1, err := client.SponsoredGuestPortal.GetSponsoredGuestPortals(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		items1 := getAllItemsSponsoredGuestPortalGetSponsoredGuestPortals(m, response1, &queryParams1)
		item1, err := searchSponsoredGuestPortalGetSponsoredGuestPortals(m, items1, vvName, vvID)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}
		vItem1 := flattenSponsoredGuestPortalGetSponsoredGuestPortalByIDItem(item1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSponsoredGuestPortals search response",
				err))
			return diags
		}

	}
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetSponsoredGuestPortalByID")
		vvID := vID

		response2, restyResp2, err := client.SponsoredGuestPortal.GetSponsoredGuestPortalByID(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			d.SetId("")
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
		return diags

	}
	return diags
}

func resourceSponsoredGuestPortalUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning SponsoredGuestPortal update for id=[%s]", d.Id())
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vID, okID := resourceMap["id"]
	vName, okName := resourceMap["name"]

	method1 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	var vvID string
	// NOTE: Added getAllItems and search function to get missing params
	if selectedMethod == 2 {
		queryParams1 := isegosdk.GetSponsoredGuestPortalsQueryParams{}

		getResp1, _, err := client.SponsoredGuestPortal.GetSponsoredGuestPortals(&queryParams1)
		if err == nil && getResp1 != nil {
			items1 := getAllItemsSponsoredGuestPortalGetSponsoredGuestPortals(m, getResp1, &queryParams1)
			item1, err := searchSponsoredGuestPortalGetSponsoredGuestPortals(m, items1, vName, vID)
			if err == nil && item1 != nil {
				if vID != item1.ID {
					vvID = item1.ID
				} else {
					vvID = vID
				}
			}
		}
	}
	if selectedMethod == 1 {
		vvID = vID
	}
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] ID used for update operation %s", vvID)
		request1 := expandRequestSponsoredGuestPortalUpdateSponsoredGuestPortalByID(ctx, "parameters.0", d)
		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
		response1, restyResp1, err := client.SponsoredGuestPortal.UpdateSponsoredGuestPortalByID(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateSponsoredGuestPortalByID", err, restyResp1.String(),
					"Failure at UpdateSponsoredGuestPortalByID, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateSponsoredGuestPortalByID", err,
				"Failure at UpdateSponsoredGuestPortalByID, unexpected response", ""))
			return diags
		}
	}

	return resourceSponsoredGuestPortalRead(ctx, d, m)
}

func resourceSponsoredGuestPortalDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning SponsoredGuestPortal delete for id=[%s]", d.Id())
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID, okID := resourceMap["id"]
	vName, okName := resourceMap["name"]

	method1 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	var vvID string
	// REVIEW: Add getAllItems and search function to get missing params
	if selectedMethod == 2 {
		queryParams1 := isegosdk.GetSponsoredGuestPortalsQueryParams{}

		getResp1, _, err := client.SponsoredGuestPortal.GetSponsoredGuestPortals(&queryParams1)
		if err != nil || getResp1 == nil {
			// Assume that element it is already gone
			return diags
		}
		items1 := getAllItemsSponsoredGuestPortalGetSponsoredGuestPortals(m, getResp1, &queryParams1)
		item1, err := searchSponsoredGuestPortalGetSponsoredGuestPortals(m, items1, vName, vID)
		if err != nil || item1 == nil {
			// Assume that element it is already gone
			return diags
		}
		if vID != item1.ID {
			vvID = item1.ID
		} else {
			vvID = vID
		}
	}
	if selectedMethod == 1 {
		vvID = vID
		getResp, _, err := client.SponsoredGuestPortal.GetSponsoredGuestPortalByID(vvID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	restyResp1, err := client.SponsoredGuestPortal.DeleteSponsoredGuestPortalByID(vvID)
	if err != nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteSponsoredGuestPortalByID", err, restyResp1.String(),
				"Failure at DeleteSponsoredGuestPortalByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteSponsoredGuestPortalByID", err,
			"Failure at DeleteSponsoredGuestPortalByID, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestSponsoredGuestPortalCreateSponsoredGuestPortal(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsoredGuestPortalCreateSponsoredGuestPortal {
	request := isegosdk.RequestSponsoredGuestPortalCreateSponsoredGuestPortal{}
	request.SponsoredGuestPortal = expandRequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortal(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortal(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortal {
	request := isegosdk.RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortal{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".portal_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".portal_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".portal_type")))) {
		request.PortalType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".portal_test_url")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".portal_test_url")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".portal_test_url")))) {
		request.PortalTestURL = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".settings")))) {
		request.Settings = expandRequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettings(ctx, key+".settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".customizations")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".customizations")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".customizations")))) {
		request.Customizations = expandRequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizations(ctx, key+".customizations.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettings {
	request := isegosdk.RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".portal_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".portal_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".portal_settings")))) {
		request.PortalSettings = expandRequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsPortalSettings(ctx, key+".portal_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".login_page_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".login_page_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".login_page_settings")))) {
		request.LoginPageSettings = expandRequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsLoginPageSettings(ctx, key+".login_page_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".aup_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".aup_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".aup_settings")))) {
		request.AupSettings = expandRequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsAupSettings(ctx, key+".aup_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".guest_change_password_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".guest_change_password_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".guest_change_password_settings")))) {
		request.GuestChangePasswordSettings = expandRequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsGuestChangePasswordSettings(ctx, key+".guest_change_password_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".guest_device_registration_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".guest_device_registration_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".guest_device_registration_settings")))) {
		request.GuestDeviceRegistrationSettings = expandRequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsGuestDeviceRegistrationSettings(ctx, key+".guest_device_registration_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".byod_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".byod_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".byod_settings")))) {
		request.ByodSettings = expandRequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsByodSettings(ctx, key+".byod_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".post_access_banner_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".post_access_banner_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".post_access_banner_settings")))) {
		request.PostAccessBannerSettings = expandRequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsPostAccessBannerSettings(ctx, key+".post_access_banner_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".auth_success_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".auth_success_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".auth_success_settings")))) {
		request.AuthSuccessSettings = expandRequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsAuthSuccessSettings(ctx, key+".auth_success_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".post_login_banner_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".post_login_banner_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".post_login_banner_settings")))) {
		request.PostLoginBannerSettings = expandRequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsPostLoginBannerSettings(ctx, key+".post_login_banner_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".support_info_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".support_info_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".support_info_settings")))) {
		request.SupportInfoSettings = expandRequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsSupportInfoSettings(ctx, key+".support_info_settings.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsPortalSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsPortalSettings {
	request := isegosdk.RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsPortalSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".https_port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".https_port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".https_port")))) {
		request.HTTPSPort = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allowed_interfaces")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allowed_interfaces")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allowed_interfaces")))) {
		request.AllowedInterfaces = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".certificate_group_tag")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".certificate_group_tag")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".certificate_group_tag")))) {
		request.CertificateGroupTag = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".authentication_method")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".authentication_method")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".authentication_method")))) {
		request.AuthenticationMethod = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".assigned_guest_type_for_employee")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".assigned_guest_type_for_employee")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".assigned_guest_type_for_employee")))) {
		request.AssignedGuestTypeForEmployee = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".display_lang")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".display_lang")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".display_lang")))) {
		request.DisplayLang = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".fallback_language")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".fallback_language")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".fallback_language")))) {
		request.FallbackLanguage = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".always_used_language")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".always_used_language")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".always_used_language")))) {
		request.AlwaysUsedLanguage = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsLoginPageSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsLoginPageSettings {
	request := isegosdk.RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsLoginPageSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".require_access_code")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".require_access_code")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".require_access_code")))) {
		request.RequireAccessCode = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".max_failed_attempts_before_rate_limit")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".max_failed_attempts_before_rate_limit")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".max_failed_attempts_before_rate_limit")))) {
		request.MaxFailedAttemptsBeforeRateLimit = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".time_between_logins_during_rate_limit")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".time_between_logins_during_rate_limit")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".time_between_logins_during_rate_limit")))) {
		request.TimeBetweenLoginsDuringRateLimit = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include_aup")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include_aup")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include_aup")))) {
		request.IncludeAup = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".aup_display")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".aup_display")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".aup_display")))) {
		request.AupDisplay = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".require_aup_acceptance")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".require_aup_acceptance")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".require_aup_acceptance")))) {
		request.RequireAupAcceptance = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".access_code")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".access_code")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".access_code")))) {
		request.AccessCode = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_guest_to_create_accounts")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_guest_to_create_accounts")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_guest_to_create_accounts")))) {
		request.AllowGuestToCreateAccounts = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_forgot_password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_forgot_password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_forgot_password")))) {
		request.AllowForgotPassword = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_guest_to_change_password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_guest_to_change_password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_guest_to_change_password")))) {
		request.AllowGuestToChangePassword = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_alternate_guest_portal")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_alternate_guest_portal")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_alternate_guest_portal")))) {
		request.AllowAlternateGuestPortal = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".social_configs")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".social_configs")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".social_configs")))) {
		request.SocialConfigs = expandRequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsLoginPageSettingsSocialConfigsArray(ctx, key+".social_configs", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsLoginPageSettingsSocialConfigsArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsLoginPageSettingsSocialConfigs {
	request := []isegosdk.RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsLoginPageSettingsSocialConfigs{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsLoginPageSettingsSocialConfigs(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsLoginPageSettingsSocialConfigs(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsLoginPageSettingsSocialConfigs {
	request := isegosdk.RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsLoginPageSettingsSocialConfigs{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".social_media_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".social_media_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".social_media_type")))) {
		request.SocialMediaType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".social_media_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".social_media_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".social_media_value")))) {
		request.SocialMediaValue = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsAupSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsAupSettings {
	request := isegosdk.RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsAupSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include_aup")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include_aup")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include_aup")))) {
		request.IncludeAup = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".require_aup_scrolling")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".require_aup_scrolling")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".require_aup_scrolling")))) {
		request.RequireAupScrolling = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".use_diff_aup_for_employees")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".use_diff_aup_for_employees")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".use_diff_aup_for_employees")))) {
		request.UseDiffAupForEmployees = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".skip_aup_for_employees")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".skip_aup_for_employees")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".skip_aup_for_employees")))) {
		request.SkipAupForEmployees = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".display_frequency_interval_days")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".display_frequency_interval_days")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".display_frequency_interval_days")))) {
		request.DisplayFrequencyIntervalDays = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".require_scrolling")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".require_scrolling")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".require_scrolling")))) {
		request.RequireScrolling = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".display_frequency")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".display_frequency")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".display_frequency")))) {
		request.DisplayFrequency = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsGuestChangePasswordSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsGuestChangePasswordSettings {
	request := isegosdk.RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsGuestChangePasswordSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_change_passwd_at_first_login")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_change_passwd_at_first_login")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_change_passwd_at_first_login")))) {
		request.AllowChangePasswdAtFirstLogin = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsGuestDeviceRegistrationSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsGuestDeviceRegistrationSettings {
	request := isegosdk.RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsGuestDeviceRegistrationSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".auto_register_guest_devices")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".auto_register_guest_devices")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".auto_register_guest_devices")))) {
		request.AutoRegisterGuestDevices = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_guests_to_register_devices")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_guests_to_register_devices")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_guests_to_register_devices")))) {
		request.AllowGuestsToRegisterDevices = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsByodSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsByodSettings {
	request := isegosdk.RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsByodSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".byod_welcome_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".byod_welcome_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".byod_welcome_settings")))) {
		request.ByodWelcomeSettings = expandRequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsByodSettingsByodWelcomeSettings(ctx, key+".byod_welcome_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".byod_registration_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".byod_registration_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".byod_registration_settings")))) {
		request.ByodRegistrationSettings = expandRequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsByodSettingsByodRegistrationSettings(ctx, key+".byod_registration_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".byod_registration_success_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".byod_registration_success_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".byod_registration_success_settings")))) {
		request.ByodRegistrationSuccessSettings = expandRequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsByodSettingsByodRegistrationSuccessSettings(ctx, key+".byod_registration_success_settings.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsByodSettingsByodWelcomeSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsByodSettingsByodWelcomeSettings {
	request := isegosdk.RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsByodSettingsByodWelcomeSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_byo_d")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_byo_d")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_byo_d")))) {
		request.EnableByod = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_guest_access")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_guest_access")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_guest_access")))) {
		request.EnableGuestAccess = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".require_mdm")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".require_mdm")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".require_mdm")))) {
		request.RequireMdm = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include_aup")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include_aup")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include_aup")))) {
		request.IncludeAup = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".aup_display")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".aup_display")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".aup_display")))) {
		request.AupDisplay = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".require_aup_acceptance")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".require_aup_acceptance")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".require_aup_acceptance")))) {
		request.RequireAupAcceptance = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".require_scrolling")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".require_scrolling")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".require_scrolling")))) {
		request.RequireScrolling = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsByodSettingsByodRegistrationSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsByodSettingsByodRegistrationSettings {
	request := isegosdk.RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsByodSettingsByodRegistrationSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".show_device_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".show_device_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".show_device_id")))) {
		request.ShowDeviceID = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_point_identity_group_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_point_identity_group_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_point_identity_group_id")))) {
		request.EndPointIDentityGroupID = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsByodSettingsByodRegistrationSuccessSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsByodSettingsByodRegistrationSuccessSettings {
	request := isegosdk.RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsByodSettingsByodRegistrationSuccessSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".success_redirect")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".success_redirect")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".success_redirect")))) {
		request.SuccessRedirect = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".redirect_url")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".redirect_url")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".redirect_url")))) {
		request.RedirectURL = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsPostAccessBannerSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsPostAccessBannerSettings {
	request := isegosdk.RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsPostAccessBannerSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include_post_access_banner")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include_post_access_banner")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include_post_access_banner")))) {
		request.IncludePostAccessBanner = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsAuthSuccessSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsAuthSuccessSettings {
	request := isegosdk.RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsAuthSuccessSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".success_redirect")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".success_redirect")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".success_redirect")))) {
		request.SuccessRedirect = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".redirect_url")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".redirect_url")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".redirect_url")))) {
		request.RedirectURL = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsPostLoginBannerSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsPostLoginBannerSettings {
	request := isegosdk.RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsPostLoginBannerSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include_post_access_banner")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include_post_access_banner")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include_post_access_banner")))) {
		request.IncludePostAccessBanner = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsSupportInfoSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsSupportInfoSettings {
	request := isegosdk.RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalSettingsSupportInfoSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include_support_info_page")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include_support_info_page")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include_support_info_page")))) {
		request.IncludeSupportInfoPage = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include_mac_addr")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include_mac_addr")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include_mac_addr")))) {
		request.IncludeMacAddr = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include_ip_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include_ip_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include_ip_address")))) {
		request.IncludeIPAddress = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include_browser_user_agent")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include_browser_user_agent")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include_browser_user_agent")))) {
		request.IncludeBrowserUserAgent = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include_policy_server")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include_policy_server")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include_policy_server")))) {
		request.IncludePolicyServer = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include_failure_code")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include_failure_code")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include_failure_code")))) {
		request.IncludeFailureCode = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".empty_field_display")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".empty_field_display")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".empty_field_display")))) {
		request.EmptyFieldDisplay = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".default_empty_field_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".default_empty_field_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".default_empty_field_value")))) {
		request.DefaultEmptyFieldValue = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizations(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizations {
	request := isegosdk.RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizations{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".portal_theme")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".portal_theme")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".portal_theme")))) {
		request.PortalTheme = expandRequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizationsPortalTheme(ctx, key+".portal_theme.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".portal_tweak_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".portal_tweak_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".portal_tweak_settings")))) {
		request.PortalTweakSettings = expandRequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizationsPortalTweakSettings(ctx, key+".portal_tweak_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".language")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".language")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".language")))) {
		request.Language = expandRequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizationsLanguage(ctx, key+".language.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".global_customizations")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".global_customizations")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".global_customizations")))) {
		request.GlobalCustomizations = expandRequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizationsGlobalCustomizations(ctx, key+".global_customizations.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".page_customizations")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".page_customizations")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".page_customizations")))) {
		request.PageCustomizations = expandRequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizationsPageCustomizations(ctx, key+".page_customizations.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizationsPortalTheme(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizationsPortalTheme {
	request := isegosdk.RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizationsPortalTheme{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".theme_data")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".theme_data")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".theme_data")))) {
		request.ThemeData = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizationsPortalTweakSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizationsPortalTweakSettings {
	request := isegosdk.RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizationsPortalTweakSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".banner_color")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".banner_color")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".banner_color")))) {
		request.BannerColor = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".banner_text_color")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".banner_text_color")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".banner_text_color")))) {
		request.BannerTextColor = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".page_background_color")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".page_background_color")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".page_background_color")))) {
		request.PageBackgroundColor = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".page_label_and_text_color")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".page_label_and_text_color")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".page_label_and_text_color")))) {
		request.PageLabelAndTextColor = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizationsLanguage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizationsLanguage {
	request := isegosdk.RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizationsLanguage{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".view_language")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".view_language")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".view_language")))) {
		request.ViewLanguage = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizationsGlobalCustomizations(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizationsGlobalCustomizations {
	request := isegosdk.RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizationsGlobalCustomizations{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mobile_logo_image")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mobile_logo_image")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mobile_logo_image")))) {
		request.MobileLogoImage = expandRequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizationsGlobalCustomizationsMobileLogoImage(ctx, key+".mobile_logo_image.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".desktop_logo_image")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".desktop_logo_image")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".desktop_logo_image")))) {
		request.DesktopLogoImage = expandRequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizationsGlobalCustomizationsDesktopLogoImage(ctx, key+".desktop_logo_image.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".banner_image")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".banner_image")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".banner_image")))) {
		request.BannerImage = expandRequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizationsGlobalCustomizationsBannerImage(ctx, key+".banner_image.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".background_image")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".background_image")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".background_image")))) {
		request.BackgroundImage = expandRequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizationsGlobalCustomizationsBackgroundImage(ctx, key+".background_image.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".banner_title")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".banner_title")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".banner_title")))) {
		request.BannerTitle = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".contact_text")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".contact_text")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".contact_text")))) {
		request.ContactText = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".footer_element")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".footer_element")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".footer_element")))) {
		request.FooterElement = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizationsGlobalCustomizationsMobileLogoImage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizationsGlobalCustomizationsMobileLogoImage {
	request := isegosdk.RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizationsGlobalCustomizationsMobileLogoImage{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data")))) {
		request.Data = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizationsGlobalCustomizationsDesktopLogoImage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizationsGlobalCustomizationsDesktopLogoImage {
	request := isegosdk.RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizationsGlobalCustomizationsDesktopLogoImage{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data")))) {
		request.Data = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizationsGlobalCustomizationsBannerImage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizationsGlobalCustomizationsBannerImage {
	request := isegosdk.RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizationsGlobalCustomizationsBannerImage{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data")))) {
		request.Data = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizationsGlobalCustomizationsBackgroundImage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizationsGlobalCustomizationsBackgroundImage {
	request := isegosdk.RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizationsGlobalCustomizationsBackgroundImage{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data")))) {
		request.Data = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizationsPageCustomizations(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizationsPageCustomizations {
	request := isegosdk.RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizationsPageCustomizations{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data")))) {
		request.Data = expandRequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizationsPageCustomizationsDataArray(ctx, key+".data", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizationsPageCustomizationsDataArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizationsPageCustomizationsData {
	request := []isegosdk.RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizationsPageCustomizationsData{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizationsPageCustomizationsData(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizationsPageCustomizationsData(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizationsPageCustomizationsData {
	request := isegosdk.RequestSponsoredGuestPortalCreateSponsoredGuestPortalSponsoredGuestPortalCustomizationsPageCustomizationsData{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".key")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".key")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".key")))) {
		request.Key = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".value")))) {
		request.Value = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsoredGuestPortalUpdateSponsoredGuestPortalByID(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByID {
	request := isegosdk.RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByID{}
	request.SponsoredGuestPortal = expandRequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortal(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortal(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortal {
	request := isegosdk.RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortal{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".portal_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".portal_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".portal_type")))) {
		request.PortalType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".portal_test_url")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".portal_test_url")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".portal_test_url")))) {
		request.PortalTestURL = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".settings")))) {
		request.Settings = expandRequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettings(ctx, key+".settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".customizations")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".customizations")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".customizations")))) {
		request.Customizations = expandRequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizations(ctx, key+".customizations.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettings {
	request := isegosdk.RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".portal_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".portal_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".portal_settings")))) {
		request.PortalSettings = expandRequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsPortalSettings(ctx, key+".portal_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".login_page_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".login_page_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".login_page_settings")))) {
		request.LoginPageSettings = expandRequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsLoginPageSettings(ctx, key+".login_page_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".aup_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".aup_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".aup_settings")))) {
		request.AupSettings = expandRequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsAupSettings(ctx, key+".aup_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".guest_change_password_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".guest_change_password_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".guest_change_password_settings")))) {
		request.GuestChangePasswordSettings = expandRequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsGuestChangePasswordSettings(ctx, key+".guest_change_password_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".guest_device_registration_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".guest_device_registration_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".guest_device_registration_settings")))) {
		request.GuestDeviceRegistrationSettings = expandRequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsGuestDeviceRegistrationSettings(ctx, key+".guest_device_registration_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".byod_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".byod_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".byod_settings")))) {
		request.ByodSettings = expandRequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsByodSettings(ctx, key+".byod_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".post_access_banner_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".post_access_banner_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".post_access_banner_settings")))) {
		request.PostAccessBannerSettings = expandRequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsPostAccessBannerSettings(ctx, key+".post_access_banner_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".auth_success_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".auth_success_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".auth_success_settings")))) {
		request.AuthSuccessSettings = expandRequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsAuthSuccessSettings(ctx, key+".auth_success_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".post_login_banner_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".post_login_banner_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".post_login_banner_settings")))) {
		request.PostLoginBannerSettings = expandRequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsPostLoginBannerSettings(ctx, key+".post_login_banner_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".support_info_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".support_info_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".support_info_settings")))) {
		request.SupportInfoSettings = expandRequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsSupportInfoSettings(ctx, key+".support_info_settings.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsPortalSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsPortalSettings {
	request := isegosdk.RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsPortalSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".https_port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".https_port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".https_port")))) {
		request.HTTPSPort = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allowed_interfaces")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allowed_interfaces")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allowed_interfaces")))) {
		request.AllowedInterfaces = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".certificate_group_tag")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".certificate_group_tag")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".certificate_group_tag")))) {
		request.CertificateGroupTag = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".authentication_method")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".authentication_method")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".authentication_method")))) {
		request.AuthenticationMethod = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".assigned_guest_type_for_employee")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".assigned_guest_type_for_employee")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".assigned_guest_type_for_employee")))) {
		request.AssignedGuestTypeForEmployee = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".display_lang")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".display_lang")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".display_lang")))) {
		request.DisplayLang = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".fallback_language")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".fallback_language")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".fallback_language")))) {
		request.FallbackLanguage = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".always_used_language")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".always_used_language")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".always_used_language")))) {
		request.AlwaysUsedLanguage = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsLoginPageSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsLoginPageSettings {
	request := isegosdk.RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsLoginPageSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".require_access_code")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".require_access_code")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".require_access_code")))) {
		request.RequireAccessCode = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".max_failed_attempts_before_rate_limit")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".max_failed_attempts_before_rate_limit")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".max_failed_attempts_before_rate_limit")))) {
		request.MaxFailedAttemptsBeforeRateLimit = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".time_between_logins_during_rate_limit")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".time_between_logins_during_rate_limit")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".time_between_logins_during_rate_limit")))) {
		request.TimeBetweenLoginsDuringRateLimit = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include_aup")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include_aup")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include_aup")))) {
		request.IncludeAup = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".aup_display")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".aup_display")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".aup_display")))) {
		request.AupDisplay = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".require_aup_acceptance")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".require_aup_acceptance")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".require_aup_acceptance")))) {
		request.RequireAupAcceptance = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".access_code")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".access_code")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".access_code")))) {
		request.AccessCode = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_guest_to_create_accounts")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_guest_to_create_accounts")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_guest_to_create_accounts")))) {
		request.AllowGuestToCreateAccounts = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_forgot_password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_forgot_password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_forgot_password")))) {
		request.AllowForgotPassword = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_guest_to_change_password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_guest_to_change_password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_guest_to_change_password")))) {
		request.AllowGuestToChangePassword = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_alternate_guest_portal")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_alternate_guest_portal")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_alternate_guest_portal")))) {
		request.AllowAlternateGuestPortal = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".social_configs")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".social_configs")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".social_configs")))) {
		request.SocialConfigs = expandRequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsLoginPageSettingsSocialConfigsArray(ctx, key+".social_configs", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsLoginPageSettingsSocialConfigsArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsLoginPageSettingsSocialConfigs {
	request := []isegosdk.RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsLoginPageSettingsSocialConfigs{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsLoginPageSettingsSocialConfigs(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsLoginPageSettingsSocialConfigs(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsLoginPageSettingsSocialConfigs {
	request := isegosdk.RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsLoginPageSettingsSocialConfigs{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".social_media_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".social_media_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".social_media_type")))) {
		request.SocialMediaType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".social_media_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".social_media_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".social_media_value")))) {
		request.SocialMediaValue = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsAupSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsAupSettings {
	request := isegosdk.RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsAupSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include_aup")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include_aup")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include_aup")))) {
		request.IncludeAup = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".require_aup_scrolling")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".require_aup_scrolling")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".require_aup_scrolling")))) {
		request.RequireAupScrolling = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".use_diff_aup_for_employees")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".use_diff_aup_for_employees")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".use_diff_aup_for_employees")))) {
		request.UseDiffAupForEmployees = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".skip_aup_for_employees")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".skip_aup_for_employees")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".skip_aup_for_employees")))) {
		request.SkipAupForEmployees = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".display_frequency_interval_days")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".display_frequency_interval_days")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".display_frequency_interval_days")))) {
		request.DisplayFrequencyIntervalDays = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".require_scrolling")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".require_scrolling")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".require_scrolling")))) {
		request.RequireScrolling = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".display_frequency")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".display_frequency")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".display_frequency")))) {
		request.DisplayFrequency = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsGuestChangePasswordSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsGuestChangePasswordSettings {
	request := isegosdk.RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsGuestChangePasswordSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_change_passwd_at_first_login")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_change_passwd_at_first_login")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_change_passwd_at_first_login")))) {
		request.AllowChangePasswdAtFirstLogin = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsGuestDeviceRegistrationSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsGuestDeviceRegistrationSettings {
	request := isegosdk.RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsGuestDeviceRegistrationSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".auto_register_guest_devices")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".auto_register_guest_devices")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".auto_register_guest_devices")))) {
		request.AutoRegisterGuestDevices = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_guests_to_register_devices")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_guests_to_register_devices")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_guests_to_register_devices")))) {
		request.AllowGuestsToRegisterDevices = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsByodSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsByodSettings {
	request := isegosdk.RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsByodSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".byod_welcome_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".byod_welcome_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".byod_welcome_settings")))) {
		request.ByodWelcomeSettings = expandRequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsByodSettingsByodWelcomeSettings(ctx, key+".byod_welcome_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".byod_registration_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".byod_registration_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".byod_registration_settings")))) {
		request.ByodRegistrationSettings = expandRequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsByodSettingsByodRegistrationSettings(ctx, key+".byod_registration_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".byod_registration_success_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".byod_registration_success_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".byod_registration_success_settings")))) {
		request.ByodRegistrationSuccessSettings = expandRequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsByodSettingsByodRegistrationSuccessSettings(ctx, key+".byod_registration_success_settings.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsByodSettingsByodWelcomeSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsByodSettingsByodWelcomeSettings {
	request := isegosdk.RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsByodSettingsByodWelcomeSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_byo_d")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_byo_d")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_byo_d")))) {
		request.EnableByod = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_guest_access")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_guest_access")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_guest_access")))) {
		request.EnableGuestAccess = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".require_mdm")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".require_mdm")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".require_mdm")))) {
		request.RequireMdm = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include_aup")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include_aup")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include_aup")))) {
		request.IncludeAup = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".aup_display")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".aup_display")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".aup_display")))) {
		request.AupDisplay = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".require_aup_acceptance")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".require_aup_acceptance")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".require_aup_acceptance")))) {
		request.RequireAupAcceptance = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".require_scrolling")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".require_scrolling")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".require_scrolling")))) {
		request.RequireScrolling = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsByodSettingsByodRegistrationSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsByodSettingsByodRegistrationSettings {
	request := isegosdk.RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsByodSettingsByodRegistrationSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".show_device_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".show_device_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".show_device_id")))) {
		request.ShowDeviceID = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_point_identity_group_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_point_identity_group_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_point_identity_group_id")))) {
		request.EndPointIDentityGroupID = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsByodSettingsByodRegistrationSuccessSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsByodSettingsByodRegistrationSuccessSettings {
	request := isegosdk.RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsByodSettingsByodRegistrationSuccessSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".success_redirect")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".success_redirect")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".success_redirect")))) {
		request.SuccessRedirect = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".redirect_url")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".redirect_url")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".redirect_url")))) {
		request.RedirectURL = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsPostAccessBannerSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsPostAccessBannerSettings {
	request := isegosdk.RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsPostAccessBannerSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include_post_access_banner")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include_post_access_banner")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include_post_access_banner")))) {
		request.IncludePostAccessBanner = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsAuthSuccessSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsAuthSuccessSettings {
	request := isegosdk.RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsAuthSuccessSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".success_redirect")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".success_redirect")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".success_redirect")))) {
		request.SuccessRedirect = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".redirect_url")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".redirect_url")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".redirect_url")))) {
		request.RedirectURL = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsPostLoginBannerSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsPostLoginBannerSettings {
	request := isegosdk.RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsPostLoginBannerSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include_post_access_banner")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include_post_access_banner")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include_post_access_banner")))) {
		request.IncludePostAccessBanner = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsSupportInfoSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsSupportInfoSettings {
	request := isegosdk.RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalSettingsSupportInfoSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include_support_info_page")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include_support_info_page")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include_support_info_page")))) {
		request.IncludeSupportInfoPage = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include_mac_addr")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include_mac_addr")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include_mac_addr")))) {
		request.IncludeMacAddr = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include_ip_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include_ip_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include_ip_address")))) {
		request.IncludeIPAddress = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include_browser_user_agent")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include_browser_user_agent")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include_browser_user_agent")))) {
		request.IncludeBrowserUserAgent = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include_policy_server")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include_policy_server")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include_policy_server")))) {
		request.IncludePolicyServer = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include_failure_code")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include_failure_code")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include_failure_code")))) {
		request.IncludeFailureCode = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".empty_field_display")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".empty_field_display")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".empty_field_display")))) {
		request.EmptyFieldDisplay = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".default_empty_field_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".default_empty_field_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".default_empty_field_value")))) {
		request.DefaultEmptyFieldValue = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizations(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizations {
	request := isegosdk.RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizations{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".portal_theme")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".portal_theme")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".portal_theme")))) {
		request.PortalTheme = expandRequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsPortalTheme(ctx, key+".portal_theme.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".portal_tweak_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".portal_tweak_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".portal_tweak_settings")))) {
		request.PortalTweakSettings = expandRequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsPortalTweakSettings(ctx, key+".portal_tweak_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".language")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".language")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".language")))) {
		request.Language = expandRequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsLanguage(ctx, key+".language.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".global_customizations")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".global_customizations")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".global_customizations")))) {
		request.GlobalCustomizations = expandRequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsGlobalCustomizations(ctx, key+".global_customizations.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".page_customizations")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".page_customizations")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".page_customizations")))) {
		request.PageCustomizations = expandRequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsPageCustomizations(ctx, key+".page_customizations.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsPortalTheme(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsPortalTheme {
	request := isegosdk.RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsPortalTheme{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".theme_data")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".theme_data")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".theme_data")))) {
		request.ThemeData = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsPortalTweakSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsPortalTweakSettings {
	request := isegosdk.RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsPortalTweakSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".banner_color")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".banner_color")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".banner_color")))) {
		request.BannerColor = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".banner_text_color")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".banner_text_color")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".banner_text_color")))) {
		request.BannerTextColor = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".page_background_color")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".page_background_color")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".page_background_color")))) {
		request.PageBackgroundColor = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".page_label_and_text_color")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".page_label_and_text_color")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".page_label_and_text_color")))) {
		request.PageLabelAndTextColor = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsLanguage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsLanguage {
	request := isegosdk.RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsLanguage{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".view_language")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".view_language")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".view_language")))) {
		request.ViewLanguage = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsGlobalCustomizations(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsGlobalCustomizations {
	request := isegosdk.RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsGlobalCustomizations{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mobile_logo_image")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mobile_logo_image")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mobile_logo_image")))) {
		request.MobileLogoImage = expandRequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsGlobalCustomizationsMobileLogoImage(ctx, key+".mobile_logo_image.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".desktop_logo_image")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".desktop_logo_image")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".desktop_logo_image")))) {
		request.DesktopLogoImage = expandRequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsGlobalCustomizationsDesktopLogoImage(ctx, key+".desktop_logo_image.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".banner_image")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".banner_image")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".banner_image")))) {
		request.BannerImage = expandRequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsGlobalCustomizationsBannerImage(ctx, key+".banner_image.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".background_image")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".background_image")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".background_image")))) {
		request.BackgroundImage = expandRequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsGlobalCustomizationsBackgroundImage(ctx, key+".background_image.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".banner_title")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".banner_title")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".banner_title")))) {
		request.BannerTitle = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".contact_text")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".contact_text")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".contact_text")))) {
		request.ContactText = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".footer_element")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".footer_element")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".footer_element")))) {
		request.FooterElement = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsGlobalCustomizationsMobileLogoImage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsGlobalCustomizationsMobileLogoImage {
	request := isegosdk.RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsGlobalCustomizationsMobileLogoImage{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data")))) {
		request.Data = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsGlobalCustomizationsDesktopLogoImage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsGlobalCustomizationsDesktopLogoImage {
	request := isegosdk.RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsGlobalCustomizationsDesktopLogoImage{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data")))) {
		request.Data = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsGlobalCustomizationsBannerImage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsGlobalCustomizationsBannerImage {
	request := isegosdk.RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsGlobalCustomizationsBannerImage{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data")))) {
		request.Data = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsGlobalCustomizationsBackgroundImage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsGlobalCustomizationsBackgroundImage {
	request := isegosdk.RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsGlobalCustomizationsBackgroundImage{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data")))) {
		request.Data = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsPageCustomizations(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsPageCustomizations {
	request := isegosdk.RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsPageCustomizations{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data")))) {
		request.Data = expandRequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsPageCustomizationsDataArray(ctx, key+".data", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsPageCustomizationsDataArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsPageCustomizationsData {
	request := []isegosdk.RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsPageCustomizationsData{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsPageCustomizationsData(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsPageCustomizationsData(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsPageCustomizationsData {
	request := isegosdk.RequestSponsoredGuestPortalUpdateSponsoredGuestPortalByIDSponsoredGuestPortalCustomizationsPageCustomizationsData{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".key")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".key")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".key")))) {
		request.Key = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".value")))) {
		request.Value = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func getAllItemsSponsoredGuestPortalGetSponsoredGuestPortals(m interface{}, response *isegosdk.ResponseSponsoredGuestPortalGetSponsoredGuestPortals, queryParams *isegosdk.GetSponsoredGuestPortalsQueryParams) []isegosdk.ResponseSponsoredGuestPortalGetSponsoredGuestPortalsSearchResultResources {
	client := m.(*isegosdk.Client)
	var respItems []isegosdk.ResponseSponsoredGuestPortalGetSponsoredGuestPortalsSearchResultResources
	for response.SearchResult != nil && response.SearchResult.Resources != nil && len(*response.SearchResult.Resources) > 0 {
		respItems = append(respItems, *response.SearchResult.Resources...)
		if response.SearchResult.NextPage != nil && response.SearchResult.NextPage.Rel == "next" {
			href := response.SearchResult.NextPage.Href
			page, size, err := getNextPageAndSizeParams(href)
			if err != nil {
				break
			}
			if queryParams != nil {
				queryParams.Page = page
				queryParams.Size = size
			}
			response, _, err = client.SponsoredGuestPortal.GetSponsoredGuestPortals(queryParams)
			if err != nil {
				break
			}
			// All is good, continue to the next page
			continue
		}
		// Does not have next page finish iteration
		break
	}
	return respItems
}

func searchSponsoredGuestPortalGetSponsoredGuestPortals(m interface{}, items []isegosdk.ResponseSponsoredGuestPortalGetSponsoredGuestPortalsSearchResultResources, name string, id string) (*isegosdk.ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortal, error) {
	client := m.(*isegosdk.Client)
	var err error
	var foundItem *isegosdk.ResponseSponsoredGuestPortalGetSponsoredGuestPortalByIDSponsoredGuestPortal
	for _, item := range items {
		if id != "" && item.ID == id {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseSponsoredGuestPortalGetSponsoredGuestPortalByID
			getItem, _, err = client.SponsoredGuestPortal.GetSponsoredGuestPortalByID(id)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetSponsoredGuestPortalByID")
			}
			foundItem = getItem.SponsoredGuestPortal
			return foundItem, err
		} else if name != "" && item.Name == name {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseSponsoredGuestPortalGetSponsoredGuestPortalByID
			getItem, _, err = client.SponsoredGuestPortal.GetSponsoredGuestPortalByID(item.ID)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetSponsoredGuestPortalByID")
			}
			foundItem = getItem.SponsoredGuestPortal
			return foundItem, err
		}
	}
	return foundItem, err
}
