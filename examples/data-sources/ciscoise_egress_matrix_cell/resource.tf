
data "ciscoise_egress_matrix_cell" "example" {
  provider    = ciscoise
  filter      = ["string"]
  filter_type = "string"
  page        = 1
  size        = 1
  sortasc     = "string"
  sortdsc     = "string"
}

output "ciscoise_egress_matrix_cell_example" {
  value = data.ciscoise_egress_matrix_cell.example.items
}

data "ciscoise_egress_matrix_cell" "example" {
  provider = ciscoise
  id       = "string"
}

output "ciscoise_egress_matrix_cell_example" {
  value = data.ciscoise_egress_matrix_cell.example.item
}
