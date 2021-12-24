terraform {
  required_providers {
    ciscoise = {
      version = "0.1.0-rc.1"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

data "ciscoise_guest_type" "example" {
  provider = ciscoise

}

output "ciscoise_guest_type_example" {
  value = data.ciscoise_guest_type.example.items
}

data "ciscoise_guest_type" "example1" {
  provider = ciscoise
  id       = data.ciscoise_guest_type.example.items[0].id
}

output "ciscoise_guest_type_example1" {
  value = data.ciscoise_guest_type.example1.item
}
