
resource "ciscoise_authorization_profile" "example" {
  provider = ciscoise
  item {

    access_type = "string"
    acl         = "string"
    advanced_attributes {

      left_hand_side_dictionary_attribue {

        advanced_attribute_value_type = "string"
        attribute_name                = "string"
        dictionary_name               = "string"
        value                         = "string"
      }
      right_hand_side_attribue_value {

        advanced_attribute_value_type = "string"
        attribute_name                = "string"
        dictionary_name               = "string"
        value                         = "string"
      }
    }
    agentless_posture           = "false"
    airespace_acl               = "string"
    airespace_ipv6_acl          = "string"
    asa_vpn                     = "string"
    authz_profile_type          = "string"
    auto_smart_port             = "string"
    avc_profile                 = "string"
    dacl_name                   = "string"
    description                 = "string"
    easywired_session_candidate = "false"
    id                          = "string"
    interface_template          = "string"
    ipv6_acl_filter             = "string"
    ipv6_dacl_name              = "string"
    mac_sec_policy              = "string"
    name                        = "string"
    neat                        = "false"
    profile_name                = "string"
    reauth {

      connectivity = "string"
      timer        = 1
    }
    service_template = "false"
    track_movement   = "false"
    vlan {

      name_id = "string"
      tag_id  = 1
    }
    voice_domain_permission = "false"
    web_auth                = "false"
    web_redirection {

      web_redirection_type                  = "string"
      acl                                   = "string"
      display_certificates_renewal_messages = "false"
      portal_name                           = "string"
      static_iphost_name_fqd_n              = "string"
    }
  }
}

output "ciscoise_authorization_profile_example" {
  value = ciscoise_authorization_profile.example
}