terraform {
  required_providers {
    ciscoise = {
      version = "0.8.0-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

resource "ciscoise_device_administration_local_exception_rules_reset_hitcount" "example" {
  provider = ciscoise
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    policy_id = "cb32c3bc-c720-40c3-83e4-8897f9dd6943"
  }

}