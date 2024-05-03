
data "ciscoise_user_equipment" "example" {
  provider    = ciscoise
  filter      = "string"
  filter_type = "string"
  page        = 1
  size        = 1
  sort        = "string"
  sort_by     = "string"
}

output "ciscoise_user_equipment_example" {
  value = data.ciscoise_user_equipment.example.items
}

data "ciscoise_user_equipment" "example" {
  provider          = ciscoise
  user_equipment_id = "string"
}

output "ciscoise_user_equipment_example" {
  value = data.ciscoise_user_equipment.example.item
}
