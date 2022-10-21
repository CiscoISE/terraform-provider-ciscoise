terraform {
  required_providers {
    ciscoise = {
      version = "0.6.9-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

resource "ciscoise_network_access_authentication_rules" "example" {
  provider = ciscoise
  parameters {
    identity_source_name = "Internal Endpoints"
    if_auth_fail         = "REJECT"
    if_process_fail      = "DROP"
    if_user_not_found    = "REJECT"

    policy_id = "d2da0d1d-5d32-41e2-a88d-7dd2107bf0ca"
    rule {
      condition {
        condition_type = "ConditionReference"
        is_negate      = false
        name           = "Wired_MAB"
        id             = "9aab0da7-e3e3-4cd7-81c2-18c3ebbe6a96"
      }
      default = "false"
      #   id         = "9aab0da7-e3e3-4cd7-81c2-18c3ebbe6a96"
      name  = "auth_wim"
      rank  = 1
      state = "enabled"
    }
  }
}
