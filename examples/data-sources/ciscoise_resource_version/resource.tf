
data "ciscoise_resource_version" "example" {
    provider = ciscoise
    resource = "string"
}

output "ciscoise_resource_version_example" {
    value = data.ciscoise_resource_version.example.item
}
