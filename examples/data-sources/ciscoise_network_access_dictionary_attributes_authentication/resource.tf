
data "ciscoise_network_access_dictionary_attributes_authentication" "example" {
    provider = ciscoise
}

output "ciscoise_network_access_dictionary_attributes_authentication_example" {
    value = data.ciscoise_network_access_dictionary_attributes_authentication.example.items
}
