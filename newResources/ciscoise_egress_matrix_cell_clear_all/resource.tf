
resource "ciscoise_egress_matrix_cell_clear_all" "example" {
  provider = ciscoise
  lifecycle {
    create_before_destroy = true
  }
  parameters{
    
  }
}