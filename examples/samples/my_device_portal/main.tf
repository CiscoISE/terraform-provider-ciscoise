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

data "ciscoise_my_device_portal" "example" {
  provider = ciscoise

}

output "ciscoise_my_device_portal_example" {
  value = data.ciscoise_my_device_portal.example.items
}

data "ciscoise_my_device_portal" "example1" {
  provider = ciscoise
  id       = data.ciscoise_my_device_portal.example.items[0].id
}

output "ciscoise_my_device_portal_example1" {
  value = data.ciscoise_my_device_portal.example1.item
}
