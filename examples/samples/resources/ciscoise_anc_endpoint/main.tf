terraform {
  required_providers {
    ciscoise = {
      version = "0.6.15-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}


resource "ciscoise_anc_endpoint" "example" {
  provider = ciscoise
  parameters {
    # ip_address = ""
    # mac_address = "DEBB.3567.0173"
    mac_address = "D0:27:0F:E6:D7:14"
    policy_name = "Test"
  }
}

output "ciscoise_anc_endpoint_example" {
  value = ciscoise_anc_endpoint.example.item
}
