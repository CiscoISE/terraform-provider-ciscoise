
data "ciscoise_device_administration_dictionary_attributes_authentication" "example" {
    provider = ciscoise
}

output "ciscoise_device_administration_dictionary_attributes_authentication_example" {
    value = data.ciscoise_device_administration_dictionary_attributes_authentication.example.items
}
