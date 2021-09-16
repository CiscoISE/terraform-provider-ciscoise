
data "ciscoise_mnt_session_by_nas_ip" "example" {
  provider = ciscoise
  nas_ipv4 = "string"
}

output "ciscoise_mnt_session_by_nas_ip_example" {
  value = data.ciscoise_mnt_session_by_nas_ip.example.item
}
