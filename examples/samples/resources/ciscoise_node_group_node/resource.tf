terraform {
  required_providers {
    ciscoise = {
      version = "0.2.0"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}


resource "ciscoise_node_group_node" "example" {
  provider = ciscoise
  parameters {
    node_group_name = "isegroup"
    hostname        = "ise"
  }
}

output "ciscoise_node_group_node_example" {
  value = ciscoise_node_group_node.example.item
}
