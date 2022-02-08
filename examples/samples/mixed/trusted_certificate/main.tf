terraform {
  required_providers {
    ciscoise = {
      version = "0.2.0"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

data "ciscoise_trusted_certificate" "example" {
  provider = ciscoise
}

output "ciscoise_trusted_certificate_example" {
  value = data.ciscoise_trusted_certificate.example.items
}

data "ciscoise_trusted_certificate" "example_id" {
  provider = ciscoise
  id       = data.ciscoise_trusted_certificate.example.items[0].id
}

output "ciscoise_trusted_certificate_example_id" {
  value = data.ciscoise_trusted_certificate.example_id.item
}


resource "ciscoise_trusted_certificate" "iden_trust" {
  provider = ciscoise
  parameters {
    # id                                 = "7865ac6a-64c6-4e65-865e-d1b093ee0b10"
    name                               = "IdenTrust Commercial Root CA 1"
    authenticate_before_crl_received   = "false" #"off"
    automatic_crl_update               = "true"  #"on"
    automatic_crl_update_period        = 6
    automatic_crl_update_units         = "Minutes"
    crl_distribution_url               = " "
    crl_download_failure_retries       = 10
    crl_download_failure_retries_units = "Minutes"
    description                        = "IdenTrust Commercial Root CA 1"
    download_crl                       = "false" #"off"
    enable_ocsp_validation             = "false" #"off"
    enable_server_identity_check       = "false" #"off"
    ignore_crl_expiration              = "false" #"off"
    non_automatic_crl_update_period    = 1
    non_automatic_crl_update_units     = "Hours"
    reject_if_no_status_from_ocs_p     = "false" #"off"
    reject_if_unreachable_from_ocs_p   = "false" #"off"
    selected_ocsp_service              = " "
    status                             = "Enabled"
  }
}
