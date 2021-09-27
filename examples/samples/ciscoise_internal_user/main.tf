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


data "ciscoise_internal_user" "example" {
  provider = ciscoise

}

output "ciscoise_internal_user_example" {
  value = data.ciscoise_internal_user.example.items
}

data "ciscoise_internal_user" "example1" {
  provider = ciscoise
  name     = data.ciscoise_internal_user.example.items[0].name
}

output "ciscoise_internal_user_example1" {
  value = data.ciscoise_internal_user.example1.item_name
}

data "ciscoise_internal_user" "example2" {
  provider = ciscoise
  id       = data.ciscoise_internal_user.example.items[0].id
}

output "ciscoise_internal_user_example2" {
  value = data.ciscoise_internal_user.example2.item_id
}