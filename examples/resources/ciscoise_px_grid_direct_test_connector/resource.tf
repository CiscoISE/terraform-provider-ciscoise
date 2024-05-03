
resource "ciscoise_px_grid_direct_test_connector" "example" {
  provider = ciscoise
  parameters {

    password  = "******"
    user_name = "string"
  }
}

output "ciscoise_px_grid_direct_test_connector_example" {
  value = ciscoise_px_grid_direct_test_connector.example
}