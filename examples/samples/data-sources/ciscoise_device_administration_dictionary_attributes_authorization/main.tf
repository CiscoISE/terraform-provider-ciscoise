terraform {
  required_providers {
    ciscoise = {
      version = "0.8.1-beta "
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

data "ciscoise_device_administration_dictionary_attributes_authorization" "example" {
  provider = ciscoise
}

output "ciscoise_device_administration_dictionary_attributes_authorization_example" {
  value = data.ciscoise_device_administration_dictionary_attributes_authorization.example.items
}
