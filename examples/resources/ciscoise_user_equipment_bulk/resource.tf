
resource "ciscoise_user_equipment_bulk" "example" {
  provider = ciscoise
  parameters {

    description  = "string"
    device_group = "string"
    id           = "string"
    imei         = "string"
  }
}

output "ciscoise_user_equipment_bulk_example" {
  value = ciscoise_user_equipment_bulk.example
}