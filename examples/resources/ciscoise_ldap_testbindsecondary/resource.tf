
resource "ciscoise_ldap_testbindsecondary" "example" {
    provider = ciscoise
    parameters {
      
      id = "string"
    }
}

output "ciscoise_ldap_testbindsecondary_example" {
    value = ciscoise_ldap_testbindsecondary.example
}