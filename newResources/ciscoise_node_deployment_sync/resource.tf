
resource "ciscoise_node_deployment_sync" "example" {
  provider = ciscoise
  lifecycle {
    create_before_destroy = true
  }
  parameters{
    hostname = "string"
  }

}