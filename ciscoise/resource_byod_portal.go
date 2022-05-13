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

func resourceByodPortal() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on BYODPortal.

- This resource allows the client to update a BYOD portal by ID.

- This resource deletes a BYOD portal by ID.

- This resource creates a BYOD portal.
`,

		CreateContext: resourceByodPortalCreate,
		ReadContext:   resourceByodPortalRead,
		UpdateContext: resourceByodPortalUpdate,
		DeleteContext: resourceByodPortalDelete,
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
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"customizations": &schema.Schema{
							Description: `Defines all of the Portal Customizations available for a BYOD`,
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"global_customizations": &schema.Schema{
										Description: `Represent the portal Global customizations`,
										Type:        schema.TypeList,
										Optional:    true,
										MaxItems:    1,
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
										Description: `This property is supported only for Read operation and it allows to show the customizations in English.
Other languages are not supported`,
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
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
										Description: `Defines the configuration for portal theme`,
										Type:        schema.TypeList,
										Optional:    true,
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"id": &schema.Schema{
													Description: `The unique internal identifier of the portal theme`,
													Type:        schema.TypeString,
													Optional:    true,
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
							Description: `Resource UUID, mandatory for update`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"name": &schema.Schema{
							Description: `Resource Name`,
							Type:        schema.TypeString,
							Optional:    true,
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
							Description: `Defines all of the settings groups available for a BYOD`,
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"byod_settings": &schema.Schema{
										Description: `Configuration of BYOD Device Welcome, Registration and Success steps`,
										Type:        schema.TypeList,
										Optional:    true,
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"byod_registration_settings": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"end_point_identity_group_id": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
															"show_device_id": &schema.Schema{
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
																Description: `After an Authentication Success where should device be redirected. Allowed values:`,
																Type:        schema.TypeString,
																Optional:    true,
															},
														},
													},
												},
												"byod_welcome_settings": &schema.Schema{
													Description: `Configuration of BYOD endpoint welcome step configuration`,
													Type:        schema.TypeList,
													Optional:    true,
													MaxItems:    1,
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
																Description:  `Require BYOD devices to scroll down to the bottom of the AUP, Only valid if includeAup = true`,
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
													Description: `Used when displayLang = ALWAYSUSE`,
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
													Description: `Specifies how empty fields are handled on the Support Information Page.
Allowed values:
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

func resourceByodPortalCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning ByodPortal create")
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestByodPortalCreateByodPortal(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	vID, okID := resourceItem["id"]
	vName, _ := resourceItem["name"]
	vvID := interfaceToString(vID)
	vvName := interfaceToString(vName)
	if okID && vvID != "" {
		getResponse2, _, err := client.ByodPortal.GetByodPortalByID(vvID)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return resourceByodPortalRead(ctx, d, m)
		}
	} else {
		queryParams2 := isegosdk.GetByodPortalQueryParams{}

		response2, _, err := client.ByodPortal.GetByodPortal(&queryParams2)
		if response2 != nil && err == nil {
			items2 := getAllItemsByodPortalGetByodPortal(m, response2, &queryParams2)
			item2, err := searchByodPortalGetByodPortal(m, items2, vvName, vvID)
			if err == nil && item2 != nil {
				resourceMap := make(map[string]string)
				resourceMap["id"] = vvID
				resourceMap["name"] = vvName
				d.SetId(joinResourceID(resourceMap))
				return resourceByodPortalRead(ctx, d, m)
			}
		}
	}
	restyResp1, err := client.ByodPortal.CreateByodPortal(request1)
	if err != nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateByodPortal", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateByodPortal", err))
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
	return resourceByodPortalRead(ctx, d, m)
}

func resourceByodPortalRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning ByodPortal read for id=[%s]", d.Id())
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

		log.Printf("[DEBUG] Selected method: GetByodPortal")
		queryParams1 := isegosdk.GetByodPortalQueryParams{}
		response1, restyResp1, err := client.ByodPortal.GetByodPortal(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		items1 := getAllItemsByodPortalGetByodPortal(m, response1, &queryParams1)
		item1, err := searchByodPortalGetByodPortal(m, items1, vvName, vvID)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}
		vItem1 := flattenByodPortalGetByodPortalByIDItem(item1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetByodPortal search response",
				err))
			return diags
		}

	}
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetByodPortalByID")
		vvID := vID

		response2, restyResp2, err := client.ByodPortal.GetByodPortalByID(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			d.SetId("")
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
		return diags

	}
	return diags
}

func resourceByodPortalUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning ByodPortal update for id=[%s]", d.Id())
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
		vvName := vName
		vvID := vID

		log.Printf("[DEBUG] Selected method: GetByodPortal")
		queryParams1 := isegosdk.GetByodPortalQueryParams{}
		response1, _, err := client.ByodPortal.GetByodPortal(&queryParams1)

		if err == nil && response1 != nil {
			items1 := getAllItemsByodPortalGetByodPortal(m, response1, &queryParams1)
			item1, err := searchByodPortalGetByodPortal(m, items1, vvName, vvID)
			if err == nil && item1 != nil {
				vvID = item1.ID
			}
		}
	}
	if selectedMethod == 1 {
		vvID = vID
	}
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] ID used for update operation %s", vvID)
		request1 := expandRequestByodPortalUpdateByodPortalByID(ctx, "parameters.0", d)
		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
		response1, restyResp1, err := client.ByodPortal.UpdateByodPortalByID(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateByodPortalByID", err, restyResp1.String(),
					"Failure at UpdateByodPortalByID, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateByodPortalByID", err,
				"Failure at UpdateByodPortalByID, unexpected response", ""))
			return diags
		}
		_ = d.Set("last_updated", getUnixTimeString())
	}

	return resourceByodPortalRead(ctx, d, m)
}

func resourceByodPortalDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning ByodPortal delete for id=[%s]", d.Id())
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
		queryParams1 := isegosdk.GetByodPortalQueryParams{}

		getResp1, _, err := client.ByodPortal.GetByodPortal(&queryParams1)
		if err != nil || getResp1 == nil {
			// Assume that element it is already gone
			return diags
		}
		items1 := getAllItemsByodPortalGetByodPortal(m, getResp1, &queryParams1)
		item1, err := searchByodPortalGetByodPortal(m, items1, vName, vID)
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
		getResp, _, err := client.ByodPortal.GetByodPortalByID(vvID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	restyResp1, err := client.ByodPortal.DeleteByodPortalByID(vvID)
	if err != nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteByodPortalByID", err, restyResp1.String(),
				"Failure at DeleteByodPortalByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteByodPortalByID", err,
			"Failure at DeleteByodPortalByID, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestByodPortalCreateByodPortal(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestByodPortalCreateByodPortal {
	request := isegosdk.RequestByodPortalCreateByodPortal{}
	request.ByodPortal = expandRequestByodPortalCreateByodPortalByodPortal(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestByodPortalCreateByodPortalByodPortal(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestByodPortalCreateByodPortalByodPortal {
	request := isegosdk.RequestByodPortalCreateByodPortalByodPortal{}
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
		request.Settings = expandRequestByodPortalCreateByodPortalByodPortalSettings(ctx, key+".settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".customizations")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".customizations")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".customizations")))) {
		request.Customizations = expandRequestByodPortalCreateByodPortalByodPortalCustomizations(ctx, key+".customizations.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestByodPortalCreateByodPortalByodPortalSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestByodPortalCreateByodPortalByodPortalSettings {
	request := isegosdk.RequestByodPortalCreateByodPortalByodPortalSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".portal_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".portal_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".portal_settings")))) {
		request.PortalSettings = expandRequestByodPortalCreateByodPortalByodPortalSettingsPortalSettings(ctx, key+".portal_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".byod_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".byod_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".byod_settings")))) {
		request.ByodSettings = expandRequestByodPortalCreateByodPortalByodPortalSettingsByodSettings(ctx, key+".byod_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".support_info_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".support_info_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".support_info_settings")))) {
		request.SupportInfoSettings = expandRequestByodPortalCreateByodPortalByodPortalSettingsSupportInfoSettings(ctx, key+".support_info_settings.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestByodPortalCreateByodPortalByodPortalSettingsPortalSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestByodPortalCreateByodPortalByodPortalSettingsPortalSettings {
	request := isegosdk.RequestByodPortalCreateByodPortalByodPortalSettingsPortalSettings{}
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

func expandRequestByodPortalCreateByodPortalByodPortalSettingsByodSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestByodPortalCreateByodPortalByodPortalSettingsByodSettings {
	request := isegosdk.RequestByodPortalCreateByodPortalByodPortalSettingsByodSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".byod_welcome_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".byod_welcome_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".byod_welcome_settings")))) {
		request.ByodWelcomeSettings = expandRequestByodPortalCreateByodPortalByodPortalSettingsByodSettingsByodWelcomeSettings(ctx, key+".byod_welcome_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".byod_registration_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".byod_registration_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".byod_registration_settings")))) {
		request.ByodRegistrationSettings = expandRequestByodPortalCreateByodPortalByodPortalSettingsByodSettingsByodRegistrationSettings(ctx, key+".byod_registration_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".byod_registration_success_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".byod_registration_success_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".byod_registration_success_settings")))) {
		request.ByodRegistrationSuccessSettings = expandRequestByodPortalCreateByodPortalByodPortalSettingsByodSettingsByodRegistrationSuccessSettings(ctx, key+".byod_registration_success_settings.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestByodPortalCreateByodPortalByodPortalSettingsByodSettingsByodWelcomeSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestByodPortalCreateByodPortalByodPortalSettingsByodSettingsByodWelcomeSettings {
	request := isegosdk.RequestByodPortalCreateByodPortalByodPortalSettingsByodSettingsByodWelcomeSettings{}
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

func expandRequestByodPortalCreateByodPortalByodPortalSettingsByodSettingsByodRegistrationSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestByodPortalCreateByodPortalByodPortalSettingsByodSettingsByodRegistrationSettings {
	request := isegosdk.RequestByodPortalCreateByodPortalByodPortalSettingsByodSettingsByodRegistrationSettings{}
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

func expandRequestByodPortalCreateByodPortalByodPortalSettingsByodSettingsByodRegistrationSuccessSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestByodPortalCreateByodPortalByodPortalSettingsByodSettingsByodRegistrationSuccessSettings {
	request := isegosdk.RequestByodPortalCreateByodPortalByodPortalSettingsByodSettingsByodRegistrationSuccessSettings{}
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

func expandRequestByodPortalCreateByodPortalByodPortalSettingsSupportInfoSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestByodPortalCreateByodPortalByodPortalSettingsSupportInfoSettings {
	request := isegosdk.RequestByodPortalCreateByodPortalByodPortalSettingsSupportInfoSettings{}
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

func expandRequestByodPortalCreateByodPortalByodPortalCustomizations(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestByodPortalCreateByodPortalByodPortalCustomizations {
	request := isegosdk.RequestByodPortalCreateByodPortalByodPortalCustomizations{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".portal_theme")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".portal_theme")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".portal_theme")))) {
		request.PortalTheme = expandRequestByodPortalCreateByodPortalByodPortalCustomizationsPortalTheme(ctx, key+".portal_theme.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".portal_tweak_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".portal_tweak_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".portal_tweak_settings")))) {
		request.PortalTweakSettings = expandRequestByodPortalCreateByodPortalByodPortalCustomizationsPortalTweakSettings(ctx, key+".portal_tweak_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".language")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".language")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".language")))) {
		request.Language = expandRequestByodPortalCreateByodPortalByodPortalCustomizationsLanguage(ctx, key+".language.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".global_customizations")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".global_customizations")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".global_customizations")))) {
		request.GlobalCustomizations = expandRequestByodPortalCreateByodPortalByodPortalCustomizationsGlobalCustomizations(ctx, key+".global_customizations.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".page_customizations")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".page_customizations")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".page_customizations")))) {
		request.PageCustomizations = expandRequestByodPortalCreateByodPortalByodPortalCustomizationsPageCustomizations(ctx, key+".page_customizations.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestByodPortalCreateByodPortalByodPortalCustomizationsPortalTheme(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestByodPortalCreateByodPortalByodPortalCustomizationsPortalTheme {
	request := isegosdk.RequestByodPortalCreateByodPortalByodPortalCustomizationsPortalTheme{}
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

func expandRequestByodPortalCreateByodPortalByodPortalCustomizationsPortalTweakSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestByodPortalCreateByodPortalByodPortalCustomizationsPortalTweakSettings {
	request := isegosdk.RequestByodPortalCreateByodPortalByodPortalCustomizationsPortalTweakSettings{}
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

func expandRequestByodPortalCreateByodPortalByodPortalCustomizationsLanguage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestByodPortalCreateByodPortalByodPortalCustomizationsLanguage {
	request := isegosdk.RequestByodPortalCreateByodPortalByodPortalCustomizationsLanguage{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".view_language")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".view_language")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".view_language")))) {
		request.ViewLanguage = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestByodPortalCreateByodPortalByodPortalCustomizationsGlobalCustomizations(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestByodPortalCreateByodPortalByodPortalCustomizationsGlobalCustomizations {
	request := isegosdk.RequestByodPortalCreateByodPortalByodPortalCustomizationsGlobalCustomizations{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mobile_logo_image")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mobile_logo_image")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mobile_logo_image")))) {
		request.MobileLogoImage = expandRequestByodPortalCreateByodPortalByodPortalCustomizationsGlobalCustomizationsMobileLogoImage(ctx, key+".mobile_logo_image.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".desktop_logo_image")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".desktop_logo_image")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".desktop_logo_image")))) {
		request.DesktopLogoImage = expandRequestByodPortalCreateByodPortalByodPortalCustomizationsGlobalCustomizationsDesktopLogoImage(ctx, key+".desktop_logo_image.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".banner_image")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".banner_image")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".banner_image")))) {
		request.BannerImage = expandRequestByodPortalCreateByodPortalByodPortalCustomizationsGlobalCustomizationsBannerImage(ctx, key+".banner_image.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".background_image")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".background_image")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".background_image")))) {
		request.BackgroundImage = expandRequestByodPortalCreateByodPortalByodPortalCustomizationsGlobalCustomizationsBackgroundImage(ctx, key+".background_image.0", d)
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

func expandRequestByodPortalCreateByodPortalByodPortalCustomizationsGlobalCustomizationsMobileLogoImage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestByodPortalCreateByodPortalByodPortalCustomizationsGlobalCustomizationsMobileLogoImage {
	request := isegosdk.RequestByodPortalCreateByodPortalByodPortalCustomizationsGlobalCustomizationsMobileLogoImage{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data")))) {
		request.Data = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestByodPortalCreateByodPortalByodPortalCustomizationsGlobalCustomizationsDesktopLogoImage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestByodPortalCreateByodPortalByodPortalCustomizationsGlobalCustomizationsDesktopLogoImage {
	request := isegosdk.RequestByodPortalCreateByodPortalByodPortalCustomizationsGlobalCustomizationsDesktopLogoImage{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data")))) {
		request.Data = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestByodPortalCreateByodPortalByodPortalCustomizationsGlobalCustomizationsBannerImage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestByodPortalCreateByodPortalByodPortalCustomizationsGlobalCustomizationsBannerImage {
	request := isegosdk.RequestByodPortalCreateByodPortalByodPortalCustomizationsGlobalCustomizationsBannerImage{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data")))) {
		request.Data = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestByodPortalCreateByodPortalByodPortalCustomizationsGlobalCustomizationsBackgroundImage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestByodPortalCreateByodPortalByodPortalCustomizationsGlobalCustomizationsBackgroundImage {
	request := isegosdk.RequestByodPortalCreateByodPortalByodPortalCustomizationsGlobalCustomizationsBackgroundImage{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data")))) {
		request.Data = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestByodPortalCreateByodPortalByodPortalCustomizationsPageCustomizations(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestByodPortalCreateByodPortalByodPortalCustomizationsPageCustomizations {
	request := isegosdk.RequestByodPortalCreateByodPortalByodPortalCustomizationsPageCustomizations{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data")))) {
		request.Data = expandRequestByodPortalCreateByodPortalByodPortalCustomizationsPageCustomizationsDataArray(ctx, key+".data", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestByodPortalCreateByodPortalByodPortalCustomizationsPageCustomizationsDataArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestByodPortalCreateByodPortalByodPortalCustomizationsPageCustomizationsData {
	request := []isegosdk.RequestByodPortalCreateByodPortalByodPortalCustomizationsPageCustomizationsData{}
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
		i := expandRequestByodPortalCreateByodPortalByodPortalCustomizationsPageCustomizationsData(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestByodPortalCreateByodPortalByodPortalCustomizationsPageCustomizationsData(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestByodPortalCreateByodPortalByodPortalCustomizationsPageCustomizationsData {
	request := isegosdk.RequestByodPortalCreateByodPortalByodPortalCustomizationsPageCustomizationsData{}
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

func expandRequestByodPortalUpdateByodPortalByID(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestByodPortalUpdateByodPortalByID {
	request := isegosdk.RequestByodPortalUpdateByodPortalByID{}
	request.ByodPortal = expandRequestByodPortalUpdateByodPortalByIDByodPortal(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestByodPortalUpdateByodPortalByIDByodPortal(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestByodPortalUpdateByodPortalByIDByodPortal {
	request := isegosdk.RequestByodPortalUpdateByodPortalByIDByodPortal{}
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
		request.Settings = expandRequestByodPortalUpdateByodPortalByIDByodPortalSettings(ctx, key+".settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".customizations")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".customizations")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".customizations")))) {
		request.Customizations = expandRequestByodPortalUpdateByodPortalByIDByodPortalCustomizations(ctx, key+".customizations.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestByodPortalUpdateByodPortalByIDByodPortalSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestByodPortalUpdateByodPortalByIDByodPortalSettings {
	request := isegosdk.RequestByodPortalUpdateByodPortalByIDByodPortalSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".portal_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".portal_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".portal_settings")))) {
		request.PortalSettings = expandRequestByodPortalUpdateByodPortalByIDByodPortalSettingsPortalSettings(ctx, key+".portal_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".byod_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".byod_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".byod_settings")))) {
		request.ByodSettings = expandRequestByodPortalUpdateByodPortalByIDByodPortalSettingsByodSettings(ctx, key+".byod_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".support_info_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".support_info_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".support_info_settings")))) {
		request.SupportInfoSettings = expandRequestByodPortalUpdateByodPortalByIDByodPortalSettingsSupportInfoSettings(ctx, key+".support_info_settings.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestByodPortalUpdateByodPortalByIDByodPortalSettingsPortalSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestByodPortalUpdateByodPortalByIDByodPortalSettingsPortalSettings {
	request := isegosdk.RequestByodPortalUpdateByodPortalByIDByodPortalSettingsPortalSettings{}
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

func expandRequestByodPortalUpdateByodPortalByIDByodPortalSettingsByodSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestByodPortalUpdateByodPortalByIDByodPortalSettingsByodSettings {
	request := isegosdk.RequestByodPortalUpdateByodPortalByIDByodPortalSettingsByodSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".byod_welcome_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".byod_welcome_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".byod_welcome_settings")))) {
		request.ByodWelcomeSettings = expandRequestByodPortalUpdateByodPortalByIDByodPortalSettingsByodSettingsByodWelcomeSettings(ctx, key+".byod_welcome_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".byod_registration_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".byod_registration_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".byod_registration_settings")))) {
		request.ByodRegistrationSettings = expandRequestByodPortalUpdateByodPortalByIDByodPortalSettingsByodSettingsByodRegistrationSettings(ctx, key+".byod_registration_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".byod_registration_success_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".byod_registration_success_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".byod_registration_success_settings")))) {
		request.ByodRegistrationSuccessSettings = expandRequestByodPortalUpdateByodPortalByIDByodPortalSettingsByodSettingsByodRegistrationSuccessSettings(ctx, key+".byod_registration_success_settings.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestByodPortalUpdateByodPortalByIDByodPortalSettingsByodSettingsByodWelcomeSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestByodPortalUpdateByodPortalByIDByodPortalSettingsByodSettingsByodWelcomeSettings {
	request := isegosdk.RequestByodPortalUpdateByodPortalByIDByodPortalSettingsByodSettingsByodWelcomeSettings{}
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

func expandRequestByodPortalUpdateByodPortalByIDByodPortalSettingsByodSettingsByodRegistrationSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestByodPortalUpdateByodPortalByIDByodPortalSettingsByodSettingsByodRegistrationSettings {
	request := isegosdk.RequestByodPortalUpdateByodPortalByIDByodPortalSettingsByodSettingsByodRegistrationSettings{}
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

func expandRequestByodPortalUpdateByodPortalByIDByodPortalSettingsByodSettingsByodRegistrationSuccessSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestByodPortalUpdateByodPortalByIDByodPortalSettingsByodSettingsByodRegistrationSuccessSettings {
	request := isegosdk.RequestByodPortalUpdateByodPortalByIDByodPortalSettingsByodSettingsByodRegistrationSuccessSettings{}
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

func expandRequestByodPortalUpdateByodPortalByIDByodPortalSettingsSupportInfoSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestByodPortalUpdateByodPortalByIDByodPortalSettingsSupportInfoSettings {
	request := isegosdk.RequestByodPortalUpdateByodPortalByIDByodPortalSettingsSupportInfoSettings{}
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

func expandRequestByodPortalUpdateByodPortalByIDByodPortalCustomizations(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestByodPortalUpdateByodPortalByIDByodPortalCustomizations {
	request := isegosdk.RequestByodPortalUpdateByodPortalByIDByodPortalCustomizations{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".portal_theme")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".portal_theme")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".portal_theme")))) {
		request.PortalTheme = expandRequestByodPortalUpdateByodPortalByIDByodPortalCustomizationsPortalTheme(ctx, key+".portal_theme.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".portal_tweak_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".portal_tweak_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".portal_tweak_settings")))) {
		request.PortalTweakSettings = expandRequestByodPortalUpdateByodPortalByIDByodPortalCustomizationsPortalTweakSettings(ctx, key+".portal_tweak_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".language")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".language")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".language")))) {
		request.Language = expandRequestByodPortalUpdateByodPortalByIDByodPortalCustomizationsLanguage(ctx, key+".language.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".global_customizations")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".global_customizations")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".global_customizations")))) {
		request.GlobalCustomizations = expandRequestByodPortalUpdateByodPortalByIDByodPortalCustomizationsGlobalCustomizations(ctx, key+".global_customizations.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".page_customizations")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".page_customizations")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".page_customizations")))) {
		request.PageCustomizations = expandRequestByodPortalUpdateByodPortalByIDByodPortalCustomizationsPageCustomizations(ctx, key+".page_customizations.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestByodPortalUpdateByodPortalByIDByodPortalCustomizationsPortalTheme(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestByodPortalUpdateByodPortalByIDByodPortalCustomizationsPortalTheme {
	request := isegosdk.RequestByodPortalUpdateByodPortalByIDByodPortalCustomizationsPortalTheme{}
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

func expandRequestByodPortalUpdateByodPortalByIDByodPortalCustomizationsPortalTweakSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestByodPortalUpdateByodPortalByIDByodPortalCustomizationsPortalTweakSettings {
	request := isegosdk.RequestByodPortalUpdateByodPortalByIDByodPortalCustomizationsPortalTweakSettings{}
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

func expandRequestByodPortalUpdateByodPortalByIDByodPortalCustomizationsLanguage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestByodPortalUpdateByodPortalByIDByodPortalCustomizationsLanguage {
	request := isegosdk.RequestByodPortalUpdateByodPortalByIDByodPortalCustomizationsLanguage{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".view_language")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".view_language")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".view_language")))) {
		request.ViewLanguage = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestByodPortalUpdateByodPortalByIDByodPortalCustomizationsGlobalCustomizations(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestByodPortalUpdateByodPortalByIDByodPortalCustomizationsGlobalCustomizations {
	request := isegosdk.RequestByodPortalUpdateByodPortalByIDByodPortalCustomizationsGlobalCustomizations{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mobile_logo_image")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mobile_logo_image")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mobile_logo_image")))) {
		request.MobileLogoImage = expandRequestByodPortalUpdateByodPortalByIDByodPortalCustomizationsGlobalCustomizationsMobileLogoImage(ctx, key+".mobile_logo_image.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".desktop_logo_image")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".desktop_logo_image")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".desktop_logo_image")))) {
		request.DesktopLogoImage = expandRequestByodPortalUpdateByodPortalByIDByodPortalCustomizationsGlobalCustomizationsDesktopLogoImage(ctx, key+".desktop_logo_image.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".banner_image")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".banner_image")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".banner_image")))) {
		request.BannerImage = expandRequestByodPortalUpdateByodPortalByIDByodPortalCustomizationsGlobalCustomizationsBannerImage(ctx, key+".banner_image.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".background_image")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".background_image")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".background_image")))) {
		request.BackgroundImage = expandRequestByodPortalUpdateByodPortalByIDByodPortalCustomizationsGlobalCustomizationsBackgroundImage(ctx, key+".background_image.0", d)
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

func expandRequestByodPortalUpdateByodPortalByIDByodPortalCustomizationsGlobalCustomizationsMobileLogoImage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestByodPortalUpdateByodPortalByIDByodPortalCustomizationsGlobalCustomizationsMobileLogoImage {
	request := isegosdk.RequestByodPortalUpdateByodPortalByIDByodPortalCustomizationsGlobalCustomizationsMobileLogoImage{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data")))) {
		request.Data = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestByodPortalUpdateByodPortalByIDByodPortalCustomizationsGlobalCustomizationsDesktopLogoImage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestByodPortalUpdateByodPortalByIDByodPortalCustomizationsGlobalCustomizationsDesktopLogoImage {
	request := isegosdk.RequestByodPortalUpdateByodPortalByIDByodPortalCustomizationsGlobalCustomizationsDesktopLogoImage{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data")))) {
		request.Data = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestByodPortalUpdateByodPortalByIDByodPortalCustomizationsGlobalCustomizationsBannerImage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestByodPortalUpdateByodPortalByIDByodPortalCustomizationsGlobalCustomizationsBannerImage {
	request := isegosdk.RequestByodPortalUpdateByodPortalByIDByodPortalCustomizationsGlobalCustomizationsBannerImage{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data")))) {
		request.Data = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestByodPortalUpdateByodPortalByIDByodPortalCustomizationsGlobalCustomizationsBackgroundImage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestByodPortalUpdateByodPortalByIDByodPortalCustomizationsGlobalCustomizationsBackgroundImage {
	request := isegosdk.RequestByodPortalUpdateByodPortalByIDByodPortalCustomizationsGlobalCustomizationsBackgroundImage{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data")))) {
		request.Data = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestByodPortalUpdateByodPortalByIDByodPortalCustomizationsPageCustomizations(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestByodPortalUpdateByodPortalByIDByodPortalCustomizationsPageCustomizations {
	request := isegosdk.RequestByodPortalUpdateByodPortalByIDByodPortalCustomizationsPageCustomizations{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data")))) {
		request.Data = expandRequestByodPortalUpdateByodPortalByIDByodPortalCustomizationsPageCustomizationsDataArray(ctx, key+".data", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestByodPortalUpdateByodPortalByIDByodPortalCustomizationsPageCustomizationsDataArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestByodPortalUpdateByodPortalByIDByodPortalCustomizationsPageCustomizationsData {
	request := []isegosdk.RequestByodPortalUpdateByodPortalByIDByodPortalCustomizationsPageCustomizationsData{}
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
		i := expandRequestByodPortalUpdateByodPortalByIDByodPortalCustomizationsPageCustomizationsData(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestByodPortalUpdateByodPortalByIDByodPortalCustomizationsPageCustomizationsData(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestByodPortalUpdateByodPortalByIDByodPortalCustomizationsPageCustomizationsData {
	request := isegosdk.RequestByodPortalUpdateByodPortalByIDByodPortalCustomizationsPageCustomizationsData{}
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

func getAllItemsByodPortalGetByodPortal(m interface{}, response *isegosdk.ResponseByodPortalGetByodPortal, queryParams *isegosdk.GetByodPortalQueryParams) []isegosdk.ResponseByodPortalGetByodPortalSearchResultResources {
	client := m.(*isegosdk.Client)
	var respItems []isegosdk.ResponseByodPortalGetByodPortalSearchResultResources
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
			response, _, err = client.ByodPortal.GetByodPortal(queryParams)
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

func searchByodPortalGetByodPortal(m interface{}, items []isegosdk.ResponseByodPortalGetByodPortalSearchResultResources, name string, id string) (*isegosdk.ResponseByodPortalGetByodPortalByIDByodPortal, error) {
	client := m.(*isegosdk.Client)
	var err error
	var foundItem *isegosdk.ResponseByodPortalGetByodPortalByIDByodPortal
	for _, item := range items {
		if id != "" && item.ID == id {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseByodPortalGetByodPortalByID
			getItem, _, err = client.ByodPortal.GetByodPortalByID(id)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetByodPortalByID")
			}
			foundItem = getItem.ByodPortal
			return foundItem, err
		} else if name != "" && item.Name == name {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseByodPortalGetByodPortalByID
			getItem, _, err = client.ByodPortal.GetByodPortalByID(item.ID)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetByodPortalByID")
			}
			foundItem = getItem.ByodPortal
			return foundItem, err
		}
	}
	return foundItem, err
}
