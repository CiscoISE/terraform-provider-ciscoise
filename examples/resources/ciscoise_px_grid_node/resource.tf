
resource "ciscoise_px_grid_node" "example" {
  provider = ciscoise
  parameters {
    name = "string"
  }
}

output "ciscoise_px_grid_node_example" {
  value = ciscoise_px_grid_node.example
}
