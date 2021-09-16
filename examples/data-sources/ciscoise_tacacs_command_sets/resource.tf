
data "ciscoise_tacacs_command_sets" "example" {
    provider = ciscoise
    name = "string"
}

output "ciscoise_tacacs_command_sets_example" {
    value = data.ciscoise_tacacs_command_sets.example.item_name
}

data "ciscoise_tacacs_command_sets" "example" {
    provider = ciscoise
    id = "string"
}

output "ciscoise_tacacs_command_sets_example" {
    value = data.ciscoise_tacacs_command_sets.example.item_id
}

data "ciscoise_tacacs_command_sets" "example" {
    provider = ciscoise
    page = 1
    size = 1
}

output "ciscoise_tacacs_command_sets_example" {
    value = data.ciscoise_tacacs_command_sets.example.items
}
