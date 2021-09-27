
data "ciscoise_endpoint_register" "example" {
  provider = ciscoise
  custom_attributes {
    custom_attributes = {
      key1 = "value1"
      key2 = "value2"
    }
  }
  description       = "string"
  group_id          = "string"
  id                = "string"
  identity_store    = "string"
  identity_store_id = "string"
  mac               = "string"
  mdm_attributes {

    mdm_compliance_status = "false"
    mdm_encrypted         = "false"
    mdm_enrolled          = "false"
    mdm_ime_i             = "string"
    mdm_jail_broken       = "false"
    mdm_manufacturer      = "string"
    mdm_model             = "string"
    mdm_os                = "string"
    mdm_phone_number      = "string"
    mdm_pinlock           = "false"
    mdm_reachable         = "false"
    mdm_serial            = "string"
    mdm_server_name       = "string"
  }
  name                      = "string"
  portal_user               = "string"
  profile_id                = "string"
  static_group_assignment   = "false"
  static_profile_assignment = "false"
}