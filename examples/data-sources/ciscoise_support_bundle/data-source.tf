
data "ciscoise_support_bundle" "example" {
  provider    = ciscoise
  description = "string"
  host_name   = "string"
  name        = "string"
  support_bundle_include_options {

    from_date           = "string"
    include_config_db   = "false"
    include_core_files  = "false"
    include_debug_logs  = "false"
    include_local_logs  = "false"
    include_system_logs = "false"
    mnt_logs            = "false"
    policy_xml          = "false"
    to_date             = "string"
  }
}