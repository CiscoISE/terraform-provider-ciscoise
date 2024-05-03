
data "ciscoise_endpoints_device_type_info" "example" {
  provider = ciscoise
}

output "ciscoise_endpoints_device_type_info_example" {
  value = data.ciscoise_endpoints_device_type_info.example.items
}
