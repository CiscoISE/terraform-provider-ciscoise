
resource "ciscoise_identity_group" "example" {
  provider = ciscoise
  item {

    description = "string"
    id          = "string"
    name        = "string"
    parent      = "string"
  }
}

output "ciscoise_identity_group_example" {
  value = ciscoise_identity_group.example
}