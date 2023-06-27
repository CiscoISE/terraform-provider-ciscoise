terraform {
  required_providers {
    ciscoise = {
      version = "0.6.21-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

data "ciscoise_guest_location" "items" {
  provider = ciscoise

}

output "ciscoise_guest_location_items" {
  value = data.ciscoise_guest_location.items.items
}

# WARNING: It responds with 405 method not allowed
# data "ciscoise_guest_location" "example" {
#   provider = ciscoise
#   id       = data.ciscoise_guest_location.items.items[0].id
# }

# output "ciscoise_guest_location_example" {
#   value = data.ciscoise_guest_location.example.item
# }
