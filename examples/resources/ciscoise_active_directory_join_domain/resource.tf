
resource "ciscoise_active_directory_join_domain" "example" {
  provider = ciscoise

  lifecycle {
    create_before_destroy = true
  }
  parameters {
    id = "string"
    additional_data {

      name  = "string"
      value = "string"
    }
  }

}