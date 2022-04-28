
resource "ciscoise_pxgrid_service_unregister" "example" {
  provider = ciscoise
  lifecycle {
    create_before_destroy = true
  }

  parameters{
    
  }
}