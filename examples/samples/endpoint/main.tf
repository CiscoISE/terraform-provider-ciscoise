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

data "ciscoise_endpoint_group" "found" {
  provider = ciscoise
  name     = "Sony-Device"
}

output "ciscoise_endpoint_group_found" {
  value = data.ciscoise_endpoint_group.found.item_name[0].id
}

resource "ciscoise_endpoint" "example" {
  provider   = ciscoise
  depends_on = [data.ciscoise_endpoint_group.found]
  parameters {
    name                      = "11:22:33:44:55:66"
    description               = "My Test Endpoint 1"
    mac                       = "11:22:33:44:55:66"
    profile_id                = "67a6ca50-edc9-4236-ada4-225559ed54d6"
    group_id                  = data.ciscoise_endpoint_group.found.item_name[0].id
    static_profile_assignment = "false"
    static_group_assignment   = "false"
    portal_user               = "portalUser"
    identity_store            = "identityStore"
    identity_store_id         = "identityStoreId"
    custom_attributes {
      custom_attributes = {
        Authorization = "Internet"
        Owner         = "wilhelm"
        Department    = "Eng"
        Model         = "1111"
        Manufacturer  = "Cisco"
        iPSK          = "abc123"
        Created       = 1234567890
        Expired       = 2134567890
      }
    }
  }
}

output "ciscoise_endpoint_example" {
  value = ciscoise_endpoint.example
}