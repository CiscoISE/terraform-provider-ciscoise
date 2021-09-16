
data "ciscoise_network_access_security_groups" "example" {
    provider = ciscoise
}

output "ciscoise_network_access_security_groups_example" {
    value = data.ciscoise_network_access_security_groups.example.items
}
