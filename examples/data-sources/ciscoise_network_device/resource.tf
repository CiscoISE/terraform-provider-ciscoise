
data "ciscoise_network_device" "example" {
    provider = ciscoise
    name = "string"
}

output "ciscoise_network_device_example" {
    value = data.ciscoise_network_device.example.item_name
}

data "ciscoise_network_device" "example" {
    provider = ciscoise
    id = "string"
}

output "ciscoise_network_device_example" {
    value = data.ciscoise_network_device.example.item_id
}

data "ciscoise_network_device" "example" {
    provider = ciscoise
    filter = ["string"]
    filter_type = "string"
    page = 1
    size = 1
    sortasc = "string"
    sortdsc = "string"
}

output "ciscoise_network_device_example" {
    value = data.ciscoise_network_device.example.items
}
