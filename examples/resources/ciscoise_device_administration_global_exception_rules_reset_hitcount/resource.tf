
resource "ciscoise_device_administration_global_exception_rules_reset_hitcount" "example" {
  provider = ciscoise
  lifecycle {
    create_before_destroy = true
  }

  parameters {

  }
}