
data "ciscoise_guest_ssid" "example" {
  provider    = ciscoise
  filter      = ["string"]
  filter_type = "string"
  page        = 1
  size        = 1
  sortasc     = "string"
  sortdsc     = "string"
}

output "ciscoise_guest_ssid_example" {
  value = data.ciscoise_guest_ssid.example.items
}

data "ciscoise_guest_ssid" "example" {
  provider = ciscoise
  id       = "string"
}

output "ciscoise_guest_ssid_example" {
  value = data.ciscoise_guest_ssid.example.item
}
