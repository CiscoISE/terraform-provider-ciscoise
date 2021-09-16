
data "ciscoise_network_access_dictionary_attribute" "example" {
    provider = ciscoise
    dictionary_name = "string"
}

output "ciscoise_network_access_dictionary_attribute_example" {
    value = data.ciscoise_network_access_dictionary_attribute.example.items
}

data "ciscoise_network_access_dictionary_attribute" "example" {
    provider = ciscoise
    dictionary_name = "string"
}

output "ciscoise_network_access_dictionary_attribute_example" {
    value = data.ciscoise_network_access_dictionary_attribute.example.item
}
