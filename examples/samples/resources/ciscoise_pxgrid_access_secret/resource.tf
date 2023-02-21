terraform {
  required_providers {
    ciscoise = {
      version = "0.6.16-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

resource "ciscoise_pxgrid_access_secret" "example" {
  provider = ciscoise
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    peer_node_name = "string"
  }
}