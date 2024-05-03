
resource "ciscoise_trustsec_vn_bulk_delete" "example" {
  provider   = ciscoise
  parameters = ["string"]
}

output "ciscoise_trustsec_vn_bulk_delete_example" {
  value = ciscoise_trustsec_vn_bulk_delete.example
}