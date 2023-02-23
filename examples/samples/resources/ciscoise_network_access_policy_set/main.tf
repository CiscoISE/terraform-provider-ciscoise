terraform {
  required_providers {
    ciscoise = {
      version = "0.6.17-beta"
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

resource "ciscoise_network_access_policy_set" "wired-mm-test" {
  provider = ciscoise
  parameters {

    condition {
      condition_type = "ConditionAndBlock"
      is_negate      = "false"
      children {
        dictionary_name = "Radius"
        attribute_name  = "NAS-Port-Type"
        operator        = "equals"
        attribute_value = "Ethernet"
      }
      children {
        dictionary_name = "DEVICE"
        attribute_name  = "Deployment Stage"
        operator        = "equals"
        attribute_value = "Deployment Stage#Monitor Mode"
      }
    }
    default      = "false"
    description  = "Wired Monitor Mode TEST"
    is_proxy     = "false"
    name         = "Wired_MM_TEST"
    rank         = 0
    service_name = "Default Network Access"
    state        = "enabled"
  }
}


# output "ciscoise_network_access_policy_set_example" {
#   value = ciscoise_network_access_policy_set.example
# }
