terraform {
  required_providers {
    ciscoise = {
      version = "0.4.0-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

data "ciscoise_network_access_profiles" "example" {
  provider = ciscoise
}

output "ciscoise_network_access_profiles_example" {
  value = data.ciscoise_network_access_profiles.example.items
}
