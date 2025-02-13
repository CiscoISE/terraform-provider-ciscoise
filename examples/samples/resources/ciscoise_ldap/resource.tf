terraform {
  required_providers {
    ciscoise = {
      version = "0.8.1-beta "
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

resource "ciscoise_ldap" "example" {
    provider = ciscoise
    parameters {
      name = "Test_LDAP_TF"
    }
}

output "ciscoise_ldap_example" {
    value = ciscoise_ldap.example
}