terraform {
  required_providers {
    ciscoise = {
      version = "0.6.17-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}
resource "ciscoise_tacacs_command_sets" "asa_read_only" {
  parameters {
    description      = "ASA Read Only 2"
    name             = "ASA Operator 2"
    permit_unmatched = "false"
    commands {
      command_list {
        grant     = "PERMIT"
        command   = "show"
        arguments = ""
      }
    }
  }
}