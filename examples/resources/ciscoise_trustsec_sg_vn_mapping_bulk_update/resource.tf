
resource "ciscoise_trustsec_sg_vn_mapping_bulk_update" "example" {
  provider = ciscoise
  parameters {

    id          = "string"
    last_update = "string"
    sg_name     = "string"
    sgt_id      = "string"
    vn_id       = "string"
    vn_name     = "string"
  }
}

output "ciscoise_trustsec_sg_vn_mapping_bulk_update_example" {
  value = ciscoise_trustsec_sg_vn_mapping_bulk_update.example
}