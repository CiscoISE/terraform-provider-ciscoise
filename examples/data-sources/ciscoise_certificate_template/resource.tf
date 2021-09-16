
data "ciscoise_certificate_template" "example" {
    provider = ciscoise
    name = "string"
}

output "ciscoise_certificate_template_example" {
    value = data.ciscoise_certificate_template.example.item_name
}

data "ciscoise_certificate_template" "example" {
    provider = ciscoise
    id = "string"
}

output "ciscoise_certificate_template_example" {
    value = data.ciscoise_certificate_template.example.item_id
}

data "ciscoise_certificate_template" "example" {
    provider = ciscoise
    page = 1
    size = 1
}

output "ciscoise_certificate_template_example" {
    value = data.ciscoise_certificate_template.example.items
}
