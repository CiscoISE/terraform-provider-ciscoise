
data "ciscoise_network_access_network_condition" "example" {
    provider = ciscoise
}

output "ciscoise_network_access_network_condition_example" {
    value = data.ciscoise_network_access_network_condition.example.items
}

data "ciscoise_network_access_network_condition" "example" {
    provider = ciscoise
    id = "string"
}

output "ciscoise_network_access_network_condition_example" {
    value = data.ciscoise_network_access_network_condition.example.item
}
