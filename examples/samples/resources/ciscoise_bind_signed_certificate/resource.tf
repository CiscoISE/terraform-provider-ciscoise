
terraform {
  required_providers {
    ciscoise = {
      version = "0.6.11-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

resource "ciscoise_bind_signed_certificate" "example" {
  provider = ciscoise
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    admin                                 = "false"
    allow_extended_validity               = "false"
    allow_out_of_date_cert                = "false"
    allow_replacement_of_certificates     = "false"
    allow_replacement_of_portal_group_tag = "false"
    data                                  = "string"
    eap                                   = "false"
    //host_name                             = "string"
    id  = "string"
    ims = "false"

    name                            = "string"
    portal                          = "false"
    portal_group_tag                = "string"
    pxgrid                          = "false"
    radius                          = "false"
    saml                            = "false"
    validate_certificate_extensions = "false"
  }
}