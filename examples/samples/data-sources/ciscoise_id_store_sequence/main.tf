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


data "ciscoise_id_store_sequence" "example" {
  provider = ciscoise
  page     = 1
  size     = 1
}

output "ciscoise_id_store_sequence_example" {
  value = data.ciscoise_id_store_sequence.example.items
}


data "ciscoise_id_store_sequence" "example1" {
  provider = ciscoise
  name     = data.ciscoise_id_store_sequence.example.items[0].name
}

output "ciscoise_id_store_sequence_example1" {
  value = data.ciscoise_id_store_sequence.example1.item_name
}

data "ciscoise_id_store_sequence" "example2" {
  provider = ciscoise
  id       = data.ciscoise_id_store_sequence.example.items[0].id
}

output "ciscoise_id_store_sequence_example2" {
  value = data.ciscoise_id_store_sequence.example2.item_id
}
