
resource "ciscoise_system_certificate_import" "example" {
  provider                                   = ciscoise
  admin                                      = "false"
  allow_extended_validity                    = "false"
  allow_out_of_date_cert                     = "false"
  allow_portal_tag_transfer_for_same_subject = "false"
  allow_replacement_of_certificates          = "false"
  allow_replacement_of_portal_group_tag      = "false"
  allow_role_transfer_for_same_subject       = "false"
  allow_sha1_certificates                    = "false"
  allow_wild_card_certificates               = "false"
  data                                       = "string"
  eap                                        = "false"
  ims                                        = "false"

  name                            = "string"
  password                        = "******"
  portal                          = "false"
  portal_group_tag                = "string"
  private_key_data                = "string"
  pxgrid                          = "false"
  radius                          = "false"
  saml                            = "false"
  validate_certificate_extensions = "false"
}