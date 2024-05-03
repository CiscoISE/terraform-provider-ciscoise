
data "ciscoise_trustsec_nbar_app" "example" {
  provider    = ciscoise
  filter      = "string"
  filter_type = "string"
  page        = 1
  size        = 1
  sort        = "string"
  sort_by     = "string"
}

output "ciscoise_trustsec_nbar_app_example" {
  value = data.ciscoise_trustsec_nbar_app.example.items
}

data "ciscoise_trustsec_nbar_app" "example" {
  provider = ciscoise
  id       = "string"
}

output "ciscoise_trustsec_nbar_app_example" {
  value = data.ciscoise_trustsec_nbar_app.example.item
}
