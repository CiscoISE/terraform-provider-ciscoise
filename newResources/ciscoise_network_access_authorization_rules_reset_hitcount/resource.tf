resource "ciscoise_network_access_authorization_rules_reset_hitcount" "example" {
  provider  = ciscoise
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    policy_id = "string"
  }  

}