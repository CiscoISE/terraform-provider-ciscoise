
resource "ciscoise_trustsec_vn_vlan_mapping_bulk_delete" "example" {
  provider   = ciscoise
  parameters = ["string"]
}

output "ciscoise_trustsec_vn_vlan_mapping_bulk_delete_example" {
  value = ciscoise_trustsec_vn_vlan_mapping_bulk_delete.example
}