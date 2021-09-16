
data "ciscoise_trusted_certificate_import" "example" {
  provider                               = ciscoise
  allow_basic_constraint_cafalse         = false
  allow_out_of_date_cert                 = false
  allow_sha1_certificates                = false
  data                                   = "string"
  description                            = "string"
  name                                   = "string"
  trust_for_certificate_based_admin_auth = false
  trust_for_cisco_services_auth          = false
  trust_for_client_auth                  = false
  trust_for_ise_auth                     = false
  validate_certificate_extensions        = false
}