
data "ciscoise_active_directory" "example" {
    provider = ciscoise
    name = "string"
}

output "ciscoise_active_directory_example" {
    value = data.ciscoise_active_directory.example.item_name
}

data "ciscoise_active_directory" "example" {
    provider = ciscoise
    id = "string"
}

output "ciscoise_active_directory_example" {
    value = data.ciscoise_active_directory.example.item_id
}

data "ciscoise_active_directory" "example" {
    provider = ciscoise
    page = 1
    size = 1
}

output "ciscoise_active_directory_example" {
    value = data.ciscoise_active_directory.example.items
}
