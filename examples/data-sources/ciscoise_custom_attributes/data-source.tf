
data "ciscoise_custom_attributes" "example" {
  provider = ciscoise
}

output "ciscoise_custom_attributes_example" {
  value = data.ciscoise_custom_attributes.example.items
}

data "ciscoise_custom_attributes" "example" {
  provider = ciscoise
  name     = "string"
}

output "ciscoise_custom_attributes_example" {
  value = data.ciscoise_custom_attributes.example.item
}
