terraform {
  required_providers {
    ciscoise = {
      version = "0.6.8-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

data "ciscoise_sponsor_portal" "example" {
  provider = ciscoise

}

output "ciscoise_sponsor_portal_example" {
  value = data.ciscoise_sponsor_portal.example.items
}

data "ciscoise_sponsor_portal" "example1" {
  provider = ciscoise
  id       = data.ciscoise_sponsor_portal.example.items[0].id
}

output "ciscoise_sponsor_portal_example1" {
  value = data.ciscoise_sponsor_portal.example1.item
}
