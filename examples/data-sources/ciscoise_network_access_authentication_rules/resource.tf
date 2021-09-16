
data "ciscoise_network_access_authentication_rules" "example" {
  provider  = ciscoise
  policy_id = "string"
}

output "ciscoise_network_access_authentication_rules_example" {
  value = data.ciscoise_network_access_authentication_rules.example.items
}

data "ciscoise_network_access_authentication_rules" "example" {
  provider  = ciscoise
  policy_id = "string"
}

output "ciscoise_network_access_authentication_rules_example" {
  value = data.ciscoise_network_access_authentication_rules.example.item
}
