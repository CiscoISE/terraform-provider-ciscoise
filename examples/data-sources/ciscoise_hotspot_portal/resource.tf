
data "ciscoise_hotspot_portal" "example" {
    provider = ciscoise
    filter = ["string"]
    filter_type = "string"
    page = 1
    size = 1
    sortasc = "string"
    sortdsc = "string"
}

output "ciscoise_hotspot_portal_example" {
    value = data.ciscoise_hotspot_portal.example.items
}

data "ciscoise_hotspot_portal" "example" {
    provider = ciscoise
    id = "string"
}

output "ciscoise_hotspot_portal_example" {
    value = data.ciscoise_hotspot_portal.example.item
}
