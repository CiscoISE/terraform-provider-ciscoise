
resource "ciscoise_endpoint_certificate" "example" {
  provider           = ciscoise
  dirpath            = "string"
  cert_template_name = "string"
  certificate_request {

    cn  = "string"
    san = "string"
  }
  format   = "string"
  password = "******"
}