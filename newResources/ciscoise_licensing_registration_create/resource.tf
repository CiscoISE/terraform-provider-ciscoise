
resource "ciscoise_licensing_registration_create" "example" {
  provider        = ciscoise
  connection_type = "string"

  registration_type  = "string"
  ssm_on_prem_server = "string"
  tier               = ["string"]
  token              = "string"
}