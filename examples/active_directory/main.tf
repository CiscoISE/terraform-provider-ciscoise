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

data "ciscoise_active_directory" "response" {
  provider = ciscoise
  page     = 1
  size     = 20
}
output "ciscoise__active_directory_response" {
  value = data.ciscoise_active_directory.response
}

data "ciscoise_active_directory" "single_response" {
  provider = ciscoise
  id       = data.ciscoise_active_directory.response.items[0].id
}

output "ciscoise__active_directory_single_response" {
  value = data.ciscoise_active_directory.single_response
}
