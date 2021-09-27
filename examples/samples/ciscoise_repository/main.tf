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


resource "ciscoise_repository" "example" {
  provider = ciscoise
  item {

    enable_pki = "false"
    name       = "Test_Repo"
    # password    = ""
    path        = "/test"
    protocol    = "HTTPS"
    server_name = "example.org"
    # user_name   = ""
  }
}
# output "ciscoise_repository_example" {
#   value = ciscoise_repository.example
# }

# data "ciscoise_repository" "example" {
#   provider = ciscoise
# }

# output "ciscoise_repository_example" {
#   value = data.ciscoise_repository.example.items
# }

# data "ciscoise_repository" "example1" {
#   provider = ciscoise
#   name     = data.ciscoise_repository.example.items[0].name
# }

# output "ciscoise_repository_example1" {
#   value = data.ciscoise_repository.example1.item
# }
