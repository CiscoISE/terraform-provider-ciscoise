
data "ciscoise_user_equipment_imei_info" "example" {
  provider = ciscoise
  imei     = "string"
}

output "ciscoise_user_equipment_imei_info_example" {
  value = data.ciscoise_user_equipment_imei_info.example.item
}
