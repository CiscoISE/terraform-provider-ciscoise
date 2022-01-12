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

func resourceSelfRegisteredPortal() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on SelfRegisteredPortal.

- This resource allows the client to update a self registered portal by ID.

- This resource deletes a self registered portal by ID.

- This resource creates a self registered portal.
`,

		CreateContext: resourceSelfRegisteredPortalCreate,
		ReadContext:   resourceSelfRegisteredPortalRead,
		UpdateContext: resourceSelfRegisteredPortalUpdate,
		DeleteContext: resourceSelfRegisteredPortalDelete,
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
													Description: `How the AUP should be displayed, either on page or as a link.
Only valid if includeAup = true.
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
												"require_aup_scrolling": &schema.Schema{
													Description: `Require the portal user to scroll to the end of the AUP.
Only valid if requireAupAcceptance = true`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"require_scrolling": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"skip_aup_for_employees": &schema.Schema{
													Description: `Only valid if requireAupAcceptance = trueG`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"use_diff_aup_for_employees": &schema.Schema{
													Description: `Only valid if requireAupAcceptance = trueG`,
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
									"byod_settings": &schema.Schema{
										Description: `Configuration of BYOD Device Welcome, Registration and Success steps`,
										Type:        schema.TypeList,
										Computed:    true,
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
													Description: `Configuration of BYOD endpoint Registration Success step configuration`,
													Type:        schema.TypeList,
													Computed:    true,
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
																Description: `Require BYOD devices to scroll down to the bottom of the AUP, 
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
												"allow_guest_to_use_social_accounts": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"allow_show_guest_form": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"alternate_guest_portal": &schema.Schema{
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
													Description: `Require the portal user to enter an access code`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"require_aup_acceptance": &schema.Schema{
													Description: `Require the portal user to accept the AUP. Only valid if includeAup = true`,
													Type:        schema.TypeString,
													Computed:    true,
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
									"self_reg_page_settings": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"account_validity_duration": &schema.Schema{
													Description: `Self-registered guest account is valid for this many account_validity_time_units`,
													Type:        schema.TypeInt,
													Computed:    true,
												},
												"account_validity_time_units": &schema.Schema{
													Description: `Time units for account_validity_duration.
Allowed Values:
- DAYS,
- HOURS,
- MINUTES`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"allow_grace_access": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"approval_email_addresses": &schema.Schema{
													Description: `Only valid if requireGuestApproval = true and sendApprovalRequestTo = SELECTEDEMAILADDRESSES`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"approve_deny_links_time_units": &schema.Schema{
													Description: `This attribute, along with approveDenyLinksValidFor, specifies how long the link can be used.
Only valid if requireGuestApproval = true.
Allowed Values:
- DAYS,
- HOURS,
- MINUTES`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"approve_deny_links_valid_for": &schema.Schema{
													Description: `This attribute, along with approveDenyLinksTimeUnits, specifies how long the link can be used.
Only valid if requireGuestApproval = true`,
													Type:     schema.TypeInt,
													Computed: true,
												},
												"assign_guests_to_guest_type": &schema.Schema{
													Description: `Guests are assigned to this guest type`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"aup_display": &schema.Schema{
													Description: `How the AUP should be displayed, either on page or as a link.
Only valid if includeAup = true.
Allowed values:
- ONPAGE,
- ASLINK`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"authenticate_sponsors_using_portal_list": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"auto_login_self_wait": &schema.Schema{
													Description: `Allow guests to login automatically from self-registration after sponsor's approval.
No need to provide the credentials by guest to login`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"auto_login_time_period": &schema.Schema{
													Description: `Waiting period for auto login until sponsor's approval.
If time exceeds, guest has to login manually by providing the credentials.
Default value is 5 minutes`,
													Type:     schema.TypeInt,
													Computed: true,
												},
												"credential_notification_using_email": &schema.Schema{
													Description: `If true, send credential notification upon approval using email.
Only valid if requireGuestApproval = true`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"credential_notification_using_sms": &schema.Schema{
													Description: `If true, send credential notification upon approval using SMS.
Only valid if requireGuestApproval = true`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"enable_guest_email_blacklist": &schema.Schema{
													Description: `Disallow guests with an e-mail address from selected domains`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"enable_guest_email_whitelist": &schema.Schema{
													Description: `Allow guests with an e-mail address from selected domains`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"field_company": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"include": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"require": &schema.Schema{
																Description: `Only applicable if include = true`,
																Type:        schema.TypeString,
																Computed:    true,
															},
														},
													},
												},
												"field_email_addr": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"include": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"require": &schema.Schema{
																Description: `Only applicable if include = true`,
																Type:        schema.TypeString,
																Computed:    true,
															},
														},
													},
												},
												"field_first_name": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"include": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"require": &schema.Schema{
																Description: `Only applicable if include = true`,
																Type:        schema.TypeString,
																Computed:    true,
															},
														},
													},
												},
												"field_last_name": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"include": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"require": &schema.Schema{
																Description: `Only applicable if include = true`,
																Type:        schema.TypeString,
																Computed:    true,
															},
														},
													},
												},
												"field_location": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"include": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"require": &schema.Schema{
																Description: `Only applicable if include = true`,
																Type:        schema.TypeString,
																Computed:    true,
															},
														},
													},
												},
												"field_person_being_visited": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"include": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"require": &schema.Schema{
																Description: `Only applicable if include = true`,
																Type:        schema.TypeString,
																Computed:    true,
															},
														},
													},
												},
												"field_phone_no": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"include": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"require": &schema.Schema{
																Description: `Only applicable if include = true`,
																Type:        schema.TypeString,
																Computed:    true,
															},
														},
													},
												},
												"field_reason_for_visit": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"include": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"require": &schema.Schema{
																Description: `Only applicable if include = true`,
																Type:        schema.TypeString,
																Computed:    true,
															},
														},
													},
												},
												"field_sms_provider": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"include": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"require": &schema.Schema{
																Description: `Only applicable if include = true`,
																Type:        schema.TypeString,
																Computed:    true,
															},
														},
													},
												},
												"field_user_name": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"include": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"require": &schema.Schema{
																Description: `Only applicable if include = true`,
																Type:        schema.TypeString,
																Computed:    true,
															},
														},
													},
												},
												"grace_access_expire_interval": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"grace_access_send_account_expiration": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"guest_email_blacklist_domains": &schema.Schema{
													Description: `Disallow guests with an e-mail address from selected domains`,
													Type:        schema.TypeList,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"guest_email_whitelist_domains": &schema.Schema{
													Description: `Self-registered guests whose e-mail address is in one of these domains will be allowed.
Only valid if enableGuestEmailWhitelist = true`,
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"include_aup": &schema.Schema{
													Description: `Include an Acceptable Use Policy (AUP) that should be displayed during login`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"post_registration_redirect": &schema.Schema{
													Description: `After the registration submission direct the guest user to one of the following pages.
Only valid if requireGuestApproval = true.
Allowed Values:
- SELFREGISTRATIONSUCCESS,
- LOGINPAGEWITHINSTRUCTIONS
- URL`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"post_registration_redirect_url": &schema.Schema{
													Description: `URL where guest user is redirected after registration.
Only valid if requireGuestApproval = true and postRegistrationRedirect = URL`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"registration_code": &schema.Schema{
													Description: `The registration code that the guest user must enter`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"require_approver_to_authenticate": &schema.Schema{
													Description: `When self-registered guests require approval, an approval request is e-mailed to one or more sponsor users.
If the Cisco ISE Administrator chooses to include an approval link in the e-mail,
a sponsor user who clicks the link will be required to enter their username and password if this attribute is true.
Only valid if requireGuestApproval = true`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"require_aup_acceptance": &schema.Schema{
													Description: `Require the portal user to accept the AUP. Only valid if includeAup = true`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"require_guest_approval": &schema.Schema{
													Description: `Require self-registered guests to be approved if true`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"require_registration_code": &schema.Schema{
													Description: `Self-registered guests are required to enter a registration code`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"selectable_locations": &schema.Schema{
													Description: `Guests can choose from these locations to set their time zone`,
													Type:        schema.TypeList,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"selectable_sms_providers": &schema.Schema{
													Description: `This attribute is an array of SMS provider names`,
													Type:        schema.TypeList,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"send_approval_request_to": &schema.Schema{
													Description: `Specifies where approval requests are sent.
Only valid if requireGuestApproval = true.
Allowed Values:
- SELECTEDEMAILADDRESSES,
- PERSONBEINGVISITED`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"sponsor_portal_list": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
									"self_reg_success_settings": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"allow_guest_login_from_selfreg_success_page": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"allow_guest_send_self_using_email": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"allow_guest_send_self_using_print": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"allow_guest_send_self_using_sms": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"aup_on_page": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"include_aup": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"include_company": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"include_email_addr": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"include_first_name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"include_last_name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"include_location": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"include_password": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"include_person_being_visited": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"include_phone_no": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"include_reason_for_visit": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"include_sms_provider": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"include_user_name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"require_aup_acceptance": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"require_aup_scrolling": &schema.Schema{
													Type:     schema.TypeString,
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
				Optional: true,
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
													Description: `How the AUP should be displayed, either on page or as a link.
Only valid if includeAup = true.
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
												"require_aup_scrolling": &schema.Schema{
													Description: `Require the portal user to scroll to the end of the AUP.
Only valid if requireAupAcceptance = true`,
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
													Description:  `Only valid if requireAupAcceptance = trueG`,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"use_diff_aup_for_employees": &schema.Schema{
													Description:  `Only valid if requireAupAcceptance = trueG`,
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
													Type:     schema.TypeString,
													Optional: true,
												},
												"success_redirect": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"byod_settings": &schema.Schema{
										Description: `Configuration of BYOD Device Welcome, Registration and Success steps`,
										Type:        schema.TypeList,
										Optional:    true,
										MaxItems:    1,
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
													Description: `Configuration of BYOD endpoint Registration Success step configuration`,
													Type:        schema.TypeList,
													Optional:    true,
													MaxItems:    1,
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
																Description: `Require BYOD devices to scroll down to the bottom of the AUP, 
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
												"allow_guest_to_use_social_accounts": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"allow_show_guest_form": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"alternate_guest_portal": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
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
													Description:  `Require the portal user to enter an access code`,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"require_aup_acceptance": &schema.Schema{
													Description:  `Require the portal user to accept the AUP. Only valid if includeAup = true`,
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
									"self_reg_page_settings": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"account_validity_duration": &schema.Schema{
													Description: `Self-registered guest account is valid for this many account_validity_time_units`,
													Type:        schema.TypeInt,
													Optional:    true,
												},
												"account_validity_time_units": &schema.Schema{
													Description: `Time units for account_validity_duration.
Allowed Values:
- DAYS,
- HOURS,
- MINUTES`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"allow_grace_access": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"approval_email_addresses": &schema.Schema{
													Description: `Only valid if requireGuestApproval = true and sendApprovalRequestTo = SELECTEDEMAILADDRESSES`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"approve_deny_links_time_units": &schema.Schema{
													Description: `This attribute, along with approveDenyLinksValidFor, specifies how long the link can be used.
Only valid if requireGuestApproval = true.
Allowed Values:
- DAYS,
- HOURS,
- MINUTES`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"approve_deny_links_valid_for": &schema.Schema{
													Description: `This attribute, along with approveDenyLinksTimeUnits, specifies how long the link can be used.
Only valid if requireGuestApproval = true`,
													Type:     schema.TypeInt,
													Optional: true,
												},
												"assign_guests_to_guest_type": &schema.Schema{
													Description: `Guests are assigned to this guest type`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"aup_display": &schema.Schema{
													Description: `How the AUP should be displayed, either on page or as a link.
Only valid if includeAup = true.
Allowed values:
- ONPAGE,
- ASLINK`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"authenticate_sponsors_using_portal_list": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"auto_login_self_wait": &schema.Schema{
													Description: `Allow guests to login automatically from self-registration after sponsor's approval.
No need to provide the credentials by guest to login`,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"auto_login_time_period": &schema.Schema{
													Description: `Waiting period for auto login until sponsor's approval.
If time exceeds, guest has to login manually by providing the credentials.
Default value is 5 minutes`,
													Type:     schema.TypeInt,
													Optional: true,
												},
												"credential_notification_using_email": &schema.Schema{
													Description: `If true, send credential notification upon approval using email.
Only valid if requireGuestApproval = true`,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"credential_notification_using_sms": &schema.Schema{
													Description: `If true, send credential notification upon approval using SMS.
Only valid if requireGuestApproval = true`,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"enable_guest_email_blacklist": &schema.Schema{
													Description:  `Disallow guests with an e-mail address from selected domains`,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"enable_guest_email_whitelist": &schema.Schema{
													Description:  `Allow guests with an e-mail address from selected domains`,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"field_company": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"include": &schema.Schema{
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
															},
															"require": &schema.Schema{
																Description:  `Only applicable if include = true`,
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
															},
														},
													},
												},
												"field_email_addr": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"include": &schema.Schema{
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
															},
															"require": &schema.Schema{
																Description:  `Only applicable if include = true`,
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
															},
														},
													},
												},
												"field_first_name": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"include": &schema.Schema{
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
															},
															"require": &schema.Schema{
																Description:  `Only applicable if include = true`,
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
															},
														},
													},
												},
												"field_last_name": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"include": &schema.Schema{
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
															},
															"require": &schema.Schema{
																Description:  `Only applicable if include = true`,
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
															},
														},
													},
												},
												"field_location": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"include": &schema.Schema{
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
															},
															"require": &schema.Schema{
																Description:  `Only applicable if include = true`,
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
															},
														},
													},
												},
												"field_person_being_visited": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"include": &schema.Schema{
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
															},
															"require": &schema.Schema{
																Description:  `Only applicable if include = true`,
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
															},
														},
													},
												},
												"field_phone_no": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"include": &schema.Schema{
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
															},
															"require": &schema.Schema{
																Description:  `Only applicable if include = true`,
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
															},
														},
													},
												},
												"field_reason_for_visit": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"include": &schema.Schema{
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
															},
															"require": &schema.Schema{
																Description:  `Only applicable if include = true`,
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
															},
														},
													},
												},
												"field_sms_provider": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"include": &schema.Schema{
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
															},
															"require": &schema.Schema{
																Description:  `Only applicable if include = true`,
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
															},
														},
													},
												},
												"field_user_name": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"include": &schema.Schema{
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
															},
															"require": &schema.Schema{
																Description:  `Only applicable if include = true`,
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
															},
														},
													},
												},
												"grace_access_expire_interval": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
												},
												"grace_access_send_account_expiration": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"guest_email_blacklist_domains": &schema.Schema{
													Description: `Disallow guests with an e-mail address from selected domains`,
													Type:        schema.TypeList,
													Optional:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"guest_email_whitelist_domains": &schema.Schema{
													Description: `Self-registered guests whose e-mail address is in one of these domains will be allowed.
Only valid if enableGuestEmailWhitelist = true`,
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"include_aup": &schema.Schema{
													Description:  `Include an Acceptable Use Policy (AUP) that should be displayed during login`,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"post_registration_redirect": &schema.Schema{
													Description: `After the registration submission direct the guest user to one of the following pages.
Only valid if requireGuestApproval = true.
Allowed Values:
- SELFREGISTRATIONSUCCESS,
- LOGINPAGEWITHINSTRUCTIONS
- URL`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"post_registration_redirect_url": &schema.Schema{
													Description: `URL where guest user is redirected after registration.
Only valid if requireGuestApproval = true and postRegistrationRedirect = URL`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"registration_code": &schema.Schema{
													Description: `The registration code that the guest user must enter`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"require_approver_to_authenticate": &schema.Schema{
													Description: `When self-registered guests require approval, an approval request is e-mailed to one or more sponsor users.
If the Cisco ISE Administrator chooses to include an approval link in the e-mail,
a sponsor user who clicks the link will be required to enter their username and password if this attribute is true.
Only valid if requireGuestApproval = true`,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"require_aup_acceptance": &schema.Schema{
													Description:  `Require the portal user to accept the AUP. Only valid if includeAup = true`,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"require_guest_approval": &schema.Schema{
													Description:  `Require self-registered guests to be approved if true`,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"require_registration_code": &schema.Schema{
													Description:  `Self-registered guests are required to enter a registration code`,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"selectable_locations": &schema.Schema{
													Description: `Guests can choose from these locations to set their time zone`,
													Type:        schema.TypeList,
													Optional:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"selectable_sms_providers": &schema.Schema{
													Description: `This attribute is an array of SMS provider names`,
													Type:        schema.TypeList,
													Optional:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"send_approval_request_to": &schema.Schema{
													Description: `Specifies where approval requests are sent.
Only valid if requireGuestApproval = true.
Allowed Values:
- SELECTEDEMAILADDRESSES,
- PERSONBEINGVISITED`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"sponsor_portal_list": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
									"self_reg_success_settings": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"allow_guest_login_from_selfreg_success_page": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"allow_guest_send_self_using_email": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"allow_guest_send_self_using_print": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"allow_guest_send_self_using_sms": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"aup_on_page": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"include_aup": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"include_company": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"include_email_addr": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"include_first_name": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"include_last_name": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"include_location": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"include_password": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"include_person_being_visited": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"include_phone_no": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"include_reason_for_visit": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"include_sms_provider": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"include_user_name": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"require_aup_acceptance": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"require_aup_scrolling": &schema.Schema{
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

func resourceSelfRegisteredPortalCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning SelfRegisteredPortal create")
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestSelfRegisteredPortalCreateSelfRegisteredPortal(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vName, _ := resourceItem["name"]
	vvName := interfaceToString(vName)
	if okID && vvID != "" {
		getResponse2, _, err := client.SelfRegisteredPortal.GetSelfRegisteredPortalByID(vvID)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return resourceSelfRegisteredPortalRead(ctx, d, m)
		}
	} else {
		queryParams2 := isegosdk.GetSelfRegisteredPortalsQueryParams{}

		response2, _, err := client.SelfRegisteredPortal.GetSelfRegisteredPortals(&queryParams2)
		if response2 != nil && err == nil {
			items2 := getAllItemsSelfRegisteredPortalGetSelfRegisteredPortals(m, response2, &queryParams2)
			item2, err := searchSelfRegisteredPortalGetSelfRegisteredPortals(m, items2, vvName, vvID)
			if err == nil && item2 != nil {
				resourceMap := make(map[string]string)
				resourceMap["id"] = vvID
				resourceMap["name"] = vvName
				d.SetId(joinResourceID(resourceMap))
				return resourceSelfRegisteredPortalRead(ctx, d, m)
			}
		}
	}
	restyResp1, err := client.SelfRegisteredPortal.CreateSelfRegisteredPortal(request1)
	if err != nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateSelfRegisteredPortal", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateSelfRegisteredPortal", err))
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
	return resourceSelfRegisteredPortalRead(ctx, d, m)
}

func resourceSelfRegisteredPortalRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning SelfRegisteredPortal read for id=[%s]", d.Id())
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
		log.Printf("[DEBUG] Selected method: GetSelfRegisteredPortals")
		queryParams1 := isegosdk.GetSelfRegisteredPortalsQueryParams{}

		response1, restyResp1, err := client.SelfRegisteredPortal.GetSelfRegisteredPortals(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		items1 := getAllItemsSelfRegisteredPortalGetSelfRegisteredPortals(m, response1, &queryParams1)
		item1, err := searchSelfRegisteredPortalGetSelfRegisteredPortals(m, items1, vvName, vvID)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}
		vItem1 := flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItem(item1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSelfRegisteredPortals search response",
				err))
			return diags
		}

	}
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetSelfRegisteredPortalByID")
		vvID := vID

		response2, restyResp2, err := client.SelfRegisteredPortal.GetSelfRegisteredPortalByID(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItem(response2.SelfRegPortal)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSelfRegisteredPortalByID response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceSelfRegisteredPortalUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning SelfRegisteredPortal update for id=[%s]", d.Id())
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
		queryParams1 := isegosdk.GetSelfRegisteredPortalsQueryParams{}

		getResp1, _, err := client.SelfRegisteredPortal.GetSelfRegisteredPortals(&queryParams1)
		if err == nil && getResp1 != nil {
			items1 := getAllItemsSelfRegisteredPortalGetSelfRegisteredPortals(m, getResp1, &queryParams1)
			item1, err := searchSelfRegisteredPortalGetSelfRegisteredPortals(m, items1, vName, vID)
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
		request1 := expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByID(ctx, "parameters.0", d)
		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
		response1, restyResp1, err := client.SelfRegisteredPortal.UpdateSelfRegisteredPortalByID(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateSelfRegisteredPortalByID", err, restyResp1.String(),
					"Failure at UpdateSelfRegisteredPortalByID, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateSelfRegisteredPortalByID", err,
				"Failure at UpdateSelfRegisteredPortalByID, unexpected response", ""))
			return diags
		}
	}

	return resourceSelfRegisteredPortalRead(ctx, d, m)
}

func resourceSelfRegisteredPortalDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning SelfRegisteredPortal delete for id=[%s]", d.Id())
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
		queryParams1 := isegosdk.GetSelfRegisteredPortalsQueryParams{}

		getResp1, _, err := client.SelfRegisteredPortal.GetSelfRegisteredPortals(&queryParams1)
		if err != nil || getResp1 == nil {
			// Assume that element it is already gone
			return diags
		}
		items1 := getAllItemsSelfRegisteredPortalGetSelfRegisteredPortals(m, getResp1, &queryParams1)
		item1, err := searchSelfRegisteredPortalGetSelfRegisteredPortals(m, items1, vName, vID)
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
		getResp, _, err := client.SelfRegisteredPortal.GetSelfRegisteredPortalByID(vvID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	restyResp1, err := client.SelfRegisteredPortal.DeleteSelfRegisteredPortalByID(vvID)
	if err != nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteSelfRegisteredPortalByID", err, restyResp1.String(),
				"Failure at DeleteSelfRegisteredPortalByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteSelfRegisteredPortalByID", err,
			"Failure at DeleteSelfRegisteredPortalByID, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestSelfRegisteredPortalCreateSelfRegisteredPortal(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortal {
	request := isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortal{}
	request.SelfRegPortal = expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortal(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortal(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortal {
	request := isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortal{}
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
		request.Settings = expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettings(ctx, key+".settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".customizations")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".customizations")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".customizations")))) {
		request.Customizations = expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizations(ctx, key+".customizations.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettings {
	request := isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".portal_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".portal_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".portal_settings")))) {
		request.PortalSettings = expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsPortalSettings(ctx, key+".portal_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".login_page_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".login_page_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".login_page_settings")))) {
		request.LoginPageSettings = expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsLoginPageSettings(ctx, key+".login_page_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".self_reg_page_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".self_reg_page_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".self_reg_page_settings")))) {
		request.SelfRegPageSettings = expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegPageSettings(ctx, key+".self_reg_page_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".self_reg_success_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".self_reg_success_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".self_reg_success_settings")))) {
		request.SelfRegSuccessSettings = expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegSuccessSettings(ctx, key+".self_reg_success_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".aup_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".aup_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".aup_settings")))) {
		request.AupSettings = expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsAupSettings(ctx, key+".aup_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".guest_change_password_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".guest_change_password_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".guest_change_password_settings")))) {
		request.GuestChangePasswordSettings = expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsGuestChangePasswordSettings(ctx, key+".guest_change_password_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".guest_device_registration_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".guest_device_registration_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".guest_device_registration_settings")))) {
		request.GuestDeviceRegistrationSettings = expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsGuestDeviceRegistrationSettings(ctx, key+".guest_device_registration_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".byod_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".byod_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".byod_settings")))) {
		request.ByodSettings = expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsByodSettings(ctx, key+".byod_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".post_login_banner_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".post_login_banner_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".post_login_banner_settings")))) {
		request.PostLoginBannerSettings = expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsPostLoginBannerSettings(ctx, key+".post_login_banner_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".post_access_banner_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".post_access_banner_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".post_access_banner_settings")))) {
		request.PostAccessBannerSettings = expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsPostAccessBannerSettings(ctx, key+".post_access_banner_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".auth_success_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".auth_success_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".auth_success_settings")))) {
		request.AuthSuccessSettings = expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsAuthSuccessSettings(ctx, key+".auth_success_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".support_info_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".support_info_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".support_info_settings")))) {
		request.SupportInfoSettings = expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSupportInfoSettings(ctx, key+".support_info_settings.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsPortalSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsPortalSettings {
	request := isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsPortalSettings{}
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

func expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsLoginPageSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsLoginPageSettings {
	request := isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsLoginPageSettings{}
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
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".alternate_guest_portal")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".alternate_guest_portal")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".alternate_guest_portal")))) {
		request.AlternateGuestPortal = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_guest_to_use_social_accounts")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_guest_to_use_social_accounts")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_guest_to_use_social_accounts")))) {
		request.AllowGuestToUseSocialAccounts = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_show_guest_form")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_show_guest_form")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_show_guest_form")))) {
		request.AllowShowGuestForm = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".social_configs")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".social_configs")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".social_configs")))) {
		request.SocialConfigs = expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsLoginPageSettingsSocialConfigsArray(ctx, key+".social_configs", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsLoginPageSettingsSocialConfigsArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsLoginPageSettingsSocialConfigs {
	request := []isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsLoginPageSettingsSocialConfigs{}
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
		i := expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsLoginPageSettingsSocialConfigs(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsLoginPageSettingsSocialConfigs(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsLoginPageSettingsSocialConfigs {
	request := isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsLoginPageSettingsSocialConfigs{}
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

func expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegPageSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegPageSettings {
	request := isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegPageSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".assign_guests_to_guest_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".assign_guests_to_guest_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".assign_guests_to_guest_type")))) {
		request.AssignGuestsToGuestType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".account_validity_duration")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".account_validity_duration")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".account_validity_duration")))) {
		request.AccountValidityDuration = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".account_validity_time_units")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".account_validity_time_units")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".account_validity_time_units")))) {
		request.AccountValidityTimeUnits = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".require_registration_code")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".require_registration_code")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".require_registration_code")))) {
		request.RequireRegistrationCode = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".registration_code")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".registration_code")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".registration_code")))) {
		request.RegistrationCode = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".field_user_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".field_user_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".field_user_name")))) {
		request.FieldUserName = expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegPageSettingsFieldUserName(ctx, key+".field_user_name.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".field_first_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".field_first_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".field_first_name")))) {
		request.FieldFirstName = expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegPageSettingsFieldFirstName(ctx, key+".field_first_name.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".field_last_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".field_last_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".field_last_name")))) {
		request.FieldLastName = expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegPageSettingsFieldLastName(ctx, key+".field_last_name.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".field_email_addr")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".field_email_addr")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".field_email_addr")))) {
		request.FieldEmailAddr = expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegPageSettingsFieldEmailAddr(ctx, key+".field_email_addr.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".field_phone_no")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".field_phone_no")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".field_phone_no")))) {
		request.FieldPhoneNo = expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegPageSettingsFieldPhoneNo(ctx, key+".field_phone_no.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".field_company")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".field_company")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".field_company")))) {
		request.FieldCompany = expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegPageSettingsFieldCompany(ctx, key+".field_company.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".field_location")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".field_location")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".field_location")))) {
		request.FieldLocation = expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegPageSettingsFieldLocation(ctx, key+".field_location.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selectable_locations")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selectable_locations")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selectable_locations")))) {
		request.SelectableLocations = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".field_sms_provider")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".field_sms_provider")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".field_sms_provider")))) {
		request.FieldSmsProvider = expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegPageSettingsFieldSmsProvider(ctx, key+".field_sms_provider.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selectable_sms_providers")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selectable_sms_providers")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selectable_sms_providers")))) {
		request.SelectableSmsProviders = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".field_person_being_visited")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".field_person_being_visited")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".field_person_being_visited")))) {
		request.FieldPersonBeingVisited = expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegPageSettingsFieldPersonBeingVisited(ctx, key+".field_person_being_visited.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".field_reason_for_visit")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".field_reason_for_visit")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".field_reason_for_visit")))) {
		request.FieldReasonForVisit = expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegPageSettingsFieldReasonForVisit(ctx, key+".field_reason_for_visit.0", d)
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
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_guest_email_whitelist")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_guest_email_whitelist")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_guest_email_whitelist")))) {
		request.EnableGuestEmailWhitelist = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".guest_email_whitelist_domains")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".guest_email_whitelist_domains")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".guest_email_whitelist_domains")))) {
		request.GuestEmailWhitelistDomains = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_guest_email_blacklist")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_guest_email_blacklist")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_guest_email_blacklist")))) {
		request.EnableGuestEmailBlacklist = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".guest_email_blacklist_domains")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".guest_email_blacklist_domains")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".guest_email_blacklist_domains")))) {
		request.GuestEmailBlacklistDomains = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".require_guest_approval")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".require_guest_approval")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".require_guest_approval")))) {
		request.RequireGuestApproval = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".auto_login_self_wait")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".auto_login_self_wait")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".auto_login_self_wait")))) {
		request.AutoLoginSelfWait = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".auto_login_time_period")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".auto_login_time_period")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".auto_login_time_period")))) {
		request.AutoLoginTimePeriod = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_grace_access")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_grace_access")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_grace_access")))) {
		request.AllowGraceAccess = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".grace_access_expire_interval")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".grace_access_expire_interval")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".grace_access_expire_interval")))) {
		request.GraceAccessExpireInterval = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".grace_access_send_account_expiration")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".grace_access_send_account_expiration")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".grace_access_send_account_expiration")))) {
		request.GraceAccessSendAccountExpiration = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".send_approval_request_to")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".send_approval_request_to")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".send_approval_request_to")))) {
		request.SendApprovalRequestTo = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".approval_email_addresses")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".approval_email_addresses")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".approval_email_addresses")))) {
		request.ApprovalEmailAddresses = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".post_registration_redirect")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".post_registration_redirect")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".post_registration_redirect")))) {
		request.PostRegistrationRedirect = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".post_registration_redirect_url")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".post_registration_redirect_url")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".post_registration_redirect_url")))) {
		request.PostRegistrationRedirectURL = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".credential_notification_using_email")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".credential_notification_using_email")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".credential_notification_using_email")))) {
		request.CredentialNotificationUsingEmail = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".credential_notification_using_sms")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".credential_notification_using_sms")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".credential_notification_using_sms")))) {
		request.CredentialNotificationUsingSms = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".approve_deny_links_valid_for")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".approve_deny_links_valid_for")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".approve_deny_links_valid_for")))) {
		request.ApproveDenyLinksValidFor = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".approve_deny_links_time_units")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".approve_deny_links_time_units")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".approve_deny_links_time_units")))) {
		request.ApproveDenyLinksTimeUnits = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".require_approver_to_authenticate")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".require_approver_to_authenticate")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".require_approver_to_authenticate")))) {
		request.RequireApproverToAuthenticate = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".authenticate_sponsors_using_portal_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".authenticate_sponsors_using_portal_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".authenticate_sponsors_using_portal_list")))) {
		request.AuthenticateSponsorsUsingPortalList = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sponsor_portal_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sponsor_portal_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sponsor_portal_list")))) {
		vSponsorPortalList := v.([]interface{})
		request.SponsorPortalList = &vSponsorPortalList
	}
	return &request
}

func expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegPageSettingsFieldUserName(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegPageSettingsFieldUserName {
	request := isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegPageSettingsFieldUserName{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include")))) {
		request.Include = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".require")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".require")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".require")))) {
		request.Require = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegPageSettingsFieldFirstName(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegPageSettingsFieldFirstName {
	request := isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegPageSettingsFieldFirstName{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include")))) {
		request.Include = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".require")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".require")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".require")))) {
		request.Require = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegPageSettingsFieldLastName(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegPageSettingsFieldLastName {
	request := isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegPageSettingsFieldLastName{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include")))) {
		request.Include = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".require")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".require")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".require")))) {
		request.Require = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegPageSettingsFieldEmailAddr(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegPageSettingsFieldEmailAddr {
	request := isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegPageSettingsFieldEmailAddr{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include")))) {
		request.Include = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".require")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".require")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".require")))) {
		request.Require = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegPageSettingsFieldPhoneNo(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegPageSettingsFieldPhoneNo {
	request := isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegPageSettingsFieldPhoneNo{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include")))) {
		request.Include = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".require")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".require")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".require")))) {
		request.Require = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegPageSettingsFieldCompany(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegPageSettingsFieldCompany {
	request := isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegPageSettingsFieldCompany{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include")))) {
		request.Include = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".require")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".require")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".require")))) {
		request.Require = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegPageSettingsFieldLocation(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegPageSettingsFieldLocation {
	request := isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegPageSettingsFieldLocation{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include")))) {
		request.Include = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".require")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".require")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".require")))) {
		request.Require = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegPageSettingsFieldSmsProvider(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegPageSettingsFieldSmsProvider {
	request := isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegPageSettingsFieldSmsProvider{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include")))) {
		request.Include = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".require")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".require")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".require")))) {
		request.Require = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegPageSettingsFieldPersonBeingVisited(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegPageSettingsFieldPersonBeingVisited {
	request := isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegPageSettingsFieldPersonBeingVisited{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include")))) {
		request.Include = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".require")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".require")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".require")))) {
		request.Require = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegPageSettingsFieldReasonForVisit(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegPageSettingsFieldReasonForVisit {
	request := isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegPageSettingsFieldReasonForVisit{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include")))) {
		request.Include = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".require")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".require")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".require")))) {
		request.Require = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegSuccessSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegSuccessSettings {
	request := isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSelfRegSuccessSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include_user_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include_user_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include_user_name")))) {
		request.IncludeUserName = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include_password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include_password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include_password")))) {
		request.IncludePassword = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include_first_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include_first_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include_first_name")))) {
		request.IncludeFirstName = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include_last_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include_last_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include_last_name")))) {
		request.IncludeLastName = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include_email_addr")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include_email_addr")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include_email_addr")))) {
		request.IncludeEmailAddr = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include_phone_no")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include_phone_no")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include_phone_no")))) {
		request.IncludePhoneNo = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include_company")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include_company")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include_company")))) {
		request.IncludeCompany = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include_location")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include_location")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include_location")))) {
		request.IncludeLocation = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include_sms_provider")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include_sms_provider")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include_sms_provider")))) {
		request.IncludeSmsProvider = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include_person_being_visited")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include_person_being_visited")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include_person_being_visited")))) {
		request.IncludePersonBeingVisited = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include_reason_for_visit")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include_reason_for_visit")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include_reason_for_visit")))) {
		request.IncludeReasonForVisit = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_guest_send_self_using_print")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_guest_send_self_using_print")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_guest_send_self_using_print")))) {
		request.AllowGuestSendSelfUsingPrint = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_guest_send_self_using_email")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_guest_send_self_using_email")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_guest_send_self_using_email")))) {
		request.AllowGuestSendSelfUsingEmail = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_guest_send_self_using_sms")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_guest_send_self_using_sms")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_guest_send_self_using_sms")))) {
		request.AllowGuestSendSelfUsingSms = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include_aup")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include_aup")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include_aup")))) {
		request.IncludeAup = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".aup_on_page")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".aup_on_page")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".aup_on_page")))) {
		request.AupOnPage = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".require_aup_acceptance")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".require_aup_acceptance")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".require_aup_acceptance")))) {
		request.RequireAupAcceptance = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".require_aup_scrolling")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".require_aup_scrolling")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".require_aup_scrolling")))) {
		request.RequireAupScrolling = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_guest_login_from_selfreg_success_page")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_guest_login_from_selfreg_success_page")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_guest_login_from_selfreg_success_page")))) {
		request.AllowGuestLoginFromSelfregSuccessPage = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsAupSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsAupSettings {
	request := isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsAupSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include_aup")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include_aup")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include_aup")))) {
		request.IncludeAup = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".use_diff_aup_for_employees")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".use_diff_aup_for_employees")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".use_diff_aup_for_employees")))) {
		request.UseDiffAupForEmployees = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".skip_aup_for_employees")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".skip_aup_for_employees")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".skip_aup_for_employees")))) {
		request.SkipAupForEmployees = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".require_scrolling")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".require_scrolling")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".require_scrolling")))) {
		request.RequireScrolling = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".require_aup_scrolling")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".require_aup_scrolling")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".require_aup_scrolling")))) {
		request.RequireAupScrolling = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".display_frequency")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".display_frequency")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".display_frequency")))) {
		request.DisplayFrequency = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".display_frequency_interval_days")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".display_frequency_interval_days")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".display_frequency_interval_days")))) {
		request.DisplayFrequencyIntervalDays = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsGuestChangePasswordSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsGuestChangePasswordSettings {
	request := isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsGuestChangePasswordSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_change_passwd_at_first_login")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_change_passwd_at_first_login")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_change_passwd_at_first_login")))) {
		request.AllowChangePasswdAtFirstLogin = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsGuestDeviceRegistrationSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsGuestDeviceRegistrationSettings {
	request := isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsGuestDeviceRegistrationSettings{}
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

func expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsByodSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsByodSettings {
	request := isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsByodSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".byod_welcome_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".byod_welcome_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".byod_welcome_settings")))) {
		request.ByodWelcomeSettings = expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsByodSettingsByodWelcomeSettings(ctx, key+".byod_welcome_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".byod_registration_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".byod_registration_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".byod_registration_settings")))) {
		request.ByodRegistrationSettings = expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsByodSettingsByodRegistrationSettings(ctx, key+".byod_registration_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".byod_registration_success_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".byod_registration_success_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".byod_registration_success_settings")))) {
		request.ByodRegistrationSuccessSettings = expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsByodSettingsByodRegistrationSuccessSettings(ctx, key+".byod_registration_success_settings.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsByodSettingsByodWelcomeSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsByodSettingsByodWelcomeSettings {
	request := isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsByodSettingsByodWelcomeSettings{}
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

func expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsByodSettingsByodRegistrationSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsByodSettingsByodRegistrationSettings {
	request := isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsByodSettingsByodRegistrationSettings{}
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

func expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsByodSettingsByodRegistrationSuccessSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsByodSettingsByodRegistrationSuccessSettings {
	request := isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsByodSettingsByodRegistrationSuccessSettings{}
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

func expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsPostLoginBannerSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsPostLoginBannerSettings {
	request := isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsPostLoginBannerSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include_post_access_banner")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include_post_access_banner")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include_post_access_banner")))) {
		request.IncludePostAccessBanner = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsPostAccessBannerSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsPostAccessBannerSettings {
	request := isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsPostAccessBannerSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include_post_access_banner")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include_post_access_banner")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include_post_access_banner")))) {
		request.IncludePostAccessBanner = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsAuthSuccessSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsAuthSuccessSettings {
	request := isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsAuthSuccessSettings{}
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

func expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSupportInfoSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSupportInfoSettings {
	request := isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalSettingsSupportInfoSettings{}
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

func expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizations(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizations {
	request := isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizations{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".portal_theme")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".portal_theme")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".portal_theme")))) {
		request.PortalTheme = expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizationsPortalTheme(ctx, key+".portal_theme.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".portal_tweak_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".portal_tweak_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".portal_tweak_settings")))) {
		request.PortalTweakSettings = expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizationsPortalTweakSettings(ctx, key+".portal_tweak_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".language")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".language")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".language")))) {
		request.Language = expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizationsLanguage(ctx, key+".language.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".global_customizations")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".global_customizations")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".global_customizations")))) {
		request.GlobalCustomizations = expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizationsGlobalCustomizations(ctx, key+".global_customizations.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".page_customizations")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".page_customizations")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".page_customizations")))) {
		request.PageCustomizations = expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizationsPageCustomizations(ctx, key+".page_customizations.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizationsPortalTheme(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizationsPortalTheme {
	request := isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizationsPortalTheme{}
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

func expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizationsPortalTweakSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizationsPortalTweakSettings {
	request := isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizationsPortalTweakSettings{}
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

func expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizationsLanguage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizationsLanguage {
	request := isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizationsLanguage{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".view_language")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".view_language")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".view_language")))) {
		request.ViewLanguage = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizationsGlobalCustomizations(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizationsGlobalCustomizations {
	request := isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizationsGlobalCustomizations{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mobile_logo_image")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mobile_logo_image")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mobile_logo_image")))) {
		request.MobileLogoImage = expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizationsGlobalCustomizationsMobileLogoImage(ctx, key+".mobile_logo_image.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".desktop_logo_image")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".desktop_logo_image")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".desktop_logo_image")))) {
		request.DesktopLogoImage = expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizationsGlobalCustomizationsDesktopLogoImage(ctx, key+".desktop_logo_image.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".banner_image")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".banner_image")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".banner_image")))) {
		request.BannerImage = expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizationsGlobalCustomizationsBannerImage(ctx, key+".banner_image.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".background_image")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".background_image")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".background_image")))) {
		request.BackgroundImage = expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizationsGlobalCustomizationsBackgroundImage(ctx, key+".background_image.0", d)
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

func expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizationsGlobalCustomizationsMobileLogoImage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizationsGlobalCustomizationsMobileLogoImage {
	request := isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizationsGlobalCustomizationsMobileLogoImage{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data")))) {
		request.Data = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizationsGlobalCustomizationsDesktopLogoImage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizationsGlobalCustomizationsDesktopLogoImage {
	request := isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizationsGlobalCustomizationsDesktopLogoImage{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data")))) {
		request.Data = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizationsGlobalCustomizationsBannerImage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizationsGlobalCustomizationsBannerImage {
	request := isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizationsGlobalCustomizationsBannerImage{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data")))) {
		request.Data = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizationsGlobalCustomizationsBackgroundImage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizationsGlobalCustomizationsBackgroundImage {
	request := isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizationsGlobalCustomizationsBackgroundImage{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data")))) {
		request.Data = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizationsPageCustomizations(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizationsPageCustomizations {
	request := isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizationsPageCustomizations{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data")))) {
		request.Data = expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizationsPageCustomizationsDataArray(ctx, key+".data", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizationsPageCustomizationsDataArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizationsPageCustomizationsData {
	request := []isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizationsPageCustomizationsData{}
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
		i := expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizationsPageCustomizationsData(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizationsPageCustomizationsData(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizationsPageCustomizationsData {
	request := isegosdk.RequestSelfRegisteredPortalCreateSelfRegisteredPortalSelfRegPortalCustomizationsPageCustomizationsData{}
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

func expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByID(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByID {
	request := isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByID{}
	request.SelfRegPortal = expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortal(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortal(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortal {
	request := isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortal{}
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
		request.Settings = expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettings(ctx, key+".settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".customizations")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".customizations")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".customizations")))) {
		request.Customizations = expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizations(ctx, key+".customizations.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettings {
	request := isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".portal_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".portal_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".portal_settings")))) {
		request.PortalSettings = expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsPortalSettings(ctx, key+".portal_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".login_page_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".login_page_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".login_page_settings")))) {
		request.LoginPageSettings = expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsLoginPageSettings(ctx, key+".login_page_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".self_reg_page_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".self_reg_page_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".self_reg_page_settings")))) {
		request.SelfRegPageSettings = expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettings(ctx, key+".self_reg_page_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".self_reg_success_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".self_reg_success_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".self_reg_success_settings")))) {
		request.SelfRegSuccessSettings = expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegSuccessSettings(ctx, key+".self_reg_success_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".aup_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".aup_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".aup_settings")))) {
		request.AupSettings = expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsAupSettings(ctx, key+".aup_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".guest_change_password_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".guest_change_password_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".guest_change_password_settings")))) {
		request.GuestChangePasswordSettings = expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsGuestChangePasswordSettings(ctx, key+".guest_change_password_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".guest_device_registration_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".guest_device_registration_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".guest_device_registration_settings")))) {
		request.GuestDeviceRegistrationSettings = expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsGuestDeviceRegistrationSettings(ctx, key+".guest_device_registration_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".byod_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".byod_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".byod_settings")))) {
		request.ByodSettings = expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsByodSettings(ctx, key+".byod_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".post_login_banner_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".post_login_banner_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".post_login_banner_settings")))) {
		request.PostLoginBannerSettings = expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsPostLoginBannerSettings(ctx, key+".post_login_banner_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".post_access_banner_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".post_access_banner_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".post_access_banner_settings")))) {
		request.PostAccessBannerSettings = expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsPostAccessBannerSettings(ctx, key+".post_access_banner_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".auth_success_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".auth_success_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".auth_success_settings")))) {
		request.AuthSuccessSettings = expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsAuthSuccessSettings(ctx, key+".auth_success_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".support_info_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".support_info_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".support_info_settings")))) {
		request.SupportInfoSettings = expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSupportInfoSettings(ctx, key+".support_info_settings.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsPortalSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsPortalSettings {
	request := isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsPortalSettings{}
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

func expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsLoginPageSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsLoginPageSettings {
	request := isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsLoginPageSettings{}
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
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".alternate_guest_portal")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".alternate_guest_portal")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".alternate_guest_portal")))) {
		request.AlternateGuestPortal = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_guest_to_use_social_accounts")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_guest_to_use_social_accounts")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_guest_to_use_social_accounts")))) {
		request.AllowGuestToUseSocialAccounts = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_show_guest_form")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_show_guest_form")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_show_guest_form")))) {
		request.AllowShowGuestForm = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".social_configs")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".social_configs")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".social_configs")))) {
		request.SocialConfigs = expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsLoginPageSettingsSocialConfigsArray(ctx, key+".social_configs", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsLoginPageSettingsSocialConfigsArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsLoginPageSettingsSocialConfigs {
	request := []isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsLoginPageSettingsSocialConfigs{}
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
		i := expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsLoginPageSettingsSocialConfigs(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsLoginPageSettingsSocialConfigs(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsLoginPageSettingsSocialConfigs {
	request := isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsLoginPageSettingsSocialConfigs{}
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

func expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettings {
	request := isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".assign_guests_to_guest_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".assign_guests_to_guest_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".assign_guests_to_guest_type")))) {
		request.AssignGuestsToGuestType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".account_validity_duration")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".account_validity_duration")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".account_validity_duration")))) {
		request.AccountValidityDuration = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".account_validity_time_units")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".account_validity_time_units")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".account_validity_time_units")))) {
		request.AccountValidityTimeUnits = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".require_registration_code")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".require_registration_code")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".require_registration_code")))) {
		request.RequireRegistrationCode = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".registration_code")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".registration_code")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".registration_code")))) {
		request.RegistrationCode = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".field_user_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".field_user_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".field_user_name")))) {
		request.FieldUserName = expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldUserName(ctx, key+".field_user_name.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".field_first_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".field_first_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".field_first_name")))) {
		request.FieldFirstName = expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldFirstName(ctx, key+".field_first_name.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".field_last_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".field_last_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".field_last_name")))) {
		request.FieldLastName = expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldLastName(ctx, key+".field_last_name.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".field_email_addr")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".field_email_addr")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".field_email_addr")))) {
		request.FieldEmailAddr = expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldEmailAddr(ctx, key+".field_email_addr.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".field_phone_no")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".field_phone_no")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".field_phone_no")))) {
		request.FieldPhoneNo = expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldPhoneNo(ctx, key+".field_phone_no.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".field_company")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".field_company")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".field_company")))) {
		request.FieldCompany = expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldCompany(ctx, key+".field_company.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".field_location")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".field_location")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".field_location")))) {
		request.FieldLocation = expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldLocation(ctx, key+".field_location.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selectable_locations")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selectable_locations")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selectable_locations")))) {
		request.SelectableLocations = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".field_sms_provider")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".field_sms_provider")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".field_sms_provider")))) {
		request.FieldSmsProvider = expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldSmsProvider(ctx, key+".field_sms_provider.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selectable_sms_providers")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selectable_sms_providers")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selectable_sms_providers")))) {
		request.SelectableSmsProviders = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".field_person_being_visited")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".field_person_being_visited")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".field_person_being_visited")))) {
		request.FieldPersonBeingVisited = expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldPersonBeingVisited(ctx, key+".field_person_being_visited.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".field_reason_for_visit")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".field_reason_for_visit")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".field_reason_for_visit")))) {
		request.FieldReasonForVisit = expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldReasonForVisit(ctx, key+".field_reason_for_visit.0", d)
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
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_guest_email_whitelist")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_guest_email_whitelist")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_guest_email_whitelist")))) {
		request.EnableGuestEmailWhitelist = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".guest_email_whitelist_domains")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".guest_email_whitelist_domains")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".guest_email_whitelist_domains")))) {
		request.GuestEmailWhitelistDomains = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_guest_email_blacklist")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_guest_email_blacklist")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_guest_email_blacklist")))) {
		request.EnableGuestEmailBlacklist = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".guest_email_blacklist_domains")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".guest_email_blacklist_domains")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".guest_email_blacklist_domains")))) {
		request.GuestEmailBlacklistDomains = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".require_guest_approval")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".require_guest_approval")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".require_guest_approval")))) {
		request.RequireGuestApproval = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".auto_login_self_wait")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".auto_login_self_wait")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".auto_login_self_wait")))) {
		request.AutoLoginSelfWait = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".auto_login_time_period")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".auto_login_time_period")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".auto_login_time_period")))) {
		request.AutoLoginTimePeriod = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_grace_access")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_grace_access")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_grace_access")))) {
		request.AllowGraceAccess = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".grace_access_expire_interval")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".grace_access_expire_interval")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".grace_access_expire_interval")))) {
		request.GraceAccessExpireInterval = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".grace_access_send_account_expiration")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".grace_access_send_account_expiration")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".grace_access_send_account_expiration")))) {
		request.GraceAccessSendAccountExpiration = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".send_approval_request_to")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".send_approval_request_to")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".send_approval_request_to")))) {
		request.SendApprovalRequestTo = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".approval_email_addresses")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".approval_email_addresses")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".approval_email_addresses")))) {
		request.ApprovalEmailAddresses = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".post_registration_redirect")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".post_registration_redirect")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".post_registration_redirect")))) {
		request.PostRegistrationRedirect = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".post_registration_redirect_url")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".post_registration_redirect_url")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".post_registration_redirect_url")))) {
		request.PostRegistrationRedirectURL = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".credential_notification_using_email")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".credential_notification_using_email")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".credential_notification_using_email")))) {
		request.CredentialNotificationUsingEmail = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".credential_notification_using_sms")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".credential_notification_using_sms")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".credential_notification_using_sms")))) {
		request.CredentialNotificationUsingSms = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".approve_deny_links_valid_for")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".approve_deny_links_valid_for")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".approve_deny_links_valid_for")))) {
		request.ApproveDenyLinksValidFor = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".approve_deny_links_time_units")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".approve_deny_links_time_units")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".approve_deny_links_time_units")))) {
		request.ApproveDenyLinksTimeUnits = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".require_approver_to_authenticate")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".require_approver_to_authenticate")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".require_approver_to_authenticate")))) {
		request.RequireApproverToAuthenticate = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".authenticate_sponsors_using_portal_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".authenticate_sponsors_using_portal_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".authenticate_sponsors_using_portal_list")))) {
		request.AuthenticateSponsorsUsingPortalList = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sponsor_portal_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sponsor_portal_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sponsor_portal_list")))) {
		vSponsorPortalList := v.([]interface{})
		request.SponsorPortalList = &vSponsorPortalList
	}
	return &request
}

func expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldUserName(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldUserName {
	request := isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldUserName{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include")))) {
		request.Include = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".require")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".require")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".require")))) {
		request.Require = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldFirstName(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldFirstName {
	request := isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldFirstName{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include")))) {
		request.Include = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".require")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".require")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".require")))) {
		request.Require = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldLastName(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldLastName {
	request := isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldLastName{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include")))) {
		request.Include = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".require")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".require")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".require")))) {
		request.Require = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldEmailAddr(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldEmailAddr {
	request := isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldEmailAddr{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include")))) {
		request.Include = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".require")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".require")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".require")))) {
		request.Require = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldPhoneNo(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldPhoneNo {
	request := isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldPhoneNo{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include")))) {
		request.Include = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".require")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".require")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".require")))) {
		request.Require = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldCompany(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldCompany {
	request := isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldCompany{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include")))) {
		request.Include = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".require")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".require")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".require")))) {
		request.Require = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldLocation(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldLocation {
	request := isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldLocation{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include")))) {
		request.Include = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".require")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".require")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".require")))) {
		request.Require = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldSmsProvider(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldSmsProvider {
	request := isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldSmsProvider{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include")))) {
		request.Include = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".require")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".require")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".require")))) {
		request.Require = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldPersonBeingVisited(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldPersonBeingVisited {
	request := isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldPersonBeingVisited{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include")))) {
		request.Include = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".require")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".require")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".require")))) {
		request.Require = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldReasonForVisit(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldReasonForVisit {
	request := isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldReasonForVisit{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include")))) {
		request.Include = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".require")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".require")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".require")))) {
		request.Require = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegSuccessSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegSuccessSettings {
	request := isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegSuccessSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include_user_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include_user_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include_user_name")))) {
		request.IncludeUserName = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include_password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include_password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include_password")))) {
		request.IncludePassword = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include_first_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include_first_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include_first_name")))) {
		request.IncludeFirstName = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include_last_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include_last_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include_last_name")))) {
		request.IncludeLastName = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include_email_addr")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include_email_addr")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include_email_addr")))) {
		request.IncludeEmailAddr = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include_phone_no")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include_phone_no")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include_phone_no")))) {
		request.IncludePhoneNo = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include_company")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include_company")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include_company")))) {
		request.IncludeCompany = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include_location")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include_location")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include_location")))) {
		request.IncludeLocation = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include_sms_provider")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include_sms_provider")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include_sms_provider")))) {
		request.IncludeSmsProvider = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include_person_being_visited")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include_person_being_visited")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include_person_being_visited")))) {
		request.IncludePersonBeingVisited = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include_reason_for_visit")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include_reason_for_visit")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include_reason_for_visit")))) {
		request.IncludeReasonForVisit = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_guest_send_self_using_print")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_guest_send_self_using_print")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_guest_send_self_using_print")))) {
		request.AllowGuestSendSelfUsingPrint = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_guest_send_self_using_email")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_guest_send_self_using_email")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_guest_send_self_using_email")))) {
		request.AllowGuestSendSelfUsingEmail = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_guest_send_self_using_sms")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_guest_send_self_using_sms")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_guest_send_self_using_sms")))) {
		request.AllowGuestSendSelfUsingSms = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include_aup")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include_aup")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include_aup")))) {
		request.IncludeAup = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".aup_on_page")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".aup_on_page")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".aup_on_page")))) {
		request.AupOnPage = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".require_aup_acceptance")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".require_aup_acceptance")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".require_aup_acceptance")))) {
		request.RequireAupAcceptance = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".require_aup_scrolling")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".require_aup_scrolling")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".require_aup_scrolling")))) {
		request.RequireAupScrolling = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_guest_login_from_selfreg_success_page")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_guest_login_from_selfreg_success_page")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_guest_login_from_selfreg_success_page")))) {
		request.AllowGuestLoginFromSelfregSuccessPage = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsAupSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsAupSettings {
	request := isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsAupSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include_aup")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include_aup")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include_aup")))) {
		request.IncludeAup = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".use_diff_aup_for_employees")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".use_diff_aup_for_employees")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".use_diff_aup_for_employees")))) {
		request.UseDiffAupForEmployees = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".skip_aup_for_employees")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".skip_aup_for_employees")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".skip_aup_for_employees")))) {
		request.SkipAupForEmployees = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".require_scrolling")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".require_scrolling")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".require_scrolling")))) {
		request.RequireScrolling = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".require_aup_scrolling")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".require_aup_scrolling")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".require_aup_scrolling")))) {
		request.RequireAupScrolling = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".display_frequency")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".display_frequency")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".display_frequency")))) {
		request.DisplayFrequency = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".display_frequency_interval_days")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".display_frequency_interval_days")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".display_frequency_interval_days")))) {
		request.DisplayFrequencyIntervalDays = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsGuestChangePasswordSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsGuestChangePasswordSettings {
	request := isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsGuestChangePasswordSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allow_change_passwd_at_first_login")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allow_change_passwd_at_first_login")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allow_change_passwd_at_first_login")))) {
		request.AllowChangePasswdAtFirstLogin = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsGuestDeviceRegistrationSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsGuestDeviceRegistrationSettings {
	request := isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsGuestDeviceRegistrationSettings{}
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

func expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsByodSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsByodSettings {
	request := isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsByodSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".byod_welcome_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".byod_welcome_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".byod_welcome_settings")))) {
		request.ByodWelcomeSettings = expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsByodSettingsByodWelcomeSettings(ctx, key+".byod_welcome_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".byod_registration_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".byod_registration_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".byod_registration_settings")))) {
		request.ByodRegistrationSettings = expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsByodSettingsByodRegistrationSettings(ctx, key+".byod_registration_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".byod_registration_success_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".byod_registration_success_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".byod_registration_success_settings")))) {
		request.ByodRegistrationSuccessSettings = expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsByodSettingsByodRegistrationSuccessSettings(ctx, key+".byod_registration_success_settings.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsByodSettingsByodWelcomeSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsByodSettingsByodWelcomeSettings {
	request := isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsByodSettingsByodWelcomeSettings{}
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

func expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsByodSettingsByodRegistrationSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsByodSettingsByodRegistrationSettings {
	request := isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsByodSettingsByodRegistrationSettings{}
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

func expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsByodSettingsByodRegistrationSuccessSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsByodSettingsByodRegistrationSuccessSettings {
	request := isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsByodSettingsByodRegistrationSuccessSettings{}
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

func expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsPostLoginBannerSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsPostLoginBannerSettings {
	request := isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsPostLoginBannerSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include_post_access_banner")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include_post_access_banner")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include_post_access_banner")))) {
		request.IncludePostAccessBanner = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsPostAccessBannerSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsPostAccessBannerSettings {
	request := isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsPostAccessBannerSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include_post_access_banner")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include_post_access_banner")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include_post_access_banner")))) {
		request.IncludePostAccessBanner = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsAuthSuccessSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsAuthSuccessSettings {
	request := isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsAuthSuccessSettings{}
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

func expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSupportInfoSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSupportInfoSettings {
	request := isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalSettingsSupportInfoSettings{}
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

func expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizations(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizations {
	request := isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizations{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".portal_theme")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".portal_theme")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".portal_theme")))) {
		request.PortalTheme = expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizationsPortalTheme(ctx, key+".portal_theme.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".portal_tweak_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".portal_tweak_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".portal_tweak_settings")))) {
		request.PortalTweakSettings = expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizationsPortalTweakSettings(ctx, key+".portal_tweak_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".language")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".language")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".language")))) {
		request.Language = expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizationsLanguage(ctx, key+".language.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".global_customizations")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".global_customizations")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".global_customizations")))) {
		request.GlobalCustomizations = expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizationsGlobalCustomizations(ctx, key+".global_customizations.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".page_customizations")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".page_customizations")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".page_customizations")))) {
		request.PageCustomizations = expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizationsPageCustomizations(ctx, key+".page_customizations.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizationsPortalTheme(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizationsPortalTheme {
	request := isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizationsPortalTheme{}
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

func expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizationsPortalTweakSettings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizationsPortalTweakSettings {
	request := isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizationsPortalTweakSettings{}
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

func expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizationsLanguage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizationsLanguage {
	request := isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizationsLanguage{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".view_language")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".view_language")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".view_language")))) {
		request.ViewLanguage = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizationsGlobalCustomizations(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizationsGlobalCustomizations {
	request := isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizationsGlobalCustomizations{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mobile_logo_image")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mobile_logo_image")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mobile_logo_image")))) {
		request.MobileLogoImage = expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizationsGlobalCustomizationsMobileLogoImage(ctx, key+".mobile_logo_image.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".desktop_logo_image")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".desktop_logo_image")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".desktop_logo_image")))) {
		request.DesktopLogoImage = expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizationsGlobalCustomizationsDesktopLogoImage(ctx, key+".desktop_logo_image.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".banner_image")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".banner_image")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".banner_image")))) {
		request.BannerImage = expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizationsGlobalCustomizationsBannerImage(ctx, key+".banner_image.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".background_image")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".background_image")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".background_image")))) {
		request.BackgroundImage = expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizationsGlobalCustomizationsBackgroundImage(ctx, key+".background_image.0", d)
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

func expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizationsGlobalCustomizationsMobileLogoImage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizationsGlobalCustomizationsMobileLogoImage {
	request := isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizationsGlobalCustomizationsMobileLogoImage{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data")))) {
		request.Data = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizationsGlobalCustomizationsDesktopLogoImage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizationsGlobalCustomizationsDesktopLogoImage {
	request := isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizationsGlobalCustomizationsDesktopLogoImage{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data")))) {
		request.Data = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizationsGlobalCustomizationsBannerImage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizationsGlobalCustomizationsBannerImage {
	request := isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizationsGlobalCustomizationsBannerImage{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data")))) {
		request.Data = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizationsGlobalCustomizationsBackgroundImage(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizationsGlobalCustomizationsBackgroundImage {
	request := isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizationsGlobalCustomizationsBackgroundImage{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data")))) {
		request.Data = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizationsPageCustomizations(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizationsPageCustomizations {
	request := isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizationsPageCustomizations{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data")))) {
		request.Data = expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizationsPageCustomizationsDataArray(ctx, key+".data", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizationsPageCustomizationsDataArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizationsPageCustomizationsData {
	request := []isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizationsPageCustomizationsData{}
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
		i := expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizationsPageCustomizationsData(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizationsPageCustomizationsData(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizationsPageCustomizationsData {
	request := isegosdk.RequestSelfRegisteredPortalUpdateSelfRegisteredPortalByIDSelfRegPortalCustomizationsPageCustomizationsData{}
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

func getAllItemsSelfRegisteredPortalGetSelfRegisteredPortals(m interface{}, response *isegosdk.ResponseSelfRegisteredPortalGetSelfRegisteredPortals, queryParams *isegosdk.GetSelfRegisteredPortalsQueryParams) []isegosdk.ResponseSelfRegisteredPortalGetSelfRegisteredPortalsSearchResultResources {
	client := m.(*isegosdk.Client)
	var respItems []isegosdk.ResponseSelfRegisteredPortalGetSelfRegisteredPortalsSearchResultResources
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
			response, _, err = client.SelfRegisteredPortal.GetSelfRegisteredPortals(queryParams)
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

func searchSelfRegisteredPortalGetSelfRegisteredPortals(m interface{}, items []isegosdk.ResponseSelfRegisteredPortalGetSelfRegisteredPortalsSearchResultResources, name string, id string) (*isegosdk.ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortal, error) {
	client := m.(*isegosdk.Client)
	var err error
	var foundItem *isegosdk.ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortal
	for _, item := range items {
		if id != "" && item.ID == id {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseSelfRegisteredPortalGetSelfRegisteredPortalByID
			getItem, _, err = client.SelfRegisteredPortal.GetSelfRegisteredPortalByID(id)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetSelfRegisteredPortalByID")
			}
			foundItem = getItem.SelfRegPortal
			return foundItem, err
		} else if name != "" && item.Name == name {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseSelfRegisteredPortalGetSelfRegisteredPortalByID
			getItem, _, err = client.SelfRegisteredPortal.GetSelfRegisteredPortalByID(item.ID)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetSelfRegisteredPortalByID")
			}
			foundItem = getItem.SelfRegPortal
			return foundItem, err
		}
	}
	return foundItem, err
}
