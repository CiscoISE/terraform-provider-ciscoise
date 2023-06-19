
resource "ciscoise_endpoint_group" "example" {
  provider = ciscoise
  parameters {

    description    = "string"
    id             = "string"
    name           = "string"
    system_defined = "false"
    parent_id      = "string"
  }
}

output "ciscoise_endpoint_group_example" {
  value = ciscoise_endpoint_group.example
}