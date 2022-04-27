
resource "ciscoise_device_administration_authorization_reset_hitcount" "example" {
  provider  = ciscoise
  lifecycle {
    create_before_destroy = true
  }
  parameters{
    policy_id = "string"
  }

}