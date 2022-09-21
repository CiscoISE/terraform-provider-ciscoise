terraform {
  required_providers {
    ciscoise = {
      version = "0.6.6-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

data "ciscoise_device_administration_identity_stores" "example" {
  provider = ciscoise
}

output "ciscoise_device_administration_identity_stores_example" {
  value = data.ciscoise_device_administration_identity_stores.example.items
}
