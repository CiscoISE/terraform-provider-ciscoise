
data "ciscoise_transport_gateway_settings" "example" {
  provider = ciscoise
}

output "ciscoise_transport_gateway_settings_example" {
  value = data.ciscoise_transport_gateway_settings.example.item
}
