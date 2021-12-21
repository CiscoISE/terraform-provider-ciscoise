
data "ciscoise_node_group_node" "example" {
  provider        = ciscoise
  node_group_name = "string"
}

output "ciscoise_node_group_node_example" {
  value = data.ciscoise_node_group_node.example.items
}
