
resource "ciscoise_pxgrid_access_secret" "example" {
  provider = ciscoise
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    peer_node_name = "string"
  }
}