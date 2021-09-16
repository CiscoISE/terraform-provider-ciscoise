
resource "ciscoise_network_device" "example" {
  provider = ciscoise
  item {

    network_device_group_list = ["string"]
    network_device_iplist {

      get_ipaddress_exclude = "string"
      ipaddress             = "string"
      mask                  = 1
    }
    authentication_settings {

      dtls_required                  = false
      enable_key_wrap                = false
      enable_multi_secret            = "string"
      enabled                        = false
      key_encryption_key             = "string"
      key_input_format               = "string"
      message_authenticator_code_key = "string"
      network_protocol               = "string"
      radius_shared_secret           = "string"
      second_radius_shared_secret    = "string"
    }
    coa_port      = 1
    description   = "string"
    dtls_dns_name = "string"
    id            = "string"
    model_name    = "string"
    name          = "string"
    profile_name  = "string"
    snmpsettings {

      link_trap_query                  = false
      mac_trap_query                   = false
      originating_policy_services_node = "string"
      polling_interval                 = 1
      ro_community                     = "string"
      version                          = "string"
    }
    software_version = "string"
    tacacs_settings {

      connect_mode_options = "string"
      shared_secret        = "string"
    }
    trustsecsettings {

      device_authentication_settings {

        sga_device_id       = "string"
        sga_device_password = "string"
      }
      device_configuration_deployment {

        enable_mode_password               = "string"
        exec_mode_password                 = "string"
        exec_mode_username                 = "string"
        include_when_deploying_sgt_updates = false
      }
      push_id_support = false
      sga_notification_and_updates {

        coa_source_host                                    = "string"
        downlaod_environment_data_every_x_seconds          = 1
        downlaod_peer_authorization_policy_every_x_seconds = 1
        download_sga_cllists_every_x_seconds               = 1
        other_sga_devices_to_trust_this_device             = false
        re_authentication_every_x_seconds                  = 1
        send_configuration_to_device                       = false
        send_configuration_to_device_using                 = "string"
      }
    }
  }
}

output "ciscoise_network_device_example" {
  value = ciscoise_network_device.example
}