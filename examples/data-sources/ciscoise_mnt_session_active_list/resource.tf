
data "ciscoise_mnt_session_active_list" "example" {
    provider = ciscoise
}

output "ciscoise_mnt_session_active_list_example" {
    value = data.ciscoise_mnt_session_active_list.example.item
}
