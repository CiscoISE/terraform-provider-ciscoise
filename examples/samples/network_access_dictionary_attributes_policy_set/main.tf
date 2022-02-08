terraform {
  required_providers {
    ciscoise = {
      version = "0.2.0"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

data "ciscoise_network_access_dictionary_attributes_policy_set" "example" {
  provider = ciscoise
}

output "ciscoise_network_access_dictionary_attributes_policy_set_example" {
  value = data.ciscoise_network_access_dictionary_attributes_policy_set.example.items
}
