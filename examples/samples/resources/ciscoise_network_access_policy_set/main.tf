terraform {
  required_providers {
    ciscoise = {
      version = "0.5.0-beta"
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
      id             = "898c705e-d80c-4c88-bc2a-f0900488e430"
      is_negate      = "false"
      name           = "My New Condition"
    }
    default      = "false"
    description  = "New Policy Set test 1"
    hit_counts   = 0
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
