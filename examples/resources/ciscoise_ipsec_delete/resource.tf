
resource "ciscoise_ipsec_delete" "example" {
  provider = ciscoise
  parameters {

    host_name = "string"
    nad_ip    = "string"
  }
}

output "ciscoise_ipsec_delete_example" {
  value = ciscoise_ipsec_delete.example
}