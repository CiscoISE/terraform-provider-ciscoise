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

data "ciscoise_support_bundle_status" "items" {
  provider = ciscoise
  page     = 1
  size     = 1
}

output "ciscoise_support_bundle_status_items" {
  value = data.ciscoise_support_bundle_status.items.items
}

data "ciscoise_support_bundle_status" "item" {
  provider = ciscoise
  id       = data.ciscoise_support_bundle_status.items.items[0].id
}

output "ciscoise_support_bundle_status_item" {
  value = data.ciscoise_support_bundle_status.item.item
}
