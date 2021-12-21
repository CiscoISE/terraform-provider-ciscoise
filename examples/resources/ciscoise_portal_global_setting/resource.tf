
resource "ciscoise_portal_global_setting" "example" {
  provider = ciscoise
  parameters {

    customization = "string"
    id            = "string"
  }
}

output "ciscoise_portal_global_setting_example" {
  value = ciscoise_portal_global_setting.example
}