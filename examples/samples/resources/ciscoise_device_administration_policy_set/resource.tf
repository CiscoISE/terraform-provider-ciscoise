terraform {
  required_providers {
    ciscoise = {
      version = "0.6.0-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
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