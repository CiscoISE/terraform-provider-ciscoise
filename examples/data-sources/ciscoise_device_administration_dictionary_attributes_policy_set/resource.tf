
data "ciscoise_device_administration_dictionary_attributes_policy_set" "example" {
  provider = ciscoise
}

output "ciscoise_device_administration_dictionary_attributes_policy_set_example" {
  value = data.ciscoise_device_administration_dictionary_attributes_policy_set.example.items
}
