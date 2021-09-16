
data "ciscoise_node_group" "example" {
  provider = ciscoise
}

output "ciscoise_node_group_example" {
  value = data.ciscoise_node_group.example.items
}

data "ciscoise_node_group" "example" {
  provider        = ciscoise
  node_group_name = "string"
}

output "ciscoise_node_group_example" {
  value = data.ciscoise_node_group.example.item
}
