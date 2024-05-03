
resource "ciscoise_custom_attributes_rename" "example" {
  provider = ciscoise
  parameters {

    current_name = "string"
    new_name     = "string"
  }
}

output "ciscoise_custom_attributes_rename_example" {
  value = ciscoise_custom_attributes_rename.example
}