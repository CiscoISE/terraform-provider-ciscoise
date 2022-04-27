terraform {
  required_providers {
    ciscoise = {
      version = "0.3.0-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

resource "ciscoise_network_access_global_exception_rules_reset_hitcount" "example" {
  provider = ciscoise
  lifecycle {
    create_before_destroy = true
  }
  parameters {}
}