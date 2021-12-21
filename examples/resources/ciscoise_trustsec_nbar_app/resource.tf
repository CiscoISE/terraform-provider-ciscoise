
resource "ciscoise_trustsec_nbar_app" "example" {
  provider = ciscoise
  parameters {

    ports    = "string"
    protocol = "string"
  }
}

output "ciscoise_trustsec_nbar_app_example" {
  value = ciscoise_trustsec_nbar_app.example
}