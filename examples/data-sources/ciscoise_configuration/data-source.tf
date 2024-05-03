
data "ciscoise_configuration" "example" {
  provider = ciscoise
}

output "ciscoise_configuration_example" {
  value = data.ciscoise_configuration.example.item
}
