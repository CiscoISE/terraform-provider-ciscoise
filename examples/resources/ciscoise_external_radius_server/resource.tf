
resource "ciscoise_external_radius_server" "example" {
    provider = ciscoise
    item {
      
      accounting_port = 1
      authentication_port = 1
      authenticator_key = "string"
      description = "string"
      enable_key_wrap = false
      encryption_key = "string"
      host_ip = "string"
      id = "string"
      key_input_format = "string"
      name = "string"
      proxy_timeout = 1
      retries = 1
      shared_secret = "string"
      timeout = 1
    }
}

output "ciscoise_external_radius_server_example" {
    value = ciscoise_external_radius_server.example
}