terraform {
  required_providers {
    ciscoise = {
      version = "0.6.7-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

resource "ciscoise_px_grid_settings_auto_approve" "example" {
  provider = ciscoise

  lifecycle {
    create_before_destroy = true
  }
  parameters {
    allow_password_based_accounts    = "false"
    auto_approve_cert_based_accounts = "false"
  }


}