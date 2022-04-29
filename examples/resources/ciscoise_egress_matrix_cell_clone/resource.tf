resource "ciscoise_egress_matrix_cell_clone" "example" {
  provider = ciscoise
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    dst_sgt_id = "string"
    id         = "string"
    src_sgt_id = "string"
  }
}