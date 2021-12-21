
data "ciscoise_licensing_registration" "example" {
  provider = ciscoise
}

output "ciscoise_licensing_registration_example" {
  value = data.ciscoise_licensing_registration.example.item
}
