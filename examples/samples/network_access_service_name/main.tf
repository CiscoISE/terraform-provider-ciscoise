terraform {
  required_providers {
    ciscoise = {
      version = "0.1.0"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

data "ciscoise_network_access_service_name" "example" {
  provider = ciscoise
}

output "ciscoise_network_access_service_name_example" {
  value = data.ciscoise_network_access_service_name.example.items
}
