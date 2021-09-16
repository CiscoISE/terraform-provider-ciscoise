
resource "ciscoise_network_access_dictionary_attribute" "example" {
    provider = ciscoise
    item {
      
      allowed_values {
        
        is_default = false
        key = "string"
        value = "string"
      }
      data_type = "string"
      description = "string"
      dictionary_name = "string"
      direction_type = "string"
      id = "string"
      internal_name = "string"
      name = "string"
    }
}

output "ciscoise_network_access_dictionary_attribute_example" {
    value = ciscoise_network_access_dictionary_attribute.example
}