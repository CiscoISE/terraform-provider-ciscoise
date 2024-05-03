
resource "ciscoise_duo_identity_sync_status" "example" {
  provider = ciscoise
  parameters {

    reason    = "string"
    sync_name = "string"
    user      = {}
  }
}

output "ciscoise_duo_identity_sync_status_example" {
  value = ciscoise_duo_identity_sync_status.example
}