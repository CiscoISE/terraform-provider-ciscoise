terraform {
  required_providers {
    ciscoise = {
      version = "0.6.4-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

data "ciscoise_allowed_protocols" "response" {
  provider = ciscoise
  page     = 1
  size     = 1
}
output "ciscoise__allowed_protocols_response" {
  value = data.ciscoise_allowed_protocols.response
}

data "ciscoise_allowed_protocols" "single_response" {
  provider = ciscoise
  id       = data.ciscoise_allowed_protocols.response.items[0].id
}

output "ciscoise__allowed_protocols_single_response" {
  value = data.ciscoise_allowed_protocols.single_response.item_id
}
