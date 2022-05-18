terraform {
  required_providers {
    ciscoise = {
      version = "0.6.1-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

data "ciscoise_mnt_account_status" "example" {
  provider = ciscoise
  duration = "3"
  mac      = "2C:54:91:88:C9:E3"
}

output "ciscoise_mnt_account_status_example" {
  value = data.ciscoise_mnt_account_status.example.item
}
