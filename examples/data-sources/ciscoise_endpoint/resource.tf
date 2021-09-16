
data "ciscoise_endpoint" "example" {
    provider = ciscoise
    name = "string"
}

output "ciscoise_endpoint_example" {
    value = data.ciscoise_endpoint.example.item_name
}

data "ciscoise_endpoint" "example" {
    provider = ciscoise
    id = "string"
}

output "ciscoise_endpoint_example" {
    value = data.ciscoise_endpoint.example.item_id
}

data "ciscoise_endpoint" "example" {
    provider = ciscoise
    filter = ["string"]
    filter_type = "string"
    page = 1
    size = 1
    sortasc = "string"
    sortdsc = "string"
}

output "ciscoise_endpoint_example" {
    value = data.ciscoise_endpoint.example.items
}
