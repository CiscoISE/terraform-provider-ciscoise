terraform {
  required_providers {
    ciscoise = {
      version = "0.6.21-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

resource "ciscoise_internal_user" "example" {
  provider = ciscoise
  parameters {

    # change_password     = "false"
    description     = "string"
    email           = "string"
    enable_password = "Lionel.messi1019"
    enabled         = "false"
    # expiry_date         = "01-01-2025"
    # expiry_date_enabled = "false"
    first_name = "string"
    # id                  = "string"
    identity_groups  = "NewGroup"
    last_name        = "string"
    name             = "string"
    password         = "Lionel.messi1019"
    password_idstore = "Lionel.messi1019"
    custom_attributes = {
      "a" : "A"
    }
  }
}

output "ciscoise_internal_user_example" {
  value     = ciscoise_internal_user.example
  sensitive = true
}