
resource "ciscoise_certificate_profile" "example" {
    provider = ciscoise
    item {
      
      allowed_as_user_name = false
      certificate_attribute_name = "string"
      description = "string"
      external_identity_store_name = "string"
      id = "string"
      match_mode = "string"
      name = "string"
      username_from = "string"
    }
}

output "ciscoise_certificate_profile_example" {
    value = ciscoise_certificate_profile.example
}