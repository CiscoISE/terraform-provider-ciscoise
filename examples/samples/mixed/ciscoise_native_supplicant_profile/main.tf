terraform {
  required_providers {
    ciscoise = {
      version = "0.6.1-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

data "ciscoise_native_supplicant_profile" "items" {
  provider = ciscoise
}

output "ciscoise_native_supplicant_profile_items" {
  value = data.ciscoise_native_supplicant_profile.items
}

data "ciscoise_native_supplicant_profile" "item" {
  provider = ciscoise
  id       = data.ciscoise_native_supplicant_profile.items.items[0].id
}

output "ciscoise_native_supplicant_profile_item" {
  value = data.ciscoise_native_supplicant_profile.item.item
}

resource "ciscoise_native_supplicant_profile" "example" {
  provider = ciscoise
  parameters {
    id          = "67a6ca50-edc9-4236-ada4-225559ed54d6"
    name        = "Cisco-ISE-Chrome-NSP"
    description = "Pre-configured Ncd For Chrome OS"
    wireless_profiles {
      ssid                    = "ChromeDummySSID"
      previous_ssid           = "ChromeDummySSID"
      allowed_protocol        = "TLS"
      certificate_template_id = "0ca8f1b6-500d-560b-e053-75189a0ab0d1"
      action_type             = "UPDATE"
    }
  }
}


output "ciscoise_native_supplicant_profile_example" {
  value = ciscoise_native_supplicant_profile.example
}
