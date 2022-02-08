terraform {
  required_providers {
    ciscoise = {
      version = "0.2.0"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

data "ciscoise_sponsor_group_member" "example" {
  provider = ciscoise

}

output "ciscoise_sponsor_group_member_example" {
  value = data.ciscoise_sponsor_group_member.example.items
}
