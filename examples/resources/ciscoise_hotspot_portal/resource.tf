
resource "ciscoise_hotspot_portal" "example" {
  provider = ciscoise
  parameters {

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

        access_code         = "string"
        include_aup         = "false"
        require_access_code = "false"
        require_scrolling   = "false"
      }
      auth_success_settings {

        redirect_url     = "string"
        success_redirect = "string"
      }
      portal_settings {

        allowed_interfaces      = ["string"]
        always_used_language    = "string"
        certificate_group_tag   = "string"
        coa_type                = "string"
        display_lang            = "string"
        endpoint_identity_group = "string"
        fallback_language       = "string"
        https_port              = 1
      }
      post_access_banner_settings {

        include_post_access_banner = "false"
      }
      post_login_banner_settings {

        include_post_access_banner = "false"
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

output "ciscoise_hotspot_portal_example" {
  value = ciscoise_hotspot_portal.example
}