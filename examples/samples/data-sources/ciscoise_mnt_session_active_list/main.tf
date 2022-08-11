terraform {
  required_providers {
    ciscoise = {
      version = "0.6.5-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

data "ciscoise_mnt_session_active_list" "example" {
  provider = ciscoise
}

output "ciscoise_mnt_session_active_list_example" {
  value = data.ciscoise_mnt_session_active_list.example.item
}
