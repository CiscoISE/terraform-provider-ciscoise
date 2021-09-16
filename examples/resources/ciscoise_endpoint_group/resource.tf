
resource "ciscoise_endpoint_group" "example" {
  provider = ciscoise
  item {

    description    = "string"
    id             = "string"
    name           = "string"
    system_defined = false
  }
}

output "ciscoise_endpoint_group_example" {
  value = ciscoise_endpoint_group.example
}