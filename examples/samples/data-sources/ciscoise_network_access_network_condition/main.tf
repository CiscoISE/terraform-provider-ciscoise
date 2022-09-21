terraform {
  required_providers {
    ciscoise = {
      version = "0.6.6-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

data "ciscoise_network_access_network_condition" "example" {
  provider = ciscoise
}

output "ciscoise_network_access_network_condition_example" {
  value = data.ciscoise_network_access_network_condition.example.items
}

data "ciscoise_network_access_network_condition" "example1" {
  provider = ciscoise
  id       = data.ciscoise_network_access_network_condition.example.items[0].id
}

output "ciscoise_network_access_network_condition_example1" {
  value = data.ciscoise_network_access_network_condition.example1.item
}
