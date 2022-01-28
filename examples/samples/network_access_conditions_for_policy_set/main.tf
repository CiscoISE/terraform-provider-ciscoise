terraform {
  required_providers {
    ciscoise = {
      version = "0.1.0"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

data "ciscoise_network_access_conditions_for_policy_set" "example" {
  provider = ciscoise
}

output "ciscoise_network_access_conditions_for_policy_set_example" {
  value = data.ciscoise_network_access_conditions_for_policy_set.example.items
}
