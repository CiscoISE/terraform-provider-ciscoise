terraform {
  required_providers {
    ciscoise = {
      version = "0.1.0-rc.4"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

resource "ciscoise_sxp_connections" "name" {
  provider = ciscoise
  parameters {
    sxp_peer = "Test2"
    # id          = "string"
    description = "Test2"
    enabled     = "false"
    ip_address  = "11.30.10.1"
    sxp_mode    = "LISTENER"
    sxp_node    = "ise"
    sxp_version = "VERSION_4"
    sxp_vpn     = "default"
  }
}

output "ciscoise_sxp_connections_name" {
  value = ciscoise_sxp_connections.name
}

data "ciscoise_sxp_connections" "example" {
  provider = ciscoise
}

output "ciscoise_sxp_connections_example" {
  value = data.ciscoise_sxp_connections.example.items
}

data "ciscoise_sxp_connections" "example1" {
  provider = ciscoise
  id       = data.ciscoise_sxp_connections.example.items[0].id
}

output "ciscoise_sxp_connections_example1" {
  value = data.ciscoise_sxp_connections.example1.item
}
