terraform {
  required_providers {
    ciscoise = {
      version = "0.1.0-rc.4"
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

