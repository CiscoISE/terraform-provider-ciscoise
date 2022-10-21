terraform {
  required_providers {
    ciscoise = {
      version = "0.6.10-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

data "ciscoise_sgt" "example" {
  provider = ciscoise

}

output "ciscoise_sgt_example" {
  value = data.ciscoise_sgt.example.items
}

data "ciscoise_sgt" "example1" {
  provider = ciscoise
  id       = data.ciscoise_sgt.example.items[0].id
}

output "ciscoise_sgt_example1" {
  value = data.ciscoise_sgt.example1.item
}
