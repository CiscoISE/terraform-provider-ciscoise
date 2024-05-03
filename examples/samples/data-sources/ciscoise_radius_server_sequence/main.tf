terraform {
  required_providers {
    ciscoise = {
      version = "0.8.0-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

data "ciscoise_radius_server_sequence" "example" {
  provider = ciscoise
  page     = 1
  size     = 1
}

output "ciscoise_radius_server_sequence_example" {
  value = data.ciscoise_radius_server_sequence.example.items
}

data "ciscoise_radius_server_sequence" "example1" {
  provider = ciscoise
  id       = data.ciscoise_radius_server_sequence.example.items[0].id
}

output "ciscoise_radius_server_sequence_example1" {
  value = data.ciscoise_radius_server_sequence.example1.item
}
