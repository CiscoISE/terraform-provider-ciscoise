terraform {
  required_providers {
    ciscoise = {
      version = "0.0.3-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

resource "ciscoise_device_administration_global_exception_rules" "example" {
  provider = ciscoise
  parameters {
    commands = ["DenyAllCommands"]
    profile  = "Default Shell Profile"

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
      name       = "Test2"
      rank       = 0
      state      = "enabled"
    }
  }
}

output "ciscoise_device_administration_global_exception_rules_example" {
  value = ciscoise_device_administration_global_exception_rules.example
}
