
resource "ciscoise_device_administration_network_conditions" "example" {
  provider = ciscoise
  parameters {

    condition_type = "string"
    conditions {

      cli_dnis_list     = ["string"]
      condition_type    = "string"
      description       = "string"
      device_group_list = ["string"]
      device_list       = ["string"]
      id                = "string"
      ip_addr_list      = ["string"]

      mac_addr_list = ["string"]
      name          = "string"
    }
    description = "string"
    id          = "string"

    name = "string"
  }
}

output "ciscoise_device_administration_network_conditions_example" {
  value = ciscoise_device_administration_network_conditions.example
}