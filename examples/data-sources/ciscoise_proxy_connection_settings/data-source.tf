
data "ciscoise_proxy_connection_settings" "example" {
  provider = ciscoise
}

output "ciscoise_proxy_connection_settings_example" {
  value = data.ciscoise_proxy_connection_settings.example.item
}
