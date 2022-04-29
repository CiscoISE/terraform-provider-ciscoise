resource "ciscoise_guest_user_reinstate" "example" {
  provider = ciscoise
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    id   = "string"
    name = "string"
  }
}