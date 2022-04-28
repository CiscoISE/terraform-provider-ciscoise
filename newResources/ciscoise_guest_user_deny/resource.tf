resource "ciscoise_guest_user_deny" "example" {
  provider = ciscoise
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    id       = "string"
  }
}