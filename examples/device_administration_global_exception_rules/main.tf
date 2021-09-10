terraform {
  required_providers {
    ciscoise = {
      version = "1.0.0"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

data "ciscoise_device_administration_global_exception_rules" "response" {
  provider = ciscoise
}
output "ciscoise__device_administration_global_exception_rules_response" {
  value = data.ciscoise_device_administration_global_exception_rules.response
}

data "ciscoise_device_administration_global_exception_rules" "single_response" {
  provider = ciscoise
  id       = data.ciscoise_device_administration_global_exception_rules.response.items[0].rule[0].id
}

output "ciscoise__device_administration_global_exception_rules_single_response" {
  value = data.ciscoise_device_administration_global_exception_rules.single_response
}
