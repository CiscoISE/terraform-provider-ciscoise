
resource "ciscoise_sxp_connections" "example" {
  provider = ciscoise
  item {

    description = "string"
    enabled     = "false"
    id          = "string"
    ip_address  = "string"
    sxp_mode    = "string"
    sxp_node    = "string"
    sxp_peer    = "string"
    sxp_version = "string"
    sxp_vpn     = "string"
  }
}

output "ciscoise_sxp_connections_example" {
  value = ciscoise_sxp_connections.example
}