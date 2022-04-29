resource "ciscoise_guest_user_suspend" "example" {
  provider = ciscoise
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    id   = "string"
    name = "string"
    additional_data {

      name  = "string"
      value = "string"
    }
  }
}