terraform {
  required_providers {
    ciscoise = {
      version = "0.6.0-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

resource "ciscoise_active_directory_join_domain_with_all_nodes" "example" {
  provider = ciscoise
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    id = "string"
    additional_data {

      name  = "string"
      value = "string"
    }
  }

}