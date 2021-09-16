
data "ciscoise_portal_global_setting" "example" {
    provider = ciscoise
    filter = ["string"]
    filter_type = "string"
    page = 1
    size = 1
    sortasc = "string"
    sortdsc = "string"
}

output "ciscoise_portal_global_setting_example" {
    value = data.ciscoise_portal_global_setting.example.items
}

data "ciscoise_portal_global_setting" "example" {
    provider = ciscoise
    id = "string"
}

output "ciscoise_portal_global_setting_example" {
    value = data.ciscoise_portal_global_setting.example.item
}
