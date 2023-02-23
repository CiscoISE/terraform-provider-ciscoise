terraform {
  required_providers {
    ciscoise = {
      version = "0.6.17-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

resource "ciscoise_device_administration_global_exception_rules_reset_hitcount" "example" {
  provider = ciscoise
  lifecycle {
    create_before_destroy = true
  }

  parameters {

  }
}