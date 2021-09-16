
data "ciscoise_mnt_session_by_mac" "example" {
  provider = ciscoise
  mac      = "string"
}

output "ciscoise_mnt_session_by_mac_example" {
  value = data.ciscoise_mnt_session_by_mac.example.item
}
