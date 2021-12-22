
data "ciscoise_mnt_authentication_status" "example" {
  provider  = ciscoise
  mac       = "string"
  rec_ord_s = "string"
  sec_ond_s = "string"
}

output "ciscoise_mnt_authentication_status_example" {
  value = data.ciscoise_mnt_authentication_status.example.item
}
