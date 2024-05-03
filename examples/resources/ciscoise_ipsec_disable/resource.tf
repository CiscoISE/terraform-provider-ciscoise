
resource "ciscoise_ipsec_disable" "example" {
  provider = ciscoise
  parameters {

    host_name = "string"
    nad_ip    = "string"
  }
}

output "ciscoise_ipsec_disable_example" {
  value = ciscoise_ipsec_disable.example
}