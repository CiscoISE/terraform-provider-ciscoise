
data "ciscoise_licensing_eval_license" "example" {
  provider = ciscoise
}

output "ciscoise_licensing_eval_license_example" {
  value = data.ciscoise_licensing_eval_license.example.item
}
