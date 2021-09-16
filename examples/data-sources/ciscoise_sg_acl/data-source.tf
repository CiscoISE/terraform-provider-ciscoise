
data "ciscoise_sg_acl" "example" {
  provider    = ciscoise
  filter      = ["string"]
  filter_type = "string"
  page        = 1
  size        = 1
  sortasc     = "string"
  sortdsc     = "string"
}

output "ciscoise_sg_acl_example" {
  value = data.ciscoise_sg_acl.example.items
}

data "ciscoise_sg_acl" "example" {
  provider = ciscoise
  id       = "string"
}

output "ciscoise_sg_acl_example" {
  value = data.ciscoise_sg_acl.example.item
}
