
data "ciscoise_node" "example" {
  provider = ciscoise
  name     = "string"
}

output "ciscoise_node_example" {
  value = data.ciscoise_node.example.item_name
}

data "ciscoise_node" "example" {
  provider = ciscoise
  id       = "string"
}

output "ciscoise_node_example" {
  value = data.ciscoise_node.example.item_id
}

data "ciscoise_node" "example" {
  provider    = ciscoise
  filter      = ["string"]
  filter_type = "string"
  page        = 1
  size        = 1
}

output "ciscoise_node_example" {
  value = data.ciscoise_node.example.items
}
