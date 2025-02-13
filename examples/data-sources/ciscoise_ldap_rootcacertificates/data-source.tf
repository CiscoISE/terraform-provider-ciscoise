
data "ciscoise_ldap_rootcacertificates" "example" {
    provider = ciscoise
}

output "ciscoise_ldap_rootcacertificates_example" {
    value = data.ciscoise_ldap_rootcacertificates.example.item
}
