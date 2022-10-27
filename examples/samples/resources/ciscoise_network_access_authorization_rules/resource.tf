terraform {
  required_providers {
    ciscoise = {
      version = "0.6.11-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

resource "ciscoise_network_access_authorization_rules" "example" {
  provider = ciscoise
  parameters {
    profile = ["authp_uc6"]
    rule {
      condition {
        condition_type = "ConditionReference"
        name           = "lc_uc6"
        id             = "716b2dfd-a1e0-4d94-9802-eb4d3bc614d3"
        is_negate      = "true"
      }
      default = "true"
      # hit_counts = 1
      # id         = "716b2dfd-a1e0-4d94-9802-eb4d3bc614d3"
      name  = "MyRuleName_1"
      rank  = 1
      state = "enabled"
    }
    security_group = "BYOD"
  }
}