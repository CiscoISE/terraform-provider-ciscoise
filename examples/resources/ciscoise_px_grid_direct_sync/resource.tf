
resource "ciscoise_px_grid_direct_sync" "example" {
  provider = ciscoise
  parameters {

    sync_type      = "string"
    connector_name = "string"
    description    = "string"
  }
}

output "ciscoise_px_grid_direct_sync_example" {
  value = ciscoise_px_grid_direct_sync.example
}