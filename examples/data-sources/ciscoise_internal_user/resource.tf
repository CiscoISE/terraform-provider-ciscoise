
data "ciscoise_internal_user" "example" {
    provider = ciscoise
    name = "string"
}

output "ciscoise_internal_user_example" {
    value = data.ciscoise_internal_user.example.item_name
}

data "ciscoise_internal_user" "example" {
    provider = ciscoise
    id = "string"
}

output "ciscoise_internal_user_example" {
    value = data.ciscoise_internal_user.example.item_id
}

data "ciscoise_internal_user" "example" {
    provider = ciscoise
    filter = ["string"]
    filter_type = "string"
    page = 1
    size = 1
    sortasc = "string"
    sortdsc = "string"
}

output "ciscoise_internal_user_example" {
    value = data.ciscoise_internal_user.example.items
}
