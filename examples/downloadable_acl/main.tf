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

data "ciscoise_downloadable_acl" "response" {
  provider = ciscoise
  page     = 1
  size     = 20
}
output "ciscoise__downloadable_acl_response" {
  value = data.ciscoise_downloadable_acl.response
}


data "ciscoise_downloadable_acl" "single_response" {
  provider = ciscoise
  id       = data.ciscoise_downloadable_acl.response.items[0].id
}

output "ciscoise__downloadable_acl_single_response" {
  value = data.ciscoise_downloadable_acl.single_response
}
