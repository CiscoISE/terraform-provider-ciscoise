
data "ciscoise_device_administration_time_date_conditions" "example" {
  provider = ciscoise
}

output "ciscoise_device_administration_time_date_conditions_example" {
  value = data.ciscoise_device_administration_time_date_conditions.example.items
}

data "ciscoise_device_administration_time_date_conditions" "example" {
  provider = ciscoise
  id       = "string"
}

output "ciscoise_device_administration_time_date_conditions_example" {
  value = data.ciscoise_device_administration_time_date_conditions.example.item
}
