
data "ciscoise_ldap_name" "example" {
    provider = ciscoise
    name = "string"
}

output "ciscoise_ldap_name_example" {
    value = data.ciscoise_ldap_name.example.item
}
