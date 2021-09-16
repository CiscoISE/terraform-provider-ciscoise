
data "ciscoise_network_access_dictionary" "example" {
  provider = ciscoise
}

output "ciscoise_network_access_dictionary_example" {
  value = data.ciscoise_network_access_dictionary.example.items
}

data "ciscoise_network_access_dictionary" "example" {
  provider = ciscoise
  name     = "string"
}

output "ciscoise_network_access_dictionary_example" {
  value = data.ciscoise_network_access_dictionary.example.item
}
