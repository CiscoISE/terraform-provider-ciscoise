
data "ciscoise_network_device_bulk_monitor_status" "example" {
    provider = ciscoise
    bulkid = "string"
}

output "ciscoise_network_device_bulk_monitor_status_example" {
    value = data.ciscoise_network_device_bulk_monitor_status.example.item
}
