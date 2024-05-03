
resource "ciscoise_endpoints_bulk" "example" {
  provider = ciscoise
  item {























  }
  parameters {

    connected_links           = {}
    custom_attributes         = {}
    description               = "string"
    device_type               = "string"
    group_id                  = "string"
    hardware_revision         = "string"
    id                        = "string"
    identity_store            = "string"
    identity_store_id         = "string"
    ip_address                = "string"
    mac                       = "string"
    mdm_attributes            = {}
    name                      = "string"
    portal_user               = "string"
    product_id                = "string"
    profile_id                = "string"
    protocol                  = "string"
    serial_number             = "string"
    software_revision         = "string"
    static_group_assignment   = "false"
    static_profile_assignment = "false"
    value                     = "string"
    vendor                    = "string"
  }
}

output "ciscoise_endpoints_bulk_example" {
  value = ciscoise_endpoints_bulk.example
}