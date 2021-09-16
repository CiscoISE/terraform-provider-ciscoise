
data "ciscoise_device_administration_policy_set" "example" {
  provider = ciscoise
}

output "ciscoise_device_administration_policy_set_example" {
  value = data.ciscoise_device_administration_policy_set.example.items
}

data "ciscoise_device_administration_policy_set" "example" {
  provider = ciscoise
  id       = "string"
}

output "ciscoise_device_administration_policy_set_example" {
  value = data.ciscoise_device_administration_policy_set.example.item
}
