
terraform {
  required_providers {
    ciscoise = {
      version = "0.8.1-beta "
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

resource "ciscoise_node_deployment_sync" "example" {
  provider = ciscoise
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    hostname = "nad"
  }

}