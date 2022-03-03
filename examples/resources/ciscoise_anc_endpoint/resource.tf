
resource "ciscoise_anc_endpoint" "example" {
  provider = ciscoise
  parameters {
    id          = "string"
    ip_address  = "string"
    mac_address = "string"
    policy_name = "string"
  }
}

output "ciscoise_anc_endpoint_example" {
  value = ciscoise_anc_endpoint.example.item
}
