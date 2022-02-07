terraform {
  required_providers {
    ciscoise = {
      version = "0.2.0"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

data "ciscoise_network_access_identity_stores" "example" {
  provider = ciscoise
}

output "ciscoise_network_access_identity_stores_example" {
  value = data.ciscoise_network_access_identity_stores.example.items
}
