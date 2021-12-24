terraform {
  required_providers {
    ciscoise = {
      version = "0.1.0-rc.1"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

data "ciscoise_network_access_dictionary_attributes_authorization" "example" {
  provider = ciscoise
}

output "ciscoise_network_access_dictionary_attributes_authorization_example" {
  value = data.ciscoise_network_access_dictionary_attributes_authorization.example.items
}
