terraform {
  required_providers {
    ciscoise = {
      version = "0.3.0-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

data "ciscoise_network_access_policy_set" "example" {
  provider = ciscoise
}


# output "ciscoise_network_access_policy_set_example" {
#   value = data.ciscoise_network_access_policy_set.example.items
# }

data "ciscoise_network_access_authentication_rules" "example" {
  provider  = ciscoise
  policy_id = data.ciscoise_network_access_policy_set.example.items[0].id
}

# output "ciscoise_network_access_authentication_rules_example" {
#   value = data.ciscoise_network_access_authentication_rules.example.items
# }

data "ciscoise_network_access_authentication_rules" "example1" {
  provider  = ciscoise
  policy_id = data.ciscoise_network_access_policy_set.example.items[0].id
  id        = data.ciscoise_network_access_authentication_rules.example.items[0].rule[0].id
}

output "ciscoise_network_access_authentication_rules_example1" {
  value = data.ciscoise_network_access_authentication_rules.example1
}
