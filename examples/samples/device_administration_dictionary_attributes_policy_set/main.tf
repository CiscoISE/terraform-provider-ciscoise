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

data "ciscoise_device_administration_dictionary_attributes_policy_set" "example" {
  provider = ciscoise
}

output "ciscoise_device_administration_dictionary_attributes_policy_set_example" {
  value = data.ciscoise_device_administration_dictionary_attributes_policy_set.example.items
}
