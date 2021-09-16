
data "ciscoise_telemetry_info" "example" {
  provider    = ciscoise
  filter      = ["string"]
  filter_type = "string"
  page        = 1
  size        = 1
}

output "ciscoise_telemetry_info_example" {
  value = data.ciscoise_telemetry_info.example.items
}

data "ciscoise_telemetry_info" "example" {
  provider = ciscoise
  id       = "string"
}

output "ciscoise_telemetry_info_example" {
  value = data.ciscoise_telemetry_info.example.item
}
