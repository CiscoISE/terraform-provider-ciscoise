terraform {
  required_providers {
    ciscoise = {
      version = "0.1.0-rc.3"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}


data "ciscoise_network_device_group" "example" {
  provider = ciscoise

}

output "ciscoise_network_device_group_example" {
  value = data.ciscoise_network_device_group.example.items
}

data "ciscoise_network_device_group" "example1" {
  provider = ciscoise
  name     = data.ciscoise_network_device_group.example.items[0].name
}

output "ciscoise_network_device_group_example1" {
  value = data.ciscoise_network_device_group.example1.item_name
}

data "ciscoise_network_device_group" "example2" {
  provider = ciscoise
  id       = data.ciscoise_network_device_group.example.items[0].id
}

output "ciscoise_network_device_group_example2" {
  value = data.ciscoise_network_device_group.example2.item_id
}
