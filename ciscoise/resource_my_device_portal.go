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

func resourceMyDevicePortal() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on MyDevicePortal.

- This resource allows the client to update a my device portal by ID.

- This resource deletes a my device portal by ID.

- This resource creates a my device portal.
`,

		CreateContext: resourceMyDevicePortalCreate,
		ReadContext:   resourceMyDevicePortalRead,
		UpdateContext: resourceMyDevicePortalUpdate,
		DeleteContext: resourceMyDevicePortalDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Description: `Unix timestamp records the last time that the resource was updated.`,
				Type:        schema.TypeString,
				Computed:    true,
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
							Description: `Defines all of the settings groups available for a Mydevice portal`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"aup_settings": &schema.Schema{
										Description: `Configuration of the Acceptable Use Policy (AUP) for a portal`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"display_frequency": &schema.Schema{
													Description: `How the AUP should be displayed, either on page or as a link. Only valid if includeAup = true.
Allowed Values:
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
													Description: `Require the portal user to read and accept an AUP`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"require_scrolling": &schema.Schema{
													Description: `Require the portal user to scroll to the end of the AUP. Only valid if requireAupAcceptance = true`,
													Type:        schema.TypeString,
													Computed:    true,
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
												"require_scrolling": &schema.Schema{
													Description: `Require the portal user to scroll to the end of the AUP. Only valid if requireAupAcceptance = true`,
													Type:        schema.TypeString,
													Computed:    true,
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
												"always_used_language": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
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
							Description: `Defines all of the settings groups available for a Mydevice portal`,
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"aup_settings": &schema.Schema{
										Description: `Configuration of the Acceptable Use Policy (AUP) for a portal`,
										Type:        schema.TypeList,
										Optional:    true,
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"display_frequency": &schema.Schema{
													Description: `How the AUP should be displayed, either on page or as a link. Only valid if includeAup = true.
Allowed Values:
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
													Description:  `Require the portal user to read and accept an AUP`,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"require_scrolling": &schema.Schema{
													Description:  `Require the portal user to scroll to the end of the AUP. Only valid if requireAupAcceptance = true`,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
											},
										},
									},
									"employee_change_password_settings": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"allow_employee_to_change_pwd": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
											},
										},
									},
									"login_page_settings": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

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
												"require_aup_acceptance": &schema.Schema{
													Description: `Require the portal user to accept the AUP.
Only valid if includeAup = true`,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"require_scrolling": &schema.Schema{
													Description:  `Require the portal user to scroll to the end of the AUP. Only valid if requireAupAcceptance = true`,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"social_configs": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
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
												"endpoint_identity_group": &schema.Schema{
													Description: `Unique Id of the endpoint identity group where user's devices will be added. Used only in Hotspot Portal`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"fallback_language": &schema.Schema{
													Description: `Used when displayLang = USEBROWSERLOCALE`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"https_port": &schema.Schema{
													Description: `The port number that the allowed interfaces will listen on. Range from 8000 to 8999`,
													Type:        schema.TypeInt,
													Optional:    true,
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

func resourceMyDevicePortalCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning MyDevicePortal create")
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestMyDevicePortalCreateMyDevicePortal(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vName, _ := resourceItem["name"]
	vvName := interfaceToString(vName)
	if okID && vvID != "" {
		getResponse2, _, err := client.MyDevicePortal.GetMyDevicePortalByID(vvID)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return resourceMyDevicePortalRead(ctx, d, m)
		}
	} else {
		queryParams2 := isegosdk.GetMyDevicePortalQueryParams{}

		response2, _, err := client.MyDevicePortal.GetMyDevicePortal(&queryParams2)
		if response2 != nil && err == nil {
			items2 := getAllItemsMyDevicePortalGetMyDevicePortal(m, response2, &queryParams2)
			item2, err := searchMyDevicePortalGetMyDevicePortal(m, items2, vvName, vvID)
			if err == nil && item2 != nil {
				resourceMap := make(map[string]string)
				resourceMap["id"] = vvID
				resourceMap["name"] = vvName
				d.SetId(joinResourceID(resourceMap))
				return resourceMyDevicePortalRead(ctx, d, m)
			}
		}
	}
	restyResp1, err := client.MyDevicePortal.CreateMyDevicePortal(request1)
	if err != nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateMyDevicePortal", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateMyDevicePortal", err))
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
	return resourceMyDevicePortalRead(ctx, d, m)
}

func resourceMyDevicePortalRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning MyDevicePortal read for id=[%s]", d.Id())
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
		vvName := vName
		vvID := vID
		log.Printf("[DEBUG] Selected method: GetMyDevicePortal")
		queryParams1 := isegosdk.GetMyDevicePortalQueryParams{}

		response1, restyResp1, err := client.MyDevicePortal.GetMyDevicePortal(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		items1 := getAllItemsMyDevicePortalGetMyDevicePortal(m, response1, &queryParams1)
		item1, err := searchMyDevicePortalGetMyDevicePortal(m, items1, vvName, vvID)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}
		vItem1 := flattenMyDevicePortalGetMyDevicePortalByIDItem(item1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetMyDevicePortal search response",
				err))
			return diags
		}

	}
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetMyDevicePortalByID")
		vvID := vID

		response2, restyResp2, err := client.MyDevicePortal.GetMyDevicePortalByID(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenMyDevicePortalGetMyDevicePortalByIDItem(response2.MyDevicePortal)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetMyDevicePortalByID response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceMyDevicePortalUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning MyDevicePortal update for id=[%s]", d.Id())
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
		queryParams1 := isegosdk.GetMyDevicePortalQueryParams{}

		getResp1, _, err := client.MyDevicePortal.GetMyDevicePortal(&queryParams1)
		if err == nil && getResp1 != nil {
			items1 := getAllItemsMyDevicePortalGetMyDevicePortal(m, getResp1, &queryParams1)
			item1, err := searchMyDevicePortalGetMyDevicePortal(m, items1, vName, vID)
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
		request1 := expandRequestMyDevicePortalUpdateMyDevicePortalByID(ctx, "parameters.0", d)
		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
		response1, restyResp1, err := client.MyDevicePortal.UpdateMyDevicePortalByID(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateMyDevicePortalByID", err, restyResp1.String(),
					"Failure at UpdateMyDevicePortalByID, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateMyDevicePortalByID", err,
				"Failure at UpdateMyDevicePortalByID, unexpected response", ""))
			return diags
		}
		d.Set("last_updated", getUnixTimeString())
	}

	return resourceMyDevicePortalRead(ctx, d, m)
}

func resourceMyDevicePortalDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning MyDevicePortal delete for id=[%s]", d.Id())
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
		queryParams1 := isegosdk.GetMyDevicePortalQueryParams{}

		getResp1, _, err := client.MyDevicePortal.GetMyDevicePortal(&queryParams1)
		if err != nil || getResp1 == nil {
			// Assume that element it is already gone
			return diags
		}
		items1 := getAllItemsMyDevicePortalGetMyDevicePortal(m, getResp1, &queryParams1)
		item1, err := searchMyDevicePortalGetMyDevicePortal(m, items1, vName, vID)
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
		getResp, _, err := client.MyDevicePortal.GetMyDevicePortalByID(vvID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	restyResp1, err := client.MyDevicePortal.DeleteMyDevicePortalByID(vvID)
	if err != nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteMyDevicePortalByID", err, restyResp1.String(),
				"Failure at DeleteMyDevicePortalByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteMyDevicePortalByID", err,
			"Failure at DeleteMyDevicePortalByID, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestMyDevicePortalCreateMyDevicePortal(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalCreateMyDevicePortal {
	request := isegosdk.RequestMyDevicePortalCreateMyDevicePortal{}
	request.MyDevicePortal = expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortal(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortal(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortal {
	request := isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortal{}
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
		request.Settings = expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettings(ctx, key+".settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".customizations")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".customizations")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".customizations")))) {
		request.Customizations = expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizations(ctx, key+".customizations.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettings {
	request := isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".portal_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".portal_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".portal_settings")))) {
		request.PortalSettings = expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsPortalSettings(ctx, key+".portal_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".login_page_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".login_page_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".login_page_settings")))) {
		request.LoginPageSettings = expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsLoginPageSettings(ctx, key+".login_page_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".aup_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".aup_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".aup_settings")))) {
		request.AupSettings = expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsAupSettings(ctx, key+".aup_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".employee_change_password_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".employee_change_password_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".employee_change_password_settings")))) {
		request.EmployeeChangePasswordSettings = expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsEmployeeChangePasswordSettings(ctx, key+".employee_change_password_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".post_login_banner_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".post_login_banner_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".post_login_banner_settings")))) {
		request.PostLoginBannerSettings = expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsPostLoginBannerSettings(ctx, key+".post_login_banner_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".post_access_banner_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".post_access_banner_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".post_access_banner_settings")))) {
		request.PostAccessBannerSettings = expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsPostAccessBannerSettings(ctx, key+".post_access_banner_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".support_info_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".support_info_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".support_info_settings")))) {
		request.SupportInfoSettings = expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsSupportInfoSettings(ctx, key+".support_info_settings.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsPortalSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsPortalSettings {
	request := isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsPortalSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".https_port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".https_port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".https_port")))) {
		request.HTTPSPort = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allowed_interfaces")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allowed_interfaces")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allowed_interfaces")))) {
		request.AllowedInterfaces = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".certificate_group_tag")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".certificate_group_tag")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".certificate_group_tag")))) {
		request.CertificateGroupTag = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".endpoint_identity_group")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".endpoint_identity_group")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".endpoint_identity_group")))) {
		request.EndpointIDentityGroup = interfaceToString(v)
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

func expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsLoginPageSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsLoginPageSettings {
	request := isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsLoginPageSettings{}
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
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".require_scrolling")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".require_scrolling")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".require_scrolling")))) {
		request.RequireScrolling = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".social_configs")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".social_configs")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".social_configs")))) {
		request.SocialConfigs = expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsLoginPageSettingsSocialConfigsArray(ctx, key+".social_configs", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsLoginPageSettingsSocialConfigsArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsLoginPageSettingsSocialConfigs {
	request := []isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsLoginPageSettingsSocialConfigs{}
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
		i := expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsLoginPageSettingsSocialConfigs(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsLoginPageSettingsSocialConfigs(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsLoginPageSettingsSocialConfigs {
	var request isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsLoginPageSettingsSocialConfigs
	keyValue := d.Get(fixKeyAccess(key))
	request = requestStringToInterface(interfaceToString(keyValue))
	return &request
}

func expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsAupSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsAupSettings {
	request := isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsAupSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".display_frequency_interval_days")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".display_frequency_interval_days")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".display_frequency_interval_days")))) {
		request.DisplayFrequencyIntervalDays = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".display_frequency")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".display_frequency")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".display_frequency")))) {
		request.DisplayFrequency = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include_aup")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include_aup")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include_aup")))) {
		request.IncludeAup = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".require_scrolling")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".require_scrolling")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".require_scrolling")))) {
		request.RequireScrolling = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsEmployeeChangePasswordSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsEmployeeChangePasswordSettings {
	request := isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsEmployeeChangePasswordSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_employee_to_change_pwd")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_employee_to_change_pwd")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_employee_to_change_pwd")))) {
		request.AllowEmployeeToChangePwd = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsPostLoginBannerSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsPostLoginBannerSettings {
	request := isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsPostLoginBannerSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include_post_access_banner")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include_post_access_banner")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include_post_access_banner")))) {
		request.IncludePostAccessBanner = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsPostAccessBannerSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsPostAccessBannerSettings {
	request := isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsPostAccessBannerSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include_post_access_banner")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include_post_access_banner")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include_post_access_banner")))) {
		request.IncludePostAccessBanner = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsSupportInfoSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsSupportInfoSettings {
	request := isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalSettingsSupportInfoSettings{}
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

func expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizations(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizations {
	request := isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizations{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".portal_theme")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".portal_theme")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".portal_theme")))) {
		request.PortalTheme = expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsPortalTheme(ctx, key+".portal_theme.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".portal_tweak_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".portal_tweak_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".portal_tweak_settings")))) {
		request.PortalTweakSettings = expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsPortalTweakSettings(ctx, key+".portal_tweak_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".language")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".language")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".language")))) {
		request.Language = expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsLanguage(ctx, key+".language.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".global_customizations")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".global_customizations")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".global_customizations")))) {
		request.GlobalCustomizations = expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsGlobalCustomizations(ctx, key+".global_customizations.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".page_customizations")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".page_customizations")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".page_customizations")))) {
		request.PageCustomizations = expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsPageCustomizations(ctx, key+".page_customizations.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsPortalTheme(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsPortalTheme {
	request := isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsPortalTheme{}
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

func expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsPortalTweakSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsPortalTweakSettings {
	request := isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsPortalTweakSettings{}
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

func expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsLanguage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsLanguage {
	request := isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsLanguage{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".view_language")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".view_language")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".view_language")))) {
		request.ViewLanguage = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsGlobalCustomizations(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsGlobalCustomizations {
	request := isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsGlobalCustomizations{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mobile_logo_image")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mobile_logo_image")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mobile_logo_image")))) {
		request.MobileLogoImage = expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsGlobalCustomizationsMobileLogoImage(ctx, key+".mobile_logo_image.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".desktop_logo_image")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".desktop_logo_image")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".desktop_logo_image")))) {
		request.DesktopLogoImage = expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsGlobalCustomizationsDesktopLogoImage(ctx, key+".desktop_logo_image.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".banner_image")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".banner_image")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".banner_image")))) {
		request.BannerImage = expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsGlobalCustomizationsBannerImage(ctx, key+".banner_image.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".background_image")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".background_image")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".background_image")))) {
		request.BackgroundImage = expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsGlobalCustomizationsBackgroundImage(ctx, key+".background_image.0", d)
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

func expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsGlobalCustomizationsMobileLogoImage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsGlobalCustomizationsMobileLogoImage {
	request := isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsGlobalCustomizationsMobileLogoImage{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data")))) {
		request.Data = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsGlobalCustomizationsDesktopLogoImage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsGlobalCustomizationsDesktopLogoImage {
	request := isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsGlobalCustomizationsDesktopLogoImage{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data")))) {
		request.Data = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsGlobalCustomizationsBannerImage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsGlobalCustomizationsBannerImage {
	request := isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsGlobalCustomizationsBannerImage{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data")))) {
		request.Data = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsGlobalCustomizationsBackgroundImage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsGlobalCustomizationsBackgroundImage {
	request := isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsGlobalCustomizationsBackgroundImage{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data")))) {
		request.Data = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsPageCustomizations(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsPageCustomizations {
	request := isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsPageCustomizations{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data")))) {
		request.Data = expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsPageCustomizationsDataArray(ctx, key+".data", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsPageCustomizationsDataArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsPageCustomizationsData {
	request := []isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsPageCustomizationsData{}
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
		i := expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsPageCustomizationsData(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsPageCustomizationsData(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsPageCustomizationsData {
	request := isegosdk.RequestMyDevicePortalCreateMyDevicePortalMyDevicePortalCustomizationsPageCustomizationsData{}
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

func expandRequestMyDevicePortalUpdateMyDevicePortalByID(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByID {
	request := isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByID{}
	request.MyDevicePortal = expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortal(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortal(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortal {
	request := isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortal{}
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
		request.Settings = expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettings(ctx, key+".settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".customizations")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".customizations")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".customizations")))) {
		request.Customizations = expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizations(ctx, key+".customizations.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettings {
	request := isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".portal_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".portal_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".portal_settings")))) {
		request.PortalSettings = expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsPortalSettings(ctx, key+".portal_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".login_page_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".login_page_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".login_page_settings")))) {
		request.LoginPageSettings = expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsLoginPageSettings(ctx, key+".login_page_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".aup_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".aup_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".aup_settings")))) {
		request.AupSettings = expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsAupSettings(ctx, key+".aup_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".employee_change_password_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".employee_change_password_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".employee_change_password_settings")))) {
		request.EmployeeChangePasswordSettings = expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsEmployeeChangePasswordSettings(ctx, key+".employee_change_password_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".post_login_banner_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".post_login_banner_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".post_login_banner_settings")))) {
		request.PostLoginBannerSettings = expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsPostLoginBannerSettings(ctx, key+".post_login_banner_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".post_access_banner_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".post_access_banner_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".post_access_banner_settings")))) {
		request.PostAccessBannerSettings = expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsPostAccessBannerSettings(ctx, key+".post_access_banner_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".support_info_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".support_info_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".support_info_settings")))) {
		request.SupportInfoSettings = expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsSupportInfoSettings(ctx, key+".support_info_settings.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsPortalSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsPortalSettings {
	request := isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsPortalSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".https_port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".https_port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".https_port")))) {
		request.HTTPSPort = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allowed_interfaces")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allowed_interfaces")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allowed_interfaces")))) {
		request.AllowedInterfaces = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".certificate_group_tag")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".certificate_group_tag")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".certificate_group_tag")))) {
		request.CertificateGroupTag = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".endpoint_identity_group")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".endpoint_identity_group")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".endpoint_identity_group")))) {
		request.EndpointIDentityGroup = interfaceToString(v)
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

func expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsLoginPageSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsLoginPageSettings {
	request := isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsLoginPageSettings{}
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
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".require_scrolling")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".require_scrolling")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".require_scrolling")))) {
		request.RequireScrolling = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".social_configs")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".social_configs")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".social_configs")))) {
		request.SocialConfigs = expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsLoginPageSettingsSocialConfigsArray(ctx, key+".social_configs", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsLoginPageSettingsSocialConfigsArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsLoginPageSettingsSocialConfigs {
	request := []isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsLoginPageSettingsSocialConfigs{}
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
		i := expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsLoginPageSettingsSocialConfigs(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsLoginPageSettingsSocialConfigs(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsLoginPageSettingsSocialConfigs {
	var request isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsLoginPageSettingsSocialConfigs
	keyValue := d.Get(fixKeyAccess(key))
	request = requestStringToInterface(interfaceToString(keyValue))
	return &request
}

func expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsAupSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsAupSettings {
	request := isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsAupSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".display_frequency_interval_days")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".display_frequency_interval_days")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".display_frequency_interval_days")))) {
		request.DisplayFrequencyIntervalDays = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".display_frequency")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".display_frequency")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".display_frequency")))) {
		request.DisplayFrequency = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include_aup")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include_aup")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include_aup")))) {
		request.IncludeAup = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".require_scrolling")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".require_scrolling")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".require_scrolling")))) {
		request.RequireScrolling = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsEmployeeChangePasswordSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsEmployeeChangePasswordSettings {
	request := isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsEmployeeChangePasswordSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_employee_to_change_pwd")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_employee_to_change_pwd")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_employee_to_change_pwd")))) {
		request.AllowEmployeeToChangePwd = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsPostLoginBannerSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsPostLoginBannerSettings {
	request := isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsPostLoginBannerSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include_post_access_banner")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include_post_access_banner")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include_post_access_banner")))) {
		request.IncludePostAccessBanner = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsPostAccessBannerSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsPostAccessBannerSettings {
	request := isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsPostAccessBannerSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include_post_access_banner")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include_post_access_banner")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include_post_access_banner")))) {
		request.IncludePostAccessBanner = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsSupportInfoSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsSupportInfoSettings {
	request := isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalSettingsSupportInfoSettings{}
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

func expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizations(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizations {
	request := isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizations{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".portal_theme")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".portal_theme")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".portal_theme")))) {
		request.PortalTheme = expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsPortalTheme(ctx, key+".portal_theme.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".portal_tweak_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".portal_tweak_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".portal_tweak_settings")))) {
		request.PortalTweakSettings = expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsPortalTweakSettings(ctx, key+".portal_tweak_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".language")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".language")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".language")))) {
		request.Language = expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsLanguage(ctx, key+".language.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".global_customizations")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".global_customizations")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".global_customizations")))) {
		request.GlobalCustomizations = expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsGlobalCustomizations(ctx, key+".global_customizations.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".page_customizations")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".page_customizations")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".page_customizations")))) {
		request.PageCustomizations = expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsPageCustomizations(ctx, key+".page_customizations.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsPortalTheme(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsPortalTheme {
	request := isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsPortalTheme{}
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

func expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsPortalTweakSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsPortalTweakSettings {
	request := isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsPortalTweakSettings{}
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

func expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsLanguage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsLanguage {
	request := isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsLanguage{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".view_language")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".view_language")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".view_language")))) {
		request.ViewLanguage = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsGlobalCustomizations(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsGlobalCustomizations {
	request := isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsGlobalCustomizations{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mobile_logo_image")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mobile_logo_image")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mobile_logo_image")))) {
		request.MobileLogoImage = expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsGlobalCustomizationsMobileLogoImage(ctx, key+".mobile_logo_image.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".desktop_logo_image")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".desktop_logo_image")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".desktop_logo_image")))) {
		request.DesktopLogoImage = expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsGlobalCustomizationsDesktopLogoImage(ctx, key+".desktop_logo_image.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".banner_image")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".banner_image")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".banner_image")))) {
		request.BannerImage = expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsGlobalCustomizationsBannerImage(ctx, key+".banner_image.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".background_image")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".background_image")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".background_image")))) {
		request.BackgroundImage = expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsGlobalCustomizationsBackgroundImage(ctx, key+".background_image.0", d)
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

func expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsGlobalCustomizationsMobileLogoImage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsGlobalCustomizationsMobileLogoImage {
	request := isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsGlobalCustomizationsMobileLogoImage{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data")))) {
		request.Data = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsGlobalCustomizationsDesktopLogoImage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsGlobalCustomizationsDesktopLogoImage {
	request := isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsGlobalCustomizationsDesktopLogoImage{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data")))) {
		request.Data = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsGlobalCustomizationsBannerImage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsGlobalCustomizationsBannerImage {
	request := isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsGlobalCustomizationsBannerImage{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data")))) {
		request.Data = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsGlobalCustomizationsBackgroundImage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsGlobalCustomizationsBackgroundImage {
	request := isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsGlobalCustomizationsBackgroundImage{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data")))) {
		request.Data = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsPageCustomizations(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsPageCustomizations {
	request := isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsPageCustomizations{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data")))) {
		request.Data = expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsPageCustomizationsDataArray(ctx, key+".data", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsPageCustomizationsDataArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsPageCustomizationsData {
	request := []isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsPageCustomizationsData{}
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
		i := expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsPageCustomizationsData(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsPageCustomizationsData(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsPageCustomizationsData {
	request := isegosdk.RequestMyDevicePortalUpdateMyDevicePortalByIDMyDevicePortalCustomizationsPageCustomizationsData{}
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

func getAllItemsMyDevicePortalGetMyDevicePortal(m interface{}, response *isegosdk.ResponseMyDevicePortalGetMyDevicePortal, queryParams *isegosdk.GetMyDevicePortalQueryParams) []isegosdk.ResponseMyDevicePortalGetMyDevicePortalSearchResultResources {
	client := m.(*isegosdk.Client)
	var respItems []isegosdk.ResponseMyDevicePortalGetMyDevicePortalSearchResultResources
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
			response, _, err = client.MyDevicePortal.GetMyDevicePortal(queryParams)
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

func searchMyDevicePortalGetMyDevicePortal(m interface{}, items []isegosdk.ResponseMyDevicePortalGetMyDevicePortalSearchResultResources, name string, id string) (*isegosdk.ResponseMyDevicePortalGetMyDevicePortalByIDMyDevicePortal, error) {
	client := m.(*isegosdk.Client)
	var err error
	var foundItem *isegosdk.ResponseMyDevicePortalGetMyDevicePortalByIDMyDevicePortal
	for _, item := range items {
		if id != "" && item.ID == id {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseMyDevicePortalGetMyDevicePortalByID
			getItem, _, err = client.MyDevicePortal.GetMyDevicePortalByID(id)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetMyDevicePortalByID")
			}
			foundItem = getItem.MyDevicePortal
			return foundItem, err
		} else if name != "" && item.Name == name {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseMyDevicePortalGetMyDevicePortalByID
			getItem, _, err = client.MyDevicePortal.GetMyDevicePortalByID(item.ID)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetMyDevicePortalByID")
			}
			foundItem = getItem.MyDevicePortal
			return foundItem, err
		}
	}
	return foundItem, err
}
