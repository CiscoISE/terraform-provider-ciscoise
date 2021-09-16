
data "ciscoise_radius_server_sequence" "example" {
  provider = ciscoise
  page     = 1
  size     = 1
}

output "ciscoise_radius_server_sequence_example" {
  value = data.ciscoise_radius_server_sequence.example.items
}

data "ciscoise_radius_server_sequence" "example" {
  provider = ciscoise
  id       = "string"
}

output "ciscoise_radius_server_sequence_example" {
  value = data.ciscoise_radius_server_sequence.example.item
}
