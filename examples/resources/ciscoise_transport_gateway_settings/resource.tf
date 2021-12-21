
resource "ciscoise_transport_gateway_settings" "example" {
  provider = ciscoise
  parameters {

    enable_transport_gateway = "false"
    url                      = "string"
  }
}

output "ciscoise_transport_gateway_settings_example" {
  value = ciscoise_transport_gateway_settings.example
}