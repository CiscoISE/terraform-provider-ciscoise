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

data "ciscoise_identity_group" "example" {
  provider = ciscoise

}

output "ciscoise_identity_group_example" {
  value = data.ciscoise_identity_group.example.items
}

# GetIdentityGroupByName's API call fails with 404
# data "ciscoise_identity_group" "example1" {
#   provider = ciscoise
#   name     = data.ciscoise_identity_group.example.items[1].name
# }

# output "ciscoise_identity_group_example1" {
#   value = data.ciscoise_identity_group.example1.item_name
# }

data "ciscoise_identity_group" "example2" {
  provider = ciscoise
  id       = data.ciscoise_identity_group.example.items[0].id
}

output "ciscoise_identity_group_example2" {
  value = data.ciscoise_identity_group.example2.item_id
}
