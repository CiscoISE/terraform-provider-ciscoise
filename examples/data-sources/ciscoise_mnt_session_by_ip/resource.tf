
data "ciscoise_mnt_session_by_ip" "example" {
    provider = ciscoise
    endpoint_ipv4 = "string"
}

output "ciscoise_mnt_session_by_ip_example" {
    value = data.ciscoise_mnt_session_by_ip.example.item
}
