
resource "ciscoise_trustsec_vn_bulk_update" "example" {
  provider = ciscoise
  parameters {

    additional_attributes = "string"
    id                    = "string"
    last_update           = "string"
    name                  = "string"
  }
}

output "ciscoise_trustsec_vn_bulk_update_example" {
  value = ciscoise_trustsec_vn_bulk_update.example
}