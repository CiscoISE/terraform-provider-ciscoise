
resource "ciscoise_network_access_network_condition" "example" {
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

output "ciscoise_network_access_network_condition_example" {
    value = ciscoise_network_access_network_condition.example
}