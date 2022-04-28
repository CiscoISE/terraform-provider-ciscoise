
resource "ciscoise_device_administration_policy_set_reset_hitcount" "example" {
  provider = ciscoise

  lifecycle {
    create_before_destroy = true
  }

  parameters {

  }
}