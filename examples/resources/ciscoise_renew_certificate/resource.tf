
resource "ciscoise_renew_certificate" "example" {
  provider = ciscoise
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    cert_type = "string"
  }
}