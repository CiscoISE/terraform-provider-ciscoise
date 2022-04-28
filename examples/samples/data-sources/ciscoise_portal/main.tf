terraform {
  required_providers {
    ciscoise = {
      version = "0.4.0-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

data "ciscoise_portal" "example" {
  provider = ciscoise

}

output "ciscoise_portal_example" {
  value = data.ciscoise_portal.example.items
}

data "ciscoise_portal" "single_response" {
  provider = ciscoise
  id       = data.ciscoise_portal.example.items[0].id
}

output "ciscoise_portal_single_response" {
  value = data.ciscoise_portal.single_response.item
}
