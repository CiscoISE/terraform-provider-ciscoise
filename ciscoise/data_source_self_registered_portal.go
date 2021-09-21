package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/CiscoISE/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSelfRegisteredPortal() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on SelfRegisteredPortal.

- This data source allows the client to get a self registered portal by ID.

- This data source allows the client to get all the self registered portals.

Filter:

[name]

Sorting:

[name, description]
`,

		ReadContext: dataSourceSelfRegisteredPortalRead,
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
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"require_aup_scrolling": &schema.Schema{
													Description: `Require the portal user to scroll to the end of the AUP.
Only valid if requireAupAcceptance = true`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"require_scrolling": &schema.Schema{
													// Type:     schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"skip_aup_for_employees": &schema.Schema{
													Description: `Only valid if requireAupAcceptance = trueG`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"use_diff_aup_for_employees": &schema.Schema{
													Description: `Only valid if requireAupAcceptance = trueG`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
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
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
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
																// Type:     schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},
															"enable_guest_access": &schema.Schema{
																// Type:     schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},
															"include_aup": &schema.Schema{
																// Type:     schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},
															"require_aup_acceptance": &schema.Schema{
																// Type:     schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},
															"require_mdm": &schema.Schema{
																// Type:     schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},
															"require_scrolling": &schema.Schema{
																Description: `Require BYOD devices to scroll down to the bottom of the AUP, 
Only valid if includeAup = true`,
																// Type:        schema.TypeBool,
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
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
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

												"allow_guests_to_register_devices": &schema.Schema{
													Description: `Allow guests to register devices`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"auto_register_guest_devices": &schema.Schema{
													Description: `Automatically register guest devices`,
													// Type:        schema.TypeBool,
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

												"access_code": &schema.Schema{
													Description: `Access code that must be entered by the portal user (only valid if requireAccessCode = true)`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"allow_alternate_guest_portal": &schema.Schema{
													// Type:     schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"allow_forgot_password": &schema.Schema{
													// Type:     schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"allow_guest_to_change_password": &schema.Schema{
													Description: `Require the portal user to enter an access code`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"allow_guest_to_create_accounts": &schema.Schema{
													// Type:     schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"allow_guest_to_use_social_accounts": &schema.Schema{
													// Type:     schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"allow_show_guest_form": &schema.Schema{
													// Type:     schema.TypeBool,
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
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"max_failed_attempts_before_rate_limit": &schema.Schema{
													Description: `Maximum failed login attempts before rate limiting`,
													Type:        schema.TypeInt,
													Computed:    true,
												},
												"require_access_code": &schema.Schema{
													Description: `Require the portal user to enter an access code`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"require_aup_acceptance": &schema.Schema{
													Description: `Require the portal user to accept the AUP. Only valid if includeAup = true`,
													// Type:        schema.TypeBool,
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
													Type:     schema.TypeString,
													Computed: true,
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
													// Type:     schema.TypeBool,
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
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
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
													// Type:     schema.TypeBool,
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
													// Type:        schema.TypeBool,
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
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"credential_notification_using_sms": &schema.Schema{
													Description: `If true, send credential notification upon approval using SMS.
Only valid if requireGuestApproval = true`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"enable_guest_email_blacklist": &schema.Schema{
													Description: `Disallow guests with an e-mail address from selected domains`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"enable_guest_email_whitelist": &schema.Schema{
													Description: `Allow guests with an e-mail address from selected domains`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"field_company": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"include": &schema.Schema{
																// Type:     schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},
															"require": &schema.Schema{
																Description: `Only applicable if include = true`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
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
																// Type:     schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},
															"require": &schema.Schema{
																Description: `Only applicable if include = true`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
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
																// Type:     schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},
															"require": &schema.Schema{
																Description: `Only applicable if include = true`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
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
																// Type:     schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},
															"require": &schema.Schema{
																Description: `Only applicable if include = true`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
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
																// Type:     schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},
															"require": &schema.Schema{
																Description: `Only applicable if include = true`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
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
																// Type:     schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},
															"require": &schema.Schema{
																Description: `Only applicable if include = true`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
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
																// Type:     schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},
															"require": &schema.Schema{
																Description: `Only applicable if include = true`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
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
																// Type:     schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},
															"require": &schema.Schema{
																Description: `Only applicable if include = true`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
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
																// Type:     schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},
															"require": &schema.Schema{
																Description: `Only applicable if include = true`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
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
																// Type:     schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},
															"require": &schema.Schema{
																Description: `Only applicable if include = true`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"grace_access_expire_interval": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"grace_access_send_account_expiration": &schema.Schema{
													// Type:     schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"guest_email_blacklist_domains": &schema.Schema{
													Description: `Disallow guests with an e-mail address from selected domains`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"guest_email_whitelist_domains": &schema.Schema{
													Description: `Self-registered guests whose e-mail address is in one of these domains will be allowed.
Only valid if enableGuestEmailWhitelist = true`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"include_aup": &schema.Schema{
													Description: `Include an Acceptable Use Policy (AUP) that should be displayed during login`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
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
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"require_aup_acceptance": &schema.Schema{
													Description: `Require the portal user to accept the AUP. Only valid if includeAup = true`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"require_guest_approval": &schema.Schema{
													Description: `Require self-registered guests to be approved if true`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"require_registration_code": &schema.Schema{
													Description: `Self-registered guests are required to enter a registration code`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
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
													// Type:     schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"allow_guest_send_self_using_email": &schema.Schema{
													// Type:     schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"allow_guest_send_self_using_print": &schema.Schema{
													// Type:     schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"allow_guest_send_self_using_sms": &schema.Schema{
													// Type:     schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"aup_on_page": &schema.Schema{
													// Type:     schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"include_aup": &schema.Schema{
													// Type:     schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"include_company": &schema.Schema{
													// Type:     schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"include_email_addr": &schema.Schema{
													// Type:     schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"include_first_name": &schema.Schema{
													// Type:     schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"include_last_name": &schema.Schema{
													// Type:     schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"include_location": &schema.Schema{
													// Type:     schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"include_password": &schema.Schema{
													// Type:     schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"include_person_being_visited": &schema.Schema{
													// Type:     schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"include_phone_no": &schema.Schema{
													// Type:     schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"include_reason_for_visit": &schema.Schema{
													// Type:     schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"include_sms_provider": &schema.Schema{
													// Type:     schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"include_user_name": &schema.Schema{
													// Type:     schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"require_aup_acceptance": &schema.Schema{
													// Type:     schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"require_aup_scrolling": &schema.Schema{
													// Type:     schema.TypeBool,
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
													// Type:     schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"include_failure_code": &schema.Schema{
													// Type:     schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"include_ip_address": &schema.Schema{
													// Type:     schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"include_mac_addr": &schema.Schema{
													// Type:     schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"include_policy_server": &schema.Schema{
													// Type:     schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"include_support_info_page": &schema.Schema{
													// Type:     schema.TypeBool,
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

func dataSourceSelfRegisteredPortalRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		log.Printf("[DEBUG] Selected method 1: GetSelfRegisteredPortals")
		queryParams1 := isegosdk.GetSelfRegisteredPortalsQueryParams{}

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

		response1, _, err := client.SelfRegisteredPortal.GetSelfRegisteredPortals(&queryParams1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetSelfRegisteredPortals", err,
				"Failure at GetSelfRegisteredPortals, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		var items1 []isegosdk.ResponseSelfRegisteredPortalGetSelfRegisteredPortalsSearchResultResources
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
				response1, _, err = client.SelfRegisteredPortal.GetSelfRegisteredPortals(&queryParams1)
				if err != nil {
					break
				}
				// All is good, continue to the next page
				continue
			}
			// Does not have next page finish iteration
			break
		}
		vItems1 := flattenSelfRegisteredPortalGetSelfRegisteredPortalsItems(&items1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSelfRegisteredPortals response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: GetSelfRegisteredPortalByID")
		vvID := vID.(string)

		response2, _, err := client.SelfRegisteredPortal.GetSelfRegisteredPortalByID(vvID)

		if err != nil || response2 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetSelfRegisteredPortalByID", err,
				"Failure at GetSelfRegisteredPortalByID, unexpected response", ""))
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
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSelfRegisteredPortalGetSelfRegisteredPortalsItems(items *[]isegosdk.ResponseSelfRegisteredPortalGetSelfRegisteredPortalsSearchResultResources) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItem["description"] = item.Description
		respItem["link"] = flattenSelfRegisteredPortalGetSelfRegisteredPortalsItemsLink(item.Link)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSelfRegisteredPortalGetSelfRegisteredPortalsItemsLink(item *isegosdk.ResponseSelfRegisteredPortalGetSelfRegisteredPortalsSearchResultResourcesLink) []map[string]interface{} {
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

func flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItem(item *isegosdk.ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortal) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["description"] = item.Description
	respItem["portal_type"] = item.PortalType
	respItem["portal_test_url"] = item.PortalTestURL
	respItem["settings"] = flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemSettings(item.Settings)
	respItem["customizations"] = flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemCustomizations(item.Customizations)
	respItem["link"] = flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemLink(item.Link)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemSettings(item *isegosdk.ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["portal_settings"] = flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemSettingsPortalSettings(item.PortalSettings)
	respItem["login_page_settings"] = flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemSettingsLoginPageSettings(item.LoginPageSettings)
	respItem["self_reg_page_settings"] = flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemSettingsSelfRegPageSettings(item.SelfRegPageSettings)
	respItem["self_reg_success_settings"] = flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemSettingsSelfRegSuccessSettings(item.SelfRegSuccessSettings)
	respItem["aup_settings"] = flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemSettingsAupSettings(item.AupSettings)
	respItem["guest_change_password_settings"] = flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemSettingsGuestChangePasswordSettings(item.GuestChangePasswordSettings)
	respItem["guest_device_registration_settings"] = flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemSettingsGuestDeviceRegistrationSettings(item.GuestDeviceRegistrationSettings)
	respItem["byod_settings"] = flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemSettingsByodSettings(item.ByodSettings)
	respItem["post_login_banner_settings"] = flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemSettingsPostLoginBannerSettings(item.PostLoginBannerSettings)
	respItem["post_access_banner_settings"] = flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemSettingsPostAccessBannerSettings(item.PostAccessBannerSettings)
	respItem["auth_success_settings"] = flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemSettingsAuthSuccessSettings(item.AuthSuccessSettings)
	respItem["support_info_settings"] = flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemSettingsSupportInfoSettings(item.SupportInfoSettings)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemSettingsPortalSettings(item *isegosdk.ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsPortalSettings) []map[string]interface{} {
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

func flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemSettingsLoginPageSettings(item *isegosdk.ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsLoginPageSettings) []map[string]interface{} {
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
	respItem["alternate_guest_portal"] = item.AlternateGuestPortal
	respItem["allow_guest_to_use_social_accounts"] = boolPtrToString(item.AllowGuestToUseSocialAccounts)
	respItem["allow_show_guest_form"] = boolPtrToString(item.AllowShowGuestForm)
	respItem["social_configs"] = flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemSettingsLoginPageSettingsSocialConfigs(item.SocialConfigs)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemSettingsLoginPageSettingsSocialConfigs(items *[]isegosdk.ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsLoginPageSettingsSocialConfigs) []map[string]interface{} {
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

func flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemSettingsSelfRegPageSettings(item *isegosdk.ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["assign_guests_to_guest_type"] = item.AssignGuestsToGuestType
	respItem["account_validity_duration"] = item.AccountValidityDuration
	respItem["account_validity_time_units"] = item.AccountValidityTimeUnits
	respItem["require_registration_code"] = boolPtrToString(item.RequireRegistrationCode)
	respItem["registration_code"] = item.RegistrationCode
	respItem["field_user_name"] = flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemSettingsSelfRegPageSettingsFieldUserName(item.FieldUserName)
	respItem["field_first_name"] = flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemSettingsSelfRegPageSettingsFieldFirstName(item.FieldFirstName)
	respItem["field_last_name"] = flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemSettingsSelfRegPageSettingsFieldLastName(item.FieldLastName)
	respItem["field_email_addr"] = flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemSettingsSelfRegPageSettingsFieldEmailAddr(item.FieldEmailAddr)
	respItem["field_phone_no"] = flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemSettingsSelfRegPageSettingsFieldPhoneNo(item.FieldPhoneNo)
	respItem["field_company"] = flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemSettingsSelfRegPageSettingsFieldCompany(item.FieldCompany)
	respItem["field_location"] = flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemSettingsSelfRegPageSettingsFieldLocation(item.FieldLocation)
	respItem["selectable_locations"] = item.SelectableLocations
	respItem["field_sms_provider"] = flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemSettingsSelfRegPageSettingsFieldSmsProvider(item.FieldSmsProvider)
	respItem["selectable_sms_providers"] = item.SelectableSmsProviders
	respItem["field_person_being_visited"] = flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemSettingsSelfRegPageSettingsFieldPersonBeingVisited(item.FieldPersonBeingVisited)
	respItem["field_reason_for_visit"] = flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemSettingsSelfRegPageSettingsFieldReasonForVisit(item.FieldReasonForVisit)
	respItem["include_aup"] = boolPtrToString(item.IncludeAup)
	respItem["aup_display"] = item.AupDisplay
	respItem["require_aup_acceptance"] = boolPtrToString(item.RequireAupAcceptance)
	respItem["enable_guest_email_whitelist"] = boolPtrToString(item.EnableGuestEmailWhitelist)
	respItem["guest_email_whitelist_domains"] = item.GuestEmailWhitelistDomains
	respItem["enable_guest_email_blacklist"] = boolPtrToString(item.EnableGuestEmailBlacklist)
	respItem["guest_email_blacklist_domains"] = item.GuestEmailBlacklistDomains
	respItem["require_guest_approval"] = boolPtrToString(item.RequireGuestApproval)
	respItem["auto_login_self_wait"] = boolPtrToString(item.AutoLoginSelfWait)
	respItem["auto_login_time_period"] = item.AutoLoginTimePeriod
	respItem["allow_grace_access"] = boolPtrToString(item.AllowGraceAccess)
	respItem["grace_access_expire_interval"] = item.GraceAccessExpireInterval
	respItem["grace_access_send_account_expiration"] = boolPtrToString(item.GraceAccessSendAccountExpiration)
	respItem["send_approval_request_to"] = item.SendApprovalRequestTo
	respItem["approval_email_addresses"] = item.ApprovalEmailAddresses
	respItem["post_registration_redirect"] = item.PostRegistrationRedirect
	respItem["post_registration_redirect_url"] = item.PostRegistrationRedirectURL
	respItem["credential_notification_using_email"] = boolPtrToString(item.CredentialNotificationUsingEmail)
	respItem["credential_notification_using_sms"] = boolPtrToString(item.CredentialNotificationUsingSms)
	respItem["approve_deny_links_valid_for"] = item.ApproveDenyLinksValidFor
	respItem["approve_deny_links_time_units"] = item.ApproveDenyLinksTimeUnits
	respItem["require_approver_to_authenticate"] = boolPtrToString(item.RequireApproverToAuthenticate)
	respItem["authenticate_sponsors_using_portal_list"] = item.AuthenticateSponsorsUsingPortalList
	respItem["sponsor_portal_list"] = responseInterfaceToSliceString(item.SponsorPortalList)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemSettingsSelfRegPageSettingsFieldUserName(item *isegosdk.ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldUserName) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["include"] = boolPtrToString(item.Include)
	respItem["require"] = boolPtrToString(item.Require)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemSettingsSelfRegPageSettingsFieldFirstName(item *isegosdk.ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldFirstName) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["include"] = boolPtrToString(item.Include)
	respItem["require"] = boolPtrToString(item.Require)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemSettingsSelfRegPageSettingsFieldLastName(item *isegosdk.ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldLastName) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["include"] = boolPtrToString(item.Include)
	respItem["require"] = boolPtrToString(item.Require)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemSettingsSelfRegPageSettingsFieldEmailAddr(item *isegosdk.ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldEmailAddr) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["include"] = boolPtrToString(item.Include)
	respItem["require"] = boolPtrToString(item.Require)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemSettingsSelfRegPageSettingsFieldPhoneNo(item *isegosdk.ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldPhoneNo) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["include"] = boolPtrToString(item.Include)
	respItem["require"] = boolPtrToString(item.Require)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemSettingsSelfRegPageSettingsFieldCompany(item *isegosdk.ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldCompany) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["include"] = boolPtrToString(item.Include)
	respItem["require"] = boolPtrToString(item.Require)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemSettingsSelfRegPageSettingsFieldLocation(item *isegosdk.ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldLocation) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["include"] = boolPtrToString(item.Include)
	respItem["require"] = boolPtrToString(item.Require)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemSettingsSelfRegPageSettingsFieldSmsProvider(item *isegosdk.ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldSmsProvider) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["include"] = boolPtrToString(item.Include)
	respItem["require"] = boolPtrToString(item.Require)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemSettingsSelfRegPageSettingsFieldPersonBeingVisited(item *isegosdk.ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldPersonBeingVisited) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["include"] = boolPtrToString(item.Include)
	respItem["require"] = boolPtrToString(item.Require)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemSettingsSelfRegPageSettingsFieldReasonForVisit(item *isegosdk.ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldReasonForVisit) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["include"] = boolPtrToString(item.Include)
	respItem["require"] = boolPtrToString(item.Require)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemSettingsSelfRegSuccessSettings(item *isegosdk.ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegSuccessSettings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["include_user_name"] = boolPtrToString(item.IncludeUserName)
	respItem["include_password"] = boolPtrToString(item.IncludePassword)
	respItem["include_first_name"] = boolPtrToString(item.IncludeFirstName)
	respItem["include_last_name"] = boolPtrToString(item.IncludeLastName)
	respItem["include_email_addr"] = boolPtrToString(item.IncludeEmailAddr)
	respItem["include_phone_no"] = boolPtrToString(item.IncludePhoneNo)
	respItem["include_company"] = boolPtrToString(item.IncludeCompany)
	respItem["include_location"] = boolPtrToString(item.IncludeLocation)
	respItem["include_sms_provider"] = boolPtrToString(item.IncludeSmsProvider)
	respItem["include_person_being_visited"] = boolPtrToString(item.IncludePersonBeingVisited)
	respItem["include_reason_for_visit"] = boolPtrToString(item.IncludeReasonForVisit)
	respItem["allow_guest_send_self_using_print"] = boolPtrToString(item.AllowGuestSendSelfUsingPrint)
	respItem["allow_guest_send_self_using_email"] = boolPtrToString(item.AllowGuestSendSelfUsingEmail)
	respItem["allow_guest_send_self_using_sms"] = boolPtrToString(item.AllowGuestSendSelfUsingSms)
	respItem["include_aup"] = boolPtrToString(item.IncludeAup)
	respItem["aup_on_page"] = boolPtrToString(item.AupOnPage)
	respItem["require_aup_acceptance"] = boolPtrToString(item.RequireAupAcceptance)
	respItem["require_aup_scrolling"] = boolPtrToString(item.RequireAupScrolling)
	respItem["allow_guest_login_from_selfreg_success_page"] = boolPtrToString(item.AllowGuestLoginFromSelfregSuccessPage)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemSettingsAupSettings(item *isegosdk.ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsAupSettings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["include_aup"] = boolPtrToString(item.IncludeAup)
	respItem["use_diff_aup_for_employees"] = boolPtrToString(item.UseDiffAupForEmployees)
	respItem["skip_aup_for_employees"] = boolPtrToString(item.SkipAupForEmployees)
	respItem["require_scrolling"] = boolPtrToString(item.RequireScrolling)
	respItem["require_aup_scrolling"] = boolPtrToString(item.RequireAupScrolling)
	respItem["display_frequency"] = item.DisplayFrequency
	respItem["display_frequency_interval_days"] = item.DisplayFrequencyIntervalDays

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemSettingsGuestChangePasswordSettings(item *isegosdk.ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsGuestChangePasswordSettings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["allow_change_passwd_at_first_login"] = boolPtrToString(item.AllowChangePasswdAtFirstLogin)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemSettingsGuestDeviceRegistrationSettings(item *isegosdk.ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsGuestDeviceRegistrationSettings) []map[string]interface{} {
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

func flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemSettingsByodSettings(item *isegosdk.ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsByodSettings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["byod_welcome_settings"] = flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemSettingsByodSettingsByodWelcomeSettings(item.ByodWelcomeSettings)
	respItem["byod_registration_settings"] = flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemSettingsByodSettingsByodRegistrationSettings(item.ByodRegistrationSettings)
	respItem["byod_registration_success_settings"] = flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemSettingsByodSettingsByodRegistrationSuccessSettings(item.ByodRegistrationSuccessSettings)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemSettingsByodSettingsByodWelcomeSettings(item *isegosdk.ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsByodSettingsByodWelcomeSettings) []map[string]interface{} {
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

func flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemSettingsByodSettingsByodRegistrationSettings(item *isegosdk.ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsByodSettingsByodRegistrationSettings) []map[string]interface{} {
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

func flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemSettingsByodSettingsByodRegistrationSuccessSettings(item *isegosdk.ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsByodSettingsByodRegistrationSuccessSettings) []map[string]interface{} {
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

func flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemSettingsPostLoginBannerSettings(item *isegosdk.ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsPostLoginBannerSettings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["include_post_access_banner"] = boolPtrToString(item.IncludePostAccessBanner)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemSettingsPostAccessBannerSettings(item *isegosdk.ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsPostAccessBannerSettings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["include_post_access_banner"] = boolPtrToString(item.IncludePostAccessBanner)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemSettingsAuthSuccessSettings(item *isegosdk.ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsAuthSuccessSettings) []map[string]interface{} {
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

func flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemSettingsSupportInfoSettings(item *isegosdk.ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsSupportInfoSettings) []map[string]interface{} {
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

func flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemCustomizations(item *isegosdk.ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalCustomizations) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["portal_theme"] = flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemCustomizationsPortalTheme(item.PortalTheme)
	respItem["portal_tweak_settings"] = flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemCustomizationsPortalTweakSettings(item.PortalTweakSettings)
	respItem["language"] = flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemCustomizationsLanguage(item.Language)
	respItem["global_customizations"] = flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemCustomizationsGlobalCustomizations(item.GlobalCustomizations)
	respItem["page_customizations"] = flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemCustomizationsPageCustomizations(item.PageCustomizations)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemCustomizationsPortalTheme(item *isegosdk.ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalCustomizationsPortalTheme) []map[string]interface{} {
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

func flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemCustomizationsPortalTweakSettings(item *isegosdk.ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalCustomizationsPortalTweakSettings) []map[string]interface{} {
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

func flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemCustomizationsLanguage(item *isegosdk.ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalCustomizationsLanguage) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["view_language"] = item.ViewLanguage

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemCustomizationsGlobalCustomizations(item *isegosdk.ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalCustomizationsGlobalCustomizations) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["mobile_logo_image"] = flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemCustomizationsGlobalCustomizationsMobileLogoImage(item.MobileLogoImage)
	respItem["desktop_logo_image"] = flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemCustomizationsGlobalCustomizationsDesktopLogoImage(item.DesktopLogoImage)
	respItem["banner_image"] = flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemCustomizationsGlobalCustomizationsBannerImage(item.BannerImage)
	respItem["background_image"] = flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemCustomizationsGlobalCustomizationsBackgroundImage(item.BackgroundImage)
	respItem["banner_title"] = item.BannerTitle
	respItem["contact_text"] = item.ContactText
	respItem["footer_element"] = item.FooterElement

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemCustomizationsGlobalCustomizationsMobileLogoImage(item *isegosdk.ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalCustomizationsGlobalCustomizationsMobileLogoImage) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["data"] = item.Data

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemCustomizationsGlobalCustomizationsDesktopLogoImage(item *isegosdk.ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalCustomizationsGlobalCustomizationsDesktopLogoImage) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["data"] = item.Data

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemCustomizationsGlobalCustomizationsBannerImage(item *isegosdk.ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalCustomizationsGlobalCustomizationsBannerImage) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["data"] = item.Data

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemCustomizationsGlobalCustomizationsBackgroundImage(item *isegosdk.ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalCustomizationsGlobalCustomizationsBackgroundImage) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["data"] = item.Data

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemCustomizationsPageCustomizations(item *isegosdk.ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalCustomizationsPageCustomizations) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["data"] = flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemCustomizationsPageCustomizationsData(item.Data)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemCustomizationsPageCustomizationsData(items *[]isegosdk.ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalCustomizationsPageCustomizationsData) []map[string]interface{} {
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

func flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemLink(item *isegosdk.ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalLink) []map[string]interface{} {
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
