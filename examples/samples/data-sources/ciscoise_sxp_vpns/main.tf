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

data "ciscoise_sxp_vpns" "example" {
  provider = ciscoise

}

output "ciscoise_sxp_vpns_example" {
  value = data.ciscoise_sxp_vpns.example.items
}

data "ciscoise_sxp_vpns" "example1" {
  provider = ciscoise
  id       = data.ciscoise_sxp_vpns.example.items[0].id
}

output "ciscoise_sxp_vpns_example1" {
  value = data.ciscoise_sxp_vpns.example1.item
}

resource "ciscoise_sxp_vpns" "name" {
  provider = ciscoise
  parameters {
    sxp_vpn_name = "testing"
  }
}
