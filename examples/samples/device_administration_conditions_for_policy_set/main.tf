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

data "ciscoise_device_administration_conditions_for_policy_set" "example" {
  provider = ciscoise
}

output "ciscoise_device_administration_conditions_for_policy_set_example" {
  value = data.ciscoise_device_administration_conditions_for_policy_set.example.items
}
