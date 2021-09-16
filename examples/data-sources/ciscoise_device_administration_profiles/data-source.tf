
data "ciscoise_device_administration_profiles" "example" {
  provider = ciscoise
}

output "ciscoise_device_administration_profiles_example" {
  value = data.ciscoise_device_administration_profiles.example.items
}
