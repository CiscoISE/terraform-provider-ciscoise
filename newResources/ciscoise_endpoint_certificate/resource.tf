resource "ciscoise_endpoint_certificate" "example" {
  provider           = ciscoise
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    dirpath            = "string"
    cert_template_name = "string"
    certificate_request {

      cn  = "string"
      san = "string"
    }
    format   = "string"
    password = "******"
  }
}