
resource "ciscoise_network_access_dictionary" "example" {
    provider = ciscoise
    item {
      
      description = "string"
      dictionary_attr_type = "string"
      id = "string"
      name = "string"
      version = "string"
    }
}

output "ciscoise_network_access_dictionary_example" {
    value = ciscoise_network_access_dictionary.example
}