
resource "ciscoise_ipsec_enable" "example" {
  provider = ciscoise
  parameters {

    host_name = "string"
    nad_ip    = "string"
  }
}

output "ciscoise_ipsec_enable_example" {
  value = ciscoise_ipsec_enable.example
}