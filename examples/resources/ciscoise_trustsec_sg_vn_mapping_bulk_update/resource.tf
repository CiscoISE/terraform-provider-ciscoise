
resource "ciscoise_trustsec_sg_vn_mapping_bulk_update" "example" {
  provider = ciscoise

  lifecycle {
    create_before_destroy = true
  }
  parameters {
    payload {
      id          = "string"
      last_update = "string"
      sg_name     = "string"
      sgt_id      = "string"
      vn_id       = "string"
      vn_name     = "string"
    }
  }
}