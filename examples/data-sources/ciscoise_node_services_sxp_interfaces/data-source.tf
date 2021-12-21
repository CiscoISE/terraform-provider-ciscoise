
data "ciscoise_node_services_sxp_interfaces" "example" {
  provider = ciscoise
  hostname = "string"
}

output "ciscoise_node_services_sxp_interfaces_example" {
  value = data.ciscoise_node_services_sxp_interfaces.example.item
}
