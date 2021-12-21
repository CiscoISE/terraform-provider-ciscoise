
resource "ciscoise_network_device_group" "example" {
  provider = ciscoise
  parameters {

    description = "string"
    id          = "string"
    name        = "string"
    othername   = "string"
  }
}

output "ciscoise_network_device_group_example" {
  value = ciscoise_network_device_group.example
}