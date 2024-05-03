
resource "ciscoise_duo_mfa" "example" {
  provider = ciscoise
  item {



  }
  parameters {

    account_configurations = {}
    connection_name        = "string"
    description            = "string"
    identity_sync          = "string"

    type = "string"

  }
}

output "ciscoise_duo_mfa_example" {
  value = ciscoise_duo_mfa.example
}