
resource "ciscoise_node_services_sxp_interfaces" "example" {
  provider = ciscoise
  parameters {

    hostname  = "string"
    interface = "string"
  }
}

output "ciscoise_node_services_sxp_interfaces_example" {
  value = ciscoise_node_services_sxp_interfaces.example
}