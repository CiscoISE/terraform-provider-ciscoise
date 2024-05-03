
data "ciscoise_duo_identity_sync_cancel_info" "example" {
  provider  = ciscoise
  sync_name = "string"
}

output "ciscoise_duo_identity_sync_cancel_info_example" {
  value = data.ciscoise_duo_identity_sync_cancel_info.example.item
}
