
data "ciscoise_portal" "example" {
    provider = ciscoise
    filter = ["string"]
    filter_type = "string"
    page = 1
    size = 1
    sortasc = "string"
    sortdsc = "string"
}

output "ciscoise_portal_example" {
    value = data.ciscoise_portal.example.items
}

data "ciscoise_portal" "example" {
    provider = ciscoise
    id = "string"
}

output "ciscoise_portal_example" {
    value = data.ciscoise_portal.example.item
}
