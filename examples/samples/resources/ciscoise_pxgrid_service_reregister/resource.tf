terraform {
  required_providers {
    ciscoise = {
      version = "0.8.2-beta "
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

resource "ciscoise_pxgrid_service_reregister" "example" {
  provider = ciscoise

  lifecycle {
    create_before_destroy = true
  }
  parameters {

  }
}