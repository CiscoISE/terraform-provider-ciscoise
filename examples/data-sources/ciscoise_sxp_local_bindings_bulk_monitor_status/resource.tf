
data "ciscoise_sxp_local_bindings_bulk_monitor_status" "example" {
    provider = ciscoise
    bulkid = "string"
}

output "ciscoise_sxp_local_bindings_bulk_monitor_status_example" {
    value = data.ciscoise_sxp_local_bindings_bulk_monitor_status.example.item
}
