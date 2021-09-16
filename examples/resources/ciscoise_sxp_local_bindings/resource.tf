
resource "ciscoise_sxp_local_bindings" "example" {
  provider = ciscoise
  item {

    binding_name       = "string"
    description        = "string"
    id                 = "string"
    ip_address_or_host = "string"
    sgt                = "string"
    sxp_vpn            = "string"
    vns                = "string"
  }
}

output "ciscoise_sxp_local_bindings_example" {
  value = ciscoise_sxp_local_bindings.example
}