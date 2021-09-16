
data "ciscoise_device_administration_command_set" "example" {
  provider = ciscoise
}

output "ciscoise_device_administration_command_set_example" {
  value = data.ciscoise_device_administration_command_set.example.items
}
