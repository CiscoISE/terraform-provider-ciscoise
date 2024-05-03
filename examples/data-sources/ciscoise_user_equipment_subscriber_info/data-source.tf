
data "ciscoise_user_equipment_subscriber_info" "example" {
  provider      = ciscoise
  subscriber_id = "string"
}

output "ciscoise_user_equipment_subscriber_info_example" {
  value = data.ciscoise_user_equipment_subscriber_info.example.items
}
