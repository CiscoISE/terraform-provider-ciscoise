
data "ciscoise_downloadable_acl" "example" {
  provider = ciscoise
  page     = 1
  size     = 1
}

output "ciscoise_downloadable_acl_example" {
  value = data.ciscoise_downloadable_acl.example.items
}

data "ciscoise_downloadable_acl" "example" {
  provider = ciscoise
  id       = "string"
}

output "ciscoise_downloadable_acl_example" {
  value = data.ciscoise_downloadable_acl.example.item
}
