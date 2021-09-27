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

# data "ciscoise_my_device_portal" "items" {
#   provider = ciscoise
#   id       = "5e97b665-385f-4f3e-a9f0-721664ea7f06"
# }

# output "ciscoise_my_device_portal_items" {
#   value = data.ciscoise_my_device_portal.items.item
# }