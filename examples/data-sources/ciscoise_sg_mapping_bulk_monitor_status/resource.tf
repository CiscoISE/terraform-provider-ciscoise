
data "ciscoise_sg_mapping_bulk_monitor_status" "example" {
  provider = ciscoise
  bulkid   = "string"
}

output "ciscoise_sg_mapping_bulk_monitor_status_example" {
  value = data.ciscoise_sg_mapping_bulk_monitor_status.example.item
}
