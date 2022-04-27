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

resource "ciscoise_network_access_authentication_rules_reset_hitcount" "example" {
  provider  = ciscoise
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    policy_id = "70836740-8bff-11e6-996c-525400b48521"
  } 
}