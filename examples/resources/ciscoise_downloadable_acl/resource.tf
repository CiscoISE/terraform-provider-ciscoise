
resource "ciscoise_downloadable_acl" "example" {
  provider = ciscoise
  parameters {

    dacl        = "string"
    dacl_type   = "string"
    description = "string"
    id          = "string"
    name        = "string"
  }
}

output "ciscoise_downloadable_acl_example" {
  value = ciscoise_downloadable_acl.example
}