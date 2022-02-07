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

data "ciscoise_mnt_session_auth_list" "example" {
  provider = ciscoise
}

output "ciscoise_mnt_session_auth_list_example" {
  value = data.ciscoise_mnt_session_auth_list.example.item
}
