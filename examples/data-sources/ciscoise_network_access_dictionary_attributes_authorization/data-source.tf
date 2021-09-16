
data "ciscoise_network_access_dictionary_attributes_authorization" "example" {
  provider = ciscoise
}

output "ciscoise_network_access_dictionary_attributes_authorization_example" {
  value = data.ciscoise_network_access_dictionary_attributes_authorization.example.items
}
