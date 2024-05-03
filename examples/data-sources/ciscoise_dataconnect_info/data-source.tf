
data "ciscoise_dataconnect_info" "example" {
  provider = ciscoise
}

output "ciscoise_dataconnect_info_example" {
  value = data.ciscoise_dataconnect_info.example.item
}
