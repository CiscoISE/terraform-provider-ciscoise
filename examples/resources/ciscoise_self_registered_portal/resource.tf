
resource "ciscoise_self_registered_portal" "example" {
  provider = ciscoise
  item {

    customizations {

      global_customizations {

        background_image {

          data = "string"
        }
        banner_image {

          data = "string"
        }
        banner_title = "string"
        contact_text = "string"
        desktop_logo_image {

          data = "string"
        }
        footer_element = "string"
        mobile_logo_image {

          data = "string"
        }
      }
      language {

        view_language = "string"
      }
      page_customizations {

        data {

          key   = "string"
          value = "string"
        }
      }
      portal_theme {

        id         = "string"
        name       = "string"
        theme_data = "string"
      }
      portal_tweak_settings {

        banner_color              = "string"
        banner_text_color         = "string"
        page_background_color     = "string"
        page_label_and_text_color = "string"
      }
    }
    description     = "string"
    id              = "string"
    name            = "string"
    portal_test_url = "string"
    portal_type     = "string"
    settings {

      aup_settings {

        display_frequency               = "string"
        display_frequency_interval_days = 1
        include_aup                     = "false"
        require_aup_scrolling           = "false"
        require_scrolling               = "false"
        skip_aup_for_employees          = "false"
        use_diff_aup_for_employees      = "false"
      }
      auth_success_settings {

        redirect_url     = "string"
        success_redirect = "string"
      }
      byod_settings {

        byod_registration_settings {

          end_point_identity_group_id = "string"
          show_device_id              = "false"
        }
        byod_registration_success_settings {

          redirect_url     = "string"
          success_redirect = "string"
        }
        byod_welcome_settings {

          aup_display            = "string"
          enable_byo_d           = "false"
          enable_guest_access    = "false"
          include_aup            = "false"
          require_aup_acceptance = "false"
          require_mdm            = "false"
          require_scrolling      = "false"
        }
      }
      guest_change_password_settings {

        allow_change_passwd_at_first_login = "false"
      }
      guest_device_registration_settings {

        allow_guests_to_register_devices = "false"
        auto_register_guest_devices      = "false"
      }
      login_page_settings {

        access_code                           = "string"
        allow_alternate_guest_portal          = "false"
        allow_forgot_password                 = "false"
        allow_guest_to_change_password        = "false"
        allow_guest_to_create_accounts        = "false"
        allow_guest_to_use_social_accounts    = "false"
        allow_show_guest_form                 = "false"
        alternate_guest_portal                = "string"
        aup_display                           = "string"
        include_aup                           = "false"
        max_failed_attempts_before_rate_limit = 1
        require_access_code                   = "false"
        require_aup_acceptance                = "false"
        social_configs {

          social_media_type  = "string"
          social_media_value = "string"
        }
        time_between_logins_during_rate_limit = 1
      }
      portal_settings {

        allowed_interfaces               = "string"
        always_used_language             = "string"
        assigned_guest_type_for_employee = "string"
        authentication_method            = "string"
        certificate_group_tag            = "string"
        display_lang                     = "string"
        fallback_language                = "string"
        https_port                       = 1
      }
      post_access_banner_settings {

        include_post_access_banner = "false"
      }
      post_login_banner_settings {

        include_post_access_banner = "false"
      }
      self_reg_page_settings {

        account_validity_duration               = 1
        account_validity_time_units             = "string"
        allow_grace_access                      = "false"
        approval_email_addresses                = "string"
        approve_deny_links_time_units           = "string"
        approve_deny_links_valid_for            = 1
        assign_guests_to_guest_type             = "string"
        aup_display                             = "string"
        authenticate_sponsors_using_portal_list = "string"
        auto_login_self_wait                    = "false"
        auto_login_time_period                  = 1
        credential_notification_using_email     = "false"
        credential_notification_using_sms       = "false"
        enable_guest_email_blacklist            = "false"
        enable_guest_email_whitelist            = "false"
        field_company {

          include = "false"
          require = "false"
        }
        field_email_addr {

          include = "false"
          require = "false"
        }
        field_first_name {

          include = "false"
          require = "false"
        }
        field_last_name {

          include = "false"
          require = "false"
        }
        field_location {

          include = "false"
          require = "false"
        }
        field_person_being_visited {

          include = "false"
          require = "false"
        }
        field_phone_no {

          include = "false"
          require = "false"
        }
        field_reason_for_visit {

          include = "false"
          require = "false"
        }
        field_sms_provider {

          include = "false"
          require = "false"
        }
        field_user_name {

          include = "false"
          require = "false"
        }
        grace_access_expire_interval         = 1
        grace_access_send_account_expiration = "false"
        guest_email_blacklist_domains        = "string"
        guest_email_whitelist_domains        = "string"
        include_aup                          = "false"
        post_registration_redirect           = "string"
        post_registration_redirect_url       = "string"
        registration_code                    = "string"
        require_approver_to_authenticate     = "false"
        require_aup_acceptance               = "false"
        require_guest_approval               = "false"
        require_registration_code            = "false"
        selectable_locations                 = ["string"]
        selectable_sms_providers             = ["string"]
        send_approval_request_to             = "string"
        sponsor_portal_list                  = ["string"]
      }
      self_reg_success_settings {

        allow_guest_login_from_selfreg_success_page = "false"
        allow_guest_send_self_using_email           = "false"
        allow_guest_send_self_using_print           = "false"
        allow_guest_send_self_using_sms             = "false"
        aup_on_page                                 = "false"
        include_aup                                 = "false"
        include_company                             = "false"
        include_email_addr                          = "false"
        include_first_name                          = "false"
        include_last_name                           = "false"
        include_location                            = "false"
        include_password                            = "false"
        include_person_being_visited                = "false"
        include_phone_no                            = "false"
        include_reason_for_visit                    = "false"
        include_sms_provider                        = "false"
        include_user_name                           = "false"
        require_aup_acceptance                      = "false"
        require_aup_scrolling                       = "false"
      }
      support_info_settings {

        default_empty_field_value  = "string"
        empty_field_display        = "string"
        include_browser_user_agent = "false"
        include_failure_code       = "false"
        include_ip_address         = "false"
        include_mac_addr           = "false"
        include_policy_server      = "false"
        include_support_info_page  = "false"
      }
    }
  }
}

output "ciscoise_self_registered_portal_example" {
  value = ciscoise_self_registered_portal.example
}