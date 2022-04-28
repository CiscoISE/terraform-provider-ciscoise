resource "ciscoise_egress_matrix_cell_set_all_status" "example" {
  provider = ciscoise
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    status   = "string"
  }
}