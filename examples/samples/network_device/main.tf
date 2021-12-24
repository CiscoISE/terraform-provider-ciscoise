terraform {
  required_providers {
    ciscoise = {
      version = "0.1.0-rc"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

data "ciscoise_network_device" "response" {
  provider = ciscoise
  page     = 1
  size     = 20
}
output "ciscoise_network_device_response" {
  value = data.ciscoise_network_device.response
}

data "ciscoise_network_device" "single_response_id" {
  provider = ciscoise
  id       = data.ciscoise_network_device.response.items[0].id
}

output "ciscoise_network_device_single_response_id" {
  value = data.ciscoise_network_device.single_response_id
}


data "ciscoise_network_device" "single_response_name" {
  provider = ciscoise
  name     = data.ciscoise_network_device.response.items[0].name
}
output "ciscoise_network_device_single_response_name" {
  value = data.ciscoise_network_device.single_response_name
}
