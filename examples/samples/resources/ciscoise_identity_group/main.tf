terraform {
  required_providers {
    ciscoise = {
      version = "0.3.0-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}


resource "ciscoise_identity_group" "example" {
  provider = ciscoise
  parameters {
    description = "NewIdGroup"
    name        = "NewGroup"
    parent      = "NAC Group:NAC:IdentityGroups:User Identity Groups"
  }
}

output "ciscoise_identity_group_example" {
  value = ciscoise_identity_group.example
}
