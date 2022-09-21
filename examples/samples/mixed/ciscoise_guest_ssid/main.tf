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

data "ciscoise_guest_ssid" "example" {
  provider = ciscoise

}

output "ciscoise_guest_ssid_example" {
  value = data.ciscoise_guest_ssid.example.items
}

resource "ciscoise_guest_ssid" "item1" {
  parameters {
    name = "guest_ise"
  }
}

output "ciscoise_guest_ssid_item1" {
  value = ciscoise_guest_ssid.item1
}

data "ciscoise_guest_ssid" "example2" {
  provider = ciscoise
  id       = ciscoise_guest_ssid.item1.item[0].id
}

output "ciscoise_guest_ssid_example2" {
  value = data.ciscoise_guest_ssid.example2.item
}
