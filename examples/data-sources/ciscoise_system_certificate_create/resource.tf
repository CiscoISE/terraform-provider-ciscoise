
data "ciscoise_system_certificate_create" "example" {
    provider = ciscoise
    ers_local_cert_stub {
      
      allow_wildcard_certs = "string"
      certificate_policies = "string"
      certificate_san_dns = "string"
      certificate_san_ip = "string"
      certificate_san_uri = "string"
      digest = "string"
      ers_subject_stub {
        
        common_name = "string"
        country_name = "string"
        locality_name = "string"
        organization_name = "string"
        organizational_unit_name = "string"
        state_or_province_name = "string"
      }
      expiration_ttl = 1
      friendly_name = "string"
      group_tag_dd = "string"
      key_length = "string"
      key_type = "string"
      saml_certificate = "string"
      selected_expiration_ttl_unit = "string"
      xgrid_certificate = "string"
    }
    node_id = "string"
}