
resource "ciscoise_egress_matrix_cell" "example" {
  provider = ciscoise
  parameters {

    default_rule       = "string"
    description        = "string"
    destination_sgt_id = "string"
    id                 = "string"
    matrix_cell_status = "string"
    name               = "string"
    sgacls             = ["string"]
    source_sgt_id      = "string"
  }
}

output "ciscoise_egress_matrix_cell_example" {
  value = ciscoise_egress_matrix_cell.example
}