
data "ciscoise_ipsec_certificates_info" "example" {
  provider = ciscoise
}

output "ciscoise_ipsec_certificates_info_example" {
  value = data.ciscoise_ipsec_certificates_info.example.items
}
