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

resource "ciscoise_network_access_conditions" "conditions" {
  provider = ciscoise

  parameters {
    name           = "condtion_11"
    condition_type = "LibraryConditionAndBlock"
    is_negate      = false

    children {
      condition_type = "ConditionReference"
      #id             = "ff6008e0-5c35-48a3-9fab-e0e709983369"
    }
    children {
      condition_type = "ConditionReference"
      #id             = "2f1a717d-cc09-41e3-a2ac-ecea19a841f8"
    }

  }
}

# output "ciscoise_network_access_conditions_example" {
#   value = ciscoise_network_access_conditions.example
# }

# data "ciscoise_network_access_conditions" "found" {
#   depends_on = [
#     ciscoise_network_access_conditions.example
#   ]
#   provider = ciscoise
#   name     = "My New Condition"
# }

# output "ciscoise_network_access_conditions_found" {
#   value = data.ciscoise_network_access_conditions.found.item_name
# }
