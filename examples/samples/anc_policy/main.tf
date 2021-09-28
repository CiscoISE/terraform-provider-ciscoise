terraform {
  required_providers {
    ciscoise = {
      version = "0.0.2-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

resource "ciscoise_anc_policy" "example" {
  provider = ciscoise
  item {
    name    = "policy1"
    actions = ["QUARANTINE"]
  }
}

output "ciscoise_anc_policy_example" {
  value = ciscoise_anc_policy.example
}

