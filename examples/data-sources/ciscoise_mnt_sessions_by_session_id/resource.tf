
data "ciscoise_mnt_sessions_by_session_id" "example" {
    provider = ciscoise
    session_id = "string"
}

output "ciscoise_mnt_sessions_by_session_id_example" {
    value = data.ciscoise_mnt_sessions_by_session_id.example.item
}
