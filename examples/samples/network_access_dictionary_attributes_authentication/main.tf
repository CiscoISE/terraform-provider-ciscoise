terraform {
  required_providers {
    ciscoise = {
      version = "0.0.2-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

data "ciscoise_network_access_dictionary_attributes_authentication" "example" {
  provider = ciscoise
}

output "ciscoise_network_access_dictionary_attributes_authentication_example" {
  value = data.ciscoise_network_access_dictionary_attributes_authentication.example.items
}
