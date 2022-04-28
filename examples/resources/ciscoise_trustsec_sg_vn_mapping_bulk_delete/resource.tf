
resource "ciscoise_trustsec_sg_vn_mapping_bulk_delete" "example" {
  provider = ciscoise
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    payload = ["string"]
  }
}