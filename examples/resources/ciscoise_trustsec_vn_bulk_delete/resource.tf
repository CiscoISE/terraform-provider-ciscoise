
resource "ciscoise_trustsec_vn_bulk_delete" "example" {
  provider = ciscoise
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    payload = ["string"]
  }
}