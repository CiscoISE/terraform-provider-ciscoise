
resource "ciscoise_sg_to_vn_to_vlan_bulk_request" "example" {
  provider = ciscoise

  parameters{
    operation_type      = "string"
    resource_media_type = "string"
  }
  
}