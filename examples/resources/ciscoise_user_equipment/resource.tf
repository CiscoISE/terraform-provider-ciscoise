
resource "ciscoise_user_equipment" "example" {
  provider = ciscoise
  item {



  }
  parameters {

    description  = "string"
    device_group = "string"
    imei         = "string"

    user_equipment_id = "string"

  }
}

output "ciscoise_user_equipment_example" {
  value = ciscoise_user_equipment.example
}