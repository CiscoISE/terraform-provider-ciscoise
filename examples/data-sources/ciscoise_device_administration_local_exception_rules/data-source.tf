
data "ciscoise_device_administration_local_exception_rules" "example" {
  provider  = ciscoise
  policy_id = "string"
}

output "ciscoise_device_administration_local_exception_rules_example" {
  value = data.ciscoise_device_administration_local_exception_rules.example.items
}

data "ciscoise_device_administration_local_exception_rules" "example" {
  provider  = ciscoise
  policy_id = "string"
  id        = "string"
}

output "ciscoise_device_administration_local_exception_rules_example" {
  value = data.ciscoise_device_administration_local_exception_rules.example.item
}
