
resource "ciscoise_pxgrid_service_reregister" "example" {
  provider = ciscoise

  lifecycle {
    create_before_destroy = true
  }
  parameters {

  }
}