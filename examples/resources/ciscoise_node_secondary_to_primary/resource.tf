
resource "ciscoise_node_secondary_to_primary" "example" {
  provider = ciscoise

  lifecycle {
    create_before_destroy = true
  }

  parameters {

  }
}