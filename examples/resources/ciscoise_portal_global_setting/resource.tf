
resource "ciscoise_portal_global_setting" "example" {
    provider = ciscoise
    item {
      
      customization = "string"
      id = "string"
    }
}

output "ciscoise_portal_global_setting_example" {
    value = ciscoise_portal_global_setting.example
}