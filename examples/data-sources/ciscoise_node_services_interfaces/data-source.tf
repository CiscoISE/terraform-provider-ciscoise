
data "ciscoise_node_services_interfaces" "example" {
  provider = ciscoise
  hostname = "string"
}

output "ciscoise_node_services_interfaces_example" {
  value = data.ciscoise_node_services_interfaces.example.items
}
