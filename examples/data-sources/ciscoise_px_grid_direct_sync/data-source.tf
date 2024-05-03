
data "ciscoise_px_grid_direct_sync" "example" {
  provider       = ciscoise
  connector_name = "string"
}

output "ciscoise_px_grid_direct_sync_example" {
  value = data.ciscoise_px_grid_direct_sync.example.item
}
