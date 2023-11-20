terraform {
  required_providers {
    ciscoise = {
      version = "0.7.0-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

resource "ciscoise_device_administration_policy_set" "example" {
  provider = ciscoise
  parameters {
    rank        = 0
    state       = "enabled"
    name        = "ASA Firewalls"
    description = "ASA Firewalls"
    condition {
      condition_type  = "ConditionAttributes"
      is_negate       = "false"
      dictionary_name = "DEVICE"
      attribute_name  = "Device Type"
      operator        = "startsWith"
      attribute_value = "All Device Types"
    }
    service_name = "Default Device Admin"
  }
}

output "ciscoise_device_administration_policy_set_example" {
  value = ciscoise_device_administration_policy_set.example
}

resource "ciscoise_device_administration_authentication_rules" "example" {
  provider = ciscoise
  parameters {
    policy_id            = ciscoise_device_administration_policy_set.example.item[0].id
    identity_source_name = "Internal Users"
    if_auth_fail         = "REJECT"
    if_process_fail      = "DROP"
    if_user_not_found    = "REJECT"
    rule {
      name    = "Default"
      state   = "enabled"
      default = "true"
    }
  }
}

output "ciscoise_device_administration_authentication_rules_example" {
  value = ciscoise_device_administration_authentication_rules.example
}
