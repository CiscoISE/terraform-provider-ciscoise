
data "ciscoise_guest_location" "example" {
    provider = ciscoise
    filter = ["string"]
    filter_type = "string"
    page = 1
    size = 1
    sortasc = "string"
    sortdsc = "string"
}

output "ciscoise_guest_location_example" {
    value = data.ciscoise_guest_location.example.items
}

data "ciscoise_guest_location" "example" {
    provider = ciscoise
    id = "string"
}

output "ciscoise_guest_location_example" {
    value = data.ciscoise_guest_location.example.item
}
