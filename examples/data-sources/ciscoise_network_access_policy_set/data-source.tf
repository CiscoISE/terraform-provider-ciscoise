
data "ciscoise_network_access_policy_set" "example" {
  provider = ciscoise
}

output "ciscoise_network_access_policy_set_example" {
  value = data.ciscoise_network_access_policy_set.example.items
}

data "ciscoise_network_access_policy_set" "example" {
  provider = ciscoise
  id       = "string"
}

output "ciscoise_network_access_policy_set_example" {
  value = data.ciscoise_network_access_policy_set.example.item
}
