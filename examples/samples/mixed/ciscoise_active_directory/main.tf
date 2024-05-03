terraform {
  required_providers {
    ciscoise = {
      version = "0.8.0-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

resource "ciscoise_active_directory" "example" {
  provider = ciscoise
  parameters {
    name        = "cisco.com"
    description = "Cisco Active Directory Test"
    domain      = "cisco.com"
    adgroups {
      groups {
        name = "cisco.com/operators"
        sid  = "S-1-5-32-548"
        type = "GLOBAL"
      }
      # Allows to add groups
      # groups {
      #   name = "cisco.com/administrators"
      #   sid  = "S-1-5-32-549"
      #   type = "GLOBAL"
      # }
      # groups {
      #   name = "cisco.com/users"
      #   sid  = "S-1-5-32-550"
      #   type = "GLOBAL"
      # }
    }
    advanced_settings {
      aging_time = 5
      country    = "cr"
      # auth_protection_type = ""
      department                        = "department"
      email                             = "dvargas@cisco.com"
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

data "ciscoise_active_directory" "single_response" {
  provider = ciscoise
  depends_on = [
    ciscoise_active_directory.example
  ]
  # name     = ciscoise_active_directory.example.item[0].name
  # id     = ciscoise_active_directory.example.item[0].id
}

output "ciscoise_active_directory_single_response" {
  value = data.ciscoise_active_directory.single_response
}
