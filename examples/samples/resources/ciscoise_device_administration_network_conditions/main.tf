terraform {
  required_providers {
    ciscoise = {
      version = "0.6.19-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

resource "ciscoise_device_administration_network_conditions" "example" {
  provider = ciscoise
  parameters {

    condition_type = "EndstationCondition"
    conditions {

      cli_dnis_list = ["TBD"]
      ip_addr_list  = ["1.1.1.1", "2.2.2.2"]
      mac_addr_list = [
        "00-0E-A6-A7-63-F7,00-0E-A6-A7-63-F8",
        "00-0E-A6-A7-63-F7,-ANY-",
        "-ANY-,00-0E-A6-A7-63-F8"
      ]
    }
    description = "Optional description for 2"
    name        = "Endstation condition 2"
  }
}

output "ciscoise_device_administration_network_conditions_example" {
  value = ciscoise_device_administration_network_conditions.example
}
