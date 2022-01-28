terraform {
  required_providers {
    ciscoise = {
      version = "0.1.0"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

resource "ciscoise_sgt" "example" {
  provider = ciscoise
  parameters {

    default_sgacls    = []
    description       = "BYOD Security Temp Group 1"
    generation_id     = 0
    is_read_only      = "false"
    name              = "BYOD_Temp"
    propogate_to_apic = "true"
    value             = 17
  }
}

output "ciscoise_sgt_example" {
  value = ciscoise_sgt.example
}