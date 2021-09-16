
data "ciscoise_sxp_vpns" "example" {
  provider    = ciscoise
  filter      = ["string"]
  filter_type = "string"
  page        = 1
  size        = 1
  sortasc     = "string"
  sortdsc     = "string"
}

output "ciscoise_sxp_vpns_example" {
  value = data.ciscoise_sxp_vpns.example.items
}

data "ciscoise_sxp_vpns" "example" {
  provider = ciscoise
  id       = "string"
}

output "ciscoise_sxp_vpns_example" {
  value = data.ciscoise_sxp_vpns.example.item
}
