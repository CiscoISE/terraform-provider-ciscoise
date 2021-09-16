
data "ciscoise_tacacs_server_sequence" "example" {
  provider = ciscoise
  name     = "string"
}

output "ciscoise_tacacs_server_sequence_example" {
  value = data.ciscoise_tacacs_server_sequence.example.item_name
}

data "ciscoise_tacacs_server_sequence" "example" {
  provider = ciscoise
  id       = "string"
}

output "ciscoise_tacacs_server_sequence_example" {
  value = data.ciscoise_tacacs_server_sequence.example.item_id
}

data "ciscoise_tacacs_server_sequence" "example" {
  provider = ciscoise
  page     = 1
  size     = 1
}

output "ciscoise_tacacs_server_sequence_example" {
  value = data.ciscoise_tacacs_server_sequence.example.items
}
