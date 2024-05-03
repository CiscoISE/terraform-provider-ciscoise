
resource "ciscoise_duo_mfa_testconnection" "example" {
  provider = ciscoise
  parameters {

    connection_name = "string"
    ikey            = "string"
    s_key           = "string"
  }
}

output "ciscoise_duo_mfa_testconnection_example" {
  value = ciscoise_duo_mfa_testconnection.example
}