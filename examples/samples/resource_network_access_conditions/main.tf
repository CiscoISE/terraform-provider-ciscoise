terraform {
  required_providers {
    ciscoise = {
      version = "0.0.1-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

# resource "ciscoise_network_access_conditions" "example" {
#   provider = ciscoise
#   item {
#       condition_type = "LibraryConditionAttributes"
#       is_negate = "false"
#       name = "My New Condition"
#       description = "New optional Description"
#       dictionary_name = "Radius"
#       attribute_name = "Service-Type"
#       operator = "equals"
#       attribute_value = "Call Check"
#   }
# }

# output "ciscoise_network_access_conditions_example" {
#     value = ciscoise_network_access_conditions.example
# }

data "ciscoise_network_access_conditions" "found" {
  provider = ciscoise
  # id = "37a5b141-f00a-4793-bcc6-805fa4e2c427"
  # name = "My New Condition"
}

output "ciscoise_network_access_conditions_found" {
  value = data.ciscoise_network_access_conditions.found
}