terraform {
  required_providers {
    ciscoise = {
      version = "0.6.10-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

data "ciscoise_network_access_conditions_for_policy_set" "items" {
  provider = ciscoise
}
output "ciscoise_network_access_conditions_for_policy_set_items" {
  value = data.ciscoise_network_access_conditions_for_policy_set.items.items
}

resource "ciscoise_network_access_policy_set" "example" {
  provider = ciscoise
  parameters {
    condition {
      condition_type = "ConditionReference"
      is_negate      = "false"
      name           = "Wired_MAB"
      id             = "9aab0da7-e3e3-4cd7-81c2-18c3ebbe6a96"
      description    = "A condition to match MAC Authentication Bypass service based authentication requests from switches, according to the corresponding MAB attributes defined in the device profile."
    }
    description  = "New Policy Set test 1"
    is_proxy     = "false"
    name         = "New Policy Set 1"
    rank         = 0
    service_name = "Default Network Access"
    state        = "disabled"
  }
}

output "ciscoise_network_access_policy_set_example" {
  value = ciscoise_network_access_policy_set.example
}
