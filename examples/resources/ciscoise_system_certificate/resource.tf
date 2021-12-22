
resource "ciscoise_system_certificate" "example" {
  provider = ciscoise
  parameters {

    admin                                      = "false"
    allow_portal_tag_transfer_for_same_subject = "false"
    allow_replacement_of_portal_group_tag      = "false"
    allow_role_transfer_for_same_subject       = "false"
    description                                = "string"
    eap                                        = "false"
    expiration_ttl_period                      = 1
    expiration_ttl_units                       = "string"
    host_name                                  = "string"
    id                                         = "string"
    ims                                        = "false"
    name                                       = "string"
    portal                                     = "false"
    portal_group_tag                           = "string"
    pxgrid                                     = "false"
    radius                                     = "false"
    renew_self_signed_certificate              = "false"
    saml                                       = "false"
  }
}

output "ciscoise_system_certificate_example" {
  value = ciscoise_system_certificate.example
}