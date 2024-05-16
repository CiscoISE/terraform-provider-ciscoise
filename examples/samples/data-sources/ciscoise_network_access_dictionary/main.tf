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

data "ciscoise_network_access_dictionary" "example" {
  provider = ciscoise
}

output "ciscoise_network_access_dictionary_example" {
  value = data.ciscoise_network_access_dictionary.example.items
}

data "ciscoise_network_access_dictionary" "example1" {
  provider = ciscoise
  name     = data.ciscoise_network_access_dictionary.example.items[0].name
}

output "ciscoise_network_access_dictionary_example1" {
  value = data.ciscoise_network_access_dictionary.example1.item
}
