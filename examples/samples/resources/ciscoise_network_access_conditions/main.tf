terraform {
  required_providers {
    ciscoise = {
      version = "0.6.1-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

resource "ciscoise_network_access_conditions" "example" {
  provider = ciscoise
  parameters {
    condition_type  = "LibraryConditionAttributes"
    is_negate       = "true"
    name            = "My New Condition"
    description     = "New Test Description"
    dictionary_name = "Radius"
    attribute_name  = "Service-Type"
    operator        = "equals"
    attribute_value = "Call Check"
  }
}

output "ciscoise_network_access_conditions_example" {
  value = ciscoise_network_access_conditions.example
}

data "ciscoise_network_access_conditions" "found" {
  depends_on = [
    ciscoise_network_access_conditions.example
  ]
  provider = ciscoise
  name     = "My New Condition"
}

output "ciscoise_network_access_conditions_found" {
  value = data.ciscoise_network_access_conditions.found.item_name
}
