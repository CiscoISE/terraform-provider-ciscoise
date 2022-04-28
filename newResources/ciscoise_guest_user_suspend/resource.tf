resource "ciscoise_guest_user_suspend" "example" {
  provider = ciscoise
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    id       = "1"
    #name     = "1"
    additional_data {

      name  = "string"
      value = "string"
    }
  }
}