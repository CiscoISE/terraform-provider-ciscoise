terraform {
  required_providers {
    ciscoise = {
      version = "0.1.0-rc.2"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

data "ciscoise_external_radius_server" "first" {
  provider = ciscoise
  name     = "externalRadiusServer1"
}

output "ciscoise_external_radius_server_first" {
  value = data.ciscoise_external_radius_server.first.item_name
}

resource "ciscoise_radius_server_sequence" "example" {
  provider = ciscoise
  parameters {
    radius_server_list      = ["externalRadiusServer1"]
    description             = "TEST Sequence for ERS1"
    name                    = "TESTSequence"
    strip_prefix            = "false"
    strip_suffix            = "false"
    prefix_separator        = "\\"
    suffix_separator        = "@"
    remote_accounting       = "true"
    local_accounting        = "false"
    use_attr_set_before_acc = "false"
    use_attr_set_on_request = "false"
    continue_authorz_policy = "false"
    before_accept_attr_manipulators_list {
    }
    on_request_attr_manipulator_list {
    }
  }
}

output "ciscoise_radius_server_sequence_example" {
  value = ciscoise_radius_server_sequence.example
}


