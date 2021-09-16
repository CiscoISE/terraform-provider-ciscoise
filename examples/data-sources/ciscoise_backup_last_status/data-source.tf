
data "ciscoise_backup_last_status" "example" {
  provider = ciscoise
}

output "ciscoise_backup_last_status_example" {
  value = data.ciscoise_backup_last_status.example.item
}
