
resource "ciscoise_tacacs_server_sequence" "example" {
    provider = ciscoise
    item {
      
      description = "string"
      id = "string"
      local_accounting = false
      name = "string"
      prefix_delimiter = "string"
      prefix_strip = false
      remote_accounting = false
      server_list = "string"
      suffix_delimiter = "string"
      suffix_strip = false
    }
}

output "ciscoise_tacacs_server_sequence_example" {
    value = ciscoise_tacacs_server_sequence.example
}