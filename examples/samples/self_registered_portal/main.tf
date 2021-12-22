terraform {
  required_providers {
    ciscoise = {
      version = "0.0.3-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

data "ciscoise_self_registered_portal" "example" {
  provider = ciscoise

}

output "ciscoise_self_registered_portal_example" {
  value = data.ciscoise_self_registered_portal.example.items
}

data "ciscoise_self_registered_portal" "example1" {
  provider = ciscoise
  id       = data.ciscoise_self_registered_portal.example.items[0].id
}

output "ciscoise_self_registered_portal_example1" {
  value = data.ciscoise_self_registered_portal.example1.item
}
