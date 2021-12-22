terraform {
  required_providers {
    ciscoise = {
      version = "0.0.2-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

data "ciscoise_internal_user" "items" {
  provider = ciscoise
}
output "ciscoise_internal_user_items" {
  value = data.ciscoise_internal_user.items.items
}

# data "ciscoise_internal_user" "example1" {
#   provider = ciscoise
#   name     = data.ciscoise_internal_user.items.items[0].name
# }

# output "ciscoise_internal_user_example1" {
#   value = data.ciscoise_internal_user.example1.item_name
# }

# data "ciscoise_internal_user" "example2" {
#   provider = ciscoise
#   id       = data.ciscoise_internal_user.items.items[0].id
# }

# output "ciscoise_internal_user_example2" {
#   value = data.ciscoise_internal_user.example2.item_id
# }

resource "ciscoise_internal_user" "example" {
  provider = ciscoise
  parameters {

    change_password = "true"
    # custom_attributes = {
    #   Created = "1616961914"
    #   Department = "EN"
    #   Expired = "1617566728"
    #   Country = "US"
    # }
    description         = "Recommended attrs to update an account."
    enable_password     = "C1sco1234!3"
    password            = "C1sco1234!3"
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