
data "ciscoise_device_administration_conditions" "example" {
    provider = ciscoise
    name = "string"
}

output "ciscoise_device_administration_conditions_example" {
    value = data.ciscoise_device_administration_conditions.example.item_name
}

data "ciscoise_device_administration_conditions" "example" {
    provider = ciscoise
    id = "string"
}

output "ciscoise_device_administration_conditions_example" {
    value = data.ciscoise_device_administration_conditions.example.item_id
}

data "ciscoise_device_administration_conditions" "example" {
    provider = ciscoise
}

output "ciscoise_device_administration_conditions_example" {
    value = data.ciscoise_device_administration_conditions.example.items
}
