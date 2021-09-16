
data "ciscoise_network_access_global_exception_rules" "example" {
    provider = ciscoise
}

output "ciscoise_network_access_global_exception_rules_example" {
    value = data.ciscoise_network_access_global_exception_rules.example.items
}

data "ciscoise_network_access_global_exception_rules" "example" {
    provider = ciscoise
    id = "string"
}

output "ciscoise_network_access_global_exception_rules_example" {
    value = data.ciscoise_network_access_global_exception_rules.example.item
}
