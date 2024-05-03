
data "ciscoise_dataconnect_settings" "example" {
  provider = ciscoise
}

output "ciscoise_dataconnect_settings_example" {
  value = data.ciscoise_dataconnect_settings.example.item
}
