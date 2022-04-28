resource "ciscoise_sxp_vpns_bulk_request" "example" {
  provider = ciscoise
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    operation_type      = "test"
    resource_media_type = "test"
  }
}
