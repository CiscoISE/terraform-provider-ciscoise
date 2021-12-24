terraform {
  required_providers {
    ciscoise = {
      version = "0.1.0-rc.1"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

data "ciscoise_endpoint_get_rejected_endpoints" "example" {
  provider = ciscoise
}

output "ciscoise_endpoint_get_rejected_endpoints_example" {
  value = data.ciscoise_endpoint_get_rejected_endpoints.example.item
}
