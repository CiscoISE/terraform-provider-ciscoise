
resource "ciscoise_duo_identity_sync" "example" {
  provider = ciscoise
  item {



  }
  parameters {

    ad_groups {

      name   = "string"
      source = "string"
    }
    configurations = {}
    last_sync      = "string"

    sync_name     = "string"
    sync_schedule = {}
    sync_status   = "string"

  }
}

output "ciscoise_duo_identity_sync_example" {
  value = ciscoise_duo_identity_sync.example
}