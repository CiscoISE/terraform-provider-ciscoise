
resource "ciscoise_network_access_network_condition" "example" {
  provider = ciscoise
  parameters {

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
}

output "ciscoise_network_access_network_condition_example" {
  value = ciscoise_network_access_network_condition.example
}