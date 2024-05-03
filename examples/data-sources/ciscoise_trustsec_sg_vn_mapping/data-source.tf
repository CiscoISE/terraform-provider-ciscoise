
data "ciscoise_trustsec_sg_vn_mapping" "example" {
  provider    = ciscoise
  filter      = "string"
  filter_type = "string"
  page        = 1
  size        = 1
  sort        = "string"
  sort_by     = "string"
}

output "ciscoise_trustsec_sg_vn_mapping_example" {
  value = data.ciscoise_trustsec_sg_vn_mapping.example.items
}

data "ciscoise_trustsec_sg_vn_mapping" "example" {
  provider = ciscoise
  id       = "string"
}

output "ciscoise_trustsec_sg_vn_mapping_example" {
  value = data.ciscoise_trustsec_sg_vn_mapping.example.item
}
