
data "ciscoise_network_access_service_name" "example" {
    provider = ciscoise
}

output "ciscoise_network_access_service_name_example" {
    value = data.ciscoise_network_access_service_name.example.items
}
