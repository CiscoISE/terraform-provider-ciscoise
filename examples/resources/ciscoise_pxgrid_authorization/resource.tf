
resource "ciscoise_pxgrid_authorization" "example" {
  provider = ciscoise
  lifecycle {
    create_before_destroy = true
  }

  parameters {

  }
}