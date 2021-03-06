
resource "ciscoise_node_group" "example" {
  provider = ciscoise
  parameters {

    description = "string"
    mar_cache {

      query_attempts       = 1
      query_timeout        = 1
      replication_attempts = 1
      replication_timeout  = 1
    }
    name            = "string"
    node_group_name = "string"
  }
}

output "ciscoise_node_group_example" {
  value = ciscoise_node_group.example
}