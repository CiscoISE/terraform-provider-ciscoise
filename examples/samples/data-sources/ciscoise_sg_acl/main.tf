terraform {
  required_providers {
    ciscoise = {
      version = "0.6.18-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

data "ciscoise_sg_acl" "example" {
  provider = ciscoise

}

output "ciscoise_sg_acl_example" {
  value = data.ciscoise_sg_acl.example.items
}

data "ciscoise_sg_acl" "example1" {
  provider = ciscoise
  id       = data.ciscoise_sg_acl.example.items[0].id
}

output "ciscoise_sg_acl_example1" {
  value = data.ciscoise_sg_acl.example1.item
}
