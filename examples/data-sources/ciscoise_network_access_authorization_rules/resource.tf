
data "ciscoise_network_access_authorization_rules" "example" {
  provider  = ciscoise
  policy_id = "string"
}

output "ciscoise_network_access_authorization_rules_example" {
  value = data.ciscoise_network_access_authorization_rules.example.items
}

data "ciscoise_network_access_authorization_rules" "example" {
  provider  = ciscoise
  policy_id = "string"
}

output "ciscoise_network_access_authorization_rules_example" {
  value = data.ciscoise_network_access_authorization_rules.example.item
}
