
data "ciscoise_mnt_version" "example" {
    provider = ciscoise
}

output "ciscoise_mnt_version_example" {
    value = data.ciscoise_mnt_version.example.item
}
