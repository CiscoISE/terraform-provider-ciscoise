
data "ciscoise_trusted_certificate" "example" {
    provider = ciscoise
    filter = ["string"]
    filter_type = "string"
    page = 1
    size = 1
    sort = "string"
    sort_by = "string"
}

output "ciscoise_trusted_certificate_example" {
    value = data.ciscoise_trusted_certificate.example.items
}

data "ciscoise_trusted_certificate" "example" {
    provider = ciscoise
    id = "string"
}

output "ciscoise_trusted_certificate_example" {
    value = data.ciscoise_trusted_certificate.example.item
}
