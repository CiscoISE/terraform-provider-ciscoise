terraform {
  required_providers {
    ciscoise = {
      version = "0.7.0-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}


resource "ciscoise_node_deployment" "example" {
  provider = ciscoise
  parameters {

    allow_cert_import = "false"
    fqdn              = "ise31-2.ise.trappedundrise.com"
    hostname          = "ise31-2"
    password          = "redacted"
    # roles             = ["string"]
    services  = ["Session", "Profiler", "DeviceAdmin"]
    user_name = "admin"
  }
}

output "ciscoise_node_deployment_example" {
  value     = ciscoise_node_deployment.example
  sensitive = true
}