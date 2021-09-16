
data "ciscoise_sms_provider" "example" {
  provider    = ciscoise
  filter      = ["string"]
  filter_type = "string"
  page        = 1
  size        = 1
  sortasc     = "string"
  sortdsc     = "string"
}

output "ciscoise_sms_provider_example" {
  value = data.ciscoise_sms_provider.example.items
}
