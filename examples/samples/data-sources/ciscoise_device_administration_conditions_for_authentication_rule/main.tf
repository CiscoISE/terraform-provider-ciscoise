terraform {
  required_providers {
    ciscoise = {
      version = "0.6.1-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

data "ciscoise_device_administration_conditions_for_authentication_rule" "example" {
  provider = ciscoise
}

output "ciscoise_device_administration_conditions_for_authentication_rule_example" {
  value = data.ciscoise_device_administration_conditions_for_authentication_rule.example.items
}
