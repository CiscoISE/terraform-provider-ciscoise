
data "ciscoise_ldap" "example" {
    provider = ciscoise
    page = 1
    size = 1
}

output "ciscoise_ldap_example" {
    value = data.ciscoise_ldap.example.items
}

data "ciscoise_ldap" "example" {
    provider = ciscoise
    id = "string"
}

output "ciscoise_ldap_example" {
    value = data.ciscoise_ldap.example.item
}
