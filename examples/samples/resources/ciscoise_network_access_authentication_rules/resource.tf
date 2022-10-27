

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

resource "ciscoise_network_access_authentication_rules" "mm-authc-eaptls" {
  provider = ciscoise
  parameters {
    identity_source_name = "Internal Endpoints"
    if_auth_fail         = "REJECT"
    if_process_fail      = "DROP"
    if_user_not_found    = "REJECT"

    policy_id = "acd4b55d-dca3-4b93-a160-8a2d01669827"
    rule {
      default    = "false"
      name       = "Dot1x EAP-TLS"
      rank       = 0
      state      = "enabled"
      condition {
        condition_type = "ConditionAndBlock"
        is_negate = "false"
      children {
        dictionary_name  = "Radius"
        attribute_name  = "NAS-Port-Type"
        operator = "equals"
        attribute_value = "Ethernet"
        is_negate      = "false"
        condition_type = "ConditionAttributes"
      }
      children {
        dictionary_name  = "Network Access"
        attribute_name  = "EapAuthentication"
        operator = "equals"
        attribute_value = "EAP-TLS"
        is_negate = "false"
        condition_type = "ConditionAttributes"
      }
    }
    }
  }
}
