
data "ciscoise_duo_identity_sync" "example" {
  provider = ciscoise
}

output "ciscoise_duo_identity_sync_example" {
  value = data.ciscoise_duo_identity_sync.example.items
}

data "ciscoise_duo_identity_sync" "example" {
  provider  = ciscoise
  sync_name = "string"
}

output "ciscoise_duo_identity_sync_example" {
  value = data.ciscoise_duo_identity_sync.example.item
}
