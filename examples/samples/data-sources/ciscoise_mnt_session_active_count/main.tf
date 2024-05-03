terraform {
  required_providers {
    ciscoise = {
      version = "0.8.0-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

data "ciscoise_mnt_session_active_count" "example" {
  provider = ciscoise
}

output "ciscoise_mnt_session_active_count_example" {
  value = data.ciscoise_mnt_session_active_count.example.item
}
