terraform {
  required_providers {
    ciscoise = {
      version = "0.3.0-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

resource "ciscoise_active_directory_leave_domain" "example" {
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