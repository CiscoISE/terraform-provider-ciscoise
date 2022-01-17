terraform {
  required_providers {
    ciscoise = {
      version = "0.1.0-rc.3"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

data "ciscoise_sgt" "sgt_src" {
  provider = ciscoise
  filter   = ["name.EQ.Guests"]
}
output "ciscoise_sgt_sgt_src_id" {
  value = data.ciscoise_sgt.sgt_src.items[0].id
}

resource "ciscoise_filter_policy" "example" {
  provider = ciscoise
  parameters {
    domains = "default"
    sgt     = data.ciscoise_sgt.sgt_src.items[0].name
    subnet  = "121.11.8.0/22"
  }
}

output "ciscoise_filter_policy_example" {
  value = ciscoise_filter_policy.example
}