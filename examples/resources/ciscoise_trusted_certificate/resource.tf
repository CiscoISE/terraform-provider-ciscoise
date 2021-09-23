
resource "ciscoise_trusted_certificate" "example" {
  provider = ciscoise
  item {

    authenticate_before_crl_received       = "false"
    automatic_crl_update                   = "false"
    automatic_crl_update_period            = 1
    automatic_crl_update_units             = "string"
    crl_distribution_url                   = "string"
    crl_download_failure_retries           = 1
    crl_download_failure_retries_units     = "string"
    description                            = "string"
    download_crl                           = "false"
    enable_ocsp_validation                 = "false"
    enable_server_identity_check           = "false"
    ignore_crl_expiration                  = "false"
    name                                   = "string"
    non_automatic_crl_update_period        = 1
    non_automatic_crl_update_units         = "string"
    reject_if_no_status_from_ocs_p         = "false"
    reject_if_unreachable_from_ocs_p       = "false"
    selected_ocsp_service                  = "string"
    status                                 = "string"
    trust_for_certificate_based_admin_auth = "false"
    trust_for_cisco_services_auth          = "false"
    trust_for_client_auth                  = "false"
    trust_for_ise_auth                     = "false"
  }
}

output "ciscoise_trusted_certificate_example" {
  value = ciscoise_trusted_certificate.example
}