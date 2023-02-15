
terraform {
  required_providers {
    ciscoise = {
      version = "0.6.15-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

resource "ciscoise_selfsigned_certificate_generate" "example" {
  provider = ciscoise
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    admin                                      = "false"
    allow_extended_validity                    = "false"
    allow_portal_tag_transfer_for_same_subject = "false"
    allow_replacement_of_certificates          = "false"
    allow_replacement_of_portal_group_tag      = "false"
    allow_role_transfer_for_same_subject       = "false"
    allow_san_dns_bad_name                     = "false"
    allow_san_dns_non_resolvable               = "false"
    allow_wild_card_certificates               = "false"
    certificate_policies                       = "string"
    digest_type                                = "string"
    eap                                        = "false"
    expiration_ttl                             = 1
    expiration_ttl_unit                        = "string"
    host_name                                  = "string"

    key_length          = "string"
    key_type            = "string"
    name                = "string"
    portal              = "false"
    portal_group_tag    = "string"
    pxgrid              = "false"
    radius              = "false"
    saml                = "false"
    san_dns             = ["string"]
    san_ip              = ["string"]
    san_uri             = ["string"]
    subject_city        = "string"
    subject_common_name = "string"
    subject_country     = "string"
    subject_org         = "string"
    subject_org_unit    = "string"
    subject_state       = "string"
  }
}