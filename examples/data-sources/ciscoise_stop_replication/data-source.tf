
data "ciscoise_stop_replication" "example" {
  provider = ciscoise
}

output "ciscoise_stop_replication_example" {
  value = data.ciscoise_stop_replication.example.item
}
