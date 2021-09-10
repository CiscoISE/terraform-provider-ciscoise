terraform {
  required_providers {
    ciscoise = {
      version = "1.0.0"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

data "ciscoise_admin_user" "response" {
  provider = ciscoise
  page     = 1
  size     = 20
}
output "ciscoise__admin_user_response" {
  value = data.ciscoise_admin_user.response
}

data "ciscoise_admin_user" "single_response" {
  provider = ciscoise
  id       = data.ciscoise_admin_user.response.items[0].id
}

output "ciscoise__admin_user_single_response" {
  value = data.ciscoise_admin_user.single_response
}
