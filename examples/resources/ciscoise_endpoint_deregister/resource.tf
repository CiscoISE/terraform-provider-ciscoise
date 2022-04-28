resource "ciscoise_endpoint_deregister" "example" {
  provider = ciscoise
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    id = "string"
  }
}