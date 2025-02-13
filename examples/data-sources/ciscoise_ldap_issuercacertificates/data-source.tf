
data "ciscoise_ldap_issuercacertificates" "example" {
    provider = ciscoise
}

output "ciscoise_ldap_issuercacertificates_example" {
    value = data.ciscoise_ldap_issuercacertificates.example.item
}
