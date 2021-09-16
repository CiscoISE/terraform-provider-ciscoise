
data "ciscoise_admin_user" "example" {
    provider = ciscoise
    filter = ["string"]
    filter_type = "string"
    page = 1
    size = 1
    sortasc = "string"
    sortdsc = "string"
}

output "ciscoise_admin_user_example" {
    value = data.ciscoise_admin_user.example.items
}

data "ciscoise_admin_user" "example" {
    provider = ciscoise
    id = "string"
}

output "ciscoise_admin_user_example" {
    value = data.ciscoise_admin_user.example.item
}
