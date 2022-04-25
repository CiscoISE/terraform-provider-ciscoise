
resource "ciscoise_network_device_bulk_request" "example" {
  provider = ciscoise

  operation_type      = "string"
  resource_media_type = "string"
}