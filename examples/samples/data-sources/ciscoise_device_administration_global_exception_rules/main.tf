terraform {
  required_providers {
    ciscoise = {
      version = "0.6.20-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

data "ciscoise_device_administration_global_exception_rules" "response" {
  provider = ciscoise
}
output "ciscoise_device_administration_global_exception_rules_response" {
  value = data.ciscoise_device_administration_global_exception_rules.response
}

data "ciscoise_device_administration_global_exception_rules" "single_response" {
  provider = ciscoise
  id       = data.ciscoise_device_administration_global_exception_rules.response.items[0].rule[0].id
}

output "ciscoise_device_administration_global_exception_rules_single_response" {
  value = data.ciscoise_device_administration_global_exception_rules.single_response
}
