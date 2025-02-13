

terraform {
  required_providers {
    ciscoise = {
      version = "0.8.2-beta "
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}
resource "ciscoise_ise_root_ca_regenerate" "example" {
  provider = ciscoise
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    remove_existing_ise_intermediate_csr = "true"
  }
}