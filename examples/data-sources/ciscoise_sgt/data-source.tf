
data "ciscoise_sgt" "example" {
  provider    = ciscoise
  filter      = ["string"]
  filter_type = "string"
  page        = 1
  size        = 1
  sortasc     = "string"
  sortdsc     = "string"
}

output "ciscoise_sgt_example" {
  value = data.ciscoise_sgt.example.items
}

data "ciscoise_sgt" "example" {
  provider = ciscoise
  id       = "string"
}

output "ciscoise_sgt_example" {
  value = data.ciscoise_sgt.example.item
}
