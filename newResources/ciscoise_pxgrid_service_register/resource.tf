
resource "ciscoise_pxgrid_service_register" "example" {
  provider = ciscoise
  lifecycle {
    create_before_destroy = true
  }
  parameters{
    name = "string"
  }
}