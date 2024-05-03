
data "ciscoise_endpoints" "example" {
  provider    = ciscoise
  filter      = "string"
  filter_type = "string"
  page        = 1
  size        = 1
  sort        = "string"
  sort_by     = "string"
}

output "ciscoise_endpoints_example" {
  value = data.ciscoise_endpoints.example.items
}

data "ciscoise_endpoints" "example" {
  provider = ciscoise
  value    = "string"
}

output "ciscoise_endpoints_example" {
  value = data.ciscoise_endpoints.example.item
}
