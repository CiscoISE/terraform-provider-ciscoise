package ciscoise

import (
	"context"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSelfRegisteredPortal() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSelfRegisteredPortalRead,
		Schema: map[string]*schema.Schema{
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
																Type:     schema.TypeString,
																Computed: true,
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
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
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
							Type:     schema.TypeString,
							Computed: true,
						},
						"portal_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"settings": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"aup_settings": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"display_frequency": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"display_frequency_interval_days": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"include_aup": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"require_aup_scrolling": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"require_scrolling": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"skip_aup_for_employees": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"use_diff_aup_for_employees": &schema.Schema{
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
										Type:     schema.TypeList,
										Computed: true,
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
																Type:     schema.TypeBool,
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
												"byod_welcome_settings": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"aup_display": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"enable_byo_d": &schema.Schema{
																Type:     schema.TypeBool,
																Computed: true,
															},
															"enable_guest_access": &schema.Schema{
																Type:     schema.TypeBool,
																Computed: true,
															},
															"include_aup": &schema.Schema{
																Type:     schema.TypeBool,
																Computed: true,
															},
															"require_aup_acceptance": &schema.Schema{
																Type:     schema.TypeBool,
																Computed: true,
															},
															"require_mdm": &schema.Schema{
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

												"allow_guests_to_register_devices": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"auto_register_guest_devices": &schema.Schema{
													Type:     schema.TypeBool,
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

												"access_code": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"allow_alternate_guest_portal": &schema.Schema{
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
												"allow_guest_to_create_accounts": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"allow_guest_to_use_social_accounts": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"allow_show_guest_form": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"alternate_guest_portal": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"aup_display": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"include_aup": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"max_failed_attempts_before_rate_limit": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"require_access_code": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"require_aup_acceptance": &schema.Schema{
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
												"time_between_logins_during_rate_limit": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},
									"portal_settings": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"allowed_interfaces": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"always_used_language": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"assigned_guest_type_for_employee": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"authentication_method": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"certificate_group_tag": &schema.Schema{
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
												"https_port": &schema.Schema{
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
									"self_reg_page_settings": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"account_validity_duration": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"account_validity_time_units": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"allow_grace_access": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"approval_email_addresses": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"approve_deny_links_time_units": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"approve_deny_links_valid_for": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"assign_guests_to_guest_type": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"aup_display": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"authenticate_sponsors_using_portal_list": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"auto_login_self_wait": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"auto_login_time_period": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"credential_notification_using_email": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"credential_notification_using_sms": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"enable_guest_email_blacklist": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"enable_guest_email_whitelist": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"field_company": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"include": &schema.Schema{
																Type:     schema.TypeBool,
																Computed: true,
															},
															"require": &schema.Schema{
																Type:     schema.TypeBool,
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
																Type:     schema.TypeBool,
																Computed: true,
															},
															"require": &schema.Schema{
																Type:     schema.TypeBool,
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
																Type:     schema.TypeBool,
																Computed: true,
															},
															"require": &schema.Schema{
																Type:     schema.TypeBool,
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
																Type:     schema.TypeBool,
																Computed: true,
															},
															"require": &schema.Schema{
																Type:     schema.TypeBool,
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
																Type:     schema.TypeBool,
																Computed: true,
															},
															"require": &schema.Schema{
																Type:     schema.TypeBool,
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
																Type:     schema.TypeBool,
																Computed: true,
															},
															"require": &schema.Schema{
																Type:     schema.TypeBool,
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
																Type:     schema.TypeBool,
																Computed: true,
															},
															"require": &schema.Schema{
																Type:     schema.TypeBool,
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
																Type:     schema.TypeBool,
																Computed: true,
															},
															"require": &schema.Schema{
																Type:     schema.TypeBool,
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
																Type:     schema.TypeBool,
																Computed: true,
															},
															"require": &schema.Schema{
																Type:     schema.TypeBool,
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
																Type:     schema.TypeBool,
																Computed: true,
															},
															"require": &schema.Schema{
																Type:     schema.TypeBool,
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
													Type:     schema.TypeBool,
													Computed: true,
												},
												"guest_email_blacklist_domains": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"guest_email_whitelist_domains": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"include_aup": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"post_registration_redirect": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"post_registration_redirect_url": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"registration_code": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"require_approver_to_authenticate": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"require_aup_acceptance": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"require_guest_approval": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"require_registration_code": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"selectable_locations": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"selectable_sms_providers": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"send_approval_request_to": &schema.Schema{
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
													Type:     schema.TypeBool,
													Computed: true,
												},
												"allow_guest_send_self_using_email": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"allow_guest_send_self_using_print": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"allow_guest_send_self_using_sms": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"aup_on_page": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"include_aup": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"include_company": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"include_email_addr": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"include_first_name": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"include_last_name": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"include_location": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"include_password": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"include_person_being_visited": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"include_phone_no": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"include_reason_for_visit": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"include_sms_provider": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"include_user_name": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"require_aup_acceptance": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"require_aup_scrolling": &schema.Schema{
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

												"default_empty_field_value": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"empty_field_display": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"include_browser_user_agent": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"include_failure_code": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"include_ip_address": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"include_mac_addr": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"include_policy_server": &schema.Schema{
													Type:     schema.TypeBool,
													Computed: true,
												},
												"include_support_info_page": &schema.Schema{
													Type:     schema.TypeBool,
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
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

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

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

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

		log.Printf("[DEBUG] Retrieved response %+v", *response2)

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
	respItem["alternate_guest_portal"] = item.AlternateGuestPortal
	respItem["allow_guest_to_use_social_accounts"] = item.AllowGuestToUseSocialAccounts
	respItem["allow_show_guest_form"] = item.AllowShowGuestForm
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
	respItem["require_registration_code"] = item.RequireRegistrationCode
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
	respItem["include_aup"] = item.IncludeAup
	respItem["aup_display"] = item.AupDisplay
	respItem["require_aup_acceptance"] = item.RequireAupAcceptance
	respItem["enable_guest_email_whitelist"] = item.EnableGuestEmailWhitelist
	respItem["guest_email_whitelist_domains"] = item.GuestEmailWhitelistDomains
	respItem["enable_guest_email_blacklist"] = item.EnableGuestEmailBlacklist
	respItem["guest_email_blacklist_domains"] = item.GuestEmailBlacklistDomains
	respItem["require_guest_approval"] = item.RequireGuestApproval
	respItem["auto_login_self_wait"] = item.AutoLoginSelfWait
	respItem["auto_login_time_period"] = item.AutoLoginTimePeriod
	respItem["allow_grace_access"] = item.AllowGraceAccess
	respItem["grace_access_expire_interval"] = item.GraceAccessExpireInterval
	respItem["grace_access_send_account_expiration"] = item.GraceAccessSendAccountExpiration
	respItem["send_approval_request_to"] = item.SendApprovalRequestTo
	respItem["approval_email_addresses"] = item.ApprovalEmailAddresses
	respItem["post_registration_redirect"] = item.PostRegistrationRedirect
	respItem["post_registration_redirect_url"] = item.PostRegistrationRedirectURL
	respItem["credential_notification_using_email"] = item.CredentialNotificationUsingEmail
	respItem["credential_notification_using_sms"] = item.CredentialNotificationUsingSms
	respItem["approve_deny_links_valid_for"] = item.ApproveDenyLinksValidFor
	respItem["approve_deny_links_time_units"] = item.ApproveDenyLinksTimeUnits
	respItem["require_approver_to_authenticate"] = item.RequireApproverToAuthenticate
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
	respItem["include"] = item.Include
	respItem["require"] = item.Require

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemSettingsSelfRegPageSettingsFieldFirstName(item *isegosdk.ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldFirstName) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["include"] = item.Include
	respItem["require"] = item.Require

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemSettingsSelfRegPageSettingsFieldLastName(item *isegosdk.ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldLastName) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["include"] = item.Include
	respItem["require"] = item.Require

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemSettingsSelfRegPageSettingsFieldEmailAddr(item *isegosdk.ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldEmailAddr) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["include"] = item.Include
	respItem["require"] = item.Require

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemSettingsSelfRegPageSettingsFieldPhoneNo(item *isegosdk.ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldPhoneNo) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["include"] = item.Include
	respItem["require"] = item.Require

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemSettingsSelfRegPageSettingsFieldCompany(item *isegosdk.ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldCompany) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["include"] = item.Include
	respItem["require"] = item.Require

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemSettingsSelfRegPageSettingsFieldLocation(item *isegosdk.ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldLocation) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["include"] = item.Include
	respItem["require"] = item.Require

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemSettingsSelfRegPageSettingsFieldSmsProvider(item *isegosdk.ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldSmsProvider) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["include"] = item.Include
	respItem["require"] = item.Require

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemSettingsSelfRegPageSettingsFieldPersonBeingVisited(item *isegosdk.ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldPersonBeingVisited) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["include"] = item.Include
	respItem["require"] = item.Require

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemSettingsSelfRegPageSettingsFieldReasonForVisit(item *isegosdk.ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegPageSettingsFieldReasonForVisit) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["include"] = item.Include
	respItem["require"] = item.Require

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemSettingsSelfRegSuccessSettings(item *isegosdk.ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsSelfRegSuccessSettings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["include_user_name"] = item.IncludeUserName
	respItem["include_password"] = item.IncludePassword
	respItem["include_first_name"] = item.IncludeFirstName
	respItem["include_last_name"] = item.IncludeLastName
	respItem["include_email_addr"] = item.IncludeEmailAddr
	respItem["include_phone_no"] = item.IncludePhoneNo
	respItem["include_company"] = item.IncludeCompany
	respItem["include_location"] = item.IncludeLocation
	respItem["include_sms_provider"] = item.IncludeSmsProvider
	respItem["include_person_being_visited"] = item.IncludePersonBeingVisited
	respItem["include_reason_for_visit"] = item.IncludeReasonForVisit
	respItem["allow_guest_send_self_using_print"] = item.AllowGuestSendSelfUsingPrint
	respItem["allow_guest_send_self_using_email"] = item.AllowGuestSendSelfUsingEmail
	respItem["allow_guest_send_self_using_sms"] = item.AllowGuestSendSelfUsingSms
	respItem["include_aup"] = item.IncludeAup
	respItem["aup_on_page"] = item.AupOnPage
	respItem["require_aup_acceptance"] = item.RequireAupAcceptance
	respItem["require_aup_scrolling"] = item.RequireAupScrolling
	respItem["allow_guest_login_from_selfreg_success_page"] = item.AllowGuestLoginFromSelfregSuccessPage

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemSettingsAupSettings(item *isegosdk.ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsAupSettings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["include_aup"] = item.IncludeAup
	respItem["use_diff_aup_for_employees"] = item.UseDiffAupForEmployees
	respItem["skip_aup_for_employees"] = item.SkipAupForEmployees
	respItem["require_scrolling"] = item.RequireScrolling
	respItem["require_aup_scrolling"] = item.RequireAupScrolling
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
	respItem["allow_change_passwd_at_first_login"] = item.AllowChangePasswdAtFirstLogin

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemSettingsGuestDeviceRegistrationSettings(item *isegosdk.ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsGuestDeviceRegistrationSettings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["auto_register_guest_devices"] = item.AutoRegisterGuestDevices
	respItem["allow_guests_to_register_devices"] = item.AllowGuestsToRegisterDevices

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

func flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemSettingsByodSettingsByodRegistrationSettings(item *isegosdk.ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsByodSettingsByodRegistrationSettings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["show_device_id"] = item.ShowDeviceID
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
	respItem["include_post_access_banner"] = item.IncludePostAccessBanner

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSelfRegisteredPortalGetSelfRegisteredPortalByIDItemSettingsPostAccessBannerSettings(item *isegosdk.ResponseSelfRegisteredPortalGetSelfRegisteredPortalByIDSelfRegPortalSettingsPostAccessBannerSettings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["include_post_access_banner"] = item.IncludePostAccessBanner

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
