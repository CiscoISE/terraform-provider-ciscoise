
resource "ciscoise_network_access_local_exception_rules_reset_hitcounts" "example" {
  provider = ciscoise
  lifecycle {
    create_before_destroy = true
  }

  parameters {
    policy_id = "string"
  }
}