terraform {
  required_providers {
    ciscoise = {
      version = "0.6.16-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

resource "ciscoise_tacacs_server_sequence" "sample" {
  provider = ciscoise
  parameters {

    description       = "Test_TACACS_Server_1"
    local_accounting  = "false"
    name              = "Test_TACACS_S1"
    prefix_delimiter  = "\\"
    prefix_strip      = "true"
    remote_accounting = "true"
    server_list       = "TestTES"
    suffix_delimiter  = "@"
    suffix_strip      = "true"
  }
}

output "ciscoise_tacacs_server_sequence_sample" {
  value = ciscoise_tacacs_server_sequence.sample
}


data "ciscoise_tacacs_server_sequence" "example" {
  provider = ciscoise
  depends_on = [
    ciscoise_tacacs_server_sequence.sample
  ]
  page = 1
  size = 1
}

output "ciscoise_tacacs_server_sequence_example" {
  value = data.ciscoise_tacacs_server_sequence.example.items
}


data "ciscoise_tacacs_server_sequence" "example_name" {
  provider = ciscoise
  depends_on = [
    ciscoise_tacacs_server_sequence.sample
  ]
  name = data.ciscoise_tacacs_server_sequence.example.items[0].name
}

output "ciscoise_tacacs_server_sequence_example_name" {
  value = data.ciscoise_tacacs_server_sequence.example_name.item_name
}

data "ciscoise_tacacs_server_sequence" "example_id" {
  provider = ciscoise
  depends_on = [
    ciscoise_tacacs_server_sequence.sample
  ]
  id = data.ciscoise_tacacs_server_sequence.example.items[0].id
}

output "ciscoise_tacacs_server_sequence_example_id" {
  value = data.ciscoise_tacacs_server_sequence.example_id.item_id
}
