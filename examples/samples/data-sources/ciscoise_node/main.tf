terraform {
  required_providers {
    ciscoise = {
      version = "0.5.0-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

data "ciscoise_node" "example" {
  provider = ciscoise
  page     = 1
  size     = 1
}

output "ciscoise_node_example" {
  value = data.ciscoise_node.example.items
}

data "ciscoise_node" "item_name" {
  provider = ciscoise
  name     = data.ciscoise_node.example.items[0].name
}

output "ciscoise_node_item_name" {
  value = data.ciscoise_node.item_name.item_name
}

data "ciscoise_node" "item_id" {
  provider = ciscoise
  id       = data.ciscoise_node.example.items[0].id
}

output "ciscoise_node_item_id" {
  value = data.ciscoise_node.item_id.item_id
}
