
resource "ciscoise_user_equipment_csv" "example" {
  provider = ciscoise
}

output "ciscoise_user_equipment_csv_example" {
  value = ciscoise_user_equipment_csv.example
}