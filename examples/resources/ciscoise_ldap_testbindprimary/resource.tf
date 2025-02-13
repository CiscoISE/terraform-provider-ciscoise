
resource "ciscoise_ldap_testbindprimary" "example" {
    provider = ciscoise
    parameters {
      
      id = "string"
    }
}

output "ciscoise_ldap_testbindprimary_example" {
    value = ciscoise_ldap_testbindprimary.example
}