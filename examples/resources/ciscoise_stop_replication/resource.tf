
resource "ciscoise_stop_replication" "example" {
  provider = ciscoise
  parameters {

    is_enabled = "false"
  }
}

output "ciscoise_stop_replication_example" {
  value = ciscoise_stop_replication.example
}