
data "ciscoise_network_access_identity_stores" "example" {
  provider = ciscoise
}

output "ciscoise_network_access_identity_stores_example" {
  value = data.ciscoise_network_access_identity_stores.example.items
}
