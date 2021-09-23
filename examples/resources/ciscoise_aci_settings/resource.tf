
resource "ciscoise_aci_settings" "example" {
  provider = ciscoise
  item {

    aci50                     = "false"
    aci51                     = "false"
    aciipaddress              = "string"
    acipassword               = "******"
    aciuser_name              = "string"
    admin_name                = "string"
    admin_password            = "string"
    all_sxp_domain            = "false"
    default_sgt_name          = "string"
    enable_aci                = "false"
    enable_data_plane         = "false"
    enable_elements_limit     = "false"
    id                        = "string"
    ip_address_host_name      = "string"
    l3_route_network          = "string"
    max_num_iepg_from_aci     = 1
    max_num_sgt_to_aci        = 1
    specific_sxp_domain       = "false"
    specifix_sxp_domain_list  = ["string"]
    suffix_to_epg             = "string"
    suffix_to_sgt             = "string"
    tenant_name               = "string"
    untagged_packet_iepg_name = "string"
  }
}

output "ciscoise_aci_settings_example" {
  value = ciscoise_aci_settings.example
}