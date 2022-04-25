
resource "ciscoise_anc_policy_bulk_request" "example" {
  provider = ciscoise

  operation_type      = "string"
  resource_media_type = "string"
}