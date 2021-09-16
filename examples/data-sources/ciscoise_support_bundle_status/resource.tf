
data "ciscoise_support_bundle_status" "example" {
  provider = ciscoise
  page     = 1
  size     = 1
}

output "ciscoise_support_bundle_status_example" {
  value = data.ciscoise_support_bundle_status.example.items
}

data "ciscoise_support_bundle_status" "example" {
  provider = ciscoise
  id       = "string"
}

output "ciscoise_support_bundle_status_example" {
  value = data.ciscoise_support_bundle_status.example.item
}
