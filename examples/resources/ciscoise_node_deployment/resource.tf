
resource "ciscoise_node_deployment" "example" {
  provider = ciscoise
  parameters {

    allow_cert_import = "false"
    fqdn              = "string"
    hostname          = "string"
    password          = "******"
    roles             = ["string"]
    services          = ["string"]
    user_name         = "string"
  }
}

output "ciscoise_node_deployment_example" {
  value = ciscoise_node_deployment.example
}