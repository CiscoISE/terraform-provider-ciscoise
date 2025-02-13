
resource "ciscoise_ldap" "example" {
    provider = ciscoise
    parameters {
      attributes {
        attributes = ["string"]
      }
      connection_settings {
        always_access_primary_first = "false"
        failback_retry_delay = 1.0
        failover_to_secondary = "false"
        ldap_node_data {
          
          name = "string"
          primary_hostname = "string"
          primary_port = 1.0
          secondary_hostname = "string"
          secondary_port = 1.0
        }
        primary_server {
          admin_dn = "string"
          admin_password = "string"
          enable_force_reconnect = "false"
          enable_secure_connection = "false"
          enable_server_identity_check = "false"
          force_reconnect_time = 1.0
          host_name = "string"
          issuer_cacertificate = "string"
          max_connections = 1.0
          port = 9090
          server_timeout = 1.0
          trust_certificate = "string"
          use_admin_access = "false"
        }
        secondary_server {
          admin_dn = "string"
          admin_password = "string"
          enable_force_reconnect = "false"
          enable_secure_connection = "false"
          enable_server_identity_check = "false"
          force_reconnect_time = 1.0
          host_name = "string"
          issuer_cacertificate = "string"
          max_connections = 1.0
          port = 9090
          server_timeout = 1.0
          trust_certificate = "string"
          use_admin_access = "false"
        }
        specify_server_for_each_ise_node = "false"
      }
      description = "string"
      directory_organization {
        group_directory_subtree = "string"
        mac_format = ------
        prefix_separator = "string"
        strip_prefix = "false"
        strip_suffix = "string"
        suffix_separator = "string"
        user_directory_subtree = "string"
      }
      enable_password_change_lda_p = "false"
      general_settings {
        certificate = "string"
        group_map_attribute_name = "string"
        group_member_reference = ------
        group_name_attribute = "string"
        group_object_class = "string"
        schema = ------
        user_info_attributes {
          additional_attribute = "string"
          country = "string"
          department = "string"
          email = "string"
          first_name = "string"
          job_title = "string"
          last_name = "string"
          locality = "string"
          organizational_unit = "string"
          state_or_province = "string"
          street_address = "string"
          telephone = "string"
        }
        user_name_attribute = "string"
        user_object_class = "string"
      }
      groups {
        groups_names = ["string"]
      }
      id = "string"
      link {
        rel = "string"
        type = "string"
        uri = "string"
      }
      name = "string"
    }
}

output "ciscoise_ldap_example" {
    value = ciscoise_ldap.example
}