
data "ciscoise_ipsec_info" "example" {
  provider    = ciscoise
  filter      = "string"
  filter_type = "string"
  page        = 1
  size        = 1
  sort        = "string"
  sort_by     = "string"
}

output "ciscoise_ipsec_info_example" {
  value = data.ciscoise_ipsec_info.example.items
}

data "ciscoise_ipsec_info" "example" {
  provider  = ciscoise
  host_name = "string"
  nad_ip    = "string"
}

output "ciscoise_ipsec_info_example" {
  value = data.ciscoise_ipsec_info.example.item
}
