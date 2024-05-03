
resource "ciscoise_trustsec_vn_vlan_mapping_bulk_update" "example" {
  provider = ciscoise
  parameters {

    id              = "string"
    is_data         = "false"
    is_default_vlan = "false"
    last_update     = "string"
    max_value       = 1
    name            = "string"
    vn_id           = "string"
    vn_name         = "string"
  }
}

output "ciscoise_trustsec_vn_vlan_mapping_bulk_update_example" {
  value = ciscoise_trustsec_vn_vlan_mapping_bulk_update.example
}