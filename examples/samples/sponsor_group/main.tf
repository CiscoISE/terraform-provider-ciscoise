terraform {
  required_providers {
    ciscoise = {
      version = "0.1.0-rc.2"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

data "ciscoise_sponsor_group" "example" {
  provider = ciscoise

}

output "ciscoise_sponsor_group_example" {
  value = data.ciscoise_sponsor_group.example.items
}

data "ciscoise_sponsor_group" "example1" {
  provider = ciscoise
  id       = data.ciscoise_sponsor_group.example.items[0].id
}

output "ciscoise_sponsor_group_example1" {
  value = data.ciscoise_sponsor_group.example1.item
}
