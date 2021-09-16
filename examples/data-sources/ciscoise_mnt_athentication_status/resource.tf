
data "ciscoise_mnt_athentication_status" "example" {
  provider  = ciscoise
  mac       = "string"
  rec_ord_s = "string"
  sec_ond_s = "string"
}

output "ciscoise_mnt_athentication_status_example" {
  value = data.ciscoise_mnt_athentication_status.example.item
}
