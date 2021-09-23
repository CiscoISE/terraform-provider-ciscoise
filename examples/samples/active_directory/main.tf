terraform {
  required_providers {
    ciscoise = {
      version = "0.0.1-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

data "ciscoise_active_directory" "response" {
  provider = ciscoise
  page     = 1
  size     = 20
}

output "ciscoise_active_directory_response" {
  value = data.ciscoise_active_directory.response
}

# data "ciscoise_active_directory" "single_response" {
#   provider = ciscoise
#   id       = data.ciscoise_active_directory.response.items[0].id
# }

# output "ciscoise_active_directory_single_response" {
#   value = data.ciscoise_active_directory.single_response
# }

resource "ciscoise_active_directory" "example" {
  provider = ciscoise
  item {
    name        = "cisco.com"
    description = "Cisco Active Directory"
    domain      = "cisco.com"
    adgroups {
      groups {
        name = "cisco.com/operators"
        sid  = "S-1-5-32-548"
        type = "GLOBAL"
      }
    }
    advanced_settings {
      aging_time = 5
      country    = "cr"
      # auth_protection_type = ""
      department                        = "department"
      email                             = "mail"
      enable_callback_for_dialin_client = "false"
      enable_pass_change                = "true"
      enable_machine_auth               = "true"
      enable_machine_access             = "true"
      enable_dialin_permission_check    = "false"
      plaintext_auth                    = "false"
      identity_not_in_ad_behaviour      = "SEARCH_JOINED_FOREST"
      unreachable_domains_behaviour     = "PROCEED"
      enable_rewrites                   = "true"
      first_name                        = "givenName"
      last_name                         = "sn"
      organizational_unit               = "company"
      job_title                         = "title"
      locality                          = "l"
      state_or_province                 = "st"
      telephone                         = "00171240593"
      street_address                    = "streetAddress"
      schema                            = "ACTIVE_DIRECTORY"
    }
    ad_attributes {
      attributes {
        name          = "name1"
        type          = "STRING"
        default_value = "defaultString"
        internal_name = "internalName1"
      }
    }
    ad_scopes_names = "Default_Scope"
  }
}

output "ciscoise_active_directory_example" {
  value = ciscoise_active_directory.example
}