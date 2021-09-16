
data "ciscoise_device_administration_conditions_for_authorization_rule" "example" {
    provider = ciscoise
}

output "ciscoise_device_administration_conditions_for_authorization_rule_example" {
    value = data.ciscoise_device_administration_conditions_for_authorization_rule.example.items
}
