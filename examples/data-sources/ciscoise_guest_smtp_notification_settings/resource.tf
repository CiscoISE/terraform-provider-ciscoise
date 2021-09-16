
data "ciscoise_guest_smtp_notification_settings" "example" {
    provider = ciscoise
    filter = ["string"]
    filter_type = "string"
    page = 1
    size = 1
    sortasc = "string"
    sortdsc = "string"
}

output "ciscoise_guest_smtp_notification_settings_example" {
    value = data.ciscoise_guest_smtp_notification_settings.example.items
}

data "ciscoise_guest_smtp_notification_settings" "example" {
    provider = ciscoise
    id = "string"
}

output "ciscoise_guest_smtp_notification_settings_example" {
    value = data.ciscoise_guest_smtp_notification_settings.example.item
}
