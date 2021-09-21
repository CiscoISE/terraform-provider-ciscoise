
resource "ciscoise_tacacs_external_servers" "example" {
  provider = ciscoise
  item {

    connection_port = 1
    description     = "string"
    host_ip         = "string"
    id              = "string"
    name            = "string"
    shared_secret   = "string"
    single_connect  = "false"
    timeout         = 1
  }
}

output "ciscoise_tacacs_external_servers_example" {
  value = ciscoise_tacacs_external_servers.example
}