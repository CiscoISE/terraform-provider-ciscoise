
data "ciscoise_network_access_conditions" "example" {
  provider = ciscoise
  name     = "string"
}

output "ciscoise_network_access_conditions_example" {
  value = data.ciscoise_network_access_conditions.example.item_name
}

data "ciscoise_network_access_conditions" "example" {
  provider = ciscoise
  id       = "string"
}

output "ciscoise_network_access_conditions_example" {
  value = data.ciscoise_network_access_conditions.example.item_id
}

data "ciscoise_network_access_conditions" "example" {
  provider = ciscoise
}

output "ciscoise_network_access_conditions_example" {
  value = data.ciscoise_network_access_conditions.example.items
}
