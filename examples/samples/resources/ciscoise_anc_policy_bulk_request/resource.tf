terraform {
  required_providers {
    ciscoise = {
      version = "0.6.18-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

resource "ciscoise_anc_policy_bulk_request" "example" {
  provider = ciscoise
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    operation_type      = "string"
    resource_media_type = "string"
  }
}