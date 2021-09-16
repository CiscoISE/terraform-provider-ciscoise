
data "ciscoise_network_access_profiles" "example" {
    provider = ciscoise
}

output "ciscoise_network_access_profiles_example" {
    value = data.ciscoise_network_access_profiles.example.items
}
