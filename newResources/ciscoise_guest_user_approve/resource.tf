resource "ciscoise_guest_user_approve" "example" {
  provider = ciscoise
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    id       = "1234"
  }
}