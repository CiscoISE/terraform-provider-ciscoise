terraform {
  required_providers {
    ciscoise = {
      version = "0.6.16-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

data "ciscoise_network_access_conditions_for_authentication_rule" "example" {
  provider = ciscoise
}

output "ciscoise_network_access_conditions_for_authentication_rule_example" {
  value = data.ciscoise_network_access_conditions_for_authentication_rule.example.items
}
