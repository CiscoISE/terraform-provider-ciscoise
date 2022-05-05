
terraform {
  required_providers {
    ciscoise = {
      version = "0.4.0-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

resource "ciscoise_trusted_certificate_import" "example" {
  provider = ciscoise
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    allow_basic_constraint_cafalse = "false"
    allow_out_of_date_cert         = "false"
    allow_sha1_certificates        = "false"
    data                           = "string"
    description                    = "string"
    name                                   = "string"
    trust_for_certificate_based_admin_auth = "false"
    trust_for_cisco_services_auth          = "false"
    trust_for_client_auth                  = "false"
    trust_for_ise_auth                     = "false"
    validate_certificate_extensions        = "false"
  }
}

output "ciscoise_trusted_certificate_import_example" {
  value = ciscoise_trusted_certificate_import.example
}