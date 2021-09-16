
data "ciscoise_active_directory_add_groups" "example" {
  provider = ciscoise
  id       = "string"
  ad_attributes {

    attributes {

      default_value = "string"
      internal_name = "string"
      name          = "string"
      type          = "string"
    }
  }
  ad_scopes_names = "string"
  adgroups {

    groups {

      name = "string"
      sid  = "string"
      type = "string"
    }
  }
  advanced_settings {

    aging_time                        = 1
    auth_protection_type              = "string"
    country                           = "string"
    department                        = "string"
    email                             = "string"
    enable_callback_for_dialin_client = false
    enable_dialin_permission_check    = false
    enable_failed_auth_protection     = false
    enable_machine_access             = false
    enable_machine_auth               = false
    enable_pass_change                = false
    enable_rewrites                   = false
    failed_auth_threshold             = 1
    first_name                        = "string"
    identity_not_in_ad_behaviour      = "string"
    job_title                         = "string"
    last_name                         = "string"
    locality                          = "string"
    organizational_unit               = "string"
    plaintext_auth                    = false
    rewrite_rules {

      rewrite_match  = "string"
      rewrite_result = "string"
      row_id         = 1
    }
    schema                        = "string"
    state_or_province             = "string"
    street_address                = "string"
    telephone                     = "string"
    unreachable_domains_behaviour = "string"
  }
  description              = "string"
  domain                   = "string"
  enable_domain_white_list = false
  name                     = "string"
}