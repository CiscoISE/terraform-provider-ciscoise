
resource "ciscoise_trustsec_vn_bulk_create" "example" {
  provider = ciscoise
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    payload {

      additional_attributes = "string"
      id                    = "string"
      last_update           = "string"
      name                  = "string"
    }
  }
}