
data "ciscoise_mnt_account_status" "example" {
    provider = ciscoise
    duration = "string"
    mac = "string"
}

output "ciscoise_mnt_account_status_example" {
    value = data.ciscoise_mnt_account_status.example.item
}
