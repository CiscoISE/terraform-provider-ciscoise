terraform {
  required_providers {
    ciscoise = {
      version = "0.6.7-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

data "ciscoise_tacacs_profile" "example" {
  provider = ciscoise
  page     = 1
  size     = 3
}

output "ciscoise_tacacs_profile_example" {
  value = data.ciscoise_tacacs_profile.example.items
}


data "ciscoise_tacacs_profile" "example1" {
  provider = ciscoise
  name     = data.ciscoise_tacacs_profile.example.items[0].name
}

output "ciscoise_tacacs_profile_example1" {
  value = data.ciscoise_tacacs_profile.example1.item_name
}

data "ciscoise_tacacs_profile" "example2" {
  provider = ciscoise
  id       = data.ciscoise_tacacs_profile.example.items[0].id
}

output "ciscoise_tacacs_profile_example2" {
  value = data.ciscoise_tacacs_profile.example2.item_id
}
