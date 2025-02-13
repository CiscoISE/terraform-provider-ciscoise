terraform {
  required_providers {
    ciscoise = {
      version = "0.8.2-beta "
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

data "ciscoise_hotspot_portal" "example" {
  provider = ciscoise

}

output "ciscoise_hotspot_portal_example" {
  value = data.ciscoise_hotspot_portal.example.items
}

data "ciscoise_hotspot_portal" "example1" {
  provider = ciscoise
  id       = data.ciscoise_hotspot_portal.example.items[0].id
}

output "ciscoise_hotspot_portal_example1" {
  value = data.ciscoise_hotspot_portal.example1.item
}
