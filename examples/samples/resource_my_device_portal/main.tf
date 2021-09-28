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

resource "ciscoise_my_device_portal" "example" {
  provider = ciscoise
  item {
    name        = "Devices Portal"
    description = "Test portal used by developers to register and manage their personal devices"
    portal_type = "MYDEVICE"
    # settings {
    #   login_page_settings {
    #     social_configs = ["{\"socialMediaType\": \"\", \"socialMediaValue\": \"\"}"]
    #   }
    # }
  }
}

output "ciscoise_my_device_portal_example" {
  value = ciscoise_my_device_portal.example.id
}

data "ciscoise_my_device_portal" "item" {
  provider = ciscoise
  id       = ciscoise_my_device_portal.example.item[0].id
}

output "ciscoise_my_device_portal_item" {
  value = data.ciscoise_my_device_portal.item.item
}