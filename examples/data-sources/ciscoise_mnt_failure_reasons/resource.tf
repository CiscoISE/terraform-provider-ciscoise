
data "ciscoise_mnt_failure_reasons" "example" {
    provider = ciscoise
}

output "ciscoise_mnt_failure_reasons_example" {
    value = data.ciscoise_mnt_failure_reasons.example.item
}
