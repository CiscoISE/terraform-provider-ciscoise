
data "ciscoise_active_directories_info" "example" {
  provider = ciscoise
}

output "ciscoise_active_directories_info_example" {
  value = data.ciscoise_active_directories_info.example.items
}
