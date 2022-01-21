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

resource "ciscoise_device_administration_authentication_rules" "example" {
  provider = ciscoise
  parameters {
    if_auth_fail      = "REJECT"
    if_process_fail   = "DROP"
    if_user_not_found = "REJECT"
    policy_id         = "cb32c3bc-c720-40c3-83e4-8897f9dd6943"
    rule {
      condition {
        attribute_name  = "EapAuthentication"
        attribute_value = "EAP-MSCHAPv2"
        condition_type  = "ConditionReference"
        dictionary_name = "Network Access"
        id              = "c456a490-0429-4fd4-91d7-efd1eb1f855a"
        is_negate       = "false"
        name            = "EAP-MSCHAPv2"
        operator        = "equals"
      }
      default    = "false"
      hit_counts = 0
      name       = "Test1"
      rank       = 0
      state      = "disabled"
    }
  }
}

output "ciscoise_device_administration_authentication_rules_example" {
  value = ciscoise_device_administration_authentication_rules.example
}