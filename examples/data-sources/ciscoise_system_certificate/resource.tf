
data "ciscoise_system_certificate" "example" {
    provider = ciscoise
    filter = ["string"]
    filter_type = "string"
    host_name = "string"
    page = 1
    size = 1
    sort = "string"
    sort_by = "string"
}

output "ciscoise_system_certificate_example" {
    value = data.ciscoise_system_certificate.example.items
}

data "ciscoise_system_certificate" "example" {
    provider = ciscoise
    filter = ["string"]
    filter_type = "string"
    host_name = "string"
    page = 1
    size = 1
    sort = "string"
    sort_by = "string"
}

output "ciscoise_system_certificate_example" {
    value = data.ciscoise_system_certificate.example.item
}
