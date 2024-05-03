
data "ciscoise_ad_groups_info" "example" {
  provider         = ciscoise
  active_directory = "string"
}

output "ciscoise_ad_groups_info_example" {
  value = data.ciscoise_ad_groups_info.example.items
}
