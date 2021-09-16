
resource "ciscoise_sxp_vpns" "example" {
  provider = ciscoise
  item {
    id           = "string"
    sxp_vpn_name = "string"
  }
}

output "ciscoise_sxp_vpns_example" {
  value = ciscoise_sxp_vpns.example
}