
data "ciscoise_anc_endpoint_bulk_monitor_status" "example" {
  provider = ciscoise
  bulkid   = "string"
}

output "ciscoise_anc_endpoint_bulk_monitor_status_example" {
  value = data.ciscoise_anc_endpoint_bulk_monitor_status.example.item
}
