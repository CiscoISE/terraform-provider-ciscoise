terraform {
  required_providers {
    ciscoise = {
      version = "0.0.1-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

resource "ciscoise_internal_user" "example" {
    provider = ciscoise
    item {
  
      change_password = "false"
    #   custom_attributes = {
    #     Created = "1616961914"
    #     Department = "EN"
    #     Expired = "1617566728"
    #     Country = "US"
    #   }
      description         = "Recommended attributes to update an account."
      enable_password     = "C1sco1234!2"
      password            = "C1sco1234!2"
      enabled             = "true"
      name                = "thomas"
      expiry_date_enabled = "false"
    }
}

data "ciscoise_internal_user" "response" {
    depends_on = [
      ciscoise_internal_user.example
    ]
}

output "ciscoise_internal_user_response" {
  value = data.ciscoise_internal_user.response
}