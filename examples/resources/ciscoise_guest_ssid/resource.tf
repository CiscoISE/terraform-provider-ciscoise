
resource "ciscoise_guest_ssid" "example" {
  provider = ciscoise
  parameters {

    id   = "string"
    name = "string"
  }
}

output "ciscoise_guest_ssid_example" {
  value = ciscoise_guest_ssid.example
}