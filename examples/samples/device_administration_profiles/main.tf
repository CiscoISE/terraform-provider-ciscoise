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

data "ciscoise_device_administration_profiles" "example" {
  provider = ciscoise
}

output "ciscoise_device_administration_profiles_example" {
  value = data.ciscoise_device_administration_profiles.example.items
}
