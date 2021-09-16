
data "ciscoise_px_grid_node" "example" {
  provider = ciscoise
  name     = "string"
}

output "ciscoise_px_grid_node_example" {
  value = data.ciscoise_px_grid_node.example.item_name
}

data "ciscoise_px_grid_node" "example" {
  provider = ciscoise
  id       = "string"
}

output "ciscoise_px_grid_node_example" {
  value = data.ciscoise_px_grid_node.example.item_id
}

data "ciscoise_px_grid_node" "example" {
  provider = ciscoise
  page     = 1
  size     = 1
}

output "ciscoise_px_grid_node_example" {
  value = data.ciscoise_px_grid_node.example.items
}
