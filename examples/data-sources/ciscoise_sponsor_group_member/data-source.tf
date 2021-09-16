
data "ciscoise_sponsor_group_member" "example" {
  provider    = ciscoise
  filter      = ["string"]
  filter_type = "string"
  page        = 1
  size        = 1
  sortasc     = "string"
  sortdsc     = "string"
}

output "ciscoise_sponsor_group_member_example" {
  value = data.ciscoise_sponsor_group_member.example.items
}
