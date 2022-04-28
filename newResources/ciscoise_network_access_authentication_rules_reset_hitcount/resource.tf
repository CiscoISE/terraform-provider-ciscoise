resource "ciscoise_network_access_authentication_rules_reset_hitcount" "example" {
  provider  = ciscoise
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    policy_id = "string"
  } 
}