terraform {
  required_providers {
    ciscoise = {
      version = "0.1.0-rc"
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
