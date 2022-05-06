terraform {
  required_providers {
    ciscoise = {
      version = "0.5.0-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

data "ciscoise_sxp_local_bindings" "example" {
  provider = ciscoise

}

output "ciscoise_sxp_local_bindings_example" {
  value = data.ciscoise_sxp_local_bindings.example.items
}

data "ciscoise_sxp_local_bindings" "example1" {
  provider = ciscoise
  id       = data.ciscoise_sxp_local_bindings.example.items[0].id
}

output "ciscoise_sxp_local_bindings_example1" {
  value = data.ciscoise_sxp_local_bindings.example1.item
}

resource "ciscoise_sxp_local_bindings" "name" {
  provider = ciscoise
  parameters {
    binding_name       = data.ciscoise_sxp_local_bindings.example1.item[0].binding_name
    description        = data.ciscoise_sxp_local_bindings.example1.item[0].description
    id                 = data.ciscoise_sxp_local_bindings.example1.item[0].id
    ip_address_or_host = "10.10.20.1"
    sgt                = data.ciscoise_sxp_local_bindings.example1.item[0].sgt
    sxp_vpn            = data.ciscoise_sxp_local_bindings.example1.item[0].sxp_vpn
    vns                = data.ciscoise_sxp_local_bindings.example1.item[0].vns
  }
}
