
resource "ciscoise_sponsored_guest_portal" "example" {
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
            
            key = "string"
            value = "string"
          }
        }
        portal_theme {
          
          id = "string"
          name = "string"
          theme_data = "string"
        }
        portal_tweak_settings {
          
          banner_color = "string"
          banner_text_color = "string"
          page_background_color = "string"
          page_label_and_text_color = "string"
        }
      }
      description = "string"
      id = "string"
      name = "string"
      portal_test_url = "string"
      portal_type = "string"
      settings {
        
        aup_settings {
          
          display_frequency = "string"
          display_frequency_interval_days = 1
          include_aup = false
          require_aup_scrolling = false
          require_scrolling = false
          skip_aup_for_employees = false
          use_diff_aup_for_employees = false
        }
        auth_success_settings {
          
          redirect_url = "string"
          success_redirect = "string"
        }
        byod_settings {
          
          byod_registration_settings {
            
            end_point_identity_group_id = "string"
            show_device_id = false
          }
          byod_registration_success_settings {
            
            redirect_url = "string"
            success_redirect = "string"
          }
          byod_welcome_settings {
            
            aup_display = "string"
            enable_byo_d = false
            enable_guest_access = false
            include_aup = false
            require_aup_acceptance = false
            require_mdm = false
            require_scrolling = false
          }
        }
        guest_change_password_settings {
          
          allow_change_passwd_at_first_login = false
        }
        guest_device_registration_settings {
          
          allow_guests_to_register_devices = false
          auto_register_guest_devices = false
        }
        login_page_settings {
          
          access_code = "string"
          allow_alternate_guest_portal = false
          allow_forgot_password = false
          allow_guest_to_change_password = false
          allow_guest_to_create_accounts = false
          aup_display = "string"
          include_aup = false
          max_failed_attempts_before_rate_limit = 1
          require_access_code = false
          require_aup_acceptance = false
          social_configs {
            
            social_media_type = "string"
            social_media_value = "string"
          }
          time_between_logins_during_rate_limit = 1
        }
        portal_settings {
          
          allowed_interfaces = "string"
          always_used_language = "string"
          assigned_guest_type_for_employee = "string"
          authentication_method = "string"
          certificate_group_tag = "string"
          display_lang = "string"
          fallback_language = "string"
          https_port = 1
        }
        post_access_banner_settings {
          
          include_post_access_banner = false
        }
        post_login_banner_settings {
          
          include_post_access_banner = false
        }
        support_info_settings {
          
          default_empty_field_value = "string"
          empty_field_display = "string"
          include_browser_user_agent = false
          include_failure_code = false
          include_ip_address = false
          include_mac_addr = false
          include_policy_server = false
          include_support_info_page = false
        }
      }
    }
}

output "ciscoise_sponsored_guest_portal_example" {
    value = ciscoise_sponsored_guest_portal.example
}