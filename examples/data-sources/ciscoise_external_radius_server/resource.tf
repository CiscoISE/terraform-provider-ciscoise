
data "ciscoise_external_radius_server" "example" {
  provider = ciscoise
  name     = "string"
}

output "ciscoise_external_radius_server_example" {
  value = data.ciscoise_external_radius_server.example.item_name
}

data "ciscoise_external_radius_server" "example" {
  provider = ciscoise
  id       = "string"
}

output "ciscoise_external_radius_server_example" {
  value = data.ciscoise_external_radius_server.example.item_id
}

data "ciscoise_external_radius_server" "example" {
  provider = ciscoise
  page     = 1
  size     = 1
}

output "ciscoise_external_radius_server_example" {
  value = data.ciscoise_external_radius_server.example.items
}
