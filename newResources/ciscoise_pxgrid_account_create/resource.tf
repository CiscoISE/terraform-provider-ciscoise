
resource "ciscoise_pxgrid_account_create" "example" {
  provider = ciscoise

  lifecycle {
    create_before_destroy = true
  }
  parameters{
    node_name = "string"
  }
  
}