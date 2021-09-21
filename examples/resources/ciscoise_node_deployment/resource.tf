
resource "ciscoise_node_deployment" "example" {
  provider = ciscoise
  item {

    administration {

      is_enabled = "false"
      role       = "string"
    }
    fdqn = "string"
    general_settings {

      monitoring {

        enable_pxgrid         = "false"
        is_enabled            = "false"
        is_mnt_dedicated      = "false"
        other_monitoring_node = "string"
        policyservice {

          enable_device_admin_service     = "false"
          enable_nac_service              = "false"
          enable_passive_identity_service = "false"
          enable_profiling_service        = "false"
          enabled                         = "false"
          session_service {

            is_enabled = "false"
            nodegroup  = "string"
          }
          sxpservice {

            is_enabled     = "false"
            user_interface = "string"
          }
        }
        role = "string"
      }
    }
    hostname = "string"
    password = "******"
    profile_configuration {

      active_directory {

        days_before_rescan = 1
        description        = "string"
        enabled            = "false"
      }
      dhcp {

        description = "string"
        enabled     = "false"
        interface   = "string"
        port        = 1
      }
      dhcp_span {

        description = "string"
        enabled     = "false"
        interface   = "string"
      }
      dns {

        description = "string"
        enabled     = "false"
      }
      http {

        description = "string"
        enabled     = "false"
        interface   = "string"
      }
      netflow {

        description = "string"
        enabled     = "false"
        interface   = "string"
        port        = 1
      }
      nmap {

        description = "string"
        enabled     = "false"
      }
      pxgrid {

        description = "string"
        enabled     = "false"
      }
      radius {

        description = "string"
        enabled     = "false"
      }
      snmp_query {

        description   = "string"
        enabled       = "false"
        event_timeout = 1
        retries       = 1
        timeout       = 1
      }
      snmp_trap {

        description     = "string"
        interface       = "string"
        link_trap_query = "false"
        mac_trap_query  = "false"
        port            = 1
      }
    }
    response {

      general_settings {

        monitoring {

          enable_pxgrid         = "false"
          is_enabled            = "false"
          is_mnt_dedicated      = "false"
          other_monitoring_node = "string"
          policyservice {

            enable_device_admin_service     = "false"
            enable_nac_service              = "false"
            enable_passive_identity_service = "false"
            enable_profiling_service        = "false"
            enabled                         = "false"
            session_service {

              is_enabled = "false"
              nodegroup  = "string"
            }
            sxpservice {

              is_enabled     = "false"
              user_interface = "string"
            }
          }
          role = "string"
        }
      }
      profile_configuration {

        active_directory {

          days_before_rescan = 1
          description        = "string"
          enabled            = "false"
        }
        dhcp {

          description = "string"
          enabled     = "false"
          interface   = "string"
          port        = 1
        }
        dhcp_span {

          description = "string"
          enabled     = "false"
          interface   = "string"
        }
        dns {

          description = "string"
          enabled     = "false"
        }
        http {

          description = "string"
          enabled     = "false"
          interface   = "string"
        }
        netflow {

          description = "string"
          enabled     = "false"
          interface   = "string"
          port        = 1
        }
        nmap {

          description = "string"
          enabled     = "false"
        }
        pxgrid {

          description = "string"
          enabled     = "false"
        }
        radius {

          description = "string"
          enabled     = "false"
        }
        snmp_query {

          description   = "string"
          enabled       = "false"
          event_timeout = 1
          retries       = 1
          timeout       = 1
        }
        snmp_trap {

          description     = "string"
          interface       = "string"
          link_trap_query = "false"
          mac_trap_query  = "false"
          port            = 1
        }
      }
    }
    user_name = "string"
  }
}

output "ciscoise_node_deployment_example" {
  value = ciscoise_node_deployment.example
}