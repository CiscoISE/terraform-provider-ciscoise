
data "ciscoise_device_administration_conditions_for_policy_set" "example" {
  provider = ciscoise
}

output "ciscoise_device_administration_conditions_for_policy_set_example" {
  value = data.ciscoise_device_administration_conditions_for_policy_set.example.items
}
