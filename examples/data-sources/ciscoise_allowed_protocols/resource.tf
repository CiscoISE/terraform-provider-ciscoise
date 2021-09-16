
data "ciscoise_allowed_protocols" "example" {
    provider = ciscoise
    name = "string"
}

output "ciscoise_allowed_protocols_example" {
    value = data.ciscoise_allowed_protocols.example.item_name
}

data "ciscoise_allowed_protocols" "example" {
    provider = ciscoise
    id = "string"
}

output "ciscoise_allowed_protocols_example" {
    value = data.ciscoise_allowed_protocols.example.item_id
}

data "ciscoise_allowed_protocols" "example" {
    provider = ciscoise
    page = 1
    size = 1
}

output "ciscoise_allowed_protocols_example" {
    value = data.ciscoise_allowed_protocols.example.items
}
