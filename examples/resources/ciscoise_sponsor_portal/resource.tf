
resource "ciscoise_sponsor_portal" "example" {
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
        require_scrolling               = "false"
      }
      login_page_settings {

        aup_display                           = "string"
        include_aup                           = "false"
        max_failed_attempts_before_rate_limit = 1
        require_aup_acceptance                = "false"
        require_aup_scrolling                 = "false"
        social_configs                        = ["string"]
        time_between_logins_during_rate_limit = 1
      }
      portal_settings {

        allowed_interfaces    = ["string"]
        authentication_method = "string"
        available_ssids       = ["string"]
        certificate_group_tag = "string"
        display_lang          = "string"
        fallback_language     = "string"
        fqdn                  = "string"
        https_port            = 1
        idle_timeout          = 1
      }
      post_access_banner_settings {

        include_post_access_banner = "false"
      }
      post_login_banner_settings {

        include_post_access_banner = "false"
      }
      sponsor_change_password_settings {

        allow_sponsor_to_change_pwd = "false"
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

output "ciscoise_sponsor_portal_example" {
  value = ciscoise_sponsor_portal.example
}