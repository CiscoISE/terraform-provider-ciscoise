terraform {
  required_providers {
    ciscoise = {
      version = "0.6.22-beta"
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

resource "ciscoise_network_access_policy_set" "policy_sets" {
  provider = ciscoise

  parameters {
    description  = "Test Policy"
    is_proxy     = false
    name         = "New Policy Set 6"
    service_name = "Default Network Access"
    state        = "enabled"
    condition {
      condition_type = "ConditionAndBlock"
      is_negate      = false

      # children {
      #     condition_type = "ConditionAttributes"
      #     is_negate =  false
      #     dictionary_name =  "DEVICE"
      #     attribute_name = "Device Type"
      #     operator = "equals"
      #     dictionary_value =  null
      #     attribute_value = "All Device Types"
      # }
      children {
        id             = "ff6008e0-5c35-48a3-9fab-e0e709983369"
        condition_type = "ConditionReference"
      }
    }
  }
}


# output "ciscoise_network_access_policy_set_example" {
#   value = ciscoise_network_access_policy_set.example
# }
