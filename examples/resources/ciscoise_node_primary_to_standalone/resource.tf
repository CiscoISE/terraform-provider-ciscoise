
resource "ciscoise_node_primary_to_standalone" "example" {
  provider = ciscoise

  lifecycle {
    create_before_destroy = true
  }

  parameters {

  }
}