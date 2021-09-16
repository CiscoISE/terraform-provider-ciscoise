
resource "ciscoise_device_administration_network_conditions" "example" {
    provider = ciscoise
    item {
      
      condition_type = "string"
      conditions {
        
        cli_dnis_list = ["string"]
        device_group_list = ["string"]
        device_list = ["string"]
        ip_addr_list = ["string"]
        mac_addr_list = ["string"]
      }
      description = "string"
      id = "string"
      name = "string"
    }
}

output "ciscoise_device_administration_network_conditions_example" {
    value = ciscoise_device_administration_network_conditions.example
}