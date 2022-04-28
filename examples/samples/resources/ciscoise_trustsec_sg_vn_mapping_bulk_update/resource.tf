terraform {
  required_providers {
    ciscoise = {
      version = "0.4.0-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

resource "ciscoise_trustsec_sg_vn_mapping_bulk_update" "example" {
  provider = ciscoise

  lifecycle {
    create_before_destroy = true
  }
  parameters {
    payload {
      id          = "string"
      last_update = "string"
      sg_name     = "a"
      sgt_id      = "string"
      vn_id       = "string"
      vn_name     = "string"
    }
  }
}