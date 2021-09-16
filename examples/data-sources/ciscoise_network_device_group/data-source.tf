
data "ciscoise_network_device_group" "example" {
  provider = ciscoise
  name     = "string"
}

output "ciscoise_network_device_group_example" {
  value = data.ciscoise_network_device_group.example.item_name
}

data "ciscoise_network_device_group" "example" {
  provider = ciscoise
  id       = "string"
}

output "ciscoise_network_device_group_example" {
  value = data.ciscoise_network_device_group.example.item_id
}

data "ciscoise_network_device_group" "example" {
  provider    = ciscoise
  filter      = ["string"]
  filter_type = "string"
  page        = 1
  size        = 1
  sortasc     = "string"
  sortdsc     = "string"
}

output "ciscoise_network_device_group_example" {
  value = data.ciscoise_network_device_group.example.items
}
