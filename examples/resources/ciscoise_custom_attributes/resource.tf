
resource "ciscoise_custom_attributes" "example" {
  provider = ciscoise
  item {



  }
  parameters {

    attribute_name = "string"
    attribute_type = "string"
    name           = "string"
  }
}

output "ciscoise_custom_attributes_example" {
  value = ciscoise_custom_attributes.example
}