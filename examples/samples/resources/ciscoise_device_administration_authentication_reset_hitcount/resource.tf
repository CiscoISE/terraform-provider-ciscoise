
terraform {
  required_providers {
    ciscoise = {
      version = "0.6.6-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}


resource "ciscoise_device_administration_authentication_reset_hitcount" "example" {
  provider = ciscoise
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    policy_id = "1a0e4fb0-c56f-11ec-ba31-463ee0624939"
  }

}