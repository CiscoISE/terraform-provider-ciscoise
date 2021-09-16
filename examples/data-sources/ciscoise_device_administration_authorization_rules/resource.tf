
data "ciscoise_device_administration_authorization_rules" "example" {
    provider = ciscoise
    policy_id = "string"
}

output "ciscoise_device_administration_authorization_rules_example" {
    value = data.ciscoise_device_administration_authorization_rules.example.items
}

data "ciscoise_device_administration_authorization_rules" "example" {
    provider = ciscoise
    policy_id = "string"
}

output "ciscoise_device_administration_authorization_rules_example" {
    value = data.ciscoise_device_administration_authorization_rules.example.item
}
