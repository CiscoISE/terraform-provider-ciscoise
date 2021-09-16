
data "ciscoise_tacacs_external_servers" "example" {
  provider = ciscoise
  name     = "string"
}

output "ciscoise_tacacs_external_servers_example" {
  value = data.ciscoise_tacacs_external_servers.example.item_name
}

data "ciscoise_tacacs_external_servers" "example" {
  provider = ciscoise
  id       = "string"
}

output "ciscoise_tacacs_external_servers_example" {
  value = data.ciscoise_tacacs_external_servers.example.item_id
}

data "ciscoise_tacacs_external_servers" "example" {
  provider = ciscoise
  page     = 1
  size     = 1
}

output "ciscoise_tacacs_external_servers_example" {
  value = data.ciscoise_tacacs_external_servers.example.items
}
