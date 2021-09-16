
data "ciscoise_guest_type" "example" {
  provider    = ciscoise
  filter      = ["string"]
  filter_type = "string"
  page        = 1
  size        = 1
  sortasc     = "string"
  sortdsc     = "string"
}

output "ciscoise_guest_type_example" {
  value = data.ciscoise_guest_type.example.items
}

data "ciscoise_guest_type" "example" {
  provider = ciscoise
  id       = "string"
}

output "ciscoise_guest_type_example" {
  value = data.ciscoise_guest_type.example.item
}
