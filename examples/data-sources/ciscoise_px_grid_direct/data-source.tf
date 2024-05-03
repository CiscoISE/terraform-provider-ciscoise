
data "ciscoise_px_grid_direct" "example" {
  provider = ciscoise
}

output "ciscoise_px_grid_direct_example" {
  value = data.ciscoise_px_grid_direct.example.items
}

data "ciscoise_px_grid_direct" "example" {
  provider       = ciscoise
  connector_name = "string"
}

output "ciscoise_px_grid_direct_example" {
  value = data.ciscoise_px_grid_direct.example.item
}
