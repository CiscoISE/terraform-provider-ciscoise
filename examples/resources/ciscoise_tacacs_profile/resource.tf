
resource "ciscoise_tacacs_profile" "example" {
  provider = ciscoise
  parameters {

    description = "string"
    id          = "string"
    name        = "string"
    session_attributes {

      session_attribute_list {

        name  = "string"
        type  = "string"
        value = "string"
      }
    }
  }
}

output "ciscoise_tacacs_profile_example" {
  value = ciscoise_tacacs_profile.example
}