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

data "ciscoise_system_config_version" "example" {
  provider = ciscoise
}

output "ciscoise_system_config_version_example" {
  value = data.ciscoise_system_config_version.example.item
}
