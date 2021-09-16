
resource "ciscoise_trusted_certificate" "example" {
    provider = ciscoise
    item {
      
      automatic_crl_update_units = "string"
      crl_distribution_url = "string"
      crl_download_failure_retries_units = "string"
      description = "string"
      id = "string"
      name = "string"
      non_automatic_crl_update_units = "string"
      selected_ocsp_service = "string"
      status = "string"
      trust_for_certificate_based_admin_auth = false
      trust_for_cisco_services_auth = false
      trust_for_client_auth = false
      trust_for_ise_auth = false
    }
}

output "ciscoise_trusted_certificate_example" {
    value = ciscoise_trusted_certificate.example
}