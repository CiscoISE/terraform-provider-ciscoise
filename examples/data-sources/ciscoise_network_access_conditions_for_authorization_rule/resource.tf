
data "ciscoise_network_access_conditions_for_authorization_rule" "example" {
  provider = ciscoise
}

output "ciscoise_network_access_conditions_for_authorization_rule_example" {
  value = data.ciscoise_network_access_conditions_for_authorization_rule.example.items
}
