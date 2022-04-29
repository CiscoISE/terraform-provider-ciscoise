
resource "ciscoise_egress_matrix_cell_bulk_request" "example" {
  provider = ciscoise
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    operation_type      = "string"
    resource_media_type = "string"
  }
}