terraform {
  required_providers {
    ciscoise = {
      version = "0.5.0-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

resource "ciscoise_trustsec_vn_vlan_mapping_bulk_update" "example" {
  provider = ciscoise
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    payload {
      id              = "92e287ae-3dbc-4d52-ae96-9f43f7025723"
      is_data         = "false"
      is_default_vlan = "false"
      max_value       = 12
      name            = "vlan1_3"
      vn_name         = "vn1"
    }
  }
}