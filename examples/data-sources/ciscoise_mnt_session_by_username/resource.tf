
data "ciscoise_mnt_session_by_username" "example" {
  provider = ciscoise
  username = "string"
}

output "ciscoise_mnt_session_by_username_example" {
  value = data.ciscoise_mnt_session_by_username.example.item
}
