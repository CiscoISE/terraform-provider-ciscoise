terraform {
  required_providers {
    ciscoise = {
      version = "0.6.4-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

data "ciscoise_device_administration_command_set" "example" {
  provider = ciscoise
}

output "ciscoise_device_administration_command_set_example" {
  value = data.ciscoise_device_administration_command_set.example.items
}
