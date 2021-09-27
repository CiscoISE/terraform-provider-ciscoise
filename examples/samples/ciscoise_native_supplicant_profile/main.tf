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

data "ciscoise_native_supplicant_profile" "example" {
  provider = ciscoise
  page     = 1
  size     = 1
}

output "ciscoise_native_supplicant_profile_example" {
  value = data.ciscoise_native_supplicant_profile.example.items
}

data "ciscoise_native_supplicant_profile" "example1" {
  provider = ciscoise
  id       = data.ciscoise_native_supplicant_profile.example.items[0].id
}

output "ciscoise_native_supplicant_profile_example1" {
  value = data.ciscoise_native_supplicant_profile.example1.item
}
