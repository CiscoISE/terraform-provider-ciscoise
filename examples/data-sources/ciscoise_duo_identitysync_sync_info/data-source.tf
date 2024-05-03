
data "ciscoise_duo_identitysync_sync_info" "example" {
  provider  = ciscoise
  sync_name = "string"
}

output "ciscoise_duo_identitysync_sync_info_example" {
  value = data.ciscoise_duo_identitysync_sync_info.example.item
}
