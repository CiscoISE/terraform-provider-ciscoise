
data "ciscoise_anc_endpoint" "example" {
  provider    = ciscoise
  filter      = ["string"]
  filter_type = "string"
  page        = 1
  size        = 1
  sortasc     = "string"
  sortdsc     = "string"
}

output "ciscoise_anc_endpoint_example" {
  value = data.ciscoise_anc_endpoint.example.items
}

data "ciscoise_anc_endpoint" "example" {
  provider = ciscoise
  id       = "string"
}

output "ciscoise_anc_endpoint_example" {
  value = data.ciscoise_anc_endpoint.example.item
}
