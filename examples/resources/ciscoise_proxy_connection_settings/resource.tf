
resource "ciscoise_proxy_connection_settings" "example" {
  provider = ciscoise
  parameters {

    bypass_hosts      = "string"
    fqdn              = "string"
    password          = "******"
    password_required = "false"
    port              = 1
    user_name         = "string"
  }
}

output "ciscoise_proxy_connection_settings_example" {
  value = ciscoise_proxy_connection_settings.example
}