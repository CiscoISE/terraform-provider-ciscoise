
resource "ciscoise_trustsec_vn_vlan_mapping_bulk_delete" "example" {
  provider = ciscoise
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    payload = ["string"]
  }
}