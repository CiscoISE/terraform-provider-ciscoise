resource "ciscoise_node_group_node" "example" {
  provider = ciscoise
  parameters {
    node_group_name = "string"
    hostname        = "string"
  }
}

output "ciscoise_node_group_node_example" {
  value = ciscoise_node_group_node.example.item
}
