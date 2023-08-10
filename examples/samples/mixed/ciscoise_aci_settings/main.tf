terraform {
  required_providers {
    ciscoise = {
      version = "0.6.22-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

data "ciscoise_aci_settings" "response" {
  provider = ciscoise
}

output "ciscoise_aci_settings_response" {
  value = data.ciscoise_aci_settings.response.item[0]
}

resource "ciscoise_aci_settings" "example" {
  provider = ciscoise
  parameters {
    aci50                     = data.ciscoise_aci_settings.response.item[0].aci50
    aci51                     = data.ciscoise_aci_settings.response.item[0].aci51
    aciipaddress              = data.ciscoise_aci_settings.response.item[0].aciipaddress
    acipassword               = data.ciscoise_aci_settings.response.item[0].acipassword
    aciuser_name              = data.ciscoise_aci_settings.response.item[0].aciuser_name
    admin_name                = data.ciscoise_aci_settings.response.item[0].admin_name
    admin_password            = data.ciscoise_aci_settings.response.item[0].admin_password
    all_sxp_domain            = data.ciscoise_aci_settings.response.item[0].all_sxp_domain
    default_sgt_name          = data.ciscoise_aci_settings.response.item[0].default_sgt_name
    enable_aci                = "true"
    enable_data_plane         = data.ciscoise_aci_settings.response.item[0].enable_data_plane
    enable_elements_limit     = data.ciscoise_aci_settings.response.item[0].enable_elements_limit
    id                        = data.ciscoise_aci_settings.response.item[0].id
    ip_address_host_name      = data.ciscoise_aci_settings.response.item[0].ip_address_host_name
    l3_route_network          = data.ciscoise_aci_settings.response.item[0].l3_route_network
    max_num_iepg_from_aci     = 1000
    max_num_sgt_to_aci        = data.ciscoise_aci_settings.response.item[0].max_num_sgt_to_aci
    specific_sxp_domain       = data.ciscoise_aci_settings.response.item[0].specific_sxp_domain
    specifix_sxp_domain_list  = data.ciscoise_aci_settings.response.item[0].specifix_sxp_domain_list
    suffix_to_epg             = data.ciscoise_aci_settings.response.item[0].suffix_to_epg
    suffix_to_sgt             = data.ciscoise_aci_settings.response.item[0].suffix_to_sgt
    tenant_name               = data.ciscoise_aci_settings.response.item[0].tenant_name
    untagged_packet_iepg_name = data.ciscoise_aci_settings.response.item[0].untagged_packet_iepg_name
  }
}

output "ciscoise_aci_settings_example" {
  value     = ciscoise_aci_settings.example
  sensitive = true
}
