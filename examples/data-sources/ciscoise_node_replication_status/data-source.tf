
data "ciscoise_node_replication_status" "example" {
  provider = ciscoise
  node     = "string"
}

output "ciscoise_node_replication_status_example" {
  value = data.ciscoise_node_replication_status.example.item
}
