terraform {
  required_providers {
    ciscoise = {
      version = "0.6.22-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

resource "ciscoise_network_access_authorization_rules_update" "example" {
  provider = ciscoise
  parameters {
    profile = ["Blackhole_Wireless_Access"]
    rule {
      default = "true"
      state   = "enabled"
      name    = "Default"
      rank    = 0
    }
    policy_id = "32966dce-589e-4aa1-a1ec-7ec7936ab3ae"
    id        = "16f69107-0f32-44c9-af90-f26a30ef11d6"
    # security_group = "BYOD"
  }
}