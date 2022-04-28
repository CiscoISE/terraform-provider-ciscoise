resource "ciscoise_guest_user_deny" "example" {
  provider = ciscoise
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    id       = "872c4e3f-8b3f-4462-98ba-00519a033cce"
  }
}