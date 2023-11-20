
data "ciscoise_node_deployment" "example" {
  provider    = ciscoise
  filter      = "string"
  filter_type = "string"
}

output "ciscoise_node_deployment_example" {
  value = data.ciscoise_node_deployment.example.items
}

data "ciscoise_node_deployment" "example" {
  provider = ciscoise
  hostname = "string"
}

output "ciscoise_node_deployment_example" {
  value = data.ciscoise_node_deployment.example.item
}
