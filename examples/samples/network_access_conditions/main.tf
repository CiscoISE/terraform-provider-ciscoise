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



data "ciscoise_network_access_conditions" "example" {
  provider = ciscoise
}

output "ciscoise_network_access_conditions_example" {
  value = data.ciscoise_network_access_conditions.example.items
}

data "ciscoise_network_access_conditions" "example1" {
  provider = ciscoise
  name     = data.ciscoise_network_access_conditions.example.items[0].name
}

output "ciscoise_network_access_conditions_example1" {
  value = data.ciscoise_network_access_conditions.example1.item_name
}

data "ciscoise_network_access_conditions" "example2" {
  provider = ciscoise
  id       = data.ciscoise_network_access_conditions.example.items[0].id
}

output "ciscoise_network_access_conditions_example2" {
  value = data.ciscoise_network_access_conditions.example2.item_id
}