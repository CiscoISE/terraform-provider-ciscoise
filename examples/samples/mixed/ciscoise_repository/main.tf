terraform {
  required_providers {
    ciscoise = {
      version = "0.6.1-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}


resource "ciscoise_repository" "example" {
  provider = ciscoise
  parameters {

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

data "ciscoise_repository" "items" {
  provider = ciscoise
  depends_on = [
    ciscoise_repository.example
  ]
}

output "ciscoise_repository_items" {
  value = data.ciscoise_repository.items.items
}

data "ciscoise_repository" "item1" {
  provider = ciscoise
  depends_on = [
    data.ciscoise_repository.items
  ]
  repository_name = data.ciscoise_repository.items.items[0].name
}

output "ciscoise_repository_item1" {
  value = data.ciscoise_repository.item1.item
}
