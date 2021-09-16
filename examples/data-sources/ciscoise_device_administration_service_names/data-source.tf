
data "ciscoise_device_administration_service_names" "example" {
  provider = ciscoise
}

output "ciscoise_device_administration_service_names_example" {
  value = data.ciscoise_device_administration_service_names.example.items
}
