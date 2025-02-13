
data "ciscoise_ldap_hosts" "example" {
    provider = ciscoise
}

output "ciscoise_ldap_hosts_example" {
    value = data.ciscoise_ldap_hosts.example.item
}
