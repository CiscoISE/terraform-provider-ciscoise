
data "ciscoise_mnt_session_auth_list" "example" {
  provider = ciscoise
}

output "ciscoise_mnt_session_auth_list_example" {
  value = data.ciscoise_mnt_session_auth_list.example.item
}
