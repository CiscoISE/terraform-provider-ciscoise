terraform {
  required_providers {
    ciscoise = {
      version = "0.7.0-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

data "ciscoise_authorization_profile" "response" {
  provider = ciscoise
  page     = 1
  size     = 20
}
output "ciscoise_authorization_profile_response" {
  value = data.ciscoise_authorization_profile.response
}

data "ciscoise_authorization_profile" "single_response_id" {
  provider = ciscoise
  id       = data.ciscoise_authorization_profile.response.items[0].id
}

output "ciscoise_authorization_profile_single_response_id" {
  value = data.ciscoise_authorization_profile.single_response_id
}


data "ciscoise_authorization_profile" "single_response_name" {
  provider = ciscoise
  name     = data.ciscoise_authorization_profile.single_response_id.item_id[0].name
}
output "ciscoise_authorization_profile_single_response_name" {
  value = data.ciscoise_authorization_profile.single_response_name
}
