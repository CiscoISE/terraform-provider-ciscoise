
resource "ciscoise_native_supplicant_profile" "example" {
  provider = ciscoise
  item {

    description = "string"
    id          = "string"
    name        = "string"
    wireless_profiles {

      action_type             = "string"
      allowed_protocol        = "string"
      certificate_template_id = "string"
      previous_ssid           = "string"
      ssid                    = "string"
    }
  }
}

output "ciscoise_native_supplicant_profile_example" {
  value = ciscoise_native_supplicant_profile.example
}