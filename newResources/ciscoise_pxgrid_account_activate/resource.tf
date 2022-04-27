
resource "ciscoise_pxgrid_account_activate" "example" {
  provider    = ciscoise
  lifecycle {
    create_before_destroy = true
  }

  parameters{
    description = "string"
  }
}