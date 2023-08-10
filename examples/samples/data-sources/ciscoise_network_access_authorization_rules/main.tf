terraform {
  required_providers {
    ciscoise = {
      version = "0.6.22-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

data "ciscoise_network_access_policy_set" "example" {
  provider = ciscoise
}
data "ciscoise_network_access_authorization_rules" "example" {
  provider  = ciscoise
  policy_id = data.ciscoise_network_access_policy_set.example.items[0].id
}

output "ciscoise_network_access_authorization_rules_example" {
  value = data.ciscoise_network_access_authorization_rules.example.items
}

data "ciscoise_network_access_authorization_rules" "example1" {
  provider  = ciscoise
  policy_id = data.ciscoise_network_access_policy_set.example.items[0].id
  id        = data.ciscoise_network_access_authorization_rules.example.items[0].rule[0].id
}

output "ciscoise_network_access_authorization_rules_example1" {
  value = data.ciscoise_network_access_authorization_rules.example1.item
}
