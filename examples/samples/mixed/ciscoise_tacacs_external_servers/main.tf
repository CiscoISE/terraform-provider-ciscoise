terraform {
  required_providers {
    ciscoise = {
      version = "0.6.3-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}


resource "ciscoise_tacacs_external_servers" "example" {
  provider = ciscoise
  parameters {

    connection_port = 49
    description     = "Test Tacacs External Servers"
    host_ip         = "10.30.10.1"
    # id              = "string"
    name           = "TestTES"
    shared_secret  = "C1sco123!"
    single_connect = "true"
    timeout        = 20
  }
}

output "ciscoise_tacacs_external_servers_example" {
  value = ciscoise_tacacs_external_servers.example
}

data "ciscoise_tacacs_external_servers" "examples" {
  provider = ciscoise
  depends_on = [
    ciscoise_tacacs_external_servers.example
  ]
  page = 1
  size = 1
}

output "ciscoise_tacacs_external_servers_examples" {
  value = data.ciscoise_tacacs_external_servers.examples.items
}

data "ciscoise_tacacs_external_servers" "example1" {
  provider = ciscoise
  depends_on = [
    ciscoise_tacacs_external_servers.example
  ]
  name = data.ciscoise_tacacs_external_servers.examples.items[0].name
}

output "ciscoise_tacacs_external_servers_example1" {
  value = data.ciscoise_tacacs_external_servers.example1.item_name
}

data "ciscoise_tacacs_external_servers" "example2" {
  provider = ciscoise
  depends_on = [
    ciscoise_tacacs_external_servers.example
  ]
  id = data.ciscoise_tacacs_external_servers.examples.items[0].id
}

output "ciscoise_tacacs_external_servers_example2" {
  value = data.ciscoise_tacacs_external_servers.example2.item_id
}
