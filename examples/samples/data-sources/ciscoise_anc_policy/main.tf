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

resource "ciscoise_anc_policy" "example" {
  provider = ciscoise
  parameters {
    name    = "policy1"
    actions = ["QUARANTINE"]
    # actions = ["PORTBOUNCE"]
  }
}

output "ciscoise_anc_policy_example" {
  value = ciscoise_anc_policy.example
}

