
data "ciscoise_licensing_connection_type" "example" {
  provider = ciscoise
}

output "ciscoise_licensing_connection_type_example" {
  value = data.ciscoise_licensing_connection_type.example.item
}
