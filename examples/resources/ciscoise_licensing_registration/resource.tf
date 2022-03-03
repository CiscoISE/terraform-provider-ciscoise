resource "ciscoise_licensing_registration" "example" {
  provider = ciscoise
  parameters {
    connection_type    = "string" # "HTTP_DIRECT", "PROXY", "SSM_ONPREM_SERVER", "TRANSPORT_GATEWAY"
    registration_type  = "string" # "DEREGISTER", "REGISTER", "RENEW", "UPDATE"
    ssm_on_prem_server = "string"
    tier               = ["string"] # "ADVANTAGE", "DEVICEADMIN", "ESSENTIAL", "PREMIER", "VM"
    token              = "string"
  }
}

output "ciscoise_licensing_registration_example" {
  value = ciscoise_licensing_registration.example
}
